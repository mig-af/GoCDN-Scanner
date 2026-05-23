package config

import "net"
import "context"



var Android bool = false


//Configuracion de Resolver y Dialer para android
func ConfigDialerAndResolver() *net.Dialer{
	resolver := &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			d := &net.Dialer{}
			return d.DialContext(ctx, "udp", "8.8.8.8:53")
		},
	}
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
