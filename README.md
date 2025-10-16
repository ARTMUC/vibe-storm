# VibeStorm

A modern web application built with Go Echo, Templ templates, TypeScript, and MySQL, following Domain-Driven Design (DDD) principles and clean architecture.

## 🚀 Features

- **Backend**: Go with Echo framework
- **Frontend**: Templ templates with TypeScript support
- **Database**: MySQL with GORM ORM
- **Architecture**: Domain-Driven Design (DDD) with clean architecture
- **Development**: Hot reload, TypeScript compilation, and modern tooling
- **Deployment**: Docker support and production-ready configuration

## 📁 Project Structure

```
vibe-storm/
├── cmd/                    # Application entry points
│   └── server/            # Main server application
├── internal/              # Private application code
│   ├── domain/           # Domain layer (entities, value objects, domain services)
│   ├── application/      # Application layer (use cases, DTOs, application services)
│   └── interfaces/       # Interface adapters (controllers, presenters, middleware)
├── pkg/                  # Public packages
│   ├── common/          # Common utilities
│   ├── config/          # Configuration management
│   ├── database/        # Database connection and utilities
│   └── middleware/      # HTTP middleware
├── web/                 # Frontend assets
│   ├── templates/       # Templ templates
│   ├── static/         # Static assets (CSS, JS, images)
│   └── package.json    # Frontend dependencies and scripts
├── migrations/         # Database migrations
├── configs/           # Configuration files
├── api/              # API documentation
├── .env.example      # Environment variables template
├── go.mod           # Go module definition
├── Makefile        # Build automation
└── README.md       # Project documentation
```

## 🛠️ Technology Stack

### Backend
- **Go 1.25** - Programming language
- **Echo** - Web framework
- **GORM** - ORM for database operations
- **Viper** - Configuration management
- **Logrus** - Logging
- **Google Wire** - Dependency injection

### Frontend
- **Templ** - Type-safe templates for Go
- **TypeScript** - Typed JavaScript
- **Webpack** - Module bundler
- **CSS3** - Styling

### Database
- **MySQL** - Relational database
- **Golang-Migrate** - Database migrations

### Development Tools
- **Docker** - Containerization
- **Make** - Build automation
- **ESLint** - JavaScript linting
- **Prettier** - Code formatting

## 🚀 Quick Start

### Prerequisites

- Go 1.25 or later
- Node.js 16+ and npm
- Docker and Docker Compose (for local development)
- Make (optional, for using Makefile commands)

### Development Setup

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd vibe-storm
   ```

2. **Set up environment**
   ```bash
   make dev-setup
   # or manually:
   # cp .env.example .env
   # go mod tidy
   # cd web && npm install
   ```

3. **Local Database Setup**
   
   You can use either a local MySQL installation or Docker for the database.

   **Option 1: Using Docker (Recommended for local development)**
   ```bash
   docker-compose up -d
   ```
   
   Update your `.env` file with:
   ```env
   DB_HOST=localhost
   DB_PORT=3306
   DB_USERNAME=vibe_user
   DB_PASSWORD=vibe_password
   DB_DATABASE=vibe_storm
   ```

   **Option 2: Using local MySQL installation**
   - Install MySQL 8.0+ on your system
   - Create a database and user
   - Update `.env` file with your MySQL credentials

4. **Run database migrations**
   ```bash
   make migrate-up
   ```

5. **Build and run**
   ```bash
   make build
   make run
   ```

   The application will be available at:
   - API: http://localhost:8080/api/v1
   - Frontend: http://localhost:8080

## 📝 Available Commands

### Development
```bash
make help          # Show all available commands
make dev-setup     # Set up development environment
make build         # Build the Go application
make run           # Run the application
make test          # Run tests
make clean         # Clean build artifacts
```

### Database
```bash
make migrate-up    # Run database migrations
make migrate-down  # Rollback migrations
make migrate-create # Create new migration
```

### Frontend
```bash
make build-frontend # Build frontend assets
cd web && npm run dev  # Watch mode for frontend
```

### Docker
```bash
make docker-build  # Build Docker image
make docker-run    # Run in Docker container
```

## 🔧 Configuration

### Environment Variables

Copy `.env.example` to `.env` and update the values:

```env
# Server Configuration
SERVER_PORT=8080
SERVER_READ_TIMEOUT=10s
SERVER_WRITE_TIMEOUT=10s

