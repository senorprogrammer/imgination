package modules

import (
	"sort"
)

type SearchResult struct {
	Results []string
}

func (searchResult *SearchResult) Sorted() []string {
	sort.Strings(searchResult.Results)
	return searchResult.Results
}

func (searchResult *SearchResult) Append(result string) {
	searchResult.Results = append(searchResult.Results, result)
}

func (searchResult *SearchResult) Len() int {
	return len(searchResult.Results)
}
