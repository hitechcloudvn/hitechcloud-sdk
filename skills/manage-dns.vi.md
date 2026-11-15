> [English](manage-dns.md) | **Tiếng Việt**

---
name: manage-hitechcloud-dns
description: Quản lý vùng DNS và bản ghi DNS trên HiTechCloud
---

# Quản lý DNS trên HiTechCloud

Dùng skill này khi người dùng muốn quản lý vùng DNS hoặc bản ghi DNS.

## Các bước

1. **Liệt kê vùng**: Lấy tất cả vùng DNS của tài khoản
2. **Xem chi tiết vùng**: Lấy các bản ghi hiện tại của một vùng
3. **Thêm/sửa/xóa bản ghi**: Thực hiện các thay đổi được yêu cầu
4. **Xác minh thay đổi**: Xác nhận các thay đổi đã được áp dụng

## Mẫu code

```python
# Liệt kê tất cả vùng
zones = client.dns.list_zones()

# Lấy bản ghi của một vùng
records = client.dns.get_zone(id=zone_id)

# Thêm một bản ghi A
client.dns.create_record(zone_id=zone_id, data={
    'type': 'A',
    'name': 'www',
    'value': '1.2.3.4',
    'ttl': 3600
})
```
