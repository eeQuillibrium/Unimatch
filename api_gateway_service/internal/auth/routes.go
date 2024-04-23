package auth


func (a *authHandlers) MapRoutes() {
	a.echogroup.POST("/signUp", a.signUpHandler())
	a.echogroup.POST("/signIn", a.signInHandler())
}