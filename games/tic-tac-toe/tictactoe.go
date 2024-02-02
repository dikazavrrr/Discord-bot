package tictactoe

import (
	"math/rand"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type TicTacToeGame struct {
	Board      [3][3]string
	IsGameOver bool
}

func NewTicTacToeGame() *TicTacToeGame {
	return &TicTacToeGame{}
}

// Create game board with empty cells
func initTicTacToeBoard(game *TicTacToeGame) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			game.Board[i][j] = "-"
		}
	}
}

// Send a message to the channel displaying the current game board
func displayTicTacToeBoard(s *discordgo.Session, channelID string, game *TicTacToeGame) {
	boardString := "```\n"
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			boardString += game.Board[i][j] + " "
		}
		boardString += "\n"
	}
	boardString += "```"
	s.ChannelMessageSend(channelID, boardString)
}

// Make a random move for the Tic-Tac-Toe bot (O)
func makeTicTacToeBotMove(s *discordgo.Session, channelID string, game *TicTacToeGame) {
	for {
		row := rand.Intn(3)
		col := rand.Intn(3)

		if game.Board[row][col] == "-" {
			game.Board[row][col] = "O"
			displayTicTacToeBoard(s, channelID, game)
			return
		}
	}
}

func checkTicTacToeWin(board [3][3]string, symbol string) bool {
	for i := 0; i < 3; i++ {
		if (board[i][0] == symbol && board[i][1] == symbol && board[i][2] == symbol) ||
			(board[0][i] == symbol && board[1][i] == symbol && board[2][i] == symbol) {
			return true
		}
	}

	if (board[0][0] == symbol && board[1][1] == symbol && board[2][2] == symbol) ||
		(board[0][2] == symbol && board[1][1] == symbol && board[2][0] == symbol) {
		return true
	}

	return false
}

// check if the board full
func checkTicTacToeFull(board [3][3]string) bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == "-" {
				return false
			}
		}
	}
	return true
}

// Start game
func (game *TicTacToeGame) StartGameTicTacToe(s *discordgo.Session, m *discordgo.MessageCreate) {
	initTicTacToeBoard(game)
	displayTicTacToeBoard(s, m.ChannelID, game)

	s.ChannelMessageSend(m.ChannelID, "Let's play Tic-Tac-Toe! You are X. Enter your moves using coordinates (e.g., A1, B2, C3).")

	s.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID || m.Content == "!stop" || game.IsGameOver || m.Content == "!play tic-tac-toe help" {
			return
		}

		move := strings.ToUpper(m.Content)
		if len(move) != 2 || (move[0] < 'A' || move[0] > 'C') || (move[1] < '1' || move[1] > '3') {
			s.ChannelMessageSend(m.ChannelID, "Invalid move. Please use coordinates (e.g., A1, B2, C3).")
			return
		}

		row := int(move[0] - 'A')
		col, _ := strconv.Atoi(string(move[1]))
		col--

		if game.Board[row][col] != "-" {
			s.ChannelMessageSend(m.ChannelID, "Cell already occupied. Please choose an empty cell.")
			return
		}

		game.Board[row][col] = "X"
		displayTicTacToeBoard(s, m.ChannelID, game)

		if checkTicTacToeWin(game.Board, "X") {
			s.ChannelMessageSend(m.ChannelID, "Congratulations! You win!")
			game.IsGameOver = true
			return
		}

		if checkTicTacToeFull(game.Board) {
			s.ChannelMessageSend(m.ChannelID, "It's a draw! The board is full.")
			game.IsGameOver = true
			return
		}

		makeTicTacToeBotMove(s, m.ChannelID, game)

		if checkTicTacToeWin(game.Board, "O") {
			s.ChannelMessageSend(m.ChannelID, "Sorry, you lost. Better luck next time!")
			return
		}

		if checkTicTacToeFull(game.Board) {
			s.ChannelMessageSend(m.ChannelID, "It's a draw! The board is full.")
			return
		}
	})

}
