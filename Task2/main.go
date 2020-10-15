package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/julienschmidt/httprouter"
)

type Article struct {
	Id                 string    `json:"Id"`
	Title              string    `json:"Title"`
	Subtitle           string    `json:"Subtitle"`
	Content            string    `json:"Content"`
	Creation_Timestamp time.Time `json:Timestamp`
}

var Articles []Article

// var value []vars

// func handler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "Hello world!")
// }

func homePage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	fmt.Fprint(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func addArticles(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// uid := r.FormValue("uid")
	uid := ps.ByName("uid")
	fmt.Fprintf(w, "you are add user %s", uid)
}

func returnAllArticles(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	uid := ps.ByName("uid")
	for _, art := range Articles {
		if art.Id == uid {
			json.NewEncoder(w).Encode(art)
		}
	}
}
func convArticleList(a *Article) []string {
	p := make([]string, 0)
	result := make([]string, 0)
	tSplit := strings.Split(a.Title, " ")
	sSplit := strings.Split(a.Subtitle, " ")
	cSplit := strings.Split(a.Content, " ")
	p = append(p, tSplit...)
	p = append(p, sSplit...)
	p = append(p, cSplit...)
	for _, e := range p {
		result = append(result, strings.ToLower(e))
	}
	return result
}
func returnsearchArticle(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["q"]

	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'key' is missing")
		return
	}

	key := keys[0]
	var key_lower string = strings.ToLower(string(key))

	fmt.Println("Url Param 'key' is: " + string(key_lower))
	for _, article := range Articles {
		a := &article
		d := convArticleList(a)
		fmt.Println(d)
		temp := make([]string, 0)
		for i := range d {
			if d[i] == key_lower {
				temp = append(temp, d[i])
			}
		}
		if len(temp) > 0 {
			json.NewEncoder(w).Encode(article)
		}
	}
}

func handleRequests() {

	router := httprouter.New()
	router.GET("/", homePage)
	// http.HandleFunc("/", homePage)
	router.POST("/articles", addArticles)
	router.GET("/articles", returnAllArticles)
	http.HandleFunc("/articles/search", returnsearchArticle)
	// router.GET("/articles/search?q=<s>", returnsearchArticle)
	router.GET("/articles/:uid", returnSingleArticle)

	log.Fatal(http.ListenAndServe(":8080", router))
}
func main() {

	Articles = []Article{
		Article{Id: "a", Title: "Hello", Subtitle: "Article Description", Content: "Article Content", Creation_Timestamp: time.Now()},
		Article{Id: "2", Title: "Hello 2", Subtitle: "Article Description", Content: "Article Content", Creation_Timestamp: time.Now()},
	}

	handleRequests()

}
