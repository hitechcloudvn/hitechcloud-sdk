#!/usr/bin/env php
/**
 * Ví dụ sử dụng cơ bản cho HiTechCloud Php SDK / Basic usage example for HiTechCloud Php SDK
 * 
 * Đặt API key của bạn: / Set your API key:
 *   export HITECHCLOUD_API_KEY=your-api-key
 */
<?php
use HiTechCloud\HiTechCloud;

$client = new HiTechCloud('your-api-key');

// Liệt kê tất cả dịch vụ / List all services
$services = client->services()->listServices();
foreach ($services['data'] as $svc) {
    echo $svc['id'] . ': ' . $svc['name'] . PHP_EOL;
}

// Lấy một dịch vụ cụ thể / Get a specific service
$service = client->services()->getService(['id' => 123]);
print_r($service);

// Các thao tác nguồn / Power operations
client->services()->startService(['id' => 123]);
client->services()->stopService(['id' => 123]);
client->services()->rebootService(['id' => 123]);
