# LIPSUMO - HEAVYWEIGHT LOREM IPSUM

Instead of using auto-generated text, why not use something more interesting? This project takes the top books off Project Gutenberg and makes them available to you in small (random, thought provoking) pieces. Make sure you check out the [website](http://lipsumo.hound.io).

Because the world is nothing but APIs all the way down, make sure you check out the Lipsumo API. Then, integrate all this amazing (mind blowing, world changing) functionality into your own projects!

# API

There's only one endpoint for the API and the response contains everything you need:

    curl http://lipsumo.hound.io/paragraphs

    {
      "Author": "Frederick Douglass",
      "Title": "The Narrative of the Life of Frederick Douglass",
      "Id": 23,
      "Data": [ 
        "paragraph1", 
        "paragraph2", 
        "paragraph3", 
        "paragraph4" 
      ]
    }

If you'd like to get a different number of paragraphs outside the default of four, it is a query parameter:

    curl http://lipsumo.hound.io/paragraphs?num=2

For a live example, take a look at the [website](http://lipsumo.hound.io). It uses the standard API every time a different number of paragraphs are selected.

# Development

For anyone interested in using this project themselves or doing some development, it is pretty easy to get setup and running locally.

1. Follow the instructions to install [GAE GO SDK](https://developers.google.com/appengine/docs/go/gettingstarted/devenvironment)
2. Setup the local app config

    cp gae/app.yaml.example gae/app.yaml

3. Fetch the dependencies

    goapp get github.com/ant0ine/go-json-rest

4. Serve away!

    goapp serve gae

# Adding new books

The books are all compiled into json files that get checked in (and deployed). You can either add new books or change the structure of the existing ones by modifying `convert.py`.

1. Download the text version of the book you're interested in from [Project Gutenberg](http://www.gutenberg.org/)
2. Save the downloaded book into `./raw_books/` and name it like `pg100.txt` where 100 is the ID of the book.
3. Run `convert.py`

# Deploying your own version

App Engine makes it easy for anyone to run their own version of Lipsumo. Follow the [instructions](https://developers.google.com/appengine/docs/go/gettingstarted/uploading) and upload via. `goapp deploy gae`. You'll have your own version running and will be able to do whatever you'd like with it.


