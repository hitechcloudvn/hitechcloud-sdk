> [English](README.md) | **Tiếng Việt**

# Ví dụ Java

## Sử dụng cơ bản

```java
// Liệt kê tất cả dịch vụ
var services = client.services().listServices();
System.out.println(services);

// Lấy một dịch vụ cụ thể
var service = client.services().getService(123);
System.out.println(service);

// Các thao tác nguồn
client.services().startService(123);
client.services().stopService(123);
client.services().rebootService(123);
```

## Cách chạy

```bash
export HITECHCLOUD_API_KEY=your-api-key
java BasicUsage
```

## Thêm ví dụ

- `basic_usage.java` — Liệt kê, xem và quản lý dịch vụ
- `error_handling.java` — Xử lý lỗi API một cách nhẹ nhàng
- `pagination.java` — Phân trang kết quả
