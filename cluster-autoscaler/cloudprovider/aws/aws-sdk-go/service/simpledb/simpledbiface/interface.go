// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package simpledbiface provides an interface to enable mocking the Amazon SimpleDB service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package simpledbiface

import (
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws/request"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/service/simpledb"
)

// SimpleDBAPI provides an interface to enable mocking the
// simpledb.SimpleDB service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//	// myFunc uses an SDK service client to make a request to
//	// Amazon SimpleDB.
//	func myFunc(svc simpledbiface.SimpleDBAPI) bool {
//	    // Make svc.BatchDeleteAttributes request
//	}
//
//	func main() {
//	    sess := session.New()
//	    svc := simpledb.New(sess)
//
//	    myFunc(svc)
//	}
//
// In your _test.go file:
//
//	// Define a mock struct to be used in your unit tests of myFunc.
//	type mockSimpleDBClient struct {
//	    simpledbiface.SimpleDBAPI
//	}
//	func (m *mockSimpleDBClient) BatchDeleteAttributes(input *simpledb.BatchDeleteAttributesInput) (*simpledb.BatchDeleteAttributesOutput, error) {
//	    // mock response/functionality
//	}
//
//	func TestMyFunc(t *testing.T) {
//	    // Setup Test
//	    mockSvc := &mockSimpleDBClient{}
//
//	    myfunc(mockSvc)
//
//	    // Verify myFunc's functionality
//	}
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type SimpleDBAPI interface {
	BatchDeleteAttributes(*simpledb.BatchDeleteAttributesInput) (*simpledb.BatchDeleteAttributesOutput, error)
	BatchDeleteAttributesWithContext(aws.Context, *simpledb.BatchDeleteAttributesInput, ...request.Option) (*simpledb.BatchDeleteAttributesOutput, error)
	BatchDeleteAttributesRequest(*simpledb.BatchDeleteAttributesInput) (*request.Request, *simpledb.BatchDeleteAttributesOutput)

	BatchPutAttributes(*simpledb.BatchPutAttributesInput) (*simpledb.BatchPutAttributesOutput, error)
	BatchPutAttributesWithContext(aws.Context, *simpledb.BatchPutAttributesInput, ...request.Option) (*simpledb.BatchPutAttributesOutput, error)
	BatchPutAttributesRequest(*simpledb.BatchPutAttributesInput) (*request.Request, *simpledb.BatchPutAttributesOutput)

	CreateDomain(*simpledb.CreateDomainInput) (*simpledb.CreateDomainOutput, error)
	CreateDomainWithContext(aws.Context, *simpledb.CreateDomainInput, ...request.Option) (*simpledb.CreateDomainOutput, error)
	CreateDomainRequest(*simpledb.CreateDomainInput) (*request.Request, *simpledb.CreateDomainOutput)

	DeleteAttributes(*simpledb.DeleteAttributesInput) (*simpledb.DeleteAttributesOutput, error)
	DeleteAttributesWithContext(aws.Context, *simpledb.DeleteAttributesInput, ...request.Option) (*simpledb.DeleteAttributesOutput, error)
	DeleteAttributesRequest(*simpledb.DeleteAttributesInput) (*request.Request, *simpledb.DeleteAttributesOutput)

	DeleteDomain(*simpledb.DeleteDomainInput) (*simpledb.DeleteDomainOutput, error)
	DeleteDomainWithContext(aws.Context, *simpledb.DeleteDomainInput, ...request.Option) (*simpledb.DeleteDomainOutput, error)
	DeleteDomainRequest(*simpledb.DeleteDomainInput) (*request.Request, *simpledb.DeleteDomainOutput)

	DomainMetadata(*simpledb.DomainMetadataInput) (*simpledb.DomainMetadataOutput, error)
	DomainMetadataWithContext(aws.Context, *simpledb.DomainMetadataInput, ...request.Option) (*simpledb.DomainMetadataOutput, error)
	DomainMetadataRequest(*simpledb.DomainMetadataInput) (*request.Request, *simpledb.DomainMetadataOutput)

	GetAttributes(*simpledb.GetAttributesInput) (*simpledb.GetAttributesOutput, error)
	GetAttributesWithContext(aws.Context, *simpledb.GetAttributesInput, ...request.Option) (*simpledb.GetAttributesOutput, error)
	GetAttributesRequest(*simpledb.GetAttributesInput) (*request.Request, *simpledb.GetAttributesOutput)

	ListDomains(*simpledb.ListDomainsInput) (*simpledb.ListDomainsOutput, error)
	ListDomainsWithContext(aws.Context, *simpledb.ListDomainsInput, ...request.Option) (*simpledb.ListDomainsOutput, error)
	ListDomainsRequest(*simpledb.ListDomainsInput) (*request.Request, *simpledb.ListDomainsOutput)

	ListDomainsPages(*simpledb.ListDomainsInput, func(*simpledb.ListDomainsOutput, bool) bool) error
	ListDomainsPagesWithContext(aws.Context, *simpledb.ListDomainsInput, func(*simpledb.ListDomainsOutput, bool) bool, ...request.Option) error

	PutAttributes(*simpledb.PutAttributesInput) (*simpledb.PutAttributesOutput, error)
	PutAttributesWithContext(aws.Context, *simpledb.PutAttributesInput, ...request.Option) (*simpledb.PutAttributesOutput, error)
	PutAttributesRequest(*simpledb.PutAttributesInput) (*request.Request, *simpledb.PutAttributesOutput)

	Select(*simpledb.SelectInput) (*simpledb.SelectOutput, error)
	SelectWithContext(aws.Context, *simpledb.SelectInput, ...request.Option) (*simpledb.SelectOutput, error)
	SelectRequest(*simpledb.SelectInput) (*request.Request, *simpledb.SelectOutput)

	SelectPages(*simpledb.SelectInput, func(*simpledb.SelectOutput, bool) bool) error
	SelectPagesWithContext(aws.Context, *simpledb.SelectInput, func(*simpledb.SelectOutput, bool) bool, ...request.Option) error
}

var _ SimpleDBAPI = (*simpledb.SimpleDB)(nil)
