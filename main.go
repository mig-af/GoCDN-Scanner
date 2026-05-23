package main

import (
	"gatoscanner/IPs"
	"gatoscanner/menu"
	"gatoscanner/style"
	"fmt"
	"os"
	"runtime"
	"gatoscanner/config"
	//
)




func main(){
	if(runtime.GOOS == "android"){
		config.Android = true
	}

	fmt.Println(style.Banner2)
	
	SaveinfFile := false
	option := os.Args
	if(len(option) < 2){
		
		fmt.Println("Invalid option, please type ./recondomain --help")
		fmt.Println()

		return
	}
	for _, v := range option{
		if(v == "--save" || v == "-s"){
			SaveinfFile = true
		}
	}

	//ips := &IPs.IpRanges{IPsPath: "./IPs/IPfiles"}
	ips := &IPs.IpRanges{}
	ips.Load()
	cdnList := ips.GetListCdn()

	switch option[1]{
	
	case "--cdn":
		
		if(len(option) < 3){
			fmt.Println()
			fmt.Println("Invalid <argument>, please type --help")
			fmt.Println()
			return
		}
		menu.CheckCdnOnly(&cdnList, option[2])	

		
	case "--subdomain":
		if(len(option) < 3){
			fmt.Println()
			fmt.Println("Invalid <argument>, please type --help")
			fmt.Println()
			return
		}
		menu.CheckAllSubdomain(&cdnList, os.Args[2], SaveinfFile)
		
	case "--help":
		fmt.Println(menu.Help())
		

	default:
		fmt.Println("Invalid option, please type ./recondomain --help")
		fmt.Println()
	}


	
	

	
}







