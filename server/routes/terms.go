package routes

import (
	"net/http"
	"text/template"

	"github.com/julienschmidt/httprouter"
	"github.com/repejota/logger"
)

// Terms is the /terms route
func Terms(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	l := logger.New("default")

	t, err := template.ParseFiles(
		"templates/terms.html",
		"templates/partials/facebook.html",
		"templates/partials/footer.html",
		"templates/partials/javascript.html",
		"templates/partials/css.html",
	)
	if err != nil {
		l.Errorf(err.Error())
	}
	context := struct {
		Title string
	}{
		"tvt.io",
	}
	t.Execute(w, context)
}
