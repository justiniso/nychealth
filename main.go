package main

import (
    // "encoding/json";
    "log";
)

var (
    path = "/resource/xx67-kt59.json"
)

func main() {
    var queryParams map[string]string

    queryParams = make(map[string]string)
    queryParams["$q"] = "Sido Gourmet"

    client := NewSodaClient(path)
    body := client.Get(queryParams)

    agg := AggregateInspections(body)

    log.Println(agg)


}