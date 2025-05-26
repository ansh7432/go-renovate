# 🔒 Renovate Security Demo - Go Web Server

[![Test](https://github.com/ansh7432/go-renovate/actions/workflows/test.yml/badge.svg)](https://github.com/ansh7432/go-renovate/actions/workflows/test.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/ansh7432/go-renovate)](https://goreportcard.com/report/github.com/ansh7432/go-renovate)
[![Security](https://img.shields.io/badge/Security-Vulnerable%20by%20Design-red)](./SECURITY.md)

This project demonstrates **Renovate's vulnerability monitoring capabilities** by intentionally using vulnerable Go dependencies and showcasing automated security updates.

## 🎯 What This Program Does

This is a simple HTTP web server built with Go that:
- **Serves two endpoints**: `/` (home) and `/health` (health check)
- **Uses deliberately vulnerable dependencies** to showcase Renovate's security features
- **Demonstrates automated dependency management** in a Go project
- **Includes comprehensive testing** and CI/CD pipeline

### Endpoints

| Endpoint | Method | Description |
|----------|---------|-------------|
| `/` | GET | Home page with welcome message and current route |
| `/health` | GET | Health check endpoint returning server status |

## 🚀 How to Run

### Prerequisites
- Go 1.21 or later
- Docker (optional)

### Local Development

1. **Clone the repository**
   ```bash
   git clone https://github.com/ansh7432/go-renovate.git
   cd go-renovate
   ```

2. **Install dependencies**
   ```bash
   go mod tidy
   ```

3. **Run the application**
   ```bash
   go run main.go
   ```

4. **Test the endpoints**
   ```bash
   # Test in browser
   open http://localhost:8080
   
   # Or use curl
   curl http://localhost:8080/
   curl http://localhost:8080/health
   ```

### Using Docker

1. **Build the image**
   ```bash
   docker build -t renovate-demo .
   ```

2. **Run the container**
   ```bash
   docker run -p 8080:8080 renovate-demo
   ```

### Running Tests

```bash
# Run all tests
go test -v ./...

# Run tests with coverage
go test -cover ./...

# Run benchmarks
go test -bench=. ./...
```

## 🔥 Vulnerable Dependencies (By Design)

This project intentionally uses **older versions** of dependencies with known security vulnerabilities to demonstrate Renovate's capabilities:

### 1. Gorilla Mux v1.7.3
- **CVE**: [CVE-2020-13443](https://nvd.nist.gov/vuln/detail/CVE-2020-13443)
- **Severity**: High (CVSS 7.5)
- **Description**: Memory exhaustion via crafted HTTP requests
- **Fix Available**: v1.8.0+
- **Used in**: HTTP routing and request handling

### 2. Gin v1.6.3
- **CVE**: [CVE-2020-28483](https://nvd.nist.gov/vuln/detail/CVE-2020-28483)
- **Severity**: High (CVSS 7.5)
- **Description**: Directory traversal vulnerability
- **Fix Available**: v1.7.7+
- **Status**: Listed but not actively used (for demonstration)

### 3. YAML v2.2.8
- **CVE**: [CVE-2022-28948](https://nvd.nist.gov/vuln/detail/CVE-2022-28948)
- **Severity**: High (CVSS 7.5)
- **Description**: Stack overflow in Unmarshal function
- **Fix Available**: v2.4.0+
- **Status**: Listed but not actively used (for demonstration)

> ⚠️ **Security Notice**: These vulnerabilities are intentional for educational purposes. See [SECURITY.md](./SECURITY.md) for detailed vulnerability information.

## 🤖 Renovate Configuration

This repository is configured with **advanced Renovate settings** to:

### Key Features
- ✅ **Immediate vulnerability detection** via OSV database
- ✅ **Automatic PR creation** for security issues
- ✅ **Semantic commit messages** with proper categorization
- ✅ **Dependency dashboard** for oversight
- ✅ **Container dependency monitoring** (Dockerfile)
- ✅ **Custom labeling and assignment** for security PRs

### Configuration Highlights
```json
{
  "vulnerabilityAlerts": { "enabled": true },
  "osvVulnerabilityAlerts": true,
  "prCreation": "immediate",
  "schedule": ["at any time"]
}
```

### Expected Behavior
1. **Detection**: Renovate scans for vulnerabilities daily
2. **PR Creation**: Security PRs created within hours of detection
3. **Labeling**: Automatic labels: `security`, `vulnerability`, `critical-security`
4. **Assignment**: Auto-assigns to repository maintainers
5. **Priority**: Security updates get highest priority

## 📊 Testing and Quality Assurance

### Automated Testing
- **Unit Tests**: Handler and router functionality
- **Integration Tests**: End-to-end endpoint testing
- **Benchmark Tests**: Performance measurement
- **Security Scanning**: Vulnerability detection in CI/CD

### CI/CD Pipeline
```yaml
- Build and test on every push/PR
- Security scanning with gosec
- Vulnerability checking with govulncheck
- Docker container testing
- Automated dependency updates via Renovate
```

## 🔬 How to Test the Assignment

### 1. Verify Vulnerabilities
```bash
# Check for known vulnerabilities
go list -json -m all | grep -E "(gorilla|gin|yaml)"

# Run security scanner
go install golang.org/x/vuln/cmd/govulncheck@latest
govulncheck ./...
```

### 2. Monitor Renovate Activity
1. **Enable Renovate** on your GitHub repository
2. **Check for PRs** within 24 hours
3. **Verify PR content** includes CVE references
4. **Test the fixes** by merging security updates

### 3. Validate Configuration
```bash
# Validate JSON configuration
python3 -m json.tool renovate.json

# Test Docker build
docker build -t test .

# Run comprehensive tests
go test -v -cover ./...
```

## 📈 Expected Results

After enabling Renovate, you should see:

1. **🔄 Dependency Dashboard**: Issue created with security overview
2. **🚨 Security PRs**: Immediate PRs for vulnerable dependencies
3. **🏷️ Proper Labeling**: PRs tagged with security labels
4. **📝 Detailed Descriptions**: CVE information and fix details
5. **⚡ Fast Response**: PRs created within hours of detection

## 🏗️ Project Structure

```
.
├── .github/
│   └── workflows/
│       └── test.yml          # CI/CD pipeline
├── Dockerfile                # Container configuration
├── SECURITY.md              # Security vulnerability report
├── README.md                # This file
├── go.mod                   # Go module with vulnerable deps
├── go.sum                   # Dependency checksums
├── main.go                  # Main application code
├── main_test.go             # Comprehensive tests
└── renovate.json            # Renovate configuration
```

## 🎯 Assignment Objectives Demonstrated

✅ **Go Module Creation**: Clean module structure following best practices  
✅ **Vulnerable Dependencies**: Multiple CVEs with real security impact  
✅ **Functional Application**: Working web server with proper routing  
✅ **Renovate Configuration**: Advanced security-focused setup  
✅ **Documentation**: Comprehensive README and security reports  
✅ **Testing**: Unit tests, integration tests, and benchmarks  
✅ **CI/CD**: Automated testing and security scanning  
✅ **Containerization**: Docker support with multi-stage builds  

## 🔗 Links and References

- **Repository**: [github.com/ansh7432/go-renovate](https://github.com/ansh7432/go-renovate)
- **Renovate Docs**: [docs.renovatebot.com](https://docs.renovatebot.com)
- **CVE Database**: [nvd.nist.gov](https://nvd.nist.gov)
- **Go Vulnerability Database**: [pkg.go.dev/vuln](https://pkg.go.dev/vuln)

## 📞 Support

For questions about this assignment or Renovate configuration:
- 📧 Create an issue in this repository
- 📖 Check the [SECURITY.md](./SECURITY.md) for vulnerability details
- 🔧 Review the [renovate.json](./renovate.json) configuration

---
