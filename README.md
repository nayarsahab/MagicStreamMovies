# MagicStream - Full-Stack Movie Streaming & Recommendation Platform

MagicStream is a modern, full-stack movie discovery and streaming web application built with a **React + Vite** frontend and a **Go (Golang) + Gin** backend, backed by **MongoDB** and integrated with **OpenAI (via LangChainGo)** for automated review sentiment ranking.

---

## 📑 Table of Contents

- [Overview](#overview)
- [Features](#features)
- [Architecture & Tech Stack](#architecture--tech-stack)
- [Project Structure](#project-structure)
- [Prerequisites](#prerequisites)
- [Environment Configuration](#environment-configuration)
  - [Backend (.env)](#backend-environment-variables)
  - [Frontend (.env)](#frontend-environment-variables)
- [Getting Started](#getting-started)
  - [1. Backend Setup](#1-backend-setup-servergo)
  - [2. Frontend Setup](#2-frontend-setup-clientreact)
- [API Endpoints](#api-endpoints)
  - [Public Endpoints](#public-routes)
  - [Protected Endpoints (JWT Required)](#protected-routes-jwt-required)
- [License](#license)

---

## 🎬 Overview

MagicStream allows users to browse an extensive catalog of movies, stream trailers/videos, register with favorite genres to receive personalized recommendations, and manage reviews. An AI sentiment analysis pipeline automatically ranks movies based on admin reviews using OpenAI's language models.

---

## ✨ Features

- **Movie Discovery & Catalog**: Explore movies with posters, backdrops, genres, and details.
- **Embedded Video Streaming**: Watch trailers and content using an integrated video player (`react-player`).
- **User Authentication & Authorization**:
  - Secure registration and login using JWT (Access & Refresh tokens).
  - Password hashing via `bcrypt`.
  - Token refresh handling with Axios private interceptors.
- **Personalized Recommendations**: Dynamic movie suggestions tailored to user-selected favorite genres.
- **AI-Powered Review Ranking**: Admin reviews are evaluated by OpenAI (via `LangChainGo`) to automatically assign sentiment and ranking values.
- **Responsive UI**: Built with React 19, React Bootstrap, and FontAwesome for seamless cross-device viewing.

---

## 🛠️ Architecture & Tech Stack

### Frontend (`ClientReact/magic-stream-client`)
- **Framework**: React 19 + Vite
- **Routing**: React Router v7
- **Styling & UI**: Bootstrap 5, React-Bootstrap, FontAwesome Icons
- **HTTP Client**: Axios with custom interceptors for automatic JWT refresh
- **Media**: React Player

### Backend (`ServerGo/MagicStreamMoviesServer`)
- **Language**: Go (Golang 1.27+)
- **Web Framework**: Gin (`github.com/gin-gonic/gin`)
- **Database**: MongoDB (using the official MongoDB Go Driver v2)
- **AI Integration**: LangChainGo (`github.com/tmc/langchaingo/llms/openai`)
- **Authentication**: JWT (`github.com/golang-jwt/jwt/v5`) & `golang.org/x/crypto/bcrypt`
- **Validation**: `go-playground/validator/v10`
- **CORS**: `github.com/gin-contrib/cors`

---

## 📁 Project Structure

```text
MovieStreaming/
├── .gitignore
├── LICENSE
├── README.md
├── ClientReact/
│   └── magic-stream-client/
│       ├── public/
│       ├── src/
│       │   ├── api/
│       │   │   ├── axiosConfig.jsx         # Axios base instance
│       │   │   └── axiosPrivateConfig.jsx  # Axios instance with auth interceptors
│       │   ├── assets/                     # Images, logos, SVG icons
│       │   ├── components/
│       │   │   ├── header/                 # Navigation header
│       │   │   ├── home/                   # Home page & movie carousel/cards
│       │   │   ├── login/                  # Login component
│       │   │   ├── register/               # Registration & favorite genres selection
│       │   │   ├── movie/                  # Movie detail view
│       │   │   ├── movies/                 # Movie list display
│       │   │   ├── recommended/            # Recommended movies view
│       │   │   ├── review/                 # Movie review & rating management
│       │   │   ├── spinner/                # Loading indicator
│       │   │   ├── stream/                 # Video streaming player
│       │   │   ├── Layout.jsx              # Main layout wrapper
│       │   │   └── RequiredAuth.jsx        # Protected route guard
│       │   ├── context/
│       │   │   └── AuthProvider.jsx        # Global authentication state
│       │   ├── hooks/
│       │   │   ├── useAuth.jsx             # Auth hook
│       │   │   └── useAxiosPrivate.jsx      # Authenticated API request hook
│       │   ├── App.jsx                     # Root application routes
│       │   └── main.jsx                    # Application entry point
│       ├── .env                            # Client environment variables
│       ├── package.json                    # Node dependencies & scripts
│       └── vite.config.js                  # Vite configuration
│
└── ServerGo/
    └── MagicStreamMoviesServer/
        ├── controllers/
        │   ├── movieController.go          # Movie retrieval, genres, recommendations & AI review
        │   └── userController.go           # Registration, authentication & token management
        ├── database/
        │   └── databaseConnection.go       # MongoDB client initialization & collection helpers
        ├── middleware/
        │   └── authMiddleware.go           # JWT verification middleware
        ├── models/
        │   ├── movieModel.go               # Movie and Genre data models
        │   └── userModel.go                # User and Token data models
        ├── routes/
        │   ├── protectedRoutes.go          # Authenticated routes
        │   └── unprotectedRoutes.go        # Public routes
        ├── utils/
        │   └── tokenUtil.go                # JWT generation, validation & claims handling
        ├── .env                            # Server environment variables
        ├── go.mod                          # Go module dependencies
        ├── go.sum                          # Go checksums
        └── main.go                         # Server entry point & CORS configuration
```

---

## 📋 Prerequisites

Before running the project locally, ensure you have installed:
- **Node.js**: v18.x or higher
- **Go**: v1.22 or higher (v1.27+ supported)
- **MongoDB**: Local MongoDB instance or MongoDB Atlas cluster URI
- **OpenAI API Key**: (Optional, required for AI review sentiment ranking)

---

## ⚙️ Environment Configuration

### Backend Environment Variables
Create or update `ServerGo/MagicStreamMoviesServer/.env`:

```env
# Server Port & CORS
PORT=8080
ALLOWED_ORIGINS=http://localhost:5173

# MongoDB Connection
MONGODB_URI=mongodb+srv://<username>:<password>@<cluster>.mongodb.net/?retryWrites=true&w=majority
DATABASE_NAME=MagicStreamDB

# JWT Secrets
SECRET_KEY=your_jwt_secret_key_here
SECRET_REFRESH_KEY=your_jwt_refresh_secret_key_here

# Recommendation Settings
RECOMMENDED_MOVIE_LIMIT=5

# OpenAI & AI Ranking Configuration
OPENAI_API_KEY=your_openai_api_key_here
BASE_PROMPT_TEMPLATE="You are a sentiment analyzer. Given the list of sentiments: [{rankings}], analyze the following movie review and respond with ONLY one matching sentiment word without punctuation: "
```

### Frontend Environment Variables
Create or update `ClientReact/magic-stream-client/.env`:

```env
VITE_API_BASE_URL=http://localhost:8080
```

---

## 🚀 Getting Started

### 1. Backend Setup (`ServerGo`)

1. Open a terminal and navigate to the server directory:
   ```bash
   cd ServerGo/MagicStreamMoviesServer
   ```

2. Download Go dependencies:
   ```bash
   go mod download
   ```

3. Start the backend server:
   ```bash
   go run main.go
   ```
   *The server will start listening on port `8080` (e.g., `http://localhost:8080`).*

---

### 2. Frontend Setup (`ClientReact`)

1. Open a new terminal and navigate to the client directory:
   ```bash
   cd ClientReact/magic-stream-client
   ```

2. Install dependencies:
   ```bash
   npm install
   ```

3. Launch the Vite development server:
   ```bash
   npm run dev
   ```
   *Open `http://localhost:5173` in your browser to view the application.*

---

## 🔌 API Endpoints

### Public Routes
| Method | Endpoint    | Description |
| :---   | :---        | :---        |
| `GET`  | `/hello`    | Health check endpoint |
| `GET`  | `/movies`   | Fetch all movies |
| `GET`  | `/genres`   | Fetch all available movie genres |
| `POST` | `/register` | Register a new user account with favorite genres |
| `POST` | `/login`    | Authenticate user and receive access & refresh tokens |
| `POST` | `/refresh`  | Issue a new access token using a valid refresh token |
| `POST` | `/logout`   | Invalidate user session and tokens |

### Protected Routes (JWT Required)
*Header: `Authorization: Bearer <access_token>`*

| Method  | Endpoint                 | Description |
| :---    | :---                     | :---        |
| `GET`   | `/movie/:imdb_id`        | Retrieve single movie details by IMDb ID |
| `POST`  | `/addmovie`              | Add a new movie to the database |
| `GET`   | `/recommendedmovies`     | Retrieve movies recommended based on user's favorite genres |
| `PATCH` | `/updatereview/:imdb_id` | Submit an admin review; triggers OpenAI sentiment analysis & ranking |

---

## 🔒 Security Best Practices
- Never commit `.env` files containing production secrets or database credentials to version control.
- Ensure unique and cryptographically secure values for `SECRET_KEY` and `SECRET_REFRESH_KEY`.

---

## 📄 License

This project is licensed under the [MIT License](LICENSE). See the [LICENSE](LICENSE) file for full details.

