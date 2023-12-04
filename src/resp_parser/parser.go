package respparser

// TODO: IMPLEMENT

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
	_type    RESP_TYPE
	_bulk    string
	_string  string
	_error   string
	_int     int
	_boolean bool
}
