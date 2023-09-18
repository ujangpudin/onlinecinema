package database

import (
	"backend_project/models"
	"backend_project/pkg/mysql"
	"fmt"
)

// Automatic Migration if Running App
func RunMigration() {
	mysql.DB.AutoMigrate(&models.User{}, &models.Category{}, &models.Film{}, &models.Transaction{})
	// mysql.DB.AutoMigrate(&models.User{})
	// mysql.DB.AutoMigrate(&models.Profile{})

	//   if err != nil {
	//     fmt.Println(err)
	//     panic("Migration Failed")
	//   }

	fmt.Println("Migration Success")
}
