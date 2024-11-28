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
	fmt.Print("Enter your Domain : ")
	for scanner.Scan() {
		checkDomain(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Printf("Error: Could not read from input: %v\n", err)
	}
}

func checkDomain(domain string) {
	var hasMX, hasSPF, hasDMARC bool
	var SPF_Record, DMARC_Record string

	mxRecord, err := net.LookupMX(domain)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}
	if len(mxRecord) > 0 {
		hasMX = true
	}

	txtRecords, err := net.LookupTXT(domain)
	if err != nil {
		log.Printf("Error : %v\n", err)
	}
	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			SPF_Record = record
		}
	}

	dmarcRecord, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		log.Printf("Error : %v\n", err)
	}
	for _, record := range dmarcRecord {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			DMARC_Record = record
		}
	}

	fmt.Printf("\n\nDomain: %v\n\nhasMX: %v\n\nhasSPF: %v\n\nSPF Record: %v\n\nhasDMARC: %v\n\nDMARC Record: %v", domain, hasMX, hasSPF, SPF_Record, hasDMARC, DMARC_Record)
}
