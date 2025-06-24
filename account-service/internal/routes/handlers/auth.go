package handlers

import (
	"net/http"

	"github.com/chainpulse/backend/account/internal/models"
	"github.com/chainpulse/backend/account/internal/routes/handlers/dto"
	service "github.com/chainpulse/backend/account/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Handlers struct {
	AuthService service.IAuthSvc
}

func NewHandlers(authSvc service.IAuthSvc) *Handlers {
	return &Handlers{AuthService: authSvc}
}

func (h *Handlers) CreateBuyer(c *gin.Context) {
	var input dto.CreateBuyerRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		dto.WriteErrorResponse(c.Writer, http.StatusBadRequest, "Validation failed: "+err.Error())
		return
	}

	// Transform DTO to model
	buyer := &models.Buyer{
		FirstName:   &input.FirstName,
		LastName:    &input.LastName,
		Email:       &input.Email,
		PhoneNumber: input.Phone,
		Password:    input.Password, // hash later
	}

	if err := h.AuthService.CreateBuyer(c.Request.Context(), buyer); err != nil {
		dto.WriteErrorResponse(c.Writer, http.StatusInternalServerError, err.Error())
		return
	}
	dto.WriteSuccessResponse(c.Writer, buyer, "Buyer created successfully")
}

func (h *Handlers) LoginBuyer(c *gin.Context) {
	var loginReq dto.LoginRequest
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		dto.WriteErrorResponse(c.Writer, http.StatusBadRequest, "Invalid JSON")
		return
	}
	user, err := h.AuthService.LoginBuyer(c.Request.Context(), loginReq.Email, loginReq.Password)
	if err != nil {
		dto.WriteErrorResponse(c.Writer, http.StatusUnauthorized, "Invalid credentials")
		return
	}
	c.SetCookie("access_token", "mocked-access-token", 3600, "/", "", false, true)
	c.SetCookie("refresh_token", "mocked-refresh-token", 86400, "/", "", false, true)
	dto.WriteSuccessResponse(c.Writer, user, "Login successful")
}

func (h *Handlers) GetBuyerByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		dto.WriteErrorResponse(c.Writer, http.StatusBadRequest, "Invalid UUID")
		return
	}
	buyer, err := h.AuthService.GetBuyerByID(c.Request.Context(), id)
	if err != nil {
		dto.WriteErrorResponse(c.Writer, http.StatusNotFound, "Buyer not found")
		return
	}
	dto.WriteSuccessResponse(c.Writer, buyer, "Buyer found")
}

func (h *Handlers) GetBuyerByEmail(c *gin.Context) {
	email := c.Query("email")
	if email == "" {
		dto.WriteErrorResponse(c.Writer, http.StatusBadRequest, "Email is required")
		return
	}
	buyer, err := h.AuthService.GetBuyerByEmail(c.Request.Context(), email)
	if err != nil {
		dto.WriteErrorResponse(c.Writer, http.StatusNotFound, "Buyer not found")
		return
	}
	dto.WriteSuccessResponse(c.Writer, buyer, "Buyer found")
}

func (h *Handlers) UpdateBuyer(c *gin.Context) {
	var input models.Buyer
	if err := c.ShouldBindJSON(&input); err != nil {
		dto.WriteErrorResponse(c.Writer, http.StatusBadRequest, "Invalid JSON")
		return
	}
	err := h.AuthService.UpdateBuyer(c.Request.Context(), &input)
	if err != nil {
		dto.WriteErrorResponse(c.Writer, http.StatusInternalServerError, err.Error())
		return
	}
	dto.WriteSuccessResponse(c.Writer, input, "Buyer updated")
}

func (h *Handlers) DeleteBuyer(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		dto.WriteErrorResponse(c.Writer, http.StatusBadRequest, "Invalid UUID")
		return
	}
	err = h.AuthService.DeleteBuyer(c.Request.Context(), id)
	if err != nil {
		dto.WriteErrorResponse(c.Writer, http.StatusInternalServerError, err.Error())
		return
	}
	dto.WriteSuccessResponse(c.Writer, nil, "Buyer deleted")
}

// --- Settings Handlers ---

func (h *Handlers) GetSettings(c *gin.Context) {
	settings, err := h.AuthService.GetSettings(c.Request.Context())
	if err != nil {
		dto.WriteErrorResponse(c.Writer, http.StatusInternalServerError, err.Error())
		return
	}
	dto.WriteSuccessResponse(c.Writer, settings, "Settings retrieved")
}

func (h *Handlers) UpdateSettings(c *gin.Context) {
	var input models.Settings
	if err := c.ShouldBindJSON(&input); err != nil {
		dto.WriteErrorResponse(c.Writer, http.StatusBadRequest, "Invalid JSON")
		return
	}
	err := h.AuthService.UpdateSettings(c.Request.Context(), &input)
	if err != nil {
		dto.WriteErrorResponse(c.Writer, http.StatusInternalServerError, err.Error())
		return
	}
	dto.WriteSuccessResponse(c.Writer, input, "Settings updated")
}

func (h *Handlers) InsertSettings(c *gin.Context) {
	var input models.Settings
	if err := c.ShouldBindJSON(&input); err != nil {
		dto.WriteErrorResponse(c.Writer, http.StatusBadRequest, "Invalid JSON")
		return
	}
	err := h.AuthService.InsertSettings(c.Request.Context(), &input)
	if err != nil {
		dto.WriteErrorResponse(c.Writer, http.StatusInternalServerError, err.Error())
		return
	}
	dto.WriteSuccessResponse(c.Writer, input, "Settings inserted")
}
