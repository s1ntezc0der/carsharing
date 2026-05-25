# Car Rental App

Premium car sharing service offering exclusive vehicles from Lamborghini to Ferrari. 
Book your dream car in 3 simple steps.

## System Requirements
- Docker 20.10+
- Docker Compose 2.0+
- (Optional for local development):
  - Go 1.24+ (for backend development)
  - Node.js 20+ and Angular CLI (for frontend development)

## Windows Setup Guide (if not on Linux)

1. **Install Prerequisites**:
   - Install [Docker Desktop](https://www.docker.com/products/docker-desktop/)
   - Enable WSL2 backend (recommended) during installation

2. **First-Time Setup**:

```bash
# Clone the repository
git clone https://github.com/s1ntezc0der/carsharing.git
cd carsharing-go
``` 

## Short Project Structure

- `backend/`: Go server with API endpoints
- `frontend/`: Angular frontend application
- `docker-compose.yaml`: Docker configuration

## Detailed Project Structure Tree

```bash
carsharing-go/
├── backend/
│   ├── cmd/
│   │   └── main.go
│   ├── internal/
│   │   ├── config/
│   │   │   └── config.go
│   │   ├── controller/
│   │   │   ├── handlers.go
│   │   │   └── routers.go
│   │   ├── entity/
│   │   │   └── models.go
│   │   ├── repository/
│   │   │   └── repo.go
│   │   ├── usecase/
│   │   │   └── services.go
│   │   └── run/
│   │       └── run.go
│   ├── images/
│   │   └── cars/
│   │       ├── 1.png
│   │       ├── 2.png
│   │       └── ...
│   ├── .env
│   ├── Dockerfile
│   └── wait-for-it.sh
│
├── frontend/
│   ├── src/
│   │   ├── app/
│   │   │   └── cars/
│   │   │       ├── cars.component.css
│   │   │       ├── cars.component.html
│   │   │       └── cars.component.ts
│   │   ├── global_styles.css
│   │   └── index.html
│   ├── nginx.conf
│   ├── Dockerfile
│   ├── package.json
│   └── proxy.conf.json
│
├── docker-compose.yaml
├── Makefile
├── make.bat
├── go.mod
├── go.sum
└── README.md
```

## Setup

1. Build and run with Docker:
```bash
# full project
docker-compose up --build

# Or try makefiles:
make cars

# CLEANUP
docker-compose down -v
docker builder prune -af
docker system prune -af
```

2. `make cars` will run both backend and frontend:

- Backend can be accessed via: [http://localhost:8080/api/cars]

- Frontend can be checked here: [http://localhost:3000]

3. cURL:

```curl
curl http://localhost:8080/api/cars
curl http://localhost:8080/api/cars?filter=Lamborghini
```

4. Database connection:

```bash
# Check active containers (779793098e50):
docker ps

# Enter postgres:
docker exec -it 779793098e50 psql -U postgres -d carsharing

# List all tables:
\dt
```

5. SQL:

```sql
# ASCEND ^, DESCEND v
SELECT * FROM orders ORDER BY created_at DESC LIMIT 5;
```

## Ports

- `Frontend`: carsharing-go-frontend-1 (port 3000)
- `Backend`: carsharing-go-backend-1 (port 8080)
- `Database`: carsharing-go-db-1 (port 5432)

## API Endpoints (Routes For Frontend Developers)

**Base URL:** `http://localhost:8080/api`

| Method  | Endpoint             | Parameters           | Headers                          | Description                          |
|---------|----------------------|----------------------|----------------------------------|--------------------------------------|
| `GET`   | `/cars`              | `?filter={brand}`    | -                                | Fetch all cars (filterable by brand) |
| `POST`  | `/orders`            | JSON payload         | `Content-Type: application/json` | Create booking                       |
| `GET`   | `/images/{filename}` | -                    | -                                | Get car images                       |

### GET /api/cars

**Test with curl:**

```bash
# GET Ferrari:
curl "http://localhost:8080/api/cars?filter=Ferrari"

# Response:
[
  {
    "id": "4",
    "image": "/images/cars/4.png",
    "title": "Ferrari F8 Spider",
    "text": "Мечта на колесах. Ferrari F8 Spider: 720 л.с., аэродинамика F1 и открытая кабина для тех, кто живет на полной скорости.",
    "prices": [3500, 3200, 2900],
    "brand": "Ferrari"
  }
]
```

### POST /api/orders

**Test with curl:**
```bash
# POST - make an order:
curl -X POST http://localhost:8080/api/orders \
  -H "Content-Type: application/json" \
  -d '{
    "car": "Ferrari 488 GTB",
    "name": "Alex",
    "phone": "+971501234567"
  }'

# Response: 
{
  "message": "Заказ успешно создан! Мы свяжемся с вами в ближайшее время."
}
```

## Tests for Golang:

```bash
# Run tests with coverage and save to file
go test ./... -coverprofile=coverage.out

# View coverage report in terminal
go tool cover -func=coverage.out

# For a quick summary:
go test ./... -cover
```

## DockerHub updates:

Run these commands from root directory where docker-compose is:

```bash
# Login to Docker Hub
docker login

# Build and push backend
docker build -t s1ntezc0der/carsharing-backend -f backend/Dockerfile .
docker push s1ntezc0der/carsharing-backend

# Build and push frontend
docker build -t s1ntezc0der/carsharing-frontend -f frontend/Dockerfile ./frontend
docker push s1ntezc0der/carsharing-frontend
```

Enjoy, Go (Chi) + TS (Angular) := fullstack!!! 
