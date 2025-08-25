package boxer

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// Boxer represents a boxer in the database.
type Boxer struct {
	// --- System Fields ---
	ID        uuid.UUID `gorm:"type:char(36);primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// --- Basic Info ---
	Name    string `gorm:"type:varchar(100);not null;index" json:"name"`
	Sex     string `gorm:"type:varchar(10)" json:"sex"`
	Country string `gorm:"type:varchar(50)" json:"country"`

	// --- Fight Record ---
	Bouts      int `json:"bouts"`
	Wins       int `json:"wins"`
	WinsKO     int `json:"wko"`
	Losses     int `json:"losts"`
	Draws      int `json:"draw"`
	TitleWins  int `json:"title-wins"`
	TitleDraws int `json:"title-draw"`

	// --- Weight Class Info ---
	// CurrentWeightClass stores the name of the boxer's current division.
	CurrentWeightClass string `gorm:"type:varchar(50)" json:"current_weight_class"`

	// Divisions stores a list of WeightClassIDs where the boxer has fought.
	// It's stored as a JSON array in the database.
	Divisions datatypes.JSON `gorm:"type:json" json:"divisions"`

	ImageURL string `json:"image_url"`
}

// BeforeCreate is a GORM hook that runs before a new record is created.
// It generates a new UUID for the Boxer's ID.
func (b *Boxer) BeforeCreate(tx *gorm.DB) (err error) {
	if b.ID == uuid.Nil {
		b.ID = uuid.New()
	}
	return
}

// AddDivision adds a new weight class ID to the boxer's division list, avoiding duplicates.
func (b *Boxer) AddDivision(newDivisionID int) error {
	var divisions []int
	if b.Divisions != nil {
		if err := json.Unmarshal(b.Divisions, &divisions); err != nil {
			return err
		}
	}

	exists := false
	for _, id := range divisions {
		if id == newDivisionID {
			exists = true
			break
		}
	}

	if !exists {
		divisions = append(divisions, newDivisionID)
	}

	updatedDivisions, err := json.Marshal(divisions)
	if err != nil {
		return err
	}
	b.Divisions = updatedDivisions
	return nil
}
