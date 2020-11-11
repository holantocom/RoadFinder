package loader

import (
    "bufio"
    "net/http"
    "encoding/json"
)

func LoadSegments() *[](map[string]string) {

    client := &http.Client{}
    req, err := http.NewRequest("GET", SEGMENTS_URL, nil)
    req.SetBasicAuth(AUTH_USER, AUTH_PASS)
    resp, err := client.Do(req)

    if err != nil{
        panic(err)
    }

    scanner := bufio.NewScanner(resp.Body)
    buf := make([]byte, 0, 64*1024)
    scanner.Buffer(buf, 1024*1024)
    var segmentsStr string

    for scanner.Scan() {
        segmentsStr = segmentsStr + scanner.Text()
    }

    if err := scanner.Err(); err != nil {
        panic(err)
    }

    var dat [](map[string]string)
    if err := json.Unmarshal([]byte(segmentsStr), &dat); err != nil {
        panic(err)
    }

    return &dat
}