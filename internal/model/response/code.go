package response

import "github.com/cynx-io/cynx-core/src/response"

const (
	// Expected Error
	CodeSuccess            response.Code = "00"
	codeValidationError    response.Code = "VE"
	codeUnauthorized       response.Code = "UA"
	codeNotAllowed         response.Code = "NA"
	codeNotFound           response.Code = "NF"
	codeInvalidCredentials response.Code = "IC"

	// [I] Internal
	codeInternalError response.Code = "I-IE"

	// [E] External Errors

	// [M] Microservice Errors

	// [D] Database Errors

)

var responseCodeNames = map[response.Code]string{
	// Expected Error
	CodeSuccess:            "Success",
	codeValidationError:    "Validation Error",
	codeUnauthorized:       "Not Authorized",
	codeNotAllowed:         "Not Allowed",
	codeNotFound:           "Not Found",
	codeInvalidCredentials: "Invalid Credentials",

	// Internal
	codeInternalError: "Internal Error",

	// External Errors

	// Microservice Errors

	// Database Errors

}
