# URL Shortener
![Go](https://img.shields.io/badge/Go-1.24-00ADD8?style=for-the-badge&logo=go)
![Fiber](https://img.shields.io/badge/Fiber-v3-00C7B7?style=for-the-badge)
![Redis](https://img.shields.io/badge/Redis-Database-DC382D?style=for-the-badge&logo=redis)
![JWT](https://img.shields.io/badge/JWT-Authentication-000000?style=for-the-badge&logo=jsonwebtokens)
![Docker](https://img.shields.io/badge/Docker-Containerized-2496ED?style=for-the-badge&logo=docker)
![License](https://img.shields.io/badge/License-MIT-green?style=for-the-badge)

A production-ready URL Shortener built with Go, Fiber, Redis, and JWT Authentication. The application allows users to create custom short URLs, generate QR codes, track click analytics, and securely manage their links through an authenticated dashboard.

---

## Features

### Authentication

- User registration
- Secure login using JWT
- Protected API routes
- Logout functionality

### URL Management

- Shorten long URLs
- Create custom aliases
- Redirect using short URLs
- Delete existing URLs
- View all created URLs

### Analytics

- Track total click count
- View analytics for each URL

### QR Code

- Generate QR codes for every shortened URL
- Open QR code directly from the dashboard

### User Experience

- Responsive interface
- Toast notifications
- Loading states
- Delete confirmation modal
- Empty state UI
- Form validation
- Guest mode for unauthenticated users

---

## Tech Stack

### Backend

- Go
- Fiber
- Redis
- JWT Authentication

### Frontend

- HTML
- CSS
- JavaScript

### Tools

- Docker
- Git
- GitHub

---

## Project Structure

```
URL-Shortner
│
├── database/
├── helpers/
├── middleWare/
├── Models/
├── routes/
│
├── static/
│   ├── css/
│   ├── js/
│   ├── index.html
│   ├── login.html
│   └── signup.html
│
├── main.go
├── Dockerfile
├── go.mod
└── go.sum
```

---

## API Endpoints

### Authentication

| Method | Endpoint | Description |
|----------|-----------|-------------|
| POST | `/signup` | Register a new user |
| POST | `/login` | Authenticate user |

---

### URL Operations

| Method | Endpoint | Description |
|----------|-----------|-------------|
| POST | `/api/v1` | Create a shortened URL |
| GET | `/go/:url` | Redirect to original URL |
| GET | `/myurls` | Get all URLs of authenticated user |
| DELETE | `/url/:id` | Delete a shortened URL |

---

### Analytics

| Method | Endpoint | Description |
|----------|-----------|-------------|
| GET | `/analytics/:url` | Retrieve click statistics |

---

### QR Code

| Method | Endpoint | Description |
|----------|-----------|-------------|
| GET | `/qr/:url` | Generate QR code |

---

## Installation

### Clone the repository

```bash
git clone https://github.com/AnshuKashyap01/URL-Shortner.git
```

```bash
cd URL-Shortner/api
```

---

### Install dependencies

```bash
go mod download
```

---

### Configure Environment Variables

Create a `.env` file in the project root.

Example:

```env
APP_PORT=:3000
DB_ADDR=localhost:6379
JWT_SECRET=your_secret_key
API_QUOTA=10
DOMAIN=http://localhost:3000
```

---

### Start Redis

Ensure Redis is running before starting the application.

---

### Run the application

```bash
go run main.go
```

Visit:

```
http://localhost:3000
```

---

## Running with Docker

Build the image

```bash
docker build -t url-shortener .
```

Run the container

```bash
docker run -p 3000:3000 \
-e APP_PORT=:3000 \
-e DB_ADDR=host.docker.internal:6379 \
-e JWT_SECRET=your_secret_key \
-e API_QUOTA=10 \
-e DOMAIN=http://localhost:3000 \
url-shortener
```

---

## Screenshots

Add screenshots of the following pages.

- Home Page
- Login
- Signup
- Dashboard
- URL Created
- Analytics
- QR Code
- Delete Confirmation

---

## Future Improvements

- URL expiration management
- Password reset
- Email verification
- User profile management
- Search and filtering
- Custom domains
- Download analytics
- Dark mode

---

## Learning Outcomes

This project demonstrates practical experience with:

- REST API development
- Authentication using JWT
- Redis as a data store
- Backend development with Go
- Fiber framework
- Docker containerization
- Frontend integration with a Go backend
- Responsive user interface design
- CRUD operations
- Client-side validation

---

## License

This project is available under the MIT License.
