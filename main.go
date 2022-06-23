// To run
// go run .
package main

// To import all at once
// go get .
import (
	"net/http"

	"github.com/gin-gonic/gin" //go get -u github.com/gin-gonic/gin
)

// superhero represents the type of data we have about a superhero.
type superhero struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Power   int64  `json:"power"`
	Special string `json:"special"`
}

// superheroes slice to initialize a few superheroes in our variable. Not using any database.
// Slice is just like an array having an index value and length, but the size of the slice is resized they are not in fixed-size just like an array.
var superheroes = []superhero{
	{ID: "1", Name: "Super-Man", Power: 2000, Special: "Flight, Superhuman Strength, X-Ray Vision, Super Speed"},
	{ID: "2", Name: "Bat-Man", Power: 500, Special: "Rich"},
	{ID: "3", Name: "Wonder-Woman", Power: 1600, Special: "Superhuman Strength, Speed"},
	{ID: "4", Name: "Iron-Man", Power: 1500, Special: "Flight, Weapons, Armor"},
	{ID: "5", Name: "Spider-Man", Power: 1100, Special: "Web, Armor, Cling to walls"},
}

func home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{ // H is a shortcut for map[string]interface{}
		"instructions": "Add '/superheroes' to the link",
	})
}

func getSuperheroes(c *gin.Context) {
	// Printing all the superheroes available in the data
	c.JSON(http.StatusOK, superheroes)
}

func addSuperhero(c *gin.Context) {
	// Creating a new object to structure superhero
	var newSuperhero superhero

	// Call BindJSON to bind the received JSON to newSuperhero
	// BindJSON adds the data provided by user to newSuperhero
	// This is kind of "try catch" concept
	if err := c.ShouldBindJSON(&newSuperhero); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "Bad Request",
		})
		return
	}

	// Add the new superhero to the slice.
	superheroes = append(superheroes, newSuperhero)
	// Printing the added superhero
	c.JSON(http.StatusCreated, newSuperhero)
}

func editSuperhero(c *gin.Context) {
	id := c.Param("id")

	// Creating a new object to structure superhero
	var editSuperhero superhero

	// Call BindJSON to bind the received JSON to newSuperhero
	// BindJSON adds the data provided by user to newSuperhero
	// This is kind of "try catch" concept
	if err := c.ShouldBindJSON(&editSuperhero); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "Bad Request",
		})
		return
	}

	for i, hero := range superheroes {
		if hero.ID == id {
			superheroes[i].Name = editSuperhero.Name
			superheroes[i].Power = editSuperhero.Power
			superheroes[i].Special = editSuperhero.Special

			c.JSON(http.StatusOK, editSuperhero)
			return
		}
	}

	// If the above statement doesn't return anything, that means the id is invalid
	c.JSON(http.StatusBadRequest, gin.H{
		"error":   true,
		"message": "Invalid",
	})
}

func removeSuperhero(c *gin.Context) {
	id := c.Param("id")

	for i, hero := range superheroes {
		if hero.ID == id {
			// arr := [100, 200, 300, 400, 500]
			// arr[:2] = [100, 200]
			// arr[2:] = [300, 400, 500]
			// arr[2+1:] = [400. 500]
			// [100, 200][400, 500]
			superheroes = append(superheroes[:i], superheroes[i+1:]...) // ... is required when writing 2 slices in append function

			c.JSON(http.StatusOK, gin.H{
				"message": "Item Deleted",
			})

			return
		}
	}

	// If the above statement doesn't return anything, that means the id is invalid
	c.JSON(http.StatusBadRequest, gin.H{
		"error":   true,
		"message": "Invalid",
	})
}

func main() {
	// Creates a gin router with default middleware: logger and recovery (crash-free) middleware
	router := gin.Default()
	router.GET("/", home)
	router.GET("/superheroes", getSuperheroes)
	router.POST("/superheroes", addSuperhero)
	router.PUT("/superheroes/:id", editSuperhero)
	router.DELETE("/superheroes/:id", removeSuperhero)

	router.Run() // By default it serves on :8080 unless a PORT environment variable was defined.
	// router.Run(":3000") for a hard coded port
}
