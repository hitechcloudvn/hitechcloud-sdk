> [English](README.md) | **Tiếng Việt**

# HiTechCloud Go SDK

SDK Go cho HiTechCloud UserAPI - Quản lý hosting, tên miền, DNS, SSL, thanh toán và nhiều hơn nữa.

## Cài đặt

```bash
go get github.com/hitechcloudvn/hitechcloud-sdk/sdk/go
```

## Bắt đầu nhanh

```go
package main

import (
    "fmt"
    "log"

    hitechcloud "github.com/hitechcloudvn/hitechcloud-sdk/sdk/go"
)

func main() {
    // Đăng nhập bằng email/mật khẩu
    client, err := hitechcloud.NewClient(
        hitechcloud.WithEmail("user@example.com", "password"),
    )
    if err != nil {
        log.Fatal(err)
    }
    defer client.Logout()

    // Hoặc bằng API token
    client, err = hitechcloud.NewClient(
        hitechcloud.WithToken("your-api-token"),
    )
    if err != nil {
        log.Fatal(err)
    }

    // Liệt kê dịch vụ
    services, err := client.Services.List(nil)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(services)
}
```

## Tùy chọn cấu hình

```go
client, err := hitechcloud.NewClient(
    hitechcloud.WithToken("token"),           // API token
    hitechcloud.WithEmail("email", "pass"),   // Đăng nhập bằng email/mật khẩu
    hitechcloud.WithBaseURL("https://..."),   // URL gốc tùy chỉnh
    hitechcloud.WithTimeout(30*time.Second),  // Thời gian chờ của request
    hitechcloud.WithMaxRetries(3),            // Số lần thử lại tối đa
    hitechcloud.WithSandbox(),                // Sử dụng môi trường sandbox
)
```

## Tài nguyên

### Xác thực
```go
client.Auth.Login(email, password)
client.Auth.Logout()
client.Auth.RefreshToken()
client.Auth.RevokeToken()
client.Auth.PasswordReset(email)
client.Auth.Signup(email, password, firstName, lastName)
```

### Người dùng
```go
client.Users.GetDetails()
client.Users.UpdateDetails(data)
client.Users.GetLogs()
```

### Dịch vụ
```go
client.Services.List(params)
client.Services.Get(serviceID)
client.Services.ListMethods()
client.Services.Cancel(serviceID, reason)
client.Services.GetLabel(serviceID)
client.Services.SetLabel(serviceID, label)
client.Services.Renew(serviceID)
client.Services.ListBillingCycles(serviceID)
client.Services.ChangeBillingCycle(serviceID, billingCycle)
```

### Tên miền
```go
client.Domains.List(params)
client.Domains.Get(domainID)
client.Domains.GetByName(name)
client.Domains.Whois(domain)
client.Domains.WhoisLookup(domain)
client.Domains.CheckAvailability(domain)
client.Domains.GetNameservers(domainID)
client.Domains.UpdateNameservers(domainID, ns)
client.Domains.GetEppCode(domainID)
client.Domains.Sync(domainID)
client.Domains.GetLock(domainID)
client.Domains.UpdateLock(domainID, reglock)
client.Domains.UpdateIDProtection(domainID, idprotection)
client.Domains.GetContact(domainID)
client.Domains.UpdateContact(domainID, data)
client.Domains.GetEmailForwarding(domainID)
client.Domains.UpdateEmailForwarding(domainID, entries)
client.Domains.GetAvailableTlds()
client.Domains.Order(domain, years, data)
client.Domains.Renew(domainID, years)
client.Domains.GetTldForm(tldID)
client.Domains.GetDocuments(domainID)
```

### Quản lý DNS
```go
client.DnsManage.GetRecords(domainID)
client.DnsManage.CreateRecord(domainID, data)
client.DnsManage.UpdateRecord(domainID, recordID, data)
client.DnsManage.DeleteRecord(domainID, recordID)
client.DnsManage.GetRecordTypes()
client.DnsManage.GetDnssecFlags()
client.DnsManage.GetDnssecRecords(domainID)
client.DnsManage.CreateDnssecRecord(domainID, data)
client.DnsManage.DeleteDnssecRecord(domainID, recordID)
client.DnsManage.RegisterNameserver(domainID, hostname, ips)
```

