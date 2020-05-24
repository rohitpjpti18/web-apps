package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Article struct
type Article struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

// Articles let's declare a global Articles array
// that we can then populate in our main function
// to simulate a database
var Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: homePage")
	json.NewEncoder(w).Encode(Articles)
}

func articleByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	//fmt.Fprintf(w, "key: "+key)

	notFound := "not found! "

	// loop over all articles
	for _, article := range Articles {
		if article.ID == key {
			json.NewEncoder(w).Encode(article)
			return
		}
	}
	json.NewEncoder(w).Encode(notFound)
	return
}

func createNewArticle(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var article Article
	//fmt.Fprintf(w, "%+v", string(reqBody))

	json.Unmarshal(reqBody, &article)
	Articles = append(Articles, article)

	json.NewEncoder(w).Encode(article)
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	for index, article := range Articles {
		if article.ID == id {
			Articles = append(Articles[:index], Articles[index+1:]...)
		}
	}
	json.NewEncoder(w).Encode("successfull")
}

func requestsHandler() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/articles", allArticles)
	myRouter.HandleFunc("/article", createNewArticle).Methods("POST")

	myRouter.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")
	myRouter.HandleFunc("/article/{id}", articleByID)
	fmt.Println("attempting to start server... at port no: 8080")
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")
	Articles = []Article{
		Article{ID: "1", Title: "Hello-harry", Desc: "Article Description", Content: "Article Content"},
		Article{ID: "2", Title: "Hello-ronald", Desc: "Article Description", Content: "Article Content"},
		Article{ID: "3", Title: "Hello-hermoine", Desc: "Article Description", Content: "Article Content"},
		Article{ID: "4", Title: "Hello-ginny", Desc: "Article Description", Content: "Article Content"},
	}
	requestsHandler()
}
