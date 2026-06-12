
# ReconDomain
```
    /\_____/\
   /  o   o  \
  ( ==  ^  == )
   )         (
  (           )
 ( (  )   (  ) )
(__(__)___(__)__)
  
```
> A passive subdomain scanner that detects CDN providers.

---

## What it does

- Enumerates subdomains using multiple passive sources
- Detects CDN provider for each subdomain (Cloudflare, Akamai, AWS Cloudfront, Google, Fastly, Sucuri, BunnyCDN)
- No active requests to the target — fully passive


---

## Installation

Precompiled binaries available for Linux and Android (Termux).

**Linux**
```bash
wget https://github.com/mig-af/GoCDN-Scanner/raw/refs/heads/main/gocdn
chmod +x gocdn
./gocdn
```

**Android (Termux)**
```bash
wget https://github.com/mig-af/GoCDN-Scanner/raw/refs/heads/main/gocdn-android 
chmod +x gocdn-android
./gocdn-android 
```

---

## Usage

```
gocdn [options] <arguments>

Options:
  --cdn <IP>          Scan all CDNs for a given IP
  --subdomain <DOMAIN>  Scan all CDNs and subdomains for a domain
  --save | -s        Save results to ./subdomains-your-domain.txt
  --help             Show help
```

**Examples**
```bash
./gocdn --cdn 123.123.123.123
./gocdn --subdomain mydomain.com
./gocdn --subdomain mydomain.net --save
```

---

## Supported CDNs

| CDN|
|---|
| Cloudflare |
 Akamai |
| AWS Cloudfront |
 Google | 
| Fastly | 
Sucuri | 
| BunnyCDN |

---



## Author

- [@mig-af](https://www.github.com/mig-af)

