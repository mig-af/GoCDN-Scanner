package cert

import (
	
	"gatoscanner/dnsmikis/requests"
	"gatoscanner/domain"
	"gatoscanner/funcs"
	"fmt"
	"io"
	"strings"
)


type CrtSh struct{
	NameService string 
	Domain string
	Url string
}



func (c *CrtSh) CheckSubdomain() (domain.SubDomains, error){
	//fmt.Println("Check subdomain of ")
	var result domain.SubDomains
	resp, err := requests.Get(c.Url)
	if(err != nil){
		return result, err
	}

	body, err:= io.ReadAll(resp.Body)
	if(err != nil){
		// fmt.Println("Error: ")
		// fmt.Println(err.Error())
		return result, err
		
	}
	//fmt.Println(string(body))
	subdomains, err := funcs.Parser[[]SubDomainCrt](string(body))
	if(err != nil){
		fmt.Println("Error parse: "+err.Error())
		return result, err
	}
	//subdomains := ParseData(string(body))
	for _,v := range subdomains{
		result.SubDomains = append(result.SubDomains, strings.ReplaceAll(v.CommonName, "*.", ""))
	}
	return result, nil


}


func (c *CrtSh) ServiceName()string{
	return c.NameService
}




// func ParseData( content string)[]SubDomainCrt{
// 	var subdomain []SubDomainCrt
// 	resp := json.Unmarshal([]byte(content), &subdomain)
// 	if(resp != nil){
// 		fmt.Println("Error ParserData")
// 		fmt.Println(resp.Error())
		
// 	}
// 	return subdomain

// }











