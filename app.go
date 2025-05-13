package main

// All Necessary imports
import (
	"net/http" // HTTP Package to allow http requests
	"slices"   // Slices package to allow slice operations

	"github.com/labstack/echo/v4" // Echo Framework
)

// Flight_array is a struct or custom object
// represented with one key inside called "Flight_array"

type flight struct {
	Flight_array [][]string `json:"Flight_array"`
}

// Main Function
func main() {
	e := echo.New() // New instance of echo framework
	e.POST("/flights", func(c echo.Context) error { // Post request with the endpoint "/flights"
		my_flights := new(flight) // New instance of flight struct (will be used to bind the input from the user)
		if err := c.Bind(my_flights); err != nil {
			return c.JSON(http.StatusBadRequest,
				map[string]string{"error": "Invalid input"})
		} // Bind (Get the input as payload) from the user and first check if its valid and doesn't have any errors
		final_arr := []string{}
		// final_arr a new empty array which will be used to store the unique results in order for the output
		for i := 0; i < len(my_flights.Flight_array); i++ { // loop through the first array (which contains sub arrays)
			for j := 0; j < len(my_flights.Flight_array[i]); j++ { // loop through each sub array
				// check if the current element is in the final_arr (this is to avoid duplicates)
				found := slices.Contains(final_arr, my_flights.Flight_array[i][j])
				// if it exist, don't do anything.
				if found {

				} else { // if it doesn't exist, add it to the final_arr (append result)
					final_arr = append(final_arr, my_flights.Flight_array[i][j])
				}
			}
		}
		return c.JSON(http.StatusOK, final_arr) // make sure we return the final_arr as a JSON
	})
	e.Logger.Fatal(e.Start(":1323")) // Start the server on port 1323
}
