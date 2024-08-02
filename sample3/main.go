package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

var query = "SELECT id, appVersion, launchTime, asVersion, deviceModel, createdAt FROM launch_time_tracking %s"

func handler2(w http.ResponseWriter, r *http.Request) {
	// Open an in-memory SQLite database
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// format := r.URL.Query().Get("format")

	filterParams := map[string]string{}

	if appVersionFilterParam := r.URL.Query().Get("app_version"); appVersionFilterParam != "" {
		filterParams["app_version"] = appVersionFilterParam
	}

	if len(filterParams) > 0 {
		conditions := []string{}

		if appVersionFilter, ok := filterParams["app_version"]; ok {
			conditions = append(conditions, fmt.Sprintf("app_version = '%s'", appVersionFilter))
		}

		//execute the query
		db.Query(fmt.Sprintf(query, conditions[0]))
	}

}

func main() {
	fmt.Println("Hello, World!")
	// use handler2 function
	http.HandleFunc("/handler2", handler2)

}
