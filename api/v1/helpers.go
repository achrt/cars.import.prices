package v1

import (
	"encoding/json"
	"net/http"

	"cars.import.prices/presenters/jsonapi"
)

func JsonError(w http.ResponseWriter, status int, err error) {
	j, _ := json.Marshal(jsonapi.NewErrorResponse(status, err))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(j)
}

func JsonSuccessCreated(w http.ResponseWriter, r interface{}) {
	j, _ := json.Marshal(r)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(j)
}

func JsonSuccess(w http.ResponseWriter, r interface{}) {
	j, _ := json.Marshal(r)

	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

func JsonSuccessAccepted(w http.ResponseWriter, r interface{}) {
	j, _ := json.Marshal(r)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(j)
}

func BadRequest(err error) (int, interface{}, error) {
	return http.StatusBadRequest, nil, err
}

func InternalServerError(err error) (int, interface{}, error) {
	return http.StatusInternalServerError, nil, err
}

func OK(body interface{}) (int, interface{}, error) {
	return http.StatusOK, body, nil
}

func Created(body interface{}) (int, interface{}, error) {
	return http.StatusCreated, body, nil
}
