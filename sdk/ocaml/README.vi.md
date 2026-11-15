> [English](README.md) | **Tiếng Việt**

# HiTechCloud hitechcloud-sdk — OCaml SDK

> **Trạng thái:** Dự kiến — hoan nghênh đóng góp!

## Cài đặt

```bash
opam install hitechcloud_sdk
```

## Tài nguyên

| Tài nguyên | Mô tả | Phương thức |
|----------|-------------|---------|
| `affiliate` | Affiliate: quản lý chương trình affiliate | 6 |
| `affiliatesAdvanced` | AffiliatesAdvanced: quản lý affiliate nâng cao với client_id | 14 |
| `aiFactory` | HiTechCloud AI Factory: templates, instances, khóa SSH, volumes, clusters | 36 |
| `auth` | Xác thực tài khoản: đăng nhập, đăng xuất, làm mới token, đăng ký, đặt lại mật khẩu | 6 |
| `bareMetal` | Bare Metal & Colocation: cài đặt lại, chẩn đoán, rescue, thông tin server, IPs, VLANs, nguồn | 21 |
| `billing` | Thanh toán & Hợp đồng: số dư, hóa đơn, phương thức thanh toán | 6 |
| `cart` | Cart: quản lý giỏ hàng | 6 |
| `cephS3` | Ceph S3: kết nối S3, credentials, lưu lượng sử dụng, buckets, sub-users, xoay khóa | 11 |
| `cloudGpu` | Cloud GPU: quản lý GPU instance | 6 |
| `cloudInstance` | Cloud Instance: quản lý toàn bộ vòng đời VM | 25 |
| `cloudService` | Cloud Service: tắt VM, reset, hostname, IP pools, networks, interfaces | 10 |
| `cloudVm` | Cloud Virtual Machine: nâng cấp, tài nguyên, templates, ISO, boot, PXE | 9 |
| `collocation` | Dịch vụ Collocation: quản lý cổng PDU | 3 |
| `contacts` | Contacts: quản lý liên hệ của tài khoản | 5 |
| `dns` | DNS: quản lý zones và records cho các dịch vụ | 8 |
| `dnsManage` | DNS Manage: DNS records của domain, DNSSEC, đăng ký nameserver | 10 |
| `domains` | Domains: WHOIS, nameservers, EPP, đăng ký, liên hệ, chuyển tiếp email | 22 |
| `emailMfaV2` | EmailMfaV2: quản lý MFA dựa trên email | 6 |
| `hosting` | Dịch vụ Hosting: cài đặt lại, trạng thái, băng thông, reset | 6 |
| `ipam` | HiTechCloudIPAM: địa chỉ IP, subnets, reverse DNS | 4 |
| `locationV2` | LocationV2: tra cứu quốc gia, bang/tỉnh, thành phố | 4 |
| `networkServices` | Dịch vụ Network: địa chỉ IP và reverse DNS cho các dịch vụ | 3 |
| `notifications` | Notifications: quản lý tùy chọn thông báo | 3 |
| `partner` | Partner: chương trình đối tác, tiers, giá, khách hàng, thu nhập, ví, payouts, leads, referral, tỷ giá | 13 |
| `passkeyV2` | PasskeyV2: quản lý MFA dựa trên passkey | 6 |
| `pmg` | HiTechCloudPMG: cấu hình và quản lý lọc email | 4 |
| `proxmox` | HiTechCloudProxmox: trạng thái máy, nguồn, VMs, IPs, backups, snapshots, băng thông | 10 |
| `proxmoxBackup` | Proxmox Backup: kết nối PBS, credentials, lưu lượng sử dụng, snapshots, groups, mật khẩu, token | 9 |
| `serviceStatus` | Service Status: liệt kê và quản lý trạng thái dịch vụ | 2 |
| `services` | Services: danh sách, chi tiết, hủy, gia hạn, chu kỳ thanh toán | 9 |
| `ssl` | Chứng chỉ SSL: liệt kê, đặt mua, tải chứng chỉ | 6 |
| `support` | Support: tickets, phòng ban, tin tức, knowledgebase | 13 |
| `urlShortener` | URL Shortener: rút gọn URL, quản lý links, thống kê | 6 |
| `users` | Hồ sơ người dùng: xem/cập nhật thông tin tài khoản, logs | 3 |
| `vcloudstack` | vCloudStack Public Cloud: rescue, unrescue, console, lưu lượng sử dụng | 4 |
| `virtualizor` | Virtualizor: quản lý VPS qua bảng điều khiển Virtualizor | 4 |
| `vneidekyc` | VNeIDEKYC: xác minh danh tính eKYC và xác minh tổ chức | 16 |
| `willExpired` | WillExpired: quản lý dịch vụ và domain sắp hết hạn | 10 |

## Giấy phép

MIT
