package auth_test

import (
	"testing"
	"tourism-monitoring/entities"
	"tourism-monitoring/services"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockAuthRepo implements the AuthRepoInterface
type MockAuthRepo struct {
	mock.Mock
}

// Login implements auth.AuthRepoInterface.
func (m *MockAuthRepo) Login(user entities.User) (entities.User, error) {
	panic("unimplemented")
}

func (m *MockAuthRepo) GetUserByEmail(email string) (entities.User, error) {
	args := m.Called(email)
	return args.Get(0).(entities.User), args.Error(1)
}

func (m *MockAuthRepo) Register(user entities.User) (entities.User, error) {
	args := m.Called(user)
	return args.Get(0).(entities.User), args.Error(1)
}

func (m *MockAuthRepo) GetLastUserID() (int, error) {
	args := m.Called()
	return args.Int(0), args.Error(1)
}

// MockJwt implements the JwtInterface
type MockJwt struct {
	mock.Mock
}

func (m *MockJwt) GenerateJWT(userID int, name, role string) (string, error) {
	args := m.Called(userID, name, role)
	return args.String(0), args.Error(1)
}

// Test for Login function
func TestAuthService_Login(t *testing.T) {
    mockRepo := new(MockAuthRepo)
    mockJwt := new(MockJwt)
    authService := services.NewAuthService(mockRepo, mockJwt)

    email := "test@example.com"
    password := "correct_password"
    hashedPassword, _ := services.HashPassword(password)

    mockRepo.On("GetUserByEmail", email).Return(entities.User{
        ID:       1,
        Email:    email,
        Password: hashedPassword,
    }, nil)

    mockJwt.On("GenerateJWT", mock.Anything, mock.Anything, mock.Anything).Return("mockToken", nil)

    user, err := authService.Login(entities.User{
        Email:    email,
        Password: password,
    })

    assert.NoError(t, err)
    assert.Equal(t, "mockToken", user.Token)
    mockRepo.AssertExpectations(t)
    mockJwt.AssertExpectations(t)
}


// Test for Register function
func TestAuthService_Register(t *testing.T) {
	mockRepo := new(MockAuthRepo)
	mockJwt := new(MockJwt)
	authService := services.NewAuthService(mockRepo, mockJwt)

	newUser := entities.User{
		Nama:     "Test User",
		Email:    "testuser@example.com",
		Password: "password123",
	}

	hashedPassword := "$2a$14$PpltLlH2QCGQuY1byM4uwu9k5P34Kbfgh3CJS2ThdueC2oyRAsnde" // hashed password
	createdUser := newUser
	createdUser.ID = 1
	createdUser.Password = hashedPassword

	mockRepo.On("GetLastUserID").Return(0, nil)
	mockRepo.On("Register", mock.Anything).Return(createdUser, nil)
	mockJwt.On("GenerateJWT", mock.Anything, mock.Anything, mock.Anything).Return("mockToken", nil)

	result, err := authService.Register(newUser)

	assert.Nil(t, err)
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "mockToken", result.Token)
	mockRepo.AssertExpectations(t)
	mockJwt.AssertExpectations(t)
}
