package urlscan

import (
	"gatoscanner/dnsmikis/requests"
	"gatoscanner/domain"
	"gatoscanner/funcs"
	"fmt"
	"io"
)



type UrlScan struct{
	NameService string 
	Domain string 
	Url string
}


func (u *UrlScan) CheckSubdomain()(domain.SubDomains, error){
	var subdomains domain.SubDomains

	resp, err := requests.Get(u.Url)
	if(err != nil){
		fmt.Println("Error: ", err.Error())
		return subdomains, err
	}

	
	if(resp.StatusCode != 200){
		return subdomains, err 
	}

	body, _ := io.ReadAll(resp.Body)
	
	data, err := funcs.Parser[Results](string(body))
	if(err != nil){
		return subdomains, err
	}

	
	for _, v := range data.Result{
		subdomains.SubDomains = append(subdomains.SubDomains, v.Page.Domain)
	}

	return subdomains, nil

}

func (u *UrlScan) ServiceName()string{
	return u.NameService
}




