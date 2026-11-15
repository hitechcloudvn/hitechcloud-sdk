> [English](rate-limiting.md) | **Tiếng Việt**

# Giới hạn tốc độ

## Giới hạn

- **Request đã xác thực**: 1000 request mỗi phút
- **Request chưa xác thực**: 60 request mỗi phút

## Header

Thông tin giới hạn tốc độ được trả về trong header:

| Header | Mô tả |
|--------|-------|
| `X-RateLimit-Limit` | Số request tối đa mỗi cửa sổ |
| `X-RateLimit-Remaining` | Số request còn lại |
| `X-RateLimit-Reset` | Unix timestamp khi cửa sổ được đặt lại |

## Xử lý giới hạn tốc độ

Khi bạn vượt giới hạn tốc độ, bạn sẽ nhận được phản hồi 429. SDK xử lý điều này tự động với backoff luỹ thừa.

```python
# SDK sẽ tự động thử lại với backoff
result = client.services.get_service(id=123)
```

## Thực hành tốt nhất

- Cache phản hồi khi có thể
- Dùng endpoint hàng loạt thay vì gọi từng cái
- Triển khai hàng đợi request cho ứng dụng lưu lượng lớn
