#![feature(slice_split_once)]

use std::sync::{Arc, Mutex};
use std::collections::HashMap;
use std::time::Duration;
use std::net::SocketAddr;

use rand::SeedableRng;
use rand::rngs::StdRng;

use russh::{Channel, ChannelId, CryptoVec, Pty, MethodSet, MethodKind};
use russh::server::{Config, Msg, Session, Auth, Response, Server as _};
use russh::keys::{Algorithm, PublicKey, PrivateKey, Certificate};

#[tokio::main]
async fn main() {
	let config = Arc::new(Config {
		inactivity_timeout: Some(Duration::from_secs(3600)),
		auth_rejection_time: Duration::from_secs(3),
		auth_rejection_time_initial: Some(Duration::from_secs(0)),
		keys: vec![
			PrivateKey::random(&mut StdRng::seed_from_u64(0), Algorithm::Ed25519).unwrap(),
		],
		preferred: Default::default(),
		methods: MethodSet::from([MethodKind::None].as_slice()),
		..Default::default()
	});
	let mut sh = Server {
		clients: Arc::new(Mutex::new(HashMap::new())),
		id: 0,
	};
	sh.run_on_address(config, ("0.0.0.0", 2025)).await.unwrap();
}

#[derive(Clone)]
struct ClientData {
	width: u32,
	height: u32,
	scroll: u32,
	mouse_x: u32,
	mouse_y: u32,
	selection: Option<u32>,
}

#[derive(Clone)]
struct Server {
	clients: Arc<Mutex<HashMap<ChannelId, ClientData>>>,
	id: usize,
}

impl russh::server::Server for Server {
	type Handler = Self;
	fn new_client(&mut self, _: Option<SocketAddr>) -> Self {
		let s = self.clone();
		self.id += 1;
		s
	}
	fn handle_session_error(&mut self, error: <Self::Handler as russh::server::Handler>::Error) {
		eprintln!("Session error: {:#?}", error);
	}
}

impl russh::server::Handler for Server {
	type Error = russh::Error;

	async fn channel_open_session(
		&mut self,
		channel: Channel<Msg>,
		_session: &mut Session,
	) -> Result<bool, Self::Error> {
		self.clients.lock().unwrap().insert(channel.id(), ClientData {
			width: 0,
			height: 0,
			scroll: 0,
			mouse_x: 0,
			mouse_y: 0,
			selection: None,
		});
		Ok(true)
	}

	async fn auth_none(&mut self, _user: &str) -> Result<Auth, Self::Error> {
		Ok(Auth::Accept)
	}

	async fn auth_password(&mut self, _user: &str, _password: &str) -> Result<Auth, Self::Error> {
		Ok(Auth::Accept)
	}

	async fn auth_publickey(
		&mut self,
		_: &str,
		_key: &PublicKey,
	) -> Result<Auth, Self::Error> {
		Ok(Auth::Accept)
	}

	async fn auth_openssh_certificate(
		&mut self,
		_user: &str,
		_certificate: &Certificate,
	) -> Result<Auth, Self::Error> {
		Ok(Auth::Accept)
	}

	async fn auth_keyboard_interactive(
		&mut self,
		_user: &str,
		_submethods: &str,
		_response: Option<Response<'_>>
	) -> Result<Auth, Self::Error> {
		Ok(Auth::Accept)
	}

	async fn pty_request(
		&mut self,
		channel: ChannelId,
		term: &str,
		col_width: u32,
		row_height: u32,
		_pix_width: u32,
		_pix_height: u32,
		_modes: &[(Pty, u32)],
		session: &mut Session,
	) -> Result<(), Self::Error> {
		session.channel_success(channel)?;

		//println!("{term} {col_width} x {row_height}");

		if let Some(client) = self.clients.lock().unwrap().get_mut(&channel) {
			client.width = col_width;
			client.height = row_height;
		}

		let client = self.clients.lock().unwrap().get(&channel).unwrap().clone();
		render(session, channel, client, false)?;

		Ok(())
	}

	async fn window_change_request(
		&mut self,
		channel: ChannelId,
		col_width: u32,
		row_height: u32,
		_pix_width: u32,
		_pix_height: u32,
		session: &mut Session,
	) -> Result<(), Self::Error> {
		session.channel_success(channel)?;

		//println!("{col_width} x {row_height}");

		if let Some(client) = self.clients.lock().unwrap().get_mut(&channel) {
			client.width = col_width;
			client.height = row_height;
		}

		let client = self.clients.lock().unwrap().get(&channel).unwrap().clone();
		render(session, channel, client, false)?;

		Ok(())
	}

