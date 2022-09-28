package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lucas-braga-souza/challenge_api-go/controllers"
)

func HandleRequets() {
	r := gin.Default()
	r.GET("/usuario", controllers.ExibeTodosUsuario)
	r.GET("/usuario/username/:username", controllers.BuscaUsuarioPorUsername)
	r.GET("/usuario/:id", controllers.BuscaUsuarioPorID)
	r.POST("/usuario", controllers.CriaNovoUsuario)
	r.PUT("/usuario/:id", controllers.EditaUsuario)
	r.DELETE("/usuario/:id", controllers.DeletaUsuario)
	r.Run()
}
