package profile

func (p *profileHandlers) MapRoutes() {
	p.group.PUT("/", p.setProfileHandler())
	p.group.GET("/", p.getProfileHandler())
}