# Security Vulnerabilities Report

## Known Vulnerabilities in Dependencies

### 1. Gorilla Mux v1.7.3
- **CVE**: CVE-2020-13443
- **Severity**: High (7.5)
- **Description**: Memory exhaustion via crafted HTTP requests
- **Fixed in**: v1.8.0+
- **CVSS**: 7.5

### 2. Gin v1.6.3
- **CVE**: CVE-2020-28483
- **Severity**: High (7.5)
- **Description**: Directory traversal vulnerability
- **Fixed in**: v1.7.7+

### 3. YAML v2.2.8
- **CVE**: CVE-2022-28948
- **Severity**: High (7.5)
- **Description**: Stack overflow in Unmarshal
- **Fixed in**: v2.4.0+

## Mitigation Strategy
This project uses Renovate to automatically detect and propose fixes for these vulnerabilities.