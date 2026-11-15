> [English](README.md) | **Tiếng Việt**

# HiTechCloud PHP SDK

[![Packagist Version](https://img.shields.io/packagist/v/hitechcloud/sdk)](https://packagist.org/packages/hitechcloud/sdk)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

PHP SDK chính thức cho [HiTechCloud UserAPI](https://docs.hitechcloud.vn).

## Cài đặt

```bash
composer require hitechcloud/sdk
```

## Bắt đầu nhanh

```php
<?php

use HiTechCloud\SDK\HiTechCloud;

// Sử dụng thông tin đăng nhập
$client = new HiTechCloud(
    baseUrl: 'https://api.hitechcloud.vn',
    username: 'your@email.com',
    password: 'your_password'
);

// Hoặc sử dụng token có sẵn
$client = new HiTechCloud(
    token: 'your_api_token'
);

// Liệt kê dịch vụ
$services = $client->services->list();

// Lấy chi tiết tên miền
$domain = $client->domains->get(123);

// Kiểm tra số dư
$balance = $client->billing->getBalance();
```

## Môi trường

| Môi trường | Base URL |
|-------------|----------|
| Production  | `https://api.hitechcloud.vn` |
| Sandbox (OTE) | `https://ote.hitechcloud.vn` |

## Các nhóm API

### Tài khoản & Xác thực
```php
$client->auth->login('email', 'password');
$client->auth->logout();
$client->auth->refreshToken();
$client->auth->revokeToken();
$client->auth->passwordReset('email');
$client->auth->signup('email', 'password', 'firstName', 'lastName');
```

### Hồ sơ người dùng
```php
$client->users->getDetails();
$client->users->updateDetails(['first_name' => 'New Name']);
$client->users->getLogs();
```

### Dịch vụ
```php
$client->services->list();
$client->services->get($serviceId);
$client->services->cancel($serviceId, 'No longer needed');
$client->services->renew($serviceId);
```

### Tên miền
```php
$client->domains->list();
$client->domains->get($domainId);
$client->domains->whois('example.com');
$client->domains->checkAvailability('example.com');
$client->domains->getNameservers($domainId);
$client->domains->updateNameservers($domainId, ['ns1' => 'ns1.example.com']);
$client->domains->getEppCode($domainId);
```

### Quản lý DNS
```php
// Bản ghi DNS của tên miền
$client->dnsManage->getRecords($domainId);
$client->dnsManage->createRecord($domainId, ['name' => 'www', 'type' => 'A', 'content' => '1.2.3.4']);

// Zone DNS của dịch vụ
$client->dns->list();
$client->dns->addZone($serviceId, 'example.com');
```

### Chứng chỉ SSL
```php
$client->ssl->list();
$client->ssl->get($certificateId);
$client->ssl->download($certificateId);
```

### Thanh toán
```php
$client->billing->getBalance();
$client->billing->listInvoices();
$client->billing->getInvoice($invoiceId);
$client->billing->applyCredit($invoiceId);
```

### Hỗ trợ
```php
$client->support->listTickets();
$client->support->createTicket(['subject' => 'Help', 'message' => '...', 'department_id' => 1]);
$client->support->getTicket('TKT-001');
$client->support->reply('TKT-001', 'Additional info');
```

### Danh bạ
```php
$client->contacts->list();
$client->contacts->create(['first_name' => 'John', 'last_name' => 'Doe']);
$client->contacts->update($contactId, ['phone' => '+84901234567']);
$client->contacts->delete($contactId);
```

### Giỏ hàng
```php
$client->cart->get();
$client->cart->add(1, ['billing_cycle' => 'monthly']);
$client->cart->applyPromo('DISCOUNT20');
$client->cart->checkout(['payment_method' => 'sepay']);
```

### Thông báo
```php
$client->notifications->list();
$client->notifications->getPreferences();
$client->notifications->updatePreferences(['email' => true, 'sms' => false]);
```

### Tiếp thị liên kết (Affiliate)
```php
$client->affiliate->getDetails();
$client->affiliate->getBalance();
$client->affiliate->listReferrals();
```

### Cloud GPU
```php
$client->cloudGpu->list();
$client->cloudGpu->get($gpuId);
$client->cloudGpu->start($gpuId);
```

### Virtualizor
```php
$client->virtualizor->list();
$client->virtualizor->get($vpsId);
$client->virtualizor->start($vpsId);
```

## Xử lý lỗi

```php
use HiTechCloud\SDK\AuthenticationException;
use HiTechCloud\SDK\NotFoundException;
use HiTechCloud\SDK\RateLimitException;
use HiTechCloud\SDK\ValidationException;
use HiTechCloud\SDK\HiTechCloudException;

try {
    $client->domains->get(99999);
} catch (AuthenticationException $e) {
    echo 'Thông tin đăng nhập không hợp lệ hoặc token đã hết hạn';
} catch (NotFoundException $e) {
    echo 'Không tìm thấy tài nguyên';
} catch (RateLimitException $e) {
    echo 'Vượt giới hạn tần suất. Thử lại sau ' . $e->getRetryAfter() . ' giây';
} catch (ValidationException $e) {
    echo 'Lỗi xác thực dữ liệu: ' . $e->getMessage();
} catch (HiTechCloudException $e) {
    echo 'Lỗi API: ' . $e->getMessage();
}
```

## Giấy phép

Giấy phép MIT - xem [LICENSE](../../LICENSE) để biết chi tiết.
