> [English](README.md) | **Tiếng Việt**

# Ví dụ Csharp

## Sử dụng cơ bản

```csharp
// Liệt kê tất cả dịch vụ
var services = await client.Services.ListServicesAsync();
foreach (var svc in services.Data)
{
    Console.WriteLine($"{svc.Id}: {svc.Name}");
}

// Lấy một dịch vụ cụ thể
var service = await client.Services.GetServiceAsync(123);
Console.WriteLine(service);

// Các thao tác nguồn
await client.Services.StartServiceAsync(123);
await client.Services.StopServiceAsync(123);
await client.Services.RebootServiceAsync(123);
```

## Cách chạy

```bash
export HITECHCLOUD_API_KEY=your-api-key
dotnet run
```

## Thêm ví dụ

- `basic_usage.cs` — Liệt kê, xem và quản lý dịch vụ
- `error_handling.cs` — Xử lý lỗi API một cách nhẹ nhàng
- `pagination.cs` — Phân trang kết quả
