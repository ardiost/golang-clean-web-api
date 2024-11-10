package routers

import (
	"github.com/ardiost/golang-clean-web-api/api/handlers"
	"github.com/gin-gonic/gin"
)

func TesterRouter(r *gin.RouterGroup) {
	x := handlers.NewTestHandler()
	r.GET("/", x.TestHa)
	r.GET("/users", x.GetUsers)
	r.GET("/user/:id", x.GetUserById)
	r.GET("/user/get-user-by-username/:username", x.GetUserByUserName)
	r.GET("/user/:id/accounts", x.GetAccounts)
	r.POST("/add-user", x.AddUser)
	r.POST("/bind/binder1", x.Binder1)
	r.POST("/bind/binder2", x.Binder2)
	r.POST("/query/query1", x.Query1)
	r.POST("/query/query2", x.Query2)
	r.POST("/body", x.BodyBind)

}
