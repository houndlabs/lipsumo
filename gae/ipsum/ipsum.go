package hello

import (
  "appengine"
  "appengine/datastore"
  "github.com/ant0ine/go-json-rest"
  "net/http"
  "strconv"
)

type Book struct {
  Author string
  Title string
  Id int
  Paragraphs []string `datastore:",noindex"`
}

func GetParagraphs(w *rest.ResponseWriter, r *rest.Request) {
  w.WriteJson(map[string]string{"a": "b"})
}

func StoreBook(w *rest.ResponseWriter, r *rest.Request) {
  book := Book{}
  r.DecodeJsonPayload(&book)
  context := appengine.NewContext(r.Request)
  key := datastore.NewKey(context, "Book", strconv.Itoa(book.Id), 0, nil)
  if _, err := datastore.Put(context, key, &book); err != nil {
    rest.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }


  w.WriteJson(map[string]string{ "success": "true" })
}

func init() {
  handler := rest.ResourceHandler{}
  handler.SetRoutes(
    rest.Route{"GET", "/paragraphs", GetParagraphs},
    rest.Route{"POST", "/paragraphs", StoreBook},
  )
  http.Handle("/", &handler)
}
