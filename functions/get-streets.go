package functions

import (
	"encoding/json"
	"fmt"
	"os"

	"dobefu/search-page/structs"
)

func GetStreets() []structs.Places {
	var streetData structs.StreetData
	streets := make([]structs.Places, 0)

	data, err := os.ReadFile("db/metatopos-places.json")

	if err != nil {
		fmt.Println(err)
		return streets
	}

	err = json.Unmarshal(data, &streetData)

	for i := 0; i < len(streetData.Places); i++ {
		streets = append(streets, streetData.Places[i])
	}

	return streets
}
