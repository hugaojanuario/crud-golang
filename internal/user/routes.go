package user

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine, handler *Handler){
	users := r.Group("/users")

	users.POST("/", handler.Create)
	users.GET("/", handler.GetAll)
	users.GET("/:id", handler.GetById)
	users.PUT("/:id", handler.Update)
	users.DELETE("/:id", handler.Delete)

}