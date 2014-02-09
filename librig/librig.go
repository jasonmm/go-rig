package librig

import (
	"bufio"
	libGoWc "github.com/jasonmm/gowc/libgowc"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	FEMALE_FIRST_NAMES = "fnames.idx"
	MALE_FIRST_NAMES   = "mnames.idx"
	LAST_NAMES         = "lnames.idx"
	LOCATION_DATA      = "locdata.idx"
	STREET_DATA        = "street.idx"

	MALE   = 1
	FEMALE = 2
	EITHER = 3
)

// The directory where the data files are found.
var DataDirectory = "/usr/share/rig/"

// The gender of the fake identity's first name.
var NameGender = EITHER

// Whether the non-area code part of the phone number has x's.
var PhoneHasX = true

// Struct holding the new identity.
type Identity struct {
	FirstName string
	LastName  string
	City      string
	State     string
	Zip       string
	Street    string
	Phone     string
}

// Return the libGoWc.Metrics struct for the given file.
func countLines(filePath string) libGoWc.Metrics {
	metric, err := libGoWc.ProcessSingleFile(filePath)
	if err != nil {
		panic(err)
	}
	return metric
}

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

// The city, state, zip, and area code are stored all on the same line in
// the data file.  Separate these and create a phone number, then return
// all four as separate variables: city, state, zip, phone.
func getCityStateZipPhone(loc string) (city string, state string, zip string, phone string) {
	parts := strings.Split(loc, " ")

	phone = "(" + parts[2] + ") "
	if PhoneHasX {
		phone = phone + "xxx-xxxx"
	} else {
		exchange := rand.Intn(898) + 101
		subscriber := rand.Intn(8998) + 1001
		phone = phone + strconv.Itoa(exchange) + "-" + strconv.Itoa(subscriber)
	}

	return parts[0], parts[1], parts[3], phone
}

// Returns a random identity.
func GetIdentity() Identity {

	femaleFilename := DataDirectory + FEMALE_FIRST_NAMES
	maleFilename := DataDirectory + MALE_FIRST_NAMES
	lastNameFilename := DataDirectory + LAST_NAMES
	locationFilename := DataDirectory + LOCATION_DATA
	streetFilename := DataDirectory + STREET_DATA

	// Get the number of lines in all the data files.
	femaleMetric := countLines(femaleFilename)
	maleMetric := countLines(maleFilename)
	lastNameMetric := countLines(lastNameFilename)
	locMetric := countLines(locationFilename)
	streetMetric := countLines(streetFilename)

	// Seed the random number generator.
	rand.Seed(time.Now().UTC().UnixNano())

	// Get the line number we will choose from.
	femaleIndex := rand.Intn(femaleMetric.Lines)
	maleIndex := rand.Intn(maleMetric.Lines)
	lastNameIndex := rand.Intn(lastNameMetric.Lines)
	locIndex := rand.Intn(locMetric.Lines)
	streetIndex := rand.Intn(streetMetric.Lines)

	// Get the chosen random line from the files.
	femaleName := getLine(femaleFilename, femaleIndex)
	maleName := getLine(maleFilename, maleIndex)
	lastName := getLine(lastNameFilename, lastNameIndex)
	locStr := getLine(locationFilename, locIndex)
	street := getLine(streetFilename, streetIndex)

	// Determine if we display a male or female first name.
	name := femaleName
	if NameGender == MALE {
		name = maleName
	} else if NameGender == EITHER {
		if rand.Intn(2) == 1 {
			name = maleName
		}
	}

	// Create a street number.
	streetNumber := rand.Intn(9900) + 100

	// Get city/state/zip and phone number.
	city, state, zip, phone := getCityStateZipPhone(locStr)

	// Create and return the identity.
	ident := Identity{
		FirstName: name,
		LastName:  lastName,
		City:      city,
		State:     state,
		Zip:       zip,
		Street:    strconv.Itoa(streetNumber) + " " + street,
		Phone:     phone,
	}
	return ident
}
