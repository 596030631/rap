package core

var stx = []byte{STX, 0}
var end = []byte{ETX, EOT}

func Encoder(msg string) []byte {
	return []byte (msg)
}

func Pick() ([]byte, []byte) {
	return stx, end
}

func Decoder(b []byte) string {
	return string(b)
}
