package dto

import (
	"context"
	"errors"
	"mime/multipart"
	"strconv"
)

type Profile struct {
	UserId int64
	Name   string
	Age    int64
	About  string
	Avatar []byte
}

func AccessProfile(
	ctx context.Context,
	mForm *multipart.Form,
) (*Profile, error) {
	name, ok := mForm.Value["name"]
	if !ok {
		return nil, errors.New("accessprofile() name")
	}

	ages, _ := mForm.Value["age"]
	age, err := strconv.Atoi(ages[0])
	if err != nil {
		return nil, errors.New("accessprofile() age")
	}

	about, ok := mForm.Value["about"]
	if !ok {
		return nil, errors.New("accessprofile() about")
	}

	avatarH, ok := mForm.File["avatar"]
	if !ok {
		return nil, errors.New("accessprofile() avatarh")
	}
	avatarFile, err := avatarH[0].Open()
	if err != nil {
		return nil, err
	}
	avatarB := make([]byte, avatarH[0].Size)
	avatarFile.Read(avatarB)
	return &Profile{
		Name:   name[0],
		Age:    int64(age),
		About:  about[0],
		Avatar: avatarB,
	}, nil
}
