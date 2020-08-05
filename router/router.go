package router

import (
	"github.com/labstack/echo/v4"
)

// InitRoutes Initialize Routing
func InitRoutes(v1 *echo.Group) {
	SetCampaignRoutes(v1)
	SetUserRoutes(v1)
	SetStaffRoute(v1)
}
