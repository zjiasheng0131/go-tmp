package api

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func HandleSuccess(ctx *gin.Context, data interface{}) {
	if data == nil {
		data = map[string]interface{}{}
	}
	resp := Response{Code: errorCodeMap[ErrSuccess], Message: ErrSuccess.Error(), Data: data}
	if _, ok := errorCodeMap[ErrSuccess]; !ok {
		resp = Response{Code: http.StatusOK, Message: "", Data: data}
	}
	ctx.JSON(http.StatusOK, resp)
}

func HandleError(ctx *gin.Context, httpCode int, err error, data interface{}) {
	if data == nil {
		data = map[string]string{}
	}
	resp := Response{Code: errorCodeMap[err], Message: err.Error(), Data: data}
	if _, ok := errorCodeMap[ErrSuccess]; !ok {
		resp = Response{Code: 500, Message: "unknown error", Data: data}
	}
	ctx.JSON(httpCode, resp)
}

type Error struct {
	Code    int
	Message string
}

var errorCodeMap = map[error]int{}

func newError(code int, msg string) error {
	err := errors.New(msg)
	errorCodeMap[err] = code
	return err
}
func (e Error) Error() string {
	return e.Message
}

var (
	// common errors
	ErrSuccess             = newError(200, "ok")
	ErrBadRequest          = newError(400, "Bad Request")
	ErrUnauthorized        = newError(401, "Unauthorized")
	ErrNotFound            = newError(404, "Not Found")
	ErrInternalServerError = newError(500, "Internal Server Error")

	// more biz errors
	ErrEmailAlreadyUse              = newError(1001, "The email is already in use.")
	DuplicatedRequest               = newError(1002, "Sent duplicated request.")
	DeleteCreatingVM                = newError(1003, "Can't delete VM which is under creation")
	DuplicateVMDeleteRequest        = newError(1004, "Duplicate deletion request for the same VM.")
	AlreadyDeletedVM                = newError(1005, "The VM has already been deleted.")
	ErrVMDoesNotExist               = newError(1006, "VM does not exist.")
	TemplateDoesNotExist            = newError(1007, "Template doesn't exist in cluster.")
	TaskWithoutCluster              = newError(1008, "The task doesn't have cluster assigned.")
	ErrNetworkDoesNotExist          = newError(1009, "Network not found.")
	ErrImageDoesNotExist            = newError(1010, "Image not found.")
	ErrSameTemplateExist            = newError(1011, "Another image with the same TemplateId already exists.")
	ErrBaseClusterNotExist          = newError(1012, "Base cluster not found for template transfer.")
	ErrTemplateSourceIPNotFound     = newError(1013, "The src ip is not found for template or template doesn't exist.")
	ErrImageExistsWithSameOSVersion = newError(1014, "Another image with the same OSVersion and OSTemplateVersion already exists.")

	ValInvailidCharacter      = newError(2001, "Input contains invalid characters")
	ValInvailidEnder          = newError(2002, "Input cannot end with a hyphen or period")
	ValInvailidadjacent       = newError(2003, "Hyphens and periods cannot be adjacent")
	ValInvailidFortiCSPPwdLen = newError(2004, "FortiCSP password must be at least 8 characters long")
	ValInvailidWindowsPwd     = newError(2005, "Windows password must contain characters from three of the "+
		"following four categories: "+
		"1. Uppercase letters (A-Z) "+
		"2. Lowercase letters (a-z) "+
		"3. Digits (0-9)"+
		"4. Non-alphanumeric characters (e.g., !, $, #, %)")
	ValInvalidNetNameCharacter      = newError(2006, "Network name should only contain characters as alphanumdashdot")
	ValInvalidDHCPRange             = newError(2007, "DHCP range must be in the format 'ip1-ip2' with valid IP addresses, separated by commas if multiple ranges")
	ValInvalidVlanRange             = newError(2008, "VLAN must be between 0 and 4095")
	ValFTPFileNotExist              = newError(2009, "The ftp file for specified version of FortiCSP/FortiCSPManager/FortiGate doesn't exist")
	ValInvalidTemplateIdRange       = newError(2010, "Auto template Id should be 8000-8999, other template Id should be 9000-9999")
	ValInvalidOSVersionCharacter    = newError(2011, "The OS version should only contain alphanum dash or dot")
	ValInvalidOSTplVersionCharacter = newError(2012, "The OS template version should only contain alphanum dash or dot")
	ValInvalidAlphanumdashCharacter = newError(2013, "The field should only contain characters as alphanumdash")
	ValImgNotExist                  = newError(2014, "The image for specified version doesn't exist")
	ValInvalidPathFormat            = newError(2015, "Invalid format: expected user@host:path")
	ValFailedDialSSH                = newError(2016, "Failed to dial SSH")
	ValFailedCreateSFTPClient       = newError(2017, "Failed to create SFTP client")
	ValFailedStatRemoteFile         = newError(2018, "Failed to stat remote file")
)
