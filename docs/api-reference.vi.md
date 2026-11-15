> [English](api-reference.md) | **Tiếng Việt**

# Tham chiếu API

URL gốc: `https://docs.hitechcloud.vn`

## Nhóm tài nguyên

| Nhóm | Lớp tài nguyên | Mô tả |
|------|----------------|-------|
| Xác thực tài khoản | `account_authentication` | Đăng nhập, đăng xuất, đặt lại mật khẩu |
| Tiếp thị liên kết | `affiliate` | Quản lý chương trình affiliate |
| Thanh toán & Hợp đồng | `billing_contracts` | Hoá đơn, thanh toán, hợp đồng |
| Giỏ hàng | `cart` | Thao tác giỏ hàng |
| Cloud GPU | `cloud_gpu` | Phiên GPU, cụm, ổ đĩa |
| Liên hệ | `contacts` | Quản lý liên hệ |
| DNS | `dns` | Quản lý vùng DNS |
| Quản lý DNS | `dns_manage` | CRUD bản ghi DNS |
| Miền | `domains` | Đăng ký, chuyển nhượng, gia hạn miền |
| Thông báo | `notifications` | Tuỳ chọn thông báo |
| Chứng chỉ SSL | `ssl_certificates` | Vòng đời chứng chỉ SSL |
| Dịch vụ | `services` | Quản lý dịch vụ (VPS, hosting, v.v.) |
| Hỗ trợ | `support` | Phiếu hỗ trợ |
| Hồ sơ người dùng | `user_profile` | Quản lý hồ sơ |
| Dịch vụ Virtualizor | `virtualizor_services` | Quản lý VPS Virtualizor |

## Tham số chung

| Tham số | Kiểu | Mô tả |
|---------|------|-------|
| `id` | integer | ID dịch vụ/tài nguyên |
| `page` | integer | Số trang để phân trang |
| `per_page` | integer | Số mục mỗi trang (mặc định: 20) |

## Định dạng phản hồi

Mọi phản hồi đều là JSON. Phản hồi thành công trả về HTTP 200 kèm dữ liệu.
Lỗi nghiệp vụ cũng trả về HTTP 200 với trường `error`.

```json
{
  "success": true,
  "data": { ... }
}
```

Phản hồi lỗi:
```json
{
  "error": "Service not found",
  "code": 404
}
```
