#!/usr/bin/env node
/**
 * Ví dụ sử dụng cơ bản cho HiTechCloud Nodejs SDK / Basic usage example for HiTechCloud Nodejs SDK
 * 
 * Đặt API key của bạn: / Set your API key:
 *   export HITECHCLOUD_API_KEY=your-api-key
 */
const HiTechCloud = require('@hitechcloud/sdk');

const client = new HiTechCloud({ apiKey: 'your-api-key' });

// Liệt kê tất cả dịch vụ / List all services
const services = await client.services.listServices();
console.log(services.data);

// Lấy một dịch vụ cụ thể / Get a specific service
const service = await client.services.getService({ id: 123 });
console.log(service);

// Các thao tác nguồn / Power operations
await client.services.startService({ id: 123 });
await client.services.stopService({ id: 123 });
await client.services.rebootService({ id: 123 });
