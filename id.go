package types

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"strings"
	"time"
)

const prettyTableSize = 34

var prettyTable = [prettyTableSize]byte{
	'1', '2', '3', '4', '5', '6', '7', '8', '9',
	'A', 'B', 'C', 'D', 'E', 'F', 'G',
	'H', 'I', 'J', 'K', 'L', 'M', 'N',
	'P', 'Q',
	'R', 'S', 'T',
	'U', 'V', 'W',
	'X', 'Y', 'Z'}

type IDFormat int

const (
	ShortIDFormat IDFormat = iota
	PrettyIDFormat
)

type ID int64

// Int converts ID into int64. Just make it easier to edit code
func (i ID) Int() int64 {
	return int64(i)
}

// Short returns a short representation of id
func (i ID) Short() string {
	if i < 0 {
		panic("invalid id")
	}
	var bytes [16]byte
	k := int64(i)
	n := 15
	for {
		j := k % 62
		switch {
		case j <= 9:
			bytes[n] = byte('0' + j)
		case j <= 35:
			bytes[n] = byte('A' + j - 10)
		default:
			bytes[n] = byte('a' + j - 36)
		}
		k /= 62
		if k == 0 {
			return string(bytes[n:])
		}
		n--
	}
}

// Pretty returns a incasesensitive pretty representation of id
func (i ID) Pretty() string {
	if i < 0 {
		panic("invalid id")
	}
	var bytes [16]byte
	k := int64(i)
	n := 15

	for {
		bytes[n] = prettyTable[k%prettyTableSize]
		k /= prettyTableSize
		if k == 0 {
			return string(bytes[n:])
		}
		n--
	}
}

func NewIDFromString(s string, f IDFormat) (ID, error) {
	switch f {
	case ShortIDFormat:
		return parseShortID(s)
	case PrettyIDFormat:
		return parsePrettyID(s)
	default:
		return 0, errors.New("invalid format")
	}
}

func (i ID) Salt(v string) string {
	sum := md5.Sum([]byte(fmt.Sprintf("%s%d", v, i)))
	return hex.EncodeToString(sum[:])
}

func parseShortID(s string) (ID, error) {
	if len(s) == 0 {
		return 0, errors.New("parse error")
	}

	var bytes = []byte(s)
	var k int64
	var v int64
	for _, b := range bytes {
		switch {
		case b >= '0' && b <= '9':
			v = int64(b - '0')
		case b >= 'A' && b <= 'Z':
			v = int64(10 + b - 'A')
		case b >= 'a' && b <= 'z':
			v = int64(36 + b - 'a')
		default:
			return 0, errors.New("parse error")
		}
		k = k*62 + v
	}
	return ID(k), nil
}

func parsePrettyID(s string) (ID, error) {
	if len(s) == 0 {
		return 0, errors.New("parse error")
	}

	s = strings.ToUpper(s)
	var bytes = []byte(s)
	var k int64
	for _, b := range bytes {
		i := searchPrettyTable(b)
		if i <= 0 {
			return 0, errors.New("parse error")
		}
		k = k*prettyTableSize + int64(i)
	}
	return ID(k), nil
}

func searchPrettyTable(v byte) int {
	left := 0
	right := prettyTableSize - 1
	for right >= left {
		mid := (left + right) / 2
		if prettyTable[mid] == v {
			return mid
		} else if prettyTable[mid] > v {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return -1
}

// ------------------------------
// IDGenerator
type IDGenerator interface {
	NextID() ID
}

type NumberGetter interface {
	GetNumber() int64
}

type SnakeIDGenerator struct {
	seqSize   uint
	shardSize uint

	clock    NumberGetter
	sharding NumberGetter

	counter Counter
}

func NewSnakeIDGenerator(shardSize, seqSize uint, clock, sharding NumberGetter) *SnakeIDGenerator {
	if seqSize < 1 || seqSize > 16 {
		panic("seqSize should be [1,16]")
	}

	if clock == nil {
		panic("clock is nil")
	}

	if shardSize > 8 {
		panic("shardSize should be [0,8]")
	}

	if shardSize > 0 && sharding == nil {
		panic("sharding is nil")
	}

	if shardSize+seqSize >= 20 {
		panic("shardSize + seqSize should be less than 20")
	}

	return &SnakeIDGenerator{
		seqSize:   seqSize,
		shardSize: shardSize,
		clock:     clock,
		sharding:  sharding,
	}
}

func (g *SnakeIDGenerator) NextID() ID {
	id := g.clock.GetNumber() << (g.seqSize + g.shardSize)
	if g.shardSize > 0 {
		id |= (g.sharding.GetNumber() % (1 << g.shardSize)) << g.seqSize
	}
	id |= g.counter.Next() % (1 << g.seqSize)
	return ID(id)
}

// JSON中若没有指定类型，number默认解析成double，double整数部分最大值为2^53，因此控制在53bit内比较好
// id由time+shard+seq组成
// 若业务多可扩充shard，并发高可扩充seq. 由于time在最高位,故扩展后的id集合与原id集合不会出现交集,可保持全局唯一

const DefaultShardBitSize = 0 // 单机版本
const DefaultSeqBitSize = 6   // 每个shard每ms不能超过64次调用

var epoch time.Time
var DefaultIDGenerator IDGenerator

func init() {
	epoch = time.Date(2019, time.January, 2, 15, 4, 5, 0, time.UTC)
	DefaultIDGenerator = NewSnakeIDGenerator(DefaultShardBitSize, DefaultSeqBitSize, nextMilliseconds, nil)
}

type numberGetterFunc func() int64

func (f numberGetterFunc) GetNumber() int64 {
	return f()
}

var nextMilliseconds numberGetterFunc = func() int64 {
	return time.Since(epoch).Nanoseconds() / 1e6
}

func NewID() ID {
	return DefaultIDGenerator.NextID()
}
