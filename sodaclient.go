package main

import (
    "fmt";
    "io/ioutil";
    "log";
    "net/http";
    "net/url";
)

const (
    uri = "https://data.cityofnewyork.us"
)

type SodaClient struct {
    resourcePath string
    requests *http.Client
}

// Instantite a new SodaClient for a given path. You should create a new 
// instance of the client for each resource. 
func NewSodaClient(resourcePath string) *SodaClient {
    sc := &SodaClient{
        resourcePath: resourcePath,
        requests: &http.Client{},
    }

    return sc
}

// Handle errors; panic if the error is not null, but this behavior could 
// change in a different mode (debug, etc.)
func HandleErr(err error) {
    if err != nil {
        log.Panic(fmt.Sprintf("[PANIC] %v", err))
        panic(err)
    }
}

// Return a query string from a map of data
func Urlencode(data map[string]string) string {
    params := url.Values{}
    for k, v := range data {
        params.Add(k, v)
    }
    return params.Encode()
}

// Accepts a response and returns the body bytestring
func ReadResponse(resp *http.Response, err error) []byte {
    HandleErr(err)
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    HandleErr(err)
    return body
}

func (sc *SodaClient) Get (params map[string]string) []byte {
    get := fmt.Sprintf("%s%s?%s", uri, sc.resourcePath, Urlencode(params))

    log.Printf("[GET] %s", get)
    resp, err := sc.requests.Get(get)

    return ReadResponse(resp, err)
}
