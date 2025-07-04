---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "aria_custom_resource Resource - aria"
subcategory: ""
description: |-
  Custom Resource resource
---

# aria_custom_resource (Resource)

Custom Resource resource

## Example Usage

```terraform
# variables.tf

variable "project_id" {
  type = string
}

# locals.tf

locals {
  source = <<EOT
import os

def handler(*args, **kwargs):
    print('Global symbols :', globals())
    print('Environment variables :', os.environ)
    print('Call Arguments: ', args, kwargs)
EOT
}

# constants.tf

resource "aria_abx_constant" "example" {
  name  = "THIS_IS_MY_CONSTANT"
  value = "42"
}

# main.tf

resource "aria_abx_action" "redis_create" {
  name            = "Custom.Redis.create"
  description     = "Provision an instance of a Redis server."
  runtime_name    = "python"
  memory_in_mb    = 128
  timeout_seconds = 60
  entrypoint      = "handler"
  dependencies    = []
  constants       = []
  secrets         = []
  source          = local.source
  shared          = true
  project_id      = var.project_id
}

resource "aria_abx_action" "redis_read" {
  name            = "Custom.Redis.read"
  description     = "Refresh properties by gathering the actual Redis instance attributes."
  runtime_name    = "python"
  memory_in_mb    = 128
  timeout_seconds = 60
  entrypoint      = "handler"
  dependencies    = []
  constants       = []
  secrets         = []
  source          = local.source
  shared          = true
  project_id      = var.project_id
}

resource "aria_abx_action" "redis_update" {
  name            = "Custom.Redis.update"
  description     = "Update Redis instance's attributes."
  runtime_name    = "python"
  memory_in_mb    = 128
  timeout_seconds = 60
  entrypoint      = "handler"
  dependencies    = []
  constants       = []
  secrets         = []
  source          = local.source
  shared          = true
  project_id      = var.project_id
}

resource "aria_abx_action" "redis_delete" {
  name            = "Custom.Redis.delete"
  description     = "Destroy the Redis instance."
  runtime_name    = "python"
  memory_in_mb    = 128
  timeout_seconds = 60
  entrypoint      = "handler"
  dependencies    = []
  constants       = []
  secrets         = []
  source          = local.source
  shared          = true
  project_id      = var.project_id
}

resource "aria_custom_resource" "redis" {
  display_name  = "Redis"
  description   = "Manage an instance of a Redis database."
  resource_type = "Custom.Redis"
  schema_type   = "ABX_USER_DEFINED"
  status        = "RELEASED"
  project_id    = var.project_id

  properties = {
    version = {
      name               = "version"
      title              = "Version"
      description        = "Instance version."
      type               = "string"
      encrypted          = false
      read_only          = false
      recreate_on_update = false
      one_of = [
        { const = "7.4", title = "7.4", encrypted = false },
        { const = "8.0", title = "8.0", encrypted = false }
      ]
    }
    description = {
      name               = "description"
      title              = "Description"
      description        = "Some description here."
      type               = "string"
      default            = jsonencode("No description given.")
      encrypted          = false
      read_only          = false
      recreate_on_update = false
    }
    storage_size = {
      name               = "storage_size"
      title              = "Storage Size"
      description        = "Storage size (MB)."
      type               = "integer"
      default            = 10 * 1024
      encrypted          = false
      read_only          = false
      recreate_on_update = false
      minimum            = 1 * 1024
      maximum            = 100 * 1024
    }
    secret = {
      name               = "secret"
      title              = "Secret"
      description        = "Secret key."
      type               = "string"
      encrypted          = true
      read_only          = false
      recreate_on_update = false
      min_length         = 16
      max_length         = 64
    }
  }

  create = {
    id                = aria_abx_action.redis_create.id
    name              = aria_abx_action.redis_create.name
    project_id        = aria_abx_action.redis_create.project_id
    type              = "abx.action"
    input_parameters  = []
    output_parameters = []
  }

  read = {
    id                = aria_abx_action.redis_read.id
    name              = aria_abx_action.redis_read.name
    project_id        = aria_abx_action.redis_read.project_id
    type              = "abx.action"
    input_parameters  = []
    output_parameters = []
  }

  update = {
    id                = aria_abx_action.redis_update.id
    name              = aria_abx_action.redis_update.name
    project_id        = aria_abx_action.redis_update.project_id
    type              = "abx.action"
    input_parameters  = []
    output_parameters = []
  }

  delete = {
    id                = aria_abx_action.redis_delete.id
    name              = aria_abx_action.redis_delete.name
    project_id        = aria_abx_action.redis_delete.project_id
    type              = "abx.action"
    input_parameters  = []
    output_parameters = []
  }
}

# Additional actions (aka Day 2), managed using relational resources
# This design is intentional for Terraform to be able to succesfully apply any changes

resource "aria_abx_action" "redis_backup" {
  name            = "Custom.Redis.backup"
  description     = "Backup the Redis database (its data)."
  runtime_name    = "python"
  memory_in_mb    = 128
  timeout_seconds = 60
  entrypoint      = "handler"
  dependencies    = []
  constants       = [aria_abx_constant.example.id]
  secrets         = []
  source          = local.source
  shared          = true
  project_id      = var.project_id
}

resource "aria_resource_action" "redis_backup" {
  name          = "backup"
  display_name  = "Backup data"
  description   = aria_abx_action.redis_backup.description
  status        = aria_custom_resource.redis.status
  resource_id   = aria_custom_resource.redis.id
  resource_type = aria_custom_resource.redis.resource_type
  project_id    = aria_custom_resource.redis.project_id
  runnable_item = {
    id                = aria_abx_action.redis_backup.id
    name              = aria_abx_action.redis_backup.name
    project_id        = aria_abx_action.redis_backup.project_id
    type              = "abx.action"
    input_parameters  = []
    output_parameters = []
  }
}

resource "aria_abx_action" "redis_restore" {
  name            = "Custom.Redis.restore"
  description     = "Restore the Redis database (its data)."
  runtime_name    = "python"
  memory_in_mb    = 128
  timeout_seconds = 60
  entrypoint      = "handler"
  dependencies    = []
  constants       = []
  secrets         = []
  source          = local.source
  shared          = true
  project_id      = var.project_id
}

resource "aria_resource_action" "redis_restore" {
  name          = "restore"
  display_name  = "Restore"
  description   = aria_abx_action.redis_restore.description
  status        = aria_custom_resource.redis.status
  resource_id   = aria_custom_resource.redis.id
  resource_type = aria_custom_resource.redis.resource_type
  project_id    = aria_custom_resource.redis.project_id
  runnable_item = {
    id                = aria_abx_action.redis_restore.id
    name              = aria_abx_action.redis_restore.name
    project_id        = aria_abx_action.redis_restore.project_id
    type              = "abx.action"
    input_parameters  = []
    output_parameters = []
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `create` (Attributes) Create action (see [below for nested schema](#nestedatt--create))
- `delete` (Attributes) Delete action (see [below for nested schema](#nestedatt--delete))
- `description` (String) Describe the resource in few sentences
- `display_name` (String) A friendly name
- `properties` (Attributes Map) Resource's properties (see [below for nested schema](#nestedatt--properties))
- `read` (Attributes) Read action (see [below for nested schema](#nestedatt--read))
- `resource_type` (String) Define the type (must be unique, e.g. `Custom.DB.PostgreSQL`) (force recreation on change)
- `update` (Attributes) Update action (see [below for nested schema](#nestedatt--update))

### Optional

- `project_id` (String) Project identifier. Empty or unset means available for all projects. (force recreation on change)
- `schema_type` (String) Type of resource, one of `ABX_USER_DEFINED` (and that's all, maybe)
- `status` (String) Resource status, one of `DRAFT`, `ON`, or `RELEASED`

### Read-Only

- `id` (String) Identifier
- `org_id` (String) Organization identifier

<a id="nestedatt--create"></a>
### Nested Schema for `create`

Required:

- `id` (String) Identifier
- `input_parameters` (Attributes List) (see [below for nested schema](#nestedatt--create--input_parameters))
- `name` (String) Runnable name
- `output_parameters` (Attributes List) (see [below for nested schema](#nestedatt--create--output_parameters))
- `project_id` (String) Project identifier
- `type` (String) Runnable type, either abx.action or vro.workflow

Optional:

- `endpoint_link` (String) Integration API endpoint (e.g. /resources/endpoints/8a430db3-924c-4d58-a29a-da811f9c992e)

<a id="nestedatt--create--input_parameters"></a>
### Nested Schema for `create.input_parameters`

Required:

- `description` (String) Describe the resource in few sentences
- `name` (String) Name
- `type` (String) Type


<a id="nestedatt--create--output_parameters"></a>
### Nested Schema for `create.output_parameters`

Required:

- `description` (String) Describe the resource in few sentences
- `name` (String) Name
- `type` (String) Type



<a id="nestedatt--delete"></a>
### Nested Schema for `delete`

Required:

- `id` (String) Identifier
- `input_parameters` (Attributes List) (see [below for nested schema](#nestedatt--delete--input_parameters))
- `name` (String) Runnable name
- `output_parameters` (Attributes List) (see [below for nested schema](#nestedatt--delete--output_parameters))
- `project_id` (String) Project identifier
- `type` (String) Runnable type, either abx.action or vro.workflow

Optional:

- `endpoint_link` (String) Integration API endpoint (e.g. /resources/endpoints/8a430db3-924c-4d58-a29a-da811f9c992e)

<a id="nestedatt--delete--input_parameters"></a>
### Nested Schema for `delete.input_parameters`

Required:

- `description` (String) Describe the resource in few sentences
- `name` (String) Name
- `type` (String) Type


<a id="nestedatt--delete--output_parameters"></a>
### Nested Schema for `delete.output_parameters`

Required:

- `description` (String) Describe the resource in few sentences
- `name` (String) Name
- `type` (String) Type



<a id="nestedatt--properties"></a>
### Nested Schema for `properties`

Required:

- `description` (String) Describe the resource in few sentences
- `encrypted` (Boolean) Encrypted?
- `name` (String) Name
- `read_only` (Boolean) Make the field read-only (in the form)
- `recreate_on_update` (Boolean) Mark this field as writable once (resource will be recreated on change)
- `title` (String) Title
- `type` (String) Type, one of `array`, `boolean`, `integer`, `object`, `number` or `string`.

Optional:

- `default` (String) Default value (JSON encoded)

We should have implemented this attribute as a dynamic type (and not JSON).
Unfortunately Terraform SDK returns this issue:
Dynamic types inside of collections are not currently supported in terraform-plugin-framework.
- `max_length` (Number) Maximum length (valid for a string)
- `maximum` (Number) Maximum value (inclusive, valid for an integer)
- `min_length` (Number) Minimum length (valid for a string)
- `minimum` (Number) Minimum value (inclusive, valid for an integer)
- `one_of` (Attributes List) Enumerate possible values (see [below for nested schema](#nestedatt--properties--one_of))
- `pattern` (String) Pattern (valid for a string)

<a id="nestedatt--properties--one_of"></a>
### Nested Schema for `properties.one_of`

Required:

- `const` (String) Technical value
- `encrypted` (Boolean) Encrypted?
- `title` (String) Display value



<a id="nestedatt--read"></a>
### Nested Schema for `read`

Required:

- `id` (String) Identifier
- `input_parameters` (Attributes List) (see [below for nested schema](#nestedatt--read--input_parameters))
- `name` (String) Runnable name
- `output_parameters` (Attributes List) (see [below for nested schema](#nestedatt--read--output_parameters))
- `project_id` (String) Project identifier
- `type` (String) Runnable type, either abx.action or vro.workflow

Optional:

- `endpoint_link` (String) Integration API endpoint (e.g. /resources/endpoints/8a430db3-924c-4d58-a29a-da811f9c992e)

<a id="nestedatt--read--input_parameters"></a>
### Nested Schema for `read.input_parameters`

Required:

- `description` (String) Describe the resource in few sentences
- `name` (String) Name
- `type` (String) Type


<a id="nestedatt--read--output_parameters"></a>
### Nested Schema for `read.output_parameters`

Required:

- `description` (String) Describe the resource in few sentences
- `name` (String) Name
- `type` (String) Type



<a id="nestedatt--update"></a>
### Nested Schema for `update`

Required:

- `id` (String) Identifier
- `input_parameters` (Attributes List) (see [below for nested schema](#nestedatt--update--input_parameters))
- `name` (String) Runnable name
- `output_parameters` (Attributes List) (see [below for nested schema](#nestedatt--update--output_parameters))
- `project_id` (String) Project identifier
- `type` (String) Runnable type, either abx.action or vro.workflow

Optional:

- `endpoint_link` (String) Integration API endpoint (e.g. /resources/endpoints/8a430db3-924c-4d58-a29a-da811f9c992e)

<a id="nestedatt--update--input_parameters"></a>
### Nested Schema for `update.input_parameters`

Required:

- `description` (String) Describe the resource in few sentences
- `name` (String) Name
- `type` (String) Type


<a id="nestedatt--update--output_parameters"></a>
### Nested Schema for `update.output_parameters`

Required:

- `description` (String) Describe the resource in few sentences
- `name` (String) Name
- `type` (String) Type
