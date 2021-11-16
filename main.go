package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/foolin/goview/supports/echoview-v4"
	"github.com/jeisson-u/web-project-rol.git/apiController"
	"github.com/jeisson-u/web-project-rol.git/database"
	"github.com/jeisson-u/web-project-rol.git/webController"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	//input database
	db_server := "localhost"
	db_user := "root"
	db_pasw := "12345"
	db_port := "1247"
	db_database := "ucompensar"

	fmt.Println("=========DATOS DE CONEXION DATABASE======")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("SERVIDOR BASE DE DATOS:")
	db_server, _ = reader.ReadString('\n')
	fmt.Println("PUERTO BASE DE DATOS:")
	db_port, _ = reader.ReadString('\n')
	fmt.Println("USUARIO BASE DE DATOS :")
	db_user, _ = reader.ReadString('\n')
	fmt.Println("CONTRASEÑA BASE DE DATOS :")
	db_pasw, _ = reader.ReadString('\n')
	fmt.Println("NOMBRE BASE DE DATOS :")
	db_database, _ = reader.ReadString('\n')



	fmt.Println( fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", db_user, db_pasw, db_server, db_port, db_database))

	os.Setenv("DATABASE_CONNECTION", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", db_user, db_pasw, db_server, db_port, db_database))

	// initialize database
	fmt.Println("connect to database...")

	initDb := database.InitializeDatabase{}

	openDb, err := initDb.Start()
	defer initDb.Close()

	if !openDb {
		fmt.Println(err.Error())
		return
	}

	/*	err = initDb.Create()

		if err != nil {
			fmt.Println(err.Error())
			return
		}*/

	fmt.Println(fmt.Printf("database created at: %s:%s (%s)", db_server, db_port, db_database))

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//Set Renderer
	e.Renderer = echoview.Default()

	securityApiController := apiController.SecurityController{}
	userApiController := apiController.UserController{}
	rolApiController := apiController.RolController{}
	securityWebController := webController.SecurityController{}
	homeWebController := webController.HomeController{}
	userWebController := webController.UserController{}
	rolWebController := webController.RolController{}
	billingWebController := webController.BillingController{}

	// static files
	e.Static("/public", "public")

	apiSecurity := e.Group("/api/v1/security")
	apiSecurity.POST("/login", securityApiController.Login)
	apiSecurity.POST("/refresh-token", securityApiController.Login)

	apiUser := e.Group("/api/v1/user")
	apiUser.POST("", userApiController.CreateUser)
	apiUser.PUT("", userApiController.UpdateUser)
	apiUser.DELETE("/:userId", userApiController.RemoveUser)
	apiUser.GET("", userApiController.GetAll)
	apiUser.GET("/:userId/apps", userApiController.GetUserApps)

	apiRol := e.Group("/api/v1/rol")
	apiRol.GET("", rolApiController.GetAll)

	web := e.Group("/")
	web.GET("", securityWebController.LoginPage)
	web.GET("home", homeWebController.HomePage)
	web.GET("app/1", userWebController.UsersPage)
	web.GET("app/2", rolWebController.RolsPage)
	web.GET("app/3", billingWebController.BillingPage)

	e.Logger.Fatal(e.Start(":1323"))
}
