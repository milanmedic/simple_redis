package respparser

import (
	"bufio"
	"errors"
	"io"
	"strconv"
)

const (
	CL = '\r'
	RF = '\n'
)

type RESP_TYPE string

const ( // most common RESP types
	ARRAY   RESP_TYPE = "*"
	BULK    RESP_TYPE = "$"
	INT     RESP_TYPE = ":"
	BOOLEAN RESP_TYPE = "#"
	ERROR   RESP_TYPE = "-"
	STRING  RESP_TYPE = "+"
)

type ParsedValue struct {
	_type  RESP_TYPE
	_bulk  string
	_array []ParsedValue
}

type RespReader struct {
	reader *bufio.Reader
}

func NewRespReader(rd io.Reader) *RespReader {
	return &RespReader{reader: bufio.NewReader(rd)}
}

func (r RespReader) Read() (ParsedValue, error) {
	sign, err := r.reader.ReadByte()
	if err != nil {
		return ParsedValue{}, err
	}

	switch sign {
	case '*':
		return r.ReadArray()
	case '$':
		return r.ReadBulk()
	default:
		return ParsedValue{}, errors.New("NOT IMPLEMENTED")
	}
}

// $5\r\nhello\r\n
func (r RespReader) ReadBulk() (ParsedValue, error) {
	val := ParsedValue{}
	val._type = BULK

	size, err := r.ReadNumber()
	if err != nil {
		return val, err
	}

	line := make([]byte, size)

	_, err = r.reader.Read(line)
	if err != nil {
		return val, err
	}

	val._bulk = string(line)

	// read trailing \r\n - meaning we've read the whole string
	r.ReadLine()

	return val, nil
}

// *2\r\n$5\r\nhello\r\n$5\r\nworld\r\n
func (r RespReader) ReadArray() (ParsedValue, error) {
	val := ParsedValue{}
	val._type = ARRAY
	size, err := r.ReadNumber()
	if err != nil {
		return val, err
	}

	val._array = make([]ParsedValue, size)
	for i := 0; i < size; i++ {
		v, err := r.Read()
		if err != nil {
			return v, err
		}

		v._array = append(v._array, v)
	}

	return val, nil
}

/*
Reading
*2 \r\n - line1
$5 \r\n - line2
hello \r\n - line3
$5 \r\n - line4
world \r\n - line5
*/
func (r RespReader) ReadLine() (line []byte, err error) {
	for {
		b, err := r.reader.ReadByte()
		if err != nil {
			return nil, err
		}

		line = append(line, b)

		if len(line) >= 2 && line[len(line)-2] == CL {
			break
		}
	}

	return line[:len(line)-2], nil
}

func (r RespReader) ReadNumber() (int, error) {
	line, err := r.ReadLine()
	if err != nil {
		return 0, nil
	}

	num, err := strconv.ParseInt(string(line), 10, 64)
	if err != nil {
		return 0, nil
	}

	return int(num), nil
}
