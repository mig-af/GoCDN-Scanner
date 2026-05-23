package funcs

import (
	//"fake/IPs"
	"gatoscanner/IPs"
	"fmt"
	"net"
	"testing"
)




func TestCheckIp(t *testing.T){
	url := "money.tigo.com.bo"
	r, _ := CheckIp(url, true, nil)
	fmt.Println("IP: ",r)
	fmt.Println("NS: ",CheckNs(url, nil))

}


func TestCheckCdn(t *testing.T){
	ranges := &IPs.IpRanges{IPsPath: "../IPs/"}
	ranges.Load()
	cdnForTest := "cloudflare"


	var cdn []string
	for _, v := range ranges.List{
		//fmt.Println(v.GetName())
		if(v.GetName() == cdnForTest){
			cdn = append(cdn, v.GetIps()...)
		}
	} 

	ji := CheckCdn(cdnForTest, net.IPv4(79,127,213,212), cdn)
	if(ji){
		fmt.Println("is: ", cdnForTest)
		
	}else{
		fmt.Println("No")
	}
	
	

}


func TestCheckBunnyCDN(t *testing.T){

	ranges := &IPs.IpRanges{IPsPath: "../IPs/"}
	ranges.Load()

	var bunnyIpsList []string
	for _,v := range ranges.GetListCdn(){
		if(v.GetName() == "bunnycdn"){
			bunnyIpsList = append(bunnyIpsList, v.GetIps()...)
		}
	}
	
	resp := CheckBunnyCDN(net.IPv4(94,20,154,22), bunnyIpsList)
	if(resp){
		t.Log("IS BUNNY CDN")
	}else{
		t.Log("NOT BUNNY CDN")
	}

}

func TestCheckNs(t *testing.T){

	er := CheckNs("181.115.186.67", nil)
	t.Log(er)
}

func TestSplitArray(t *testing.T){
	//lista := []int{23, 2,23,4,23,1,43,76,98,34,00,4, 23, 656, 12, 24, 56, 12, 67, 12, 7, 12, 7, 23, 56, 23, 5, 3, 43, 12, 34, 56, 76, 3, 45, 2, 6, 2}
	//s:=SplitArray(lista, 2)

	//fmt.Println(s)
}


func TestReverse(t *testing.T){
	letras := []string{"uno", "dos", "tres", "cuatro"}
	resp := Reverse(letras)
	t.Log(resp) //[cuatro, tres, dos, uno]


}