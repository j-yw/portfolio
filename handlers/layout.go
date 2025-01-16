package handlers

import (
	"github.com/j-yw/portfolio/view/layout"
	"github.com/labstack/echo/v4"
)

type PageHandler struct{}

func (h PageHandler) HandlePage(c echo.Context) error {
	return render(c, layout.Base())
}
