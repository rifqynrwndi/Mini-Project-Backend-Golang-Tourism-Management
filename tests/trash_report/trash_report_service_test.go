package trash_report_test

import (
	"errors"
	"testing"
	"time"
	"tourism-monitoring/entities"
	"tourism-monitoring/services/trash_report"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockTrashReportRepo for testing
type MockTrashReportRepo struct {
	mock.Mock
}

func (m *MockTrashReportRepo) GetTrashReportByPlaceID(id int) ([]entities.TrashReport, error) {
	args := m.Called(id)
	return args.Get(0).([]entities.TrashReport), args.Error(1)
}

func (m *MockTrashReportRepo) GetTrashReportByID(id int) (entities.TrashReport, error) {
	args := m.Called(id)
	return args.Get(0).(entities.TrashReport), args.Error(1)
}

func (m *MockTrashReportRepo) InsertTrashReport(report entities.TrashReport) (entities.TrashReport, error) {
	args := m.Called(report)
	return args.Get(0).(entities.TrashReport), args.Error(1)
}

func (m *MockTrashReportRepo) UpdateTrashReport(id int, report entities.TrashReport) (entities.TrashReport, error) {
	args := m.Called(id, report)
	return args.Get(0).(entities.TrashReport), args.Error(1)
}

func (m *MockTrashReportRepo) DeleteTrashReport(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockTrashReportRepo) GetTotalTrashReportsCount() (int64, error) {
	args := m.Called()
	return args.Get(0).(int64), args.Error(1)
}

func TestTrashReportService_GetTrashReportByPlaceID(t *testing.T) {
	mockRepo := new(MockTrashReportRepo)
	expectedReports := []entities.TrashReport{
		{ID: 1, ObjekWisataID: 1, JumlahSampah: 10.5, TanggalLaporan: time.Date(2024, 11, 1, 0, 0, 0, 0, time.UTC)},
		{ID: 2, ObjekWisataID: 1, JumlahSampah: 8.2, TanggalLaporan: time.Date(2024, 11, 2, 0, 0, 0, 0, time.UTC)},
	}

	mockRepo.On("GetTrashReportByPlaceID", 1).Return(expectedReports, nil)

	service := trash_report.NewTrashReportService(mockRepo)
	result, err := service.GetTrashReportByPlaceID(1)

	assert.NoError(t, err)
	assert.Equal(t, expectedReports, result)
	mockRepo.AssertExpectations(t)
}

func TestTrashReportService_GetTrashReportByID_Success(t *testing.T) {
	mockRepo := new(MockTrashReportRepo)
	expectedReport := entities.TrashReport{ID: 1, ObjekWisataID: 1, JumlahSampah: 10.5, TanggalLaporan: time.Date(2024, 11, 1, 0, 0, 0, 0, time.UTC)}

	mockRepo.On("GetTrashReportByID", 1).Return(expectedReport, nil)

	service := trash_report.NewTrashReportService(mockRepo)
	result, err := service.GetTrashReportByID(1)

	assert.NoError(t, err)
	assert.Equal(t, expectedReport, result)
	mockRepo.AssertExpectations(t)
}

func TestTrashReportService_GetTrashReportByID_NotFound(t *testing.T) {
	mockRepo := new(MockTrashReportRepo)

	mockRepo.On("GetTrashReportByID", 99).Return(entities.TrashReport{}, errors.New("record not found"))

	service := trash_report.NewTrashReportService(mockRepo)
	result, err := service.GetTrashReportByID(99)

	assert.Error(t, err)
	assert.Equal(t, entities.TrashReport{}, result)
	mockRepo.AssertExpectations(t)
}

func TestTrashReportService_InsertTrashReport(t *testing.T) {
	mockRepo := new(MockTrashReportRepo)
	newReport := entities.TrashReport{ObjekWisataID: 1, JumlahSampah: 12.5, TanggalLaporan: time.Date(2024, 11, 3, 0, 0, 0, 0, time.UTC)}
	expectedReport := entities.TrashReport{ID: 1, ObjekWisataID: 1, JumlahSampah: 12.5, TanggalLaporan: time.Date(2024, 11, 3, 0, 0, 0, 0, time.UTC)}

	mockRepo.On("InsertTrashReport", newReport).Return(expectedReport, nil)

	service := trash_report.NewTrashReportService(mockRepo)
	result, err := service.InsertTrashReport(newReport)

	assert.NoError(t, err)
	assert.Equal(t, expectedReport, result)
	mockRepo.AssertExpectations(t)
}

func TestTrashReportService_UpdateTrashReport(t *testing.T) {
	mockRepo := new(MockTrashReportRepo)
	updateReport := entities.TrashReport{ObjekWisataID: 1, JumlahSampah: 15.0, TanggalLaporan: time.Date(2024, 11, 4, 0, 0, 0, 0, time.UTC)}
	expectedReport := entities.TrashReport{ID: 1, ObjekWisataID: 1, JumlahSampah: 15.0, TanggalLaporan: time.Date(2024, 11, 4, 0, 0, 0, 0, time.UTC)}

	mockRepo.On("UpdateTrashReport", 1, updateReport).Return(expectedReport, nil)

	service := trash_report.NewTrashReportService(mockRepo)
	result, err := service.UpdateTrashReport(1, updateReport)

	assert.NoError(t, err)
	assert.Equal(t, expectedReport, result)
	mockRepo.AssertExpectations(t)
}

func TestTrashReportService_DeleteTrashReport(t *testing.T) {
	mockRepo := new(MockTrashReportRepo)

	mockRepo.On("DeleteTrashReport", 1).Return(nil)

	service := trash_report.NewTrashReportService(mockRepo)
	err := service.DeleteTrashReport(1)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestTrashReportService_GetTotalTrashReportsCount(t *testing.T) {
	mockRepo := new(MockTrashReportRepo)

	mockRepo.On("GetTotalTrashReportsCount").Return(int64(10), nil)

	service := trash_report.NewTrashReportService(mockRepo)
	count, err := service.GetTotalTrashReportsCount()

	assert.NoError(t, err)
	assert.Equal(t, int64(10), count)
	mockRepo.AssertExpectations(t)
}
