> [English](domain-operations.md) | **Tiếng Việt**

---
name: manage-hitechcloud-domains
description: Đăng ký, chuyển nhượng và quản lý miền trên HiTechCloud
---

# Quản lý miền trên HiTechCloud

Dùng skill này khi người dùng muốn đăng ký, chuyển nhượng hoặc quản lý miền.

## Các bước

1. **Tìm miền**: Kiểm tra tình trạng khả dụng và giá
2. **Đăng ký/chuyển nhượng**: Hoàn tất đăng ký hoặc chuyển nhượng
3. **Cấu hình DNS**: Thiết lập nameserver và bản ghi DNS
4. **Quản lý liên hệ**: Cập nhật thông tin liên hệ WHOIS

## Mẫu code

```python
# Tìm kiếm miền khả dụng
result = client.domains.search_domain(domain='example.com')

# Đăng ký miền
order = client.domains.register_domain(data={
    'domain': 'example.com',
    'years': 1,
    'registrant': {
        'name': 'Nguyen Van A',
        'email': 'a@example.com',
        'phone': '+84912345678'
    }
})

# Cập nhật nameserver
client.domains.update_nameservers(id=domain_id, data={
    'nameservers': ['ns1.hitechcloud.vn', 'ns2.hitechcloud.vn']
})
```
