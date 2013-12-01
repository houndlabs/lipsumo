package ipsum

import (
  "html/template"
  "io/ioutil"
  "log"
  "encoding/json"
  "fmt"
  "math/rand"
  "net/http"
  "os"
  "strconv"
  "time"

  "github.com/ant0ine/go-json-rest"
)

var (
  config *Config
)

type Config struct {
  Mailchimp string
}

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
  Id int
  Data []string
}

func loadConfig() {
  file, err := ioutil.ReadFile("config.json")
  if err != nil {
    log.Println("open config: ", err)
    os.Exit(1)
  }

  temp := new(Config)
  if err = json.Unmarshal(file, temp); err != nil {
    log.Println("parse config: ", err)
    os.Exit(1)
  }
  config = temp
}

func selectParagraphs(count int) Response {
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
  resp.Id = book.Id

  return resp
}

func getParagraphs(w *rest.ResponseWriter, r *rest.Request) {
  count, _ := strconv.Atoi(r.FormValue("num"))
  if count == 0 { count = 4 }

  w.WriteJson(selectParagraphs(count))
}

func showIndex(w *rest.ResponseWriter, r *rest.Request) {
  t, _ := template.ParseFiles(
    "tmpl/about.html",
    "tmpl/foot.html",
    "tmpl/head.html",
    "tmpl/index.html",
    "tmpl/metadata.html",
  )

  var tmpl_data struct {
    Response
    Config
  }

  tmpl_data.Response = selectParagraphs(4)
  tmpl_data.Config = *config

  log.Println(tmpl_data)
  t.ExecuteTemplate(w, "index", tmpl_data)
}

func init() {
  // -- If you don't seed the golang RNG ... it isn't random at all.
  rand.Seed(time.Now().UTC().UnixNano())

  loadConfig()

  handler := rest.ResourceHandler{}
  handler.SetRoutes(
    rest.Route{"GET", "/", showIndex},
    rest.Route{"GET", "/paragraphs", getParagraphs},
  )
  http.Handle("/", &handler)

  log.Print("started...")
}
