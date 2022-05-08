package goiplookup

import (
	"bufio"
	_ "flag"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

const Usage = signature + "\n" + `Usage of GoIPLookUp:
  -v, --verbose 	[ verbose output { default: false } ]
  -h, --help 		[ prints help information ]
  -u, --url 		[ website url ]
  -p, --path 		[ file path for bulk urls ]
  -o, --output 		[ output filename ]
`

const signature = `
  _____      _____ _____  _                 _    _    _       
 / ____|    |_   _|  __ \| |               | |  | |  | |      
| |  __  ___  | | | |__) | |     ___   ___ | | _| |  | |_ __  
| | |_ |/ _ \ | | |  ___/| |    / _ \ / _ \| |/ / |  | | '_ \ 
| |__| | (_) || |_| |    | |___| (_) | (_) |   <| |__| | |_) |
 \_____|\___/_____|_|    |______\___/ \___/|_|\_\\____/| .__/ 
                                                       | |    
                                                       |_| by heavykenny
twitter @heavykenny`

func FilterURL(url string) string {
	if strings.HasPrefix(url, "https://") {
		url = url[8:]
	}

	if strings.HasPrefix(url, "http://") {
		url = url[7:]
	}
	return url
}

func WriteToFile(output string, url string, resp string) {
	if output != "" {
		f, err := os.OpenFile(output, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}
		if _, err := f.WriteString(url + "--" + resp + "\n"); err != nil {
			log.Println(err)
		}
	}
}

func GetIPFromURL(url string, verbose bool) string {
	resp, _ := net.LookupIP(url)

	if len(resp) > 0 {
		if verbose {
			fmt.Printf("%s -- %v\n", url, resp[0])
		}
		return resp[0].String()
	}
	return ""
}

func ReadFile(path string, verbose bool, output string) {
	f, _ := os.Open(path)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := FilterURL(scanner.Text())
		resp := GetIPFromURL(line, verbose)
		if resp != "" {
			if output != "" {
				WriteToFile(output, line, resp)
			}
		}
	}
}
