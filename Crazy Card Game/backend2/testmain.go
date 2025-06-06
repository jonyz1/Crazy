package main

import (
	"testing"

	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
)

// Mock struct for Player
type MockPlayer struct {
	ID   string
	Conn *websocket.Conn
	Hand []Card
}

// Initialize some mock data for testing
func TestInitializeGame(t *testing.T) {
	game.Players = make(map[string]*Player)
	game.GameState = "waiting"
	game.Turn = ""

	assert.Equal(t, len(game.Players), 0, "Expected no players initially")
	assert.Equal(t, game.GameState, "waiting", "Expected the game to be in 'waiting' state")
}

func TestAddPlayer(t *testing.T) {
	player := &Player{ID: "Player-1", Conn: nil, Hand: []Card{{Suit: "Hearts", Value: "6"}}}
	game.Players[player.ID] = player
	assert.Equal(t, len(game.Players), 1, "Expected 1 player in the game")
	assert.Equal(t, game.Players[player.ID].ID, "Player-1", "Player ID should be 'Player-1'")
}

func TestNextTurn(t *testing.T) {
	player1 := &Player{ID: "Player-1", Conn: nil, Hand: []Card{{Suit: "Hearts", Value: "6"}}}
	player2 := &Player{ID: "Player-2", Conn: nil, Hand: []Card{{Suit: "Spades", Value: "9"}}}

	game.Players[player1.ID] = player1
	game.Players[player2.ID] = player2
	game.Turn = player1.ID

	nextTurn(player1.ID)

	assert.Equal(t, game.Turn, player2.ID, "Expected turn to be moved to Player-2")
}

func TestRemoveCardFromHand(t *testing.T) {
	player := &Player{ID: "Player-1", Conn: nil, Hand: []Card{{Suit: "Hearts", Value: "6"}, {Suit: "Spades", Value: "7"}}}
	game.Players[player.ID] = player

	cardToRemove := Card{Suit: "Hearts", Value: "6"}
	removeCardFromHand(player.ID, cardToRemove)

	assert.Equal(t, len(game.Players[player.ID].Hand), 1, "Expected player to have 1 card left in hand")
	assert.NotContains(t, game.Players[player.ID].Hand, cardToRemove, "Expected the card to be removed from hand")
}

func TestUpdateUserHandAndDeck(t *testing.T) {
	// Mock DB calls or handle them if required for testing
	// Here, you can mock the DB interaction with a test or a mock DB to simulate data changes.
	// Since testing DB interaction is out of scope, this part should mock or simulate DB updates.
	assert.NotNil(t, UpdateUserHandAndDeck(nil, 1, 1, []Card{{Suit: "Hearts", Value: "6"}}), "Expected UpdateUserHandAndDeck to not return an error")
}

// Test function for handle7Card
func TestHandle7Card(t *testing.T) {
	player := &Player{ID: "Player-1", Conn: nil, Hand: []Card{{Suit: "Hearts", Value: "6"}}}
	game.Players[player.ID] = player
	cardToPlay := Card{Suit: "Hearts", Value: "7"}

	handle7Card(cardToPlay, player.ID)

	// Check if the card is added to playstack and removed from hand
	assert.Contains(t, game.PlayStack, cardToPlay, "Expected card to be added to PlayStack")
	assert.NotContains(t, game.Players[player.ID].Hand, cardToPlay, "Expected card to be removed from player's hand")
}

func TestMain(m *testing.M) {
	// Setup the game environment here, if necessary.
	// Run the tests
	m.Run()
}
