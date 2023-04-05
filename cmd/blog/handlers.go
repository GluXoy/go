package main

import (
	"html/template"
	"log"
	"net/http"
)

type indexPage struct {
	FeaturedPosts []featuredPostData
	MostRecent    []mostRecentData
}

type featuredPostData struct {
	Title    				string
	Subtitle           		string
	BackgroundImageModifier string
	AuthorAvatar   			string
	Author   				string
	PublishDate 			string
}

type mostRecentData struct {
	PhotoImage 				string
	Title 					string
	Subtitle 				string
	Author  				string
	AuthorAvatar 			string
	PublishDate 			string
}

func index(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("pages/index.html") 
	if err != nil {
		http.Error(w, "Internal Server Error", 500) 
		log.Println(err.Error())                    
		return                                      
	}

	data := indexPage{
		FeaturedPosts: featuredPosts(),
		MostRecent:    mostRecent(),
	}

	err = ts.Execute(w, data) 
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		log.Println(err.Error())
		return
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

func featuredPosts() []featuredPostData {
	return []featuredPostData{
		{
			Title:    "The Road Ahead",
			Subtitle: "The road ahead might be paved - it might not be.",
			BackgroundImageModifier: "featured-post_background-road",
			AuthorAvatar: "static/img/matt_vogels.jpg",
			Author: "Mat Vogels",
			PublishDate: "September 25, 2015",
		},
		{
			Title:    "From Top Down",
			Subtitle: "Once a year, go someplace you’ve never been before.",
			BackgroundImageModifier: "featured-post_background-fromtop",
			AuthorAvatar: "static/img/william_wong.jpg",
			Author: "William Wong",
			PublishDate: "September 25, 2015",
		},
	}
}

func mostRecent() []mostRecentData {
	return []mostRecentData{
		{
			PhotoImage: "static/img/still_standing_tall.jpg",
			Title: "Still Standing Tall",
			Subtitle: "Life begins at the end of your comfort zone.",
			Author: "William Wong",
			AuthorAvatar: "static/img/william_wong.jpg",
			PublishDate: "9/25/2015",
		},
		{
			PhotoImage: "static/img/sunny_side_up.jpg",
			Title: "Sunny Side Up",
			Subtitle: "No place is ever as bad as they tell you it’s going to be.",
			Author: "Mat Vogels",
			AuthorAvatar: "static/img/matt_vogels.jpg",
			PublishDate: "9/25/2015",
		},
		{
			PhotoImage: "static/img/watter_falls.jpg",
			Title: "Water Falls",
			Subtitle: "We travel not to escape life, but for life not to escape us.",
			Author: "Mat Vogels",
			AuthorAvatar: "static/img/matt_vogels.jpg",
			PublishDate: "9/25/2015",
		},
		{
			PhotoImage: "static/img/throw_the_mist.jpg",
			Title: "Through the Mist",
			Subtitle: "Travel makes you see what a tiny place you occupy in the world.",
			Author: "William Wong",
			AuthorAvatar: "static/img/william_wong.jpg",
			PublishDate: "9/25/2015",
		},
		{
			PhotoImage: "static/img/awaken_early.jpg",
			Title: "Awaken Early",
			Subtitle: "Not all those who wander are lost.",
			Author: "Mat Vogels",
			AuthorAvatar: "static/img/matt_vogels.jpg",
			PublishDate: "9/25/2015",
		},
		{
			PhotoImage: "static/img/try_it_always.jpg",
			Title: "Try it Always",
			Subtitle: "The world is a book, and those who do not travel read only one page.",
			Author: "Mat Vogels",
			AuthorAvatar: "static/img/matt_vogels.jpg",
			PublishDate: "9/25/2015",
		},
	}
}