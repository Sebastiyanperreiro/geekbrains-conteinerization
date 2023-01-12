package main

import (
	"github.com/Sebastiyanperreiro/geekbrains-conteinerization/HW/8.ci-cd/app/app"
	"github.com/Sebastiyanperreiro/geekbrains-conteinerization/HW/8.ci-cd/app/config"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(":8000")
}
