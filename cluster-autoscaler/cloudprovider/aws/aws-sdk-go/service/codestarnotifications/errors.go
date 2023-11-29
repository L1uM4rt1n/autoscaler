// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codestarnotifications

import (
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeAccessDeniedException for service response error code
	// "AccessDeniedException".
	//
	// AWS CodeStar Notifications can't create the notification rule because you
	// do not have sufficient permissions.
	ErrCodeAccessDeniedException = "AccessDeniedException"

	// ErrCodeConcurrentModificationException for service response error code
	// "ConcurrentModificationException".
	//
	// AWS CodeStar Notifications can't complete the request because the resource
	// is being modified by another process. Wait a few minutes and try again.
	ErrCodeConcurrentModificationException = "ConcurrentModificationException"

	// ErrCodeConfigurationException for service response error code
	// "ConfigurationException".
	//
	// Some or all of the configuration is incomplete, missing, or not valid.
	ErrCodeConfigurationException = "ConfigurationException"

	// ErrCodeInvalidNextTokenException for service response error code
	// "InvalidNextTokenException".
	//
	// The value for the enumeration token used in the request to return the next
	// batch of the results is not valid.
	ErrCodeInvalidNextTokenException = "InvalidNextTokenException"

	// ErrCodeLimitExceededException for service response error code
	// "LimitExceededException".
	//
	// One of the AWS CodeStar Notifications limits has been exceeded. Limits apply
	// to accounts, notification rules, notifications, resources, and targets. For
	// more information, see Limits.
	ErrCodeLimitExceededException = "LimitExceededException"

	// ErrCodeResourceAlreadyExistsException for service response error code
	// "ResourceAlreadyExistsException".
	//
	// A resource with the same name or ID already exists. Notification rule names
	// must be unique in your Amazon Web Services account.
	ErrCodeResourceAlreadyExistsException = "ResourceAlreadyExistsException"

	// ErrCodeResourceNotFoundException for service response error code
	// "ResourceNotFoundException".
	//
	// AWS CodeStar Notifications can't find a resource that matches the provided
	// ARN.
	ErrCodeResourceNotFoundException = "ResourceNotFoundException"

	// ErrCodeValidationException for service response error code
	// "ValidationException".
	//
	// One or more parameter values are not valid.
	ErrCodeValidationException = "ValidationException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"AccessDeniedException":           newErrorAccessDeniedException,
	"ConcurrentModificationException": newErrorConcurrentModificationException,
	"ConfigurationException":          newErrorConfigurationException,
	"InvalidNextTokenException":       newErrorInvalidNextTokenException,
	"LimitExceededException":          newErrorLimitExceededException,
	"ResourceAlreadyExistsException":  newErrorResourceAlreadyExistsException,
	"ResourceNotFoundException":       newErrorResourceNotFoundException,
	"ValidationException":             newErrorValidationException,
}
