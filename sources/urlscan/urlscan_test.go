package urlscan

import (
	"fmt"
	"testing"
)



func TestCheckSubdomain(t *testing.T){
	url:="https://urlscan.io/api/v1/search/?q=domain:entel.bo"
	obj := UrlScan{NameService: "example", Domain: "entel.bo", Url: url} 

	resp, err := obj.CheckSubdomain()
	if(err != nil){
		t.Fatal(err)
	}
	fmt.Println(resp)

}