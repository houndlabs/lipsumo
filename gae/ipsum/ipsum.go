package ipsum

import(
  "log"
  "net/http"

  "github.com/ant0ine/go-json-rest"
)

func getParagraphs(w *rest.ResponseWriter, r *rest.Request) {
}

func init() {
  handler := rest.ResourceHandler{}
  handler.SetRoutes(
    rest.Route{"GET", "/paragraphs", getParagraphs},
  )
  http.Handle("/", &handler)

  log.Print("start")
}
