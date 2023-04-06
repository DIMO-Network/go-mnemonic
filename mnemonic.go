package mnemonic

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// const ...
const (
	InvalidStrength = "Invalid strength"
	InvalidEntropy  = "Invalid entropy"
)

var (
	chunksRe = regexp.MustCompile("[01]{11}")
	WordList = strings.Split(WordlistEnglish, "\n")
)

func EntropyToMnemonic(entropy []byte) ([]string, error) {
	length := len(entropy)
	if length < 16 || length > 32 || length%4 != 0 {
		return nil, errors.New(InvalidEntropy)
	}
	entropyBits := BytesToBinary(entropy)
	checksumBits := DeriveChecksumBits(entropy)
	bits := entropyBits + checksumBits
	chunks := chunksRe.FindAllString(bits, -1)
	words := []string{}
	for _, binary := range chunks {
		i, err := BinaryToInt(binary)
		if err != nil {
			return words, err
		}
		words = append(words, WordList[i])
	}
	return words, nil
}

func DeriveChecksumBits(bytes []byte) string {
	ENT := len(bytes) * 8
	CS := ENT / 32
	s := sha256.New()
	s.Write(bytes)
	hash := s.Sum(nil)
	return BytesToBinary(hash)[:CS]
}

func BytesToBinary(bytes []byte) string {
	bits := ""
	for _, b := range bytes {
		bits += fmt.Sprintf("%08b", b)
	}
	return bits
}

func BinaryToInt(bin string) (int64, error) {
	return strconv.ParseInt(bin, 2, 16)
}
