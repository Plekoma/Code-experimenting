package CRS

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

// idk apparently this is how i like make script wide vars or smth idk this lang is weird, also err is here bc like it dont work without
var base string
var route string
var envPath string
// Da map for the id to thing json, highway is any bc it has nested stuffs and run and regular is string bc it dont
var highway map[string]any
var regular map[string]string 
var run map[string]string
var err error

func Init() {
	establishBase()
}

// Establishes the lowest file in the app
func establishBase() {
	// Get da base and get paths
	base, err = os.Getwd()
	route = filepath.Join(base, "CRS")
	envPath = filepath.Join(base, ".env")

	// checks for errors
	if err != nil {
		log.Printf("Establishing base failed (somehow): %v", err)
	}
	
	// set the base as a terminal environment variable (ALB is App Lowest Base)
	env := map[string]string{
		"ALB": base,
		"CRS": route,
	}

	// write to da env file ig dis how you do it
	err = godotenv.Write(env, envPath)

	if err != nil{
		log.Fatalf("Error writing to env: %v", err)
	}
}

// Gets da json
func load() {
	// open da like jsons
	data, err := os.ReadFile(filepath.Join(route, "highway.json"))
	if err != nil{
		log.Fatalf("Error retrieving json file, cannot continue. %v", err)
	}

	// make a temp struct to like mrrow idk mrrow mrrow mrrow
	var mrrow struct {
		ID map[string]string `json:"id"`
	}

	if err := json.Unmarshal(data, &mrrow); err == nil {
		regular = mrrow.ID
	}

	json.Unmarshal(data, &highway)


	// get json with os and like do unmarshal whatev that is
	json.Unmarshal(data, &highway)

	// open da like jsons
	data, err = os.ReadFile(filepath.Join(route, "run.json"))
	if err != nil{
		log.Fatalf("Error retrieving json file, cannot continue. %v", err)
	}
	// get json with os and like do unmarshal whatev that is
	json.Unmarshal(data, &run)

	// js confirmation it succeeded
	fmt.Println(highway)
	fmt.Println(run)
	fmt.Println(regular)
}

// Has its own function bc it was originally like actually complicated until i realized i was stupid
func detect(path string) string {
	Type := filepath.Ext(path)
	fmt.Println(Type)
	return Type
}

func Route(file string){
	load()

	// kinda like parse the data inputed into file

}