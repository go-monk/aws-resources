CLI tool to print AWS resources and tags.

```sh
$ go install
$ aws-resources -h
```

Note: This tool uses the Resource Groups Tagging API that can sometimes return resources that already don't exist.

---

Similar `aws` commands:

```sh
# aws-resources
aws resourcegroupstaggingapi get-resources
# aws-resources -tag environment=stage 
aws resourcegroupstaggingapi get-resources --tag-filters Key=environment,Values=stage
# aws-resources -tag environment=stage -tag environment=prod
aws resourcegroupstaggingapi get-resources --tag-filters Key=environment,Values=stage,prod
```

Advantages over `aws`:

- easier to remember and use
- no dependencies, just a single binary
- portable - works on Mac, Linux, Windows, ...

---

Some use cases:

```sh
# Count resources in all available profiles.
for profile in $(aws configure list-profiles); do echo -n "$profile: "; aws-resources -profile $profile | wc -l; done

# Get resources with tags in all existing regions.
for region in $(aws ec2 describe-regions --all-regions --query 'Regions[].RegionName' --output text); do echo "---$region---"; aws-resources -region $region -tags; done
```
