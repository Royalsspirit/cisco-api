package api

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/Royalsspirit/cisco-api/internal/api/models"
	"github.com/Royalsspirit/cisco-api/internal/api/util"
)

func (s *Server) list(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	characters, err := models.AllPeoples(s.DB)

	//Add business logic to characters items if needed

	if err != nil {
		util.ResponseError(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := json.NewEncoder(w).Encode(characters); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func (s *Server) update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	body := json.NewDecoder(r.Body)

	var input models.People

	err := body.Decode(&input)
	/**
	* if body decoding failed stop execution
	**/
	if err != nil {

		util.ResponseError(w, "Internal error", http.StatusInternalServerError)
		return
	}
	// Init validator
	validate, trans := s.Validator.Val, s.Validator.Trans

	validateErr := validate.Struct(input)

	if validateErr != nil {

		errors := util.BuildErrorsFromValidationErrors(validateErr.(validator.ValidationErrors), &trans)

		util.ResponseErrors(w, errors, http.StatusBadRequest)
		return
	}

	err = models.PeopleUpdateHandler(s.DB, input, vars["id"])

	if err != nil {
		util.ResponseError(w, "Internal error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (s *Server) create(w http.ResponseWriter, r *http.Request) {

}

func (s *Server) delete(w http.ResponseWriter, r *http.Request) {

}
