package routes

import (
	"fmt"
	"log"
	"net/http"

	"main.go/database"
	"main.go/jwt"
	"main.go/model"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var Error string
var Fetch model.UserModel
var Updatefeach model.UserModel

const RoleUser = "user"

// login page GET
func Handle(c *gin.Context) {
	c.Header("Cache-Control", "no-cache,no-store,must-revalidate")
	session := sessions.Default(c)
	check := session.Get(RoleUser)
	if check == nil {
		c.HTML(http.StatusSeeOther, "login.html", Error)
		Error = ""
	} else {
		c.Redirect(http.StatusSeeOther, "/home")
	}
}

// login page POST
func Login(c *gin.Context) {

	Fetch = model.UserModel{}
	database.DB.First(&Fetch, "email=?", c.Request.FormValue("username"))
	fmt.Println("==========username:", Fetch.Password, "==========")

	plainpassword := c.Request.FormValue("password")
	err := bcrypt.CompareHashAndPassword([]byte(Fetch.Password), []byte(plainpassword))

	if err != nil {
		Error = "Invalid Username or Password"
		c.Redirect(http.StatusSeeOther, "/")
	} else {
		if Fetch.Status == "Blocked" {
			Error = "Blocked User"
			c.Redirect(http.StatusSeeOther, "/")
		} else {
			jwt.JwtToken(c, Fetch.Email, RoleUser)
			c.Redirect(http.StatusSeeOther, "/home")
		}
	}
}

func HomeHandler(c *gin.Context) {
	c.Header("Cache-Control", "no-cache,no-store,must-revalidate")
	session := sessions.Default(c)
	check := session.Get(RoleUser)
	if check != nil {
		c.HTML(http.StatusSeeOther, "home.html", Fetch)
	} else {
		c.Redirect(http.StatusSeeOther, "/")
	}
}

// logout page
func Logout(c *gin.Context) {
	c.Header("cache-control", "no-cache,no-store,must-revalidate")
	session := sessions.Default(c)
	session.Delete(RoleUser)
	session.Save()
	//{clear previous user data}
	Fetch = model.UserModel{}

	c.Redirect(http.StatusSeeOther, "/")
}

// signup page
func Signup(c *gin.Context) {
	c.Header("Cache-control", "no-cache,no-store,must-revalidate")
	session := sessions.Default(c)
	check := session.Get(RoleUser)
	if check != nil {
		c.Redirect(http.StatusSeeOther, "/home")
	} else {
		c.HTML(http.StatusSeeOther, "signup.html", Error)
		Error = ""
	}
}

func SignupPost(c *gin.Context) {

	hashedPassword, error := bcrypt.GenerateFromPassword([]byte(c.Request.PostFormValue("password")), 10)

	if error != nil {
		log.Fatal("=======WE CANT PROVIDE HASHED PASSWORD=======")
	}

	err := database.DB.Create(&model.UserModel{
		Name:     c.Request.PostFormValue("username"),
		Password: string(hashedPassword),
		Email:    c.Request.PostFormValue("email"),
		Status:   "Active",
	})
	if err.Error != nil {
		Error = "Email already exists"
		c.Redirect(http.StatusSeeOther, "/signup")
	} else {
		Error = "User Successfully Created"
		c.Redirect(http.StatusSeeOther, "/")
	}
}
