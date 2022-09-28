package main

import (
	"github.com/lucas-braga-souza/challenge_api-go/database"
	"github.com/lucas-braga-souza/challenge_api-go/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequets()
}
