> [English](CHANGELOG.md) | **Tiếng Việt**

# Nhật ký thay đổi

Mọi thay đổi đáng chú ý của HiTechCloud SDK sẽ được ghi lại trong tệp này.

Định dạng dựa trên [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
và dự án tuân theo [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.3.0] - 2026-11-16

### Thêm
- Chuẩn hoá song ngữ Việt/Anh trên toàn bộ kho lưu trữ (README, tài liệu, hướng dẫn, tài liệu SDK, ví dụ, agent, skill)
- Pipeline CI/CD GitLab (`.gitlab-ci.yml`) với kiểm thử và phát hành theo tag
- `workflow_dispatch` cho workflow Release để phát hành cho các tag đã có
- README còn thiếu cho SDK C# (`sdk/csharp/README.md` + `README.vi.md`)
- Chú thích code song ngữ trong các ví dụ

### Sửa
- Ghim `maven-resources-plugin` phiên bản `3.3.1` trong `sdk/java/pom.xml` (phiên bản mặc định `3.5.0` không tồn tại trên Maven Central) — khôi phục job CI Java
- Sửa đường dẫn phát triển trong README (`cd sdk/python`, `cd sdk/nodejs`)

## [1.2.0] - 2026-11-15

### Thay đổi
- Chuyển kho lưu trữ sang tổ chức `hitechcloudvn` trên GitHub và GitLab
- Chuẩn hoá đường dẫn module Go thành `github.com/hitechcloudvn/hitechcloud-sdk/sdk/go`
- Trỏ tham chiếu SDK Julia, Objective-C và Swift về monorepo
- Chuẩn hoá tác giả commit thành Nguyễn Thanh An (thanhan92f1)

## [1.1.0] - 2026-11-15

### Thay đổi
- Đồng bộ mã nguồn User API mới nhất
- Thêm theo dõi artifact build cho pipeline phát hành

## [1.0.0] - 2024-12-01

### Thêm
- Phát hành đầu tiên với 7 SDK ngôn ngữ (Python, Node.js, PHP, Go, Ruby, Java, C#)
- 39 module tài nguyên bao phủ toàn bộ endpoint User API
- Đặc tả OpenAPI 3.0
- Tài liệu đầy đủ
- Pipeline CI/CD cho mọi ngôn ngữ
- Ví dụ cho mọi ngôn ngữ được hỗ trợ
- Cấu hình agent và máy chủ MCP

[1.3.0]: https://github.com/hitechcloudvn/hitechcloud-sdk/compare/v1.2.0...v1.3.0
[1.2.0]: https://github.com/hitechcloudvn/hitechcloud-sdk/compare/v1.1.0...v1.2.0
[1.1.0]: https://github.com/hitechcloudvn/hitechcloud-sdk/compare/v1.0.0...v1.1.0
[1.0.0]: https://github.com/hitechcloudvn/hitechcloud-sdk/releases/tag/v1.0.0
