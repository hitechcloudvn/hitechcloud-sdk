> [English](authentication.md) | **Tiếng Việt**

# Xác thực

HiTechCloud User API hỗ trợ hai phương thức xác thực:

## API Key (Basic Auth)

```python
client = HiTechCloud(api_key="your-api-key")
```

API key được gửi dưới dạng header Basic Auth.

## Bearer Token

```python
client = HiTechCloud(token="your-bearer-token")
```

Token được gửi dưới dạng `Authorization: Bearer <token>`.

## Lấy API key của bạn

1. Đăng nhập vào [my.hitechcloud.vn](https://my.hitechcloud.vn)
2. Vào **Cài đặt** → **API Keys**
3. Tạo API key mới

## Thực hành bảo mật tốt nhất

- Không bao giờ commit API key vào kho lưu trữ mã nguồn
- Dùng biến môi trường: `HITECHCLOUD_API_KEY`
- Đổi key định kỳ
- Chỉ dùng quyền tối thiểu cần thiết
