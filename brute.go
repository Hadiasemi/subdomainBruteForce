package main

import (
	"bufio"
	"fmt"
	"math"
	"net"
	"os"
	"runtime"
	"sync"
)

func DNScheck(sub []string, wg *sync.WaitGroup, domain string) {

	defer wg.Done()
	colorCyan := "\033[36m"
	for _, line := range sub {
		subdomain := fmt.Sprintf("%s.%s", line, domain)
		_, err := net.LookupIP(subdomain)
		if err == nil {
			fmt.Println(string(colorCyan), subdomain)
		}
	}
}

func fileReader(file string) []string {

	var fileLines []string
	readFile, err := os.Open(file)

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	readFile.Close()
	return fileLines

}

func ChunkStringSlice(s []string, chunkSize int) [][]string {
	chunkNum := int(math.Ceil(float64(len(s)) / float64(chunkSize)))
	res := make([][]string, 0, chunkNum)
	for i := 0; i < chunkNum-1; i++ {
		res = append(res, s[i*chunkSize:(i+1)*chunkSize])
	}
	res = append(res, s[(chunkNum-1)*chunkSize:])
	return res
}

func main() {
	var wg sync.WaitGroup
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Printf("Usage: %s url", os.Args[0])
		os.Exit(0)
	}
	file := "./subdomains-top1million-110000.txt"
	// file := "./deepmagic.com-prefixes-top500.txt"
	domain := args[0]
	subdomains := fileReader(file)

	totalSubdomain := len(subdomains)

	totalCPU := runtime.NumCPU()

	wg.Add(totalCPU)
	parallelDomainList := ChunkStringSlice(subdomains, totalSubdomain/totalCPU)
	size := len(parallelDomainList)

	for i := 0; i < size; i++ {

		go DNScheck(parallelDomainList[i], &wg, domain)
	}
	// for i := 0; i < totalSubdomain; i += totalSubdomain - 1 / totalCPU {
	// 	end := (i+(totalSubdomain/totalCPU + 1))

	// 	go DNScheck(subdomains[i:end], &wg, domain)
	// }

	wg.Wait()

}


