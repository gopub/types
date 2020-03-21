package types

import "fmt"

type ByteUnit int64

const (
	_           = iota
	KB ByteUnit = 1 << (10 * iota) // 1 << (10*1)
	MB                             // 1 << (10*2)
	GB                             // 1 << (10*3)
	TB                             // 1 << (10*4)
	PB                             // 1 << (10*5)
)

func (b ByteUnit) HumanReadable() string {
	if b < KB {
		return fmt.Sprintf("%d B", b)
	} else if b < MB {
		return fmt.Sprintf("%.2f KB", float64(b)/float64(KB))
	} else if b < GB {
		return fmt.Sprintf("%.2f MB", float64(b)/float64(MB))
	} else if b < TB {
		return fmt.Sprintf("%.2f GB", float64(b)/float64(GB))
	} else if b < PB {
		return fmt.Sprintf("%.2f TB", float64(b)/float64(TB))
	} else {
		return fmt.Sprintf("%.2f PB", float64(b)/float64(PB))
	}
}
