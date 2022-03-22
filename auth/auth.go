package auth

import (
	"Centralized_transaction/database"
	"Centralized_transaction/models"
	"Centralized_transaction/security"
	"Centralized_transaction/utils/channels"

	"github.com/jinzhu/gorm"
)

func SignIn(email, password string) (models.User, error) {
	user := models.User{}
	var err error
	var db *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {
		db, err = database.Connect()
		if err != nil {
			ch <- false
			return
		}
		defer db.Close()

		err = db.Debug().Model(models.User{}).Where("email = ?", email).Take(&user).Error
		if err != nil {
			ch <- false
			return
		}
		err = security.VerifyPassword(user.Password, password)
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		user.Token,err = CreateToken(user.ID)
		return user,err
	}
	return user, err
}
