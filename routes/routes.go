package routes

import (
	"myeduate/ent"
	"myeduate/routes/gql"

	"github.com/labstack/echo/v4"
)

func Init(app *echo.Echo, client *ent.Client) {
	//auth.Init(app, client)
	gql.Init(app, client)
}
