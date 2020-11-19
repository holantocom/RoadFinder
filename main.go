package main

import (
    "github.com/albertorestifo/dijkstra"
    "./loader"
    "./schemas"
    "./controller"
    "net/http"
    "log"
)

var g = dijkstra.Graph{}
var p = map[string]string{}

func matching(w http.ResponseWriter, req *http.Request) {

    from, okFrom := req.URL.Query()["from"]
    to, okTo := req.URL.Query()["to"]

    if !okFrom || !okTo {
        answer := schemas.ErrorMessage{
            Code: 1, 
            Message: "URL param missing",
        }
        controller.Answer(w, 400, answer)
        return
    }

    path, cost, err := g.Path(from[0], to[0])

    if err != nil || len(path) <= 1 {
        answer := schemas.ErrorMessage{
            Code: 2, 
            Message: "Route not found",
        }
        controller.Answer(w, 400, answer)
        return
    }

    cityNames := make([]string, 0)
    for _, city := range path {
        cityNames = append(cityNames, p[city])
    }

    answer := schemas.Matching{
        GeoList: path,
        Cost: cost,
        Cities: cityNames,
    }

    controller.Answer(w, 200, answer)
}

func main() {

    go func() {
        g = *loader.LoadSegments()
        log.Println("Segments Loaded")
    }()

    go func() {
        p = *loader.LoadEndpoints()
        log.Println("Endpoints Loaded")
    }()

    http.HandleFunc("/matching", matching)

    http.ListenAndServe(":8090", nil)

}