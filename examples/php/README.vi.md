> [English](README.md) | **Tiếng Việt**

# Ví dụ Php

## Sử dụng cơ bản

```php
// Liệt kê tất cả dịch vụ
$services = client->services()->listServices();
foreach ($services['data'] as $svc) {
    echo $svc['id'] . ': ' . $svc['name'] . PHP_EOL;
}

// Lấy một dịch vụ cụ thể
$service = client->services()->getService(['id' => 123]);
print_r($service);

// Các thao tác nguồn
client->services()->startService(['id' => 123]);
client->services()->stopService(['id' => 123]);
client->services()->rebootService(['id' => 123]);
```

## Cách chạy

```bash
export HITECHCLOUD_API_KEY=your-api-key
php basic_usage.php
```

## Thêm ví dụ

- `basic_usage.php` — Liệt kê, xem và quản lý dịch vụ
- `error_handling.php` — Xử lý lỗi API một cách nhẹ nhàng
- `pagination.php` — Phân trang kết quả
