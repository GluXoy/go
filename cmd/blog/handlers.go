package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
)


type indexPage struct {
	FeaturedPosts []featuredPostData
	MostRecent    []mostRecentData
}

type featuredPostData struct {
	Title    				string `db:"title"`
	Subtitle           		string `db:"subtitle"`
	BackgroundImageModifier string `db:"image_url"`
	AuthorAvatar   			string `db:"author_url"`
	Author   				string `db:"author"`
	PublishDate 			string `db:"publish_date"`
}

type mostRecentData struct {
	PhotoImage 				string `db:"image_url"`
	Title 					string `db:"title"`
	Subtitle 				string `db:"subtitle"`
	Author  				string `db:"author"`
	AuthorAvatar 			string `db:"author_url"`
	PublishDate 			string `db:"publish_date"`
}


func index(db *sqlx.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		featuredPosts, err := featuredPosts(db)
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			log.Println(err)
			return 
		}

		recentPosts, err := mostRecent(db)
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			log.Println(err)
			return 
		}

		ts, err := template.ParseFiles("pages/index.html") 
		if err != nil {
			http.Error(w, "Internal Server Error", 500) 
			log.Println(err)
			return 
		}

		data := indexPage{
			FeaturedPosts: featuredPosts,
			MostRecent:    recentPosts,
		}

		err = ts.Execute(w, data) 
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			log.Println(err)
			return
		}

		log.Println("Request completed successfully")
	}
}

func post(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("pages/post.html") 
	if err != nil {
		http.Error(w, "Internal Server Error", 500) 
		log.Println(err.Error())                    
		return                                      
	}

	err = ts.Execute(w, nil) 
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		log.Println(err.Error())
		return
	}
}

func featuredPosts(db *sqlx.DB) ([]featuredPostData, error) {
	const query = `
		SELECT
			title,
			subtitle,
			author,
			author_url,
			publish_date,
			image_url
		FROM
			post
		WHERE featured = 1
	` 

	var posts []featuredPostData 

	err := db.Select(&posts, query) 
	if err != nil {                
		return nil, err
	}

	return posts, nil
}

func mostRecent(db *sqlx.DB) ([]mostRecentData, error) {
	const query = `
		SELECT
			title,
			subtitle,
			author,
			author_url,
			publish_date,
			image_url
		FROM
			post
		WHERE featured = 0
	` 

	var posts []mostRecentData 

	err := db.Select(&posts, query) 
	if err != nil {                 
		return nil, err
	}

	return posts, nil
}