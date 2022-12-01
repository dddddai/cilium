// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

func (c *Client) DescribeVerifiedAccessEndpoints(ctx context.Context, params *DescribeVerifiedAccessEndpointsInput, optFns ...func(*Options)) (*DescribeVerifiedAccessEndpointsOutput, error) {
	if params == nil {
		params = &DescribeVerifiedAccessEndpointsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeVerifiedAccessEndpoints", params, optFns, c.addOperationDescribeVerifiedAccessEndpointsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeVerifiedAccessEndpointsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeVerifiedAccessEndpointsInput struct {
	DryRun *bool

	Filters []types.Filter

	MaxResults *int32

	NextToken *string

	VerifiedAccessEndpointIds []string

	VerifiedAccessGroupId *string

	VerifiedAccessInstanceId *string

	noSmithyDocumentSerde
}

type DescribeVerifiedAccessEndpointsOutput struct {
	NextToken *string

	VerifiedAccessEndpoints []types.VerifiedAccessEndpoint

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeVerifiedAccessEndpointsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeVerifiedAccessEndpoints{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeVerifiedAccessEndpoints{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeVerifiedAccessEndpoints(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

// DescribeVerifiedAccessEndpointsAPIClient is a client that implements the
// DescribeVerifiedAccessEndpoints operation.
type DescribeVerifiedAccessEndpointsAPIClient interface {
	DescribeVerifiedAccessEndpoints(context.Context, *DescribeVerifiedAccessEndpointsInput, ...func(*Options)) (*DescribeVerifiedAccessEndpointsOutput, error)
}

var _ DescribeVerifiedAccessEndpointsAPIClient = (*Client)(nil)

// DescribeVerifiedAccessEndpointsPaginatorOptions is the paginator options for
// DescribeVerifiedAccessEndpoints
type DescribeVerifiedAccessEndpointsPaginatorOptions struct {
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeVerifiedAccessEndpointsPaginator is a paginator for
// DescribeVerifiedAccessEndpoints
type DescribeVerifiedAccessEndpointsPaginator struct {
	options   DescribeVerifiedAccessEndpointsPaginatorOptions
	client    DescribeVerifiedAccessEndpointsAPIClient
	params    *DescribeVerifiedAccessEndpointsInput
	nextToken *string
	firstPage bool
}

// NewDescribeVerifiedAccessEndpointsPaginator returns a new
// DescribeVerifiedAccessEndpointsPaginator
func NewDescribeVerifiedAccessEndpointsPaginator(client DescribeVerifiedAccessEndpointsAPIClient, params *DescribeVerifiedAccessEndpointsInput, optFns ...func(*DescribeVerifiedAccessEndpointsPaginatorOptions)) *DescribeVerifiedAccessEndpointsPaginator {
	if params == nil {
		params = &DescribeVerifiedAccessEndpointsInput{}
	}

	options := DescribeVerifiedAccessEndpointsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeVerifiedAccessEndpointsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeVerifiedAccessEndpointsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeVerifiedAccessEndpoints page.
func (p *DescribeVerifiedAccessEndpointsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeVerifiedAccessEndpointsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.DescribeVerifiedAccessEndpoints(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opDescribeVerifiedAccessEndpoints(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeVerifiedAccessEndpoints",
	}
}