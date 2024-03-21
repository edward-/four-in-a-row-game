# 4 in A Row [GAME]

<img src="./4inarow.png" alt="drawing" width="500"/>

## Game Description

4inARow is a digital rendition of the classic connect-four game where two players take turns dropping colored discs from the top into a seven-column, six-row vertically suspended grid. The objective is to be the first to form a horizontal, vertical, or diagonal line of four of one's own discs.

This game is built in Go, offering a a group of endpoints for the experience. 4inARow combines strategy, foresight, and a bit of luck, making each game unique and engaging.

## How to Start
### Prerequisites
- Go 1.22 or higher installed on your machine.
- Docker Desktop 4.28.0 
### Run

you could need to install `migrate-cli`:
```bash
make migrate-install
```
or
```bash
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```
#### Run local
```bash
   make compose-up
   make migrate-up
   make run
```

#### Run tests
```bash
   make compose-up
   make migrate-test-up
   make test
```

#### Run all in docker
```bash
   make compose-app-up
```

Now, you can start playing the game calling some endpoints [I need a cli, in the future :)]. endpoints:

```
GET    /ping
POST   /v1/users
POST   /v1/games
GET    /v1/games/:gameId/board
POST   /v1/games/:gameId/turn
```

## Architecture

This project adopts the Clean Architecture principles, ensuring that the game's design is independent of frameworks, UI, and databases. By segregating the system into layers (Domain, Infrastructure, Interface Adapters, and Frameworks & Drivers), we enhance maintainability, scalability, and the potential for future feature expansions. This architectural choice facilitates unit testing and decouples business logic from device-specific implementations.

## Stack Tech Used

- Programming Language: Go (Golang) - chosen for its simplicity, performance, and efficient concurrency support.
- Game Logic: The core game logic is implemented in pure Go, focusing on simplicity and performance to handle game states, player turns, and victory conditions efficiently.

## Docs&Tech
- [microservices](https://microservices.io/)
- [go.dev](https://go.dev/)
- [Go Style](https://google.github.io/styleguide/go/)
- [golang-migrate](https://github.com/golang-migrate/migrate)
- [redis](https://redis.com/)
- [postgeSQL](https://www.postgresql.org/)
- [gorm](gorm.io/gorm)
- [docker](https://www.docker.com/)
- [viper](github.com/spf13/viper)
- [ginkgo](github.com/onsi/ginkgo)
- [gin](github.com/gin-gonic/gin)

## TODO
- [ ] Implement swagger documentation
- [ ] Deploy to AWS in a lambda function with terraform
- [ ] Validate tests with github-actions
- [ ] Implement AI opponent with varying difficulty levels.
- [ ] Add cli functionality.
- [ ] Create a graphical user interface (GUI) version using a Go-based GUI library like Fyne.
- [ ] Optimize performance for handling larger grids or multiple simultaneous games.

## Contributing
We welcome contributions of all kinds from the community! If you're interested in making 4inARow even better, feel free to fork the repository, make your changes, and submit a pull request. For more details, check out our CONTRIBUTING.md file (to be created).

License
4inARow[Game] is released under the MIT License. See the LICENSE file in the repository for more details.