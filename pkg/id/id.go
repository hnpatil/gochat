package id

import (
	"fmt"
	"time"
)

type ID string

func New() ID {
	// Convert to hexadecimal and ensure it is 8 characters long by keeping only the last 32 bits
	return ID(fmt.Sprintf("%08x", time.Now().UnixMilli()&0xFFFFFFFF))
}

func (id ID) String() string {
	return string(id)
}
