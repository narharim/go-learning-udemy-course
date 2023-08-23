package main

import (
	"fmt"
	"net"
	"os"

	"github.com/coreos/go-iptables/iptables"
)

const (
	tableNAT    = "nat"
	tableFilter = "filter"

	chainOutput  = "OUTPUT"
	chainInput   = "INPUT"
	chainForward = "FORWARD"
	policyAccept = "ACCEPT"
	policyDrop   = "DROP"

	protoTCP = "tcp"

	destinationFlag = "-d"
	protoFlag       = "-p"
	matchFlag       = "-m"
	jumpFlag        = "-j"
)

func checkRootUser() bool {
	if os.Getuid() != 0 {
		fmt.Println("Error: This program should be run as root.")
		return false
	}
	fmt.Println("Running as root user.")
	return true
}

func initializeIPTables() (*iptables.IPTables, error) {
	ipt, err := iptables.New()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize IPTables instance: %w", err)
	}
	return ipt, nil
}

func addIPRule(ipt *iptables.IPTables, ip string) error {
	ipstr := fmt.Sprintf("%s%s", ip, "/32")

	ruleSpec := []string{
		destinationFlag, ipstr, protoFlag, protoTCP, matchFlag, protoTCP, jumpFlag, policyAccept,
	}

	if err := ipt.AppendUnique(tableFilter, chainForward, ruleSpec...); err != nil {
		return fmt.Errorf("failed to add rule for IP %s: %w", ipstr, err)
	}

	fmt.Println("IP rule added successfully.")
	return nil
}

func lookUpIP(domain string) (string, error) {
	ips, err := net.LookupIP(domain)
	if err != nil {
		return "", fmt.Errorf("failed to look up IP for %s: %w", domain, err)
	}
	return ips[0].String(), nil
}

func main() {
	domainList := []string{"google.com", "aws.amazon.com", "bitbucket.org", "security.ubuntu.com",
		"github.com", "sns.ap-south-1.amazonaws.com", "ap-south-1.ec2.archive.ubuntu.com",
		"pypi.org", "pypi.python.org", "files.pythonhosted.org"}

	if !checkRootUser() {
		os.Exit(1)
	}

	ipt, err := initializeIPTables()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	for _, domain := range domainList {
		ip, err := lookUpIP(domain)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}
		fmt.Printf("IPv4 %s\n", ip)

		if err := addIPRule(ipt, ip); err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}
}
