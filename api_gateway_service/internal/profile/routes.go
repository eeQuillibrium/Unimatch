package profile

func (p *profileHandlers) MapRoutes() {
	p.group.PUT("/setProfile", p.setProfileHandler())
	p.group.PUT("/getProfile", p.getProfileHandler())
}