> [English](CONTRIBUTING.md) | **Tiếng Việt**

# Đóng góp cho HiTechCloud SDK

Cảm ơn bạn đã quan tâm đóng góp!

## Thiết lập môi trường phát triển

1. Clone kho lưu trữ
2. Chọn SDK ngôn ngữ của bạn trong thư mục `sdk/`
3. Cài đặt phụ thuộc
4. Chạy kiểm thử

## Thêm ngôn ngữ mới

1. Tạo thư mục mới trong `sdk/` với tên ngôn ngữ
2. Triển khai client và tất cả lớp tài nguyên (resource)
3. Thêm kiểm thử
4. Cập nhật README.md (và README.vi.md) chính
5. Thêm workflow CI/CD trong `.github/workflows/` và `.gitlab-ci.yml`

## Phong cách code

- Tuân theo quy ước của từng ngôn ngữ
- Thêm docstring/JSDoc cho mọi phương thức công khai
- Theo cấu trúc resource hiện có
- Tài liệu song ngữ (tiếng Việt và tiếng Anh): giữ cặp `README.md` / `README.vi.md` đồng bộ

## Pull Request

1. Fork kho lưu trữ
2. Tạo nhánh tính năng
3. Thêm kiểm thử cho chức năng mới
4. Đảm bảo mọi kiểm thử đều pass (CI chạy trên GitHub Actions và GitLab CI)
5. Gửi pull request

## Thông điệp commit

- Dùng tiêu đề ngắn gọn, dạng mệnh lệnh (vd: `fix: ...`, `feat: ...`, `docs: ...`)
- Kết thúc thông điệp commit bằng dòng đồng tác giả khi phù hợp:

```
Co-Authored-By: pgt-vn <nguyenthanhan92f1@gmail.com>
Co-Authored-By: agent-hitechcloud <agent@hitechcloud.vn>
```

## Giấy phép

Khi đóng góp, bạn đồng ý rằng các đóng góp của mình sẽ được phát hành theo Giấy phép MIT.
