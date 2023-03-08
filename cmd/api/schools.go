// Filename: cmd/api/schools.go
package main

import (
	"fmt"
	"net/http"
	"strconv"
	"github.com/julienschmidt/httprouter"
)
// createSchool handler for the "post /v1/schools" endpoint
func (app *application) createSchoolHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Created a school...")
}
// showSchoolHandler for the "GET /v1/schools/:id" endpoint
func (app *application) showSchoolHandler(w http.ResponseWriter, r *http.Request) {
	params :=httprouter.ParamsFromContext(r.Context())
	// fmt.Fprintln(w, "School displayed...")
	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)

	if err != nil || id < 1 {
		http.NotFound(w,r)
		return
	} 
	// Display the school id
	fmt.Fprintf( w, "show the details for school %d\n", id)

	

}