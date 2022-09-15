package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"

	nu "github.com/devpablocristo/sitemap-generator/navigate-url"
	xml "github.com/devpablocristo/sitemap-generator/xml-file"
)

// type URL struct {
// 	url   string
// 	depth int
// }

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("\nSitemap Generator commands:")
	fmt.Println("-u starting url.")
	fmt.Println("-p number of parallel workers to navigate through site.")
	fmt.Println("-of output file path.")
	fmt.Println("-md  max depth of url navigation recursion.")
	fmt.Println("example: sitemapgen -u http://localhost:8081/index.html -p 4 -of sitemap.xml -md 0")
	fmt.Println("")
	fmt.Print("sitemapgen ")
	cmd, _ := reader.ReadString('\n')
	cmd = strings.Replace(cmd, "\n", "", 1)

	cmds := strings.Split(cmd, " ")
	//fmt.Println(cmds)

	var startURL string
	var parallel int
	var outputFile string
	var maxDepth int
	for i := 0; i < len(cmds); i++ {
		switch cmds[i] {
		case "-u":
			startURL = cmds[i+1]
		case "-p":
			parallel, _ = strconv.Atoi(cmds[i+1])
		case "-of":
			outputFile = cmds[i+1]
		case "-md":
			maxDepth, _ = strconv.Atoi(cmds[i+1])
		}
	}

	maxDepth++
	nu.ListURLs = make(map[string]bool)
	nu.ListURLs[startURL] = true

	var wg sync.WaitGroup
	var mu sync.Mutex

	if parallel > 0 {
		for i := 0; i < parallel; i++ {
			wg.Add(1)
			go nu.Worker(startURL, maxDepth, &wg, &mu)
		}
	} else {
		fmt.Println("Parallel must be greater than zero")
		os.Exit(1)
	}

	wg.Wait()

	xml, err := xml.CreateXML(outputFile)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\n" + xml + "\n")
}
