package main

func initializeRoutes() {

	router.Use(setUserStatus())

	router.GET("/", showIndexPage)

	userRoutes := router.Group("/u")
	{
		userRoutes.GET("/login", ensureNotLoggedIn(), showLoginPage)

		userRoutes.POST("/login", ensureNotLoggedIn(), performLogin)

		userRoutes.GET("/logout", ensureLoggedIn(), logout)

		userRoutes.GET("/register", ensureNotLoggedIn(), showRegistrationPage)

		userRoutes.POST("/register", ensureNotLoggedIn(), register)
	}

	productRoutes := router.Group("/product")
	{
		productRoutes.GET("/view/:product_id", getProduct)

		productRoutes.GET("/create", ensureLoggedIn(), showProductCreationPage)

		productRoutes.POST("/create", ensureLoggedIn(), createProduct)
	}
}
