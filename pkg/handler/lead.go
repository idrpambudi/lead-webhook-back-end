package handler

import (
	"leadwebhook/pkg/domain"
	"leadwebhook/pkg/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type LeadHandler struct {
	ls *service.LeadService
}

// RegisterRoutes for all http endpoints
func RegisterLeadRoutes(e *echo.Echo, service *service.LeadService) {
	handler := &LeadHandler{service}
	e.GET("/lead", handler.listLeads)
	e.GET("/user/:userID/webhook", handler.getWebhookID)
	e.POST("/webhook/:webhookID/lead", handler.addLead)
}

func (h *LeadHandler) listLeads(c echo.Context) error {
	userID := c.QueryParam("userID")
	offset, err := strconv.Atoi(c.QueryParam("offset"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	resp, err := h.ls.ListLeads(userID, offset, limit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}
	return c.JSON(http.StatusOK, resp)
}

func (h *LeadHandler) addLead(c echo.Context) error {
	webhookID := c.Param("webhookID")
	var lead domain.Lead
	if err := c.Bind(&lead); err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}
	lead.WebhookID = webhookID

	resp, err := h.ls.AddLead(lead)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}
	return c.JSON(http.StatusCreated, resp)
}

func (h *LeadHandler) getWebhookID(c echo.Context) error {
	userID := c.Param("userID")
	resp := h.ls.GetWebhookID(userID)
	return c.JSON(http.StatusOK, resp)
}
