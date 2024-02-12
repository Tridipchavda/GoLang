package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

// Struct For Accepting JSON from contact.js
type Contact struct {
	Id          int64 `json:"id"`
	ContactDets struct {
		Email string `json:"email"`
		Phone string `json:"phone"`
	}
}

// Struct For Accepting JSON from tech.js
type Tech struct {
	Id       int64 `json:"id"`
	TechDets []struct {
		Tech string  `json:"tech"`
		Exp  float32 `json:"exp"`
	}
}

// Struct For Accepting JSON from user.js
type User struct {
	Id      int64  `json:"id"`
	Name    string `json:"name"`
	Address struct {
		Area    string `json:"area"`
		Country string `json:"country"`
	}
}

// Struct For store the JSON in output/dat1.json
type Output struct {
	UserId  int64
	Name    string
	Email   string
	Phone   string
	Address struct {
		Area    string `json:"area"`
		Country string `json:"country"`
	}
	TechDetails []struct {
		Tech string  `json:"tech"`
		Exp  float32 `json:"exp"`
	}
}

// Sort Functions to sort the store JSON data based on Id
func SortUsers(users []User) {
	sort.Slice(users, func(i, j int) bool {
		return users[i].Id < users[j].Id
	})
}

func SortTechs(techs []Tech) {
	sort.Slice(techs, func(i, j int) bool {
		return techs[i].Id < techs[j].Id
	})
}

func SortContact(contact []Contact) {
	sort.Slice(contact, func(i, j int) bool {
		return contact[i].Id < contact[j].Id
	})
}

func main() {
	//fetch al the JSON file data in variables
	contactJsonData, err1 := os.ReadFile("json-data/contact.json")
	techJsonData, err2 := os.ReadFile("json-data/tech.json")
	userJsonData, err3 := os.ReadFile("json-data/user.json")

	// Check if there any error in one
	if err1 != nil || err2 != nil || err3 != nil {
		panic("File Reading Error")
	}

	// Make Arrays of Struct to store the JSON Array in Struct Array
	c := []Contact{}
	t := []Tech{}
	u := []User{}

	// Converting the JSON File data into JSON and store it in Structs
	err4 := json.Unmarshal(contactJsonData, &c)
	err5 := json.Unmarshal(techJsonData, &t)
	err6 := json.Unmarshal(userJsonData, &u)

	// Check if any Error in converting the file in Json data
	if err4 != nil || err5 != nil || err6 != nil {
		fmt.Println(err4, err5, err6)
		panic("Json Not Formatting")
	}
	// Sort all the Structs
	SortContact(c)
	SortTechs(t)
	SortUsers(u)

	// Output Struct to store the Join information on User , Contact, Tech on basis of Id
	var outs []Output = make([]Output, 0, 0)

	// Make the index Pointer of the structs starting from 0 index
	indexUser := 0
	indexTech := 0
	indexContact := 0

	// Map for Mapping Country Code to Country
	getContryCode := map[string]string{"UK": "+41-", "India": "+91-"}

	// Iterate the Struct Array Simultaeously and check for the Matching Id to JOIN
	for indexContact < len(c)-1 || indexUser < len(u)-1 || indexContact < len(c)-1 {
		// If all the IDs are same add the Combine Data to "out" Array
		if (u[indexUser].Id == t[indexTech].Id) && (u[indexUser].Id == c[indexContact].Id) {

			outs = append(outs, Output{UserId: c[indexContact].Id,
				Name: u[indexUser].Name, Email: c[indexContact].ContactDets.Email,
				Phone:   getContryCode[u[indexUser].Address.Country] + c[indexContact].ContactDets.Phone,
				Address: u[indexUser].Address, TechDetails: t[indexUser].TechDets})

			// Check if Incrementation is Valid or Out of range
			if indexContact < len(c)-1 {
				indexContact++
			}
			if indexTech < len(t)-1 {
				indexTech++
			}
			if indexUser < len(u)-1 {
				indexUser++
			}
			// Other than check which struct has misAligning Id and
			// Try to find the next Id in Other Structs If Match enter the data in out
		} else if u[indexUser].Id == t[indexTech].Id {
			if u[indexUser].Id > c[indexContact].Id {
				indexContact++
			} else {
				if indexUser < len(u)-1 {
					indexUser++
				}
				if indexTech < len(t)-1 {
					indexTech++
				}
			}
		} else if t[indexTech].Id == c[indexContact].Id {
			if t[indexUser].Id > u[indexUser].Id {
				indexContact++
			} else {
				if indexContact < len(c)-1 {
					indexContact++
				}
				if indexTech < len(t)-1 {
					indexTech++
				}
			}

		} else if c[indexContact].Id == u[indexUser].Id {
			if u[indexUser].Id > t[indexTech].Id {
				indexContact++
			} else {
				if indexContact < len(c)-1 {
					indexContact++
				}
				if indexUser < len(u)-1 {
					indexUser++
				}
			}

		}
		// fmt.Println(u[indexUser].Id, c[indexContact].Id, t[indexTech].Id)

	}

	// check for error in UnParsing the data
	b, err := json.Marshal(outs)
	if err != nil {
		panic(err)
	}
	// Write the data in file output/dat1.json with Read/Write/Execute permissions
	os.WriteFile("output/dat1.json", b, 0644)

}
