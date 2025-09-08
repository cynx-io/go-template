package response

import (
	core "github.com/cynx-io/cynx-core/proto/gen"
	"github.com/cynx-io/cynx-core/src/response"
)

func setResponse[Resp response.Generic](resp Resp, code response.Code) {
	resp.GetBase().Code = code.String()
	resp.GetBase().Desc = responseCodeNames[code]
}

func Success[Resp response.Generic](resp Resp) {
	setResponse(resp, CodeSuccess)
	resp.GetBase().Status = core.Status_SUCCESS
}

func ErrorValidation[Resp response.Generic](resp Resp) {
	setResponse(resp, codeValidationError)
}

func ErrorUnauthorized[Resp response.Generic](resp Resp) {
	setResponse(resp, codeUnauthorized)
}

func ErrorNeedAuth[Resp response.Generic](resp Resp) {
	setResponse(resp, codeNeedAuth)
	resp.GetBase().Status = core.Status_NEED_AUTH
}

func ErrorNotAllowed[Resp response.Generic](resp Resp) {
	setResponse(resp, codeNotAllowed)
}

func ErrorNotFound[Resp response.Generic](resp Resp) {
	setResponse(resp, codeNotFound)
}

func ErrorInvalidCredentials[Resp response.Generic](resp Resp) {
	setResponse(resp, codeInvalidCredentials)
}

func ErrorInternalFailSchema[Resp response.Generic](resp Resp) {
	setResponse(resp, codeInternalFailSchema)
}

func ErrorInternalMarshal[Resp response.Generic](resp Resp) {
	setResponse(resp, codeInternalMarshal)
}

func ErrorInternalUnmarshal[Resp response.Generic](resp Resp) {
	setResponse(resp, codeInternalUnmarshal)
}

func ErrorInternalParse[Resp response.Generic](resp Resp) {
	setResponse(resp, codeInternalParse)
}

func ErrorMicroservicePlutusPayment[Resp response.Generic](resp Resp) {
	setResponse(resp, codeMicroservicePlutusPayment)
}

func ErrorMicroserviceHermesUser[Resp response.Generic](resp Resp) {
	setResponse(resp, codeMicroserviceHermesUser)
}

func ErrorPaymentRequired[Resp response.Generic](resp Resp) {
	setResponse(resp, codePaymentRequired)
	resp.GetBase().Status = core.Status_NEED_TOKEN
}

func ErrorPaymentTokenInsufficient[Resp response.Generic](resp Resp) {
	setResponse(resp, codePaymentTokenInsufficient)
	resp.GetBase().Status = core.Status_INSUFFICIENT_TOKEN
}

func ErrorInternalExtractionData[Resp response.Generic](resp Resp) {
	setResponse(resp, codeInternalExtractionData)
}
