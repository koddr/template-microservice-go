package main

func main() {
	// Initialize application.
	app, err := initializeApplication()
	if err != nil {
		app.Logger.Error("failed to initialize application", "details", err.Error())
		panic(err)
	}

	// Make sure to close the DB connection when the application exits.
	defer app.Database.Connection.Close()

	// Run application.
	if err := app.Run(); err != nil {
		app.Logger.Error("failed to run application", "details", err.Error())
		panic(err)
	}
}
