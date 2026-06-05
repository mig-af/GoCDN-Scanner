package funcs

import (
	"context"
	"encoding/json"
	"fmt"
	"gatoscanner/config"
	//"math/rand"
	"net"
	"os"
	//"time"
)

//---------------
func DeleteRepeat(list []string)[]string{

	// fmt.Println("Limpiando lista?---")
	// fmt.Println(list)
	var newList []string

	for i:=0; i<len(list); i++{
		num := list[i]
		for x:=0; x<len(list); x++{
			if num != list[x] && elementInList(newList, num) == false{
				newList = append(newList, num)
			}
		}
	}

	return  newList

}


func elementInList(list []string, element string)bool{

	for i:=0; i<len(list); i++{
		if(element == list[i]){
			return true
		}
	}
	return false

}

//------------------


func CheckIp(url string, onlyIpv4 bool, resolver *net.Resolver)(*[]net.IP, error){
	var (
		resp []net.IP
		err error
	)
	
	if(config.Android){
		resp, err = resolver.LookupIP(context.Background(), "ip", url)

	}else{
		resp, err = net.LookupIP(url)

	}
	if(err != nil){
		//time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		// fmt.Printf("\r %s", err.Error())
		return &[]net.IP{}, err
	}
	if(onlyIpv4){
		
		var ipv4s []net.IP 

		//fmt.Println(len(resp))
		for i:=0; i<len(resp);i++ {
			if(len(resp[i]) != net.IPv6len){
				ipv4s = append(ipv4s, resp[i])
			}
		}
		return &ipv4s, nil
	}

	return &resp, nil
}




func CheckNs(url string, resolver *net.Resolver)[]string{
	var (
		resp []*net.NS
		err error
	) 

	if(config.Android){
		
		resp, err = resolver.LookupNS(context.Background(), url)
	}else{
		resp, err = net.LookupNS(url)

	}

	if(err != nil){
		//fmt.Println(err.Error())
		return []string{}
	}
	list := []string{}
	for _, v := range resp{
		list = append(list, v.Host)
	}
	//fmt.Println(list)
	return list
}



func CheckCdn(cdnName string,  ip net.IP, rangeIps []string)bool{

	if(cdnName == "bunnycdn"){
		return false // Seguridad, BUNNYCDN maneja IPs puras no rangos CIDR si el servicio es bunny esta funcion siempre devolvera false
		//Para bunnyCdn hay otra funcion creada
	}
	for _, v := range(rangeIps){

		
		_, ipnet, err := net.ParseCIDR(v)
		if(err != nil){
			fmt.Println("Error checkcnd: "+ v + err.Error())
			return false
		}

		ipp := net.ParseIP(ip.String())
		if(ipnet.Contains(ipp)){
			return true
		}


	}
	return false
	
}


func CheckBunnyCDN(ip net.IP, cdnRange[]string)bool{

	for _,v := range cdnRange{
		if(ip.String() == v){
			return true
		}
	}
	return false
}







//Dividimos un array en pequenos chunks
//Lo convertimos en un array bidimensional [[1,2,3], [4,5,6]]

func SplitArray(list []string, numSplit int)[][]string {
	newArray := [][]string{}


	nElementos := len(list) / numSplit
	resto := len(list) % numSplit
	
	
	//fmt.Println(list)
	indice := 0 
	final := nElementos
	for i :=1; i<=numSplit; i++{
		//time.Sleep(1 * time.Second)
		fin := final * i
		chunk := list[indice : fin]
	
		if(resto != 0 && i == numSplit){ //Si hay restantes (resto) y es el ultimo chunk se anade los elementos a este ultimo 
			for _, restante := range list[fin:]{
				chunk = append(chunk, restante)
			// 	fmt.Println(restante)
			 }
			
		}

		newArray = append(newArray, chunk)
		//fmt.Println(chunk)

		indice = fin
	}
	return newArray

}


//Parsea el json string hacia un struct 
//[T any] es el tipo de dato T que manejara y devolvera
func Parser[T any](content string)(T, error){
	var data T
	resp := json.Unmarshal([]byte(content), &data)
	if(resp!= nil){
		fmt.Println("Error marshal: ", resp.Error())
		return data, resp
		
	}
	return data, nil
}



func Save(filename string, content *string){
	// var data string

	// for _, v := range content{
	// 	//info := fmt.Sprintf("------------------------------\n%s %s %s", v.Name, v.Ip, v.Cdns)
	// 	data += string(info)
	// }
	resp, err := os.OpenFile(filename, os.O_CREATE | os.O_RDWR | os.O_APPEND, 0644)
	if(err != nil){
		fmt.Println(err)
	}
	defer resp.Close()

	resp.Write([]byte(*content))

}



func Reverse[T any](list []T)[]T{
	NewList := []T{}

	for i:=1 ; i<=len(list); i++{
		dato := list[len(list)-i:len(list)-i+1]
		//fmt.Println(dato)
		NewList = append(NewList, dato[0])
	}
	return NewList

}










