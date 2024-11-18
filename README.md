# Custom TikTok Video Player Backend

This project provides a backend API for a custom TikTok video player. The server is built with Golang, utilizing Gorilla Mux for routing, and it integrates with a Python microservice to fetch TikTok video data dynamically.

---

## Features
- **Category Search**: Search videos by hashtags or topics.
- **Trending Videos**: Fetch the latest trending TikTok videos.
- **Custom Video Playback**: Serve video data without watermarks.
- **Asynchronous Processing**: Handles background API calls for smooth playback.

---

## Prerequisites
- Go 1.20 or higher
---

## Getting Started

### 1. Clone the Repository
```bash
git clone https://github.com/your-repo/custom-tiktok-player-backend.git
cd custom-tiktok-player-backend
```

### 2. Set Up Environment Variables
Create a `.env` file in the project root with the following:
```plaintext
PYTHON_SERVER_URL=http://<python-service-url>
PORT=8000
```

### 3. Install Dependencies
```bash
go mod tidy
```

### 4. Run the Application
```bash
go run cmd/server/main.go
```

The server will start on the configured port (default: `8000`).

### 5. API Endpoints

| Method | Endpoint                              | Description                       |
|--------|---------------------------------------|-----------------------------------|
| GET    | `/category?hashtag={hashtag}`         | Fetch videos by hashtag          |
| GET    | `/trending`                           | Fetch trending videos            |

---



## Development Notes

- **Error Handling**: All handlers and services follow a consistent error-handling pattern for better debugging.
- **Python Integration**: The backend relies on a Python microservice for TikTok data. Ensure the Python service is running and accessible at the `PYTHON_SERVER_URL` specified in the `.env` file.

---

## Contributing
Contributions are welcome! Please submit a pull request or raise an issue for bugs or feature requests.

---

## License
This project is licensed under the MIT License.
```
