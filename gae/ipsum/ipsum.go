package ipsum

import(
  "io/ioutil"
  "log"
  "encoding/json"
  "fmt"
  "math/rand"
  "net/http"
  "strconv"

  "github.com/ant0ine/go-json-rest"
)

type Book struct {
  Author string
  Title string
  Id int
  Paragraphs int
  Data []Paragraph
}

type Paragraph struct {
  Text string
  Characters int
  Sentences int
}

type Response struct {
  Author string
  Title string
  Data []string
}

func getParagraphs(w *rest.ResponseWriter, r *rest.Request) {
  count, _ := strconv.Atoi(r.FormValue("num"))
  if count == 0 { count = 4 }

  available_books, _ := ioutil.ReadDir("books")

  book_name := available_books[rand.Intn(len(available_books))].Name()
  raw_book, _ := ioutil.ReadFile(fmt.Sprintf("books/%s", book_name))

  var book Book
  var resp Response
  json.Unmarshal(raw_book, &book)
  start := rand.Intn(book.Paragraphs - count)

  var subset []string
  for _, v := range book.Data[start:start+count] {
    subset = append(subset, v.Text)
  }

  resp.Author = book.Author
  resp.Title = book.Title
  resp.Data = subset

  w.WriteJson(resp)
}

func init() {
  handler := rest.ResourceHandler{}
  handler.SetRoutes(
    rest.Route{"GET", "/paragraphs", getParagraphs},
  )
  http.Handle("/", &handler)

  log.Print("start")
}
