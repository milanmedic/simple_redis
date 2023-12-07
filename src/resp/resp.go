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
)

type ParsedValue struct {
	_type  RESP_TYPE
	_bulk  string
	_array []ParsedValue
}

func (pv *ParsedValue) SetType(t RESP_TYPE) {
	pv._type = t
}

func (pv *ParsedValue) SetBulk(blk string) {
	pv._bulk = blk
}

func (pv *ParsedValue) SetArray(arr []ParsedValue) {
	pv._array = arr
}

func (pv *ParsedValue) GetArray() []ParsedValue {
	return pv._array
}
