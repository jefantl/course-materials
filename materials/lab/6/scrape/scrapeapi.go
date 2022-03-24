package scrape

// scrapeapi.go HAS TEN TODOS - TODO_5-TODO_14 and an OPTIONAL "ADVANCED" ASK

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"scrape/logging"

	"github.com/gorilla/mux"
)

//==========================================================================\\

var fileHitCount int

// Helper function walk function, modfied from Chap 7 BHG to enable passing in of
// additional parameter http responsewriter; also appends items to global Files and
// if responsewriter is passed, outputs to http

func walkFn(w http.ResponseWriter) filepath.WalkFunc {
	fileHitCount = 0
	return func(path string, f os.FileInfo, err error) error {
		w.Header().Set("Content-Type", "application/json")

		for _, r := range regexes {
			handleRegExp(w, path, r)
		}
		return nil
	}

}

// DONE
//TODO_7: One of the options for the API is a query command
//TODO_7: Create a walkFn2 function based on the walkFn function,
//TODO_7: Instead of using the regexes array, define a single regex
//TODO_7: Hint look at the logic in scrape.go to see how to do that;
//TODO_7: You won't have to itterate through the regexes for loop in this func!

func walkFn2(w http.ResponseWriter, query string) filepath.WalkFunc {
	fileHitCount = 0
	return func(path string, f os.FileInfo, err error) error {
		w.Header().Set("Content-Type", "application/json")

		r := regexp.MustCompile(query)
		handleRegExp(w, path, r)

		return nil
	}
}

func handleRegExp(w http.ResponseWriter, path string, r *regexp.Regexp) {
	if r != nil && r.MatchString(path) {
		var tfile FileInfo
		dir, filename := filepath.Split(path)
		tfile.Filename = string(filename)
		tfile.Location = string(dir)

		// Done
		//TODO_5: As it currently stands the same file can be added to the array more than once
		//TODO_5: Prevent this from happening by checking if the file AND location already exist as a single record
		for _, f := range Files {
			if f.Filename == tfile.Filename && f.Location == tfile.Location {
				return
			}
		}

		Files = append(Files, tfile)
		fileHitCount++

		if w != nil && len(Files) > 0 {

			// Done
			//TODO_6: The current key value is the LEN of Files (this terrible);
			//TODO_6: Create some variable to track how many files have been added
			w.Write([]byte(`"` + fmt.Sprint(fileHitCount) + `":  `))
			json.NewEncoder(w).Encode(tfile)
			w.Write([]byte(`,`))

		}

		logging.IfLevel(fmt.Sprintf("[+] HIT: %s\n", path), 2)

	}
}

//==========================================================================\\

