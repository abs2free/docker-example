```


POST books/_doc
{
  "name": "Snow Crash",
  "author": "Neal Stephenson",
  "release_date": "1992-06-01",
  "page_count": 470
}


POST /_bulk
{"index":{"_index":"books"}}
{"name":"Revelation Space","author":"Alastair Reynolds","release_date":"2000-03-15","page_count":585}
{"index":{"_index":"books"}}
{"name":"1984","author":"George Orwell","release_date":"1985-06-01","page_count":328}
{"index":{"_index":"books"}}
{"name":"Fahrenheit 451","author":"Ray Bradbury","release_date":"1953-10-15","page_count":227}
{"index":{"_index":"books"}}
{"name":"Brave New World","author":"Aldous Huxley","release_date":"1932-06-01","page_count":268}
{"index":{"_index":"books"}}
{"name":"The Handmaids Tale","author":"Margaret Atwood","release_date":"1985-06-01","page_count":311}


GET books

GET books/_search

GET books/_search
{
  "query": {
    "match": {
      "name": "brave"
    }
  }
}

```
