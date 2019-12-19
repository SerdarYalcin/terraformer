// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
)

type DeleteUserPolicyInput struct {
	_ struct{} `type:"structure"`

	// The name identifying the policy document to delete.
	//
	// This parameter allows (through its regex pattern (http://wikipedia.org/wiki/regex))
	// a string of characters consisting of upper and lowercase alphanumeric characters
	// with no spaces. You can also include any of the following characters: _+=,.@-
	//
	// PolicyName is a required field
	PolicyName *string `min:"1" type:"string" required:"true"`

	// The name (friendly name, not ARN) identifying the user that the policy is
	// embedded in.
	//
	// This parameter allows (through its regex pattern (http://wikipedia.org/wiki/regex))
	// a string of characters consisting of upper and lowercase alphanumeric characters
	// with no spaces. You can also include any of the following characters: _+=,.@-
	//
	// UserName is a required field
	UserName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteUserPolicyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteUserPolicyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteUserPolicyInput"}

	if s.PolicyName == nil {
		invalidParams.Add(aws.NewErrParamRequired("PolicyName"))
	}
	if s.PolicyName != nil && len(*s.PolicyName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("PolicyName", 1))
	}

	if s.UserName == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserName"))
	}
	if s.UserName != nil && len(*s.UserName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("UserName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteUserPolicyOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteUserPolicyOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteUserPolicy = "DeleteUserPolicy"

// DeleteUserPolicyRequest returns a request value for making API operation for
// AWS Identity and Access Management.
//
// Deletes the specified inline policy that is embedded in the specified IAM
// user.
//
// A user can also have managed policies attached to it. To detach a managed
// policy from a user, use DetachUserPolicy. For more information about policies,
// refer to Managed Policies and Inline Policies (https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-vs-inline.html)
// in the IAM User Guide.
//
//    // Example sending a request using DeleteUserPolicyRequest.
//    req := client.DeleteUserPolicyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/DeleteUserPolicy
func (c *Client) DeleteUserPolicyRequest(input *DeleteUserPolicyInput) DeleteUserPolicyRequest {
	op := &aws.Operation{
		Name:       opDeleteUserPolicy,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteUserPolicyInput{}
	}

	req := c.newRequest(op, input, &DeleteUserPolicyOutput{})
	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteUserPolicyRequest{Request: req, Input: input, Copy: c.DeleteUserPolicyRequest}
}

// DeleteUserPolicyRequest is the request type for the
// DeleteUserPolicy API operation.
type DeleteUserPolicyRequest struct {
	*aws.Request
	Input *DeleteUserPolicyInput
	Copy  func(*DeleteUserPolicyInput) DeleteUserPolicyRequest
}

// Send marshals and sends the DeleteUserPolicy API request.
func (r DeleteUserPolicyRequest) Send(ctx context.Context) (*DeleteUserPolicyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteUserPolicyResponse{
		DeleteUserPolicyOutput: r.Request.Data.(*DeleteUserPolicyOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteUserPolicyResponse is the response type for the
// DeleteUserPolicy API operation.
type DeleteUserPolicyResponse struct {
	*DeleteUserPolicyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteUserPolicy request.
func (r *DeleteUserPolicyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}