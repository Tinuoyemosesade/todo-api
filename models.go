package main

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Title       string `json:"Mo"`
	Description string `json:"just a simple todo"`
	Completed   bool   `json:"completed"`
}
