package binary_pack

import (
	"strings"
	"strconv"
)

type BinaryPack struct {}

func (bp *BinaryPack) Pack(format []string, args ...interface{}) []byte {
	return []byte{}
}

func (bp *BinaryPack) UnPack(format []string, msg []byte) []interface{} {
	return make([]interface{}, 1)
}

func (bp *BinaryPack) CalcSize(format []string) int {
	var size int

	for _, f := range format {
		switch f {
		case "?":
			size = size + 1
		case "h", "H":
			size = size + 2
		case "i", "I", "f":
			size = size + 4
		case "Q", "d":
			size = size + 8
		default:
			if strings.Contains(f, "s") {
				n, _ := strconv.Atoi(strings.TrimRight(f, "s"))
				size = size + n
			}
		}
	}

	return size
}