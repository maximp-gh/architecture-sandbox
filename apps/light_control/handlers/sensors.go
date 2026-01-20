package handlers

import (
	//"context"
	"net/http"
	//"strconv"

	//"lightcontrol/models"
	"lightcontrol/services"

	"github.com/gin-gonic/gin"
)

// SensorHandler handles sensor-related requests
type SensorHandler struct {
	//DB                 *db.DB
	LightService *services.LightService
}

// NewSensorHandler creates a new SensorHandler
func NewSensorHandler(lightService *services.LightService) *SensorHandler {
	return &SensorHandler{
		//DB:                 db,
		LightService: lightService,
	}
}

// RegisterRoutes registers the sensor routes
func (h *SensorHandler) RegisterRoutes(router *gin.RouterGroup) {
	sensors := router.Group("/light")
	{
		sensors.GET("", h.GetLightSensors)
		sensors.GET("/:id", h.GetSensorByID)

		sensors.POST("", h.POSToperation)
		sensors.PUT("/:id", h.PUToperation)
		sensors.DELETE("/:id", h.DELoperation)
		sensors.PATCH("/:id/value", h.PATCHoperation)

		//sensors.GET("/:location", h.GetLightByLocation)
	}
}

func BasicMessage(s string, c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": s,
	})
}

// CRUD endpoints
func (h *SensorHandler) POSToperation(c *gin.Context) {
	BasicMessage("Use old endpoint for creating light sensors", c)
}
func (h *SensorHandler) PUToperation(c *gin.Context) {
	BasicMessage("Use old endpoint for replacing light sensors", c)
}
func (h *SensorHandler) PATCHoperation(c *gin.Context) {
	BasicMessage("Use old endpoint for updating light sensors", c)
}
func (h *SensorHandler) DELoperation(c *gin.Context) {
	BasicMessage("Use old endpoint for deleting light sensors", c)
}

func (h *SensorHandler) GetLightSensors(c *gin.Context) {
	BasicMessage("Enumerating all the light sensors is not implemented yet", c)
}

func (h *SensorHandler) GetSensorByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "sensor id is required"})
		return
	}

	// TODO: Here we need to validate if requested id belongs to the authorised used
	// like SendRequestTo UserProfile API.

	resp, err := h.LightService.CrudOps(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Return the light data
	c.JSON(http.StatusOK, gin.H{
		"id":       resp.ID,
		"type":     resp.Type,
		"location": resp.Location,
		"value":    0,
		//"unit":     "n/a",
		"status": "Not available",
		//"timestamp":   tempData.Timestamp,
		"description": "light state by ID",
	})
}

// GetByLocation handles GET /api/v1/light/:location
func (h *SensorHandler) GetLightByLocation(c *gin.Context) {
	location := c.Param("location")
	if location == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Location is required"})
		return
	}

	// Return the light data
	c.JSON(http.StatusOK, gin.H{
		"location": location,
		"value":    1,
		"unit":     "n/a",
		"status":   "on",
		//"timestamp":   tempData.Timestamp,
		"description": "light state",
	})
}
