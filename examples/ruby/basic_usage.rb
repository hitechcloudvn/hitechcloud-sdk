#!/usr/bin/env ruby
/**
 * Ví dụ sử dụng cơ bản cho HiTechCloud Ruby SDK / Basic usage example for HiTechCloud Ruby SDK
 * 
 * Đặt API key của bạn: / Set your API key:
 *   export HITECHCLOUD_API_KEY=your-api-key
 */
require 'hitechcloud'

client = HiTechCloud::Client.new(api_key: 'your-api-key')

# Liệt kê tất cả dịch vụ / List all services
services = client.services.list_services
services['data'].each do |svc|
  puts "#{svc['id']}: #{svc['name']}"
end

# Lấy một dịch vụ cụ thể / Get a specific service
service = client.services.get_service(id: 123)
puts service

# Các thao tác nguồn / Power operations
client.services.start_service(id: 123)
client.services.stop_service(id: 123)
client.services.reboot_service(id: 123)
