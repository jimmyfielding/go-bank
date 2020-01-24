package registration

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gocql/gocql"

	"github.com/go-bank/internal/application/v1beta1"
	"github.com/go-bank/internal/cassandra"
	"github.com/go-bank/internal/response"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", register)
}

func register(w http.ResponseWriter, r *http.Request) {

	var app v1beta1.Application
	if err := json.NewDecoder(r.Body).Decode(&app); err != nil {
		_ = json.NewEncoder(w).Encode(response.ErrorResponse{Message: err.Error()})
		return
	}

	gocqlUuid := gocql.TimeUUID()

	if err := cassandra.Session.Query(`
		INSERT INTO users (id, firstName, surname) VALUES (?, ?, ?)`,
		gocqlUuid, app.FirstName, app.Surname, app.DateOfBirth); err != nil {
		_ = json.NewEncoder(w).Encode(response.ErrorResponse{Message: err.Error()})
		return
	}

	json.NewEncoder(w).Encode(response.OKResponse{Message: fmt.Sprintf("Created user with id: %s", gocqlUuid)})
	return
}
