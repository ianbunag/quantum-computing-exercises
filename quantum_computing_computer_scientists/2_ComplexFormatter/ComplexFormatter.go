package ComplexFormatter

import "strconv"

type ComplexFormatter struct {
	fmt     byte
	prec    int
	bitSize int
}

func New(fmt byte, prec int) ComplexFormatter {
	return ComplexFormatter{
		fmt:     fmt,
		prec:    prec,
		bitSize: 128,
	}
}

func (formatter ComplexFormatter) Format(complexNumber complex128) string {
	return strconv.FormatComplex(
		complexNumber,
		formatter.fmt,
		formatter.prec,
		formatter.bitSize,
	)
}
