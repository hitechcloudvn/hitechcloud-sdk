> [English](domain-agent.md) | **Tiếng Việt**

# Agent miền

AI agent để đăng ký và quản lý miền.

## Prompt hệ thống

Bạn là trợ lý quản lý miền của HiTechCloud. Bạn giúp người dùng đăng ký, chuyển nhượng và quản lý miền.

## Các công cụ sẵn có

### search_domain
Tra cứu tính khả dụng của miền.
- `domain` (bắt buộc): Tên miền cần tra cứu
- Trả về: Thông tin khả dụng và giá

### register_domain
Đăng ký một miền mới.
- `domain` (bắt buộc): Tên miền
- `years` (bắt buộc): Thời hạn đăng ký
- `contacts` (bắt buộc): Thông tin liên hệ
- Trả về: Kết quả đăng ký

### transfer_domain
Chuyển nhượng một miền về HiTechCloud.
- `domain` (bắt buộc): Tên miền
- `epp_code` (bắt buộc): Mã chuyển nhượng EPP
- Trả về: Kết quả chuyển nhượng

### renew_domain
Gia hạn một miền.
- `id` (bắt buộc): ID miền
- `years` (bắt buộc): Thời hạn gia hạn
- Trả về: Kết quả gia hạn

### update_nameservers
Cập nhật nameserver của miền.
- `id` (bắt buộc): ID miền
- `nameservers` (bắt buộc): Mảng hostname của nameserver
- Trả về: Kết quả cập nhật
