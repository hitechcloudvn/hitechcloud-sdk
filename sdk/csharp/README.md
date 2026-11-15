> **English** | [Tiếng Việt](README.vi.md)

# HiTechCloud C# SDK

C# SDK for HiTechCloud UserAPI - Manage hosting, domains, DNS, SSL, billing, and more.

## Installation

### NuGet

```
dotnet add package HiTechCloud.SDK
```

Or add a package reference directly:

```xml
<PackageReference Include="HiTechCloud.SDK" Version="1.0.0" />
```

## Quick Start

```csharp
using HiTechCloud.SDK;

// With API token
var client = new HiTechCloudClient("your-api-token");

// Or login with email/password
var authClient = new HiTechCloudClient();
await authClient.LoginAsync("user@example.com", "password");

// List services
var services = await client.Services.ServicesAsync();

// Get domain details
var domain = await client.Domains.DomainAsync(123);

// Logout
await client.LogoutAsync();
```

## Configuration

```csharp
// Custom base URL and settings
var client = new HiTechCloudClient(
    "token",                  // API token (null for no auth)
    "https://custom.api.com", // Custom base URL
    true,                     // Sandbox (OTE) mode
    30,                       // Timeout (seconds)
    3                         // Max retries
);
```

## Resources

| Resource | Description | Methods |
|----------|-------------|---------|
| `Affiliate` | Affiliate: affiliate program management | 6 |
| `AffiliatesAdvanced` | AffiliatesAdvanced: advanced affiliate management with client_id | 14 |
| `AiFactory` | HiTechCloud AI Factory: templates, instances, SSH keys, volumes, clusters | 36 |
| `Auth` | Account Authentication: login, logout, token refresh, signup, password reset | 6 |
| `BareMetal` | Bare Metal & Colocation: reinstall, diagnostics, rescue, server info, IPs, VLANs, power | 21 |
| `Billing` | Billing & Contracts: balance, invoices, payment methods | 6 |
| `Cart` | Cart: shopping cart management | 6 |
| `CephS3` | Ceph S3: S3 connection, credentials, usage, buckets, sub-users, key rotation | 11 |
| `CloudGpu` | Cloud GPU: GPU instance management | 6 |
| `CloudInstance` | Cloud Instance: full VM lifecycle management | 25 |
| `CloudService` | Cloud Service: VM shutdown, reset, hostname, IP pools, networks, interfaces | 10 |
| `CloudVm` | Cloud Virtual Machine: upgrade, resources, templates, ISO, boot, PXE | 9 |
| `Collocation` | Collocation Services: PDU port management | 3 |
| `Contacts` | Contacts: manage account contacts | 5 |
| `Dns` | DNS: zones and records management for services | 8 |
| `DnsManage` | DNS Manage: domain DNS records, DNSSEC, nameserver registration | 10 |
| `Domains` | Domains: WHOIS, nameservers, EPP, registration, contact, email forwarding | 22 |
| `EmailMfaV2` | EmailMfaV2: email-based MFA management | 6 |
| `Hosting` | Hosting Services: reinstall, status, bandwidth, reset | 6 |
| `Ipam` | HiTechCloudIPAM: IP addresses, subnets, reverse DNS | 4 |
| `LocationV2` | LocationV2: countries, states, cities lookup | 4 |
| `NetworkServices` | Network Services: IP addresses and reverse DNS for services | 3 |
| `Notifications` | Notifications: manage notification preferences | 3 |
| `Partner` | Partner: partner program, tiers, pricing, customers, earnings, wallet, payouts, leads, referral, rates | 13 |
| `PasskeyV2` | PasskeyV2: passkey-based MFA management | 6 |
| `Pmg` | HiTechCloudPMG: mail filtering configuration and management | 4 |
| `Proxmox` | HiTechCloudProxmox: machine status, power, VMs, IPs, backups, snapshots, bandwidth | 10 |
| `ProxmoxBackup` | Proxmox Backup: PBS connection, credentials, usage, snapshots, groups, password, token | 9 |
| `ServiceStatus` | Service Status: list and manage service statuses | 2 |
| `Services` | Services: list, details, cancel, renew, billing cycle | 9 |
| `Ssl` | SSL Certificates: list, order, download certificates | 6 |
| `Support` | Support: tickets, departments, news, knowledgebase | 13 |
| `UrlShortener` | URL Shortener: shorten URLs, manage links, stats | 6 |
| `Users` | User Profile: view/update account details, logs | 3 |
| `Vcloudstack` | vCloudStack Public Cloud: rescue, unrescue, console, usage | 4 |
| `Virtualizor` | Virtualizor: VPS management via Virtualizor panel | 4 |
| `Vneidekyc` | VNeIDEKYC: eKYC identity verification and organization verification | 16 |
| `WillExpired` | WillExpired: expiring services and domains management | 10 |

## Error Handling

```csharp
try
{
    await client.Services.ServicesAsync();
}
catch (AuthenticationException e)
{
    Console.WriteLine("Auth error: " + e.Message);
}
catch (NotFoundException e)
{
    Console.WriteLine("Not found: " + e.Message);
}
catch (RateLimitException e)
{
    Console.WriteLine("Rate limited, retry after: " + e.RetryAfter + "s");
}
catch (ValidationException e)
{
    Console.WriteLine("Validation error: " + e.Message);
}
catch (ServerException e)
{
    Console.WriteLine("Server error: " + e.Message);
}
catch (HiTechCloudException e)
{
    Console.WriteLine("API error: " + e.Message);
}
```

## Testing

Unit tests are written with xUnit and live in `Tests/`.

```
dotnet test
```

## Requirements

- .NET 8.0+
- System.Text.Json 8.0.0+

## Links

- Repository: https://github.com/hitechcloudvn/hitechcloud-sdk
- API endpoint: https://api.hitechcloud.vn
- Sandbox (OTE): https://ote.hitechcloud.vn

## License

MIT
