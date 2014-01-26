package main

import (
	"bufio"
	"fmt"
	"github.com/jasonmm/gowc/libgowc"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Default filenames.
const (
	FNAMES  = "/usr/share/rig/fnames.idx"
	MNAMES  = "/usr/share/rig/mnames.idx"
	LNAMES  = "/usr/share/rig/lnames.idx"
	LOCDATA = "/usr/share/rig/locdata.idx"
	STREET  = "/usr/share/rig/street.idx"
)

// Get the specified line number from the given file.  If the given line number
// is larger than the number of lines in the file then the last line is returned.
func getLine(filePath string, lineNum int) string {
	cnt := 0

	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		cnt = cnt + 1
		if cnt >= lineNum {
			break
		}
	}
	if err = scanner.Err(); err != nil {
		panic(err)
	}
	return scanner.Text()
}

// Return the libgowc.Metrics struct for the given file.
func countLines(filePath string) libgowc.Metrics {
	metric, err := libgowc.ProcessSingleFile(filePath)
	if err != nil {
		panic(err)
	}
	return metric
}

// Break out the city/state/zip and create a phone number from the area code.
func getCityStateZipPhone(loc string) (string, string, string, string) {
	parts := strings.Split(loc, " ")

	phone := "(" + parts[2] + ") xxx-xxxx"

	return parts[0], parts[1], parts[3], phone
}

func main() {
	// Get the number of lines in all the data files.
	femaleMetric := countLines(FNAMES)
	maleMetric := countLines(MNAMES)
	lastNameMetric := countLines(LNAMES)
	locMetric := countLines(LOCDATA)
	streetMetric := countLines(STREET)

	// Seed the random number generator.
	rand.Seed(time.Now().UTC().UnixNano())

	// Get the line number we will choose from.
	femaleIndex := rand.Intn(femaleMetric.Lines)
	maleIndex := rand.Intn(maleMetric.Lines)
	lastNameIndex := rand.Intn(lastNameMetric.Lines)
	locIndex := rand.Intn(locMetric.Lines)
	streetIndex := rand.Intn(streetMetric.Lines)

	// Get the chosen random line from the files.
	femaleName := getLine(FNAMES, femaleIndex)
	maleName := getLine(MNAMES, maleIndex)
	lastName := getLine(LNAMES, lastNameIndex)
	locStr := getLine(LOCDATA, locIndex)
	street := getLine(STREET, streetIndex)

	// Determine if we display a male or female first name.
	name := femaleName
	if rand.Intn(2) == 1 {
		name = maleName
	}

	// Create a street number.
	streetNumber := rand.Intn(9900) + 100

	// Get city/state/zip and phone number.
	city, state, zip, phone := getCityStateZipPhone(locStr)

	// Output the identity.
	fmt.Println(name, lastName)
	fmt.Println(streetNumber, street)
	fmt.Println(city, state, zip)
	fmt.Println(phone)
	fmt.Println()
}
