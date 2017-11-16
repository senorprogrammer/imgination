package modules

import ()

type SearchResult struct {
	Results []string
}

func (searchResult *SearchResult) Append(result string) {
	searchResult.Results = append(searchResult.Results, result)
}

func (searchResult *SearchResult) Len() int {
	return len(searchResult.Results)
}
