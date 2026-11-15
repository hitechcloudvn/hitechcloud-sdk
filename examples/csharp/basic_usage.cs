#!/usr/bin/env csharp
/**
 * Ví dụ sử dụng cơ bản cho HiTechCloud Csharp SDK / Basic usage example for HiTechCloud Csharp SDK
 * 
 * Đặt API key của bạn: / Set your API key:
 *   export HITECHCLOUD_API_KEY=your-api-key
 */
using HiTechCloud;

var client = new HiTechCloudClient("your-api-key");

// Liệt kê tất cả dịch vụ / List all services
var services = await client.Services.ListServicesAsync();
foreach (var svc in services.Data)
{
    Console.WriteLine($"{svc.Id}: {svc.Name}");
}

// Lấy một dịch vụ cụ thể / Get a specific service
var service = await client.Services.GetServiceAsync(123);
Console.WriteLine(service);

// Các thao tác nguồn / Power operations
await client.Services.StartServiceAsync(123);
await client.Services.StopServiceAsync(123);
await client.Services.RebootServiceAsync(123);
