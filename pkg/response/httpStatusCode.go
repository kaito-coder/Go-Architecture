package response

const (
	ErrCodeSuccess = 20001
	ErrCodeFail    = 40001
	ErrAuthFail    = 40002
	ErrCodeUserHasExists = 50001 // user has already exists
)
var msg = map[int]string{
	ErrCodeSuccess: "success",
	ErrCodeFail:    "fail",
	ErrAuthFail:    "auth fail",
	ErrCodeUserHasExists: "user has already exists",
}