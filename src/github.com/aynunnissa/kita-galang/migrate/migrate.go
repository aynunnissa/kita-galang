package main

import (
	"github.com/aynunnissa/kita-galang/initializers"
	"github.com/aynunnissa/kita-galang/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Donation{})
}
