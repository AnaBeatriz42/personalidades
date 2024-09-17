package main

import (
	"fmt"

	"github.com/AnaBeatriz42/api-personalidades/database"
	"github.com/AnaBeatriz42/api-personalidades/models"
	"github.com/AnaBeatriz42/api-personalidades/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "nome1", Historia: "historia1"},
		{Id: 2, Nome: "nome2", Historia: "historia2"},
		{Id: 3, Nome: "nome3", Historia: "historia3"},
	}

	database.ConectaBanco()

	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleRequest()
}
