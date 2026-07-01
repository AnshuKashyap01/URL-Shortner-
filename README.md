# URL Shortener API

A production-inspired URL Shortener built with Go, Fiber, Redis, JWT Authentication and Docker.

![Go](https://img.shields.io/badge/Go-1.24-00ADD8?logo=go&logoColor=white)
![Fiber](https://img.shields.io/badge/Fiber-v3-00C853)
![Redis](https://img.shields.io/badge/Redis-DC382D?logo=redis&logoColor=white)
![Docker](https://img.shields.io/badge/Docker-2496ED?logo=docker&logoColor=white)
![JWT](https://img.shields.io/badge/JWT-Authentication-orange)
![REST API](https://img.shields.io/badge/REST-API-success)
![QR Code](https://img.shields.io/badge/QR-Code-blueviolet)
![Analytics](https://img.shields.io/badge/Analytics-Enabled-green)
# URL Shortener API

A production-inspired URL Shortener built with **Go**, **Fiber**, **Redis**, **JWT Authentication**, and **Docker**. The application allows users to create and manage shortened URLs with authentication, analytics, rate limiting, QR code generation, and user-specific URL management.

---

## Features

### Authentication
- User Signup
- User Login
- JWT Authentication
- Password hashing using bcrypt
- Protected Routes

### URL Shortening
- Generate random short URLs
- Custom short URLs
- URL expiration support
- Automatic HTTPS enforcement

### User Management
- View all URLs created by the authenticated user
- Delete owned URLs
- Ownership verification using JWT

### Analytics
- Click tracking for every shortened URL
- Analytics endpoint to retrieve click count

### QR Code Generation
- Generate a QR code for every shortened URL
- Access QR codes using `/qr/:url`

### Security
- JWT Authentication
- Password hashing with bcrypt
- IP-based rate limiting using Redis

### Infrastructure
- Redis as the primary datastore
- Docker & Docker Compose support
- RESTful API architecture

---

## Tech Stack

- Go
- Fiber
- Redis
- JWT
- bcrypt
- Docker
- Docker Compose

---

## Project Structure

```
.
├── api
│   ├── database
│   ├── helpers
│   ├── middleWare
│   ├── Models
│   ├── routes
│   ├── .env
│   ├── Dockerfile
│   ├── go.mod
│   └── main.go
│
├── db
│   └── Dockerfile
│
├── docker-compose.yml
└── README.md
```

---

## API Endpoints

| Method | Endpoint | Description | Protected |
|----------|----------------|---------------------------------|------------|
| POST | `/signup` | Register a new user | No |
| POST | `/login` | Login and receive JWT | No |
| POST | `/api/v1` | Create a shortened URL | Yes |
| GET | `/:url` | Redirect to original URL | No |
| GET | `/qr/:url` | Generate QR Code | No |
| GET | `/analytics/:url` | Retrieve click analytics | No |
| GET | `/myurls` | Retrieve all URLs of logged-in user | Yes |
| DELETE | `/url/:id` | Delete a URL owned by the user | Yes |

---

## Running the Project

### Clone the repository

```bash
git clone https://github.com/<your-username>/URLShortner.git
cd URLShortner
```

### Start using Docker

```bash
docker compose up --build
```

The API will be available at

```
http://localhost:3000
```

---

## Environment Variables

Create a `.env` file inside the **api** directory.

```env
APP_PORT=:3000

DOMAIN=localhost:3000

API_QUOTA=10

JWT_SECRET=your_secret_key
```

---

## Example Workflow

### Register

```http
POST /signup
```

```json
{
    "username":"anshu",
    "email":"anshu@gmail.com",
    "password":"123456"
}
```

---

### Login

```http
POST /login
```

Response

```json
{
    "token":"<JWT_TOKEN>"
}
```

---

### Create a Short URL

```http
POST /api/v1
```

Headers

```
Authorization: Bearer <JWT_TOKEN>
```

Body

```json
{
    "url":"https://github.com",
    "short":"github"
}
```

---

### Redirect

```
GET /github
```

Redirects to

```
https://github.com
```

---

### Generate QR Code

```
GET /qr/github
```

Returns a PNG QR Code pointing to the shortened URL.

---

### View Analytics

```http
GET /analytics/github
```

Example Response

```json
{
    "short_url":"github",
    "clicks":15
}
```

---

### View User URLs

```http
GET /myurls
```

Returns all URLs created by the authenticated user.

---

### Delete URL

```http
DELETE /url/github
```

Deletes the URL if it belongs to the authenticated user.

---

## Redis Database Usage

| Redis DB | Purpose |
|-----------|------------------------------|
| DB 0 | URL Storage |
| DB 1 | Analytics & Rate Limiting |
| DB 2 | User Authentication |

---

## Security Features

- Password hashing using bcrypt
- JWT Authentication
- Protected API endpoints
- User ownership verification
- Redis-based IP rate limiting

---

## Future Improvements

- Password-protected URLs
- Browser & Device Analytics
- Unique Visitor Tracking
- Geo-location Analytics
- URL Health Monitoring
- URL Editing
- Swagger API Documentation
- CI/CD Pipeline
- Cloud Deployment

---

## License

This project is intended for educational purposes and backend development practice.
