package main

import (
	"fmt"

	"github.com/Sree9446086944/jsonCrudApi/controllers"
	"github.com/Sree9446086944/jsonCrudApi/initializers"
	"github.com/gin-gonic/gin"
)

// CompileDaemon pkg - go get github.com/githubnemo/CompileDaemon -> builds everytime there is change
// go install github.com/githubnemo/CompileDaemon

//godotenv pkg - for loading env variables
// go get github.com/joho/godotenv

//Gin - http web framework in go - https://gin-gonic.com/docs/quickstart/
// go get -u github.com/gin-gonic/gin

//GORM - Go ORM library - https://gorm.io/docs/
// go get -u gorm.io/gorm
// go get -u gorm.io/driver/postgres -> for postgres driver

// init() will run before main
func init() {
	//  check from gin doc
	/*
		//move this logic to another file and import here -> copy to loadEnvVariables.go in initializers directory
		//to load .env file
		err := godotenv.Load() // now port will be 3000, since given in .env file , loaded
		if err != nil {
			log.Fatal("Error loading .env file.")
		}
	*/

	//above logic from another file
	initializers.LoadEnvVariables() // restart if didnt get imported
	//connect to db using gorm, in another file and import here
	initializers.ConnectToDB()
}

func main() {
	fmt.Println("Hello!!!")
	//to run CompileDaemon ->  CompileDaemon -command="./jsonCrudApi"  // executable name

	//create gin router
	r := gin.Default()

	//create route
	/*r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"}) // create json response, H is map
	})*/
	//OR
	r.GET("/", controllers.Test) //handler method created in controllers pkg

	r.POST("/posts", controllers.PostsCreate)
	r.GET("/posts", controllers.PostsIndex)
	r.GET("/posts/:id", controllers.PostsShow)

	r.PUT("/posts/:id", controllers.PostsUpdate)

	r.DELETE("/posts/:id", controllers.PostsDelete)

	r.Run() // starts server and in port 8080 default, PORT changed in .env file to 3000

}
