package errors

import "errors"

var (
	ErrUnableToMigrate                = errors.New("Unable to migrate")
	ErrUnableToGetUrl                 = errors.New("Unable to get url")
	ErrRecordNotFound                 = errors.New("record not found")
	ErrUnknown                        = errors.New("unknown error")
	ErrPasswordEncryption             = errors.New("could not encrypt password")
	ErrUnableToSave                   = errors.New("unable to save data")
	ErrForgotEmail                    = errors.New("email is required")
	ErrInputValidation                = errors.New("error input validation")
	ErrUnableToDelete                 = errors.New("unable to delete data")
	ErrUnableToFetch                  = errors.New("unable to fetch data")
	ErrIDNotFound                     = errors.New("id not found ")
	ErrInvalidRequest                 = errors.New("invalid request")
	ErrUnauthorizedClient             = errors.New("unauthorized client")
	ErrAccessDenied                   = errors.New("access denied")
	ErrUnsupportedResponseType        = errors.New("unsupported response type")
	ErrInvalidScope                   = errors.New("invalid scope")
	ErrServerError                    = errors.New("server error")
	ErrTemporarilyUnavailable         = errors.New("temporarily unavailable")
	ErrInvalidClient                  = errors.New("invalid client")
	ErrInvalidGrant                   = errors.New("invalid grant")
	ErrUnsupportedGrantType           = errors.New("unsupported grant type")
	ErrCodeChallengeRquired           = errors.New("invalid request")
	ErrUnsupportedCodeChallengeMethod = errors.New("invalid request")
	ErrInvalidCodeChallengeLen        = errors.New("invalid request")
	ErrorUnableToFetch                = errors.New("unable to fetch")
	ErrorUnableToCreate               = errors.New("unable to create")
	ErrorUnableToConvert              = errors.New("unable to convert")
	ErrorUnableToBindJsonToStruct     = errors.New("unable to bind json to struct")
	ErrInvalidRedirectURI             = errors.New("invalid redirect uri")
	ErrInvalidAuthorizeCode           = errors.New("invalid authorize code")
	ErrInvalidAccessToken             = errors.New("invalid access token")
	ErrInvalidRefreshToken            = errors.New("invalid refresh token")
	ErrExpiredAccessToken             = errors.New("expired access token")
	ErrExpiredRefreshToken            = errors.New("expired refresh token")
	ErrMissingCodeVerifier            = errors.New("missing code verifier")
	ErrMissingCodeChallenge           = errors.New("missing code challenge")
	ErrInvalidCodeChallenge           = errors.New("invalid code challenge")
	ErrDataAlreayExist                = errors.New("data should be unique")
	ErrInvalidAPIKey                  = errors.New("client API Key is invalid")
	ErrInvalidMessage                 = errors.New("message is invalid")
	ErrToManyRegIDs                   = errors.New("too many registrations ids")
	ErrInvalidTimeToLive              = errors.New("messages time-to-live is invalid")
	ErrMissingRegistrationTo          = errors.New("missing registration token")
	ErrInvalidToken                   = errors.New("invalid notification token")
	ErrNotRegistered                  = errors.New("unregistered device")
	ErrInvalidPackageName             = errors.New("invalid package name")
	ErrMismatchSenderID               = errors.New("mismatched sender id")
	ErrMessageTooBig                  = errors.New("message is too big")
	ErrInvalidDataKey                 = errors.New("invalid data key")
	ErrInvalidTTL                     = errors.New("invalid time to live")
	ErrDeviceMessageRateExceeded      = errors.New("device message rate exceeded")
	ErrTopicsMessageRateExceeded      = errors.New("topics message rate exceeded")
	ErrDataIsEmpty                    = errors.New("data and notification are empty")
	ErrInvalidVariable                = errors.New("invalid variable")
	ErrorInvalidRecieverAddress       = errors.New("invalid receiver address")
	ErrorInvalidBody                  = errors.New("invalid body provided")
	ErrorInvalidSenderAddress         = errors.New("invalid sender address")
	ErrorInvalidCallBackUrl           = errors.New("invalid call back url")
	ErrUnableToSendEmailMessage       = errors.New("unable to send email message")
	ErrUnableToSendSmsMessage         = errors.New("unable to send sms message")
	ErrRoleNameISEmpty                = errors.New("role name cannot be empty")
	ErrGenerateToken                  = errors.New("unable to generate token")
	ErrPermissionAlreadyDefined       = errors.New("policy already defined")
	ErrPermissionPermissionNotFound   = errors.New("policy not found")
	ErrDatabaseConnection             = errors.New("database connection failed")
	ErrInvalidTransaction             = errors.New("invalid transaction")
	ErrUnsupportedRelation            = errors.New("unsupported relations")
	ErrInvalidUserPhoneOrPassword     = errors.New("Incorrect phone or password")
	ErrUnsupportedDriver              = errors.New("Unsupported driver")
	ErrInvalidDB                      = errors.New("invalid db")
	ErrInvalidField                   = errors.New("Empty field exist")
	ErrConflictHappened               = errors.New("conflict happened")
	ErrRequireApproval                = errors.New("Require approval")
	ErrUnSupportedSize                = errors.New("Unsupported size")
	ErrFieldAlreadyExist              = errors.New("Field already exist")
)

