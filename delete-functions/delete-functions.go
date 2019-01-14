package main

import "fmt"

func display(countryCapitalMap map[string]string) {
	for country := range countryCapitalMap {
		fmt.Println("Capital of", country, "is", countryCapitalMap[country])
	}
}

func main() {
	countryCapitalMap := map[string]string{
		"France": "Paris",
		"Italy":  "Rome",
		"Japan":  "Tokyo",
		"India":  "New Delhi",
	}

	fmt.Println("---original map---")
	display(countryCapitalMap)

	//delete an entry
	delete(countryCapitalMap, "France")
	fmt.Println("entry for France is deleted")

	fmt.Println("---Updated map---")
	display(countryCapitalMap)

}
