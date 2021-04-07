package models

import (
	"fmt"
	u "simple-api/utils"

	"github.com/jinzhu/gorm"
)

type Estate struct {
	gorm.Model
	Name   string `json:"name"`
	Phone  string `json:"phone"`
	UserId uint   `json:"user_id"` //The user that this estate belongs to
}

/*
 This struct function validate the required parameters sent through the http request body
returns message and true if the requirement is met
*/
func (estate *Estate) Validate() (map[string]interface{}, bool) {

	if estate.Name == "" {
		return u.Message(false, "estate name should be on the payload"), false
	}

	if estate.Phone == "" {
		return u.Message(false, "Phone number should be on the payload"), false
	}

	if estate.UserId <= 0 {
		return u.Message(false, "User is not recognized"), false
	}

	//All the required parameters are present
	return u.Message(true, "success"), true
}

func (estate *Estate) Create() map[string]interface{} {

	if resp, ok := estate.Validate(); !ok {
		return resp
	}

	GetDB().Create(estate)

	resp := u.Message(true, "success")
	resp["estate"] = estate
	return resp
}

func GetEstate(id uint) *Estate {

	estate := &Estate{}
	err := GetDB().Table("estates").Where("id = ?", id).First(estate).Error
	if err != nil {
		return nil
	}
	return estate
}

func GetEstates(user uint) []*Estate {

	estates := make([]*Estate, 0)
	err := GetDB().Table("estates").Where("user_id = ?", user).Find(&estates).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return estates
}
