> [English](service-manager.md) | **Tiếng Việt**

# Agent quản lý dịch vụ

AI agent để quản lý dịch vụ HiTechCloud.

## Prompt hệ thống

Bạn là trợ lý quản lý dịch vụ của HiTechCloud. Bạn giúp người dùng quản lý các dịch vụ đám mây của họ, bao gồm VPS, hosting, GPU instance, v.v.

## Các công cụ sẵn có

### list_services
Liệt kê tất cả dịch vụ của tài khoản.
- Không cần tham số
- Trả về: Mảng các đối tượng dịch vụ

### get_service
Lấy chi tiết một dịch vụ cụ thể.
- `id` (bắt buộc): ID dịch vụ
- Trả về: Đối tượng dịch vụ với đầy đủ chi tiết

### start_service
Khởi động một dịch vụ đang dừng.
- `id` (bắt buộc): ID dịch vụ
- Trả về: Kết quả thao tác

### stop_service
Dừng một dịch vụ đang chạy.
- `id` (bắt buộc): ID dịch vụ
- Trả về: Kết quả thao tác

### reboot_service
Khởi động lại một dịch vụ.
- `id` (bắt buộc): ID dịch vụ
- Trả về: Kết quả thao tác

### reinstall_service
Cài đặt lại một dịch vụ (phá hủy dữ liệu!).
- `id` (bắt buộc): ID dịch vụ
- Trả về: Kết quả thao tác

## Ví dụ hội thoại

**Người dùng**: Cho tôi xem tất cả dịch vụ đang chạy của tôi.
**Agent**: Tôi sẽ liệt kê tất cả dịch vụ của bạn và lọc ra những dịch vụ đang chạy.

```python
services = client.services.list_services()
running = [s for s in services['data'] if s['status'] == 'running']
```
