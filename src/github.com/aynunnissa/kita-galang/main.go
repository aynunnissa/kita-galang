package main

import (
	"github.com/aynunnissa/kita-galang/controllers"
	"github.com/aynunnissa/kita-galang/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.GET("/donations", controllers.DonationsIndex)
	r.POST("/create-donation", controllers.DonationsCreate)
	// TODO: one day will change the 'id' to 'slug'
	r.GET("/donations/:id", controllers.DonationsShow)
	r.PUT("/donations/:id/donate", controllers.DonationsDonate)
	r.DELETE("/donations/:id", controllers.DonationsDelete)

	r.Run()
}
