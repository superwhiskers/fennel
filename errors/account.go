/*

fennel - nintendo network utility library for golang
Copyright (C) 2018-2019 superwhiskers <whiskerdev@protonmail.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Lesser General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Lesser General Public License for more details.

You should have received a copy of the GNU Lesser General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.

*/

package errors

// AccountServerError represents an error recieved from the nintendo network account server
type AccountServerError string

// account server error constants. errors marked with pretendo-specific are error codes that appear to be specific to the pretendo implementation of nintendo network
// if an error marked this way isn't pretendo-specific, let me know
const (
	UnknownError                               AccountServerError = "0000" // pretendo-specific?
	BadParameterFormatError                    AccountServerError = "0001" // this error indicates that you provided a wrong parameter for the request
	BadRequestFormatError                      AccountServerError = "0002" // this error indicates that the request format was invalid
	MissingRequestParameterError               AccountServerError = "0003" // this error indicates that you are missing a request parameter
	UnauthorizedClientError                    AccountServerError = "0004" // this error indicates that you probably provided an incorrect/are missing something in the ClientInformation struct
	InvalidAccountTokenError                   AccountServerError = "0005" // this error indicates that the account token that you provided is invalid
	ExpiredAccountTokenError                   AccountServerError = "0006" // this error indicates that you provided an expired account token to the server
	ForbiddenRequestError                      AccountServerError = "0007"
	RequestNotFoundError                       AccountServerError = "0008"
	InvalidHTTPMethodError                     AccountServerError = "0009"
	InvalidPlatformIDError                     AccountServerError = "0010"
	SystemUpdateRequiredError                  AccountServerError = "0011"
	BannedDeviceError                          AccountServerError = "0012" // this error indicates that your device is banned from all services forever
	AccountIDExistsError                       AccountServerError = "0100" // this error indicates that the account id you specified to the request exists already
	InvalidAccountIDError                      AccountServerError = "0101"
	InvalidMailAddressError                    AccountServerError = "0103"
	UnauthorizedDeviceError                    AccountServerError = "0104"
	RegistrationLimitReachedError              AccountServerError = "0105"
	InvalidAccountPasswordError                AccountServerError = "0106"
	CountryMismatchError                       AccountServerError = "0107"
	BannedAccountError                         AccountServerError = "0108" // this error indicates that your account is banned from all services forever
	DeviceMismatchError                        AccountServerError = "0110"
	AccountIDChangedError                      AccountServerError = "0111"
	AccountDeletedError                        AccountServerError = "0112"
	COPPANotAcceptedError                      AccountServerError = "0114"
	AssociationLimitReachedError               AccountServerError = "0115"
	InvalidConfirmationCodeError               AccountServerError = "0116"
	ConfirmationCodeExpiredError               AccountServerError = "0117"
	GameServerUniqueIDNotLinkedError           AccountServerError = "0118"
	BannedAccountInApplicationError            AccountServerError = "0119"
	BannedDeviceInApplicationError             AccountServerError = "0120"
	BannedAccountInNEXServiceError             AccountServerError = "0121"
	BannedDeviceInNEXServiceError              AccountServerError = "0122"
	ServiceClosedError                         AccountServerError = "0123"
	ApplicationUpdateRequiredError             AccountServerError = "0124"
	ClientUniqueIDNotLinkedError               AccountServerError = "0125"
	BannedAccountInIndependentServiceError     AccountServerError = "0126"
	BannedDeviceInIndependentServiceError      AccountServerError = "0127"
	MailAddressNotValidatedError               AccountServerError = "0128"
	WrongBirthdateOrMailAddressError           AccountServerError = "0129"
	PIDNotFoundError                           AccountServerError = "0130"
	WrongAccountMailError                      AccountServerError = "0131"
	TempbannedAccountError                     AccountServerError = "0132" // this error indicates that your account is temporarily banned from all services
	TempbannedAccountInApplicationError        AccountServerError = "0134" // this error indicates that your account is temporarily banned from the specified application
	TempbannedAccountInNEXServiceError         AccountServerError = "0136" // this error indicates that your account is temporarily banned from the specified NEX service
	TempbannedDeviceInNEXServiceError          AccountServerError = "0137" // this error indicates that your device is temporarily banned from the specified NEX service
	TempbannedAccountInIndependentServiceError AccountServerError = "0138" // this error indicates that your account is temporarily banned from the specified independent service
	TempbannedDeviceInIndependentServiceError  AccountServerError = "0139" // this error indicates that your device is temporarily banned from the specified independent service
	COPPAAgreementCanceledError                AccountServerError = "0142"
	DeviceInactiveError                        AccountServerError = "0143"
	EULANotAcceptedError                       AccountServerError = "1004"
	InvalidUniqueIDError                       AccountServerError = "1006"
	NEXAccountNotFoundError                    AccountServerError = "1016"
	GameServerIDEnvironmentNotFoundError       AccountServerError = "1017"
	TokenGenerationFailedError                 AccountServerError = "1018"
	InvalidNEXClientIDError                    AccountServerError = "1019"
	InvalidClientKeyError                      AccountServerError = "1020"
	InvalidGameServerIDError                   AccountServerError = "1021"
	InvalidClientIDError                       AccountServerError = "1022"
	WrongMailAddressError                      AccountServerError = "1023"
	MasterPinNotFoundError                     AccountServerError = "1024"
	MailTextNotFoundError                      AccountServerError = "1025"
	MailSendFailedError                        AccountServerError = "1031"
	DomainAccountAlreadyExistsError            AccountServerError = "1032"
	ExcessiveMailSendRequestError              AccountServerError = "1033"
	CreditCardGeneralFailiure                  AccountServerError = "1035"
	CreditCardDateExpiredError                 AccountServerError = "1036"
	CreditCardDeclinedError                    AccountServerError = "1037"
	InvalidCreditCardNumberError               AccountServerError = "1038"
	CreditCardNumberWrongError                 AccountServerError = "1039"
	InvalidCreditCardDateError                 AccountServerError = "1040"
	CreditCardBlacklistedError                 AccountServerError = "1041"
	InvalidCreditCardPinError                  AccountServerError = "1042"
	WrongCreditCardPinError                    AccountServerError = "1043"
	InvalidLocationError                       AccountServerError = "1044"
	InvalidPostalCodeError                     AccountServerError = "1045"
	DeviceEULACountryMismatchError             AccountServerError = "1046"
	InvalidEULACountryError                    AccountServerError = "1100"
	InvalidEULACountryAndVersionError          AccountServerError = "1101"
	ParentalControlsRequiredError              AccountServerError = "1103"
	AccountIDFormatInvalidError                AccountServerError = "1104"
	WrongAccountPasswordOrMailAddressError     AccountServerError = "1105"
	AuthenticationLockedError                  AccountServerError = "1106"
	AccountIDPasswordSameError                 AccountServerError = "1107"
	ApprovalIDNotFoundError                    AccountServerError = "1111"
	PendingMigrationError                      AccountServerError = "1115"
	MailAddressDomainNameNotAcceptableError    AccountServerError = "1125"
	MailAddressDomainNameNotResolvedError      AccountServerError = "1126"
	CountryNotProvidedError                    AccountServerError = "1200"
	BadRequestError                            AccountServerError = "1600" // pretendo-specific?
	InternalServerError                        AccountServerError = "2001"
	UnderMaintenanceError                      AccountServerError = "2002"
	NintendoNetworkClosedError                 AccountServerError = "2999"

	// unknown: error 1134
)
