package app

import (
	"fmt"
)

var (
	router *gin.Engine
)

func init() {
	router = gin.Default()
}

//StartApp is the main function of the project
func StartApp() {
	fmt.Println("Starting API on port 8080")

	addRoutes()

	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}

