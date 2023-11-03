# SSM wargame platform

## Development setup

### Install

Install:

- Docker (for db)
- Go
- `migrate` CLI - https://github.com/golang-migrate/migrate/tree/master/cmd/migrate
- nodejs and yarn - https://yarnpkg.com/getting-started/install

### DB

Run DB:

```sh
cd backend
docker compose up --build
```

Access it with:

```sh
cd backend
docker compose exec db psql -U postgres
docker compose exec db psql ssm_wargame -U postgres
```

First, connect with `psql` as shown above and create the database:

```sql
CREATE DATABASE ssm_wargame;
```

Then, run the migrations:

```sh
cd backend
./scripts/migrate up
```

Now you are ready to go!

### Backend

Copy `backend/.env.example` to `backend/.env` and change it. The s3 stuff is
outdated. See `backend/internal/config/config.go` for the full structure. Run
backend:

```sh
cd backend
go run cmd/api/main.go
```

### Frontend

Run frontend:

```sh
cd frontend
yarn # Installs
yarn dev
```
