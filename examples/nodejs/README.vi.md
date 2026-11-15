> [English](README.md) | **Tiếng Việt**

# Ví dụ Nodejs

## Sử dụng cơ bản

```nodejs
// Liệt kê tất cả dịch vụ
const services = await client.services.listServices();
console.log(services.data);

// Lấy một dịch vụ cụ thể
const service = await client.services.getService({ id: 123 });
console.log(service);

// Các thao tác nguồn
await client.services.startService({ id: 123 });
await client.services.stopService({ id: 123 });
await client.services.rebootService({ id: 123 });
```

## Cách chạy

```bash
export HITECHCLOUD_API_KEY=your-api-key
node basic_usage.js
```

## Thêm ví dụ

- `basic_usage.js` — Liệt kê, xem và quản lý dịch vụ
- `error_handling.js` — Xử lý lỗi API một cách nhẹ nhàng
- `pagination.js` — Phân trang kết quả
