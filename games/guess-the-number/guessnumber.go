package guessthenumber

import (
	"math/rand"
	"strconv"

	"github.com/bwmarrin/discordgo"
)

type GuessTheNumberGame struct {
	TargetNumber int
	IsGameOver   bool
}

func NewGuessTheNumberGame() *GuessTheNumberGame {
	return &GuessTheNumberGame{}
}

// StartGameGuess initializes and starts the Guess The Number game.
func (game *GuessTheNumberGame) StartGameGuess(s *discordgo.Session, m *discordgo.MessageCreate) {
	game.TargetNumber = rand.Intn(1000)

	s.ChannelMessageSend(m.ChannelID, "I've chosen a number between 0 and 1000. Try to guess it!")

	s.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		//Check if the message is from the bot, if it's a stop command,
		//if the game is already over, or if the user is requesting help
		if m.Author.ID == s.State.User.ID || m.Content == "!stop" || game.IsGameOver || m.Content == "!play guess-the-number help" {
			return
		}

		//Parse the user's guess from the message content
		guess, err := strconv.Atoi(m.Content)
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "Invalid input. Please enter a number.")
			return
		}

		if guess < 0 || guess > 1000 {
			s.ChannelMessageSend(m.ChannelID, "Write a number between 0 and 1000.")
		} else if guess == game.TargetNumber {
			s.ChannelMessageSend(m.ChannelID, "Congratulations! You guessed the number!")
			game.IsGameOver = true
			return
		} else if guess < game.TargetNumber {
			s.ChannelMessageSend(m.ChannelID, "Too low! Try a higher number.")
		} else {
			s.ChannelMessageSend(m.ChannelID, "Too high! Try a lower number.")
		}

		return
	})

}
