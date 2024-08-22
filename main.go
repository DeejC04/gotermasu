// CLI STRUCTURE
// "asucli search -subject cse -catalognbr 205"


package main

 import (
	 // "os"
	 "encoding/json"
	 "flag"
	 "fmt"
	 "io"
	 "log"
	 "net/http"
	 "os"
	 "net/url"
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
	// adding flags for each search category/filter
	// only one of these four may have a specified value by the user (figure out how to do that)
	// searchCourses := searchCmd.String("course", "", "search course")
	searchCourses := flag.NewFlagSet("courses", flag.ExitOnError)
	subject := searchCourses.String("subject", "", "subject")
	catalogNbr := searchCourses.String("catalogNbr", "", "catalogNbr")
	apiUrl, _ := url.Parse("https://eadvs-cscc-catalog-api.apps.asu.edu/catalog-microservices/api/v1")

	
	queryString := apiUrl.Query()
	// depending on what the user states in the cli, more params will be added to
	// queryString (don't use .Set, use .Add)
	// queryString := url.Values{}
	// queryString.Set("refine", "Y")

	if len(os.Args) < 2 {
        fmt.Println("expected 'search' subcommand")
        os.Exit(1)
    }

	switch os.Args[1] {
	case "courses":
		searchCourses.Parse(os.Args[2:])

		//debug
		if isFlagPassed("subject", searchCourses) {
			queryString.Add("subject", *subject)
		}

		if isFlagPassed("catalogNbr", searchCourses) {
			queryString.Add("catalogNbr", *catalogNbr)
		}


		// depending on what search param (courses/classes/subjects/terms) is used
		// the nested commands will be different etc
		
		//debug
		fmt.Println(queryString)
	default:
		fmt.Println("unknown command")
		os.Exit(1)
	}
	
	apiUrl.RawQuery = queryString.Encode()

	fmt.Println(apiUrl.String())
	// fmt.Println(isFlagPassed("subject", searchCourses))
	// will be used to make requests with headers (doesn't seem like
	// default http.Get() command can use headers
	client := &http.Client{}
	// queryParams := "/search/" + searchCourses.Name() + queryString.Encode()

	// creating the request itself, afaik it doesn't actually send it yet
	// since I'm literally just creating a new request
	req, err := http.NewRequest("GET", (apiUrl.String()), nil)
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

func isFlagPassed(name string, flagSetToVisit *flag.FlagSet) bool {
    found := false
    (*flag.FlagSet).Visit(flagSetToVisit, func(f *flag.Flag) {
        if f.Name == name {
            found = true
        }
    })
    return found
}