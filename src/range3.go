package main

import "fmt"

// START OMIT
func main() {
    continents := map[string]string{
        "United States": "North America",
        "India":         "Asia",
        "Namibia":       "Africa",
        "Germany":       "Europe",
        "Mexico":        "North America",
        "Ireland":       "Europe",
        "Iceland":       "Europe",
    }
    for country, continent := range continents { // HL
        fmt.Println(country, "is in", continent)
    }
}
// END OMIT
