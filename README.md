# rest-template

A full-stack web application template using a Go REST API backend and a Vue 3 frontend.

## Stack

| Layer    | Technology |
|----------|------------|
| Backend  | Go, Gin, sqlx |
| Database | PostgreSQL 17 |
| Frontend | Vue 3, Vite, Pinia, Vue Router |
| Dev      | Docker Compose, Air (live reload) |

## Project structure

```
.
├── db/                 # Database connection, migrations, and data layer
│   └── migrations/     # SQL migration files (applied in order on startup)
├── middleware/         # Gin middleware (session auth)
├── roles/              # Role management HTTP handlers
├── users/              # User management HTTP handlers
├── ui/                 # Vue 3 frontend
│   └── src/
│       ├── stores/     # Pinia stores (auth state)
│       ├── router/     # Vue Router (route definitions + auth guard)
│       ├── utils/      # Fetch wrapper (api.js)
│       └── views/      # Page components
└── main.go             # Entry point, route registration
```

## Getting started

### Development

Start the API and database with live reload:

```bash
make up-d
```

The API will be available at `http://localhost:8080`. Air watches for Go file changes and restarts automatically.

Start the frontend dev server:

```bash
cd ui && npm install && npm run dev
```

The UI will be available at `http://localhost:5173`. API requests to `/api/*` are proxied to the backend.

### Default credentials

A root admin account is seeded on first startup (see environment variables below). The root user is automatically assigned the `admin` role.

## Environment variables

| Variable             | Default              | Description |
|----------------------|----------------------|-------------|
| `APP_ENV`            | `dev`                | Set to `production` to enable release mode and disable request logging |
| `LOG_LEVEL`          | `debug`              | Log level: `debug`, `info`, `warn`, or `error` |
| `POSTGRES_HOST`      | `db`                 | PostgreSQL host |
| `POSTGRES_PORT`      | `5432`               | PostgreSQL port |
| `POSTGRES_USER`      | `dev`                | PostgreSQL user |
| `POSTGRES_PASSWORD`  | `secret`             | PostgreSQL password |
| `POSTGRES_DB`        | `app`                | PostgreSQL database name |
| `ROOT_USER_EMAIL`    | `admin@example.com`  | Email for the seeded root admin account |
| `ROOT_USER_PASSWORD` | `password`           | Password for the seeded root admin account |

## API

All endpoints except `/login` and `/healthcheck` require an `Authorization: Token <token>` header, obtained from the login response.

### Auth

| Method | Path      | Description         |
|--------|-----------|---------------------|
| `POST` | `/login`  | Log in, returns token and user |
| `POST` | `/logout` | Invalidate the current session |

### Users

| Method  | Path                        | Description             |
|---------|-----------------------------|-------------------------|
| `POST`  | `/users`                    | Create a user           |
| `PATCH` | `/users/:userId`            | Update email/display name |
| `GET`   | `/users/:userId/roles`      | List roles for a user   |
| `POST`  | `/users/:userId/roles`      | Assign a role to a user |
| `DELETE`| `/users/:userId/roles/:roleId` | Remove a role from a user |

### Roles

| Method  | Path             | Description      |
|---------|------------------|------------------|
| `GET`   | `/roles`         | List all roles   |
| `POST`  | `/roles`         | Create a role    |
| `PATCH` | `/roles/:roleId` | Update a role    |
| `DELETE`| `/roles/:roleId` | Delete a role    |

## Database migrations

Migration files live in `db/migrations/` and are named `NNN_description.sql`. They are applied automatically in lexicographic order on startup. Applied migrations are tracked in the `schema_migrations` table and will not be re-run.

## Production build

```bash
docker build --target final -t rest-template .
```
