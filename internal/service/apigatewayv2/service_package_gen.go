// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package apigatewayv2

import (
	"context"

	aws_sdkv1 "github.com/aws/aws-sdk-go/aws"
	session_sdkv1 "github.com/aws/aws-sdk-go/aws/session"
	apigatewayv2_sdkv1 "github.com/aws/aws-sdk-go/service/apigatewayv2"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  DataSourceAPI,
			TypeName: "aws_apigatewayv2_api",
		},
		{
			Factory:  DataSourceAPIs,
			TypeName: "aws_apigatewayv2_apis",
		},
		{
			Factory:  DataSourceExport,
			TypeName: "aws_apigatewayv2_export",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceAPI,
			TypeName: "aws_apigatewayv2_api",
			Name:     "API",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
				ListTags:            listTags_,
				UpdateTags:          updateTags_,
			},
		},
		{
			Factory:  ResourceAPIMapping,
			TypeName: "aws_apigatewayv2_api_mapping",
		},
		{
			Factory:  ResourceAuthorizer,
			TypeName: "aws_apigatewayv2_authorizer",
		},
		{
			Factory:  ResourceDeployment,
			TypeName: "aws_apigatewayv2_deployment",
		},
		{
			Factory:  ResourceDomainName,
			TypeName: "aws_apigatewayv2_domain_name",
			Name:     "Domain Name",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
				ListTags:            listTags_,
				UpdateTags:          updateTags_,
			},
		},
		{
			Factory:  ResourceIntegration,
			TypeName: "aws_apigatewayv2_integration",
		},
		{
			Factory:  ResourceIntegrationResponse,
			TypeName: "aws_apigatewayv2_integration_response",
		},
		{
			Factory:  ResourceModel,
			TypeName: "aws_apigatewayv2_model",
		},
		{
			Factory:  ResourceRoute,
			TypeName: "aws_apigatewayv2_route",
		},
		{
			Factory:  ResourceRouteResponse,
			TypeName: "aws_apigatewayv2_route_response",
		},
		{
			Factory:  ResourceStage,
			TypeName: "aws_apigatewayv2_stage",
			Name:     "Stage",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
				ListTags:            listTags_,
				UpdateTags:          updateTags_,
			},
		},
		{
			Factory:  ResourceVPCLink,
			TypeName: "aws_apigatewayv2_vpc_link",
			Name:     "VPC Link",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
				ListTags:            listTags_,
				UpdateTags:          updateTags_,
			},
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.APIGatewayV2
}

// NewConn returns a new AWS SDK for Go v1 client for this service package's AWS API.
func (p *servicePackage) NewConn(ctx context.Context, config map[string]any) (*apigatewayv2_sdkv1.ApiGatewayV2, error) {
	sess := config["session"].(*session_sdkv1.Session)

	return apigatewayv2_sdkv1.New(sess.Copy(&aws_sdkv1.Config{Endpoint: aws_sdkv1.String(config["endpoint"].(string))})), nil
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
