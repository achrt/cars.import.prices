package v1

import (
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"

	"cars.import.prices/api"
)

func Wrapper() func(callable func(c *api.Context, a *api.Application, w http.ResponseWriter, r *http.Request) (int, interface{}, error)) func(w http.ResponseWriter, r *http.Request) {
	return func(callable func(c *api.Context, a *api.Application, w http.ResponseWriter, r *http.Request) (int, interface{}, error)) func(w http.ResponseWriter, r *http.Request) {
		return func(w http.ResponseWriter, r *http.Request) {
			context := r.Context().Value("context")

			if context == nil {
				JsonError(w, http.StatusInternalServerError, errors.New("invalid request context"))
			} else {
				c := context.(*api.Context)
				c.Vars = mux.Vars(r)

				a := &api.Application{}

				bodyFailed := false

				if r.Method == http.MethodPost || r.Method == http.MethodPatch || r.Method == http.MethodPut {
					body, err := ioutil.ReadAll(r.Body)

					if err != nil {
						bodyFailed = true
						JsonError(w, http.StatusBadRequest, errors.New("cannot read request"))
					} else {
						c.Body = body
					}
				}

				if !bodyFailed {
					status, response, err := callable(c, a, w, r)

					if err != nil {
						if r.Method == http.MethodGet || r.Method == http.MethodDelete {
							status = http.StatusNotFound
						}

						JsonError(w, status, err)
					} else {
						if status == http.StatusCreated {
							JsonSuccessCreated(w, response)
						} else {
							if status == http.StatusAccepted {
								JsonSuccessAccepted(w, response)
							} else {
								JsonSuccess(w, response)
							}
						}
					}
				}
			}
		}
	}
}
