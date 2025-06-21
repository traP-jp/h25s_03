package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/eraxyso/go-template/api"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

func (h *Handler) PostEvent(ctx echo.Context) error {
	requestBody := api.PostEventJSONRequestBody{}
	if err := ctx.Bind(&requestBody); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	eventID, err := h.EventService.PostEvents(ctx, requestBody)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return ctx.JSON(http.StatusOK, eventID)
}

func (h *Handler) GetEvents(ctx echo.Context, params api.GetEventsParams) error {
	ifDeleted := params.IfDeleted
	userID := h.MiddlewareService.GetUserID(ctx)
	eventSummaries, err := h.EventService.GetEvents(ctx, ifDeleted, userID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return ctx.JSON(http.StatusOK, eventSummaries)
}

func (h *Handler) GetEvent(ctx echo.Context, eventID openapi_types.UUID) error {
	userID := h.MiddlewareService.GetUserID(ctx)
	event, err := h.EventService.GetEvent(ctx, eventID, userID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return ctx.JSON(http.StatusOK, event)
}

func (h *Handler) PatchEvent(ctx echo.Context, eventID openapi_types.UUID) error {
	requestBody := api.PatchEventJSONRequestBody{}
	if err := ctx.Bind(&requestBody); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	if err := h.EventService.PatchEvent(ctx, eventID, requestBody); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return ctx.NoContent(http.StatusOK)
}

func (h *Handler) DeleteEvent(ctx echo.Context, eventID openapi_types.UUID) error {
	if err := h.EventService.DeleteEvent(ctx, eventID); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return ctx.NoContent(http.StatusOK)
}
