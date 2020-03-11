package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

/* 		ITIC 10-1
Monica Alejandra Hernandez Cordero
*/

type Alumno struct {
	Nombre       string
	Calificacion string
}

func main() {
	menu :=
		`¿Qué quiere realizar?
				1. Capturar
				2. Mostrar
				3. Salir
		`
	var seleccion int

	for seleccion != 3 {
		fmt.Print(menu)
		fmt.Scanln(&seleccion)
		scanner := bufio.NewScanner(os.Stdin)
		switch seleccion {
		case 1:
			var Alumnos []Alumno
			for i := 0; i < 5; i++ {
				var a Alumno
				fmt.Println("Nombre del alumno:")
				if scanner.Scan() {
					a.Nombre = scanner.Text()
				}
				fmt.Println("Calificaciones del alumno:")
				if scanner.Scan() {
					a.Calificacion = scanner.Text()
				}

				Alumnos = append(Alumnos, a)
			}

			for i := 0; i < len(Alumnos); i++ {
				InsercionDatos(Alumnos[i].Nombre, Alumnos[i].Calificacion)
			}

		case 2:

			data, err := ioutil.ReadFile("calificaciones.txt")
			if err != nil {
				fmt.Println(err)
			}

			fmt.Println("La nota contiene \n " + string(data))
		}
	}
}

func InsercionDatos(Nombre string, Calificacion string) {
	f, err := os.OpenFile("calificaciones.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	if _, err := f.WriteString(Nombre + " " + Calificacion + "\n"); err != nil {
		fmt.Println(err)
	}

}
