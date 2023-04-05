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

	//First interval
	rand.Seed(time.Now().UnixNano())
	water = rand.Intn(100 - 1) + 1
	wind = rand.Intn(100 - 1) + 1
	id = rand.Intn(100 - 1) + 1
	fmt.Println("-----------------",time.Now(),"-----------------")
	random(wind, water, id)


	// Start interval 15 sec
	for range time.Tick(time.Second * 15) {
		fmt.Println("-----------------",time.Now(),"-----------------")
		water = rand.Intn(100 - 1) + 1
		wind = rand.Intn(100 - 1) + 1
		id = rand.Intn(100 - 1) + 1
		random(wind, water, id)
    }
}


func random(wind, water, id int){
	
	// POST ID
	data := map[string]int{
		"userID" : id,
	}
	requestJson, err := json.Marshal(data)
	client := &http.Client{}
	if err != nil {
		log.Fatalln(err)
		
	}
	req, err := http.NewRequest("POST", "https://jsonplaceholder.typicode.com/posts/" , bytes.NewBuffer(requestJson))
	req.Header.Set("Content-type", "Application/json")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(req.Body)

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



	// Data Sensor
	dataSensor := map[string]int{
		"water" : water,
		"wind"	: wind,
	}

	jsonString, err := json.Marshal(dataSensor)
	src := []byte(jsonString)
	dst := &bytes.Buffer{}
	if err := json.Indent(dst, src, "", "  "); err != nil {
		panic(err)
	}
	fmt.Println(dst.String())

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



