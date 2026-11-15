> [English](nodejs.md) | **Tiếng Việt**

# Hướng dẫn SDK Node.js

## Cài đặt

```bash
npm install @hitechcloud/sdk
```

## Khởi tạo

```nodejs
const HiTechCloud = require('@hitechcloud/sdk');

const client = new HiTechCloud({ apiKey: 'your-api-key' });
```

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

## Kiểm thử

```nodejs
const HiTechCloud = require('@hitechcloud/sdk');

describe('HiTechCloud SDK', () => {
    let client;
    beforeAll(() => {
        client = new HiTechCloud({ apiKey: 'test-key' });
    });

    test('should initialize client', () => {
        expect(client).toBeDefined();
    });
});
```

## Biến môi trường

| Biến | Mô tả |
|----------|-------------|
| `HITECHCLOUD_API_KEY` | API key của bạn |
| `HITECHCLOUD_BASE_URL` | Base URL tùy chỉnh (mặc định: `https://docs.hitechcloud.vn`) |

## Thêm ví dụ

Xem thư mục [`examples/nodejs/`](../../examples/nodejs/) để có các ví dụ đầy đủ.
