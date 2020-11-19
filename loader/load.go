package loader

import (
    "bufio"
    "net/http"
    "encoding/json"
)

func LoadSegments() *map[string]map[string]int {

    var segments  = map[string]map[string]int{}

    parsed := *loadFromURL(SEGMENTS_URL)
    for _, seg := range parsed.([]interface{}) {

        seg := seg.(map[string]interface{})

        if segments[seg["from"].(string)] == nil {
            segments[seg["from"].(string)] = make(map[string]int)
        }

        segments[seg["from"].(string)][seg["to"].(string)] = 200
    }

    return &segments
}

func LoadEndpoints() *map[string]string {
    var pointNames = map[string]string{}

    parsed := *loadFromURL(ENDPOINTS_URL)
    for _, point := range parsed.([]interface{}){
        point := point.(map[string]interface{})
        pointNames[point["code"].(string)] = point["name"].(string)
    }

    return &pointNames
}

func loadFromURL(url string) *interface{} {

    client := &http.Client{}
    req, err := http.NewRequest("GET", url, nil)
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

    var dat interface{}
    if err := json.Unmarshal([]byte(segmentsStr), &dat); err != nil {
        panic(err)
    }

    return &dat
}