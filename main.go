package main

import (
	crypto "crypto/rand"
	"encoding/binary"
	"errors"
	"flag"
	"fmt"
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
	flag.IntVar(&n, "n", 16, "number of password to generate")
	var length int
	flag.IntVar(&length, "length", 15, "set password length")
	var minUpperLength int
	flag.IntVar(&minUpperLength, "min-uppercase-length", 2, "min uppercase length")
	var minLowerLength int
	flag.IntVar(&minLowerLength, "min-lowercase-length", 2, "min lowercase length")
	var minSpecialLength int
	flag.IntVar(&minSpecialLength, "min-special-length", 2, "min special length")
	var minDigitLength int
	flag.IntVar(&minDigitLength, "min-digit-length", 2, "min digit length")
	flag.Parse()

	for i := 1; i <= n; i++ {
		password, err := generate(length, minUpperLength, minLowerLength, minSpecialLength, minDigitLength)
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

func generate(length, minUpperLength, minLowerLength, minSpecialLength, minDigitLength int) (string, error) {
	password := strings.Builder{}

	for i := 0; i < minLowerLength; i++ {
		random := rand.Intn(len(lowercase))
		password.WriteByte(lowercase[random])
	}

	for i := 0; i < minUpperLength; i++ {
		random := rand.Intn(len(uppercase))
		password.WriteByte(uppercase[random])
	}

	for i := 0; i < minSpecialLength; i++ {
		random := rand.Intn(len(special))
		password.WriteByte(special[random])
	}

	for i := 0; i < minDigitLength; i++ {
		random := rand.Intn(len(digit))
		password.WriteByte(digit[random])
	}

	remainingLen := length - minLowerLength - minUpperLength - minSpecialLength - minDigitLength

	if remainingLen < 0 {
		return "", errors.New("password length is to small to satisfy all constraint")
	}

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
