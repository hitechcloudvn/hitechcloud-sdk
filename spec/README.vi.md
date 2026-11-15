> [English](README.md) | **Tiếng Việt**

# Đặc tả API HiTechCloud

## Các tệp

- `openapi.json` — Đặc tả OpenAPI 3.0 (gốc của kho lưu trữ)
- `postman_collection.json` — Postman collection dùng để kiểm thử
- `insomnia_config.json` — Cấu hình workspace của Insomnia

## Kiểm tra tính hợp lệ

```bash
# Kiểm tra tính hợp lệ của đặc tả OpenAPI
npx @redocly/cli lint openapi.json
```

## Tạo mã nguồn

Mã SDK được sinh từ đặc tả OpenAPI:

```bash
# Tạo SDK Python
openapi-generator generate -i openapi.json -g python -o sdk/python

# Tạo SDK Node.js
openapi-generator generate -i openapi.json -g javascript -o sdk/nodejs
```

## Tài liệu

Tài liệu API tương tác có sẵn tại:
- [docs.hitechcloud.vn](https://docs.hitechcloud.vn)
- [Swagger UI](https://docs.hitechcloud.vn/swagger)
