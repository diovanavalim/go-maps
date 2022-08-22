package maps

import "fmt"

func CreateMaps() {
	userMap := map[string]string{
		"name":     "Diovana",
		"lastName": "Valim",
	}

	fmt.Println(userMap)

	userMap["age"] = "21"

	fmt.Println(userMap)
}

func CreateNestedMaps() {
	userMap := map[string]map[string]string{
		"name": {
			"firstName": "Adriana",
			"lastName":  "Valim",
		},
		"age": {
			"age": "52",
		},
		"address": {
			"street": "Foo",
			"number": "Bar",
		},
	}

	fmt.Println(userMap["name"]["firstName"])
	fmt.Println(userMap["name"]["lastName"])
	fmt.Println(userMap["age"])
	fmt.Println(userMap["address"])
}

func MapsMethods() {
	userMap := map[string]string{
		"name":       "Diovana",
		"lastName":   "Valim",
		"nickname":   "Dio",
		"occupation": "Software Developer",
		"age":        "21",
		"address":    "Foo Bar",
	}

	fmt.Println(userMap)

	delete(userMap, "address")

	fmt.Println(userMap)

	userMap["sign"] = "aries"

	fmt.Println(userMap)
}
