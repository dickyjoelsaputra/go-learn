package database

import (
	"fmt"
	"go_learn/models"
	"go_learn/pkg/mysql"
)

// Automatic Migration if Running App
func RunMigration() {
  err := mysql.DB.AutoMigrate(&models.User{})

  if err != nil {
    fmt.Println(err)
    panic("Migration Failed")
  }

  fmt.Println("Migration Success")
}