> [English](README.md) | **Tiếng Việt**

# Công cụ phát triển

Công cụ để phát triển và kiểm thử HiTechCloud SDK.

## Các script

- `generate_sdk.py` — Sinh mã SDK từ đặc tả OpenAPI
- `validate_spec.py` — Kiểm tra tính hợp lệ của đặc tả OpenAPI
- `run_tests.sh` — Chạy kiểm thử cho SDK của mọi ngôn ngữ
- `check_endpoints.py` — Kiểm tra mọi endpoint đã được triển khai

## Cách sử dụng

```bash
# Kiểm tra tính hợp lệ của đặc tả
python3 tools/validate_spec.py

# Kiểm tra độ bao phủ endpoint
python3 tools/check_endpoints.py

# Chạy tất cả kiểm thử
bash tools/run_tests.sh
```
