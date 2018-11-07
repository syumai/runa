package main

import (
	"io"
	"unicode/utf8"
)

type RangeStream struct {
	from, to, curr int64
}

func NewRangeStream(from, to int64) *RangeStream {
	return &RangeStream{
		from: from,
		to:   to,
		curr: from,
	}
}

func (rs *RangeStream) ReadRune() (r rune, size int, err error) {
	if rs.curr > rs.to {
		return 0, 0, io.EOF
	}

	r = rune(rs.curr)
	size = utf8.RuneLen(r)
	rs.curr++
	return
}
