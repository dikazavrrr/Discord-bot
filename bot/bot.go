package bot

import (
	"discordbot/config"
	guessthenumber "discordbot/games/guess-the-number"
	tictactoe "discordbot/games/tic-tac-toe"
	"discordbot/weather"
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var (
	isGameOver bool
)

func Start() {
	//Create a new DiscordGo session
	goBot, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	//Add the message handler function to handle incoming messages
	goBot.AddHandler(messageHandler)

	//Open a connection to Discord
	err = goBot.Open()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Bot is running")
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	//Ignore messages from the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	//Check if the message starts with "!".
	//To understand that this is a command
	if !strings.HasPrefix(m.Content, "!") {
		return
	}

	args := strings.Split(m.Content[1:], " ")

	switch args[0] {
	case "weather":
		//weather command
		city, temp, con := weather.Weather(args[1])
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("City: %+v\nTemperature: %+v°C\nCondition: %+v", city, temp, con))
	case "help":
		//help command
		helpMessage := "```\n" +
			"Available Commands:\n" +
			"!weather <city> - Get weather information for the specified city.\n" +
			"!help - Display this help message.\n" +
			"!play tic-tac-toe - Start a game of tic-tac-toe.\n" +
			"!play guess-the-number - Start a game of guess-the-number.\n" +
			"!stop - End the current game.\n" +
			"```"
		s.ChannelMessageSend(m.ChannelID, helpMessage)
	case "play":
		//Start game
		if strings.ToLower(args[1]) == "tic-tac-toe" {
			isGameOver = false
			//Rules for tic-tac-toe game
			if len(args) > 2 && strings.ToLower(args[2]) == "help" {
				isGameOver = true
				helpMessage := `
**Tic-Tac-Toe Game Help:**
- To start a new game, type: ` + "`!play tic-tac-toe`" + `
- To make a move, enter the coordinates of the cell you want to mark. For example: ` + "`A1`" + `, ` + "`B2`" + `, ` + "`C3`" + `
- To stop the game at any time, type: ` + "`!stop`" + `
		
**Coordinates:**
- Rows are denoted by letters A, B, C.
- Columns are denoted by numbers 1, 2, 3.
		
**Game Rules:**
- You are X, and the bot is O.
- Get three X's in a row (horizontally, vertically, or diagonally) to win.
- If the board is full and no one has three in a row, it's a draw.
		
Have fun and enjoy the game!
`

				s.ChannelMessageSend(m.ChannelID, helpMessage)
				return
			} else {
				//Start Tic-tac-toe game
				s.ChannelMessageSend(m.ChannelID, "Let's start! Tic-tac-toe!")
				handleTicTacToeCommand(s, m)
				isGameOver = true
			}

		} else if strings.ToLower(args[1]) == "guess-the-number" {
			//Rules for guess-the-number
			isGameOver = false

			if len(args) > 2 && strings.ToLower(args[2]) == "help" {
				isGameOver = true
				helpMessage := `
**Guess The Number Game Help:**
	- To start a new game, type: ` + "`!play guess-the-number`" + `
- To make a guess, simply type the number you want to guess.
	For example: ` + "`42`" + `
- To stop the game at any time, type: ` + "`!stop`" + `			
**Game Rules:**
	- The target number is randomly chosen between 0 and 1000.
- Your goal is to guess the correct number.
			
Have fun and happy guessing!
`

				s.ChannelMessageSend(m.ChannelID, helpMessage)
				return
			} else {
				//Start Guess the number
				s.ChannelMessageSend(m.ChannelID, "Let's start! Guess the number!")
				handleGuessTheNumberCommand(s, m)
				isGameOver = true
			}
		} else {
			s.ChannelMessageSend(m.ChannelID, "Invalid game choice. Please choose either 'tic-tac-toe' or 'guess-the-number'.")
		}
	case "stop":
		//stop game command
		isGameOver = true
		s.ChannelMessageSend(m.ChannelID, "Game over.")
		return
	}
}

var (
	ticTacToeGame      *tictactoe.TicTacToeGame
	guessTheNumberGame *guessthenumber.GuessTheNumberGame
)

func init() {
	//Initialize games
	ticTacToeGame = tictactoe.NewTicTacToeGame()
	guessTheNumberGame = guessthenumber.NewGuessTheNumberGame()
}

func handleTicTacToeCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	//Start the tic-tac-toe game

	if isGameOver {
		isGameOver = true
		s.ChannelMessageSend(m.ChannelID, "Game over.")
		return
	}

	ticTacToeGame.StartGameTicTacToe(s, m)

}

func handleGuessTheNumberCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	//Start the guess-the-number game
	if isGameOver {
		isGameOver = true
		s.ChannelMessageSend(m.ChannelID, "Game over.")
		return
	}

	guessTheNumberGame.StartGameGuess(s, m)
}
