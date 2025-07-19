package privy

import (
	"encoding/hex"
)

func DecodeCF(value string) string {
	// Decode the hex string to get the raw bytes containing key and encrypted data.
	data, err := hex.DecodeString(value)
	if err != nil || len(data) < 1 {
		return ""
	}

	// Extract the XOR key from the first byte of the decoded data.
	key := data[0]
	// Create a byte slice to hold the decrypted data (excluding the key byte).
	decoded := make([]byte, len(data)-1)

	// Iterate through the encrypted bytes (starting from index 1, after the key).
	for i := 1; i < len(data); i++ {
		// XOR each encrypted byte with the key to decrypt it.
		decoded[i-1] = data[i] ^ key
	}

	// Convert the decrypted bytes back to a string and return.
	return string(decoded)
}
