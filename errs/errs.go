package errs

import "strings"

type Err string

const (
	SUCCESS          Err = "0000000"
	FAILD            Err = "4000000"
	INVALID_INPUT    Err = "4001001"
	WS_CONNECT_FAILD Err = "5001001"
	DB_OPERATR_ERROR Err = "5001002"
	Permission_Deny  Err = "4011001"
)

var errsMap map[Err]string

func init() {
	errsMap = map[Err]string{
		SUCCESS:          "success .",
		FAILD:            "faild .",
		INVALID_INPUT:    "invalid input data .",
		WS_CONNECT_FAILD: "new websocket connect faild .",
		DB_OPERATR_ERROR: "database operate error .",
		Permission_Deny:  "permission dent .",
	}

}

func (e Err) Error() string {
	return errsMap[e]
}

type ErrorComplex struct {
	StatusCode int
	Msg        string
}

func (e ErrorComplex) Error() string {
	return e.Msg
}
func NewComplexError(statusCode int, er error, extend ... string) (err ErrorComplex) {
	var msg string
	msg = strings.Join(extend, msg+er.Error())
	return ErrorComplex{StatusCode: statusCode,
		Msg: msg}
}
