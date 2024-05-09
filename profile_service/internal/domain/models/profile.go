package models

import (
	"strconv"

	kafkaMessages "github.com/eeQuillibrium/Unimatch/proto/gen/go/kafka"
)

type Profile struct {
	UserID    int    `sqlx:"id"`
	Name      string `sqlx:"name"`
	Age       int    `sqlx:"age"`
	About     string `sqlx:"about"`
	ImagePath string `sqlx:"img_path"`
}

func AccessKafkaProfile(msg *kafkaMessages.Profile, imgPath string) *Profile {
	var profile Profile
	profile.UserID = int(msg.GetUserID())
	profile.Name = msg.GetName()
	profile.Age = int(msg.GetAge())
	profile.About = msg.GetAbout()
	profile.ImagePath = imgPath

	return &profile
}
func AccessRedisProfile(m map[string]string, userID int) (*Profile, error) {
	var profile Profile
	profile.UserID = userID
	profile.Name = m["name"]
	age, err := strconv.Atoi(m["age"])
	if err != nil {
		return nil, err
	}
	profile.Age = age
	profile.About = m["about"]
	profile.ImagePath = m["imagepath"]
	return &profile, nil
}
