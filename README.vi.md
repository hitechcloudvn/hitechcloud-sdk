> [English](README.md) | **Tiếng Việt**

# HiTechCloud UserAPI SDK

[![CI](https://github.com/hitechcloudvn/hitechcloud-sdk/actions/workflows/ci.yml/badge.svg)](https://github.com/hitechcloudvn/hitechcloud-sdk/actions/workflows/ci.yml)
[![Pipeline GitLab](https://gitlab.com/hitechcloudvn/hitechcloud-sdk/badges/main/pipeline.svg)](https://gitlab.com/hitechcloudvn/hitechcloud-sdk/-/pipelines)
[![Giấy phép: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

> SDK đa ngôn ngữ chính thức cho [HiTechCloud User API](https://docs.hitechcloud.vn) — 348 endpoint thuộc 38 nhóm.

## Tổng quan

Kho lưu trữ này chứa các SDK chính thức cho HiTechCloud User API trên nhiều ngôn ngữ lập trình. Các SDK cung cấp giao diện gọn gàng, idiom theo từng ngôn ngữ để quản lý tài khoản HiTechCloud của bạn: dịch vụ, miền (domain), DNS, thanh toán, phiếu hỗ trợ, chứng chỉ SSL và nhiều hơn nữa.

## Ngôn ngữ được hỗ trợ

| Ngôn ngữ | Gói | Trạng thái |
|----------|-----|-----------|
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

## Nhóm API

| Nhóm | Endpoint | Mô tả |
|------|----------|-------|
| Xác thực tài khoản | 6 | Đăng nhập, đăng xuất, làm mới token, đăng ký |
| Hồ sơ người dùng | 3 | Xem/cập nhật thông tin tài khoản, nhật ký |
| Tiếp thị liên kết (Affiliate) | 6 | Tổng quan, chiến dịch, hoa hồng, chi trả |
| Thanh toán & Hợp đồng | 6 | Số dư, hóa đơn, phương thức thanh toán |
| Hỗ trợ | 13 | Phiếu hỗ trợ, phòng ban, tin tức, kho kiến thức |
| Liên hệ | 5 | CRUD liên hệ, phân quyền |
| Miền (Domains) | 25 | WHOIS, nameserver, EPP, đăng ký miền |
| Quản lý DNS | 10 | Bản ghi DNS, DNSSEC, đăng ký nameserver |
| Chứng chỉ SSL | 6 | Danh sách, đặt mua, tải chứng chỉ |
| Dịch vụ | 9 | Danh sách, chi tiết, hủy, gia hạn, chu kỳ thanh toán |
| Giỏ hàng | 6 | Danh mục, sản phẩm, đơn hàng, báo giá |
| DNS | 8 | Quản lý vùng DNS và bản ghi |
| Thông báo | 3 | Thông báo cổng thông tin |
| Dịch vụ Virtualizor | 4 | Tạm ngưng, mở lại, dựng lại VM, khóa SSH |
| Cloud GPU | 6 | Khởi động lại, bật, tắt VM, quy tắc tường lửa |

## Bắt đầu nhanh

### Xác thực

```bash
# Lấy token / Get token
TOKEN=$(curl -s -X POST "https://ote.hitechcloud.vn/login" \
  -H "Content-Type: application/json" \
  -d '{"username":"demo@ote.hitechcloud.vn","password":"W9vTyPuJLyTVHAJ3Fe3Coi2h"}' \
  | grep -o '"token":"[^"]*"' | cut -d'"' -f4)

# Dùng token / Use token
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

# Danh sách dịch vụ / List services
services = client.services.list()

# Xem chi tiết miền / Get domain details
domain = client.domains.get(123)

# Tạo phiếu hỗ trợ / Create support ticket
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

// Danh sách dịch vụ / List services
const services = await client.services.list();

// Xem chi tiết miền / Get domain details
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

// Danh sách dịch vụ / List services
$services = client->services()->list();

// Xem chi tiết miền / Get domain details
$domain = client->domains()->get(123);
```

## Môi trường

| Môi trường | URL gốc | Mô tả |
|------------|---------|-------|
| Production | `https://api.hitechcloud.vn` | Dữ liệu tài khoản thật |
| OTE (Sandbox) | `https://ote.hitechcloud.vn` | Môi trường kiểm thử với dữ liệu mẫu |

## Giới hạn tốc độ

- **120 request/phút** cho mỗi IP tại tầng proxy
- Thêm giới hạn tốc độ từ máy chủ gốc
- Phản hồi `429` kèm header `Retry-After` khi vượt giới hạn

## Mã lỗi

| Mã | Mô tả |
|----|-------|
| 401 | Thiếu hoặc token xác thực không hợp lệ |
| 403 | Không đủ quyền |
| 404 | Không tìm thấy tài nguyên |
| 429 | Vượt giới hạn tốc độ |
| 500 | Lỗi nội bộ máy chủ |
| 501 | Endpoint chưa được triển khai trên OTE |

## Phát triển

```bash
# Clone kho lưu trữ / Clone repository
git clone https://github.com/hitechcloudvn/hitechcloud-sdk.git
cd hitechcloud-sdk

# Chạy kiểm thử (ví dụ Python) / Run tests (Python example)
cd sdk/python && pip install -e ".[dev]" && pytest

# Chạy kiểm thử (ví dụ Node.js) / Run tests (Node.js example)
cd sdk/nodejs && npm install && npm test
```

CI/CD chạy tự động trên GitHub Actions và GitLab CI khi bạn đẩy code — không build artifact phát hành ở máy cục bộ.
CI/CD runs automatically on GitHub Actions and GitLab CI when you push — do not build release artifacts locally.

## Đóng góp

Xem [CONTRIBUTING.md](CONTRIBUTING.vi.md) / See [CONTRIBUTING.md](CONTRIBUTING.md).

## Nhật ký thay đổi

Xem [CHANGELOG.md](CHANGELOG.vi.md) / See [CHANGELOG.md](CHANGELOG.md).

## Giấy phép

Giấy phép MIT — xem [LICENSE](LICENSE) để biết chi tiết.

## Hỗ trợ

- 📧 Email: support@hitechcloud.vn
- 📖 Tài liệu: https://docs.hitechcloud.vn
- 🐛 Báo lỗi: https://github.com/hitechcloudvn/hitechcloud-sdk/issues
- 🦊 Bản mirror GitLab: https://gitlab.com/hitechcloudvn/hitechcloud-sdk
