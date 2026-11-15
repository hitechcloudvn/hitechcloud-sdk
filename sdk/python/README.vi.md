> [English](README.md) | **Tiếng Việt**

# HiTechCloud Python SDK

[![PyPI version](https://img.shields.io/pypi/v/hitechcloud-sdk)](https://pypi.org/project/hitechcloud-sdk/)
[![Python](https://img.shields.io/pypi/pyversions/hitechcloud-sdk)](https://pypi.org/project/hitechcloud-sdk/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Python SDK chính thức cho [HiTechCloud UserAPI](https://docs.hitechcloud.vn).

## Cài đặt

```bash
pip install hitechcloud-sdk
```

Để hỗ trợ bất đồng bộ:
```bash
pip install hitechcloud-sdk[async]
```

## Bắt đầu nhanh

```python
from hitechcloud import HiTechCloud

# Sử dụng thông tin đăng nhập
client = HiTechCloud(
    base_url="https://api.hitechcloud.vn",
    username="your@email.com",
    password="your_password"
)

# Hoặc sử dụng token có sẵn
client = HiTechCloud(
    base_url="https://api.hitechcloud.vn",
    token="your_api_token"
)

# Context manager (tự động đóng session)
with HiTechCloud(token="your_token") as client:
    services = client.services.list()
    domains = client.domains.list()
```

## Môi trường

| Môi trường | Base URL |
|-------------|----------|
| Production  | `https://api.hitechcloud.vn` |
| Sandbox (OTE) | `https://ote.hitechcloud.vn` |

## Các nhóm API

### Tài khoản & Xác thực
```python
client.auth.login("email", "password")
client.auth.logout()
client.auth.refresh_token()
client.auth.revoke_token()
client.auth.password_reset("email")
client.auth.signup("email", "password", "first_name", "last_name")
```

### Hồ sơ người dùng
```python
client.users.get_details()
client.users.update_details(first_name="New Name")
client.users.get_logs()
```

### Dịch vụ
```python
client.services.list()
client.services.get(service_id)
client.services.cancel(service_id, reason="No longer needed")
client.services.renew(service_id)
```

### Tên miền
```python
client.domains.list()
client.domains.get(domain_id)
client.domains.whois("example.com")
client.domains.check_availability("example.com")
client.domains.get_nameservers(domain_id)
client.domains.update_nameservers(domain_id, ns1="ns1.example.com", ns2="ns2.example.com")
client.domains.get_epp_code(domain_id)
```

### Quản lý DNS
```python
# Bản ghi DNS của tên miền
client.dns_manage.get_records(domain_id)
client.dns_manage.create_record(domain_id, name="www", record_type="A", content="1.2.3.4")
client.dns_manage.update_record(domain_id, index=0, content="5.6.7.8")
client.dns_manage.delete_record(domain_id, index=0)

# Zone DNS của dịch vụ
client.dns.list()
client.dns.add_zone(service_id, "example.com")
client.dns.add_record(service_id, zone_id, name="@",
                       record_type="A", content="1.2.3.4")
```

### Chứng chỉ SSL
```python
client.ssl.list()
client.ssl.get(certificate_id)
client.ssl.download(certificate_id)
client.ssl.order(product_id=1, csr="-----BEGIN CERTIFICATE REQUEST-----...")
```

### Thanh toán
```python
client.billing.get_balance()
client.billing.list_invoices()
client.billing.get_invoice(invoice_id)
client.billing.apply_credit(invoice_id)
client.billing.get_payment_methods()
```

### Hỗ trợ
```python
client.support.list_tickets()
client.support.create_ticket(subject="Help", message="...", department_id=1)
client.support.get_ticket("TKT-001")
client.support.reply("TKT-001", message="Additional info")
client.support.close("TKT-001")
```

### Danh bạ
```python
client.contacts.list()
client.contacts.create(first_name="John", last_name="Doe", email="john@example.com")
client.contacts.update(contact_id, phone="+84901234567")
client.contacts.delete(contact_id)
```

### Giỏ hàng
```python
client.cart.get()
client.cart.add(product_id=1, billing_cycle="monthly")
client.cart.apply_promo("DISCOUNT20")
client.cart.checkout(payment_method="sepay")
```

### Thông báo
```python
client.notifications.list()
client.notifications.get_preferences()
client.notifications.update_preferences(email=True, sms=False)
```

### Tiếp thị liên kết (Affiliate)
```python
client.affiliate.get_details()
client.affiliate.get_balance()
client.affiliate.list_referrals()
client.affiliate.list_transactions()
```

### Cloud GPU
```python
client.cloud_gpu.list()
client.cloud_gpu.get(gpu_id)
client.cloud_gpu.start(gpu_id)
client.cloud_gpu.stop(gpu_id)
client.cloud_gpu.restart(gpu_id)
```

### Virtualizor
```python
client.virtualizor.list()
client.virtualizor.get(vps_id)
client.virtualizor.start(vps_id)
client.virtualizor.stop(vps_id)
```

## Xử lý lỗi

```python
from hitechcloud import (
    HiTechCloud,
    HiTechCloudError,
    AuthenticationError,
    AuthorizationError,
    NotFoundError,
    RateLimitError,
    ValidationError,
    ServerError,
)

try:
    client.domains.get(99999)
except AuthenticationError:
    print("Thông tin đăng nhập không hợp lệ hoặc token đã hết hạn")
except NotFoundError:
    print("Không tìm thấy tài nguyên")
except RateLimitError as e:
    print(f"Vượt giới hạn tần suất. Thử lại sau {e.retry_after} giây")
except ValidationError as e:
    print(f"Lỗi xác thực dữ liệu: {e}")
except ServerError:
    print("Lỗi máy chủ, vui lòng thử lại sau")
except HiTechCloudError as e:
    print(f"Lỗi API: {e}")
```

## Giới hạn tần suất (Rate Limiting)

SDK tự động xử lý giới hạn tần suất (HTTP 429) bằng cách:
1. Đọc header `Retry-After`
2. Ném ra `RateLimitError` với thuộc tính `retry_after`
3. Áp dụng exponential backoff cho các lần thử lại

## Phát triển

```bash
# Cài đặt các dependency cho phát triển
pip install -e ".[dev]"

# Chạy các bài kiểm thử
pytest

# Chạy các bài kiểm thử kèm độ phủ
pytest --cov=hitechcloud
```

## Giấy phép

Giấy phép MIT - xem [LICENSE](../LICENSE) để biết chi tiết.
