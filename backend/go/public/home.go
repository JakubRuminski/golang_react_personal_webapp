package public

import (
	"github.com/jakubruminski/golang/golang_webapp_utils/html_utils"
	"net/http"
)

func HandleHomePage(w http.ResponseWriter, r *http.Request) {

	html_utils.ServeHtml(w, "../ui/static/html/home.html")
}