func APISTATUS(w http.ResponseWriter, r *http.Request) {

	logging.IfLevel(fmt.Sprintf("Entering %s end point", r.URL.Path), 1)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{ "status" : "API is up and running ",`))
	var regexstrings []string

	for _, regex := range regexes {
		regexstrings = append(regexstrings, regex.String())
	}

	w.Write([]byte(` "regexs" :`))
	json.NewEncoder(w).Encode(regexstrings)
	w.Write([]byte(`}`))
	logging.IfLevel(fmt.Sprint(regexes), 1)

}

func MainPage(w http.ResponseWriter, r *http.Request) {
	logging.IfLevel(fmt.Sprintf("Entering %s end point", r.URL.Path), 1)
	w.Header().Set("Content-Type", "text/html")

	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, `<div>
    <h1>This is a file scraper API</h1>
    </div>
    <p>Here are the URLs you can use to gather file information from this system:</p>
    <div>
    <ul>
    <li>/api-status</li>
    <li>/indexer</li>
    <li>/search</li>
    <li>/addsearch/{regex}</li>
    <li>/clear</li>
    <li>/reset</li>
    </ul>
    </div>
    <p>&nbsp;</p>
    <p>&nbsp;</p>`)
}

func FindFile(w http.ResponseWriter, r *http.Request) {
	logging.IfLevel(fmt.Sprintf("Entering %s end point", r.URL.Path), 1)
	q, ok := r.URL.Query()["q"]

	w.WriteHeader(http.StatusOK)
	if ok && len(q[0]) > 0 {
		logging.IfLevel(fmt.Sprintf("Entering search with query=%s", q[0]), 1)

		// ADVANCED: Create a function in scrape.go that returns a list of file locations; call and use the result here
		// e.g., func finder(query string) []string { ... }

		found := false
		for _, File := range Files {
			if File.Filename == q[0] {
				json.NewEncoder(w).Encode(File.Location)
				found = true
			}
		}

		// Done
		//TODO_9: Handle when no matches exist; print a useful json response to the user; hint you might need a "FOUND variable" to check here ...
		if !found {
			w.Write([]byte(`{ "status": "File not found" }`))
		}
	} else {
		// didn't pass in a search term, show all that you've found
		w.Write([]byte(`"files":`))
		json.NewEncoder(w).Encode(Files)
	}
}

func IndexFiles(w http.ResponseWriter, r *http.Request) {
	logging.IfLevel(fmt.Sprintf("Entering %s end point", r.URL.Path), 1)
	w.Header().Set("Content-Type", "application/json")

	locations, locOK := r.URL.Query()["location"]

	// DONE
	//TODO_10: Currently there is a huge risk with this code ... namely, we can search from the root /
	//TODO_10: Assume the location passed starts at /home/ (or in Windows pick some "safe?" location)
	//TODO_10: something like ...  rootDir string := "???"
	//TODO_10: create another variable and append location to rootDir (where appropriate) to patch this hole

	rootDir := "/home/jasonfantl/Applications/Unity/"
	location := rootDir + locations[0]

	if locOK && len(location) > 0 {
		w.WriteHeader(http.StatusOK)

	} else {
		w.WriteHeader(http.StatusFailedDependency)
		w.Write([]byte(`{ "parameters" : {"required": "location",`))
		w.Write([]byte(`"optional": "regex"},`))
		w.Write([]byte(`"examples" : { "required": "/indexer?location=/xyz",`))
		w.Write([]byte(`"optional": "/indexer?location=/xyz&regex=(i?).md"}}`))
		return
	}

	//wrapper to make "nice json"
	w.Write([]byte(`{ `))

	// DONE
	// TODO_11: Currently the code DOES NOT do anything with an optionally passed regex parameter
	// Define the logic required here to call the new function walkFn2(w,regex[0])
	// Hint, you need to grab the regex parameter (see how it's done for location above...)
	regex, regexOK := r.URL.Query()["regex"]

	if regexOK {
		if err := filepath.Walk(location, walkFn2(w, `(i?)`+regex[0])); err != nil {
			log.Panicln(err)
		}
	} else {
		// else run code to locate files matching stored regular expression
		if err := filepath.Walk(location, walkFn(w)); err != nil {
			log.Panicln(err)
		}
	}

	//wrapper to make "nice json"
	w.Write([]byte(` "status": "completed"} `))

}

// DONE
//TODO_12 create endpoint that calls resetRegEx AND *** clears the current Files found; ***
//TODO_12 Make sure to connect the name of your function back to the reset endpoint main.go!
func ResetFilesAndReg(w http.ResponseWriter, r *http.Request) {
	logging.IfLevel(fmt.Sprintf("Entering %s end point", r.URL.Path), 1)
	w.Header().Set("Content-Type", "application/json")

	resetRegEx()
	Files = nil

	w.Write([]byte(`{ "status": "reset" }`))
}

// DONE
//TODO_13 create endpoint that calls clearRegEx ;
//TODO_13 Make sure to connect the name of your function back to the clear endpoint main.go!
func ClearFilesAndReg(w http.ResponseWriter, r *http.Request) {
	logging.IfLevel(fmt.Sprintf("Entering %s end point", r.URL.Path), 1)
	w.Header().Set("Content-Type", "application/json")

	clearRegEx()

	w.Write([]byte(`{ "status": "cleared" }`))
}

// DONE
//TODO_14 create endpoint that calls addRegEx ;
//TODO_12 Make sure to connect the name of your function back to the addsearch endpoint in main.go!
// consider using the mux feature
// params := mux.Vars(r)
// params["regex"] should contain your string that you pass to addRegEx
// If you try to pass in (?i) on the command line you'll likely encounter issues
// Suggestion : prepend (?i) to the search query in this endpoint

func AddReg(w http.ResponseWriter, r *http.Request) {
	logging.IfLevel(fmt.Sprintf("Entering %s end point", r.URL.Path), 1)
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	if s, ok := params["regex"]; ok {
		addRegEx(`(?i)` + s)
	} else {
		w.WriteHeader(http.StatusFailedDependency)
		w.Write([]byte(`mux somehow failed`))
		return
	}
}
