package models

import (
	// "fmt"
	"time"

	"github.com/jinzhu/gorm"
)

type Card struct {
	Suit  string `json:"suit"`
	Value string `json:"value"`
}

type Game struct {
	ID            uint      `json:"id" gorm:"primary_key"`
	DeckRemaining []Card    `json:"deckremaining" gorm:"type:json"` // Store as JSON
	TopCard       Card      `json:"topcard" gorm:"type:json"`       // Store as JSON
	StartAt       time.Time `json:"start_at"`
	EndAt         time.Time `json:"end_at"`
	PlayersID     []uint    `json:"playersid" gorm:"type:json"` // Store as JSON
	RoomID        uint      `json:"roomid"`
	Finished      bool      `json:"finished"`
	IP            string    `json:"ip"`
	Winner        uint      `json:"winner"`
}

type User struct {
	ID            uint   `json:"id" gorm:"primary_key"`
	Username      string `json:"username"`
	Password      string `json:"password"`
	GameIDs       []uint `json:"gameids" gorm:"type:json"` // Store as JSON
	CurrentGameID uint   `json:"currentgameid"`
	CurrentHand   []Card `json:"currenthand" gorm:"type:json"` // Store as JSON
}

type Room struct {
	ID            uint   `json:"id" gorm:"primary_key"`
	Name          string `json:"name"`
	CurrentGameID uint   `json:"currentgameid"`
	PlayerCount   uint   `json:"playercount"`
}

// Migrate the user model to create the users table
func MigrateUser(db *gorm.DB) {
	// Automatically migrate the schema
	// fmt.Println("not  migrating")
	db.AutoMigrate(&User{}, &Room{}, &Game{})
	// fmt.Println("finshed migrating")
}
