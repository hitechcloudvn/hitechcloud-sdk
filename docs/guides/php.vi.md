> [English](php.md) | **Tiếng Việt**

# Hướng dẫn SDK PHP

## Cài đặt

```bash
composer require hitechcloud/sdk
```

## Khởi tạo

```php
<?php
use HiTechCloud\HiTechCloud;

$client = new HiTechCloud('your-api-key');
```

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

## Kiểm thử

```php
<?php
use PHPUnit\Framework\TestCase;
use HiTechCloud\HiTechCloud;

class HiTechCloudTest extends TestCase {
    public function testClientInit() {
        $client = new HiTechCloud('test-key');
        $this->assertNotNull($client);
    }
}
```

## Biến môi trường

| Biến | Mô tả |
|----------|-------------|
| `HITECHCLOUD_API_KEY` | API key của bạn |
| `HITECHCLOUD_BASE_URL` | Base URL tùy chỉnh (mặc định: `https://docs.hitechcloud.vn`) |

## Thêm ví dụ

Xem thư mục [`examples/php/`](../../examples/php/) để có các ví dụ đầy đủ.
