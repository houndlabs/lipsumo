package hello

import (
  "io"
  "io/ioutil"
  "log"
  "encoding/json"
  "fmt"
  "math/rand"
  "net/http"
  "strconv"

  "appengine"
  "appengine/blobstore"
  "appengine/datastore"
  "appengine/memcache"
  "appengine/urlfetch"

  "github.com/ant0ine/go-json-rest"
)

type Book struct {
  Author string
  Title string
  Id int
  Paragraphs int
  Key string
}

type BookBlob struct {
  Book

  Data []string
}

func downloadBook(c appengine.Context, i int) BookBlob {

  client := urlfetch.Client(c)
  resp, _ := client.Get(fmt.Sprintf("http://localhost:8080/books/%d", i))
  defer resp.Body.Close()
  body, _ := ioutil.ReadAll(resp.Body)

  var blob BookBlob
  json.Unmarshal(body, &blob)
  return blob
}

func memcacheBook(c appengine.Context, blob BookBlob) {

  var items []*memcache.Item
  for k, v := range blob.Data {
    paragraph := &memcache.Item{
      Key: fmt.Sprintf("%d:%d", blob.Id, k),
      Value: []byte(v),
    }
    items = append(items, paragraph)
  }

  if err := memcache.SetMulti(c, items); err != nil {
    c.Infof("Memcache error: %s", err)
  }
}

func getParagraph(c appengine.Context, book int, index int) string {

  // Get the item from the memcache
  if item, err := memcache.Get(c, fmt.Sprintf("%d:%d", book, index)); err == memcache.ErrCacheMiss {
    blob := downloadBook(c, book)
    memcacheBook(c, blob)
    return getParagraph(c, book, index)
  } else if err != nil {
    c.Infof("Error: %s", err)
  } else {
    return string(item.Value)
  }

  return ""
}

func GetParagraphs(w *rest.ResponseWriter, r *rest.Request) {
  count, _ := strconv.Atoi(r.FormValue("num"))
  if count == 0 { count = 4 }

  c := appengine.NewContext(r.Request)
  q := datastore.NewQuery("Book")

  var books []Book
  q.GetAll(c, &books)

  var blob BookBlob
  blob.Book = books[rand.Intn(len(books))]
  start := rand.Intn(blob.Paragraphs - count)

  var subset []string
  for i := start; i < start + count; i++ {
    subset = append(subset, getParagraph(c, blob.Id, i))
  }

  blob.Data = subset
  w.WriteJson(blob)
}

func GetBook(w *rest.ResponseWriter, r *rest.Request) {
  c := appengine.NewContext(r.Request)
  key := datastore.NewKey(c, "Book", r.PathParam("id"), 0, nil)
  var book Book
  datastore.Get(c, key, &book)
  blobstore.Send(w.ResponseWriter, appengine.BlobKey(book.Key))
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

func UploadID(w *rest.ResponseWriter, r *rest.Request) {
  c := appengine.NewContext(r.Request)
  uploadURL, err := blobstore.UploadURL(c, "/upload/id", nil)
  if err != nil {
    rest.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  w.WriteJson(map[string]string{ "succes": "true", "url": uploadURL.String() })
}

func BlobID(w http.ResponseWriter, r *http.Request) {
  blobs, _, _ := blobstore.ParseUpload(r)
  io.WriteString(w, fmt.Sprintf("{ \"success\":\"true\", \"key\":\"%s\" }",
    blobs["file"][0].BlobKey))
}

func init() {
  handler := rest.ResourceHandler{}
  handler.SetRoutes(
    rest.Route{"GET", "/paragraphs", GetParagraphs},
    rest.Route{"GET", "/books/:id", GetBook},
    rest.Route{"POST", "/books", StoreBook},
    rest.Route{"GET", "/upload", UploadID},
  )
  http.Handle("/", &handler)
  http.HandleFunc("/upload/id", BlobID)

  log.Print("start")
}
