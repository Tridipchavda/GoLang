package router

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// Handle Get Request for Wikipedia
func HandleSearch(w http.ResponseWriter, r *http.Request) {
	// set cross origin's
	setCors(w, r)
	// Get the user query input from the URL
	query := strings.ReplaceAll(mux.Vars(r)["query"], " ", "%20")
	// Search the user query and get the suggestion from wikipedia API
	suggestions, err := GetWikipediaSearch(query)
	if err != nil {
		http.Error(w, "Error while receiving suggestion from wikipedia", http.StatusNotFound)
		return
	}
	// Convert struct to string by Marshal and store it in suggestionResponse
	suggestionResponse, err := json.Marshal(suggestions)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Send Data to User related to Search
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(suggestionResponse))
}

func HandleWikiPediaContent(w http.ResponseWriter, r *http.Request) {
	// Set Cross origin's
	setCors(w, r)
	// Get Wikipedia content
	content,err := GetWikipediaContent(mux.Vars(r)["title"])
	if err != nil{
		http.Error(w,"Error while getting Content From Wikipedia",404)
	}
	// Show content on this page
	w.Write([]byte(content))
}
