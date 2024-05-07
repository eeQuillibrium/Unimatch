package profile

func (p *profileHandlers) MapRoutes() {
	p.group.POST("", p.setProfileHandler())
	p.group.GET("/form", p.formProfileHandler())
	p.group.GET("", p.getProfileHandler())
}
