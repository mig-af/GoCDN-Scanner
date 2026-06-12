package menu

import (
	"fmt"
	"gatoscanner/IPs"
	"gatoscanner/config"
	"gatoscanner/sources"
	"gatoscanner/sources/cert"
	"gatoscanner/sources/hacktarget"
	"gatoscanner/sources/rapiddns"
	"gatoscanner/sources/urlscan"
	"gatoscanner/sources/robtex"
	"gatoscanner/domain"
	"gatoscanner/funcs"
	"gatoscanner/style"
	"net"
	"sync"
	"time"
)




var Resolver *net.Resolver = config.ConfigResolver() //solo para android
	



//---------Buscar cdn de una sola ip
func CheckCdnOnly(ips *[]IPs.Cdn,ip string){
	fmt.Println("")
	fmt.Println("SEARCHING CDN >:",  ip)
	time.Sleep(1 * time.Second)
	for _,v := range *ips{
		b := funcs.CheckBunnyCDN(net.ParseIP(ip), v.GetIps())
		if(b){
			fmt.Println("bunnycdn")
			return
		}
		

		cdn := funcs.CheckCdn(v.GetName(), net.ParseIP(ip), v.GetIps())
		if(cdn){
			fmt.Println()
			fmt.Println("CDN: "+style.GREEN + v.GetName() + style.END)
			fmt.Println()
			return
		}
	}
	fmt.Println("")
	fmt.Println("CDN NOT FOUND")
	fmt.Println("\nIP: ", ip)
}



func CheckAllSubdomain(cdnList *[]IPs.Cdn , dominio string, savefile bool){
	
	
	if(len(funcs.CheckNs(dominio, Resolver)) == 0){
		fmt.Printf(style.RED + "Domain: %s not found\n"+style.END , dominio  )
		return
	}

	//-------Urls------------

	urlCrt := fmt.Sprintf("https://crt.sh/?q=%s&output=json", dominio)
	urlHtarget := fmt.Sprintf("https://api.hackertarget.com/hostsearch/?q=%s", dominio)
	urlUrlScanio := fmt.Sprintf("https://urlscan.io/api/v1/search/?q=domain:%s", dominio)
	urlRapidDns := fmt.Sprintf("https://rapiddns.io/subdomain/%s?full=1", dominio)
	urlRobtex := fmt.Sprintf("https://robtex.com/en/dns-lookup/%s", robtex.ChangeDomainFormat(dominio))
	

	
	//----------------------Services-------------------------------------
	var subDomainsStrings []string

	//-----------------CRT.SH--------------------------------
	crtSh := &cert.CrtSh{NameService:"crt.sh", Domain:dominio, Url: urlCrt}
	crt, err := ScanSubdomain(crtSh)
	if(err != nil){
		
		intentos := 10
		
		respCrt := make(chan domain.SubDomains)
		go func(){
			ok := false
			for v := range intentos-1{
				fmt.Printf("\rTrying: %s%d%s", style.Randcolor() ,v ,style.END)
				time.Sleep(1 * time.Second)
				crt, err = ScanSubdomain(crtSh)
				if(len(crt.SubDomains) > 0){
					respCrt <- crt
					fmt.Printf("\r%s", ".......ok.......")
					ok = true
					
					break
				}
			}
			
			if(!ok){
				fmt.Printf("\rcrt.sh > %s", err.Error())
				time.Sleep(1 * time.Second)
				respCrt <- domain.SubDomains{}
				
			}
			
		}()
	
		crt = <- respCrt
		
	}

	subDomainsStrings = append(subDomainsStrings, crt.SubDomains...)
	

	//--------------ROBTEX.COM------------------------
	rbt := &robtex.Robtext{NameService: "robtext.com", Domain: dominio, Url:urlRobtex}
	//-----------------HACKERTARGET.COM----------------------------------------
	hackTarget := &hacktarget.Htarget{NameService:"hackertarget", Domain: dominio, Url: urlHtarget}
	//-------------------URLSCAN.IO-----------------------------------------
	scanIo := &urlscan.UrlScan{NameService: "urlscan.io", Domain: dominio, Url: urlUrlScanio}
	//----------------RAPIDDNS.COM------------------------------------------
	rapid := &rapiddns.RapidDns{NameService: "rapiddns", Domain: dominio, Url: urlRapidDns}
	
	
	services := []sources.Scan{
		rbt,
		hackTarget,
		scanIo,
		rapid,
	}
	for _ , service := range services{
		serv, err := ScanSubdomain(service)
		if(err != nil){
			fmt.Printf("\n%s", err.Error())
			serv = domain.SubDomains{}
		}
		subDomainsStrings = append(subDomainsStrings, serv.SubDomains...)
	}

	//----limpiar duplicados
	listClean := funcs.DeleteRepeat(subDomainsStrings)
	

	//-----INICIO LISTA  DOMINIOS (SUBDOMINIOS)
	var subdomains []domain.Domain
	
	start := time.Now()
	fmt.Printf("\r%s", "Starting:.....................")
	


	//------ INICIO ----
	subdomains = Start(listClean, cdnList)
	
	//----guardar si es requerido---------
	if(savefile){
		var data string
		for _, v := range subdomains{
			data += fmt.Sprintf("%s, %s,  %s\n", v.Name, v.Ip, v.Cdns)
		}
		funcs.Save("subdomains-"+dominio+".txt", &data)
	}
	
	//-------RESULTADOS
	fmt.Printf("\r%s","---------------------------Results-------------------------------\n")
	fmt.Println("Domain: ", dominio)
	fmt.Println("-----------------------------------------------------------------")
	for n, v := range subdomains{
		time.Sleep(100 * time.Millisecond)
		fmt.Println(n+1,style.YELLOW, v.Name, style.END ,style.GREEN, v.Ip, style.END) 
		fmt.Println(style.Randcolor() + "   Cdn: >" + style.END, style.GREEN, v.Cdns, style.END)
		fmt.Println("-----------------------------------------------------------------")
	}
	end := time.Since(start)
	fmt.Println("Execution time:", end)
	

}


