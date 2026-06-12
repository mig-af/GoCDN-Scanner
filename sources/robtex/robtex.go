package robtex

import (
	"errors"
	"gatoscanner/config"
	"time"
	//"gatoscanner/dnsmikis/requests"
	"gatoscanner/domain"
	"net/http"
	"strings"
	"github.com/antchfx/htmlquery"
)



type Robtext struct{
	NameService string 
	Domain string 
	Url string
}



func (r *Robtext) CheckSubdomain()(domain.SubDomains, error){
	var subdomains domain.SubDomains

	dialer := config.ConfigDialerAndResolver()
	client := &http.Client{
		Timeout: 20 * time.Second,
		Transport: &http.Transport{
			DialContext: dialer.DialContext,
		},
	}


	html, err := htmlquery.LoadURLWithClient(r.Url, client)

	if(err != nil){
		return subdomains, err
	}

	div := htmlquery.Find(html, "//h2[@id='subdomains']/following-sibling::div[1]")
	if(len(div) == 0){
		return subdomains, errors.New("Robtex: subdomain not found")
	}
	aTarget := htmlquery.Find(div[0], "//a//text()")
	
	for _, v := range aTarget{
		subdomains.SubDomains = append(subdomains.SubDomains,  v.Data+r.Domain)
		//fmt.Println(v.Data+r.Domain)
	}
	return subdomains, nil

}

func (r *Robtext) ServiceName()string{
	return r.NameService
}



func ChangeDomainFormat(domain string)string{
	NewFormat := []string{}
	dmain := strings.Split(domain, ".")
	
	first := dmain[0]
	last := dmain[len(dmain)-1:][0]
	NewFormat = append(NewFormat, last)

	for i:=1; i<len(dmain)-1; i++{
		NewFormat = append(NewFormat, dmain[i])
	}
	NewFormat = append(NewFormat, first)
	return strings.Join(NewFormat, "/")
	
}
