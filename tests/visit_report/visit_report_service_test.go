package visit_report_test

import (
	"testing"
	"time"
	"tourism-monitoring/entities"
	"tourism-monitoring/services/visit_report"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockVisitReportRepo for testing
type MockVisitReportRepo struct {
	mock.Mock
}

func (m *MockVisitReportRepo) GetAllVisitReports(limit, offset int) ([]entities.VisitReport, error) {
	args := m.Called(limit, offset)
	return args.Get(0).([]entities.VisitReport), args.Error(1)
}

func (m *MockVisitReportRepo) GetVisitReportByID(id int) (entities.VisitReport, error) {
	args := m.Called(id)
	return args.Get(0).(entities.VisitReport), args.Error(1)
}

func (m *MockVisitReportRepo) InsertVisitReport(report entities.VisitReport) (entities.VisitReport, error) {
	args := m.Called(report)
	return args.Get(0).(entities.VisitReport), args.Error(1)
}

func (m *MockVisitReportRepo) UpdateVisitReport(id int, report entities.VisitReport) (entities.VisitReport, error) {
	args := m.Called(id, report)
	return args.Get(0).(entities.VisitReport), args.Error(1)
}

func (m *MockVisitReportRepo) DeleteVisitReport(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockVisitReportRepo) GetTotalVisitReportsCount() (int64, error) {
	args := m.Called()
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockVisitReportRepo) GetAverageVisitsForPlace(placeID int) (float64, error) {
	args := m.Called(placeID)
	return args.Get(0).(float64), args.Error(1)
}

func TestVisitReportService_GetAllVisitReports(t *testing.T) {
	mockRepo := new(MockVisitReportRepo)

	// Simulate expected data
	expectedReports := []entities.VisitReport{
		{
			ID:               1,
			TanggalKunjungan: time.Date(2024, 11, 1, 0, 0, 0, 0, time.UTC),
			EstimasiEmisiKarbon: 10.5,
		},
		{
			ID:               2,
			TanggalKunjungan: time.Date(2024, 11, 2, 0, 0, 0, 0, time.UTC),
			EstimasiEmisiKarbon: 8.3,
		},
	}

	mockRepo.On("GetAllVisitReports", 10, 0).Return(expectedReports, nil)

	service := visit_report.NewVisitReportService(mockRepo)
	result, err := service.GetAllVisitReports(1, 10)

	assert.NoError(t, err)
	assert.Equal(t, expectedReports, result)
	mockRepo.AssertExpectations(t)
}

func TestVisitReportService_GetVisitReportByID(t *testing.T) {
	mockRepo := new(MockVisitReportRepo)
	expectedReport := entities.VisitReport{
		ID:               1,
		TanggalKunjungan: time.Date(2024, 11, 1, 0, 0, 0, 0, time.UTC),
		EstimasiEmisiKarbon: 10.5,
	}

	mockRepo.On("GetVisitReportByID", 1).Return(expectedReport, nil)

	service := visit_report.NewVisitReportService(mockRepo)
	result, err := service.GetVisitReportByID(1)

	assert.NoError(t, err)
	assert.Equal(t, expectedReport, result)
	mockRepo.AssertExpectations(t)
}

func TestVisitReportService_InsertVisitReport(t *testing.T) {
	mockRepo := new(MockVisitReportRepo)

	report := entities.VisitReport{
		TanggalKunjungan: time.Date(2024, 11, 3, 0, 0, 0, 0, time.UTC),
		EstimasiEmisiKarbon: 12.7,
	}

	expectedReport := report
	expectedReport.ID = 1

	mockRepo.On("InsertVisitReport", report).Return(expectedReport, nil)

	service := visit_report.NewVisitReportService(mockRepo)
	result, err := service.InsertVisitReport(report, "car", 100.0)

	assert.NoError(t, err)
	assert.Equal(t, expectedReport, result)
	mockRepo.AssertExpectations(t)
}

func TestVisitReportService_DeleteVisitReport(t *testing.T) {
	mockRepo := new(MockVisitReportRepo)
	mockRepo.On("DeleteVisitReport", 1).Return(nil)

	service := visit_report.NewVisitReportService(mockRepo)
	err := service.DeleteVisitReport(1)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
