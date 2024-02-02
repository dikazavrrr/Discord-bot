# Discord Bot with Games and Weather

This Discord bot is designed to provide various features, including games like Tic-Tac-Toe and Guess the Number, as well as fetching weather information.

## Getting Started

These instructions will help you set up and run the Discord bot on your server.

### Prerequisites

1. Go installed on your machine. You can download it from [here](https://golang.org/dl/).

2. Discord bot token. Obtain a token by creating a new bot on the [Discord Developer Portal](https://discord.com/developers/applications).

3. Weather API key. Get an API key from [WeatherAPI](https://www.weatherapi.com/).

### Installation

Clone the repository:

   ```bash
   git clone https://github.com/dikazavrrr/discord-bot.git
   cd discord-bot
   ```

### Run Locally

```bash
go run main.go
```

### Commands

!weather <city>: Get weather information for the specified city.

!play tic-tac-toe: Start a game of Tic-Tac-Toe.

!play guess-the-number: Start a game of Guess the Number.

!stop: End the current game.

### Project Structure

The project is organized into several directories to maintain a clean and modular structure:

- **bot**: Contains the main Discord bot logic.
  - *bot.go*: Implements the Discord bot's initialization and message handling.
  
- **games**: Houses the various game implementations.
  - **guess-the-number**: Implements the "Guess the Number" game.
    - *guessnumber.go*: Contains the logic for the Guess the Number game.

  - **tic-tac-toe**: Implements the "Tic-Tac-Toe" game.
    - *tictactoe.go*: Contains the logic for the Tic-Tac-Toe game.

- **config**: Manages configuration settings.
  - *config.go*: Defines a configuration struct and handles reading from a configuration file.

- **weather**: Handles fetching weather information from an external API.
  - *weather.go*: Contains the logic for fetching weather information.

- *main.go*: The main entry point for the Discord bot, responsible for starting the bot and handling configuration.

- *config.json*: Configuration file where sensitive data such as Discord bot token and API keys are stored. Make sure to replace placeholder values with actual tokens and keys.
