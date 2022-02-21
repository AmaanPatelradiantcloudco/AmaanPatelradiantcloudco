package main

import (
	"encoding/json"
	"testing"
)

// func TestExample(t *testing.T) {
// 	if models.Example(4) != 16 {
// 		t.Error("TEST CASE Failed. Expected result was 16.")
// 	}

// 	if models.Example(5) != 25 {
// 		t.Error("TEST CASE Failed. Expected result was 25.")
// 	}
// }

func TestAuthorValidations(t *testing.T) { // A test function name starts with the word 'Test' and also takes in a parameter of type '*testing.T'.

	TestValues := []string{
		`{"id":1, "name":"hello world","description":"2345"}`,
		`{"id": 2, name": "amaan patel","description":"2345"}`,
	}

	// err := json.NewDecoder(js).Decode(&testAuthor) // error

	for _, value := range TestValues {

		var testAuthor Author // A new Author type variable

		err := json.Unmarshal([]byte(value), &testAuthor)
		if err != nil {
			t.Errorf("TEST (UNMARSHALLING)  FAILED : %v", err.Error())
		}

		// Testing validations :
		err = testAuthor.Validate()
		if err != nil {
			t.Errorf("TEST (VALIDATION) FAILED : %v", err.Error())
		}
	}
}

/* Testing validations for wrong values :
//func TestNegativeAuthorValidations(t *testing.T) {


	}
	// var testAuthor models.Author // If defined outside then will not have the next json value unmarshalled as a new value rather the next will be actually updating the prevuius one.
	// If any field is not in the next json string and is there in the previous one then the current one will too have the field as defining outside does not creates a new Author type variable
	//  rather the json will be unmarshalled into the already created one thus, effectively updating the previous one.

	for _, value := range TestValues {
		var testAuthor Author
		if err := json.Unmarshal([]byte(value), &testAuthor); err != nil {
			t.Errorf("FAILED TO UNMARSHAL %q TO JSON : %v", value, err.Error())
		}

		fmt.Println(testAuthor)
		if err := testAuthor.Validate(); err == nil {
			t.Errorf("Expected a validation error, none received.")
		}
	}
}*/

func TestBookValidations(t *testing.T) {

	// The following will give an 'Unmarshallling' error :
	/*
		TestValues := []string{
			`{"id": "1","title": "Hello! Go", "description": "A book on Golang."}`,
			`{"id": 1, "title": "Hello! Go", "description": "A book on Golang.", "author": "{"id": 1, "name": "Harsimran"}"}`,
		}
	*/
	/*
		TestValues := []string{
			// Validations will be performed on the fields of the 'author' (as specified for the 'Author' type) too
			//  as it is itself a field of a 'Book' type and set to 'required'.
			`{"id": 222222222222222222, "title": "Hello! JS", "author": {"id": 1, "name": "H"}}`,
		}
	*/

	TestValues := []string{
		`{"id": 1, "name": "Amaan","description":"hello world"}}`,
		// `{"id": 1, "title": "Hello! Ruby                                                                                             ",  "author": {"id": 1, "name": "Harsimran"}"}`,
		// 	`{"id": 1, "title": "Hello! Ruby..............................................................................................",  "author": {"id": 1, "name": "Harsimran"}}`,
		// `{"id": 2222222222222222222222222222, "title": "Hello! Go", "description": "A book on Golang.", "author": {"id": 1, "name": "Harsimran"}}`,
		`{"id": 2, "name": "Fahad", "description": "hello world"}`,
	}

	for _, value := range TestValues {

		var testBook Book // A new Book type variable

		if err := json.Unmarshal([]byte(value), &testBook); err != nil {
			t.Errorf("FAILED TO UNMARSHAL %q TO JSON : %v", value, err.Error())
		}

		if err := testBook.Validate(); err != nil {
			t.Errorf("TEST CASE (VALIDATION) FAILED : %v", err.Error())
		}
	}
}

func TestUserValidations(t *testing.T) {

	TestValues := []string{
		`{"username": "hello world", "email": "Moin qureshi", "password": "$2a$10$xYADUPksQV3kPzis5I0.ruBlgzmxaTx9uDJSHkxXfNB..."}`,
	}

	for i, value := range TestValues {

		var testUser User // A new User type variable

		if err := json.Unmarshal([]byte(value), &testUser); err != nil {
			t.Errorf("FAILED TO UNMARSHAL %q TO JSON : %v", value, err.Error())

		}

		if err := testUser.Validate(); err != nil {
			t.Errorf("(%v) => TEST CASE (VALIDATION) FAILED : %v", i, err.Error()) // index starts with 0 in a slice.
		}
	}
}
