package config

import "net"
import "context"



var Android bool = false

func ErrorMessage()string{
	var ErrorMsg string = "Invalid option, please type ./gocdn --help"
	if(Android){
		ErrorMsg = "Invalid option, please type ./gocdn-android --help"
	}
	return ErrorMsg
}

//Configuracion de Resolver y Dialer para android
func ConfigDialerAndResolver() *net.Dialer{
	resolver := ConfigResolver()
	dialer := &net.Dialer{
		Resolver: resolver,
	}
	return dialer
}


//Solo resolver para android
func ConfigResolver()*net.Resolver{

	resolver := &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			d := net.Dialer{}
			return d.DialContext(ctx, "udp", "8.8.8.8:53")
		},
	}
	return resolver

}
