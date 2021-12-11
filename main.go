package main

import (
	crypto "crypto/rand"
	"encoding/binary"
	"flag"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strings"
)

const (
	lowercase = "abcdedfghijklmnopqrstuvwxyz"
	uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	special   = "!@#$%&*^"
	digit     = "0123456789"
)

var allCharSet = lowercase + uppercase + special + digit

func init() {
	var b [8]byte
	if _, err := crypto.Read(b[:]); err != nil {
		panic(fmt.Sprintf("cannot init seed: %s", err))
	}
	rand.Seed(int64(binary.LittleEndian.Uint64(b[:])))
}

func main() {
	var n int
	flag.IntVar(&n, "n", 16, "set number of password to generate")
	var length int
	flag.IntVar(&length, "len", 15, "set password length")
	flag.Parse()

	fmt.Printf("Generating %d passwords of length %d\n\n", n, length)

	for i := 1; i <= n; i++ {
		password, err := generate(length)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		if i%4 == 0 {
			fmt.Println(password)
		} else if i == n {
			fmt.Println(password)
		} else {
			fmt.Print(password, "  ")
		}
	}
}

func generate(length int) (string, error) {
	password := strings.Builder{}

	min := int(math.Round(float64(length) * 0.15))

	for i := 0; i < min; i++ {
		random := rand.Intn(len(lowercase))
		password.WriteByte(lowercase[random])
	}

	for i := 0; i < min; i++ {
		random := rand.Intn(len(uppercase))
		password.WriteByte(uppercase[random])
	}

	for i := 0; i < min; i++ {
		random := rand.Intn(len(special))
		password.WriteByte(special[random])
	}

	for i := 0; i < min; i++ {
		random := rand.Intn(len(digit))
		password.WriteByte(digit[random])
	}

	remainingLen := length - min*4

	for i := 0; i < remainingLen; i++ {
		random := rand.Intn(len(allCharSet))
		password.WriteByte(allCharSet[random])
	}

	inRune := []rune(password.String())
	rand.Shuffle(len(inRune), func(i, j int) {
		inRune[i], inRune[j] = inRune[j], inRune[i]
	})

	return string(inRune), nil
}
