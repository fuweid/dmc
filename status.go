package dmc

import "github.com/fuweid/echo"

func Status() string {
	return echo.Version()
}
