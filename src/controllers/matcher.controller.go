package controllers

import (
	"matech-backend/src/models"
	"matech-backend/src/services"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type MatcherController interface {
	GetMatchSections(c *gin.Context)
}

type matcherController struct {
	matcherService services.MatcherService
}

func NewMatcherController(matcherService services.MatcherService) MatcherController {
	return &matcherController{matcherService: matcherService}
}

func (m *matcherController) GetMatchSections(c *gin.Context) {
	ids := strings.Split(c.Query("ids"), ",")
	result := []models.StudentSections{}
	if err := m.matcherService.GetStudentsSecitons(ids, &result); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	section := map[string]models.Section{}
	m.matcherService.MatchSections(result, &section)
	c.JSON(http.StatusOK, section)
}
