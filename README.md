CLI tool to print AWS resources and tags returned by the Resource Groups Tagging API.

```sh
$ go install
$ aws-resources -h
```

Note: AWS `GetResources` can return tagged or previously tagged resources, so this tool can list resources that already do not exist. When a result looks suspicious, confirm it with the service-specific API such as `aws ec2 describe-nat-gateways`.

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
