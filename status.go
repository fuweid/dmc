package dmc

import "github.com/fuweid/echo"

func Status() string {
	return echo.Version()
}

func Release() string {
	return "2019-11-11"
}
