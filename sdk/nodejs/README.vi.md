> [English](README.md) | **Tiếng Việt**

# HiTechCloud Node.js SDK

[![npm version](https://img.shields.io/npm/v/hitechcloud-sdk)](https://www.npmjs.com/package/hitechcloud-sdk)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Node.js SDK chính thức cho [HiTechCloud UserAPI](https://docs.hitechcloud.vn).

## Cài đặt

```bash
npm install hitechcloud-sdk
```

## Bắt đầu nhanh

```javascript
const { HiTechCloud } = require("hitechcloud-sdk");

// Sử dụng thông tin đăng nhập
const client = new HiTechCloud({
  baseUrl: "https://api.hitechcloud.vn",
  username: "your@email.com",
  password: "your_password",
});

// Hoặc sử dụng token có sẵn
const client = new HiTechCloud({
  token: "your_api_token",
});

// Tất cả các phương thức đều trả về Promise
async function main() {
  const services = await client.services.list();
  const domains = await client.domains.list();
  const balance = await client.billing.getBalance();
  console.log(services, domains, balance);
}

main();
```

## Môi trường

| Môi trường | Base URL |
|-------------|----------|
| Production  | `https://api.hitechcloud.vn` |
| Sandbox (OTE) | `https://ote.hitechcloud.vn` |

## Các nhóm API

### Tài khoản & Xác thực
```javascript
await client.auth.login("email", "password");
await client.auth.logout();
await client.auth.refreshToken();
await client.auth.revokeToken();
await client.auth.passwordReset("email");
await client.auth.signup("email", "password", "firstName", "lastName");
```

### Hồ sơ người dùng
```javascript
await client.users.getDetails();
await client.users.updateDetails({ first_name: "New Name" });
await client.users.getLogs();
```

### Dịch vụ
```javascript
await client.services.list();
await client.services.get(serviceId);
await client.services.cancel(serviceId, "No longer needed");
await client.services.renew(serviceId);
```

### Tên miền
```javascript
await client.domains.list();
await client.domains.get(domainId);
await client.domains.whois("example.com");
await client.domains.checkAvailability("example.com");
await client.domains.getNameservers(domainId);
await client.domains.updateNameservers(domainId, { ns1: "ns1.example.com", ns2: "ns2.example.com" });
await client.domains.getEppCode(domainId);
```

### Quản lý DNS
```javascript
// Bản ghi DNS của tên miền
await client.dnsManage.getRecords(domainId);
await client.dnsManage.createRecord(domainId, { name: "www", type: "A", content: "1.2.3.4" });

// Zone DNS của dịch vụ
await client.dns.list();
await client.dns.addZone(serviceId, "example.com");
await client.dns.addRecord(serviceId, zoneId, { name: "@", type: "A", content: "1.2.3.4" });
```

### Chứng chỉ SSL
```javascript
await client.ssl.list();
await client.ssl.get(certificateId);
await client.ssl.download(certificateId);
await client.ssl.order({ product_id: 1, csr: "-----BEGIN CERTIFICATE REQUEST-----..." });
```

### Thanh toán
```javascript
await client.billing.getBalance();
await client.billing.listInvoices();
await client.billing.getInvoice(invoiceId);
await client.billing.applyCredit(invoiceId);
await client.billing.getPaymentMethods();
```

### Hỗ trợ
```javascript
await client.support.listTickets();
await client.support.createTicket({ subject: "Help", message: "...", department_id: 1 });
await client.support.getTicket("TKT-001");
await client.support.reply("TKT-001", "Additional info");
await client.support.close("TKT-001");
```

### Danh bạ
```javascript
await client.contacts.list();
await client.contacts.create({ first_name: "John", last_name: "Doe", email: "john@example.com" });
await client.contacts.update(contactId, { phone: "+84901234567" });
await client.contacts.delete(contactId);
```

### Giỏ hàng
```javascript
await client.cart.get();
await client.cart.add(1, { billing_cycle: "monthly" });
await client.cart.applyPromo("DISCOUNT20");
await client.cart.checkout({ payment_method: "sepay" });
```

### Thông báo
```javascript
await client.notifications.list();
await client.notifications.getPreferences();
await client.notifications.updatePreferences({ email: true, sms: false });
```

### Tiếp thị liên kết (Affiliate)
```javascript
await client.affiliate.getDetails();
await client.affiliate.getBalance();
await client.affiliate.listReferrals();
await client.affiliate.listTransactions();
```

### Cloud GPU
```javascript
await client.cloudGpu.list();
await client.cloudGpu.get(gpuId);
await client.cloudGpu.start(gpuId);
await client.cloudGpu.stop(gpuId);
await client.cloudGpu.restart(gpuId);
```

### Virtualizor
```javascript
await client.virtualizor.list();
await client.virtualizor.get(vpsId);
await client.virtualizor.start(vpsId);
await client.virtualizor.stop(vpsId);
```

## Xử lý lỗi

```javascript
const {
  HiTechCloudError,
  AuthenticationError,
  NotFoundError,
  RateLimitError,
  ValidationError,
  ServerError,
} = require("hitechcloud-sdk");

try {
  await client.domains.get(99999);
} catch (error) {
  if (error instanceof AuthenticationError) {
    console.log("Thông tin đăng nhập không hợp lệ hoặc token đã hết hạn");
  } else if (error instanceof NotFoundError) {
    console.log("Không tìm thấy tài nguyên");
  } else if (error instanceof RateLimitError) {
    console.log(`Vượt giới hạn tần suất. Thử lại sau ${error.retryAfter} giây`);
  } else if (error instanceof ValidationError) {
    console.log("Lỗi xác thực dữ liệu:", error.data);
  } else if (error instanceof ServerError) {
    console.log("Lỗi máy chủ, vui lòng thử lại sau");
  }
}
```

## Giấy phép

Giấy phép MIT - xem [LICENSE](../../LICENSE) để biết chi tiết.
