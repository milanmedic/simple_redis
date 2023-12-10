package resp

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
	NULL    RESP_TYPE = "_"
)

type ParsedValue struct {
	_type   RESP_TYPE
	_string string
	_error  string
	_bulk   string
	_array  []ParsedValue
}

func (pv *ParsedValue) SetType(t RESP_TYPE) {
	pv._type = t
}

func (pv *ParsedValue) GetType() RESP_TYPE {
	return pv._type
}

func (pv *ParsedValue) SetString(s string) {
	pv._string = s
}

func (pv *ParsedValue) GetString() string {
	return pv._string
}

func (pv *ParsedValue) SetError(e error) {
	pv._error = e.Error()
}

func (pv *ParsedValue) GetError() string {
	return pv._error
}

func (pv *ParsedValue) SetBulk(blk string) {
	pv._bulk = blk
}

func (pv *ParsedValue) GetBulk() string {
	return pv._bulk
}

func (pv *ParsedValue) SetArray(arr []ParsedValue) {
	pv._array = arr
}

func (pv *ParsedValue) GetArray() []ParsedValue {
	return pv._array
}
