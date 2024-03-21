package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	topDomainNames, err := readDomainsFromCSV()

	t1 := time.Now()
	if err != nil {
		log.Fatal("Error while reading domains ", err.Error())
	}
	statusChecker(topDomainNames)

	fmt.Println("\nTotal time: ", time.Since(t1))
}

func readDomainsFromCSV() ([]string, error) {
	file, err := ioutil.ReadFile("topDomains.csv")
	if err != nil {
		log.Println("Error while reading file...", err.Error())
		return nil, err
	}

	var domains []string
	reader := csv.NewReader(bytes.NewReader(file))
	all, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error while reading records from csv...", err.Error())
		return nil, err
	}
	for _, rec := range all {
		domains = append(domains, rec[1])
	}
	if len(domains) > 0 {
		return domains[1:], nil
	}
	return domains, nil
}

func statusChecker(topDomainNames []string) {
	for idx, host := range topDomainNames {
		lookupHost(host, idx)
	}
}

func lookupHost(host string, idx int) {
	client := http.Client{
		Timeout: time.Second * 2,
	}
	_, err := client.Get(host)
	if err != nil {
		fmt.Println(fmt.Sprintf("%d, %s website is down", idx, host))
		return
	}
	fmt.Println(fmt.Sprintf("%d, %s website is up", idx, host))
}
