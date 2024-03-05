package main

import (
	"fmt"
	_ "math"
	"reflect"
)

func main() {

	course := map[string]interface{}{
		"Name":  "React bootcamp",
		"Price": 399,
		"Platform": map[string]interface{}{
			"name":        "Udemy",
			"rating":      4.2,
			"coupon-code": "xyz45",
			"courses": map[string]interface{}{
				"JS":     "JavaScript Bootcamp",
				"sql":    "SQL Bootcamp",
				"python": "python bootcamp",
				"name":   "udemy",
			},
		},
		"roll-no": 72,
		"Branch":  "CSE",
	}
	set := setKeyValue("Name", "React Guide Course", course) // call to setKeyValue function

	if set {
		fmt.Println("Name Changed")
	} else {
		fmt.Println("Not Found")

	}
	fmt.Println(course)

	unset := RemoveKey("coupon-code", course) // call to RemoveKey function

	if unset {
		fmt.Println("Deleted")
	} else {
		fmt.Println("Not Found")
	}
	fmt.Println(unset)

	// pointer to a Courses struct
	var coursePtr *Courses = &Courses{}
	// Populate the Courses struct fields using the course
	PopulateStruct(course, coursePtr)

	//  populated Courses struct
	fmt.Printf("%+v\n", *coursePtr)

}

// setKeyValue function

func setKeyValue(key string, value interface{}, source map[string]interface{}) bool {

	var res bool
	if _, val := source[key]; val {
		source[key] = value // if the key exists update the value and return true
		return true
	} else {
		for _, val1 := range source {

			if reflect.TypeOf(val1).Kind() == reflect.Map {
				res = res || setKeyValue(key, value, val1.(map[string]interface{}))

			}

		}
	}
	return res
}

// Removekey function

func RemoveKey(key string, source map[string]interface{}) bool {
	var res bool
	if _, val := source[key]; val {
		delete(source, key) // if key exists delete it
		return true
	} else { // else iterate over the map to check for submaps
		for _, val1 := range source {

			if reflect.TypeOf(val1).Kind() == reflect.Map {
				res = res || RemoveKey(key, val1.(map[string]interface{})) // recursive call to perform deletion on the sub map

			}
		}
	}
	return res
}

//PopulateStruct function

type Courses struct {
	Name     string
	Price    int
	Platform Address
}

type Address struct {
	City  string
	State string
}

func PopulateStruct(course map[string]interface{}, result interface{}) {
	resultValue := reflect.ValueOf(result).Elem()

	for key, value := range course {
		field := resultValue.FieldByName(key)

		// Check if the field is valid (exists in the struct)
		if field.IsValid() {
			// If the field is a struct itself, recursively populate its fields
			if field.Kind() == reflect.Struct {
				if nestedMap, val := value.(map[string]interface{}); val {
					nestedStruct := reflect.New(field.Type()).Interface()
					// Recursively populate the nested struct fields
					PopulateStruct(nestedMap, nestedStruct)
					field.Set(reflect.ValueOf(nestedStruct).Elem())
				}
			} else {
				field.Set(reflect.ValueOf(value))
			}
		}
	}
}
