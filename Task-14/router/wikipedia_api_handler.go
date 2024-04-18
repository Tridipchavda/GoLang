package router

import (
	"Task-14/model"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func GetWikipediaSearch(query string) (*model.SearchSuggestion, error) {
	res, err := http.Get("https://en.wikipedia.org/w/api.php?action=opensearch&search=" + query)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer res.Body.Close()

	var wikiSearchResults Search
	json.NewDecoder(res.Body).Decode(&wikiSearchResults)

	var suggestions *model.SearchSuggestion = &model.SearchSuggestion{}

	if len(wikiSearchResults) == 0 {
		suggestions = &model.SearchSuggestion{
			Name:  []interface{}{"Not Found"},
			Links: []interface{}{" "},
		}
	} else {
		suggestions = &model.SearchSuggestion{
			Name:  wikiSearchResults[1].([]interface{}),
			Links: wikiSearchResults[3].([]interface{}),
		}
	}

	return suggestions, nil
}

func GetWikipediaContent(title string) ([]byte, error) {
	rawContent, err := http.Get("https://en.wikipedia.org/wiki/" + title)
	if err != nil {
		return nil, err
	}
	defer rawContent.Body.Close()
	content, err := ioutil.ReadAll(rawContent.Body)
	if err != nil {
		return nil, err
	}
	return content, nil
}
