package routes

import "../controllers"

func StatusRoutes(c *controllers.Controller) (r Routes) {
	r = Routes{
		Route{
			"Status",
			"GET",
			c.Session.LiteConfig.Config["api"]["mainroute"] + "/status",
			c.Status,
		},
	}
	return
}
