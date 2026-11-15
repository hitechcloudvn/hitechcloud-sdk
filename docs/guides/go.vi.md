> [English](go.md) | **Tiếng Việt**

# Hướng dẫn SDK Go

## Cài đặt

```bash
go get github.com/hitechcloudvn/hitechcloud-sdk/sdk/go
```

## Khởi tạo

```go
package main

import (
    "github.com/hitechcloudvn/hitechcloud-sdk/sdk/go"
)

func main() {
    client := hitechcloud.NewClient("your-api-key")
}
```

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

## Kiểm thử

```go
package hitechcloud

import "testing"

func TestNewClient(t *testing.T) {
    client := NewClient("test-key")
    if client == nil {
        t.Fatal("client should not be nil")
    }
}
```

## Biến môi trường

| Biến | Mô tả |
|----------|-------------|
| `HITECHCLOUD_API_KEY` | API key của bạn |
| `HITECHCLOUD_BASE_URL` | Base URL tùy chỉnh (mặc định: `https://docs.hitechcloud.vn`) |

## Thêm ví dụ

Xem thư mục [`examples/go/`](../../examples/go/) để có các ví dụ đầy đủ.
