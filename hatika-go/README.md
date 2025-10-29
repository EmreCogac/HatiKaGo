# hatikago Go - ABP Framework Port

Go dilinde yazılmış, ABP Framework'ten ilham alan bir backend framework.

## Özellikler

- ✅ **Domain Driven Design (DDD)** mimarisi
- ✅ **Multi-tenancy** desteği
- ✅ **JWT Authentication & Authorization**
- ✅ **Audit Logging** (CreatedBy, CreatedAt, UpdatedBy, UpdatedAt)
- ✅ **Repository Pattern**
- ✅ **Clean Architecture**
- ✅ **RESTful API**
- ✅ **GORM ORM**
- ✅ **Validation**
- ✅ **Error Handling**

## Proje Yapısı

```
hatika-go/
├── cmd/
│   └── api/              # Main application entry point
├── internal/
│   ├── domain/           # Domain entities (Core layer)
│   │   ├── entities/     # Business entities
│   │   └── repositories/ # Repository interfaces
│   ├── application/      # Application services (Business logic)
│   │   ├── services/     # Application services
│   │   └── dtos/         # Data Transfer Objects
│   ├── infrastructure/   # Infrastructure layer
│   │   ├── persistence/  # Database implementations
│   │   └── config/       # Configuration
│   └── interfaces/       # Interface adapters
│       ├── http/         # HTTP handlers
│       └── middleware/   # HTTP middleware
├── pkg/                  # Shared packages
│   ├── auth/            # JWT, authorization utilities
│   ├── multitenancy/    # Multi-tenant support
│   ├── audit/           # Audit logging
│   └── errors/          # Error handling
└── config/              # Configuration files
```

## Kurulum

### Gereksinimler
- Go 1.21+
- PostgreSQL 13+

### Adımlar

1. Bağımlılıkları yükleyin:
```bash
go mod download
```

2. `.env` dosyası oluşturun:
```bash
cp config/.env.example config/.env
```

3. Veritabanı ayarlarını yapılandırın

4. Uygulamayı çalıştırın:
```bash
go run cmd/api/main.go
```

## API Endpoints

### Authentication
- `POST /api/auth/login` - Login
- `POST /api/auth/register` - Register

### Projects
- `GET /api/projects` - Get all projects (paginated)
- `GET /api/projects/:id` - Get project by ID
- `POST /api/projects` - Create project
- `PUT /api/projects/:id` - Update project
- `DELETE /api/projects/:id` - Delete project

### OCR Projects
- `GET /api/ocr-projects` - Get all OCR projects
- `GET /api/ocr-projects/:id` - Get OCR project by ID
- `PUT /api/ocr-projects/:id` - Update OCR project

## Kullanım Örnekleri

### Project Oluşturma
```bash
curl -X POST http://localhost:8080/api/projects \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -d '{
    "projectName": "Yeni Proje",
    "projectCode": "PRJ-001",
    "projectComment": "Test projesi",
    "yapiSahibi": "Ahmet Yılmaz",
    "adress": "İstanbul"
  }'
```

### Projeleri Listeleme
```bash
curl -X GET "http://localhost:8080/api/projects?pageNumber=1&pageSize=10" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

## Geliştirme

### Test Çalıştırma
```bash
go test ./...
```

### Build
```bash
go build -o bin/api cmd/api/main.go
```

### Docker ile Çalıştırma
```bash
docker-compose up -d
```
