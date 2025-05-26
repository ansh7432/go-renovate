# Renovate Demo - Go Web Server

This is a simple Go web server that demonstrates Renovate's vulnerability monitoring capabilities.

## What this program does

This program creates a basic HTTP server using the Gorilla Mux router. It serves two endpoints:
- `/` - Home page that displays a welcome message
- `/health` - Health check endpoint

## How to run

1. Clone this repository
2. Run the program:
   ```bash
   go run main.go
   ```
3. Open your browser and navigate to `http://localhost:8080`

## External Dependency

This project uses **Gorilla Mux v1.7.3** as its routing library. We specifically chose this older version because it contains a known vulnerability.

### Vulnerability Details

- **Package**: github.com/gorilla/mux
- **Vulnerable Version**: v1.7.3 (and earlier)
- **CVE**: [CVE-2020-13443](https://nvd.nist.gov/vuln/detail/CVE-2020-13443)
- **Severity**: High
- **Description**: The Gorilla Mux package before v1.8.0 is vulnerable to a potential memory exhaustion attack via crafted HTTP requests.

The vulnerability was fixed in version v1.8.0 and later versions.

## Renovate Configuration

This repository is configured with Renovate to automatically detect and create pull requests for dependency updates when security vulnerabilities are found.