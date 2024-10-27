package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type BlogSummary struct {
	Id             int    `json:"id"`
	Date           string `json:"date"`
	Summary        string `json:"summary"`
	FirstParagraph string `json:"first_paragraph"`
}

var randomData = []BlogSummary{
	{Id: 1, Date: "Jan 01, 2020", Summary: "<h1>Blog Summary 1</h1>", FirstParagraph: "<p>Writing code for a computer is hard enough. You take something big and fuzzy, some large vague business outcome you want to achive. Then you break it down recursively and think about all the cases until you have clear logical statements a computer can follow. Computers are very good at following logical statements</p>"},
	{Id: 2, Date: "Jan 02, 2020", Summary: "<h1>Blog Summary 2</h1>", FirstParagraph: "<p>Writing code for a computer is hard enough. You take something big and fuzzy, some large vague business outcome you want to achive. Then you break it down recursively and think about all the cases until you have clear logical statements a computer can follow. Computers are very good at following logical statements</p>"},
	{Id: 3, Date: "Jan 03, 2020", Summary: "<h1>Blog Summary 3</h1>", FirstParagraph: "<p>Writing code for a computer is hard enough. You take something big and fuzzy, some large vague business outcome you want to achive. Then you break it down recursively and think about all the cases until you have clear logical statements a computer can follow. Computers are very good at following logical statements</p>"},
}

func main() {
	port := "8080"

	jsonOutput, err := json.Marshal(randomData)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(jsonOutput))

	http.HandleFunc("/blogs", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*") // Allow all origins
		w.Write(jsonOutput)
	})

	// Start the server
	fmt.Printf("Running on port %s\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
