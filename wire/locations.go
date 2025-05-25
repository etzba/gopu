package wire

// Location is a static location in the server memory
type Location struct {
	Id int `json:"id"`
	// Name of the location
	Name string `json:"name"`
	// Address is the street and number (zip and country)
	Address string `json:"address"`
	// Longtitude location coordinate
	Longtitude float64 `json:"longtitude"`
	// Latitude location coordinate
	Latitude float64 `json:"latitude"`
}

// Distance sent from client
type Distance struct {
	// Formula calculate the distance with specific formula (pythagors, haversine)
	Formula string `json:"formula"`
	// FromPlace take a name of a place in the world to calculate the distance from
	FromPlace string `json:"fromPlace"`
	// Longtitude client coordinate
	Longtitude float64 `json:"longtitude"`
	// Latitude client coordinate
	Latitude float64 `json:"latitude"`
}
