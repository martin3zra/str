package str

import "crypto/rand"

// Random generate a alphanumeric string with the given length
func Random(times int) string {
	return string(randASCIIBytes(times))
}

func randASCIIBytes(n int) []byte {

	const encodeURL = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_"
	output := make([]byte, n)

	// We will take n bytes, one byte for each character of output.
	randomness := make([]byte, n)

	// read all random
	_, err := rand.Read(randomness)
	if err != nil {
		panic(err)
	}

	// fill output
	for pos := range output {
		// get random item
		random := uint8(randomness[pos])

		// random % 64
		randomPos := random % uint8(len(encodeURL))

		// put into output
		output[pos] = encodeURL[randomPos]
	}

	return output
}
