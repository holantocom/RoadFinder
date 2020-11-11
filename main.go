package main

import (
    "fmt"
    "github.com/albertorestifo/dijkstra"
    "./loader"
    "net/http"
    "log"
)

var g = dijkstra.Graph{}

func matching(w http.ResponseWriter, req *http.Request) {

    from, okFrom := req.URL.Query()["from"]
    to, okTo := req.URL.Query()["to"]

    if !okFrom || !okTo {
        log.Println("Url Param is missing")
        return
    }

    path, cost, err := g.Path(from[0], to[0])

    if err != nil {
        log.Println(err)
        return
    }

    fmt.Fprintf(w, "path: %v, cost: %v\n", path, cost)
}

func waitSegments(){
    var segments  = map[string]map[string]int{}

    for _, seg := range (*loader.LoadSegments()) {

        if segments[seg["from"]] == nil {
            segments[seg["from"]] = make(map[string]int)
        }

        segments[seg["from"]][seg["to"]] = 1
    }

    g = segments
}

func main() {

    go waitSegments()

    http.HandleFunc("/matching", matching)

    http.ListenAndServe(":8090", nil)

}