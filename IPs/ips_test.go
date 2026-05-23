package IPs

import (
	"fmt"
	"testing"
)



func TestLoadIPsFromAPIBunny(t *testing.T){
	bunny := &BunnyCdn{Name: "Bunny"}

	bunny.IpRange = bunny.LoadIPsFromAPIBunny()
	fmt.Println(bunny.IpRange)
	// resp := GetIps("./IPs/")
	// fmt.Println(resp[5])
}
