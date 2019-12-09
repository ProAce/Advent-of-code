package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()

	input := "bgvyzdsv"

	for i := 0; i < math.MaxInt64; i++ {
		hash := md5.Sum([]byte(input + strconv.Itoa(i)))
		hashByte := []byte{}

		for _, value := range hash {
			hashByte = append(hashByte, value)
		}

		value := hex.EncodeToString(hashByte)

		if strings.HasPrefix(value, "000000") { // Hash has leading zeroes
			fmt.Println(i)
			break
		}
	}

	fmt.Println(time.Since(start))
}
