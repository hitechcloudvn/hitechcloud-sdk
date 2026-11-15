> [English](ruby.md) | **Tiếng Việt**

# Hướng dẫn SDK Ruby

## Cài đặt

```bash
gem install hitechcloud
```

## Khởi tạo

```ruby
require 'hitechcloud'

client = HiTechCloud::Client.new(api_key: 'your-api-key')
```

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

## Kiểm thử

```ruby
require 'minitest/autorun'
require 'hitechcloud'

class HiTechCloudTest < Minitest::Test
  def test_client_init
    client = HiTechCloud::Client.new(api_key: 'test-key')
    assert client
  end
end
```

## Biến môi trường

| Biến | Mô tả |
|----------|-------------|
| `HITECHCLOUD_API_KEY` | API key của bạn |
| `HITECHCLOUD_BASE_URL` | Base URL tùy chỉnh (mặc định: `https://docs.hitechcloud.vn`) |

## Thêm ví dụ

Xem thư mục [`examples/ruby/`](../../examples/ruby/) để có các ví dụ đầy đủ.
