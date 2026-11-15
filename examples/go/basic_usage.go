#!/usr/bin/env go
/**
 * Ví dụ sử dụng cơ bản cho HiTechCloud Go SDK / Basic usage example for HiTechCloud Go SDK
 * 
 * Đặt API key của bạn: / Set your API key:
 *   export HITECHCLOUD_API_KEY=your-api-key
 */
package main

import (
    "github.com/hitechcloudvn/hitechcloud-sdk/sdk/go"
)

func main() {
    client := hitechcloud.NewClient("your-api-key")
}

// Liệt kê tất cả dịch vụ / List all services
services, err := client.Services.ListServices(nil)
if err != nil {
    log.Fatal(err)
}
fmt.Println(services)

// Lấy một dịch vụ cụ thể / Get a specific service
service, err := client.Services.GetService(123)
if err != nil {
    log.Fatal(err)
}
fmt.Println(service)

// Các thao tác nguồn / Power operations
client.Services.StartService(123)
client.Services.StopService(123)
client.Services.RebootService(123)
