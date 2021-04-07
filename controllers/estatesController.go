package controllers

import (
	"encoding/json"
	"net/http"
	"simple-api/models"
	u "simple-api/utils"
)

var CreateEstate = func(w http.ResponseWriter, r *http.Request) {

	user := r.Context().Value("user").(uint) //Grab the id of the user that send the request
	estate := &models.Estate{}

	err := json.NewDecoder(r.Body).Decode(estate)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	estate.UserId = user
	resp := estate.Create()
	u.Respond(w, resp)
}

var GetEstatesFor = func(w http.ResponseWriter, r *http.Request) {

	id := r.Context().Value("user").(uint)
	data := models.GetEstates(id)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}
