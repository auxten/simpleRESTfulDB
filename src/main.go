package main

import (
    "fmt"
    "log"
    "net/http"
    "strings"
    "sync"
    "html"
)

type db struct {
    l sync.Mutex
    db_map map[string]string
}
var DB_g = db{db_map:make(map[string]string)}

type Hello struct{}

func (h Hello) ServeHTTP( w http.ResponseWriter, r *http.Request) {
    url_list := strings.Split(html.EscapeString(r.URL.Path), "/")
    if url_list[1] == "get" && len(url_list) == 3{
        DB_g.l.Lock()
        defer DB_g.l.Unlock()
        v, exist := DB_g.db_map[url_list[2]]
        if exist {
            w.Write([]byte(v))
        } else {
            w.WriteHeader(http.StatusNotFound)
            w.Write([]byte(fmt.Sprintf("key %s not found", url_list[2])))
        }
    } else if url_list[1] == "set" && len(url_list) == 4{
        DB_g.l.Lock()
        defer DB_g.l.Unlock()
        DB_g.db_map[url_list[2]] = url_list[3]
    } else {
        w.WriteHeader(http.StatusNotImplemented)
        w.Write([]byte(fmt.Sprint("error request length")))
    }
    return 
}




func main() {
    var h Hello
    err := http.ListenAndServe("0.0.0.0:4000", h)
    if err != nil {
        log.Fatal(err)
    }
}

