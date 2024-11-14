package ai

import (
	"net/http"
	"strconv"
	"tourism-monitoring/services/ai"

	"github.com/labstack/echo/v4"
)

type AIController struct {
	aiService *ai.AIService
}

func NewAIController(aiService *ai.AIService) *AIController {
	return &AIController{aiService: aiService}
}

func (controller *AIController) PredictVisitsAndRecommend(c echo.Context) error {
	placeID, err := strconv.Atoi(c.QueryParam("place_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid Place ID"})
	}

	response, err := controller.aiService.PredictVisitsAndRecommend(placeID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"prediksi_dan_rekomendasi": response})
}
