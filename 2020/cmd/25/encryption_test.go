package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLoopSize(t *testing.T) {
	cases := []struct {
		publicKey, expectedLoopSize int
	}{
		{5764801, 8},
		{17807724, 11},
	}

	for i, tt := range cases {
		t.Run(fmt.Sprintf("case %d (%d)", i+1, tt.publicKey), func(t *testing.T) {
			loopSize := getLoopSize(tt.publicKey)

			assert.Equal(t, tt.expectedLoopSize, loopSize)
		})
	}
}

func TestGetEncryptionKey(t *testing.T) {
	cases := []struct {
		cardPublicKey, doorPublicKey, expectedEncryptionKey int
	}{
		{5764801, 17807724, 14897079},
	}

	for i, tt := range cases {
		t.Run(fmt.Sprintf("case %d", i+1), func(t *testing.T) {
			cardLoopSize := getLoopSize(tt.cardPublicKey)
			doorLoopSize := getLoopSize(tt.doorPublicKey)

			encryptionKey1 := getEncryptionKey(tt.cardPublicKey, doorLoopSize)
			encryptionKey2 := getEncryptionKey(tt.doorPublicKey, cardLoopSize)

			assert.Equal(t, encryptionKey1, tt.expectedEncryptionKey)
			assert.Equal(t, encryptionKey2, tt.expectedEncryptionKey)
		})
	}
}
