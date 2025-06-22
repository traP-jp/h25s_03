package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"traquji/api"
	"traquji/service"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

type PostEventJSONResponseBody struct {
	EventID openapi_types.UUID `json:"event_id"`
}

func (h *Handler) PostEvent(ctx echo.Context) error {
	requestBody := api.PostEventJSONRequestBody{}
	if err := ctx.Bind(&requestBody); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Errorf("bind request body (handler): %w", err))
	}
	eventOnCreate := service.EventOnCreate{
		Title:       requestBody.Title,
		Description: requestBody.Description,
		Date:        requestBody.Date.Time,
		IsOpen:      requestBody.IsOpen,
		Admins:      requestBody.Admins,
		Attendees:   requestBody.Attendees,
	}
	eventID, err := h.EventService.CreateEvent(ctx.Request().Context(), eventOnCreate)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Errorf("create event (handler): %w", err))
	}
	return ctx.JSON(http.StatusOK, PostEventJSONResponseBody{
		EventID: eventID,
	})
}

type GetEventsJSONResponseBody []api.EventSummary

func (h *Handler) GetEvents(ctx echo.Context, params api.GetEventsParams) error {
	ifDeleted := params.IfDeleted
	userID := h.MiddlewareService.GetUserID(ctx)
	eventSummaries, err := h.EventService.GetEvents(ctx.Request().Context(), ifDeleted, userID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Errorf("get events (handler): %w", err))
	}
	response := make([]api.EventSummary, 0, len(eventSummaries))
	for _, event := range eventSummaries {
		response = append(response, api.EventSummary{
			EventId:      event.EventID,
			Title:        event.Title,
			Description:  event.Description,
			Date:         openapi_types.Date{Time: event.Date},
			IsOpen:       event.IsOpen,
			IsMeAttendee: event.IsMeAttendee,
			Admins:       event.Admins,
		})
	}
	return ctx.JSON(http.StatusOK, GetEventsJSONResponseBody(response))
}

type GetEventJSONResponseBody api.Event

func (h *Handler) GetEvent(ctx echo.Context, eventID openapi_types.UUID) error {
	userID := h.MiddlewareService.GetUserID(ctx)
	event, err := h.EventService.GetEvent(ctx.Request().Context(), eventID, userID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Errorf("get event (handler): %w", err))
	}
	return ctx.JSON(http.StatusOK, GetEventJSONResponseBody{
		EventId:      event.EventID,
		Title:        event.Title,
		Description:  event.Description,
		Date:         openapi_types.Date{Time: event.Date},
		IsOpen:       event.IsOpen,
		IsMeAttendee: event.IsMeAttendee,
		IsDeleted:    event.IsDeleted,
		CreatedAt:    event.CreatedAt,
		UpdatedAt:    event.UpdatedAt,
		Admins:       event.Admins,
		Attendees:    event.Attendees,
	})
}

func (h *Handler) PatchEvent(ctx echo.Context, eventID openapi_types.UUID) error {
	requestBody := api.PatchEventJSONRequestBody{}
	if err := ctx.Bind(&requestBody); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Errorf("bind request body (handler): %w", err))
	}
	eventOnEdit := service.EventOnEdit{
		Title:       requestBody.Title,
		Description: requestBody.Description,
		Date:        requestBody.Date.Time,
		IsOpen:      requestBody.IsOpen,
		Admins:      requestBody.Admins,
		Attendees:   requestBody.Attendees,
	}
	if err := h.EventService.EditEvent(ctx.Request().Context(), eventID, eventOnEdit); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Errorf("edit event (handler): %w", err))
	}

	return ctx.NoContent(http.StatusOK)
}

func (h *Handler) DeleteEvent(ctx echo.Context, eventID openapi_types.UUID) error {
	if err := h.EventService.DeleteEvent(ctx.Request().Context(), eventID); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Errorf("delete event (handler): %w", err))
	}
	return ctx.NoContent(http.StatusOK)
}
