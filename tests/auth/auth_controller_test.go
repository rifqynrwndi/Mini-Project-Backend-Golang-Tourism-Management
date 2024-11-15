package auth_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"tourism-monitoring/controllers/auth"
	"tourism-monitoring/entities"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockAuthService implements the AuthInterface
type MockAuthService struct {
	mock.Mock
}

func (m *MockAuthService) Login(user entities.User) (entities.User, error) {
	args := m.Called(user)
	return args.Get(0).(entities.User), args.Error(1)
}

func (m *MockAuthService) Register(user entities.User) (entities.User, error) {
	args := m.Called(user)
	return args.Get(0).(entities.User), args.Error(1)
}

func TestAuthController_Login(t *testing.T) {
	mockService := new(MockAuthService)
	controller := auth.NewAuthController(mockService)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(`{"email":"test@example.com", "password":"password123"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Mock behavior
	mockUser := entities.User{ID: 1, Email: "test@example.com", Token: "mockToken"}
	mockService.On("Login", mock.Anything).Return(mockUser, nil)

	// Call handler
	if assert.NoError(t, controller.LoginController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "mockToken")
	}

	mockService.AssertExpectations(t)
}

func TestAuthController_Register(t *testing.T) {
	mockService := new(MockAuthService)
	controller := auth.NewAuthController(mockService)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/register", strings.NewReader(`{"nama":"Test User", "email":"test@example.com", "password":"password123"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Mock behavior
	mockUser := entities.User{ID: 1, Email: "test@example.com", Token: "mockToken"}
	mockService.On("Register", mock.Anything).Return(mockUser, nil)

	// Call handler
	if assert.NoError(t, controller.RegisterController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "mockToken")
	}

	mockService.AssertExpectations(t)
}
