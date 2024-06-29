package functions

import (
	"strings"

	"dobefu/search-page/structs"
)

func FilterStreets(streets []structs.Places, query structs.QueryParams) []structs.Places {
	if query.Query == "" {
		return streets
	}

	filteredStreets := make([]structs.Places, 0)

	for i := 0; i < len(streets); i++ {
		if !strings.Contains(strings.ToLower(streets[i].Place), strings.ToLower(query.Query)) &&
			!strings.Contains(streets[i].PlaceAKA, query.Query) {
			continue
		}

		filteredStreets = append(filteredStreets, streets[i])
	}

	return filteredStreets
}
