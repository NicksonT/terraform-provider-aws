// Code generated by internal/generate/tags/main.go; DO NOT EDIT.
package codebuild

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/codebuild"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
)

// listTags_Func is the type of the listTags_ function.
type listTags_Func func(context.Context, string) error

// updateTags_Func is the type of the updateTags_ function.
type updateTags_Func func(context.Context, string, any, any) error

var listTags_ listTags_Func

// []*SERVICE.Tag handling

// Tags returns codebuild service tags.
func Tags(tags tftags.KeyValueTags) []*codebuild.Tag {
	result := make([]*codebuild.Tag, 0, len(tags))

	for k, v := range tags.Map() {
		tag := &codebuild.Tag{
			Key:   aws.String(k),
			Value: aws.String(v),
		}

		result = append(result, tag)
	}

	return result
}

// KeyValueTags creates tftags.KeyValueTags from codebuild service tags.
func KeyValueTags(ctx context.Context, tags []*codebuild.Tag) tftags.KeyValueTags {
	m := make(map[string]*string, len(tags))

	for _, tag := range tags {
		m[aws.StringValue(tag.Key)] = tag.Value
	}

	return tftags.New(ctx, m)
}

// getTagsIn returns codebuild service tags from Context.
// nil is returned if there are no input tags.
func getTagsIn(ctx context.Context) []*codebuild.Tag {
	if inContext, ok := tftags.FromContext(ctx); ok {
		if tags := Tags(inContext.TagsIn.UnwrapOrDefault()); len(tags) > 0 {
			return tags
		}
	}

	return nil
}

// setTagsOut sets codebuild service tags in Context.
func setTagsOut(ctx context.Context, tags []*codebuild.Tag) {
	if inContext, ok := tftags.FromContext(ctx); ok {
		inContext.TagsOut = types.Some(KeyValueTags(ctx, tags))
	}
}

var updateTags_ updateTags_Func
