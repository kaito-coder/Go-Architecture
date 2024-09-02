package response

const (
	ErrCodeSuccess = 20001
	ErrCodeFail    = 40001
	ErrAuthFail    = 40002
)
var msg = map[int]string{
	ErrCodeSuccess: "success",
	ErrCodeFail:    "fail",
	ErrAuthFail:    "auth fail",
}