// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package ssmincidents

import (
	"context"

	aws_sdkv2 "github.com/aws/aws-sdk-go-v2/aws"
	ssmincidents_sdkv2 "github.com/aws/aws-sdk-go-v2/service/ssmincidents"
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
			Factory:  DataSourceReplicationSet,
			TypeName: "aws_ssmincidents_replication_set",
		},
		{
			Factory:  DataSourceResponsePlan,
			TypeName: "aws_ssmincidents_response_plan",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceReplicationSet,
			TypeName: "aws_ssmincidents_replication_set",
			Name:     "Replication Set",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "id",
				ListTags:            listTags_,
				UpdateTags:          updateTags_,
			},
		},
		{
			Factory:  ResourceResponsePlan,
			TypeName: "aws_ssmincidents_response_plan",
			Name:     "Response Plan",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "id",
				ListTags:            listTags_,
				UpdateTags:          updateTags_,
			},
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.SSMIncidents
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*ssmincidents_sdkv2.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws_sdkv2.Config))

	return ssmincidents_sdkv2.NewFromConfig(cfg, func(o *ssmincidents_sdkv2.Options) {
		if endpoint := config["endpoint"].(string); endpoint != "" {
			o.EndpointResolver = ssmincidents_sdkv2.EndpointResolverFromURL(endpoint)
		}
	}), nil
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
