# Product Management API

RESTful API for managing products and categories with Go, Gin, GORM, and PostgreSQL.

## 🚀 Quick Start

```bash
# Clone and setup
git clone https://github.com/sirawong/crud-arise.git
cd crud-arise
cp .env.example .env

# Start everything (API + Database + Sample Data)
make up
```

**API**: `http://localhost:8080`  
**Swagger**: `http://localhost:8080/swagger/index.html`

## 📖 API Endpoints

**Products**
- `GET /api/v1/products` - List products
- `POST /api/v1/products` - Create product
- `GET /api/v1/products/{id}` - Get product
- `PUT /api/v1/products/{id}` - Update product
- `DELETE /api/v1/products/{id}` - Delete product

**Categories**
- `GET /api/v1/categories` - List categories
- `POST /api/v1/categories` - Create category
- `GET /api/v1/categories/{id}` - Get category
- `PUT /api/v1/categories/{id}` - Update category
- `DELETE /api/v1/categories/{id}` - Delete category

## 🔧 Development

### Architecture
- **Clean Architecture**: Domain → Service → Handler layers
- **Dependency Injection**: Clean separation of concerns
- **Repository Pattern**: Abstract database operations
- **Unit Testing**: Comprehensive test coverage with mocks

### Project Structure
```
├── cmd/api/           # Application entry point
├── internal/
│   ├── domain/        # Business entities & interfaces
│   ├── services/      # Business logic
│   ├── handler/       # HTTP controllers
│   ├── repository/    # Database layer
│   └── di/           # Dependency injection
├── pkg/              # Shared utilities
├── scripts/          # Database initialization
└── docs/            # Swagger documentation
```

### Development Workflow
```bash
# 1. Start database + API + sample data
make up

# 2. For local development (API runs locally)
make down              # Stop Docker API
make postgres          # Start only PostgreSQL
make run              # Run API locally

# 3. Development commands
make test             # Run unit tests
make swagger          # Generate API docs
make gen              # Generate mocks
make lint             # Code linting
```

### Database
- **Auto Migration**: GORM handles schema changes
- **Sample Data**: Automatically loaded on first startup
- **Clean Separation**: Repository pattern abstracts DB operations

### Testing Strategy
- **Service Layer**: Business logic with mocked repositories
- **Handler Layer**: HTTP endpoints with mocked services
- **Test Coverage**: Run `go test -cover ./...`

### Adding New Features
1. Define entity in `internal/domain/entity/`
2. Create repository interface in `internal/domain/repository/`
3. Implement business logic in `internal/services/`
4. Add HTTP handlers in `internal/handler/http/`
5. Wire dependencies in `internal/di/`

## ⚙️ Configuration

Edit `.env` file:
```env
HTTP_SERVER_PORT=8080
DNS=postgresql://postgres:password@postgresql:5432/product_db?sslmode=disable
```

For local development, change `postgresql` to `localhost` in DNS.

## 📝 Usage Examples

**Create Category**
```bash
curl -X POST http://localhost:8080/api/v1/categories \
  -H "Content-Type: application/json" \
  -d '{"name": "Electronics"}'
```

**Create Product**
```bash
curl -X POST http://localhost:8080/api/v1/products \
  -H "Content-Type: application/json" \
  -d '{
    "name": "iPhone 15",
    "description": "Latest iPhone",
    "sku": "IP15-001", 
    "price": 999.99,
    "stock": 100,
    "categoryId": "category-id-here"
  }'
```

**List Products with Filters**
```bash
curl "http://localhost:8080/api/v1/products?name=iPhone&minPrice=500&maxPrice=1500"
```

## 🛠️ Tech Stack

- **Go 1.24.5** + Gin + GORM
- **PostgreSQL 15**
- **Docker & Docker Compose**
- **Swagger/OpenAPI**
