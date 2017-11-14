package wptdashboard

import (
	"html/template"
	"net/http"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func init() {
	http.HandleFunc("/tasks/populate-dev-data", populateDevData)
	http.HandleFunc("/test-runs", testRunHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/api/runs", apiTestRunsHandler)
	http.HandleFunc("/api/run", apiTestRunHandler)
	http.HandleFunc("/results", resultsRedirectHandler)
	http.HandleFunc("/", testHandler)
}