// Descriptions error description
var Descriptions = map[error]string{
	ErrUnknown:                        "Unknown  error",
	ErrForgotEmail:                    "Email is forgotten",
	ErrInputValidation:                "Error input validation",
	ErrRecordNotFound:                 "Record not found,your transaction fails",
	ErrUnableToSave:                   "Unable to save",
	ErrUnableToDelete:                 "Unable to delete",
	ErrorUnableToFetch:                "Unable to fetch something went wrong ",
	ErrInvalidRequest:                 "The request is missing a required parameter, includes an invalid parameter value, includes a parameter more than once, or is otherwise malformed",
	ErrUnauthorizedClient:             "The client is not authorized to request an authorization code using this method",
	ErrAccessDenied:                   "The resource owner or authorization server denied the request",
	ErrUnsupportedResponseType:        "The authorization server does not support obtaining an authorization code using this method",
	ErrInvalidScope:                   "The requested scope is invalid, unknown, or malformed",
	ErrServerError:                    "The authorization server encountered an unexpected condition that prevented it from fulfilling the request",
	ErrTemporarilyUnavailable:         "The authorization server is currently unable to handle the request due to a temporary overloading or maintenance of the server",
	ErrInvalidClient:                  "Client authentication failed",
	ErrInvalidGrant:                   "The provided authorization grant (e.g., authorization code, resource owner credentials) or refresh token is invalid, expired, revoked, does not match the redirection URI used in the authorization request, or was issued to another client",
	ErrUnsupportedGrantType:           "The authorization grant type is not supported by the authorization server",
	ErrCodeChallengeRquired:           "PKCE is required. code_challenge is missing",
	ErrUnsupportedCodeChallengeMethod: "Selected code_challenge_method not supported",
	ErrInvalidCodeChallengeLen:        "Code challenge length must be between 43 and 128 charachters long",
	ErrInvalidAPIKey:                  "Provided api key is not valid",
	ErrInvalidMessage:                 "Provided message is Invalid",
	ErrToManyRegIDs:                   "Too many registration id exist",
	ErrorUnableToCreate:               "Unable to create",
	ErrInvalidToken:                   "Invalid notification token was provided",
	ErrDataAlreayExist:                "Provided data is already exist",
	ErrorUnableToBindJsonToStruct:     "Unable to parse json to struct",
	ErrorUnableToConvert:              "Unable to convert type conversion",
	ErrInvalidVariable:                "Unable to process the provided variable",
	ErrorInvalidRecieverAddress:       "Unable to find receiver address",
	ErrorInvalidSenderAddress:         "Unable to send text message due to invalid sender address provided ",
	ErrorInvalidBody:                  "The body of message must not be empty ",
	ErrorInvalidCallBackUrl:           "Sms not sent due to invalid call back url is provided ",
	ErrUnableToSendEmailMessage:       "Unable to send email message please try again",
	ErrUnableToSendSmsMessage:         "Unable to send sms message please try again",
	ErrDatabaseConnection:             "Error occurred while establishing a database connection",
	ErrRoleNameISEmpty:                "Role name cannot be empty",
	ErrGenerateToken:                  "Unable to generate token",
	ErrPermissionAlreadyDefined:       "Policy already defined",
	ErrPermissionPermissionNotFound:   "Valid for user policy not found",
	ErrUnableToMigrate:                "Unable to migrate the model",
	ErrInvalidUserPhoneOrPassword:     "Incorect user phone or password",
	ErrUnableToGetUrl:                 "Invalid or unregistered url is provided",
	ErrInvalidField:                   "Field must be field a correct value",
	ErrConflictHappened:               "Conflicting error happened",
	ErrRequireApproval:                "Needs approvals of user by company",
	ErrUnSupportedSize:                "Unsupported image size",
	ErrFieldAlreadyExist:              "No new value found it is already existed",
}

