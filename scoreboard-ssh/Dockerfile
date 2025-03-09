FROM rustlang/rust:nightly-slim@sha256:80de01c4d0edf0258d14269088dc0877baf6dc734e6bb48c839903fc5c36f336

WORKDIR /app

COPY Cargo.toml Cargo.lock ./
COPY src src
COPY scoreboard.json watsup.json ./

RUN cargo build --release

CMD ["cargo", "run", "--release"]
