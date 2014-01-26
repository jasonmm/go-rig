package main

import (
	"bufio"
	"fmt"
	"github.com/jasonmm/gowc/libgowc"
	"math/rand"
	"os"
	"time"
)

const (
	FNAMES  = "/usr/share/rig/fnames.idx"
	MNAMES  = "/usr/share/rig/mnames.idx"
	LNAMES  = "/usr/share/rig/lnames.idx"
	LOCDATA = "/usr/share/rig/locdata.idx"
	STREET  = "/usr/share/rig/street.idx"
)

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

func countLines(filePath string) libgowc.Metrics {
	metric, err := libgowc.ProcessSingleFile(filePath)
	if err != nil {
		panic(err)
	}
	return metric
}

func main() {

	femaleMetric := countLines(FNAMES)
	maleMetric := countLines(MNAMES)
	lastNameMetric := countLines(LNAMES)
	locMetric := countLines(LOCDATA)
	streetMetric := countLines(STREET)

	rand.Seed(time.Now().UTC().UnixNano())

	femaleIndex := rand.Intn(femaleMetric.Lines)
	maleIndex := rand.Intn(maleMetric.Lines)
	lastNameIndex := rand.Intn(lastNameMetric.Lines)
	locIndex := rand.Intn(locMetric.Lines)
	streetIndex := rand.Intn(streetMetric.Lines)

	femaleName := getLine(FNAMES, femaleIndex)
	maleName := getLine(MNAMES, maleIndex)
	lastName := getLine(LNAMES, lastNameIndex)
	locStr := getLine(LOCDATA, locIndex)
	street := getLine(STREET, streetIndex)

	fmt.Println(femaleName)
	fmt.Println(maleName)
	fmt.Println(lastName)
	fmt.Println(locStr)
	fmt.Println(street)
}
