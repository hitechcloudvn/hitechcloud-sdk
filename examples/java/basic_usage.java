#!/usr/bin/env java
/**
 * Ví dụ sử dụng cơ bản cho HiTechCloud Java SDK / Basic usage example for HiTechCloud Java SDK
 * 
 * Đặt API key của bạn: / Set your API key:
 *   export HITECHCLOUD_API_KEY=your-api-key
 */
import vn.hitechcloud.HiTechCloud;

HiTechCloud client = new HiTechCloud("your-api-key");

// Liệt kê tất cả dịch vụ / List all services
var services = client.services().listServices();
System.out.println(services);

// Lấy một dịch vụ cụ thể / Get a specific service
var service = client.services().getService(123);
System.out.println(service);

// Các thao tác nguồn / Power operations
client.services().startService(123);
client.services().stopService(123);
client.services().rebootService(123);
