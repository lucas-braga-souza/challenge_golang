package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lucas-braga-souza/challenge_api-go/controllers"
)

func HandleRequets() {
	r := gin.Default()
	r.GET("/usuario", controllers.ExibeTodosUsuario)
	r.GET("/usuario/:username", controllers.BuscaUsuarioPorUsername)
	r.POST("/usuario", controllers.CriaNovoUsuario)
	r.PUT("/usuario/:username", controllers.EditaUsuario)
	r.DELETE("/usuario/:username", controllers.DeletaUsuario)
	r.Run()
}
