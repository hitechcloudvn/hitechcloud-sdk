> [English](ssl-certificate.md) | **Tiếng Việt**

---
name: manage-hitechcloud-ssl
description: Quản lý chứng chỉ SSL trên HiTechCloud
---

# Quản lý chứng chỉ SSL trên HiTechCloud

Dùng skill này khi người dùng muốn mua, cài đặt hoặc quản lý chứng chỉ SSL.

## Các bước

1. **Liệt kê chứng chỉ**: Lấy tất cả chứng chỉ SSL
2. **Kiểm tra trạng thái**: Xác minh trạng thái và hạn sử dụng của chứng chỉ
3. **Tạo CSR**: Tạo Certificate Signing Request
4. **Gửi DCV**: Hoàn tất Domain Control Validation
5. **Cài đặt chứng chỉ**: Tải xuống và cài đặt chứng chỉ

## Mẫu code

```python
# Liệt kê chứng chỉ
certs = client.ssl.list_certificates()

# Lấy chi tiết chứng chỉ
cert = client.ssl.get_certificate(id=cert_id)

# Tạo CSR
csr = client.ssl.generate_csr(data={
    'domain': 'example.com',
    'organization': 'My Company',
    'country': 'VN'
})
```
