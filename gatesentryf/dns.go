package gatesentryf

import (
	"fmt"

	gatesentryDnsServer "bitbucket.org/abdullah_irfan/gatesentryf/dns/server"
	gatesentry2logger "bitbucket.org/abdullah_irfan/gatesentryf/logger"
	gatesentry2storage "bitbucket.org/abdullah_irfan/gatesentryf/storage"
)

var (
	blocklists = []string{
		"https://raw.githubusercontent.com/StevenBlack/hosts/master/hosts",
		"https://easylist.to/easylist/easylist.txt",
		"https://raw.githubusercontent.com/anudeepND/blacklist/master/adservers.txt",
		"https://v.firebog.net/hosts/AdguardDNS.txt",
		"https://raw.githubusercontent.com/PolishFiltersTeam/KADhosts/master/KADhosts.txt",
		"https://raw.githubusercontent.com/FadeMind/hosts.extras/master/add.Spam/hosts",
		"https://v.firebog.net/hosts/static/w3kbl.txt",
		"https://adaway.org/hosts.txt",
		"https://v.firebog.net/hosts/RPiList-Phishing.txt",
		"https://v.firebog.net/hosts/RPiList-Malware.txt",
		"https://gitlab.com/quidsup/notrack-blocklists/raw/master/notrack-malware.txt",
		"https://pgl.yoyo.org/adservers/serverlist.php?hostformat=hosts&showintro=0&mimetype=plaintext",
		"https://bitbucket.org/ethanr/dns-blacklists/raw/8575c9f96e5b4a1308f2f12394abd86d0927a4a0/bad_lists/Mandiant_APT1_Report_Appendix_D.txt",
		// Add more blocklist URLs here
	}
)

func DNSServerThread(baseDir string, logger *gatesentry2logger.Log, c <-chan int, settings *gatesentry2storage.MapStore) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	fmt.Println("Inside DNS server thread")
	for {
		select {
		case msg := <-c:
			fmt.Println("Received message:", msg)
			if msg == 1 {
				// fmt.Println("ACTUAL Starting DNS server")
				// Start the DNS server
				go gatesentryDnsServer.StartDNSServer(baseDir, logger, blocklists, settings)
				fmt.Println("ACTUAL DNS server started")
			} else if msg == 2 {
				fmt.Println("ACTUAL Stopping DNS server")
				// Stop the DNS server
				go gatesentryDnsServer.StopDNSServer()
				fmt.Println("ACTUAL DNS server stopped")
			}
		}
	}

}

// func DNSServerThread(baseDir string, logger *gatesentry2logger.Log, c chan int) {
// 	fmt.Println("Inside DNS server thread")
// 	select {
// 	case msg := <-c:
// 		fmt.Println("Received message: " + fmt.Sprint(msg))
// 		if msg == 1 {
// 			fmt.Println("ACTUAL Starting DNS server")
// 			// Start the DNS server
// 			// go gatesentryDnsServer.StartDNSServer(baseDir, logger, blocklists)
// 			fmt.Println("ACTUAL DNS server started")
// 		} else if msg == 2 {
// 			fmt.Println("ACTUAL Stopping DNS server")
// 			// Stop the DNS server
// 			// go gatesentryDnsServer.StopDNSServer()
// 			fmt.Println("ACTUAL DNS server stopped")
// 		}
// 	}
// for {
// 	fmt.Println("Waiting for message")
// 	msg := <-c

// 	fmt.Println("Received message: " + fmt.Sprint(msg))
// 	if msg == 2 {
// 		fmt.Println("Stopping DNS server")
// 		// Stop the DNS server
// 		gatesentryDnsServer.StopDNSServer()
// 	} else if msg == 1 {
// 		fmt.Println("Starting DNS server")
// 		// Start the DNS server
// 		gatesentryDnsServer.StartDNSServer(baseDir, logger, blocklists)
// 		fmt.Println("DNS server started")
// 	}
// }
// select {
// case msg := <-c:
// 	fmt.Println("Received message: " + msg)
// 	switch msg {
// 	case "stop":
// 		fmt.Println("Stopping DNS server")
// 		// Stop the DNS server
// 		gatesentryDnsServer.StopDNSServer()

// 	case "start":
// 		fmt.Println("Starting DNS server")
// 		// Start the DNS server
// 		gatesentryDnsServer.StartDNSServer(baseDir, logger, blocklists)
// 		fmt.Println("DNS server started")

// 	default:
// 		// Do nothing
// 	}
// }
// }
