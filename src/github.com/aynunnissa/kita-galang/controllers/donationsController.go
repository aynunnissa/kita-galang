package controllers

import (
	"github.com/aynunnissa/kita-galang/initializers"
	"github.com/aynunnissa/kita-galang/models"
	"github.com/gin-gonic/gin"
)

func DonationsCreate(c *gin.Context) {
	// Get data off req body
	var body struct {
		Title       string
		Description string
		Target      int
		Collected   int
	}

	c.Bind(&body)

	// Create a donation
	donation := models.Donation{
		Title:       body.Title,
		Description: body.Description,
		Target:      body.Target,
		Collected:   body.Collected,
	}

	createdDonation := initializers.DB.Create(&donation)

	if createdDonation.Error != nil {
		c.Status(400)
		return
	}

	// Return created donation
	c.JSON(201, gin.H{
		"donation": donation,
	})
}

func DonationsIndex(c *gin.Context) {
	// Get all donations
	var donations []models.Donation
	initializers.DB.Find(&donations)

	c.JSON(200, gin.H{
		"donations": donations,
	})
}

func DonationsShow(c *gin.Context) {
	// TODO: one day change 'id' to 'slug'
	// Get donation id from url
	donationId := c.Param("id")

	// Get the donation
	var donation models.Donation
	initializers.DB.First(&donation, donationId)

	c.JSON(200, gin.H{
		"donation": donation,
	})
}

func DonationsDonate(c *gin.Context) {
	// TODO: one day change 'id' to 'slug'
	// Get donation id from url
	donationId := c.Param("id")

	// Get the nominal
	var body struct {
		Nominal int
	}

	c.Bind(&body)

	// Find the donation
	var donation models.Donation
	initializers.DB.First(&donation, donationId)

	// Update the collected nominal
	initializers.DB.Model(&donation).Update("collected", body.Nominal)

	// Send updated donation to the response
	c.JSON(200, gin.H{
		"donation": donation,
	})
}

func DonationsDelete(c *gin.Context) {
	// TODO: one day change 'id' to 'slug'
	// Get donation id from url
	donationId := c.Param("id")

	initializers.DB.Delete(&models.Donation{}, donationId)

	c.Status(200)
}
