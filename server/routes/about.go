// Copyright 2016 The front Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package routes

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/goincremental/negroni-sessions"
	"github.com/julienschmidt/httprouter"
)

// About is the /about route
func About(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	session := sessions.GetSession(r)

	// Get user
	_ = fmt.Sprint(session.Get("user"))

	t := template.Must(template.ParseFiles(
		"templates/about.html",
		"templates/partials/facebook.html",
		"templates/partials/footer.html",
		"templates/partials/javascript.html",
		"templates/partials/css.html",
	))

	context := struct {
		Title string
	}{
		"tvt.io",
	}
	t.Execute(w, context)
}
