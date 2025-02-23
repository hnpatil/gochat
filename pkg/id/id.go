package id

import (
	"crypto/sha256"
	"fmt"
	"sort"
	"strings"
)

type ID string

func New(ids ...string) ID {
	sort.Strings(ids)

	hash := sha256.Sum256([]byte(strings.Join(ids, "-")))
	bytes := make([][]byte, 4)

	for i := 0; i < len(bytes); i++ {
		s := i * 4
		if i == 1 || i == 2 {
			bytes[i] = hash[s : s+2]
		} else {
			bytes[i] = hash[s : s+4]
		}
	}

	return ID(fmt.Sprintf("%x-%x-%x-%x", bytes[0], bytes[1], bytes[2], bytes[3]))
}
