package rapiddns

import (
	"gatoscanner/dnsmikis/requests"
	"gatoscanner/domain"
	//"fmt"
	"io"
	"regexp"
	"strings"
)



type RapidDns struct{
	NameService string 
	Domain string 
	Url string 
}


func (r *RapidDns) CheckSubdomain()(domain.SubDomains, error){
	var result domain.SubDomains
	resp, err := requests.Get(r.Url)
	if(err != nil){
		return result, err
	}

	body, err:= io.ReadAll(resp.Body)
	if(err != nil){
		// fmt.Println("Error: ")
		// fmt.Println(err.Error())
		return result, err
	}
	
 	re :=  regexp.MustCompile(`<td>([a-zA-Z0-9._-]+\.[a-zA-Z]{2,})</td>`)
	listRegex := re.FindAllSubmatch([]byte(body), -1)
	for _, v := range listRegex{
		dominio := strings.ReplaceAll(string(v[0]), "</td>", "")
		dominio = strings.ReplaceAll(dominio, "<td>", "")
		result.SubDomains = append(result.SubDomains, dominio)
	}
	//fmt.Println(list)
	return result, nil
	
}

func (r *RapidDns)ServiceName()string{
	return r.NameService
}