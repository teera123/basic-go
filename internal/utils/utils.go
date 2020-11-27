package utils

import (
	"net/http"

	"github.com/labstack/echo"

	"academy/internal/model"
)

func ResponseJSON(c echo.Context, data interface{}, err error) error {
	success, code, status := true, http.StatusOK, "success"
	if err != nil {
		success, code, status = false, http.StatusInternalServerError, err.Error()
	}
	res := model.Response{
		Success: success,
		Code:    code,
		Message: status,
		Data:    data,
	}
	return c.JSON(code, res)
}
