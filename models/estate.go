package models

import (
	"fmt"
	u "simple-api/utils"

	"github.com/jinzhu/gorm"
)

type Estate struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	ModelUrl    string `json:"modelUrl`
}

/*
 This struct function validate the required parameters sent through the http request body
returns message and true if the requirement is met
*/
func (estate *Estate) Validate() (map[string]interface{}, bool) {

	if estate.Title == "" {
		return u.Message(false, "Estate title should be on the payload"), false
	}

	if estate.Description == "" {
		return u.Message(false, "Description should be on the payload"), false
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
