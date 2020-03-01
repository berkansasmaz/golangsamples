package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the sockMerchant function below.
func sockMerchant(n int32, ar []int32) int32 {
	var i, j int32
	var count, oddEven int32 = 0, 0
	for i = 0; i < n; i++ {
		for j = 0; j < n; j++ {
			fmt.Println("disi", "i:", i, "j:", j)
			if i == j {
			} else if ar[i] == ar[j] {
				fmt.Println("içi", "i:", i, "j:", j)
				oddEven++
				fmt.Println("oddEven: ", oddEven)
				if i > 0 {
					z := i - 1
					for z >= 0 {
						if ar[z] == ar[i] {
							fmt.Println("continue: ")
							if i != n-1 {
								i++
								j = 0
								oddEven = 0
							} else {
								oddEven = 0
								break
							}
						}
						z--
					}
				}
			}
			if j == n-1 {
				oddEven++
				count += oddEven / 2
				fmt.Println("oddEven: ", oddEven)
				oddEven = 0
				fmt.Println("count: ", count)
			}
		}
	}
	return count
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	arTemp := strings.Split(readLine(reader), " ")

	var ar []int32

	for i := 0; i < int(n); i++ {
		arItemTemp, err := strconv.ParseInt(arTemp[i], 10, 64)
		checkError(err)
		arItem := int32(arItemTemp)
		ar = append(ar, arItem)
	}

	result := sockMerchant(n, ar)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
