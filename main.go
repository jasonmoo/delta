package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"regexp"
	"strconv"
	"strings"
)

var (
	decompress = flag.Bool("d", false, "compress the range supplied")

	decompressRegexp = regexp.MustCompile(`^(-?\d+)(?:-(-?\d+))?$`)
)

func init() {
	flag.Parse()
}

func main() {

	s := bufio.NewScanner(os.Stdin)

	switch {

	case *decompress:

		stdout := bufio.NewWriter(os.Stdout)
		defer stdout.Flush()

		for s.Scan() {
			line := strings.TrimSpace(s.Text())
			if len(line) == 0 {
				continue
			}
			matches := decompressRegexp.FindStringSubmatch(line)
			// no match -> passthru
			if len(matches) == 0 {
				stdout.WriteString(line)
				stdout.WriteByte('\n')
				continue
			}
			// single value -> passthru
			if matches[2] == "" {
				stdout.WriteString(matches[1])
				stdout.WriteByte('\n')
				continue
			}
			start, err := strconv.ParseInt(matches[1], 10, 64)
			if err != nil {
				log.Println(err)
				continue
			}
			end, err := strconv.ParseInt(matches[2], 0, 64)
			if err != nil {
				log.Println(err)
				continue
			}
			for start <= end {
				stdout.WriteString(strconv.FormatInt(start, 10))
				stdout.WriteByte('\n')
				start++
			}
		}

	default:

		var (
			range_start, range_end int64
			capturing              bool

			final_output = func() {
				if capturing {
					if range_start == range_end {
						fmt.Print(strconv.FormatInt(range_start, 10), "\n")
					} else {
						fmt.Print(strconv.FormatInt(range_start, 10), "-", strconv.FormatInt(range_end, 10), "\n")
					}
				}
			}
			c    = make(chan os.Signal, 1)
			done = make(chan struct{})
		)

		signal.Notify(c, os.Interrupt)
		go func() {
			<-c // falls through on close
			final_output()
			done <- struct{}{}
		}()

		for s.Scan() {

			val, err := strconv.ParseInt(strings.TrimSpace(s.Text()), 0, 64)
			if err != nil {
				fmt.Print(s.Bytes(), "\n")
				break
			}

			if !capturing {
				range_start = val
				range_end = val
				capturing = true
			} else if val == range_end+1 {
				range_end = val
			} else if range_start == range_end {
				fmt.Print(strconv.FormatInt(range_start, 10), "\n")
				range_start = val
				range_end = val
			} else {
				fmt.Print(strconv.FormatInt(range_start, 10), "-", strconv.FormatInt(range_end, 10), "\n")
				range_start = val
				range_end = val
			}

		}

		close(c)
		<-done

	}

}
