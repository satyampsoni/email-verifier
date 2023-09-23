package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Domain, hasMX, hasSPF, hasDMARC, hasDKIM, spfRecord, dmarcRecord")

	for scanner.Scan() {
		checkDomain(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("could not read from input: %v\n", err)
	}
}

func checkDomain(Domain string) {
	var hasMX, hasSPF, hasDMARC, hasDKIM bool
	var spfRecord, dmarcRecord string

	mx, err := net.LookupMX(Domain)
	if err != nil {
		log.Printf("could not lookup MX records for %s: %v\n", Domain, err)
	}
	if len(mx) > 0 {
		hasMX = true
	}

	txtRecords, err := net.LookupTXT(Domain)
	if err != nil {
		log.Printf("could not lookup TXT records for %s: %v\n", Domain, err)
	}
	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfRecord = record
			break
		}
	}

	dmarcRecords, err := net.LookupTXT("dmarc." + Domain)
	if err != nil {
		log.Printf("could not lookup TXT records for %s: %v\n", Domain, err)
	}
	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = record
			break
		}
	}

	fmt.Printf("%s, %v, %v, %v, %v, %s, %s\n", Domain, hasMX, hasSPF, hasDMARC, hasDKIM, spfRecord, dmarcRecord)
}


