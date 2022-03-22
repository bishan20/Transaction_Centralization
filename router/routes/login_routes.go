package routes

import (
	"Centralized_transaction/controllers"
	"net/http"
)

var loginRoutes = []Route{
	{
		Url:          "/login",
		Method:       http.MethodPost,
		Handler:      controllers.Login,
		AuthRequired: false,
	},
}
