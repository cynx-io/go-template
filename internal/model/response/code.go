package response

import "github.com/cynx-io/cynx-core/src/response"

const (
	// Expected Error
	CodeSuccess                  response.Code = "00"
	codeValidationError          response.Code = "VE"
	codeUnauthorized             response.Code = "UA"
	codeNotAllowed               response.Code = "NA"
	codeNotFound                 response.Code = "NF"
	codeInvalidCredentials       response.Code = "IC"
	codeNeedAuth                 response.Code = "NU"
	codePaymentRequired          response.Code = "PR"
	codePaymentTokenInsufficient response.Code = "TI"

	// [I] Internal
	codeInternalFailSchema     response.Code = "I-FS"
	codeInternalMarshal        response.Code = "I-MA"
	codeInternalUnmarshal      response.Code = "I-UM"
	codeInternalParse          response.Code = "I-PA"
	codeInternalExtractionData response.Code = "I-ED"

	// [E] External Errors

	// [M] Microservice Errors
	codeMicroservicePlutusPayment response.Code = "M-PP"
	codeMicroserviceHermesUser    response.Code = "M-HU"

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
	codeNeedAuth:           "Need Authentication",

	// Internal
	codeInternalFailSchema:     "Fail Schema",
	codeInternalMarshal:        "Marshal Error",
	codeInternalUnmarshal:      "Unmarshal Error",
	codeInternalParse:          "Parse Error",
	codeInternalExtractionData: "Extraction Data Error",

	// External Errors

	// Microservice Errors
	codeMicroservicePlutusPayment: "Plutus Payment Service Error",
	codeMicroserviceHermesUser:    "Hermes User Service Error",

	// Database Errors
}
