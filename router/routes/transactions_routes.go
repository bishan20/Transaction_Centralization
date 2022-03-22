package routes

import (
	"Centralized_transaction/controllers"
	"net/http"
)

var transactionsRoutes = []Route{
	{
		Url:          "/transaction/list",
		Method:       http.MethodPost,
		Handler:      controllers.GetTransactions,
		AuthRequired: true,
	},
	{
		Url:          "/transaction",
		Method:       http.MethodPost,
		Handler:      controllers.CreateTransaction,
		AuthRequired: false,
	},
	//{
	// 	Url:          "/posts/{id}",
	// 	Method:       http.MethodGet,
	// 	Handler:      controllers.GetPost,
	// 	AuthRequired: false,
	// },
	// {
	// 	Url:          "/posts/{id}",
	// 	Method:       http.MethodPut,
	// 	Handler:      controllers.UpdatePost,
	// 	AuthRequired: true,
	// },
	// {
	// 	Url:          "/posts/{id}",
	// 	Method:       http.MethodDelete,
	// 	Handler:      controllers.DeletePost,
	// 	AuthRequired: true,
	// },
}
