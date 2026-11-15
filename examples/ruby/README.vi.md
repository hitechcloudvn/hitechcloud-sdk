> [English](README.md) | **Tiếng Việt**

# Ví dụ Ruby

## Sử dụng cơ bản

```ruby
# Liệt kê tất cả dịch vụ
services = client.services.list_services
services['data'].each do |svc|
  puts "#{svc['id']}: #{svc['name']}"
end

# Lấy một dịch vụ cụ thể
service = client.services.get_service(id: 123)
puts service

# Các thao tác nguồn
client.services.start_service(id: 123)
client.services.stop_service(id: 123)
client.services.reboot_service(id: 123)
```

## Cách chạy

```bash
export HITECHCLOUD_API_KEY=your-api-key
ruby basic_usage.rb
```

## Thêm ví dụ

- `basic_usage.rb` — Liệt kê, xem và quản lý dịch vụ
- `error_handling.rb` — Xử lý lỗi API một cách nhẹ nhàng
- `pagination.rb` — Phân trang kết quả
