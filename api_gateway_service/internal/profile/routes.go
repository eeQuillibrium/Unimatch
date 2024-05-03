package profile

func (p *profileHandlers) MapRoutes() {
	p.group.GET("/", p.profileHandler())
	p.group.PUT("/setProfile", p.setProfileHandler())
	p.group.GET("/getProfile", p.getProfileHandler())
}