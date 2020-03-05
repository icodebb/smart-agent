/**
 * Web functions.
 *
 */

package web

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// Handler for properties.
func Handler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	name := query.Get("name")
	if name == "" {
		name = "My Friend"
	}

	remote := r.RemoteAddr

	log.Printf("Received %s request for %s from %s\n", r.Method, name, remote)
	w.Write([]byte(fmt.Sprintf("Hello %s, a %s from %s\n", name, r.Method, remote)))
	w.Write([]byte(fmt.Sprintf("%s\n", time.Now())))
	w.Write([]byte("\n"))
}

// HealthHandler for health sub page.
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

// ReadinessHandler for readiness.
func ReadinessHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
