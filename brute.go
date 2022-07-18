package main

import (
	"bufio"
	"flag"
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

	file := flag.String("f", "./deepmagic.com-prefixes-top500.txt", "subdomain file")
	domain := flag.String("u", "", "specify the url")
	thread := flag.Int("t", runtime.NumCPU(), "thread")
	flag.Parse()

	if *domain == "" {
		fmt.Println("specify the url: -u url")
		os.Exit(0)
	}
	subdomains := fileReader(*file)

	totalSubdomain := len(subdomains)


	wg.Add(*thread)
	parallelDomainList := ChunkStringSlice(subdomains, totalSubdomain/ *thread)
	size := len(parallelDomainList)

	for i := 0; i < size; i++ {

		go DNScheck(parallelDomainList[i], &wg, *domain)
	}

	wg.Wait()

}
