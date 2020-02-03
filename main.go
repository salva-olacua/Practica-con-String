package main

import ("Practica-con-String/myStringUtils"
	    "fmt")
	   

func main(){
	//string de prueba
	str := "wqewqwqwrqtqwerty"
	strSinRepetidos := myStringUtils.StringSinRepetidos(str)
	substringMasCorto := myStringUtils.SubstringMasCortoConUnPatron(str,strSinRepetidos)
	fmt.Println("\n El string de ejemplo es: ",str)
	fmt.Println("\n El substring mas corto que tiene todas las letras sin repetir es: ",substringMasCorto)
}

