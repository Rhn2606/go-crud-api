package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Todo struct {
	UserID    int    `json:"userId"`
	Id        int    `json:"Id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func getRequest() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("error")
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println("Error in getting response", res.Status)
		return

	}

	var todo Todo

	err = json.NewDecoder(res.Body).Decode(&todo)
	if err != nil {
		fmt.Println("error decoding", err)
		return
	}
	fmt.Println(todo)

}

func postRequest() {
	todo := Todo{
		UserID:    23,
		Title:     "Rohan",
		Completed: true,
	}
	//convert to json
	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("error marshalling", err)
		return
	}

	//convert json data to string
	jsonString := string(jsonData)
	jsonReader := strings.NewReader(jsonString)

	myURL := "https://jsonplaceholder.typicode.com/todos"

	res, err := http.Post(myURL, "application/json", jsonReader)
	if err != nil {
		fmt.Println("Error sending request :", err)
		return
	}
	defer res.Body.Close()

	data, _ := io.ReadAll(res.Body)
	fmt.Println("response :", string(data))

}

func UpdateRequest() {
	todo := Todo{
		UserID:    22,
		Title:     "Rohit",
		Completed: false,
	}

	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("error marshalling", err)
		return
	}
	jsonString := string(jsonData)

	jsonReader := strings.NewReader(jsonString)

	myURL := "https://jsonplaceholder.typicode.com/todos/1"

	req, err := http.NewRequest(http.MethodPut, myURL, jsonReader)
	if err != nil {
		fmt.Println("error updating data", err)
		return
	}
	req.Header.Set("content-type", "application/json")

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("error sending request", err)
		return
	}

	defer res.Body.Close()
	data, _ := io.ReadAll(res.Body)
	fmt.Println("response :", string(data))
	fmt.Println("Response status", res.Status)

}

func deleteRequest() {
	myURL := "https://jsonplaceholder.typicode.com/todos/1"
	req, err := http.NewRequest(http.MethodDelete, myURL, nil)
	if err != nil {
		fmt.Println("error updating data", err)
		return
	}
	req.Header.Set("content-type", "application/json")

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("error sending request", err)
		return
	}

	defer res.Body.Close()

	fmt.Println("Response status", res.Status)

}

func main() {
	fmt.Println("CRUD in Go")
	//getRequest()
	//postRequest()
	//UpdateRequest()
	deleteRequest()

}
