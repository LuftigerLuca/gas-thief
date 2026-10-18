# ⛽ Gas Thief

> Automated gas station price tracker for Germany — polls the [Tankerkoenig](https://www.tankerkoenig.de/) API and stores historical fuel prices in MariaDB.

[![CI](https://github.com/LuftigerLuca/gas-thief/actions/workflows/ci.yml/badge.svg)](https://github.com/LuftigerLuca/gas-thief/actions/workflows/ci.yml)
[![Go](https://img.shields.io/badge/Go-1.25-00ADD8?logo=go&logoColor=white)](https://go.dev/)
[![Next.js](https://img.shields.io/badge/Next.js-16-000000?logo=nextdotjs&logoColor=white)](https://nextjs.org/)
[![Docker](https://img.shields.io/badge/Docker-Ready-2496ED?logo=docker&logoColor=white)](https://www.docker.com/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

---

## What it does

Gas Thief periodically fetches diesel, E5 and E10 prices from nearby gas stations via the Tankerkoenig API and persists snapshots into a MariaDB database. This lets you build a historical time series of fuel prices for any geographic area in Germany.

- Configurable lookup radius and coordinates
- Configurable polling interval (default: 15 min)
- Upserts station metadata, appends price calls as new rows
- Structured logging via `log/slog`

## Quick Start

### Docker Compose

```yaml
services:
  gas-thief:
    image: ghcr.io/luftigerluca/gas-thief:latest
    env_file: .env
    depends_on:
      mariadb:
        condition: service_healthy
    restart: unless-stopped

  mariadb:
    image: mariadb:latest
    ports:
      - "3306:3306"
    environment:
      MYSQL_ROOT_PASSWORD: password
      MYSQL_DATABASE: gas-thief
    volumes:
      - mariadb_data:/var/lib/mysql
    healthcheck:
      test: ["CMD", "mariadb-admin", "ping", "-h", "localhost"]
      interval: 5s
      timeout: 3s
      retries: 5

volumes:
  mariadb_data:
```

### Local

```bash
# Backend (put your .env in backend/)
cd backend
go run ./app

# Frontend (http://localhost:3000)
cd frontend
pnpm install
pnpm dev
```

## Configuration

All configuration is done via environment variables. Gas Thief looks for a `.env` file in the working directory (`backend/` when running locally), but env vars take precedence.

| Variable           | Required | Default | Description                            |
| ------------------ | -------- | ------- | -------------------------------------- |
| `API_KEY`          | ✅       | —       | Tankerkoenig API key                   |
| `LOOK_UP_LAT`      | ✅       | —       | Latitude for station lookup            |
| `LOOK_UP_LNG`      | ✅       | —       | Longitude for station lookup           |
| `LOOK_UP_RADIUS`   | ❌       | `10`    | Lookup radius in km                    |
| `LOOK_UP_INTERVAL` | ❌       | `15`    | Polling interval in minutes            |
| `DB_HOST`          | ✅       | —       | MariaDB host                           |
| `DB_PORT`          | ✅       | —       | MariaDB port                           |
| `DB_USER`          | ✅       | —       | MariaDB user                           |
| `DB_PASSWORD`      | ✅       | —       | MariaDB password                       |
| `DB_NAME`          | ✅       | —       | MariaDB database name                  |

## Development

### Prerequisites

- Go 1.25+
- [golangci-lint](https://golangci-lint.run/) v2
- Node.js 20+ and [pnpm](https://pnpm.io/)
- Docker (optional)

### Backend

The root `Makefile` delegates to `backend/Makefile`:

```bash
make check       # run all checks (format, lint, build, test)
make build       # compile binary
make lint        # run golangci-lint
make fmt         # format all files
make fmt-check   # check formatting without modifying
make test        # run tests
```

### Frontend

Next.js app with Tailwind CSS and [shadcn/ui](https://ui.shadcn.com/). Run from `frontend/`:

```bash
pnpm dev         # start dev server
pnpm build       # production build
pnpm lint        # run ESLint
pnpm typecheck   # run tsc
pnpm format      # format with Prettier
```

See [`frontend/README.md`](frontend/README.md) for details.

### Pre-commit Hook

The hook runs automatically on every commit (format check + lint). To set it up after cloning:

```bash
git config core.hooksPath .githooks
```

## Project Structure

```
gas-thief/
├── .github/workflows/       # CI + Deploy workflows
├── .githooks/               # Git hooks (pre-commit)
├── Makefile                 # Delegates to backend/Makefile
├── backend/
│   ├── app/
│   │   ├── main.go          # Entry point
│   │   ├── scheduler.go     # Polling loop
│   │   ├── persistence.go   # DB connection + save logic
│   │   ├── settings/
│   │   │   └── settings.go  # Env var loading
│   │   ├── api/
│   │   │   ├── client.go    # HTTP client for Tankerkoenig
│   │   │   ├── dto.go       # Request/response types
│   │   │   └── mapper.go    # API → domain mapping
│   │   └── domain/
│   │       └── models.go    # GORM models (Station, Call)
│   ├── Dockerfile           # Multi-stage build
│   ├── Makefile             # Build, lint, format targets
│   └── .golangci.yml        # Linter config
├── frontend/
│   ├── app/                 # Next.js App Router (layout, pages, global styles)
│   ├── components/          # React components (shadcn/ui lands in components/ui)
│   ├── hooks/               # Custom React hooks
│   ├── lib/                 # Utilities
│   ├── public/              # Static assets
│   └── package.json         # Scripts + dependencies (pnpm)
└── docker-compose.yml       # Local dev (MariaDB only)
```
