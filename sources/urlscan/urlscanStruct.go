package urlscan




type Results struct{
	Result []UrlScanObj `json:"results"`
}


type UrlScanObj struct{
	Task Task `json:"task"`
	Page Page `json:"page"`
}

type Task struct{
	Visibility string `json:"visibilty"`
	Method string `json:"method"`
	Domain string `json:"domain"`
	ApexDomain string `json:"apexDomain"`
}



type Page struct {
	Page string `json:"page"`
	Server string `json:"server"`
	Ip string `json:"ip"`
	Url string `json:"url"`
	Domain string `json:"domain"`
	ApexDomain string `json:"apexDomain"`
	AsnName string `json:"asnname"`
	Status string `json:"status"`
}

// {
//       "submitter": {},
//       "task": {
//         "visibility": "public",
//         "method": "api",
//         "domain": "coc.nube.entel.bo",
//         "apexDomain": "entel.bo",
//         "time": "2026-01-30T14:01:37.785Z",
//         "uuid": "019c0f35-774e-776f-b53f-f46e3b8f4b1d",
//         "url": "http://coc.nube.entel.bo/"
//       },
//       "stats": {
//         "uniqIPs": 2,
//         "uniqCountries": 1,
//         "dataLength": 1232931,
//         "encodedDataLength": 331888,
//         "requests": 11
//       },
//       "page": {
//         "country": "BO",
//         "server": "nginx/1.20.1",
//         "redirected": "same-domain",
//         "ip": "190.129.114.76",
//         "apexDomainAgeDays": 4606,
//         "language": "es",
//         "mimeType": "text/html",
//         "url": "https://coc.nube.entel.bo/login",
//         "tlsValidDays": 372,
//         "tlsAgeDays": 179,
//         "domainAgeDays": 0,
//         "tlsValidFrom": "2025-08-04T00:00:00.000Z",
//         "domain": "coc.nube.entel.bo",
//         "apexDomain": "entel.bo",
//         "asnname": "EMPRESA NACIONAL DE TELECOMUNICACIONES SOCIEDAD ANONIMA, BO",
//         "asn": "AS6568",
//         "tlsIssuer": "GeoTrust TLS RSA CA G1",
//         "status": "200"
//       },
//       "_id": "019c0f35-774e-776f-b53f-f46e3b8f4b1d",
//       "_score": null,
//       "sort": [
//         1769781697785,
//         "019c0f35-774e-776f-b53f-f46e3b8f4b1d"
//       ],
//       "result": "https://urlscan.io/api/v1/result/019c0f35-774e-776f-b53f-f46e3b8f4b1d/",
//       "screenshot": "https://urlscan.io/screenshots/019c0f35-774e-776f-b53f-f46e3b8f4b1d.png"
//     }

