> [English](csharp.md) | **Tiếng Việt**

# Hướng dẫn SDK C#

## Cài đặt

```csharp
dotnet add package HiTechCloud.SDK
```

## Khởi tạo

```csharp
using HiTechCloud;

var client = new HiTechCloudClient("your-api-key");
```

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

## Kiểm thử

```csharp
using Xunit;
using HiTechCloud;

public class HiTechCloudTest {
    [Fact]
    public void TestClientInit() {
        var client = new HiTechCloudClient("test-key");
        Assert.NotNull(client);
    }
}
```

## Biến môi trường

| Biến | Mô tả |
|----------|-------------|
| `HITECHCLOUD_API_KEY` | API key của bạn |
| `HITECHCLOUD_BASE_URL` | Base URL tùy chỉnh (mặc định: `https://docs.hitechcloud.vn`) |

## Thêm ví dụ

Xem thư mục [`examples/csharp/`](../../examples/csharp/) để có các ví dụ đầy đủ.
