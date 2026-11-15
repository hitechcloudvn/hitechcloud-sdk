> [English](python.md) | **Tiếng Việt**

# Hướng dẫn SDK Python

## Cài đặt

```bash
pip install hitechcloud
```

## Khởi tạo

```python
from hitechcloud import HiTechCloud

client = HiTechCloud(api_key="your-api-key")
```

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

## Kiểm thử

```python
import pytest
from hitechcloud import HiTechCloud

def test_client_init():
    client = HiTechCloud(api_key="test-key")
    assert client is not None

def test_list_services():
    client = HiTechCloud(api_key="test-key")
    result = client.services.list_services()
    assert isinstance(result, dict)
```

## Biến môi trường

| Biến | Mô tả |
|----------|-------------|
| `HITECHCLOUD_API_KEY` | API key của bạn |
| `HITECHCLOUD_BASE_URL` | Base URL tùy chỉnh (mặc định: `https://docs.hitechcloud.vn`) |

## Thêm ví dụ

Xem thư mục [`examples/python/`](../../examples/python/) để có các ví dụ đầy đủ.
