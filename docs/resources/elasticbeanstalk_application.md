---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_elasticbeanstalk_application Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  The AWS::ElasticBeanstalk::Application resource specifies an Elastic Beanstalk application.
---

# awscc_elasticbeanstalk_application (Resource)

The AWS::ElasticBeanstalk::Application resource specifies an Elastic Beanstalk application.



<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `application_name` (String) A name for the Elastic Beanstalk application. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the application name.
- `description` (String) Your description of the application.
- `resource_lifecycle_config` (Attributes) Specifies an application resource lifecycle configuration to prevent your application from accumulating too many versions. (see [below for nested schema](#nestedatt--resource_lifecycle_config))

### Read-Only

- `id` (String) Uniquely identifies the resource.

<a id="nestedatt--resource_lifecycle_config"></a>
### Nested Schema for `resource_lifecycle_config`

Optional:

- `service_role` (String) The ARN of an IAM service role that Elastic Beanstalk has permission to assume. The ServiceRole property is required the first time that you provide a ResourceLifecycleConfig for the application. After you provide it once, Elastic Beanstalk persists the Service Role with the application, and you don't need to specify it again. You can, however, specify it in subsequent updates to change the Service Role to another value.
- `version_lifecycle_config` (Attributes) Defines lifecycle settings for application versions. (see [below for nested schema](#nestedatt--resource_lifecycle_config--version_lifecycle_config))

<a id="nestedatt--resource_lifecycle_config--version_lifecycle_config"></a>
### Nested Schema for `resource_lifecycle_config.version_lifecycle_config`

Optional:

- `max_age_rule` (Attributes) Specify a max age rule to restrict the length of time that application versions are retained for an application. (see [below for nested schema](#nestedatt--resource_lifecycle_config--version_lifecycle_config--max_age_rule))
- `max_count_rule` (Attributes) Specify a max count rule to restrict the number of application versions that are retained for an application. (see [below for nested schema](#nestedatt--resource_lifecycle_config--version_lifecycle_config--max_count_rule))

<a id="nestedatt--resource_lifecycle_config--version_lifecycle_config--max_age_rule"></a>
### Nested Schema for `resource_lifecycle_config.version_lifecycle_config.max_age_rule`

Optional:

- `delete_source_from_s3` (Boolean) Set to true to delete a version's source bundle from Amazon S3 when Elastic Beanstalk deletes the application version.
- `enabled` (Boolean) Specify true to apply the rule, or false to disable it.
- `max_age_in_days` (Number) Specify the number of days to retain an application versions.


<a id="nestedatt--resource_lifecycle_config--version_lifecycle_config--max_count_rule"></a>
### Nested Schema for `resource_lifecycle_config.version_lifecycle_config.max_count_rule`

Optional:

- `delete_source_from_s3` (Boolean) Set to true to delete a version's source bundle from Amazon S3 when Elastic Beanstalk deletes the application version.
- `enabled` (Boolean) Specify true to apply the rule, or false to disable it.
- `max_count` (Number) Specify the maximum number of application versions to retain.

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_elasticbeanstalk_application.example <resource ID>
```
