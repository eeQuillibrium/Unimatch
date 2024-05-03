package auth

func (a *authHandlers) MapRoutes() {
	a.echogroup.GET("/", a.authHandler())
	a.echogroup.POST("/signUp", a.signUpHandler())
	a.echogroup.POST("/signIn", a.signInHandler())
}
