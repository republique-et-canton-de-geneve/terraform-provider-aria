---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "aria_catalog_source Resource - aria"
subcategory: ""
description: |-
  Catalog source resource
---

# aria_catalog_source (Resource)

Catalog source resource

## Example Usage

```terraform
# Publish the Cloud Templates of a Project Libray using a Catalog Source ---------------------------

# variables.tf

variable "library_project_id" {
  description = "Identifier of the project containing Cloud templates to publish."
  type        = string
}

# main.tf

resource "aria_catalog_source" "library_project_cloud_templates" {
  name        = "Cloud Templates Catalog Source"
  description = "Publish some Cloud templates from a library project."
  project_id  = var.library_project_id
  type_id     = "com.vmw.abx.actions"

  config = {
    source_project_id = var.library_project_id
  }
}

# Create a Workflow and make it available using a Catalog Source -----------------------------------

# Method 1
#
# Using only the catalog source's waiting mechanism.
#
# Workflows's integration attribute will be null (cannot be guarantee).
# The aria integration data source is used to retrieve the integration endpoint.

# data.tf

data "aria_integration" "workflows" {
  type_id = "com.vmw.vro.workflow"
}

# main.tf

resource "aria_orchestrator_category" "my_company" {
  name      = "MyCompany"
  type      = "WorkflowCategory"
  parent_id = ""
}

# Example is dummy and contains no code
resource "aria_orchestrator_workflow" "dummy" {
  name        = "Dummy Workflow for Catalog Source"
  description = "Workflows doing nothing particular."
  category_id = aria_orchestrator_category.my_company.id
  version     = "0.1.0"

  position = { x = 100, y = 50 }

  restart_mode            = 1 # resume
  resume_from_failed_mode = 0 # default

  attrib        = jsonencode([])
  presentation  = jsonencode({})
  workflow_item = jsonencode([])

  input_parameters  = []
  output_parameters = []

  input_forms = jsonencode([
    {
      layout = {
        pages = []
      }
      schema = {}
    }
  ])

  wait_imported = false
}

resource "aria_catalog_source" "dummy" {
  name        = "Dummy Workflow Catalog Source"
  description = "Publish the dummy workflow."
  type_id     = data.aria_integration.workflows.type_id

  config = {
    workflows = [
      {
        id          = aria_orchestrator_workflow.dummy.id
        name        = aria_orchestrator_workflow.dummy.name
        description = aria_orchestrator_workflow.dummy.description
        version     = aria_orchestrator_workflow.dummy.version
        integration = {
          name                        = data.aria_integration.workflows.name
          endpoint_configuration_link = data.aria_integration.workflows.endpoint_configuration_link
          endpoint_uri                = data.aria_integration.workflows.endpoint_uri
        }
      }
    ]
  }

  # Refresh the catalog source every time the workflow is changed
  import_trigger = aria_orchestrator_workflow.dummy.version_id
}

# Create a Workflow and make it available using a Catalog Source -----------------------------------

# Method 2
#
# Using both workflow's andd catalog source's waiting mechanism.
#
# Workflows's integration attribute will be set.

# main.tf

resource "aria_orchestrator_category" "my_company" {
  name      = "MyCompany"
  type      = "WorkflowCategory"
  parent_id = ""
}

# Example is dummy and contains no code
resource "aria_orchestrator_workflow" "dummy" {
  name        = "Dummy Workflow for Catalog Source"
  description = "Workflows doing nothing particular."
  category_id = aria_orchestrator_category.my_company.id
  version     = "0.1.0"

  position = { x = 100, y = 50 }

  restart_mode            = 1 # resume
  resume_from_failed_mode = 0 # default

  attrib        = jsonencode([])
  presentation  = jsonencode({})
  workflow_item = jsonencode([])

  input_parameters  = []
  output_parameters = []

  input_forms = jsonencode([
    {
      layout = {
        pages = []
      }
      schema = {}
    }
  ])
}

resource "aria_catalog_source" "dummy" {
  name        = "Dummy Workflow Catalog Source"
  description = "Publish the dummy workflow."
  type_id     = "com.vmw.vro.workflow"

  config = {
    workflows = [
      {
        id          = aria_orchestrator_workflow.dummy.id
        name        = aria_orchestrator_workflow.dummy.name
        description = aria_orchestrator_workflow.dummy.description
        version     = aria_orchestrator_workflow.dummy.version
        integration = aria_orchestrator_workflow.dummy.integration
      }
    ]
  }

  # Refresh the catalog source every time the workflow is changed
  import_trigger = aria_orchestrator_workflow.dummy.version_id
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `config` (Attributes) Configuration (see [below for nested schema](#nestedatt--config))
- `description` (String) Describe the resource in few sentences
- `name` (String) Source name (e.g. getVRAHost)
- `type_id` (String) Source type (e.g. `com.vmw.vro.workflow`)

### Optional

- `import_trigger` (String) Set it to any value changing every time you want the catalog source to be refreshed.

One use case can be to ensure workflows are refreshed in service broker every time its changed, by using `workflow.version_id` as value for this.
- `project_id` (String) Project identifier. Empty or unset means available for all projects. (force recreation on change)
- `wait_imported` (Boolean) Wait for import to be completed (up to 15 minutes, checked every 30 seconds, default is true)

### Read-Only

- `created_at` (String) Creation timestamp (RFC3339)
- `created_by` (String) User who created the resource
- `global` (Boolean) Is it globally shared?
- `id` (String) Identifier
- `items_found` (Number) Number of existing items
- `items_imported` (Number) Number of imported items
- `last_import_completed_at` (String) Last import end timestamp (RFC3339)
- `last_import_errors` (List of String) Action input parameters
- `last_import_started_at` (String) Last import start timestamp (RFC3339)
- `last_updated_at` (String) Last update timestamp (RFC3339)
- `last_updated_by` (String) Last user who updated the resource

<a id="nestedatt--config"></a>
### Nested Schema for `config`

Optional:

- `source_project_id` (String) Project to make available (required for Cloud Templates or ABX Actions catalog sources)
- `workflows` (Attributes List) Workflows to make available (required for Orchestrator Worflows catalog sources) (see [below for nested schema](#nestedatt--config--workflows))

<a id="nestedatt--config--workflows"></a>
### Nested Schema for `config.workflows`

Required:

- `description` (String) Workflow description
- `id` (String) Identifier
- `integration` (Attributes) Integration (see [below for nested schema](#nestedatt--config--workflows--integration))
- `name` (String) Workflow name
- `version` (String) Workflow version

<a id="nestedatt--config--workflows--integration"></a>
### Nested Schema for `config.workflows.integration`

Required:

- `endpoint_configuration_link` (String) Integration endpoint configuration link
- `endpoint_uri` (String) Integration endpoint URI
- `name` (String) Integration name
