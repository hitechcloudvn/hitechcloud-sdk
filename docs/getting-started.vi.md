> [English](getting-started.md) | **Tiếng Việt**

# Bắt đầu

## Cài đặt

Chọn ngôn ngữ của bạn và cài đặt SDK:

### Python
```bash
pip install hitechcloud
```

### Node.js
```bash
npm install @hitechcloud/sdk
```

### PHP
```bash
composer require hitechcloud/sdk
```

### Go
```bash
go get github.com/hitechcloudvn/hitechcloud-sdk/sdk/go
```

### Ruby
```bash
gem install hitechcloud
```

### Java
```xml
<dependency>
    <groupId>vn.hitechcloud</groupId>
    <artifactId>sdk</artifactId>
    <version>1.0.0</version>
</dependency>
```

### C#
```bash
dotnet add package HiTechCloud.SDK
```

## Xác thực

Mọi lời gọi API đều cần xác thực bằng API key hoặc Bearer token.

```python
from hitechcloud import HiTechCloud

client = HiTechCloud(api_key="your-api-key")
# hoặc / or
client = HiTechCloud(token="your-bearer-token")
```

Xem [Xác thực](authentication.vi.md) để biết chi tiết.
