> [English](java.md) | **Tiếng Việt**

# Hướng dẫn SDK Java

## Cài đặt

```java
<dependency>
    <groupId>vn.hitechcloud</groupId>
    <artifactId>sdk</artifactId>
    <version>1.0.0</version>
</dependency>
```

## Khởi tạo

```java
import vn.hitechcloud.HiTechCloud;

HiTechCloud client = new HiTechCloud("your-api-key");
```

## Sử dụng cơ bản

```java
// Liệt kê tất cả dịch vụ
var services = client.services().listServices();
System.out.println(services);

// Lấy một dịch vụ cụ thể
var service = client.services().getService(123);
System.out.println(service);

// Các thao tác nguồn
client.services().startService(123);
client.services().stopService(123);
client.services().rebootService(123);
```

## Kiểm thử

```java
import org.junit.Test;
import static org.junit.Assert.*;
import vn.hitechcloud.HiTechCloud;

public class HiTechCloudTest {
    @Test
    public void testClientInit() {
        HiTechCloud client = new HiTechCloud("test-key");
        assertNotNull(client);
    }
}
```

## Biến môi trường

| Biến | Mô tả |
|----------|-------------|
| `HITECHCLOUD_API_KEY` | API key của bạn |
| `HITECHCLOUD_BASE_URL` | Base URL tùy chỉnh (mặc định: `https://docs.hitechcloud.vn`) |

## Thêm ví dụ

Xem thư mục [`examples/java/`](../../examples/java/) để có các ví dụ đầy đủ.
