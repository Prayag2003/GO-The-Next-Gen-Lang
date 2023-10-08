package main

import (
	"encoding/json"
	"fmt"
)

// struct data with Aliases
type SIH_DATA struct {
	TeamName              string `json:"sih_team_name"`
	NumberOfMembers       int    `json:"number"`
	Lead                  string `json:"team_lead"`
	Idea                  string
	Members               []string `json:"members,omitempty"` // omitempty removes nil data
	Registration_Password string   `json:"-"`
}

func main() {
	fmt.Println("More practice on JSON")

	// convert data to JSON --> JSON Encoding
	EncodeJSON()
	DecodeJSON()
}

func EncodeJSON() {
	teams := []SIH_DATA{
		{"Tools Of Titans", 8, "Prayag Bhatt", "AI Based Training System", []string{"Prayag", "Denil", "Aum", "Rudraksh", "Bhavya", "Kartik", "Karan", "Istuti"}, "123456"},

		{"Byte Hogs", 4, "ABC", "Smart Education", nil, "1234567"},
	}

	// packaging this data as JSON Data
	// Pass the interface inside
	finalJson, err := json.MarshalIndent(teams, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)
}

func DecodeJSON() {

	// NOTE: Json Data from the web in form of bytes
	jsonDataFromWeb := []byte(`
		{
                "sih_team_name": "Tools Of Titans",
                "number": 8,
                "team_lead": "Prayag Bhatt",
                "Idea": "AI Based Training System",
                "members": [
                        "Prayag",
                        "Denil",
                        "Aum",
                        "Rudraksh",
                        "Bhavya",
                        "Kartik",
                        "Karan",
                        "Istuti"
                ]
        }
	`)

	checkIfJsonIsValid := json.Valid(jsonDataFromWeb)
	var tools_to_titans SIH_DATA

	if checkIfJsonIsValid {
		fmt.Println("Valid Data")
		json.Unmarshal(jsonDataFromWeb, &tools_to_titans)
		fmt.Printf("%#v\n", tools_to_titans)
	} else {
		fmt.Println("JSON IS INVALID")
	}

	// NOTE:  In form of key-value pairs
	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and Type is %T\n", k, v, v)
	}
}
