package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

type Config struct {
	ServeHttps bool   `json:"ServeHttps"`
	Path       string `json:"Path"`
	CertPath   string `json:"CertPath"`
	KeyPath    string `json:"KeyPath"`
	ConfigPath string
}

var config Config

func main() {
	parseConfig("configs/server-config.json", &config)

	config.Path = strings.TrimRight(config.Path, "/")

	http.HandleFunc("/", handler)

	if config.ServeHttps {
		// start HTTPS server
		log.Fatal(http.ListenAndServeTLS(config.Path, config.CertPath, config.KeyPath, nil))
	} else {
		// start HTTP server
		log.Fatal(http.ListenAndServe(":80", nil))
	}
}

func parseConfig(path string, config *Config) {
	f, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	byteContent, err := ioutil.ReadAll(f)

	if err != nil {
		panic(err)
	}

	if !json.Valid(byteContent) {
		panic("The configuration file is not a valid JSON file.")
	}

	//By default, object keys which don't have a corresponding struct field are ignored
	json.Unmarshal([]byte(byteContent), config)

	f.Close()
}

func handler(w http.ResponseWriter, r *http.Request) {
	if strings.HasPrefix(r.URL.Path, "/search") {
		//routePath(w, r, trimmedURL)
		getSearchResults(w, r)
	} else {

		if r.URL.Path == "/" {
			http.ServeFile(w, r, config.Path+"/index.html")
		} else {
			http.ServeFile(w, r, config.Path+r.URL.Path)
		}

	}
}

func getSearchResults(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")

	if query == "" {
		return
	}

	req, err := http.NewRequest("GET", "https://www.youtube.com/results", nil)

	if err != nil {
		panic(err)
	}

	q := req.URL.Query()
	q.Add("search_query", query)
	req.URL.RawQuery = q.Encode()

	client := &http.Client{}

	resp, err := client.Get(req.URL.String())

	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	regex, _ := regexp.Compile(`href="/watch\?v=[A-Za-z0-9]*"`)

	rawLinkList := regex.FindAllString(string(body), 30)

	var linkList []string

	for i := 0; i < len(rawLinkList); i += 2 {
		link := strings.Split(rawLinkList[i], "\"")[1]

		linkList = append(linkList, "https://www.youtube.com"+link)
	}

	json, err := json.Marshal(linkList)

	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}
