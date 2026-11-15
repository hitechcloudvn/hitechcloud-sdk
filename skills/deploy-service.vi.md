> [English](deploy-service.md) | **Tiếng Việt**

---
name: deploy-hitechcloud-service
description: Triển khai và quản lý dịch vụ HiTechCloud (VPS, hosting, GPU)
---

# Triển khai dịch vụ HiTechCloud

Dùng skill này khi người dùng muốn triển khai, quản lý hoặc khắc phục sự cố dịch vụ HiTechCloud.

## Các bước

1. **Xác định loại dịch vụ**: Hỏi người dùng loại dịch vụ (VPS, hosting, GPU, v.v.)
2. **Kiểm tra tính khả dụng**: Dùng API để kiểm tra sản phẩm và giá khả dụng
3. **Tạo dịch vụ**: Dùng endpoint đặt hàng để tạo dịch vụ
4. **Xác minh triển khai**: Kiểm tra trạng thái dịch vụ sau khi tạo

## Mẫu code

```python
from hitechcloud import HiTechCloud

client = HiTechCloud(api_key=os.environ['HITECHCLOUD_API_KEY'])

# Liệt kê sản phẩm khả dụng
products = client.cart.list_products()

# Tạo một đơn hàng
order = client.cart.create_order(product_id=123, config={
    'hostname': 'my-server',
    'os': 'ubuntu-22.04',
    'region': 'vn-hanoi'
})

# Kiểm tra trạng thái dịch vụ
service = client.services.get_service(id=order['service_id'])
print(f"Status: {service['status']}")
```
