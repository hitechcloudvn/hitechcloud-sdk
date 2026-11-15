> [English](error-handling.md) | **Tiếng Việt**

# Xử lý lỗi

## Mã trạng thái HTTP

| Mã | Ý nghĩa |
|----|---------|
| 200 | Thành công (hoặc lỗi nghiệp vụ — kiểm tra thân phản hồi) |
| 401 | Chưa xác thực — thông tin đăng nhập không hợp lệ hoặc bị thiếu |
| 403 | Bị cấm — không đủ quyền |
| 404 | Không tìm thấy |
| 422 | Lỗi kiểm tra dữ liệu |
| 429 | Vượt giới hạn tốc độ |
| 500 | Lỗi nội bộ máy chủ |

## Lỗi nghiệp vụ

HiTechCloud trả về lỗi nghiệp vụ trong các phản hồi HTTP 200. Luôn kiểm tra thân phản hồi:

```python
try:
    result = client.services.get_service(id=123)
    if 'error' in result:
        print(f"Lỗi nghiệp vụ / Business error: {result['error']}")
except HiTechCloudError as e:
    print(f"Lỗi API / API error: {e.status_code} - {e.message}")
```

## Logic thử lại

SDK tự động thử lại khi:
- 429 (giới hạn tốc độ) — với backoff luỹ thừa
- 500, 502, 503, 504 — tối đa 3 lần

Tắt tính năng này với: `client = HiTechCloud(retry=False)`
