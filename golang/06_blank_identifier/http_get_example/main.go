package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// Example of use of blank identifier to not treat err case
	google, _ := http.Get("https://www.google.com/") // DL web page
	google_page, _ := ioutil.ReadAll(google.Body)    // read the body of the page
	google.Body.Close()                              // Closes the DL
	fmt.Printf("%s", google_page)

}
