package controllers

import (
	"Centralized_transaction/auth"
	"Centralized_transaction/models"
	"Centralized_transaction/responses"
	"encoding/json"
	"io/ioutil"
	"net/http"
)
type TokenRequest struct{
	Token string `json:"token"`
}
func Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user.Prepare()
	err = user.Validate("login")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	user, err = auth.SignIn(user.Email, user.Password)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user.Password = ""
	responses.JSON(w, http.StatusOK, user)
}
