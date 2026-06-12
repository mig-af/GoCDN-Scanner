package robtex


import (
	"testing"
)


func TestChangeDomainFormat(t *testing.T){
	casos := []struct{
		dominioOriginal, esperado string
	}{
		{"www.example.com", "com/example/www"},
		{"test.net", "net/test"},
		{"www.sub.example.com", "com/sub/example/www"},
	}

	for _, v := range casos{
		resp := ChangeDomainFormat(v.dominioOriginal)
		if( resp != v.esperado){
			t.Errorf("%s no es igual a %s", resp, v.esperado)
		}
	}


}


func TestCheckSubdomain(t *testing.T){
	domain := "whatsapp.com"
	t.Log(ChangeDomainFormat(domain))
	url := "https://robtex.com/en/dns-lookup/"+ChangeDomainFormat(domain)
	rbt := &Robtext{NameService: "robtext", Domain: domain, Url:url}

	resp, err := rbt.CheckSubdomain() 
	if(err != nil){
		t.Log(err)
	}
	t.Log(resp)
}




