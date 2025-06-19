package application

import (
	"crypto/sha256"
	"crypto/subtle"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// BasicAuthMiddleware checks, if the user is authenticated§.
func (a *Application) BasicAuthMiddleware(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		// Get the Basic Authentication credentials.
		username, password, ok := r.BasicAuth()

		if ok {
			// Calculate SHA-256 hashes for the provided and expected
			// usernames and passwords.
			usernameHash := sha256.Sum256([]byte(username))
			passwordHash := sha256.Sum256([]byte(password))
			expectedUsernameHash := sha256.Sum256([]byte(a.Config.Server.BasicAuthUsername))
			expectedPasswordHash := sha256.Sum256([]byte(a.Config.Server.BasicAuthPassword))

			// Use the subtle.ConstantTimeCompare() function to check if
			// the provided username and password hashes equal the
			// expected username and password hashes. ConstantTimeCompare
			// will return 1 if the values are equal, or 0 otherwise.
			// Importantly, we should to do the work to evaluate both the
			// username and password before checking the return values to
			// avoid leaking information.
			usernameMatch := (subtle.ConstantTimeCompare(usernameHash[:], expectedUsernameHash[:]) == 1)
			passwordMatch := (subtle.ConstantTimeCompare(passwordHash[:], expectedPasswordHash[:]) == 1)

			// If the username and password are correct, then call
			// the next handler in the chain. Make sure to return
			// afterwards, so that none of the code below is run.
			if usernameMatch && passwordMatch {
				// Delegate request to the given handle.
				h(w, r, ps)
				return
			}
		}

		// Reject requests without correct Basic Authentication credentials.
		w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
		a.Logger.Error(http.StatusText(http.StatusUnauthorized), "details", http.StatusUnauthorized)
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
	}
}
