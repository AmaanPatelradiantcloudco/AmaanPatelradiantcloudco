package models

import (
	"encoding/json"
	"fmt"
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
	}

	// err := json.NewDecoder(js).Decode(&testAuthor) // error

	for _, value := range TestValues {

		var testAuthor Authorapi // A new Author type variable

		err := json.Unmarshal([]byte(value), &testAuthor)
		if err != nil {
			t.Errorf("TEST (UNMARSHALLING)  FAILED : %v", err.Error())
		}

		// Testing validations :
		err = testAuthor.Validate()
		if err != nil {
			t.Errorf("TEST (VALIDATION) FAILED : %v", err.Error())
			fmt.Print("Test case failed")
		}
	}
}
