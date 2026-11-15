> **English** | [Tiếng Việt](README.vi.md)

# HiTechCloud UserAPI SDK

[![CI](https://github.com/hitechcloudvn/hitechcloud-sdk/actions/workflows/ci.yml/badge.svg)](https://github.com/hitechcloudvn/hitechcloud-sdk/actions/workflows/ci.yml)
[![GitLab Pipeline](https://gitlab.com/hitechcloudvn/hitechcloud-sdk/badges/main/pipeline.svg)](https://gitlab.com/hitechcloudvn/hitechcloud-sdk/-/pipelines)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

> Official multi-language SDK for the [HiTechCloud User API](https://docs.hitechcloud.vn) — 348 endpoints across 38 groups.

## Overview

This repository contains official SDK implementations for the HiTechCloud User API in multiple programming languages. The SDKs provide a clean, idiomatic interface to manage your HiTechCloud account: services, domains, DNS, billing, support tickets, SSL certificates, and more.

## Supported Languages

| Language | Package | Status |
|----------|---------|--------|
| Python | `hitechcloud-sdk` (PyPI) | ✅ |
| Node.js | `hitechcloud-sdk` (npm) | ✅ |
| PHP | `hitechcloud/sdk` (Composer) | ✅ |
| Go | `github.com/hitechcloudvn/hitechcloud-sdk/sdk/go` | ✅ |
| Ruby | `hitechcloud_sdk` (RubyGems) | ✅ |
| Java | `vn.hitechcloud:sdk` (Maven) | ✅ |
| C# | `HiTechCloud.SDK` (NuGet) | ✅ |
| Kotlin | `vn.hitechcloud:sdk-kotlin` (Maven) | ✅ |
| Swift | `HiTechCloudSDK` (SPM) | ✅ |
| Rust | `hitechcloud-sdk` (crates.io) | ✅ |
| Dart/Flutter | `hitechcloud_sdk` (pub.dev) | ✅ |
| R | `hitechcloudsdk` (CRAN) | ✅ |
| Julia | `HiTechCloudSDK.jl` | ✅ |
| Clojure | `hitechcloud/sdk` (Clojars) | ✅ |
| OCaml | `hitechcloud-sdk` (opam) | ✅ |
| Objective-C | HiTechCloudSDK (CocoaPods) | ✅ |
| Shell (Bash) | `hitechcloud.sh` | ✅ |
| PowerShell | `HiTechCloud.SDK` | ✅ |
| C (libcurl) | `libhitechcloud` | ✅ |

## API Groups

| Group | Endpoints | Description |
|-------|-----------|-------------|
| Account Authentication | 6 | Login, logout, token refresh, signup |
| User Profile | 3 | View/update account details, logs |
| Affiliate | 6 | Summary, campaigns, commissions, payouts |
| Billing & Contracts | 6 | Balance, invoices, payment methods |
| Support | 13 | Tickets, departments, news, knowledgebase |
| Contacts | 5 | CRUD contacts, privileges |
| Domains | 25 | WHOIS, nameservers, EPP, registration |
| DNS Manage | 10 | DNS records, DNSSEC, nameserver registration |
| SSL Certificates | 6 | List, order, download certificates |
| Services | 9 | List, details, cancel, renew, billing cycle |
| Cart | 6 | Categories, products, orders, quotes |
| DNS | 8 | DNS zones and records management |
| Notifications | 3 | Portal notifications |
| Virtualizor Services | 4 | VM suspend, unsuspend, rebuild, SSH keys |
| Cloud GPU | 6 | VM reboot, start, stop, firewall rules |

## Quick Start

### Authentication

```bash
# Get token / Lấy token
TOKEN=$(curl -s -X POST "https://ote.hitechcloud.vn/login" \
  -H "Content-Type: application/json" \
  -d '{"username":"demo@ote.hitechcloud.vn","password":"W9vTyPuJLyTVHAJ3Fe3Coi2h"}' \
  | grep -o '"token":"[^"]*"' | cut -d'"' -f4)

# Use token / Dùng token
curl -s "https://ote.hitechcloud.vn/service" \
  -H "Authorization: Bearer $TOKEN"
```

### Python

```python
from hitechcloud import HiTechCloud

client = HiTechCloud(
    base_url="https://ote.hitechcloud.vn",
    username="demo@ote.hitechcloud.vn",
    password="W9vTyPuJLyTVHAJ3Fe3Coi2h"
)

# List services / Danh sách dịch vụ
services = client.services.list()

# Get domain details / Xem chi tiết miền
domain = client.domains.get(123)

# Create support ticket / Tạo phiếu hỗ trợ
ticket = client.support.create_ticket(
    subject="Need help",
    message="Having an issue with my service",
    department_id=1
)
```

### Node.js

```javascript
import { HiTechCloud } from 'hitechcloud-sdk';

const client = new HiTechCloud({
    baseUrl: 'https://ote.hitechcloud.vn',
    username: 'demo@ote.hitechcloud.vn',
    password: 'W9vTyPuJLyTVHAJ3Fe3Coi2h'
});

// List services / Danh sách dịch vụ
const services = await client.services.list();

// Get domain details / Xem chi tiết miền
const domain = await client.domains.get(123);
```

### PHP

```php
use HiTechCloud\SDK\Client;

$client = new Client([
    'base_url' => 'https://ote.hitechcloud.vn',
    'username' => 'demo@ote.hitechcloud.vn',
    'password' => 'W9vTyPuJLyTVHAJ3Fe3Coi2h',
]);

// List services / Danh sách dịch vụ
$services = client->services()->list();

// Get domain details / Xem chi tiết miền
$domain = client->domains()->get(123);
```

## Environments

| Environment | Base URL | Description |
|-------------|----------|-------------|
| Production | `https://api.hitechcloud.vn` | Real account data |
| OTE (Sandbox) | `https://ote.hitechcloud.vn` | Test environment with sample data |

## Rate Limits

- **120 requests/minute** per IP at the proxy layer
- Additional rate limits from the origin server
- `429` response with `Retry-After` header when exceeded

## Error Codes

| Code | Description |
|------|-------------|
| 401 | Missing or invalid authentication token |
| 403 | Insufficient permissions |
| 404 | Resource not found |
| 429 | Rate limit exceeded |
| 500 | Internal server error |
| 501 | Endpoint not implemented in OTE |

## Development

```bash
# Clone repository / Clone kho lưu trữ
git clone https://github.com/hitechcloudvn/hitechcloud-sdk.git
cd hitechcloud-sdk

# Run tests (Python example) / Chạy kiểm thử (ví dụ Python)
cd sdk/python && pip install -e ".[dev]" && pytest

# Run tests (Node.js example) / Chạy kiểm thử (ví dụ Node.js)
cd sdk/nodejs && npm install && npm test
```

CI/CD runs automatically on GitHub Actions and GitLab CI when you push — do not build release artifacts locally.
CI/CD chạy tự động trên GitHub Actions và GitLab CI khi bạn đẩy code — không build artifact phát hành ở máy cục bộ.

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md) / Xem [CONTRIBUTING.md](CONTRIBUTING.vi.md).

## Changelog

See [CHANGELOG.md](CHANGELOG.md) / Xem [CHANGELOG.md](CHANGELOG.vi.md).

## License

MIT License - see [LICENSE](LICENSE) for details.

## Support

- 📧 Email: support@hitechcloud.vn
- 📖 Documentation: https://docs.hitechcloud.vn
- 🐛 Issues: https://github.com/hitechcloudvn/hitechcloud-sdk/issues
- 🦊 GitLab Mirror: https://gitlab.com/hitechcloudvn/hitechcloud-sdk
