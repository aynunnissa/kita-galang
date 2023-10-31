package models

import "gorm.io/gorm"

type Donation struct {
	gorm.Model
	Title       string
	Description string
	Target      int
	Collected   int
}
