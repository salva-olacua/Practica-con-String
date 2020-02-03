package myStringUtils

import "strings"

func StringSinRepetidos(str string) string {
	sinRepetidos:= ""

	for _,r := range str {		
		s:=string(r)
		if(!strings.ContainsAny(sinRepetidos,s)) {
			sinRepetidos+=s
		}
	}
	return sinRepetidos
}


// La libreria "strings" es estricta con el orden, por ende declaro esta nueva funcion.
func containsAllIrrestrict(aComparar, str string) bool {
	for _,r := range str{
		if(!strings.Contains(aComparar,string(r))){
			return false;
		}
	}
	return true
}


func SubstringMasCortoConUnPatron(str, cadenaPatron string) string {
	tamStringMaximo := len(str)
	tamStringMinimo := len(cadenaPatron)
	i := 0

	for tamStringMinimo<=tamStringMaximo {
		//busqueda del substring mas corto
		if(containsAllIrrestrict(str[i:tamStringMinimo+i],cadenaPatron)){
			return str[i:tamStringMinimo+i]
		}
		//condicion de reinicio del algoritmo
		if tamStringMinimo+i==tamStringMaximo{
			tamStringMinimo++
			i=-1
		}
		i++
	}
	return ""
}
