package application

import (
	"github.com/julienschmidt/httprouter"
)

// router returns a new mux instance with all the routes.
func (a *Application) router() *httprouter.Router {
	// Create a new mux.
	router := httprouter.New()

	// Public API handler for health check.
	router.GET("/", a.GetHealthCheckAPIHandler)

	// Private API handler for getting all transactions from DB.
	router.GET("/api/v1/transactions", a.BasicAuthMiddleware(a.GetAllTransactionsAPIHandler))

	// Private API handler for getting transactions from DB with filter.
	router.GET("/api/v1/transactions/filter", a.BasicAuthMiddleware(a.GetTransactionsByFilterAPIHandler))

	// Private API handler for adding a new transaction to DB.
	router.POST("/api/v1/transaction", a.BasicAuthMiddleware(a.AddTransactionAPIHandler))

	return router
}
