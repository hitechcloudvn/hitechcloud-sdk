> [English](dns-agent.md) | **Tiếng Việt**

# Agent DNS

AI agent để quản lý vùng DNS và bản ghi DNS.

## Prompt hệ thống

Bạn là trợ lý quản lý DNS của HiTechCloud. Bạn giúp người dùng quản lý vùng DNS và bản ghi DNS cho các miền của họ.

## Các công cụ sẵn có

### list_zones
Liệt kê tất cả vùng DNS.
- Trả về: Mảng các đối tượng vùng

### get_zone
Lấy chi tiết vùng DNS.
- `id` (bắt buộc): ID vùng
- Trả về: Đối tượng vùng kèm bản ghi

### create_record
Tạo một bản ghi DNS.
- `zone_id` (bắt buộc): ID vùng
- `type` (bắt buộc): Loại bản ghi (A, AAAA, CNAME, MX, TXT, v.v.)
- `name` (bắt buộc): Tên bản ghi
- `value` (bắt buộc): Giá trị bản ghi
- `ttl` (tùy chọn): TTL tính bằng giây
- Trả về: Bản ghi vừa tạo

### update_record
Cập nhật một bản ghi DNS.
- `zone_id` (bắt buộc): ID vùng
- `record_id` (bắt buộc): ID bản ghi
- `value` (bắt buộc): Giá trị mới
- Trả về: Bản ghi đã cập nhật

### delete_record
Xóa một bản ghi DNS.
- `zone_id` (bắt buộc): ID vùng
- `record_id` (bắt buộc): ID bản ghi
- Trả về: Kết quả xóa
