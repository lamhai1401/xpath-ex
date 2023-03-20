package main

import (
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestJsonPath(t *testing.T) {
	xpath := NewJSONPath()
	// read large json file
	jsonFile, _ := os.Open("large.json")
	defer jsonFile.Close()
	// add new json file to xpath
	xpath.addNewDoc(jsonFile)
	// test only with query function
	Convey("Test xpath with input query", t, func() {
		Convey("Retrieve the fragment by specifying the location and attributes.", func() {
			// move to payload key and find in issue object which have id = 53222540, print login field in user object
			query := "//payload/issue[id=53222540]/user/login"
			expectValue := "No-CQRT"
			data, err := xpath.QueryAll(query)
			Convey("Result of query above", func() {
				So(err, ShouldBeNil)
				So(data[0], ShouldEqual, expectValue)
			})
		})
		Convey("Retrieve the fragment by XPath functions.", func() {
			// find all object has field number and it value is 20344
			query := "//*/number[contains(.,'20344')]"
			expectValue := 20344
			data, err := xpath.QueryAll(query)
			Convey("Result of query above", func() {
				So(err, ShouldBeNil)
				So(data[0], ShouldEqual, expectValue)
			})
		})
		Convey("Retrieve the value.", func() {
			// find all object that has object issue with field id = 53222540
			// after that it has user object, user object must contain login field
			query := "//payload/issue[id=53222540]/user[login]"
			expectedID := 9343331
			data, err := xpath.QueryAll(query)
			value, isData := data[0].(map[string]interface{})
			So(isData, ShouldBeTrue)
			Convey("Result of query above", func() {
				So(err, ShouldBeNil)
				So(value["id"], ShouldEqual, expectedID)
			})
		})
	})
}
