package response

const (
	ErrCodeSuccess = 20001
	ErrCodeFail    = 40001
	ErrAuthFail    = 40002
	ErrOTPInvalid  = 40003
	ErrorSendEmailFail = 40004
	ErrCodeUserHasExists = 50001 // user has already exists
	ErrCodeParamError = 50002
)
var msg = map[int]string{
	ErrCodeSuccess: "success",
	ErrCodeFail:    "fail",
	ErrAuthFail:    "auth fail",
	ErrCodeUserHasExists: "user has already exists",
	ErrOTPInvalid : "otp invalid",
	ErrorSendEmailFail: "send email fail",
	ErrCodeParamError: "param error",
}