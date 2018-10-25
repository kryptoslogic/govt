// vtFileReport - fetches a report from VirusTotal for the given resource. A resource can be MD5, SHA-1 or SHA-2 of a file.
//  vtFileReport -rsrc=8ac31b7350a95b0b492434f9ae2f1cde
//
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/kryptoslogic/govt"
)

var apikey string
var apiurl string
var rsrc string

// init - initializes flag variables.
func init() {
	flag.StringVar(&apikey, "apikey", os.Getenv("VT_API_KEY"), "Set environment variable VT_API_KEY to your VT API Key or specify on prompt")
	flag.StringVar(&apiurl, "apiurl", "https://www.virustotal.com/vtapi/v2/", "URL of the VirusTotal API to be used.")
	flag.StringVar(&rsrc, "rsrc", "", "resource of file to retrieve report for. A resource can be md5, sha-1 or sha-2 sum of a file.")
}

// check - an error checking function
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	flag.Parse()
	if rsrc == "" {
		fmt.Println("-rsrc=<md5|sha-1|sha-2> not given!")
		os.Exit(1)
	}
	c, err := govt.New(govt.SetApikey(apikey), govt.SetUrl(apiurl))
	check(err)

	// get a file report
	r, err := c.GetFileReport(rsrc, "0")
	check(err)
	j, err := json.MarshalIndent(r, "", "    ")
	check(err)
	fmt.Printf("FileReport: ")
	os.Stdout.Write(j)
}
