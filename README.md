# Stockyard Drifter

**API documentation server — write docs in Markdown or OpenAPI, serve as a beautiful hosted page**

Part of the [Stockyard](https://stockyard.dev) family of self-hosted developer tools.

## Quick Start

```bash
docker run -p 9130:9130 -v drifter_data:/data ghcr.io/stockyard-dev/stockyard-drifter
```

Or with docker-compose:

```bash
docker-compose up -d
```

Open `http://localhost:9130` in your browser.

## Configuration

| Variable | Default | Description |
|----------|---------|-------------|
| `PORT` | `9130` | HTTP port |
| `DATA_DIR` | `./data` | SQLite database directory |
| `DRIFTER_LICENSE_KEY` | *(empty)* | Pro license key |

## Free vs Pro

| | Free | Pro |
|-|------|-----|
| Limits | 1 project, 10 pages | Unlimited projects and pages |
| Price | Free | $2.99/mo |

Get a Pro license at [stockyard.dev/tools/](https://stockyard.dev/tools/).

## Category

Developer Tools

## License

Apache 2.0
