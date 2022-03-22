package auto

import (
	"Centralized_transaction/database"
	"Centralized_transaction/models"
	"log"
)

func Load() {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// err = db.Debug().DropTableIfExists(&models.Transaction{}, &models.User{}).Error
	// if err != nil {
	// 	log.Fatal(err)
	// }
	err = db.Debug().AutoMigrate(&models.User{}, &models.Transaction{}).Error
	if err != nil {
		log.Fatal(err)
	}

}
