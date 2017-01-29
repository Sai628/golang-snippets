package main

import (
    "encoding/json"
    "encoding/xml"
    "net/http"
    "path"
    "runtime"
)

type Profile struct {
    Name string
    Hobbies []string `xml:"Hobbies>Hobby"`
}

func sendingHeaders(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Server", "A Go Web Server")
    w.WriteHeader(http.StatusOK)
}

func renderingPlainText(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("OK"))
}

func renderingJSON(w http.ResponseWriter, r *http.Request) {
    profile := Profile{"Alex", []string{"snowboarding", "programming"}}
    js, err := json.Marshal(profile)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.Write(js)
}

func renderingXML(w http.ResponseWriter, r *http.Request) {
    profile := Profile{"Alex", []string{"snowboarding", "programming"}}
    x, err := xml.MarshalIndent(profile, "", "  ")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/xml")
    w.Write(x)
}

func sendingFile(w http.ResponseWriter, r *http.Request) {
    _, filename, _, _ := runtime.Caller(0)
    fp := path.Join(path.Dir(filename), "images", "foo.jpg")
    http.ServeFile(w, r, fp)
}

func main() {
    http.HandleFunc("/sending_headers", sendingHeaders)
    // $ curl -i localhost:3000/sending_headers
    // HTTP/1.1 200 OK
    // Server: A Go Web Server
    // Date: Sun, 29 Jan 2017 07:21:27 GMT
    // Content-Length: 0
    // Content-Type: text/plain; charset=utf-8


    http.HandleFunc("/rendering_plain_text", renderingPlainText)
    // $ curl -i localhost:3000/rendering_plain_text
    // HTTP/1.1 200 OK
    // Date: Sun, 29 Jan 2017 07:25:33 GMT
    // Content-Length: 2
    // Content-Type: text/plain; charset=utf-8
    //
    // OK


    http.HandleFunc("/rendering_json", renderingJSON)
    // $ curl -i localhost:3000/rendering_json
    // HTTP/1.1 200 OK
    // Content-Type: application/json
    // Date: Sun, 29 Jan 2017 07:30:19 GMT
    // Content-Length: 56
    //
    // {"Name":"Alex","Hobbies":["snowboarding","programming"]}


    http.HandleFunc("/rendering_xml", renderingXML)
    // $ curl -i localhost:3000/rendering_xml
    // HTTP/1.1 200 OK
    // Content-Type: application/xml
    // Date: Sun, 29 Jan 2017 08:16:26 GMT
    // Content-Length: 128
    //
    // <Profile>
    // <Name>Alex</Name>
    // <Hobbies>
    // <Hobby>snowboarding</Hobby>
    // <Hobby>programming</Hobby>
    // </Hobbies>
    // </Profile>


    http.HandleFunc("/sending_file", sendingFile)
    // $ curl -I localhost:3000/sending_file
    // HTTP/1.1 200 OK
    // Accept-Ranges: bytes
    // Content-Length: 10862
    // Content-Type: image/jpeg
    // Last-Modified: Sun, 29 Jan 2017 08:27:31 GMT
    // Date: Sun, 29 Jan 2017 08:49:51 GMT


    http.ListenAndServe(":3000", nil)
}
