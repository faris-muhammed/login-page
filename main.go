package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"main.go/database"
	"main.go/routes"
)

func init() {
	database.DBconnect()
}

func main() {

	r := gin.Default()

	r.LoadHTMLGlob("template/*")

	stor := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", stor))

	//user
	r.GET("/", routes.Handle)
	r.POST("/login", routes.Login)
	r.GET("/signup", routes.Signup)
	r.POST("/signup", routes.SignupPost)
	r.GET("/logout", routes.Logout)
	r.GET("/home", routes.HomeHandler)

	//admin
	r.GET("/admin", routes.AdminLogin)
	r.POST("/admin", routes.AdminPost)
	r.GET("/adminhome", routes.AdminHome)
	r.GET("/delete/:ID", routes.DeleteUsers)
	r.GET("/block/:ID", routes.BlockUsers)
	r.GET("/edit/:ID", routes.Edit)
	r.POST("/edit/:ID", routes.EditUser)
	r.GET("/adlogout", routes.Adlogout)

	r.Run(":1111")
}