	async fn data(
		&mut self,
		channel: ChannelId,
		data: &[u8],
		session: &mut Session,
	) -> Result<(), Self::Error> {
		// Sending Ctrl+C ends the session and disconnects the client
		if data == [3] {
			close(session, channel)?;
			//return Err(russh::Error::Disconnect);
			return Ok(());
		}

		let mut tooltip_only = false;

		if let Some(cmd) = data.strip_prefix(b"\x1b[M") {
			if cmd.len() >= 3 {
				let button = cmd[0] as u32 - 32;
				let x = cmd[1] as u32 - 32;
				let y = cmd[2] as u32 - 32;
				//println!("Mouse: {:?} {} {}", button, x, y);

				if let Some(client) = self.clients.lock().unwrap().get_mut(&channel) {
					client.mouse_x = x;
					client.mouse_y = y;

					if button == 64 {
						// scroll up
						client.scroll = client.scroll.saturating_sub(1);
					} else if button == 65 {
						// scroll down
						client.scroll = std::cmp::min(client.scroll + 1, 19);
					} else {
						tooltip_only = true;
					}
				}
			}
		}

		//println!("Data: {:?}", data);

		let client = self.clients.lock().unwrap().get(&channel).unwrap().clone();;
		render(session, channel, client, tooltip_only)?;

		Ok(())
	}
}

trait Write {
	type Error;

	fn write(&mut self, buf: &[u8]) -> Result<(), Self::Error>;
	fn write_fmt(&mut self, fmt: std::fmt::Arguments) -> Result<(), Self::Error>;
}

impl<W: std::io::Write> Write for W {
	type Error = std::io::Error;

	fn write(&mut self, buf: &[u8]) -> Result<(), Self::Error> {
		self.write_all(buf)
	}

	fn write_fmt(&mut self, fmt: std::fmt::Arguments) -> Result<(), Self::Error> {
		struct Adapter<'a, T: ?Sized + 'a> {
			inner: &'a mut T,
			error: Result<(), std::io::Error>,
		}

		impl<T: std::io::Write + ?Sized> std::fmt::Write for Adapter<'_, T> {
			fn write_str(&mut self, s: &str) -> std::fmt::Result {
				self.inner.write_all(s.as_bytes()).map_err(|e| {
					self.error = Err(e);
					std::fmt::Error
				})
			}
		}

		let mut output = Adapter { inner: self, error: Ok(()) };
		std::fmt::write(&mut output, fmt).map_err(|_| output.error.unwrap_err())
	}
}

struct Terminal<W: Write>(W);

impl<W: Write> Terminal<W> {
	/// CUP
	fn set_cursor_pos(&mut self, row: u32, col: u32) -> Result<(), W::Error> {
		match (row, col) {
			(1, 1) => write!(self.0, "\x1b[H"),
			(_, 1) => write!(self.0, "\x1b[{}H", row),
			(_, _) => write!(self.0, "\x1b[{};{}H", row, col),
		}
	}
}

fn render2() {
	/*struct Logo {
		lines: Vec<&'static str>,
	}

	trait Comp {

	}

	fn center_h(component: impl Comp) -> impl Comp {

	}

	let logo = center_h(Logo { lines: vec![
		"█▀▀▀ █▀▀▀ █▀▀█ █▀▀▄ █▀▀▀ █▀▀▄ █▀▀█ █▀▀█ █▀▀▄ █▀▀▄",
		"▀▀▀█ █    █  █ █▀▀▄ █▀▀▀ █▀▀▄ █  █ █▀▀█ █▀▀▄ █  █",
		"▀▀▀▀ ▀▀▀▀ ▀▀▀▀ ▀  ▀ ▀▀▀▀ ▀▀▀  ▀▀▀▀ ▀  ▀ ▀  ▀ ▀▀▀ ",
	] });

	let eligibility = ButtonSet {
		buttons: vec![
			Button { text: "Eligible", color: Color::Blue },
			Button { text: "Open", color: Color::Green },
			Button { text: "All", color: Color::Black },
		],
	}*/
}

