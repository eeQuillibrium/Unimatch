package authservice


func (a *authservice) MapRoutes() {
	a.echogroup.POST("/signUp", a.signUpHandler())
	a.echogroup.POST("/signIn", a.signInHandler())
}