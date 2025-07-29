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

```bash
make help          # Show all commands
make up            # Start all services + data
make down          # Stop services
make run           # Run API locally (needs PostgreSQL)
make test          # Run tests
make build         # Build and start
```

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
