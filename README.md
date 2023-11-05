## License and State

| ![WTFPL](https://upload.wikimedia.org/wikipedia/commons/0/0a/WTFPL_badge.svg) | WTFPL – Do What the Fuck You Want to Public License
| :-------- | :-------- |

| ![](https://img.shields.io/badge/Made_With-Go-blue) | ![Build](https://github.com/Cheveo/go-rest-cli/actions/workflows/go.yml/badge.svg?event=push) 
| :--------     | :-------    |


# GO REST CLI
```blue
..######....#######..........########..########..######..########..........######..##.......####
.##....##..##.....##.........##.....##.##.......##....##....##............##....##.##........##.
.##........##.....##.........##.....##.##.......##..........##............##.......##........##.
.##...####.##.....##.#######.########..######....######.....##....#######.##.......##........##.
.##....##..##.....##.........##...##...##.............##....##............##.......##........##.
.##....##..##.....##.........##....##..##.......##....##....##............##....##.##........##.
..######....#######..........##.....##.########..######.....##.............######..########.####
```

Is an opinionated tool to scaffold rest projects in golang.
One can create separate domains or whole webservice projects,
powered by:

	💎 Standard Library + gorilla mux
	💎 Gin & Gonic + Gorm

for Web Servers.


## Installation

simply use the go install command:

```bash
  go install https://github.com/Cheveo/go-rest-cli
```
    
## API Reference

```
Use "go-rest-cli [command] --help" for more information about a command.
```

#### Create a [Standard Library](https://pkg.go.dev/std) and [Gorilla Mux](https://github.com/gorilla/mux) based project

```bash
  go-rest-cli std-project
```

| Parameter     | Short       | Type      | Description             | IsRequired |
| :--------     | :-------    | :---------|:---------------         |:--------
| `--name`      | `-n`        | string    | Project Name            | **Required**
| `--module`    | `-m`        | string    | Go Module Name          | **Required**
| `--domain`    | `-d`        | string    | Domain Name             | **Required**
| `--directory` | `-p`        | string    | Path/Directory Name (if not set the current working directory is used)     | **Optional**

```
// resulting structure
.
└── my-project/
    ├── my-domain/
    │   ├── handler/
    │   │   ├── handler.go
    │   │   └── {{my-domain}}_handler.go
    │   ├── service/
    │   │   ├── service.go
    │   │   └── {{my-domain}}_service.go
    │   ├── models/
    │   │   └── {{my-domain}}_models.go
    │   └── storage/
    │       ├── storage.go
    │       ├── {{my-domain}}_sql_statements
    │       └── {{my-domain}}_storage.go
    ├── cmd/
    │   └── api/
    │       └── main.go
    ├── server/
    │   └── server.go
    ├── utils/
    │   ├── make_http_handler.go
    │   └── write_json.go
    └── go.mod 
```

#### Create a [Standard Library](https://pkg.go.dev/std) based domain

```bash
  go-rest-cli std-domain
```

| Parameter     | Short       | Type      | Description             | IsRequired |
| :--------     | :-------    | :---------|:---------------         |:--------
| `--module`    | `-m`        | string    | Go Module Name          | **Required**
| `--domain`    | `-d`        | string    | Domain Name             | **Required**
| `--directory` | `-p`        | string    | Path/Directory Name (if not set the current working directory is used)     | **Optional**

```
// Resulting directory structure
.
└── my-domain/
    ├── handler
    ├── service
    ├── models
    └── storage
```

#### Create a [Gin & Gonic](https://github.com/gin-gonic/gin) and [Gorm](https://gorm.io/index.html) based project

```bash
  go-rest-cli gin-project
```

| Parameter     | Short       | Type      | Description             | IsRequired |
| :--------     | :-------    | :---------|:---------------         |:--------
| `--name`      | `-n`        | string    | Project Name            | **Required**
| `--module`    | `-m`        | string    | Go Module Name          | **Required**
| `--domain`    | `-d`        | string    | Domain Name             | **Required**
| `--directory` | `-p`        | string    | Path/Directory Name (if not set the current working directory is used)     | **Optional**

```
// resulting structure
.
└── my-project/
    ├── my-domain/
    │   ├── handler/
    │   │   ├── handler.go
    │   │   └── {{my-domain}}_handler.go
    │   ├── service/
    │   │   ├── service.go
    │   │   └── {{my-domain}}_service.go
    │   ├── models/
    │   │   └── {{my-domain}}_models.go
    │   └── storage/
    │       ├── storage.go
    │       └── {{my-domain}}_storage.go
    ├── cmd/
    │   └── api/
    │       └── main.go
    ├── db/
    │   ├── db.go
    │   └── gorm_db.go
    ├── errors/
    │   └── http_error.go
    ├── middlewares/
    │   └── error_handler.go
    ├── responses/
    │   └── http_response.go
    └── go.mod
```

#### Create a [Gin & Gonic](https://github.com/gin-gonic/gin) and [Gorm](https://gorm.io/index.html) based domain

```bash
  go-rest-cli gin-domain
```

| Parameter     | Short       | Type      | Description             | IsRequired |
| :--------     | :-------    | :---------|:---------------         |:--------
| `--name`      | `-n`        | string    | Project Name            | **Required**
| `--module`    | `-m`        | string    | Go Module Name          | **Required**
| `--domain`    | `-d`        | string    | Domain Name             | **Required**
| `--directory` | `-p`        | string    | Path/Directory Name (if not set the current working directory is used)     | **Optional**

```
// Resulting directory structure
.
└── my-domain/
    ├── handler
    ├── service
    ├── models
    └── storage
```


## Why I created this

Since we use golang very often in our company to build web services and the basic structure is always the same, I decided to write a CLI tool that creates a web service project or even a whole domain from scratch.

The tool is influenced by the typical layered architecture in classic web services. Basically golang is often used for creating microservices, but often microservices are overkill while maintaining a good monolithic codebase has it's advantages. If you get to the point where scaling and performance are actually the bottlenecks, then it makes sense to switch to microservices, depending on the use case.

The tool is primarily intended for internal purposes and should help us to reach our goal faster. We have created the tool using the factory and strategy pattern to make the extension simple and adhere to the solid principle.

There is still a lot to do and many ways to extend it. We welcome anyone who is interested in contributing or simply using it.


## 🚀 About Me
I’m an enthusiastic web developer boasting over 6 years of experience crafting web applications.
I have a relentless drive for learning new techrelated concepts and techniques to continually enhance my engineering skills. I take pride in my strong communication abilities and enjoy engaging in discussions about effective implementations and architectural ideas. Please don't hesitate to contact me.


## 🔗 Links
[![portfolio](https://img.shields.io/badge/my_portfolio-000?style=for-the-badge&logo=ko-fi&logoColor=white)](https://cheveo.de/profiles/danyal-iqbal)

[![linkedin](https://img.shields.io/badge/linkedin-0A66C2?style=for-the-badge&logo=linkedin&logoColor=white)](https://www.linkedin.com/in/danyal-iqbal-3b46b6ab/)



## 🛠 Skills
using technologies such as:
* Angular
* Nx
* NestJS
* TypeScript
* NodeJS
* Golang
* Amazon Web Services
* Firebase


## Authors

- [@Sagato](https://github.com/Sagato)
- [@Cheveo](https://github.com/Cheveo)


## Contributing

Contributions are always welcome!

Kindly create a pull request or file an issue for feature requests.

Please adhere to this project's `code of conduct`.
