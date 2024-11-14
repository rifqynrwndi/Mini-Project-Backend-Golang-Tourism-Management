package ai

import (
	"context"
	"fmt"
	"os"
	"strings"
	"tourism-monitoring/repositories/visit_report"
	"tourism-monitoring/repositories/places"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type AIService struct {
	client          *genai.Client
	visitReportRepo *visit_report.VisitReportRepo
	placesRepo      *places.PlacesRepo
}

func NewAIService(visitReportRepo *visit_report.VisitReportRepo, placesRepo *places.PlacesRepo) (*AIService, error) {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
	if err != nil {
		return nil, fmt.Errorf("failed to create Gemini client: %w", err)
	}
	return &AIService{
		client:          client,
		visitReportRepo: visitReportRepo,
		placesRepo:      placesRepo,
	}, nil
}


func (service *AIService) GetVisitDataForPrediction(placeID int) (int, float64, error) {
	place, err := service.placesRepo.GetPlaceByID(placeID)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to retrieve place: %w", err)
	}

	avgVisits, err := service.visitReportRepo.GetAverageVisitsForPlace(placeID)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to retrieve average visits: %w", err)
	}

	return place.JumlahPengunjung, avgVisits, nil
}


func (service *AIService) PredictVisitsAndRecommend(placeID int) (result string, err error) {
	jumlahPengunjung, rataRataKunjungan, err := service.GetVisitDataForPrediction(placeID)
	if err != nil {
		return "", fmt.Errorf("failed to get visit data for prediction: %w", err)
	}

	prompt := fmt.Sprintf(
		"Berdasarkan data kunjungan historis, jumlah pengunjung saat ini di lokasi wisata adalah %d dengan rata-rata kunjungan per bulan sekitar %.2f. Prediksi jumlah kunjungan bulan depan dan berikan rekomendasi untuk manajemen kapasitas dan kegiatan ramah lingkungan bagi wisatawan.",
		jumlahPengunjung,
		rataRataKunjungan,
	)

	ctx := context.Background()
	model := service.client.GenerativeModel("gemini-1.5-flash")

	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		return "", fmt.Errorf("failed to generate content: %w", err)
	}

	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				result += fmt.Sprintf("%s", part)
			}
		}
	}

	result = strings.ReplaceAll(result, "\n", "")

	return result, nil
}