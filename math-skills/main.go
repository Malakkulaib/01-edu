package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run stats.go <filename>")
		return
	}
	filePath := os.Args[1]

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	var data []float64
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		val, err := strconv.ParseFloat(line, 64)
		if err != nil {
			continue
		}
		data = append(data, val)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	if len(data) == 0 {
		fmt.Println("Error: File contains no valid data.")
		return
	}

	sum := 0.0
	for _, num := range data {
		sum += num
	}
	mean := sum / float64(len(data))

	sort.Float64s(data)
	var median float64
	n := len(data)
	if n%2 == 1 {
		median = data[n/2]
	} else {
		median = (data[n/2-1] + data[n/2]) / 2.0
	}

	varianceSum := 0.0
	for _, num := range data {
		diff := num - mean
		varianceSum += diff * diff
	}
	variance := varianceSum / float64(n)

	stdDev := math.Sqrt(variance)

	fmt.Printf("Average: %d\n", int(math.Round(mean)))
	fmt.Printf("Median: %d\n", int(math.Round(median)))
	fmt.Printf("Variance: %d\n", int(math.Round(variance)))
	fmt.Printf("Standard Deviation: %d\n", int(math.Round(stdDev)))
}
