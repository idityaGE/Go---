package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"net/smtp"
	"os"
	"regexp"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Email Verifier")
	fmt.Println("--------------")
	fmt.Println("Type 'exit' to quit.")

	for {
		fmt.Print("\nEnter an Email Address: ")
		scanner.Scan()
		email := strings.TrimSpace(scanner.Text())

		// Exit if user types "exit"
		if strings.ToLower(email) == "exit" {
			fmt.Println("Exiting program. Goodbye!")
			break
		}

		// Validate email
		if !isValidEmail(email) {
			fmt.Println("Invalid email format. Please enter a valid email address.")
			continue
		}

		// Split into local part and domain part
		parts := strings.Split(email, "@")
		localPart, domain := parts[0], parts[1]

		// Verify domain and email
		fmt.Printf("\nValidating Email: %s\n", email)
		checkDomain(domain)
		checkEmailAddress(localPart, domain)
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Error reading input: %v\n", err)
	}
}

// Checks the domain's DNS records
func checkDomain(domain string) {
	hasMX := false

	// Check MX Records
	mxRecords, err := net.LookupMX(domain)
	if err != nil {
		fmt.Printf("Domain Check Failed: %v\n", err)
		return
	}
	if len(mxRecords) > 0 {
		hasMX = true
	}

	// Display results
	fmt.Printf("Domain: %s\n", domain)
	fmt.Printf("Has MX Records: %v\n", hasMX)
	if hasMX {
		fmt.Println("   - Mail servers found.")
	} else {
		fmt.Println("   - No mail servers found.")
	}
}

// Checks if the email address is accepted by the domain's SMTP server
func checkEmailAddress(localPart, domain string) {
	mxRecords, err := net.LookupMX(domain)
	if err != nil || len(mxRecords) == 0 {
		fmt.Println("Cannot verify email address: No MX records found.")
		return
	}

	// Connect to the SMTP server of the highest-priority MX record
	mxServer := mxRecords[0].Host
	fmt.Printf("Checking email address on SMTP server: %s\n", mxServer)

	client, err := smtp.Dial(mxServer + ":25")
	if err != nil {
		fmt.Printf("SMTP Connection Failed: %v\n", err)
		return
	}
	defer client.Close()

	// Perform handshake and set sender and recipient
	err = client.Hello("example.com")
	if err != nil {
		fmt.Printf("SMTP HELO Failed: %v\n", err)
		return
	}

	err = client.Mail("verify@example.com") // Set a dummy sender email
	if err != nil {
		fmt.Printf("SMTP MAIL FROM Failed: %v\n", err)
		return
	}

	err = client.Rcpt(localPart + "@" + domain) // Check recipient
	if err != nil {
		fmt.Printf("Recipient Rejected: %v\n", err)
		return
	}

	fmt.Printf("The email address %s@%s is valid!\n", localPart, domain)
}

// Validates the email address syntax
func isValidEmail(email string) bool {
	// Simple regex for email validation
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(regex)
	return re.MatchString(email)
}
