package modules

import (
	"fmt"
)

func Render(searchResult *SearchResult) {
	fmt.Println("\n")
	fmt.Printf("Found %d results\n\n", searchResult.Len())

	for _, result := range searchResult.Results {
		fmt.Println(result)
	}
}
