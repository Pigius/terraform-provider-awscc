---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_guardduty_ip_set Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource Type definition for AWS::GuardDuty::IPSet
---

# awscc_guardduty_ip_set (Resource)

Resource Type definition for AWS::GuardDuty::IPSet



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `format` (String)
- `location` (String)
- `name` (String)

### Optional

- `activate` (Boolean)
- `detector_id` (String)
- `tags` (Attributes List) (see [below for nested schema](#nestedatt--tags))

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Required:

- `key` (String)
- `value` (String)

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_guardduty_ip_set.example <resource ID>
```
