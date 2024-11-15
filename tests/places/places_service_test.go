package places_test

import (
	"errors"
	"testing"
	"tourism-monitoring/entities"
	"tourism-monitoring/services/places"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockPlacesRepo for testing
type MockPlacesRepo struct {
	mock.Mock
}

func (m *MockPlacesRepo) GetAllPlaces() ([]entities.Place, error) {
	args := m.Called()
	return args.Get(0).([]entities.Place), args.Error(1)
}

func (m *MockPlacesRepo) GetPlaceByID(id int) (entities.Place, error) {
	args := m.Called(id)
	return args.Get(0).(entities.Place), args.Error(1)
}

func (m *MockPlacesRepo) InsertPlace(place entities.Place) (entities.Place, error) {
	args := m.Called(place)
	return args.Get(0).(entities.Place), args.Error(1)
}

func (m *MockPlacesRepo) UpdatePlace(id int, place entities.Place) (entities.Place, error) {
	args := m.Called(id, place)
	return args.Get(0).(entities.Place), args.Error(1)
}

func (m *MockPlacesRepo) DeletePlace(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockPlacesRepo) GetTotalPlacesCount() (int64, error) {
	args := m.Called()
	return args.Get(0).(int64), args.Error(1)
}

func TestPlacesService_GetAllPlaces(t *testing.T) {
	mockRepo := new(MockPlacesRepo)
	expectedPlaces := []entities.Place{
		{ID: 1, Lokasi: "Place 1", KapasitasMaks: 50, JumlahPengunjung: 10},
		{ID: 2, Lokasi: "Place 2", KapasitasMaks: 100, JumlahPengunjung: 20},
	}

	mockRepo.On("GetAllPlaces").Return(expectedPlaces, nil)

	service := places.NewPlacesService(mockRepo)
	result, err := service.GetAllPlaces()

	assert.NoError(t, err)
	assert.Equal(t, expectedPlaces, result)
	mockRepo.AssertExpectations(t)
}

func TestPlacesService_GetPlaceByID_Success(t *testing.T) {
	mockRepo := new(MockPlacesRepo)
	expectedPlace := entities.Place{ID: 1, Lokasi: "Place 1", KapasitasMaks: 50, JumlahPengunjung: 10}

	mockRepo.On("GetPlaceByID", 1).Return(expectedPlace, nil)

	service := places.NewPlacesService(mockRepo)
	result, err := service.GetPlaceByID(1)

	assert.NoError(t, err)
	assert.Equal(t, expectedPlace, result)
	mockRepo.AssertExpectations(t)
}

func TestPlacesService_GetPlaceByID_NotFound(t *testing.T) {
	mockRepo := new(MockPlacesRepo)

	mockRepo.On("GetPlaceByID", 99).Return(entities.Place{}, errors.New("record not found"))

	service := places.NewPlacesService(mockRepo)
	result, err := service.GetPlaceByID(99)

	assert.Error(t, err)
	assert.Equal(t, entities.Place{}, result)
	mockRepo.AssertExpectations(t)
}

func TestPlacesService_InsertPlace(t *testing.T) {
	mockRepo := new(MockPlacesRepo)
	newPlace := entities.Place{Lokasi: "New Place", KapasitasMaks: 50, JumlahPengunjung: 10}
	expectedPlace := entities.Place{ID: 1, Lokasi: "New Place", KapasitasMaks: 50, JumlahPengunjung: 10}

	mockRepo.On("InsertPlace", newPlace).Return(expectedPlace, nil)

	service := places.NewPlacesService(mockRepo)
	result, err := service.InsertPlace(newPlace)

	assert.NoError(t, err)
	assert.Equal(t, expectedPlace, result)
	mockRepo.AssertExpectations(t)
}

func TestPlacesService_UpdatePlace(t *testing.T) {
	mockRepo := new(MockPlacesRepo)
	updatePlace := entities.Place{Lokasi: "Updated Place", KapasitasMaks: 100, JumlahPengunjung: 20}
	expectedPlace := entities.Place{ID: 1, Lokasi: "Updated Place", KapasitasMaks: 100, JumlahPengunjung: 20}

	mockRepo.On("UpdatePlace", 1, updatePlace).Return(expectedPlace, nil)

	service := places.NewPlacesService(mockRepo)
	result, err := service.UpdatePlace(1, updatePlace)

	assert.NoError(t, err)
	assert.Equal(t, expectedPlace, result)
	mockRepo.AssertExpectations(t)
}

func TestPlacesService_DeletePlace(t *testing.T) {
	mockRepo := new(MockPlacesRepo)

	mockRepo.On("DeletePlace", 1).Return(nil)

	service := places.NewPlacesService(mockRepo)
	err := service.DeletePlace(1)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestPlacesService_GetTotalPlacesCount(t *testing.T) {
	mockRepo := new(MockPlacesRepo)

	mockRepo.On("GetTotalPlacesCount").Return(int64(5), nil)

	service := places.NewPlacesService(mockRepo)
	count, err := service.GetTotalPlacesCount()

	assert.NoError(t, err)
	assert.Equal(t, int64(5), count)
	mockRepo.AssertExpectations(t)
}
