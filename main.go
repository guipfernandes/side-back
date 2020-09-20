package main

	import (
		"side/core"
	
		"github.com/labstack/echo/v4"
		_"github.com/lib/pq"
	)
	
	// Entry point application.
	func main() {
		api := &core.API{
			Echo: echo.New(),
		}
		api.LoadDefault().StartAPI()
	}	
