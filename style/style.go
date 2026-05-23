package style

import "math/rand"




var (

	BLUE string = "\033[1;34m"
	YELLOW string = "\033[1;33m"
	GREEN string = "\033[1;92m"
	RED string = "\033[1;31m"
	END string = "\033[0m"
	CYAN string = "\033[1;36m"
	SUB string = "\033[4m"
	WHITE string = "\033[1;37m"


)

func Randcolor()string{
	colors := []string{
				"\033[1;30m", // Negrita Negro 
				"\033[1;31m", // Negrita Rojo
				"\033[1;32m", // Negrita Verde
				"\033[1;33m", // Negrita Amarillo
				"\033[1;34m", // Negrita Azul
				"\033[1;35m", // Negrita Magenta
				"\033[1;36m", // Negrita Cyan
				"\033[1;37m", // Negrita Blanco
				"\033[1;90m", // Negrita Gris
				"\033[1;91m", // Negrita Rojo brillante
				"\033[1;92m", // Negrita Verde brillante
				"\033[1;93m", // Negrita Amarillo brillante
				"\033[1;94m", // Negrita Azul brillante
				"\033[1;95m", // Negrita Magenta brillante
				"\033[1;96m", // Negrita Cyan brillante
				"\033[1;97m", // Negrita Blanco brillante
	}
	
	return colors[rand.Intn(len(colors))]
}