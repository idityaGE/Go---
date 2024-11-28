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
	fmt.Println("Email Domain Verifier")
	fmt.Println("----------------------")
	fmt.Println("Type 'exit' to quit.")
	for {
		fmt.Print("\nEnter a Domain: ")
		scanner.Scan()
		domain := strings.TrimSpace(scanner.Text())

		// Exit if user types "exit"
		if strings.ToLower(domain) == "exit" {
			fmt.Println("Exiting program. Goodbye!")
			break
		}

		// Validate domain
		if domain == "" || !isValidDomain(domain) {
			fmt.Println("Invalid domain. Please enter a valid domain.")
			continue
		}

		// Check domain and display results
		checkDomain(domain)
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Error reading input: %v\n", err)
	}
}

// Checks the domain's DNS records
func checkDomain(domain string) {
	hasMX, hasSPF, hasDMARC := false, false, false
	SPFRecord, DMARCRecord := "", ""

	// Check MX Records
	mxRecords, err := net.LookupMX(domain)
	if err != nil {
		fmt.Printf("MX Record Error: %v\n", err)
	} else if len(mxRecords) > 0 {
		hasMX = true
	}

	// Check SPF Records
	txtRecords, err := net.LookupTXT(domain)
	if err != nil {
		fmt.Printf("SPF Record Error: %v\n", err)
	} else {
		for _, record := range txtRecords {
			if strings.HasPrefix(record, "v=spf1") {
				hasSPF = true
				SPFRecord = record
				break
			}
		}
	}

	// Check DMARC Records
	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		fmt.Printf("DMARC Record Error: %v\n", err)
	} else {
		for _, record := range dmarcRecords {
			if strings.HasPrefix(record, "v=DMARC1") {
				hasDMARC = true
				DMARCRecord = record
				break
			}
		}
	}

	// Display results
	fmt.Println("\nVerification Results:")
	fmt.Printf("Domain: %s\n", domain)
	fmt.Printf("Has MX Records: %v\n", hasMX)
	if hasMX {
		fmt.Println("   - Mail servers found.")
	} else {
		fmt.Println("   - No mail servers found.")
	}
	fmt.Printf("Has SPF Records: %v\n", hasSPF)
	if hasSPF {
		fmt.Printf("   - SPF Record: %s\n", SPFRecord)
	} else {
		fmt.Println("   - No SPF record found.")
	}
	fmt.Printf("Has DMARC Records: %v\n", hasDMARC)
	if hasDMARC {
		fmt.Printf("   - DMARC Record: %s\n", DMARCRecord)
	} else {
		fmt.Println("   - No DMARC record found.")
	}
}

// Validates the domain using basic criteria
func isValidDomain(domain string) bool {
	if len(domain) < 3 || len(domain) > 253 {
		return false
	}
	if strings.HasPrefix(domain, ".") || strings.HasSuffix(domain, ".") {
		return false
	}
	for _, part := range strings.Split(domain, ".") {
		if len(part) < 1 || len(part) > 63 {
			return false
		}
	}
	return true
}
