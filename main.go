package main

import (
	crypto "crypto/rand"
	"encoding/binary"
	"flag"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

const (
	lowercase = "abcdedfghijklmnopqrstuvwxyz"
	uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	special   = "!@#$%&*^"
	digit     = "0123456789"
)

var (
	allCharSet       = lowercase + uppercase + special + digit
	noSpecialCharSet = lowercase + uppercase + digit
)

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
	flag.IntVar(&length, "len", 20, "set password length")
	var noSpecial bool
	flag.BoolVar(&noSpecial, "no-special", false, "no special charset")
	flag.Parse()

	if length <= 0 || n <= 0 {
		fmt.Fprintln(os.Stderr, "Come on! Let's be realistic!")
		os.Exit(1)
	}

	fmt.Printf("Generating %d passwords of length %d %s\n", n, length, unquoteCodePoint("\\U0001f680"))
	if noSpecial {
		fmt.Println("- special charset disable!")
	}
	fmt.Println()

	for i := 1; i <= n; i++ {
		password := generate(length, noSpecial)
		if i%4 == 0 {
			fmt.Println(password)
		} else if i == n {
			fmt.Println(password)
		} else {
			fmt.Print(password, "  ")
		}
	}
}

func generate(length int, noSpecCharSet bool) string {
	password := strings.Builder{}

	min := int(math.Round(float64(length) * 0.15))
	remainingLen := length - min*3

	for i := 0; i < min; i++ {
		random := rand.Intn(len(lowercase))
		password.WriteByte(lowercase[random])
	}

	for i := 0; i < min; i++ {
		random := rand.Intn(len(uppercase))
		password.WriteByte(uppercase[random])
	}

	for i := 0; i < min; i++ {
		random := rand.Intn(len(digit))
		password.WriteByte(digit[random])
	}

	if !noSpecCharSet {
		for i := 0; i < min; i++ {
			random := rand.Intn(len(special))
			password.WriteByte(special[random])
		}
		remainingLen -= min
	}

	for i := 0; i < remainingLen; i++ {
		if !noSpecCharSet {
			random := rand.Intn(len(allCharSet))
			password.WriteByte(allCharSet[random])
			continue
		}
		random := rand.Intn(len(noSpecialCharSet))
		password.WriteByte(noSpecialCharSet[random])
	}

	passwordRune := []rune(password.String())
	rand.Shuffle(len(passwordRune), func(i, j int) {
		passwordRune[i], passwordRune[j] = passwordRune[j], passwordRune[i]
	})

	return string(passwordRune)
}

func unquoteCodePoint(s string) string {
	r, err := strconv.ParseInt(strings.TrimPrefix(s, "\\U"), 16, 32)
	if err != nil {
		panic(err)
	}
	return string(r)
}
