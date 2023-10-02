const VIEW_SIZE = 6 * 2 + 1;
const BLOCK_SIZE = 16 * 4;

class Game {
	constructor() {
		this.element = document.createElement("div");
		this.element.className = "game";
		
		this.canvas = document.createElement("canvas");
		this.canvas.width = VIEW_SIZE * BLOCK_SIZE * window.devicePixelRatio;
		this.canvas.height = VIEW_SIZE * BLOCK_SIZE * window.devicePixelRatio;
		this.canvas.style.width = VIEW_SIZE * BLOCK_SIZE + "px";
		this.canvas.style.height = VIEW_SIZE * BLOCK_SIZE + "px";
		this.ctx = this.canvas.getContext("2d");
		
		this.element.appendChild(this.canvas);
		
		this.map = null;
		this.entities = [];
		
		window.addEventListener("keydown", event => {
			if (!this.ws) return;
			
			if (event.code === "KeyW" || event.code === "ArrowUp") {
				this.ws.send(JSON.stringify({type: "move", x: "0", y: "1"}));
			} else if (event.code === "KeyS" || event.code === "ArrowDown") {
				this.ws.send(JSON.stringify({type: "move", x: "0", y: "-1"}));
			} else if (event.code === "KeyA" || event.code === "ArrowLeft") {
				this.ws.send(JSON.stringify({type: "move", x: "-1", y: "0"}));
			} else if (event.code === "KeyD" || event.code === "ArrowRight") {
				this.ws.send(JSON.stringify({type: "move", x: "1", y: "0"}));
			}
		});
		
		this.canvas.addEventListener("click", event => {
			let x = Math.floor(event.offsetX / BLOCK_SIZE) - 6;
			let y = 6 - Math.floor(event.offsetY / BLOCK_SIZE);
			
			this.ws.send(JSON.stringify({type: "interact", x: x.toString(), y: y.toString(), slot: "-1"}));
		});
		
		this.connect();
		this.render();
	}
	
	connect() {
		this.ws = new WebSocket("wss://daydun.com:666/");
		this.ws.addEventListener("open", () => {});
		this.ws.addEventListener("message", event => {
			let data = JSON.parse(event.data);
			
			switch (data.type) {
				case "connect": {
					if (data.version !== 1) {
						alert("Client out of date.");
						return;
					}
					
					let name = prompt("Enter username:");
					this.ws.send(JSON.stringify({type: "connect", name}));
					
					break;
				}
				//case "move":
				case "tick": {
					this.map = data.map;
					this.entities = data.entities;
					break;
				}
			}
		});
		this.ws.addEventListener("close", () => {
			alert("Server connection closed.");
		});
	}
	
