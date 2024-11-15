package tourists_test

import (
	"errors"
	"testing"
	"tourism-monitoring/entities"
	"tourism-monitoring/services/tourists"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockTouristsRepo for testing
type MockTouristsRepo struct {
	mock.Mock
}

func (m *MockTouristsRepo) GetAllTourists() ([]entities.User, error) {
	args := m.Called()
	return args.Get(0).([]entities.User), args.Error(1)
}

func (m *MockTouristsRepo) GetTouristByID(id int) (entities.User, error) {
	args := m.Called(id)
	return args.Get(0).(entities.User), args.Error(1)
}

func (m *MockTouristsRepo) InsertTourist(user entities.User) (entities.User, error) {
	args := m.Called(user)
	return args.Get(0).(entities.User), args.Error(1)
}

func (m *MockTouristsRepo) UpdateTourist(id int, user entities.User) (entities.User, error) {
	args := m.Called(id, user)
	return args.Get(0).(entities.User), args.Error(1)
}

func (m *MockTouristsRepo) DeleteTourist(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockTouristsRepo) GetTotalTouristsCount() (int64, error) {
	args := m.Called()
	return args.Get(0).(int64), args.Error(1)
}

func TestTouristsService_GetAllTourists(t *testing.T) {
	mockRepo := new(MockTouristsRepo)
	expectedTourists := []entities.User{
		{ID: 1, Nama: "Tourist 1", Email: "tourist1@example.com"},
		{ID: 2, Nama: "Tourist 2", Email: "tourist2@example.com"},
	}

	mockRepo.On("GetAllTourists").Return(expectedTourists, nil)

	service := tourists.NewTouristsService(mockRepo)
	result, err := service.GetAllTourists()

	assert.NoError(t, err)
	assert.Equal(t, expectedTourists, result)
	mockRepo.AssertExpectations(t)
}

func TestTouristsService_GetTouristByID_Success(t *testing.T) {
	mockRepo := new(MockTouristsRepo)
	expectedTourist := entities.User{ID: 1, Nama: "Tourist 1", Email: "tourist1@example.com"}

	mockRepo.On("GetTouristByID", 1).Return(expectedTourist, nil)

	service := tourists.NewTouristsService(mockRepo)
	result, err := service.GetTouristByID(1)

	assert.NoError(t, err)
	assert.Equal(t, expectedTourist, result)
	mockRepo.AssertExpectations(t)
}

func TestTouristsService_GetTouristByID_NotFound(t *testing.T) {
	mockRepo := new(MockTouristsRepo)

	mockRepo.On("GetTouristByID", 99).Return(entities.User{}, errors.New("record not found"))

	service := tourists.NewTouristsService(mockRepo)
	result, err := service.GetTouristByID(99)

	assert.Error(t, err)
	assert.Equal(t, entities.User{}, result)
	mockRepo.AssertExpectations(t)
}

func TestTouristsService_InsertTourist(t *testing.T) {
	mockRepo := new(MockTouristsRepo)
	newTourist := entities.User{Nama: "New Tourist", Email: "newtourist@example.com"}
	expectedTourist := entities.User{ID: 1, Nama: "New Tourist", Email: "newtourist@example.com"}

	mockRepo.On("InsertTourist", newTourist).Return(expectedTourist, nil)

	service := tourists.NewTouristsService(mockRepo)
	result, err := service.InsertTourist(newTourist)

	assert.NoError(t, err)
	assert.Equal(t, expectedTourist, result)
	mockRepo.AssertExpectations(t)
}

func TestTouristsService_UpdateTourist(t *testing.T) {
	mockRepo := new(MockTouristsRepo)
	updateTourist := entities.User{Nama: "Updated Tourist", Email: "updatedtourist@example.com"}
	expectedTourist := entities.User{ID: 1, Nama: "Updated Tourist", Email: "updatedtourist@example.com"}

	mockRepo.On("UpdateTourist", 1, updateTourist).Return(expectedTourist, nil)

	service := tourists.NewTouristsService(mockRepo)
	result, err := service.UpdateTourist(1, updateTourist)

	assert.NoError(t, err)
	assert.Equal(t, expectedTourist, result)
	mockRepo.AssertExpectations(t)
}

func TestTouristsService_DeleteTourist(t *testing.T) {
	mockRepo := new(MockTouristsRepo)

	mockRepo.On("DeleteTourist", 1).Return(nil)

	service := tourists.NewTouristsService(mockRepo)
	err := service.DeleteTourist(1)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestTouristsService_GetTotalTouristsCount(t *testing.T) {
	mockRepo := new(MockTouristsRepo)

	mockRepo.On("GetTotalTouristsCount").Return(int64(5), nil)

	service := tourists.NewTouristsService(mockRepo)
	count, err := service.GetTotalTouristsCount()

	assert.NoError(t, err)
	assert.Equal(t, int64(5), count)
	mockRepo.AssertExpectations(t)
}
