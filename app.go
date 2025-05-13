package main

import (
	"net/http"
	"github.com/labstack/echo/v4"
	
)

type flight struct {
  Flight_array  [][]string `json:"Flight_array"` 
}

func main() {
	e := echo.New()
	e.POST("/flights", func(c echo.Context) error {
		my_flights := new(flight)
		if err := c.Bind(my_flights); err != nil { 
        return c.JSON(http.StatusBadRequest, 
			map[string]string{"error": "Invalid input"}) 
		} 
		
		return c.JSON(http.StatusOK, my_flights) 
	})
	e.Logger.Fatal(e.Start(":1323"))
}