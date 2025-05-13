package main

// All Necessary imports
import (
	"net/http" // HTTP Package to allow http requests
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

	// Post request with the endpoint "/flights"
	e.POST("/flights", func(c echo.Context) error {
		my_flights := new(flight) // New instance of flight struct (will be used to bind the input from the user)

		// Bind (Get the input as payload) from the user and first check if it's valid and doesn't have any errors
		if err := c.Bind(my_flights); err != nil {
			return c.JSON(http.StatusBadRequest,
				map[string]string{"error": "Invalid input"})
		}

		// Create maps to help reconstruct the path
		routeMap := make(map[string]string)    // map each source to destination
		destinations := make(map[string]bool)  // track destinations

		// Loop through all tickets and fill the maps
		for i := 0; i < len(my_flights.Flight_array); i++ {
			from := my_flights.Flight_array[i][0]
			to := my_flights.Flight_array[i][1]
			routeMap[from] = to
			destinations[to] = true
		}

		// Find the starting point (a source that is not a destination)
		start := ""
		for i := 0; i < len(my_flights.Flight_array); i++ {
			src := my_flights.Flight_array[i][0]
			if !destinations[src] {
				start = src
				break
			}
		}

		// Reconstruct the full route starting from the first source
		final_arr := []string{start}
		for {
			next, ok := routeMap[start]
			if !ok {
				break // stop when there's no next destination
			}
			final_arr = append(final_arr, next)
			start = next
		}

		// return the result as JSON
		return c.JSON(http.StatusOK, final_arr)
	})

	// Start the server on port 1323
	e.Logger.Fatal(e.Start(":1323"))
}
