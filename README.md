# 🚀 Starship Stops Calculator

A Go-based application to calculate the number of resupply stops required for starships to traverse a given distance, using dynamic data from the [Star Wars API (SWAPI)](https://swapi.dev).

---

## 🌐 API Quick Test

**Live API URL**: [http://54.161.58.1:8080/calculate-stops](http://54.161.58.1:8080/calculate-stops)

### **Example Request**

**GET** `/calculate-stops/{distance}`

**Example URL**:
```
http://54.161.58.1:8080/calculate-stops/1000000
```

**Response**:
```json
{
    "distance": 1000000,
    "results": [
        {
            "name": "Calamari Cruiser",
            "stops": 0
        },
        {
            "name": "Executor",
            "stops": 0
        },
        {
            "name": "Star Destroyer",
            "stops": 0
        },
        {
            "name": "CR90 corvette",
            "stops": 1
        },
        {
            "name": "EF76 Nebulon-B escort frigate",
            "stops": 1
        },
        {
            "name": "Death Star",
            "stops": 3
        },
        {
            "name": "Millennium Falcon",
            "stops": 9
        },
        {
            "name": "Rebel transport",
            "stops": 11
        },
        {
            "name": "Imperial shuttle",
            "stops": 13
        },
        {
            "name": "Sentinel-class landing craft",
            "stops": 19
        },
        {
            "name": "Slave 1",
            "stops": 19
        },
        {
            "name": "A-wing",
            "stops": 49
        },
        {
            "name": "X-wing",
            "stops": 59
        },
        {
            "name": "B-wing",
            "stops": 65
        },
        {
            "name": "Y-wing",
            "stops": 74
        },
        {
            "name": "TIE Advanced x1",
            "stops": 79
        },
        {
            "name": "arc-170",
            "stops": 83
        }
    ]
}
```

---

## 📋 Features

- Fetches starship data from the SWAPI API with support for paginated responses.
- Calculates stops based on starship speed (`MGLT`) and consumables duration.
- Handles edge cases such as invalid input, missing data, and unreachable distances.

---

## 🛠️ Technologies

- **Language**: [Go (Golang)](https://golang.org)
- **API**: [Star Wars API (SWAPI)](https://swapi.dev)
- **Testing**: Table-driven tests, TDD methodology.
- **Docker**: Multi-stage builds for production and development.
- **Air**: Hot-reload for local development.

---

## 🚀 Getting Started

### Prerequisites

Ensure you have the following installed:

- [Go](https://golang.org/doc/install)
- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)

---

### Running Locally

1. **Clone the repository**:
    ```bash
    git clone https://github.com/your-username/starship-stops.git
    cd starship-stops
    ```

2. **Start the application in development mode**:
    ```bash
    docker compose up --build -d
    ```

3. **Access the application**:
    - **API Endpoint**: [http://localhost:8080/calculate-stops/{distance}](http://localhost:8080/calculate-stops/1000000)

---

## 🏭 Production Deployment

1. **Build the production image**:
    ```bash
    docker build -t starship-stops:prod --target production .
    ```

2. **Run the production container**:
    ```bash
    docker run -p 8080:8080 starship-stops:prod
    ```

3. **Test the API**:
    - Access the API using a web browser or a tool like `curl`:
      ```bash
      curl http://localhost:8080/calculate-stops/1000000
      ```

---

## 📂 Project Structure

```plaintext
📦 get-starships-stops
├── cmd
│   └── app
│       └── main.go           # Application entry point
├── internal
│   ├── api
│   │   ├── handlers          # HTTP handlers for API
│   │   ├── middleware        # Middleware (e.g., headers)
│   │   ├── models            # API request and response models
│   ├── config                # Application configuration
│   ├── domain                # Core business models
│   ├── parser                # Parsing utilities (distance, consumables)
│   ├── service               # Core business logic
│   │   └── swapi             # SWAPI client service and API interactions
├── tmp                       # Development artifacts (ignored in production)
└── .air.toml                 # Hot-reload configuration for development
```

---

## ✅ Tests

Run tests with:

```bash
go test ./...
```

### **Test Coverage**:

- **SWAPI Client**: Pagination, error scenarios, and response handling.
- **Business Logic**: Starship stop calculations for edge cases.
- **HTTP Handlers**: Input validation and error propagation.

---

## 💡 Design Decisions

- **Separation of Concerns**: Clear distinction between HTTP, and business logic.
- **Dependency Injection**: Mockable interfaces for testing (e.g., SWAPI client).
- **Resilience**: Graceful handling of invalid data and API errors.
- **Scalability**: Modular architecture for easy extension.

---

## 🌟 Conclusion

This application showcases a modular and scalable approach to building APIs in Go, focusing on clean architecture, robust testing, and production readiness. Whether you're calculating starship stops or exploring new possibilities, this project serves as a strong foundation for future enhancements.
