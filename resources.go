package main

import (
	"context"
	"fmt"
	"slices"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/resourcegroupstaggingapi"
	"github.com/aws/aws-sdk-go-v2/service/resourcegroupstaggingapi/types"
)

func GetResources(ctx context.Context, cfg aws.Config, tagfilters Tags) (Resources, error) {
	client := resourcegroupstaggingapi.NewFromConfig(cfg)

	in := &resourcegroupstaggingapi.GetResourcesInput{
		TagFilters:       []types.TagFilter{},
		ResourcesPerPage: aws.Int32(100), // Maximum page size for faster pagination
	}

	if len(tagfilters) > 0 {
		for _, tag := range tagfilters {
			in.TagFilters = append(in.TagFilters, types.TagFilter{
				Key:    &tag.Key,
				Values: []string{tag.Value},
			})
		}
	}

	var resources Resources

	for {
		out, err := client.GetResources(ctx, in)
		if err != nil {
			return Resources{}, err
		}

		resources = append(resources, out.ResourceTagMappingList...)

		if aws.ToString(out.PaginationToken) == "" {
			break
		}
		in.PaginationToken = out.PaginationToken
	}

	return resources, nil
}

type Resources []types.ResourceTagMapping

func (resources Resources) Print() {
	arns := make([]string, 0, len(resources))
	for _, resource := range resources {
		arns = append(arns, aws.ToString(resource.ResourceARN))
	}

	// Sort ARNs for consistent output
	slices.Sort(arns)

	for _, arn := range arns {
		fmt.Println(arn)
	}
}

func (resources Resources) PrintWithTags() {
	resourceMap := make(map[string]types.ResourceTagMapping)
	arns := make([]string, 0, len(resources))

	for _, resource := range resources {
		arn := aws.ToString(resource.ResourceARN)
		arns = append(arns, arn)
		resourceMap[arn] = resource
	}

	// Sort ARNs for consistent output
	slices.Sort(arns)

	// Print each resource with its tags
	for _, arn := range arns {
		resource := resourceMap[arn]
		fmt.Println(arn)

		if len(resource.Tags) > 0 {
			// Sort tags by key for consistent output
			tagKeys := make([]string, 0, len(resource.Tags))
			tagMap := make(map[string]string)

			for _, tag := range resource.Tags {
				key := aws.ToString(tag.Key)
				tagKeys = append(tagKeys, key)
				tagMap[key] = aws.ToString(tag.Value)
			}

			slices.Sort(tagKeys)

			for _, key := range tagKeys {
				fmt.Printf("- %s=%s\n", key, tagMap[key])
			}
		}
	}
}