### DNS
```go
client.Dns.List(params)
client.Dns.ListForService(serviceID)
client.Dns.AddZone(data)
client.Dns.GetZone(zoneID)
client.Dns.DeleteZone(zoneID)
client.Dns.AddRecord(zoneID, data)
client.Dns.EditRecord(zoneID, recordID, data)
client.Dns.DeleteRecord(zoneID, recordID)
```

### SSL
```go
client.Ssl.List(params)
client.Ssl.Get(sslID)
client.Ssl.Download(sslID)
client.Ssl.ListAvailable(params)
client.Ssl.Order(data)
client.Ssl.ListServerSoftware()
```

### Thanh toán
```go
client.Billing.GetBalance()
client.Billing.ListInvoices(params)
client.Billing.GetInvoice(invoiceID)
client.Billing.ApplyCredit(amount)
client.Billing.GetPaymentMethods()
client.Billing.GetPaymentFees()
```

### Hỗ trợ
```go
client.Support.ListDepartments()
client.Support.ListTickets(params)
client.Support.GetTicket(ticketID)
client.Support.CreateTicket(data)
client.Support.ReplyTicket(ticketID, message)
client.Support.CloseTicket(ticketID)
client.Support.ListNews(params)
client.Support.GetNews(newsID)
client.Support.ListKbCategories()
client.Support.ListKbArticles(params)
client.Support.GetKbArticle(articleID)
client.Support.SearchKb(query)
client.Support.ListAnnouncements(params)
```

### Liên hệ
```go
client.Contacts.List(params)
client.Contacts.Get(contactID)
client.Contacts.Create(data)
client.Contacts.Update(contactID, data)
client.Contacts.Delete(contactID)
```

### Giỏ hàng
```go
client.Cart.Get()
client.Cart.Add(data)
client.Cart.Remove(itemID)
client.Cart.ApplyPromo(code)
client.Cart.Checkout(data)
client.Cart.GetPaymentMethods()
```

### Thông báo
```go
client.Notifications.List(params)
client.Notifications.GetPreferences()
client.Notifications.UpdatePreferences(data)
```

### Tiếp thị liên kết
```go
client.Affiliate.GetDetails()
client.Affiliate.GetBalance()
client.Affiliate.GetPayout()
client.Affiliate.UpdatePayout(data)
client.Affiliate.ListReferrals(params)
client.Affiliate.ListTransactions(params)
```

### Cloud GPU
```go
client.CloudGpu.List(params)
client.CloudGpu.Get(gpuID)
client.CloudGpu.Start(gpuID)
client.CloudGpu.Stop(gpuID)
client.CloudGpu.Restart(gpuID)
client.CloudGpu.GetConsole(gpuID)
```

### Virtualizor
```go
client.Virtualizor.List(params)
client.Virtualizor.Get(vpsID)
client.Virtualizor.Start(vpsID)
client.Virtualizor.Stop(vpsID)
```

## Xử lý lỗi

```go
import hitechcloud "github.com/hitechcloudvn/hitechcloud-sdk/sdk/go"

result, err := client.Services.List(nil)
if err != nil {
    switch e := err.(type) {
    case *hitechcloud.AuthenticationError:
        fmt.Println("Auth error:", e.Message)
    case *hitechcloud.NotFoundError:
        fmt.Println("Not found:", e.Message)
    case *hitechcloud.RateLimitError:
        fmt.Println("Rate limited, retry after:", e.RetryAfter, "seconds")
    case *hitechcloud.ValidationError:
        fmt.Println("Validation error:", e.Message)
    case *hitechcloud.ServerError:
        fmt.Println("Server error:", e.Message)
    default:
        fmt.Println("Error:", err)
    }
}
```

## Giấy phép

MIT
