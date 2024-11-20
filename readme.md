# Tourism Monitoring System

A backend system to monitor sustainable tourism practices, including carbon emission tracking, waste management, and visitor management for tourist attractions. This project uses **clean architecture** principles and integrates external APIs for advanced functionalities like weather forecasts and visitor recommendations.

---

## 🚀 Features

### Core Features
- **Visit Report Management**:
  - Track tourist visits, calculate estimated carbon emissions, and log waste metrics.
- **Trash Report Management**:
  - Manage waste reports by type and quantity for each tourist attraction.
- **Tourist Management**:
  - Manage tourist data including type, origin, and demographic information.
- **Weather Integration**:
  - Fetch real-time weather data for tourist attractions.
- **AI Integration**:
  - Generate predictions for visitor numbers and provide recommendations for sustainable tourism practices.

### Additional Features
- **Sustainability Metrics**:
  - Calculate the environmental impact of visits and suggest greener alternatives.
- **Scalable Architecture**:
  - Designed with clean architecture for scalability and maintainability.

---

## 🛠️ Tech Stack

### Backend
- **Programming Language**: Go (Golang)
- **Frameworks**: 
  - Echo (HTTP routing)
  - GORM (Object Relational Mapper)
- **Database**: MySQL
- **Authentication**: JWT

### External Integrations
- **Weather API**: Fetches weather data for real-time monitoring.
- **Generative AI API**: Provides visitor predictions and sustainable recommendations.

### Deployment
- **Platform**: AWS EC2
- **CI/CD**: GitHub Actions for automated testing, building, and deployment.

---

## ⚙️ Project Structure

This project follows **clean architecture** principles.

```
/entities                # Core domain models
/repositories            # Data access logic
/services                # Business logic and use cases
/controllers             # HTTP handlers and request processing
/tests                   # Unit and integration tests
/config                  # Configuration files (e.g., environment variables)
/docs                    # API documentation (Swagger/Postman collections)
```

---

## 📋 API Documentation

### Postman Collection
You can import the Postman collection available in the `/docs` folder for testing the APIs.

### Swagger Documentation
Swagger documentation is automatically generated and can be accessed at:
```
[BASE_URL]/swagger/index.html
```

---

## 📦 Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/tourism-monitoring.git
   cd tourism-monitoring
   ```

2. Set up the environment variables in a `.env` file:
   ```env
   DB_HOST=localhost
   DB_USER=root
   DB_PASSWORD=yourpassword
   DB_NAME=tourism_db
   WEATHER_API_KEY=your-weather-api-key
   WEATHER_API_URL=https://api.weatherapi.com/v1
   GEMINI_API_KEY=your-ai-api-key
   ```

3. Run the application:
   ```bash
   go run main.go
   ```

---

## 🧪 Running Tests

Run the tests with coverage:
```bash
go test -v -cover ./tests/... -coverprofile=coverage.out
```

View the test coverage:
```bash
go tool cover -html=coverage.out
```

---

## 🚀 Deployment

### Docker
1. Build the Docker image:
   ```bash
   docker build -t tourism-monitoring .
   ```

2. Run the container:
   ```bash
   docker run -p 8080:8080 --env-file=.env tourism-monitoring
   ```

### CI/CD
The project uses GitHub Actions for continuous integration and deployment:
- **Pipeline Steps**:
  1. Lint and test code.
  2. Build Docker image.
  3. Deploy to AWS EC2 instance.

---

## 📖 ERD

Refer to the `/docs/ERD.png` for the Entity Relationship Diagram.

---

## 📂 Folder Structure

```
/entities            # Core domain models (e.g., VisitReport, Tourist)
/services            # Business logic (e.g., VisitReportService)
/repositories         # Data access layer (e.g., VisitReportRepo)
/controllers          # HTTP handlers (e.g., VisitReportController)
/tests                # Unit and integration tests
/config               # Configuration files (e.g., environment variables)
/docs                 # Documentation (e.g., Swagger, Postman)
```

---

## 📌 License

This project is licensed under the MIT License. See the `LICENSE` file for details.

---

## 💡 Contributors

- Muhammad Rifqy Nirwandi (Golang Mastery for Green Tech Backend Engineers)

Feel free to contribute by creating issues or submitting pull requests!