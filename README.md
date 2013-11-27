lit-ipsum
=========

1. Follow the instructions to install [GAE GO SDK](https://developers.google.com/appengine/docs/go/gettingstarted/devenvironment)
2. Start the dev server

    goapp serve gae

Rebuild Books
=============

    python convert.py

API
===

    curl "localhost:8080/paragraphs?num=10"
