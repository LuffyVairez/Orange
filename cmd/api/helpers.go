// Filename: cmd/api/helpers.go
package main

import (
	"errors"
	"net/http"
	"strconv"
     "encording/json"
	"github.com/julienschmidt/httprouter"
)

func (app *application) readIDParams(r *http.Request) (int64, error) {
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("invalid Id Parameter")
	}
	return id, nil
}