fn render(session: &mut Session, channel: ChannelId, client: ClientData, tooltip_only: bool) -> Result<(), russh::Error> {
	use std::fmt::Write;
	let mut buf = String::new();

	let watsup = json::parse(include_str!("../watsup.json")).unwrap();
	let scoreboard = json::parse(include_str!("../scoreboard.json")).unwrap();

	let scoreboard: Vec<_> = scoreboard.members().filter(|team| {
		let eligible = team["eligibility_class"].as_str().unwrap() == "Eligible";
		eligible
	}).collect();

	macro_rules! tty {
		(bg($r:expr, $g:expr, $b:expr) $($rest:tt)*) => {
			write!(buf, "\x1b[48;2;{};{};{}m", $r, $g, $b).unwrap();
			tty!($($rest)*);
		};
		(fg($r:expr, $g:expr, $b:expr) $($rest:tt)*) => {
			write!(buf, "\x1b[38;2;{};{};{}m", $r, $g, $b).unwrap();
			tty!($($rest)*);
		};
		($s:literal $($rest:tt)*) => {
			write!(buf, $s).unwrap();
			tty!($($rest)*);
		};
		() => {};
	}

	//tty!("\x1b[6n");

	if !tooltip_only {
		tty!("\x1b[?25l"); // hide cursor

		//tty!("\x1b[?1000h");
		tty!("\x1b[?1003h");
		tty!(bg(0,68,84)); // bg color
		tty!("\x1b[2J"); // clear screen (workaround)
		tty!("\x1b[3J"); // clear screen & scrollback
		tty!("\x1b[H"); // reset cursor position

		tty!(fg(255,171,28));
	}


	fn center(buf: &mut String, text: &str, width: u32) {
		for _ in 0..(width as usize).saturating_sub(text.chars().count()) / 2 {
			buf.push(' ');
		}
		buf.push_str(text);
		buf.push_str("\r\n");
	}

	if !tooltip_only {
		tty!("\r\n");

		center(&mut buf, "█▀▀▀ █▀▀▀ █▀▀█ █▀▀▄ █▀▀▀ █▀▀▄ █▀▀█ █▀▀█ █▀▀▄ █▀▀▄", client.width);
		center(&mut buf, "▀▀▀█ █    █  █ █▀▀▄ █▀▀▀ █▀▀▄ █  █ █▀▀█ █▀▀▄ █  █", client.width);
		center(&mut buf, "▀▀▀▀ ▀▀▀▀ ▀▀▀▀ ▀  ▀ ▀▀▀▀ ▀▀▀  ▀▀▀▀ ▀  ▀ ▀  ▀ ▀▀▀ ", client.width);

		tty!(" " fg(60,110,121) "▄▄▄▄▄▄▄▄▄▄▄▄" fg(0,44,54) "▄▄▄▄▄▄▄▄" "▄▄▄▄▄▄▄" "\r\n");
		tty!(" " bg(60,110,121) fg(255,255,255) "  Eligible  " bg(0,44,54) "  Open  " "  All  " bg(0,68,84) "\r\n");
		tty!(" " fg(60,110,121) "▀▀▀▀▀▀▀▀▀▀▀▀" fg(0,44,54) "▀▀▀▀▀▀▀▀" "▀▀▀▀▀▀▀" "\r\n");
	}

	let challenges = &watsup["challenges"];

	struct Category<'a> {
		name: &'a str,
		name_len: usize,
		width: usize,
		chals: Vec<&'a json::JsonValue>,
	}
	let categories = ["hardware", "web", "reversing", "forensics", "crypto", "pwn", "blockchain", "misc"].map(|name| {
		let mut chals: Vec<_> = challenges.members().filter(|chal| {
			chal["categories"].members().find(|c| !matches!(c.as_str().unwrap(), "enkel" | "medium" | "svår")).unwrap() == name
		}).collect();
		chals.sort_by_key(|chal| {
			if chal["categories"].contains("enkel") {
				0
			} else if chal["categories"].contains("medium") {
				1
			} else if chal["categories"].contains("svår") {
				2
			} else {
				panic!()
			}
		});

		let name_len = name.chars().count();
		let width = std::cmp::max(chals.len() * 4, name_len + 2);
		Category {
			name,
			name_len,
			width,
			chals,
		}
	});

	let legend = 2 + 1 + 22 + 1 + 4 + 1;
	let lpad = client.width.saturating_sub(legend + categories.iter().map(|cat| cat.width as u32).sum::<u32>()) / 2;
	let chal_x = (lpad + legend) as usize;

	let mut tooltip = None;
	if client.mouse_y >= 11 && client.mouse_y < 11 + scoreboard.len() as u32 {
		let mut x = chal_x;
		for cat in &categories {
			for (i, chal) in cat.chals.iter().enumerate() {
				let x = 1 + 1 + x + i * 4;
				if client.mouse_x as usize >= x && (client.mouse_x as usize) < x + 4 {
					let title = chal["title"].as_str().unwrap();
					tooltip = Some(title);
				}
			}
			x += cat.width;
		}
	}
	write!(buf, "\x1b[{}H", 8).unwrap(); // set cursor position
	tty!(bg(0,68,84) fg(255,255,255));
	tty!("\x1b[2K"); // clear line
	if let Some(tooltip) = tooltip {
		center(&mut buf, tooltip, client.width);
	} else {
		tty!("\r\n");
	}

	if tooltip_only {
		session.data(channel, CryptoVec::from(buf))?;
		return Ok(());
	}

	for _ in 0..chal_x {
		buf.push(' ');
	}
	for (i, cat) in categories.iter().enumerate() {
		if i % 2 == 0 {
			tty!(fg(0,58,74));
		} else {
			tty!(fg(0,68,84));
		}

		for _ in 0..cat.width {
			buf.push('▄');
		}
	}

	tty!("\r\n");

	for _ in 0..chal_x {
		buf.push(' ');
	}
	tty!(fg(255,255,255));
	for (i, cat) in categories.iter().enumerate() {
		if i % 2 == 0 {
			tty!(bg(0,58,74));
		} else {
			tty!(bg(0,68,84));
		}

		let lpad = (cat.width - cat.name_len) / 2;
		for _ in 0..lpad {
			buf.push(' ');
		}

		write!(buf, "{}", cat.name).unwrap();
		for _ in 0..cat.width - lpad - cat.name_len {
			buf.push(' ');
		}
	}

	tty!("\r\n");
	for _ in 0..lpad {
		buf.push(' ');
	}
	write!(buf, "{:>w$} ", "#", w = 2).unwrap();
	write!(buf, "{:^w$} ", "Team", w = 21).unwrap();
	write!(buf, "{:w$} ", "Score", w = 5).unwrap();

	for (i, cat) in categories.iter().enumerate() {
		if i % 2 == 0 {
			tty!(bg(0,58,74));
		} else {
			tty!(bg(0,68,84));
		}

		for _ in 0..cat.width {
			buf.push(' ');
		}
	}

	let mut x = chal_x;
	for (i, cat) in categories.iter().enumerate() {
		if i % 2 == 0 {
			tty!(bg(0,58,74));
		} else {
			tty!(bg(0,68,84));
		}

		for (i, chal) in cat.chals.iter().enumerate() {
			let title = chal["title"].as_str().unwrap();
			write!(buf, "\x1b[{};{}H{}…", 11, 1 + 1 + x + i * 4, title.chars().next().unwrap()).unwrap();
		}
		x += cat.width;
	}

	for (i, team) in scoreboard.iter().skip(client.scroll as usize).enumerate() {
		let y = 12 + i;
		let rank = client.scroll as usize + i;
		if y as u32 > client.height {
			break;
		}

		let name = team["name"].as_str().unwrap();
		let score = team["score"].as_u32().unwrap();

		tty!("\r\n");

		for _ in 0..lpad {
			buf.push(' ');
		}

		/*let rank_s = match rank {
			0 => "🥇".to_owned(),
			1 => "🥈".to_owned(),
			2 => "🥉".to_owned(),
			_ => (rank + 1).to_string(),
		};*/
		let rank_s = (rank + 1).to_string();

		write!(buf, "{rank_s:>w$} ", w = 2).unwrap();
		write!(buf, "{name:^w$} ", w = 22).unwrap();
		tty!(fg(128,160,160));
		write!(buf, "{score:w$} ", w = 4).unwrap();

		for (i, cat) in categories.iter().enumerate() {
			if i % 2 == 0 {
				tty!(bg(0,58,74));
			} else {
				tty!(bg(0,68,84));
			}

			for _ in 0..cat.width {
				buf.push(' ');
			}
		}
		let mut x = chal_x;
		for (i, cat) in categories.iter().enumerate() {
			if i % 2 == 0 {
				tty!(bg(0,58,74));
			} else {
				tty!(bg(0,68,84));
			}

			for (i, chal) in cat.chals.iter().enumerate() {
				write!(buf, "\x1b[{};{}H", y, 1 + 1 + x + i * 4).unwrap();
				if team["solves"].has_key(&chal["id"].as_u32().unwrap().to_string()) {
					tty!(fg(255,255,255) "🏳️");
				} else {
					tty!(fg(0,44,54) "●");
				}
			}
			x += cat.width;
		}
		tty!(fg(255,255,255));
	}

	session.data(channel, CryptoVec::from(buf))?;
	Ok(())
}

fn close(session: &mut Session, channel: ChannelId) -> Result<(), russh::Error> {
	session.data(channel, CryptoVec::from("\x1b[?25h\x1b[?1000l\x1b[0m\x1b[2J\x1b[H"))?;
	session.close(channel)?;
	Ok(())
}
