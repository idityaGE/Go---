This Go program performs basic email domain verification by checking three critical DNS records for a given domain: **MX (Mail Exchange), SPF (Sender Policy Framework), and DMARC (Domain-based Message Authentication, Reporting, and Conformance)**. Here's a detailed breakdown of how the program works:

---

### Code Breakdown
1. **Package Imports:**
   ```go
   import (
       "bufio"
       "fmt"
       "log"
       "net"
       "os"
       "strings"
   )
   ```
   - **`bufio`**: For buffered input from the standard input.
   - **`fmt`**: For formatted input and output.
   - **`log`**: For logging errors and other important information.
   - **`net`**: Provides network-related functions like DNS lookups.
   - **`os`**: For interacting with the operating system, like reading input.
   - **`strings`**: To manipulate strings (e.g., checking prefixes).

2. **Main Function:**
   ```go
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
   ```
   - Reads a domain name from the user input.
   - For each domain entered, it calls `checkDomain` to verify the DNS records.
   - Handles any errors during input using the scanner.

3. **`checkDomain` Function:**
   ```go
   func checkDomain(domain string) {
       var hasMX, hasSPF, hasDMARC bool
       var SPF_Record, DMARC_Record string
   ```
   - Initializes boolean flags (`hasMX`, `hasSPF`, `hasDMARC`) to track if specific records are found.
   - Prepares variables to store the actual SPF and DMARC records.

4. **MX Record Check:**
   ```go
   mxRecord, err := net.LookupMX(domain)
   if err != nil {
       log.Printf("Error: %v\n", err)
   }
   if len(mxRecord) > 0 {
       hasMX = true
   }
   ```
   - Uses `net.LookupMX` to fetch MX records for the domain.
   - If records exist, it sets `hasMX` to `true`.

5. **SPF Record Check:**
   ```go
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
   ```
   - Fetches TXT records for the domain using `net.LookupTXT`.
   - Iterates through the records to find one starting with `v=spf1`, which identifies an SPF record.
   - Sets `hasSPF` to `true` and stores the record if found.

6. **DMARC Record Check:**
   ```go
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
   ```
   - Looks up `_dmarc.<domain>` for DMARC records using `net.LookupTXT`.
   - Finds records starting with `v=DMARC1`, sets `hasDMARC` to `true`, and stores the record.

7. **Output Results:**
   ```go
   fmt.Printf("\n\nDomain: %v\n\nhasMX: %v\n\nhasSPF: %v\n\nSPF Record: %v\n\nhasDMARC: %v\n\nDMARC Record: %v", domain, hasMX, hasSPF, SPF_Record, hasDMARC, DMARC_Record)
   ```
   - Prints whether the domain has MX, SPF, and DMARC records, along with their values if available.

---

### How Email Verifiers Work
An **email verifier** checks if an email domain or address can receive emails and is configured correctly. This program focuses on verifying domain-level configurations through DNS records:

1. **MX (Mail Exchange) Records:**
   - Specify the mail servers for the domain.
   - Essential for receiving emails.
   - The presence of MX records means the domain can handle incoming emails.

2. **SPF (Sender Policy Framework):**
   - A TXT record that specifies which mail servers are authorized to send emails on behalf of the domain.
   - Helps prevent spoofing by enabling email providers to verify the sender's authenticity.

3. **DMARC (Domain-based Message Authentication, Reporting, and Conformance):**
   - Another TXT record used to specify how the domain handles spoofed messages (e.g., reject, quarantine, or none).
   - Builds on SPF and DKIM (not checked here) to improve email security.

By combining these checks, the program ensures the domain is configured to send and receive emails securely.

---

### Summary of Functionality
1. Prompts the user to enter a domain.
2. Checks for:
   - **MX Records**: Verifies the domain can receive emails.
   - **SPF Records**: Confirms the domain is authorized to send emails.
   - **DMARC Records**: Ensures proper handling of unauthorized emails.
3. Outputs the verification results.

This tool is useful for debugging email configuration issues or verifying domain readiness for email services.