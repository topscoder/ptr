package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
)

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		printUsage()
		return
	}

	arg := args[0]

	if isFile(arg) {
		processFile(arg)
	} else if arg == "-" {
		processStdin()
	} else {
		processSingleIP(arg)
	}
}

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  ptr <ip_address_or_input_file>")
	fmt.Println()
	fmt.Println("Options:")
	fmt.Println("  <ip_address_or_input_file>  Provide an IP address, an input file with one IP address per line, or use '-' to read IP addresses from stdin.")
}

func isFile(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !info.IsDir()
}

func processFile(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var wg sync.WaitGroup
	ipChannel := make(chan string)

	// Start worker goroutines
	for i := 0; i < 10; i++ { // Adjust the number of goroutines as needed
		wg.Add(1)
		go processIPs(ipChannel, &wg)
	}

	// Send IPs to worker goroutines through channel
	for scanner.Scan() {
		ip := strings.TrimSpace(scanner.Text())
		if ip != "" {
			ipChannel <- ip
		}
	}
	close(ipChannel)

	wg.Wait()

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}

func processStdin() {
	scanner := bufio.NewScanner(os.Stdin)
	var wg sync.WaitGroup
	ipChannel := make(chan string)

	// Start worker goroutines
	for i := 0; i < 10; i++ { // Adjust the number of goroutines as needed
		wg.Add(1)
		go processIPs(ipChannel, &wg)
	}

	// Send IPs to worker goroutines through channel
	for scanner.Scan() {
		ip := strings.TrimSpace(scanner.Text())
		if ip != "" {
			ipChannel <- ip
		}
	}
	close(ipChannel)

	wg.Wait()

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}
}

func processSingleIP(ip string) {
	var wg sync.WaitGroup
	ipChannel := make(chan string)

	// Start worker goroutines
	for i := 0; i < 10; i++ { // Adjust the number of goroutines as needed
		wg.Add(1)
		go processIPs(ipChannel, &wg)
	}

	// Send the single IP to worker goroutines through channel
	ipChannel <- ip
	close(ipChannel)

	wg.Wait()
}

func processIPs(ipChannel <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	for ip := range ipChannel {
		ptrDomainNames, err := net.LookupAddr(ip)
		if err != nil {
			fmt.Printf("%s\t%v\n", ip, err)
			continue
		}

		var cleanedDomainNames []string

		for _, ptrDomainName := range ptrDomainNames {
			// Remove the brackets from the PTR domain name.
			ptrDomainName = strings.Trim(ptrDomainName, "[")

			// Remove the last character if it is a dot.
			if ptrDomainName[len(ptrDomainName)-1] == '.' {
				ptrDomainName = ptrDomainName[:len(ptrDomainName)-1]
			}

			cleanedDomainNames = append(cleanedDomainNames, ptrDomainName)
		}

		fmt.Printf("%s\t%s\n", ip, strings.Join(cleanedDomainNames, "\n"))
	}
}
