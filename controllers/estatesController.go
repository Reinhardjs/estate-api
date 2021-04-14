package controllers

import (
	"encoding/json"
	"net/http"
	"simple-api/models"
	u "simple-api/utils"
	"strconv"

	"github.com/gorilla/mux"
)

var CreateEstate = func(w http.ResponseWriter, r *http.Request) {

	//user := r.Context().Value("user").(uint) //Grab the id of the user that send the request
	estate := &models.Estate{}

	err := json.NewDecoder(r.Body).Decode(estate)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	resp := estate.Create()
	u.Respond(w, resp)
}

var GetEstate = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		//The passed path parameter is not an integer
		u.Respond(w, u.Message(false, "There was an error in your request"))
		return
	}

	data := models.GetEstate(uint(id))
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

var GetEstates = func(w http.ResponseWriter, r *http.Request) {
	// id := r.Context().Value("user").(uint)
	data := models.GetEstates()
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}
