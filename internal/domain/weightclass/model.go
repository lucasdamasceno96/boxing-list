package weightclass

// WeightClassID is a type for our weight class enum.
type WeightClassID int

// Enum for all weight classes, used for internal logic.
const (
	Strawweight WeightClassID = iota + 1 // Starts with 1
	LightFlyweight
	Flyweight
	SuperFlyweight
	Bantamweight
	SuperBantamweight
	Featherweight
	SuperFeatherweight
	Lightweight
	SuperLightweight
	Welterweight
	SuperWelterweight
	Middleweight
	SuperMiddleweight
	LightHeavyweight
	Cruiserweight
	Heavyweight
)

// Info stores all details for a single weight class.
type Info struct {
	ID        WeightClassID
	Name      string  // Official name of the weight class
	WeightLBS float32 // Weight limit in pounds
}

// AllClasses is our single source of truth, a list of all official weight classes.
var AllClasses = []Info{
	{ID: Strawweight, Name: "Strawweight", WeightLBS: 105},
	{ID: LightFlyweight, Name: "Light Flyweight", WeightLBS: 108},
	{ID: Flyweight, Name: "Flyweight", WeightLBS: 112},
	{ID: SuperFlyweight, Name: "Super Flyweight", WeightLBS: 115},
	{ID: Bantamweight, Name: "Bantamweight", WeightLBS: 118},
	{ID: SuperBantamweight, Name: "Super Bantamweight", WeightLBS: 122},
	{ID: Featherweight, Name: "Featherweight", WeightLBS: 126},
	{ID: SuperFeatherweight, Name: "Super Featherweight", WeightLBS: 130},
	{ID: Lightweight, Name: "Lightweight", WeightLBS: 135},
	{ID: SuperLightweight, Name: "Super Lightweight", WeightLBS: 140},
	{ID: Welterweight, Name: "Welterweight", WeightLBS: 147},
	{ID: SuperWelterweight, Name: "Super Welterweight", WeightLBS: 154},
	{ID: Middleweight, Name: "Middleweight", WeightLBS: 160},
	{ID: SuperMiddleweight, Name: "Super Middleweight", WeightLBS: 168},
	{ID: LightHeavyweight, Name: "Light Heavyweight", WeightLBS: 175},
	{ID: Cruiserweight, Name: "Cruiserweight", WeightLBS: 200},
	{ID: Heavyweight, Name: "Heavyweight", WeightLBS: 0}, // No upper limit
}

// -- Helper Functions --

var classesByName = make(map[string]Info)

// init runs once when the package is loaded, creating a map for fast lookups.
func init() {
	for _, class := range AllClasses {
		classesByName[class.Name] = class
	}
}

// GetByName efficiently finds a weight class by its name.
func GetByName(name string) (Info, bool) {
	class, ok := classesByName[name]
	return class, ok
}