# Database Configuration
DB_HOST=localhost
DB_PORT=3306
DB_USERNAME=root
DB_PASSWORD=password
DB_DATABASE=vibe_storm

# Application Configuration
APP_NAME=VibeStorm
APP_VERSION=1.0.0
APP_ENV=development
```

### Configuration Files

- `configs/app.yaml` - Application configuration
- `web/package.json` - Frontend dependencies
- `web/tsconfig.json` - TypeScript configuration
- `web/webpack.config.js` - Frontend build configuration

## 🏗️ Architecture

### Domain-Driven Design (DDD)

The project follows DDD principles with clear separation of concerns:

- **Domain Layer** (`internal/domain/`): Contains business entities, value objects, and domain services
- **Application Layer** (`internal/application/`): Contains use cases and application services
- **Infrastructure Layer** (`internal/infrastructure/`): Contains database implementations and external services
- **Interface Layer** (`internal/interfaces/`): Contains controllers, presenters, and HTTP handlers

### Clean Architecture

The architecture follows clean architecture principles:
- **Dependency Rule**: Dependencies point inward toward the domain
- **Independent Layers**: Each layer has its own responsibilities
- **Testability**: Easy to test each layer independently

## 🔒 Security

- CORS middleware configured
- Secure headers middleware
- Request logging and monitoring
- Input validation and sanitization
- SQL injection prevention with GORM

## 📊 Monitoring

- Structured logging with Logrus
- Request/response logging middleware
- Health check endpoint at `/api/v1/health`
- Database connection monitoring

## 🚢 Deployment

### Docker Deployment

1. Build the application:
   ```bash
   make prod-build
   ```

2. Build Docker image:
   ```bash
   make docker-build
   ```

3. Run with Docker:
   ```bash
   make docker-run
   ```

### Manual Deployment

1. Set `APP_ENV=production` in your environment
2. Build the application: `make prod-build`
3. Run the binary: `./bin/server`

## 🧪 Testing

```bash
make test              # Run all tests
make test-coverage     # Run tests with coverage report
```

## 📚 API Documentation

### Swagger/OpenAPI Documentation

The API is fully documented using Swagger/OpenAPI 3.0:

- **Swagger UI**: http://localhost:8080/swagger/index.html
- **OpenAPI JSON**: http://localhost:8080/swagger/doc.json
- **ReDoc**: http://localhost:8080/swagger/swagger-ui/

### Available Endpoints

#### Health Check
- `GET /api/v1/health` - Application health status

#### User Management
- `GET /api/v1/users` - Get all users (with pagination)
- `POST /api/v1/users` - Create a new user
- `GET /api/v1/users/{id}` - Get user by ID
- `PUT /api/v1/users/{id}` - Update user by ID
- `DELETE /api/v1/users/{id}` - Delete user by ID

### Generating Documentation

To regenerate Swagger documentation after making changes:

```bash
make swagger-gen
```

### API Features

- **Interactive Documentation**: Test endpoints directly from Swagger UI
- **Request/Response Examples**: See sample requests and responses
- **Authentication**: Bearer token and cookie authentication support
- **Validation**: Input validation with detailed error messages
- **Error Handling**: Comprehensive error response documentation

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch: `git checkout -b feature/amazing-feature`
3. Commit your changes: `git commit -m 'Add amazing feature'`
4. Push to the branch: `git push origin feature/amazing-feature`
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the LICENSE file for details.

## 🙏 Acknowledgments

- [Echo Framework](https://echo.labstack.com/)
- [GORM](https://gorm.io/)
- [Templ](https://templ.guide/)
- [Domain-Driven Design](https://domainlanguage.com/ddd/)
- [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)

## 📞 Support

For support, please contact the development team or create an issue in the repository.

---

**Happy coding! 🎉**
