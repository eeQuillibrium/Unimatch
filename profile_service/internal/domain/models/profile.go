package models

import kafkaMessages "github.com/eeQuillibrium/Unimatch/proto/gen/go/kafka"

type Profile struct {
	UserID    int
	Name      string
	Age       int
	About     string
	ImagePath string
}

func AccessProfile(msg *kafkaMessages.Profile, imgPath string) *Profile {
	var profile Profile
	profile.UserID = int(msg.GetUserID())
	profile.Name = msg.GetName()
	profile.Age = int(msg.GetAge())
	profile.About = msg.GetAbout()
	profile.ImagePath = imgPath

	return &profile
}
