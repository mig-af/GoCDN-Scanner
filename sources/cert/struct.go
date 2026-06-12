package cert




type SubDomainCrt struct{
	IssuerCaId int `json:"issuer_ca_id"`
	IssueName string `json:"issuer_name"`
	CommonName string `json:"common_name"`
	NameValue string `json:"name_value"`
	Id int `json:"id"`
	EntryTimestamp string `json:"entry_timestamp"`
	NotBefore string `json:"not_before"`
	NotAfter string `json:"not_after"`
	SerialNumber string `json:"serial_number"`
	ResultCount int `json:"result_count"`
}

//  {
//     "issuer_ca_id": 295816,
//     "issuer_name": "C=US, O=Let's Encrypt, CN=R12",
//     "common_name": "verify.nube.entel.bo",
//     "name_value": "verify.nube.entel.bo",
//     "id": 21491643764,
//     "entry_timestamp": "2025-10-01T17:52:38.798",
//     "not_before": "2025-10-01T16:54:08",
//     "not_after": "2025-12-30T16:54:07",
//     "serial_number": "06d6c2b558eb75236969562745ec346cb33f",
//     "result_count": 2
//   }