func Start(lista []string, cdnlist *[]IPs.Cdn)[]domain.Domain{
	subdomains := []domain.Domain{}
	dmain := make(chan *domain.Domain, 10)
	var wg sync.WaitGroup
	limitElements := 100
	numThreads := 7


	// ---- SI HAY MAS DE 100 SUBDOMINIOS CREAMOS 7 GOROUTINES PARALELAS
	if(len(lista) > limitElements){
		fmt.Printf("\r%s", "Accelerating....\n")
		time.Sleep(1 * time.Second)
		chunksSubdomains := funcs.SplitArray(lista, numThreads)

		for _,list := range chunksSubdomains{
			wg.Add(1)
			go func(lista []string){
				defer wg.Done()
				for _, x := range lista{
					
					ip, err := funcs.CheckIp(x, true, Resolver)
					if(err != nil){
						fmt.Printf("\r%s", style.RED + err.Error()[:31] + style.END)
						dmain <- nil
						continue
					}
					domaiin := &domain.Domain{Name: x, Ip: *ip }
					domaiin.FindCdn(cdnlist)
					//subdomains = append(subdomains, domaiin)
					dmain <- domaiin
					}		
			}(list)

		}

		go func(){
			wg.Wait()
			close(dmain)
		}()

		for info := range dmain{
			if(info != nil){
				subdomains = append(subdomains, *info)
			}
		}
			
		return subdomains


	}

	//-----SI SON MENOS DE 100 SUBDOMIIOS SE HACE ESCANEO NORMAL---

	for _, x := range lista{
		ip, _ := funcs.CheckIp(x, true, Resolver)
		domaiin := domain.Domain{Name: x, Ip: *ip }
		domaiin.FindCdn(cdnlist)
		subdomains = append(subdomains, domaiin)

	}

	//-----RETORNAMOS RESULTADOS
	return subdomains

}



// ------PARA PROBAR LA INTERFACE :| 
func ScanSubdomain(s sources.Scan)(domain.SubDomains, error){

	resp, err := s.CheckSubdomain()
	if(err != nil){
		
		//fmt.Printf("\r%s",s.ServiceName())
		return resp, err
	}
	return resp, nil

}




func Help()string{
	return `
Help:
Use: 
    ./gocdn [options] <arguments>
Android:
    ./gocdn-android [options] <arguments>

Options:
  --cdn <IP>            "Scann all CDN for this ip"
  --subdomain <DOMAIN>  "Scann all CDN and Subdomains for this domain"	
  --help	        "Help"
  [--save] [-s]         "Save results in a file (./subdomains-your-domain.txt)" 

Example:
  ./gocdn --cdn  123.123.123.123
  ./gocdn --subdomain  mydomain.com
  ./gocdn --subdomain mydomain.com --save 

`
}

