// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package verifiedpermissions

import (
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeAccessDeniedException for service response error code
	// "AccessDeniedException".
	//
	// You don't have sufficient access to perform this action.
	ErrCodeAccessDeniedException = "AccessDeniedException"

	// ErrCodeConflictException for service response error code
	// "ConflictException".
	//
	// The request failed because another request to modify a resource occurred
	// at the same.
	ErrCodeConflictException = "ConflictException"

	// ErrCodeInternalServerException for service response error code
	// "InternalServerException".
	//
	// The request failed because of an internal error. Try your request again later
	ErrCodeInternalServerException = "InternalServerException"

	// ErrCodeResourceNotFoundException for service response error code
	// "ResourceNotFoundException".
	//
	// The request failed because it references a resource that doesn't exist.
	ErrCodeResourceNotFoundException = "ResourceNotFoundException"

	// ErrCodeServiceQuotaExceededException for service response error code
	// "ServiceQuotaExceededException".
	//
	// The request failed because it would cause a service quota to be exceeded.
	ErrCodeServiceQuotaExceededException = "ServiceQuotaExceededException"

	// ErrCodeThrottlingException for service response error code
	// "ThrottlingException".
	//
	// The request failed because it exceeded a throttling quota.
	ErrCodeThrottlingException = "ThrottlingException"

	// ErrCodeValidationException for service response error code
	// "ValidationException".
	//
	// The request failed because one or more input parameters don't satisfy their
	// constraint requirements. The output is provided as a list of fields and a
	// reason for each field that isn't valid.
	//
	// The possible reasons include the following:
	//
	//    * UnrecognizedEntityType The policy includes an entity type that isn't
	//    found in the schema.
	//
	//    * UnrecognizedActionId The policy includes an action id that isn't found
	//    in the schema.
	//
	//    * InvalidActionApplication The policy includes an action that, according
	//    to the schema, doesn't support the specified principal and resource.
	//
	//    * UnexpectedType The policy included an operand that isn't a valid type
	//    for the specified operation.
	//
	//    * IncompatibleTypes The types of elements included in a set, or the types
	//    of expressions used in an if...then...else clause aren't compatible in
	//    this context.
	//
	//    * MissingAttribute The policy attempts to access a record or entity attribute
	//    that isn't specified in the schema. Test for the existence of the attribute
	//    first before attempting to access its value. For more information, see
	//    the has (presence of attribute test) operator (https://docs.cedarpolicy.com/policies/syntax-operators.html#has-presence-of-attribute-test)
	//    in the Cedar Policy Language Guide.
	//
	//    * UnsafeOptionalAttributeAccess The policy attempts to access a record
	//    or entity attribute that is optional and isn't guaranteed to be present.
	//    Test for the existence of the attribute first before attempting to access
	//    its value. For more information, see the has (presence of attribute test)
	//    operator (https://docs.cedarpolicy.com/policies/syntax-operators.html#has-presence-of-attribute-test)
	//    in the Cedar Policy Language Guide.
	//
	//    * ImpossiblePolicy Cedar has determined that a policy condition always
	//    evaluates to false. If the policy is always false, it can never apply
	//    to any query, and so it can never affect an authorization decision.
	//
	//    * WrongNumberArguments The policy references an extension type with the
	//    wrong number of arguments.
	//
	//    * FunctionArgumentValidationError Cedar couldn't parse the argument passed
	//    to an extension type. For example, a string that is to be parsed as an
	//    IPv4 address can contain only digits and the period character.
	ErrCodeValidationException = "ValidationException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"AccessDeniedException":         newErrorAccessDeniedException,
	"ConflictException":             newErrorConflictException,
	"InternalServerException":       newErrorInternalServerException,
	"ResourceNotFoundException":     newErrorResourceNotFoundException,
	"ServiceQuotaExceededException": newErrorServiceQuotaExceededException,
	"ThrottlingException":           newErrorThrottlingException,
	"ValidationException":           newErrorValidationException,
}
