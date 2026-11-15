> [English](README.md) | **Tiếng Việt**

# Ví dụ Go

## Sử dụng cơ bản

```go
// Liệt kê tất cả dịch vụ
services, err := client.Services.ListServices(nil)
if err != nil {
    log.Fatal(err)
}
fmt.Println(services)

// Lấy một dịch vụ cụ thể
service, err := client.Services.GetService(123)
if err != nil {
    log.Fatal(err)
}
fmt.Println(service)

// Các thao tác nguồn
client.Services.StartService(123)
client.Services.StopService(123)
client.Services.RebootService(123)
```

## Cách chạy

```bash
export HITECHCLOUD_API_KEY=your-api-key
go run basic_usage.go
```

## Thêm ví dụ

- `basic_usage.go` — Liệt kê, xem và quản lý dịch vụ
- `error_handling.go` — Xử lý lỗi API một cách nhẹ nhàng
- `pagination.go` — Phân trang kết quả
