package respparser

import (
	"bufio"
	"errors"
	"io"
	"strconv"

	"simple_redis.com/m/src/resp"
)

type RespReader struct {
	reader *bufio.Reader
}

func NewRespReader(rd io.Reader) *RespReader {
	return &RespReader{reader: bufio.NewReader(rd)}
}

func (r RespReader) Read() (resp.ParsedValue, error) {
	sign, err := r.reader.ReadByte()
	if err != nil {
		return resp.ParsedValue{}, err
	}

	switch sign {
	case '*':
		return r.ReadArray()
	case '$':
		return r.ReadBulk()
	default:
		return resp.ParsedValue{}, errors.New("NOT IMPLEMENTED")
	}
}

// $5\r\nhello\r\n
func (r RespReader) ReadBulk() (resp.ParsedValue, error) {
	val := resp.ParsedValue{}
	val.SetType(resp.BULK)

	size, err := r.ReadNumber()
	if err != nil {
		return val, err
	}

	line := make([]byte, size)

	_, err = r.reader.Read(line)
	if err != nil {
		return val, err
	}

	val.SetBulk(string(line))

	// read trailing \r\n - meaning we've read the whole string
	r.ReadLine()

	return val, nil
}

// *2\r\n$5\r\nhello\r\n$5\r\nworld\r\n
func (r RespReader) ReadArray() (resp.ParsedValue, error) {
	v := resp.ParsedValue{}
	v.SetType(resp.ARRAY)

	size, err := r.ReadNumber()
	if err != nil {
		return v, err
	}

	arr := make([]resp.ParsedValue, size)
	v.SetArray(arr)
	for i := 0; i < size; i++ {
		v, err := r.Read()
		if err != nil {
			return v, err
		}
		arr = v.GetArray()
		arr = append(arr, v)
		v.SetArray(arr)
	}

	return v, nil
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

		if len(line) >= 2 && line[len(line)-2] == resp.CL {
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
