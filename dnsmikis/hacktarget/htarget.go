package hacktarget

import (
	"errors"
	"gatoscanner/dnsmikis/requests"
	"gatoscanner/domain"
	"fmt"
	"io"
	"strings"
)




type Htarget struct{
	NameService string
	Domain string 
	Url string
}




func (h *Htarget) CheckSubdomain()(domain.SubDomains, error){

	var subdomains domain.SubDomains

	resp, err := requests.Get(h.Url)
	if(err != nil){
		fmt.Println("Error: ", err.Error())
		return subdomains, err
	}

	
	if(resp.StatusCode != 200){
		
		return subdomains, err 
	}

	body, _ := io.ReadAll(resp.Body)
	list := [][]string{}
	for _, v := range strings.Fields(string(body)){
		
		list = append(list, strings.Split(v, ","))
	}	
	
	for _,x := range list{
		subdomains.SubDomains = append(subdomains.SubDomains, x[0])
	}

	if( subdomains.SubDomains[0] == "API" || subdomains.SubDomains[0] == "error"){
		return subdomains, errors.New("Bad response")
	}
	return subdomains, nil
}



func (t *Htarget) ServiceName()string{
	return t.NameService
}