	render() {
		let ctx = this.ctx;
		ctx.clearRect(0, 0, this.canvas.width, this.canvas.height);
		
		if (this.map !== null) {
			ctx.save();
			ctx.translate(this.canvas.width / 2, this.canvas.height / 2);
			ctx.scale(window.devicePixelRatio, window.devicePixelRatio);
			ctx.scale(BLOCK_SIZE, BLOCK_SIZE);
			
			ctx.translate(-0.5, -0.5);
			
			for (let y = 6; y >= -6; y--) {
				for (let x = -6; x <= 6; x++) {
					let block = this.map[14 - (y + 7)][x + 7];
					
					if (block.type === "air") {
						ctx.strokeStyle = "rgba(0, 0, 0, 0.1)";
						ctx.lineWidth = 0.1;
						ctx.strokeRect(x + 0.05, y + 0.05, 0.9, 0.9);
					} else if (block.type === "dirt") {
						ctx.fillStyle = "#674c39";
						ctx.fillRect(x, y, 1, 1);
						
						ctx.strokeStyle = "rgba(0, 0, 0, 0.2)";
						ctx.lineWidth = 0.1;
						ctx.strokeRect(x + 0.05, y + 0.05, 0.9, 0.9);
					} else if (block.type === "wood") {
						ctx.fillStyle = "#937240";
						ctx.fillRect(x, y, 1, 1);
						
						ctx.strokeStyle = "#4a3c25";
						ctx.lineWidth = 0.1;
						ctx.strokeRect(x + 0.05, y + 0.05, 0.9, 0.9);
					} else if (block.type === "concrete") {
						ctx.fillStyle = "#888";
						ctx.fillRect(x, y, 1, 1);
						
						ctx.strokeStyle = "rgba(0, 0, 0, 0.2)";
						ctx.lineWidth = 0.1;
						ctx.strokeRect(x + 0.05, y + 0.05, 0.9, 0.9);
					} else if (block.type === "tombstone") {
						ctx.strokeStyle = "rgba(0, 0, 0, 0.1)";
						ctx.lineWidth = 0.1;
						ctx.strokeRect(x + 0.05, y + 0.05, 0.9, 0.9);
						
						ctx.fillStyle = "#674c39";
						ctx.strokeStyle = "#423125";
						ctx.lineWidth = 0.1;
						ctx.beginPath();
						ctx.rect(x + 0.3, y + 0.2, 0.4, 0.6);
						ctx.fill();
						ctx.stroke();
						
						ctx.fillStyle = "#aaa";
						ctx.strokeStyle = "#777";
						ctx.beginPath();
						ctx.rect(x + 0.3, y + 0.1, 0.4, 0.3);
						ctx.fill();
						ctx.stroke();
					} else {
						ctx.fillStyle = "#222";
						ctx.fillRect(x, y, 1, 1);
						ctx.textAlign = "center";
						ctx.fillStyle = "#fff";
						ctx.font = "bold 0.6px Arial";
						ctx.fillText("?", x + 0.5, y + 0.7);
						
						ctx.strokeStyle = "rgba(255, 255, 255, 0.1)";
						ctx.lineWidth = 0.1;
						ctx.strokeRect(x + 0.05, y + 0.05, 0.9, 0.9);
					}
				}
			}
			
			for (let entity of this.entities) {
				ctx.save();
				ctx.translate(parseInt(entity.x), -parseInt(entity.y));
				ctx.translate(0.5, 0.5);
				if (entity.type === "player") {
					let seed = 1234;
					for (let c of entity.name) {
						seed *= c.codePointAt();
						seed %= 2**16;
					}
					
					function rand() {
						seed = (seed ** 2) % 2**16;
						return seed / 2**16;
					}
					
					let r1 = rand();
					let r2 = (r1 + 0.2 + rand() * 0.6) % 1;
					let r3 = Math.min(Math.abs(r1 - r2), Math.abs(r1 + 1 - r2)) < 0.5 ? (r1 + r2) / 2 + 0.5 : (r1 + r2) / 2;
					
					const f = r => [Math.cos(r * 2 * Math.PI) * 0.4, Math.sin(r * 2 * Math.PI) * 0.4];
					this.p1 = f(r1);
					this.p2 = f(r2);
					this.p3 = f(r3);
					
					let c = rand();
					this.color = [
						Math.sin(c * 2 * Math.PI + 0/2 * Math.PI) * 127 + 128,
						Math.sin(c * 2 * Math.PI + 2/3 * Math.PI) * 127 + 128,
						Math.sin(c * 2 * Math.PI + 4/3 * Math.PI) * 127 + 128
					];
					
					ctx.beginPath();
					ctx.moveTo(...this.p1);
					ctx.lineTo(...this.p2);
					ctx.lineTo(...this.p3);
					ctx.closePath();
					ctx.fillStyle = `rgb(${this.color.join(",")})`;
					ctx.fill();
					ctx.lineWidth = 0.1;
					ctx.strokeStyle = "#000";
					ctx.stroke();
				} else if (entity.type === "ghost") {
					const GHOST_WIDTH = 0.6;
					const GHOST_HEIGHT = 0.8;
					const WAVES = 2.5;
					const WAVE_RADIUS = GHOST_WIDTH / WAVES / 2;
					
					ctx.lineWidth = 0.15;
					ctx.beginPath();
					ctx.moveTo(-GHOST_WIDTH / 2, GHOST_HEIGHT / 2 - WAVE_RADIUS / 2);
					ctx.lineTo(-GHOST_WIDTH / 2, -GHOST_HEIGHT / 2 + GHOST_WIDTH / 2);
					ctx.arc(0, -GHOST_HEIGHT / 2 + GHOST_WIDTH / 2, GHOST_WIDTH / 2, Math.PI, 0);
					ctx.lineTo(GHOST_WIDTH / 2, GHOST_HEIGHT / 2 - WAVE_RADIUS / 2);
					for (let i = 0; i < WAVES * 2; i++) {
						if (i % 2 === 0) {
							ctx.arcTo(GHOST_WIDTH / 2 - WAVE_RADIUS * i      , GHOST_HEIGHT / 2 - WAVE_RADIUS * 0, GHOST_WIDTH / 2 - WAVE_RADIUS * (i * 2 + 1) / 2, GHOST_HEIGHT / 2 - WAVE_RADIUS * 0.0, WAVE_RADIUS / 2);
							ctx.arcTo(GHOST_WIDTH / 2 - WAVE_RADIUS * (i + 1), GHOST_HEIGHT / 2 - WAVE_RADIUS * 0, GHOST_WIDTH / 2 - WAVE_RADIUS * (i * 2 + 2) / 2, GHOST_HEIGHT / 2 - WAVE_RADIUS * 0.5, WAVE_RADIUS / 2);
						} else {
							ctx.arcTo(GHOST_WIDTH / 2 - WAVE_RADIUS * i      , GHOST_HEIGHT / 2 - WAVE_RADIUS * 1, GHOST_WIDTH / 2 - WAVE_RADIUS * (i * 2 + 1) / 2, GHOST_HEIGHT / 2 - WAVE_RADIUS * 1.0, WAVE_RADIUS / 2);
							ctx.arcTo(GHOST_WIDTH / 2 - WAVE_RADIUS * (i + 1), GHOST_HEIGHT / 2 - WAVE_RADIUS * 1, GHOST_WIDTH / 2 - WAVE_RADIUS * (i * 2 + 2) / 2, GHOST_HEIGHT / 2 - WAVE_RADIUS * 0.5, WAVE_RADIUS / 2);
						}
					}
					
					// Eyes
					const EYE_X = -0.17;
					const EYE_Y = -0.27;
					const EYE_R = 0.15;
					ctx.moveTo(EYE_X + EYE_R, EYE_Y);
					ctx.arc(EYE_X, EYE_Y, EYE_R, 0, 2 * Math.PI, true);
					ctx.moveTo(-EYE_X + EYE_R, EYE_Y);
					ctx.arc(-EYE_X, EYE_Y, EYE_R, 0, 2 * Math.PI, true);
					
					ctx.strokeStyle = "rgba(0, 0, 0, 0.6)";
					ctx.stroke();
					ctx.fillStyle = "#fff";
					ctx.fill();
				} else if (entity.type === "monster") {
					ctx.beginPath();
					ctx.rect(-0.3, -0.1, 0.6, 0.5);
					ctx.arc(0, -0.1, 0.3, 0, 2 * Math.PI, true);
					ctx.fillStyle = "#284";
					ctx.fill();
					ctx.strokeStyle = "#000";
					ctx.stroke();
				} else {
					ctx.fillStyle = "#222";
					ctx.beginPath();
					ctx.arc(0, 0, 0.5, 0, 2 * Math.PI);
					ctx.fill();
					ctx.textAlign = "center";
					ctx.fillStyle = "#fff";
					ctx.font = "bold 0.6px Arial";
					ctx.fillText("?", 0, 0.2);
					
					ctx.strokeStyle = "rgba(255, 255, 255, 0.1)";
					ctx.lineWidth = 0.1;
					ctx.beginPath();
					ctx.arc(0, 0, 0.4, 0, 2 * Math.PI);
					ctx.stroke();
				}
				ctx.restore();
			}
			
			ctx.restore();
		}
		
		window.requestAnimationFrame(() => this.render());
	}
}

window.addEventListener("load", () => {
	window.game = new Game();
	document.body.appendChild(window.game.element);
});