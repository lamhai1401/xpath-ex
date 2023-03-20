package main

import (
	"fmt"
	"os"

	"github.com/antchfx/jsonquery"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	jsonFile, err := os.Open("large.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened jsonFile as a byte array.
	doc, err := jsonquery.Parse(jsonFile)
	if err != nil {
		panic(err.Error())
	}

	// "/user[@id='Sg45wR6T']/projects"
	// query := "//payload/issue[id=9343331]/user"
	query := "//payload/issue[id=53222540]/user[login]"

	// query := "//ev[contains(@plugin, 'chatSurvey')]"
	lst, err := jsonquery.QueryAll(doc, query)
	if err != nil {
		panic(err.Error())
	}

	// spew.Dump(lst.Data, lst.Value())

	// spew.Dump(lst.Value())
	for _, data := range lst {
		spew.Dump("Value: ", data.Value())
		// spew.Dump("==Data:", data.Data)
		// spew.Dump("==Value:", data.Value())
	}
}
