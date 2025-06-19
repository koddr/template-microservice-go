package application

import (
	"context"
	"io"
	"net/http"
	"time"

	jsoniter "github.com/json-iterator/go"
	"github.com/julienschmidt/httprouter"
)

// GetHealthCheckAPIHandler checks the health of the server (GET).
func (a *Application) GetHealthCheckAPIHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	_, err := w.Write([]byte(`{"status": "ok"}`))
	if err != nil {
		a.Logger.Error("error writing JSON response", "details", err.Error())
		http.Error(w, "error writing JSON response", http.StatusInternalServerError)
		return
	}
}

// GetAllTransactionsAPIHandler gets all transactions from the database (GET).
func (a *Application) GetAllTransactionsAPIHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// SQL query to get all transactions.
	query, err := a.Attachments.Queries.ReadFile("queries/get_all_transactions.sql")
	if err != nil {
		a.Logger.Error("error reading query file from embedded filesystem", "details", err.Error())
		http.Error(w, "error reading query file from embedded filesystem", http.StatusInternalServerError)
		return
	}

	// Prepare a variable to hold the JSON response.
	var jsonResponse []byte

	// Execute the SQL statement and get JSON response directly.
	// Using QueryRow to fetch a single row and scan the result into jsonResponse.
	if err = a.Database.Connection.QueryRow(
		r.Context(),
		string(query),
	).Scan(&jsonResponse); err != nil {
		a.Logger.Error("error selecting data from database", "details", err.Error())
		http.Error(w, "error selecting data from database", http.StatusInternalServerError)
		return
	}

	// Set the content type.
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON response.
	_, err = w.Write(jsonResponse)
	if err != nil {
		a.Logger.Error("error writing JSON response", "details", err.Error())
		http.Error(w, "error writing JSON response", http.StatusInternalServerError)
		return
	}
}

// GetTransactionsByFilterAPIHandler gets transactions from the database with filter (GET).
func (a *Application) GetTransactionsByFilterAPIHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Check if there is data in the query params.
	if r.URL.Query() == nil {
		a.Logger.Error("empty 'created_at_start' or 'created_at_to' query params")
		http.Error(w, "empty 'created_at_start' or 'created_at_to' query params", http.StatusBadRequest)
		return
	}

	// Read the query params.
	createdAtStart := r.URL.Query().Get("created_at_start")
	createdAtEnd := r.URL.Query().Get("created_at_end")

	// Check if the query params are empty.
	if createdAtStart == "" || createdAtEnd == "" {
		a.Logger.Error("empty 'created_at_start' or 'created_at_end' query params")
		http.Error(w, "empty 'created_at_start' or 'created_at_end' query params", http.StatusBadRequest)
		return
	}

	// Parse the 'created_at_start' query param.
	// This is the start time for the filter, so it should be inclusive.
	createdAtStartTime, err := time.Parse(time.DateOnly, createdAtStart)
	if err != nil {
		a.Logger.Error("error parsing 'created_at_start' query param", "details", err.Error())
		http.Error(w, "error parsing 'created_at_start' query param", http.StatusBadRequest)
		return
	}

	// Parse the 'created_at_end' query param.
	// This is the end time for the filter, so it should be inclusive.
	createdAtEndTime, err := time.Parse(time.DateOnly, createdAtEnd)
	if err != nil {
		a.Logger.Error("error parsing 'created_at_end' query param", "details", err.Error())
		http.Error(w, "error parsing 'created_at_end' query param", http.StatusBadRequest)
		return
	}

	// SQL query to get transactions with 'created_at' filters.
	query, err := a.Attachments.Queries.ReadFile("queries/get_transactions_with_created_at_filters.sql")
	if err != nil {
		a.Logger.Error("error reading query file from embedded filesystem", "details", err.Error())
		http.Error(w, "error reading query file from embedded filesystem", http.StatusInternalServerError)
		return
	}
	// Prepare a variable to hold the JSON response.
	var jsonResponse []byte

	// Execute the SQL statement and get JSON response directly.
	// Using QueryRow to fetch a single row and scan the result into jsonResponse.
	if err = a.Database.Connection.QueryRow(
		r.Context(),
		string(query),
		createdAtStartTime,
		createdAtEndTime,
	).Scan(&jsonResponse); err != nil {
		a.Logger.Error("error selecting data from database", "details", err.Error())
		http.Error(w, "error selecting data from database", http.StatusInternalServerError)
		return
	}

	// Set the content type.
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON response.
	_, err = w.Write(jsonResponse)
	if err != nil {
		a.Logger.Error("error writing JSON response", "details", err.Error())
		http.Error(w, "error writing JSON response", http.StatusInternalServerError)
		return
	}
}

// AddTransactionAPIHandler adds a new transaction to the database (POST).
func (a *Application) AddTransactionAPIHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Read request body.
	body, err := io.ReadAll(r.Body)
	if err != nil {
		a.Logger.Error("error reading request body", "details", err.Error())
		http.Error(w, "error reading request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Create a new Transaction instance.
	transaction := &Transaction{}

	// Unmarshal request body.
	if err := jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal(body, &transaction); err != nil {
		a.Logger.Error("error unmarshaling request body", "details", err.Error())
		http.Error(w, "error unmarshaling request body", http.StatusBadRequest)
		return
	}

	// Write the status code.
	w.WriteHeader(http.StatusCreated)

	// Run the database operation in a goroutine.
	go func() {
		// Create a context with a timeout for the database operation.
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		// Read the SQL insert statement and use pgx to execute it.
		query, err := a.Attachments.Queries.ReadFile("queries/insert_new_transaction.sql")
		if err != nil {
			a.Logger.Error("error reading query file from embedded filesystem", "details", err.Error())
			return
		}

		// Execute the SQL statement.
		_, err = a.Database.Connection.Exec(
			ctx,
			string(query),
			transaction.ProfileID,
			transaction.MessengerID,
			transaction.MessengerName,
			transaction.EventID,
			transaction.EventType,
			transaction.UTMSource,
			transaction.UTMMedium,
			transaction.UTMCampaign,
			transaction.UTMContent,
			transaction.UTMTerm,
		)
		if err != nil {
			a.Logger.Error("error inserting data into database", "details", err.Error())
		}
	}()
}
