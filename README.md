# Go Auth API

Basic API in **Go** with **JWT** authentication and **PostgreSQL** database.  
Currently includes endpoints:

- `POST /signup` – create user  
- `POST /login` – authenticate user  
- `GET /me` – get logged-in user data  
- `GET /ping` – test endpoint with date/time

---

## Technologies

- Go (Golang)  
- Gin (HTTP framework)  
- GORM (ORM)  
- PostgreSQL  
- Docker & Docker Compose  

---

## Running the project

1. Clone the repository:

```bash
git clone <REPOSITORY_URL>
cd go-auth-api
```

2. Configure environment variables in the `.env` file (example already included).  

3. Start the containers:

```bash
docker compose up --build
```

4. Access the endpoints:

- `http://localhost:9000/ping`  
- `http://localhost:9000/signup`  
- `http://localhost:9000/login`  
- `http://localhost:9000/me`  

> The port `9000` is the one the API listens on inside the container.

---

## Notes

- The database is initialized with PostgreSQL 16 and a persistent volume (`db_data`).  
- JWT is used for authentication of private endpoints.  
- For development, the API automatically tries to connect to the database and restarts on failure (`restart: always` in Docker).
