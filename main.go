	package main

	import (
		// "encoding/json"
		"encoding/json"
		"fmt"
		"io"
		"log"
		"net/http"
		// "flag"
	)

	type Course []struct {
		Subject          string `json:"SUBJECT"`
		Catalognbr       string `json:"CATALOGNBR"`
		Coursetitlelong  string `json:"COURSETITLELONG"`
		Componentprimary string `json:"COMPONENTPRIMARY"`
		Allowmultenroll  string `json:"ALLOWMULTENROLL"`
		Crserepeatable   string `json:"CRSEREPEATABLE"`
		Gradingbasis     string `json:"GRADINGBASIS"`
		Acadorg          string `json:"ACADORG"`
		Collegemap       []struct {
			Department string `json:"DEPARTMENT"`
			Info       struct {
				Acadorg       string `json:"ACADORG"`
				Asucollegeurl string `json:"ASUCOLLEGEURL"`
				Descrformal   string `json:"DESCRFORMAL"`
				Enrollreq     string `json:"ENROLLREQ"`
				Acadgroup     string `json:"ACADGROUP"`
			} `json:"INFO"`
		} `json:"COLLEGEMAP"`
		Descrlong         string `json:"DESCRLONG"`
		Componentdescr    string `json:"COMPONENTDESCR"`
		Gradingbasisdescr string `json:"GRADINGBASISDESCR"`
		Descr4            string `json:"DESCR4"`
		Crseid            string `json:"CRSEID"`
		Hours             string `json:"HOURS"`
		Unitsmaximum      string `json:"UNITSMAXIMUM"`
		Unitsminimum      string `json:"UNITSMINIMUM"`
		Subjectdescr      string `json:"SUBJECTDESCR"`
		Topicslist        []any  `json:"TOPICSLIST"`
		Gsgold            string `json:"GSGOLD"`
		Gsmaroon          string `json:"GSMAROON"`
	}

	func main() {
		// will be used to make requests with headers (doesn't seem like
		// default http.Get() command can use headers
		client := &http.Client{}
		url := "https://eadvs-cscc-catalog-api.apps.asu.edu/catalog-microservices/api/v1"

		// creating the request itself, afaik it doesn't actually send it yet
		// since I'm literally just creating a new request
		req, err := http.NewRequest("GET", (url + "/search/courses?refine=Y&catalogNbr=205&course_id=104182&subject=CSE&term=2247"), nil)
		if err != nil {
			log.Fatal(err)
		}

		// adding the header necessary to send the request
		req.Header.Add("Authorization", "Bearer null")

		// finally, sending the actual request via the client with the headers included
		resp, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		var course Course
		err = json.Unmarshal(body, &course)
		if err != nil {
			log.Fatal(err)
		}

		// by default, response is a byte slice (research it), so it must be converted
		// to a string
		fmt.Println(course[0].Subject)

	}
