package utils

import (
	"strconv"

	"github.com/dustin/go-humanize"
)

const MaxRawDifficulty = 1000

type Difficulty uint64

func (d Difficulty) String() string {
	if d < MaxRawDifficulty {
		return strconv.FormatUint(uint64(d), 10)
	} else {
		return humanize.SIWithDigits(float64(d), 2, "")
	}
}
