# URL Shortener
![Go](https://img.shields.io/badge/Go-1.24-00ADD8?style=for-the-badge&logo=go)
![Fiber](https://img.shields.io/badge/Fiber-v3-00C7B7?style=for-the-badge)
![Redis](https://img.shields.io/badge/Redis-Database-DC382D?style=for-the-badge&logo=redis)
![JWT](https://img.shields.io/badge/JWT-Authentication-000000?style=for-the-badge&logo=jsonwebtokens)
![Docker](https://img.shields.io/badge/Docker-Containerized-2496ED?style=for-the-badge&logo=docker)
![License](https://img.shields.io/badge/License-MIT-green?style=for-the-badge)


A production-inspired URL shortening service built with Go, Fiber, and Redis. The application allows authenticated users to create and manage short URLs, generate QR codes, track click analytics, and securely manage their links through a responsive web interface.

## Live Demo

**Application:** https://url-shortner-xv7d.onrender.com

## Features

- User authentication with JWT
- Create short URLs with optional custom aliases
- Automatic URL expiration
- QR code generation for every shortened URL
- Click analytics
- Personal dashboard to manage URLs
- Delete URLs
- Rate limiting using Redis
- Responsive user interface
- Toast notifications and confirmation modals
- Docker support
- Redis Cloud integration
- Deployed on Render

---

## Tech Stack

### Backend

- Go
- Fiber
- Redis
- JWT Authentication
- bcrypt

### Frontend

- HTML
- CSS
- JavaScript

### Deployment

- Render
- Redis Cloud
- Docker

---

## Project Structure

```
.
├── database
├── helpers
├── middleWare
├── Models
├── routes
├── static
│   ├── css
│   ├── js
│   ├── index.html
│   ├── login.html
│   └── signup.html
├── main.go
├── Dockerfile
└── docker-compose.yml
```

---

## API Endpoints

| Method | Endpoint | Description |
|---------|----------|-------------|
| POST | `/signup` | Register a new user |
| POST | `/login` | Login and receive JWT |
| POST | `/api/v1` | Create a short URL |
| GET | `/go/:url` | Redirect to original URL |
| GET | `/analytics/:url` | Get click analytics |
| GET | `/myurls` | Retrieve user's URLs |
| DELETE | `/url/:id` | Delete a URL |
| GET | `/qr/:url` | Generate QR Code |

---

## Environment Variables

Create a `.env` file in the project root.

```env
APP_PORT=3000

DB_ADDR=your_redis_host:port
DB_USER=default
DB_PASS=your_redis_password

DOMAIN=http://localhost:3000

JWT_SECRET=your_secret_key

API_QUOTA=10
```

---

## Running Locally

### Clone the repository

```bash
git clone https://github.com/AnshuKashyap01/URL_Shortner.git
```

```bash
cd URL_Shortner
```

### Install dependencies

```bash
go mod download
```

### Configure environment variables

Create a `.env` file using the template above.

### Run the application

```bash
go run main.go
```

Open:

```
http://localhost:3000
```

---

## Docker

Build the image

```bash
docker build -t url-shortener .
```

Run the container

```bash
docker run -p 3000:3000 url-shortener
```

---

## Performance Optimizations

- Singleton Redis client using `sync.Once`
- Single API request for dashboard data (eliminated N+1 requests)
- Instant UI updates after URL deletion without reloading the table
- Redis connection pooling
- Client-side form validation
- JWT-protected API endpoints

---

## Security

- Password hashing using bcrypt
- JWT-based authentication
- Protected routes using middleware
- Redis-backed rate limiting
- URL validation before shortening

---

## Future Improvements

- Search and filter URLs
- URL edit functionality
- Analytics dashboard with charts
- Unit and integration testing
- Custom domains
- URL expiration management interface

---

## Screenshots

Add screenshots of:

- Home Page
- Login Page
- Signup Page
- Dashboard
- QR Code Generation

---

## Author

**Anshu Kashyap**

GitHub: https://github.com/AnshuKashyap01

LinkedIn: https://www.linkedin.com/in/anshu-kashyap01

---

## License

This project is licensed under the MIT License.