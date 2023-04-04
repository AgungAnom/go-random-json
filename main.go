package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main(){
	var water int
	var wind int
	var id int

	rand.Seed(time.Now().UnixNano())



	water = rand.Intn(100 - 1) + 1
	wind = rand.Intn(100 - 1) + 1
	id = rand.Intn(100 - 1) + 1
	fmt.Println("----------------- Interval : 1 -----------------")
	random(wind, water, id)


	// Start interval
	i := 2
	for range time.Tick(time.Second * 15) {
		fmt.Println("----------------- Interval :",i,"-----------------")
		i++
		water = rand.Intn(100 - 1) + 1
		wind = rand.Intn(100 - 1) + 1
		id = rand.Intn(100 - 1) + 1
		random(wind, water, id)
    }
}


func random(wind, water, id int){
	

	data := map[string]int{
		"userID" : id,
	}
	requestJson, err := json.Marshal(data)
	client := &http.Client{}
	if err != nil {
		log.Fatalln(err)
		
	}
	fmt.Println(string(requestJson))
	req, err := http.NewRequest("POST", "https://jsonplaceholder.typicode.com/posts/" , bytes.NewBuffer(requestJson))
	req.Header.Set("Content-type", "Application/json")
	if err != nil {
		log.Fatalln(err)
	}
	res, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(body))
	fmt.Println("{")
	fmt.Print("	water: ", water,", \n")
	fmt.Println("	wind: ",wind)
	fmt.Println("}")
	if water < 5 {
		fmt.Println("status water : aman")	
	} else if water >= 6 && water <= 8{
		fmt.Println("status water : siaga")	
	} else if water > 8 {
		fmt.Println("status water : bahaya")	
	}
	if wind < 6 {
		fmt.Println("status wind : aman")	
	} else if wind >= 7 && wind <= 15{
		fmt.Println("status wind : siaga")	
	} else if wind > 15 {
		fmt.Println("status wind : bahaya")	
	}
}



