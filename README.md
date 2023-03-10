# StarSearch

StarSearch is a simple set of API endpoints to demonstrate Golang's strength in RESTful API development. It fetches data from [Swapi API](https://swapi.dev/), transforms and send them as a ```Response``` to different endpoints. See [demo and documentation](https://marvelous-florentine-950c63.netlify.app/#/)

## Installation

### Requirements
* [Go 1.19](https://go.dev/doc/install)
* [Postgresql 15](https://www.postgresql.org/docs/current/tutorial-install.html) (You can also install with ```sh install.sh db``` on Linux)
* [Redis 7](https://redis.io/docs/getting-started/installation)

#### For Linux users
To enable hot reload, install [air](https://github.com/cosmtrek/air): ```sh install.sh air```.  
To install the demo and documentation ```sh install.sh docs```.

#### Using Docker
You can install and launch the app using ```docker-compose up```

## Serving the app
#### On Linux
```bash
sh serve.sh app
```

#### On Windows
```bash
cd app && go run main.go
```
If you installed [air](https://github.com/cosmtrek/air), then you can do
```bash
cd app && air
```

## Building the app

#### On Linux
```bash
sh build.sh app
```

#### On Windows
```bash
cd app && go build -o starsearch main.go
```