// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package appflow

import (
	"context"

	aws_sdkv1 "github.com/aws/aws-sdk-go/aws"
	session_sdkv1 "github.com/aws/aws-sdk-go/aws/session"
	appflow_sdkv1 "github.com/aws/aws-sdk-go/service/appflow"
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
	return []*types.ServicePackageSDKDataSource{}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceConnectorProfile,
			TypeName: "aws_appflow_connector_profile",
		},
		{
			Factory:  ResourceFlow,
			TypeName: "aws_appflow_flow",
			Name:     "Flow",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "id",
				ListTags:            listTags_,
				UpdateTags:          updateTags_,
			},
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.AppFlow
}

// NewConn returns a new AWS SDK for Go v1 client for this service package's AWS API.
func (p *servicePackage) NewConn(ctx context.Context, config map[string]any) (*appflow_sdkv1.Appflow, error) {
	sess := config["session"].(*session_sdkv1.Session)

	return appflow_sdkv1.New(sess.Copy(&aws_sdkv1.Config{Endpoint: aws_sdkv1.String(config["endpoint"].(string))})), nil
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
