package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
)

var indexEmail = 28
var indexName = 0
var cookieSessionIntra *http.Cookie

func getAllStudents() [][]string {
	var students [][]string

	files, err := ioutil.ReadDir("csv")
	if err != nil {
		log.Fatal(err)
	}
	for i, fileInfo := range files {
		log.Printf("Starting file %s, (%d on %d)", fileInfo.Name(), i+1, len(files))
		file, err := os.Open("csv/" + fileInfo.Name())
		if err != nil {
			log.Fatal(err)
		}
		reader := csv.NewReader(file)
		// entries, err := csvutil.ReadFile("csv/" + fileInfo.Name())
		entries, err := reader.ReadAll()
		if err != nil {
			log.Fatal(err)
		}
		for _, value := range entries[1:] {
			log.Printf("Parsed %s %s (%s)\n", value[1], value[3], value[indexEmail])
		}
		students = append(students, entries[1:]...)
	}
	return students
}

func init() {
	// var err error

	rand.Seed(int64(os.Getpid()))
	// cookieSessionIntra, err = logToIntra()
	// log.Println("Cookie created with name " + cookieSessionIntra.Name + " and value " + cookieSessionIntra.Value)
	// if err != nil {
	// 	log.Fatalf("Cannot log to intra.epitech.eu, status code : %s\n", err)
	// }
}

func main() {
	students := getAllStudents()
	dispo := make([][]string, len(students))
	// fmt.Print(students)
	copy(dispo, students)

	yetsend := []string{}
	yetget := []string{}

	for id, value := range students {
		if stringInSlice(value[indexEmail], yetsend) {
			continue
		}
		randomIndex := id
		sender := value
		for (dispo[randomIndex][indexEmail] == students[id][indexEmail]) || stringInSlice(dispo[randomIndex][indexEmail], yetget) {
			randomIndex = rand.Int() % (len(dispo) - 1)
			if randomIndex < 0 {
				log.Fatal("end of program")
			}
			fmt.Print(randomIndex)
		}
		getter := dispo[randomIndex]
		// fmt.Printf("%s => %s\n", sender[indexEmail], getter[indexEmail])
		if !sendMail(sender, getter) {
			log.Fatal("error")
		}
		dispo = append(dispo[:randomIndex], dispo[randomIndex+1:]...)
	}
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
