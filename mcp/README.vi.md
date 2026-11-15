> [English](README.md) | **Tiếng Việt**

# Máy chủ MCP của HiTechCloud

Máy chủ Model Context Protocol (MCP) cho User API của HiTechCloud.

## Thiết lập

1. Cài đặt dependencies:
   ```bash
   cd mcp && npm install
   ```

2. Cấu hình môi trường:
   ```bash
   export HITECHCLOUD_API_KEY=your-api-key
   ```

3. Chạy máy chủ:
   ```bash
   node server.js
   ```

## Các công cụ sẵn có

| Công cụ | Mô tả |
|------|-------------|
| `list_services` | Liệt kê tất cả dịch vụ |
| `get_service` | Lấy chi tiết dịch vụ |
| `start_service` | Khởi động dịch vụ |
| `stop_service` | Dừng dịch vụ |
| `reboot_service` | Khởi động lại dịch vụ |
| `list_dns_zones` | Liệt kê vùng DNS |
| `manage_dns_record` | Tạo/cập nhật/xóa bản ghi DNS |
| `list_domains` | Liệt kê miền |
| `search_domain` | Tìm miền khả dụng |
| `list_invoices` | Liệt kê hóa đơn |

## Tích hợp với Claude Desktop

Thêm vào cấu hình Claude Desktop của bạn:

```json
{
  "mcpServers": {
    "hitechcloud": {
      "command": "node",
      "args": ["/path/to/hitechcloud-sdk/mcp/server.js"],
      "env": {
        "HITECHCLOUD_API_KEY": "your-api-key"
      }
    }
  }
}
```
