package main

import (
	crypto "crypto/rand"
	"flag"
	"fmt"
	"math"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"

	"github.com/nbutton23/zxcvbn-go"
)

const (
	lowercase = "abcdedfghijklmnopqrstuvwxyz"
	uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	special   = "!@#$%&*^{}[]()-;:><_+,.="
	digit     = "0123456789"
)

var (
	allCharSet       = lowercase + uppercase + special + digit
	noSpecialCharSet = lowercase + uppercase + digit
)

var r *rand.Rand

func init() {
	var b [32]byte
	if _, err := crypto.Read(b[:]); err != nil {
		panic(fmt.Sprintf("cannot init seed: %s", err))
	}
	r = rand.New(rand.NewChaCha8(b))
}

func main() {
	var n int
	flag.IntVar(&n, "n", 16, "set the number of passwords to generate")
	var length int
	flag.IntVar(&length, "len", 20, "set passwords length")
	var noSpecial bool
	flag.BoolVar(&noSpecial, "no-special", false, "disable special charset")
	var col int
	flag.IntVar(&col, "col", 4, "set the number of column on which display passwords")
	flag.Parse()

	if length <= 0 || n <= 0 || col <= 0 {
		_, _ = fmt.Fprintln(os.Stderr, "Come on! Let's be realistic!")
		os.Exit(1)
	}

	fmt.Printf("Generating %d passwords of length %d %s\n", n, length, unquoteCodePoint("\\U0001f680"))
	if noSpecial {
		fmt.Println("- special charset disable!")
	}
	fmt.Println()

	for i := 1; i <= n; i++ {
		password := generate(length, noSpecial)
		if i%col == 0 || i == n {
			fmt.Printf("%s (%.2f)\n", password, entropy(password))
			continue
		}
		fmt.Printf("%s (%.2f)\t", password, entropy(password))
	}
	fmt.Println()
}

func generate(length int, noSpecCharSet bool) string {
	password := strings.Builder{}

	min := int(math.Round(float64(length) * 0.15))
	remainingLen := length - min*3

	for i := 0; i < min; i++ {
		random := r.IntN(len(lowercase))
		password.WriteByte(lowercase[random])
	}

	for i := 0; i < min; i++ {
		random := r.IntN(len(uppercase))
		password.WriteByte(uppercase[random])
	}

	for i := 0; i < min; i++ {
		random := r.IntN(len(digit))
		password.WriteByte(digit[random])
	}

	if !noSpecCharSet {
		for i := 0; i < min; i++ {
			random := r.IntN(len(special))
			password.WriteByte(special[random])
		}
		remainingLen -= min
	}

	for i := 0; i < remainingLen; i++ {
		if !noSpecCharSet {
			random := r.IntN(len(allCharSet))
			password.WriteByte(allCharSet[random])
			continue
		}
		random := r.IntN(len(noSpecialCharSet))
		password.WriteByte(noSpecialCharSet[random])
	}

	passwordRune := []rune(password.String())
	r.Shuffle(len(passwordRune), func(i, j int) {
		passwordRune[i], passwordRune[j] = passwordRune[j], passwordRune[i]
	})

	return string(passwordRune)
}

func unquoteCodePoint(s string) string {
	i, err := strconv.ParseInt(strings.TrimPrefix(s, "\\U"), 16, 32)
	if err != nil {
		panic(err)
	}
	return string(i)
}

func entropy(s string) (bits float64) {
	e := zxcvbn.PasswordStrength(s, nil)
	return e.Entropy
}
