package helper

import (
	"strconv"
)

func GenerateZeros(valor int, length int) string {
	num := valor

	// Convetir el número a una cadena
	str := strconv.Itoa(num)

	// Compara la longitud actual con la longitud deseada
	diff := length - len(str)

	// Agregar ceros delante del número si es necesario
	if diff > 0 {
		zeros := ""
		for i := 0; i < diff; i++ {
			zeros += "0"
		}
		str = zeros + str
	}

	// Retornamos la variable convertida en el formato ######
	return str
}
