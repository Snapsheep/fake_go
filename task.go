package task

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// User :: desc
type User struct {
	ID   int    `json:"ID`
	Name string `json:"Name"`
	Age  int    `json:"Age`
}

// AllUsers :: desc
var AllUsers []User
var pathFile string = "./users.json"

// AddUsers :: desc
func AddUsers(name string, age int) {
	// read file
	var data []User
	file, _ := ioutil.ReadFile(pathFile)

	json.Unmarshal(file, &data)

	if len(data) > 0 {
		for _, v := range data {
			AllUsers = append(AllUsers, User{ID: v.ID, Name: v.Name, Age: v.Age})
		}
		AllUsers = append(AllUsers, User{ID: len(data) + 1, Name: name, Age: age})
		b, err := json.MarshalIndent(AllUsers, "", "  ")

		err = os.Remove(pathFile)

		if err != nil {
			fmt.Println(err)
			return
		}

		f, err := os.OpenFile(pathFile,
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println(err)
		}
		defer f.Close()
		if _, err := f.WriteString(string(b)); err != nil {
			fmt.Println(err)
		}
	} else {
		f, err := os.Create(pathFile)

		if err != nil {
			fmt.Println(err)
		}
		defer f.Close()

		err = retrieveUser(name, age)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

// ListUsers :: desc
func ListUsers() {
	var data []User
	file, err := ioutil.ReadFile(pathFile)
	err = json.Unmarshal(file, &data)
	if err != nil {
		fmt.Println("Json is empty!")
	}
	fmt.Print(string(file))
}

func retrieveUser(name string, age int) error {
	AllUsers = append(AllUsers, User{ID: 1, Name: name, Age: age})
	b, err := json.MarshalIndent(AllUsers, "", "  ")

	if err != nil {
		fmt.Println(err)
	} else {
		f, err := os.OpenFile(pathFile,
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println(err)
		}
		defer f.Close()

		if _, err := f.WriteString(string(b)); err != nil {
			fmt.Println(err)
		}
	}
	return err
}
