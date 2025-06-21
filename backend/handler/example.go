package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/eraxyso/go-template/api"
)

func (h *Handler) PostExample(c echo.Context, params api.PostExampleParams) error {
	requestBody := api.PostExampleJSONRequestBody{}
	if err := c.Bind(&requestBody); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	if err := h.ExampleService.ExampleServiceMethod(c, params.ParamExample, requestBody.RequestBody); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.NoContent(http.StatusOK)
}
