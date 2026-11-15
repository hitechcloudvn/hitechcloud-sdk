> [English](README.md) | **Tiếng Việt**

# Ví dụ Python

## Sử dụng cơ bản

```python
# Liệt kê tất cả dịch vụ
services = client.services.list_services()
for svc in services.get('data', []):
    print(f"{svc['id']}: {svc['name']}")

# Lấy một dịch vụ cụ thể
service = client.services.get_service(id=123)
print(service)

# Các thao tác nguồn
client.services.start_service(id=123)
client.services.stop_service(id=123)
client.services.reboot_service(id=123)
```

## Cách chạy

```bash
export HITECHCLOUD_API_KEY=your-api-key
python3 basic_usage.py
```

## Thêm ví dụ

- `basic_usage.py` — Liệt kê, xem và quản lý dịch vụ
- `error_handling.py` — Xử lý lỗi API một cách nhẹ nhàng
- `pagination.py` — Phân trang kết quả
