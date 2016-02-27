package main 

import
(
	"fmt"
	"flag"
	"io/ioutil"
	"strings"
)

func main() {
	var htmlFile string
	var htmlString string
	
	// Get command line parameter for html file
	flag.StringVar(&htmlFile, "htmlfile", "", "Specifiy a path to the html file to be converted to a C string literal")
	flag.Parse()

	if (htmlFile == "") {
		flag.PrintDefaults()
		return
	}

	// Read html file to string
	fmt.Println("Loading HTML file: " + htmlFile)
	fileBytes, err := ioutil.ReadFile(htmlFile)

	if (err != nil) {
		fmt.Println("Failed to read file")
		return
	}

	htmlString = string(fileBytes)

	// Find all \ and convert to \\
	htmlString = strings.Replace(htmlString, "\\", "\\\\", -1)

	// Find all " and convert to \"
	htmlString = strings.Replace(htmlString, "\"", "\\\"", -1)

	// Find all ' and convert to \'
	htmlString = strings.Replace(htmlString, "'", "\\'", -1)

	// Find all turn all new lines into \n
	htmlString = strings.Replace(htmlString, "\n", "\\n", -1)

	// Remove all carriage returns
	htmlString = strings.Replace(htmlString, "\r", "", -1)

	// Remove all tabs
	htmlString = strings.Replace(htmlString, "\t", "", -1)

	// Print the output
	fmt.Println(htmlString)
}