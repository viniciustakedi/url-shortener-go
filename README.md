# 📦 URL Shortner API

A RESTful API developed in Go to genereate short URLs in order to share short links to other people. With this API you are able to create short urls and don't worry about delete them, because the API has a CronJob that delete all expired urls every day.

## 🚀 Technologies Used

- **Go (Golang)**
- **Gin** – Lightweight and fast web framework
- **MongoDB** – NO-SQL database
- **Docker** – Containerization for consistent environments
- **Godotenv** & **Viper** – To config env vars and environments
- **Go-playground/validator** – Validate request payload
- **Robfig/cron/v3** – Used to create, start and execute cronjobs

## 📁 Project Structure

```
url-shortener-go/
├── api/           # API handlers and controllers
├── config/        # Configuration and environment variables
├── cron/          # Simple folder to organlize and create and store all cronjobs
├── infra/         # Infrastructure (e.g., database connection)
├── middlewares/   # Custom middlewares (e.g., authentication)
├── router/        # API route definitions
├── server/        # Server initialization and configuration
├── utils/         # Utility functions
├── main.go        # Application entry point
├── Dockerfile     # Dockerfile for containerization
└── .env.example   # Example environment variables
```

## ⚙️ Setup and Execution

### Prerequisites

- [Go](https://golang.org/doc/install) installed
- [Docker](https://www.docker.com/get-started) (optional, for containerization)

### Running Locally

1. Clone the repository:

```bash
git clone https://github.com/viniciustakedi/url-shortener-go.git
cd url-shortener-go
```

2. Copy the `.env.example` file to `.env.development` and configure the environment variables as needed.

3. Install dependencies and run the application:

```bash
go mod tidy
go run main.go
```

The API will be available at `http://localhost:8080/api`.

### Using Docker

1. Build the Docker image:

```bash
docker build -t url-shortener .
```

2. Copy the `.env.example` file to `.env.development` and configure the environment variables as needed.

3. Run the container:

```bash
docker run -d -p 8080:8080 --env-file .env.development url-shortener
```

The API will be available at `http://localhost:8080/api`.

## 🛠️ Main Endpoints

- `GET /api/health` – Just to check if API is running.
- `POST /api/url/shorten` – Send a long URL and the endpoint will return a short one.
- `GET /api/url/:shortUrl` – Call this endpoint with a existent short url code to get the original. Useful to front.

## 🔨 TO-DO List
This project I want to build a sign-up and login system, some way to track access to link and generate a QRCode. Logged users can generate a life-time short url.

## 🤝 Contributing

Contributions are welcome! Please open an issue or submit a pull request for any enhancements or bug fixes.

## 📄 License

This project is Open Source, so if you wanna just clone and make changes in order to improve this project and help others people, feel free to open a PR! 
