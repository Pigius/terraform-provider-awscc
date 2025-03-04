---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_cleanrooms_membership Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::CleanRooms::Membership
---

# awscc_cleanrooms_membership (Data Source)

Data Source schema for AWS::CleanRooms::Membership



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `arn` (String)
- `collaboration_arn` (String)
- `collaboration_creator_account_id` (String)
- `collaboration_identifier` (String)
- `default_result_configuration` (Attributes) (see [below for nested schema](#nestedatt--default_result_configuration))
- `membership_identifier` (String)
- `query_log_status` (String)
- `tags` (Attributes Set) An arbitrary set of tags (key-value pairs) for this cleanrooms membership. (see [below for nested schema](#nestedatt--tags))

<a id="nestedatt--default_result_configuration"></a>
### Nested Schema for `default_result_configuration`

Read-Only:

- `output_configuration` (Attributes) (see [below for nested schema](#nestedatt--default_result_configuration--output_configuration))
- `role_arn` (String)

<a id="nestedatt--default_result_configuration--output_configuration"></a>
### Nested Schema for `default_result_configuration.output_configuration`

Read-Only:

- `s3` (Attributes) (see [below for nested schema](#nestedatt--default_result_configuration--output_configuration--s3))

<a id="nestedatt--default_result_configuration--output_configuration--s3"></a>
### Nested Schema for `default_result_configuration.output_configuration.s3`

Read-Only:

- `bucket` (String)
- `key_prefix` (String)
- `result_format` (String)




<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Read-Only:

- `key` (String)
- `value` (String)