// StatusCodes response error HTTP status code
var StatusCodes = map[error]int{
	ErrInvalidRequest:                 400,
	ErrUnauthorizedClient:             401,
	ErrAccessDenied:                   403,
	ErrUnsupportedResponseType:        401,
	ErrInvalidScope:                   400,
	ErrServerError:                    500,
	ErrTemporarilyUnavailable:         503,
	ErrInvalidClient:                  401,
	ErrInvalidGrant:                   401,
	ErrUnsupportedGrantType:           401,
	ErrCodeChallengeRquired:           400,
	ErrUnsupportedCodeChallengeMethod: 400,
	ErrInvalidCodeChallengeLen:        400,
	ErrIDNotFound:                     404,
	ErrForgotEmail:                    422,
	ErrUnableToSave:                   422,
	ErrUnableToDelete:                 422,
	ErrUnableToFetch:                  422,
	ErrInvalidAPIKey:                  400,
	ErrInvalidMessage:                 400,
	ErrToManyRegIDs:                   406,
	ErrorUnableToCreate:               422,
	ErrInvalidToken:                   400,
	ErrDataAlreayExist:                406,
	ErrorUnableToBindJsonToStruct:     400,
	ErrorUnableToConvert:              403,
	ErrInvalidVariable:                406,
	ErrorInvalidRecieverAddress:       400,
	ErrorInvalidSenderAddress:         400,
	ErrorInvalidBody:                  400,
	ErrorInvalidCallBackUrl:           400,
	ErrUnableToSendEmailMessage:       400,
	ErrUnableToSendSmsMessage:         400,
	ErrDatabaseConnection:             500,
	ErrRoleNameISEmpty:                400,
	ErrGenerateToken:                  500,
	ErrPermissionAlreadyDefined:       400,
	ErrPermissionPermissionNotFound:   400,
	ErrUnableToMigrate:                400,
	ErrInvalidUserPhoneOrPassword:     403,
	ErrUnableToGetUrl:                 400,
	ErrInvalidField:                   400,
	ErrConflictHappened:               409,
	ErrRequireApproval:                400,
	ErrorUnableToFetch:                400,
	ErrRecordNotFound:                 400,
	ErrUnSupportedSize:                400,
	ErrFieldAlreadyExist:              403,
}

// StatusCodes response error HTTP status code
var ErrCodes = map[error]int{
	ErrUnknown:                        5001,
	ErrPasswordEncryption:             5000,
	ErrInvalidAccessToken:             4017,
	ErrInvalidRequest:                 4000,
	ErrUnauthorizedClient:             4001,
	ErrAccessDenied:                   4002,
	ErrUnsupportedResponseType:        4003,
	ErrInvalidScope:                   4004,
	ErrServerError:                    4005,
	ErrTemporarilyUnavailable:         4006,
	ErrInvalidClient:                  4007,
	ErrInvalidGrant:                   4008,
	ErrUnsupportedGrantType:           4009,
	ErrCodeChallengeRquired:           4010,
	ErrUnsupportedCodeChallengeMethod: 4011,
	ErrInvalidCodeChallengeLen:        4012,
	ErrIDNotFound:                     4013,
	ErrUnableToSave:                   4014,
	ErrUnableToDelete:                 4015,
	ErrUnableToFetch:                  4016,
	ErrInvalidAPIKey:                  4017,
	ErrInvalidMessage:                 4018,
	ErrToManyRegIDs:                   4019,
	ErrorUnableToCreate:               4020,
	ErrInvalidToken:                   4021,
	ErrDataAlreayExist:                4022,
	ErrorUnableToBindJsonToStruct:     4023,
	ErrorUnableToConvert:              4024,
	ErrInvalidVariable:                4025,
	ErrorInvalidRecieverAddress:       4026,
	ErrorInvalidSenderAddress:         4027,
	ErrorInvalidBody:                  4028,
	ErrorInvalidCallBackUrl:           4029,
	ErrUnableToSendEmailMessage:       4030,
	ErrUnableToSendSmsMessage:         4031,
	ErrDatabaseConnection:             4032,
	ErrRoleNameISEmpty:                4033,
	ErrGenerateToken:                  4034,
	ErrPermissionAlreadyDefined:       4035,
	ErrPermissionPermissionNotFound:   4036,
	ErrUnableToMigrate:                4037,
	ErrInvalidUserPhoneOrPassword:     4038,
	ErrUnableToGetUrl:                 4039,
	ErrInvalidField:                   4040,
	ErrConflictHappened:               4041,
	ErrRequireApproval:                4042,
	ErrorUnableToFetch:                4043,
	ErrRecordNotFound:                 4044,
	ErrUnSupportedSize:                4045,
	ErrFieldAlreadyExist:              4046,
}
