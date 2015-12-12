package routes

import (
	"net/http"
	"text/template"

	"github.com/julienschmidt/httprouter"
	"github.com/tvtio/front/catalog"
	"github.com/tvtio/front/logger"
	"github.com/tvtio/front/tmdb"
)

// TV is the /tv/:id route
func TV(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	tv, err := catalog.TV(id)
	if err != nil {
		logger.Fatal(err.Error())
	}
	context := struct {
		Title string
		TV    tmdb.TV
	}{
		"tvt.io",
		tv,
	}
	t, err := template.ParseFiles(
		"templates/tv.html",
		"templates/partials/facebook.html",
		"templates/partials/footer.html",
		"templates/partials/javascript.html",
		"templates/partials/css.html",
	)
	if err != nil {
		logger.Fatal(err.Error())
	}
	t.Execute(w, context)
}
