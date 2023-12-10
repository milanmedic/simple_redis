package respmarshaler

import (
	"io"
	"strconv"

	"simple_redis.com/m/src/resp"
)

type Marshaler struct {
	writer io.Writer
}

func NewMarshaler(w io.Writer) *Marshaler {
	return &Marshaler{writer: w}
}

func (m Marshaler) Write(value resp.ParsedValue) error {
	var bytes = marshal(value)

	_, err := m.writer.Write(bytes)
	if err != nil {
		return err
	}

	return nil
}

func marshal(value resp.ParsedValue) []byte {

	switch value.GetType() {
	case resp.ARRAY:
		return marshalArr(value)
	case resp.BULK:
		return marshalBlk(value.GetBulk())
	case resp.STRING:
		return marshalStr(value.GetString())
	case resp.ERROR:
		return marshalError(value.GetError())
	case resp.NULL:
		return marshallNull()
	default:
		return nil
	}
}

func marshalBlk(value string) []byte {
	var line []byte
	line = append(line, []byte(resp.BULK)...)                //$
	line = append(line, []byte(strconv.Itoa(len(value)))...) // 5
	line = append(line, resp.CL+resp.RF)                     // \r\n
	line = append(line, value...)                            // stringVal
	line = append(line, resp.CL+resp.RF)                     // \r\n

	return line
}

func marshalArr(value resp.ParsedValue) []byte {
	var line []byte
	var arrLen string = strconv.Itoa(len(value.GetArray()))

	line = append(line, []byte(resp.ARRAY)...) // *
	line = append(line, []byte(arrLen)...)     // len
	line = append(line, resp.CL+resp.RF)       // \r\n

	for _, v := range value.GetArray() {
		line = append(line, marshal(v)...) // content
	}

	return line
}

func marshalStr(value string) []byte {
	var line []byte
	line = append(line, []byte(resp.STRING)...) // +
	line = append(line, value...)               // stringVal
	line = append(line, resp.CL+resp.RF)        // \r\n

	return line
}

func marshalError(value string) []byte {
	var bytes []byte
	bytes = append(bytes, []byte(resp.ERROR)...)
	bytes = append(bytes, value...)
	bytes = append(bytes, '\r', '\n')

	return bytes
}

func marshallNull() []byte {
	return []byte("$-1\r\n")
}
