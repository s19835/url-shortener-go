# Shortly

## URL Shortener Service in Go

![Go](https://img.shields.io/badge/Go-1.21+-blue)
![Redis](https://img.shields.io/badge/Redis-7.0+-red)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-15+-blue)

A high-performance, production-ready URL shortener service built with Go, Redis (for caching), and PostgreSQL (for persistence). Designed for scalability, security, and deployment flexibility.

## Features

- **RESTful API** with JSON responses
- **Base62 encoding** for compact URLs
- **Redis caching** for instant redirects
- **PostgreSQL persistence** for reliability
- **Custom expiration** per URL
- **Docker-ready** configuration
- **12-factor app** compliant
- **Production-ready architecture** with security best practices

## Quick Start

### Prerequisites

- Go 1.21+
- Redis 7.0+
- PostgreSQL 15+

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/url-shortener.git
   cd url-shortener
   ```

2. Set up environment variables:

   ```bash
   cp .env.example .env
   # Edit .env with your credentials
   ```

3. Run the application:

   ```bash
   go run main.go
   ```

### Using Docker

```bash
docker-compose up -d
```

## API Documentation

### Shorten a URL

**Request**:

```http
POST /shorten
Content-Type: application/json

{
  "url": "https://example.com/very-long-path",
  "expires_in": "24h"
}
```

**Response**:

```json
{
  "short_url": "http://yourdomain/abc123",
  "original_url": "https://example.com/very-long-path",
  "expires_at": "2023-12-31T23:59:59Z"
}
```

### Redirect

```http
GET /abc123
```

*Returns 301 redirect to original URL*

## Configuration

Environment Variables:

| Variable | Description | Default |
|----------|-------------|---------|
| `DB_URL` | PostgreSQL connection URL | `postgres://postgres:postgres@localhost:5432/url_shortener?sslmode=disable` |
| `REDIS_URL` | Redis connection URL | `redis://localhost:6379/0` |
| `SERVER_PORT` | HTTP port | `8080` |
| `ENVIRONMENT` | Runtime environment | `production` |

## Architecture

```
cmd/           # Main application entry point
handlers/      # HTTP endpoint handlers
services/      # Business logic
repositories/  # Database access
models/        # Data structures
config/        # Configuration loading
utils/         # Helper functions
middleware/    # Security & rate limiting
```

## Future Enhancements

### Web Application Features

1. **User Interface**
   - Web form for URL shortening
   - Dashboard for link analytics
   - QR code generation

2. **Authentication**
   - User accounts (OAuth2 support)
   - API key management
   - Rate limiting

3. **Advanced Features**
   - Custom short URLs
   - Link expiration management
   - Bulk URL shortening
   - Click analytics (geolocation, referrers)

4. **Infrastructure**
   - Kubernetes deployment manifests
   - Terraform scripts for cloud provisioning
   - CDN integration for global caching

## Deployment Guide

### Production Deployment

1. **Database Setup**:

   ```bash
   psql -c "CREATE DATABASE url_shortener_prod;"
   psql -c "CREATE USER shortener_prod WITH PASSWORD 'securepassword';"
   psql -c "GRANT ALL PRIVILEGES ON DATABASE url_shortener_prod TO shortener_prod;"
   ```

2. **Systemd Service**:

   ```ini
   # /etc/systemd/system/url-shortener.service
   [Unit]
   Description=URL Shortener Service
   After=network.target

   [Service]
   User=appuser
   WorkingDirectory=/opt/url-shortener
   EnvironmentFile=/opt/url-shortener/.env
   ExecStart=/opt/url-shortener/url-shortener
   Restart=always

   [Install]
   WantedBy=multi-user.target
   ```

3. **Reverse Proxy (Nginx)**:

   ```nginx
   server {
       listen 80;
       server_name short.example.com;

       location / {
           proxy_pass http://localhost:8080;
           proxy_set_header Host $host;
       }
   }
   ```

4. **Kubernetes Deployment** (future enhancement)

   ```yaml
   apiVersion: apps/v1
   kind: Deployment
   metadata:
     name: url-shortener
   spec:
     replicas: 3
     selector:
       matchLabels:
         app: url-shortener
     template:
       metadata:
         labels:
           app: url-shortener
       spec:
         containers:
           - name: url-shortener
             image: your-dockerhub-repo/url-shortener:latest
             envFrom:
               - configMapRef:
                   name: url-shortener-config
   ```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

MIT License. See `LICENSE` for details.
