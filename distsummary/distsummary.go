// distsummary hashes the strings found on stdin and produces a distribution summary.
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"

	"github.com/mredivo/pearson"
)

func main() {

	var bucket [256]int
	var sum, populated int
	var freqdist = make(map[int]int)
	dots := ".      .      .      .      .      .      .      .      .      .      .      .      .      .      .      .      ."
	dot := "      ."

	// Read in the strings and hash them into buckets.
	p := pearson.New(nil)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		h8 := p.Hash(scanner.Text())
		bucket[int(h8)] = bucket[int(h8)] + 1
	}

	// Print a heat map of the distribution while calculating it.
	fmt.Printf(dots)
	for i, v := range bucket {
		sum += v
		if c, found := freqdist[v]; found {
			freqdist[v] = c + 1
		} else {
			freqdist[v] = 1
		}
		if i%16 == 0 {
			fmt.Printf("%s\n.", dot)
		}
		if v != 0 {
			fmt.Printf("%7d", v)
			populated += 1
		} else {
			fmt.Printf("%7s", "")
		}
	}
	fmt.Printf("%s\n", dot)
	fmt.Printf("%s%s\n", dots, dot)

	// Summarize the bucket distribution.
	fmt.Printf("Number of input strings: %7d\n", sum)
	fmt.Printf("Number of unique hashes: %7d  %6.2f %%\n", populated, float64(populated)*100.0/256.0)
	fmt.Printf("Empty buckets:           %7d  %6.2f %%\n", 256-populated, (float64(256-populated)*100.0)/256.0)
	fmt.Printf("Total buckets:           %7d  %6.2f %%\n", 256, 100.0)
	fmt.Printf("Frequency distribution:\n  Count  Occurences   Percent\n")
	var order = make([]int, 0)
	for k, _ := range freqdist {
		order = append(order, k)
	}
	sort.Ints(order)
	for _, i := range order {
		fmt.Printf("%7d     %7d  %6.2f %%\n", i, freqdist[i], (float64(freqdist[i])*100.0)/256.0)
	}
}
