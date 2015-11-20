package main

import (
    "fmt"
    "html"
    "log"
	"time"
	"strings"
    "github.com/julienschmidt/httprouter"
	"net/http"
	"io/ioutil"
	"github.com/coreos/etcd/Godeps/_workspace/src/golang.org/x/net/context"
    "github.com/coreos/etcd/client"
	"encoding/json"
)

type Entry struct {
	host string `json:"host"`
	url string `json:"url"`
}

func register(w http.ResponseWriter, r *http.Request, _ httprouter.Params, cfg client.Config) {
    c, err := client.New(cfg)
    if err != nil {
        log.Fatal(err)
    }
    kapi := client.NewKeysAPI(c)
	log.Print("Setting ", html.EscapeString(r.URL.Path))
	resp, err := kapi.Set(context.Background(), r.URL.Path, "bar", nil)
	if err != nil {
		log.Fatal(err)
	} else {
	// print common key info
		log.Printf("Set is done. Metadata is %q\n", resp)
	}
}

func route(w http.ResponseWriter, r *http.Request, _ httprouter.Params, cfg client.Config) {
    //fmt.Fprintf(w, "Requested: %q", html.EscapeString(r.URL.Path))
	prefix:="/services"
    url:=r.URL.Path[len(prefix):]
	base:="/directory"
	
	splitted_path := strings.Split(url, "/")
	str:=""
	for i := len(splitted_path)-1;i>=0;i-- {
		str=splitted_path[i] + "/" + str
    }
	
	c, err := client.New(cfg)
    if err != nil {
        log.Fatal(err)
    }
	if len (splitted_path) > 1 {
		kapi := client.NewKeysAPI(c)
		resp, err := kapi.Get(context.Background(),base+"/"+splitted_path[1], &client.GetOptions{Recursive: true})
		if err != nil {
             log.Fatal(err)
		}
		host:=""
		p:=""
		//log.Printf(resp.Node.Nodes[1].Value)
		for _, n := range resp.Node.Nodes {
			if n.Key == base+"/"+splitted_path[1] + "/host" { 
				host=n.Value
			} 
			if n.Key == base+"/"+splitted_path[1] + "/url" { 
				p=n.Value
			}
		}
		query:=""
		if len (splitted_path) > 2 {
		   query="/"+ splitted_path[2]
		}
		request:="http://"+host + p + query
		log.Printf(request)
		response, err := http.Get(request)
		if err != nil {
             log.Println(err)
			 http.Error(w, http.StatusText(404), 404)
		} else {
			b, err := ioutil.ReadAll(response.Body)
			if err != nil {
				log.Println(err)
				http.Error(w, http.StatusText(500), 500)
			} else {
				w.Write([]byte(b))
			}
		}
	}
}



func main() {
    cfg := client.Config{
        Endpoints:               []string{"http://127.0.0.1:2379"},
        Transport:               client.DefaultTransport,
        // set timeout per request to fail fast when the target endpoint is unavailable
        HeaderTimeoutPerRequest: time.Second,
    }
	router := httprouter.New()
    router.POST("/register", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {register(w, r, ps, cfg)})
	router.GET("/services/*service", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {route(w, r, ps, cfg)})
	router.POST("/services/*service", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {route(w, r, ps, cfg)})
    log.Fatal(http.ListenAndServe(":6667", router))
}
