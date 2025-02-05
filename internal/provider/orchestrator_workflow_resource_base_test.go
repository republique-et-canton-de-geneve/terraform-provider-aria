// Copyright (c) State of Geneva (Switzerland)
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccOrchestratorWorkflowBaseResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: `
resource "aria_orchestrator_category" "root" {
  name      = "TEST_ARIA_PROVIDER"
  type      = "WorkflowCategory"
  parent_id = ""
}

locals {
  input_forms = [
    {
      layout = {
        pages = []
      }
      schema = {}
    }
  ]
}

resource "aria_orchestrator_workflow" "test" {
  name        = "Test Workflow"
  description = "Workflow generated by the acceptance tests of Aria provider."
  category_id = aria_orchestrator_category.root.id
  version     = "0.1.0"

  position = { x = 100, y = 50 }

  restart_mode            = 1 # resume
  resume_from_failed_mode = 0 # default

  attrib        = jsonencode([])
  presentation  = jsonencode({})
  workflow_item = jsonencode([])

  input_parameters  = []
  output_parameters = []

  input_forms = jsonencode(local.input_forms)

  wait_on_catalog = false # Make tests faster

  lifecycle {
    postcondition {
      condition     = jsondecode(self.input_forms) == local.input_forms
      error_message = "Attribute Input Forms is not what we expected: ${self.input_forms}"
    }
  }
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("aria_orchestrator_workflow.test", "id"),
					resource.TestCheckResourceAttr("aria_orchestrator_workflow.test", "name", "Test Workflow"),
					resource.TestCheckResourceAttr("aria_orchestrator_workflow.test", "description", "Workflow generated by the acceptance tests of Aria provider."),
					resource.TestCheckResourceAttr("aria_orchestrator_workflow.test", "version", "0.1.0"),
					resource.TestMatchResourceAttr("aria_orchestrator_workflow.test", "version_id", regexp.MustCompile("[0-9a-f]{40}")),
					resource.TestCheckResourceAttr("aria_orchestrator_workflow.test", "position.x", "100"),
					resource.TestCheckResourceAttr("aria_orchestrator_workflow.test", "position.y", "50"),
					resource.TestCheckResourceAttr("aria_orchestrator_workflow.test", "allowed_operations", "vef"),
					resource.TestCheckResourceAttr("aria_orchestrator_workflow.test", "attrib", "[]"),
					resource.TestCheckResourceAttr("aria_orchestrator_workflow.test", "object_name", "workflow:name=generic"),
					resource.TestCheckResourceAttr("aria_orchestrator_workflow.test", "presentation", "{}"),
					resource.TestCheckResourceAttr("aria_orchestrator_workflow.test", "restart_mode", "1"),
					resource.TestCheckResourceAttr("aria_orchestrator_workflow.test", "resume_from_failed_mode", "0"),
					resource.TestCheckResourceAttr("aria_orchestrator_workflow.test", "root_name", "item0"),
					resource.TestCheckResourceAttr("aria_orchestrator_workflow.test", "workflow_item", "[]"),
					resource.TestCheckResourceAttr("aria_orchestrator_workflow.test", "api_version", "6.0.0"),
					resource.TestCheckResourceAttr("aria_orchestrator_workflow.test", "editor_version", "2.0"),
					resource.TestCheckResourceAttr("aria_orchestrator_workflow.test", "force_delete", "false"),
					resource.TestCheckResourceAttr("aria_orchestrator_workflow.test", "wait_on_catalog", "false"),
					resource.TestCheckResourceAttrPair(
						"aria_orchestrator_workflow.test", "category_id",
						"aria_orchestrator_category.root", "id",
					),
				),
			},
			// Update and Read testing
			{
				Config: `
resource "aria_orchestrator_category" "root" {
  name      = "TEST_ARIA_PROVIDER"
  type      = "WorkflowCategory"
  parent_id = ""
}

locals {
  input_forms = [
    {
      layout = {
        pages = []
      }
      options = {
        externalValidations = []
      }
      schema = {}
    }
  ]
}

resource "aria_orchestrator_workflow" "test" {
  name        = "Test Workflow Renamed"
  description = "Workflow generated by the acceptance tests of Aria provider (updated)."
  category_id = aria_orchestrator_category.root.id
  version     = "0.2.0"

  position = { x = 60, y = 10 }

  restart_mode            = 0 # skip
  resume_from_failed_mode = 2 # disabled

  attrib        = jsonencode([])
  presentation  = jsonencode({})
  workflow_item = jsonencode([])

  input_parameters  = []
  output_parameters = []

  force_delete    = true
  wait_on_catalog = false # Make tests faster

  input_forms = jsonencode(local.input_forms)

  lifecycle {
    postcondition {
      condition     = jsondecode(self.input_forms) == local.input_forms
      error_message = "Attribute Input Forms is not what we expected: ${self.input_forms}"
    }
  }
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("aria_orchestrator_workflow.test", "id"),
					resource.TestCheckResourceAttr("aria_orchestrator_workflow.test", "name", "Test Workflow Renamed"),
					resource.TestCheckResourceAttr("aria_orchestrator_workflow.test", "description", "Workflow generated by the acceptance tests of Aria provider (updated)."),
					resource.TestCheckResourceAttr("aria_orchestrator_workflow.test", "version", "0.2.0"),
					resource.TestMatchResourceAttr("aria_orchestrator_workflow.test", "version_id", regexp.MustCompile("[0-9a-f]{40}")),
					resource.TestCheckResourceAttr("aria_orchestrator_workflow.test", "position.x", "60"),
					resource.TestCheckResourceAttr("aria_orchestrator_workflow.test", "position.y", "10"),
					resource.TestCheckResourceAttr("aria_orchestrator_workflow.test", "allowed_operations", "vef"),
					resource.TestCheckResourceAttr("aria_orchestrator_workflow.test", "attrib", "[]"),
					resource.TestCheckResourceAttr("aria_orchestrator_workflow.test", "object_name", "workflow:name=generic"),
					resource.TestCheckResourceAttr("aria_orchestrator_workflow.test", "presentation", "{}"),
					resource.TestCheckResourceAttr("aria_orchestrator_workflow.test", "restart_mode", "0"),
					resource.TestCheckResourceAttr("aria_orchestrator_workflow.test", "resume_from_failed_mode", "2"),
					resource.TestCheckResourceAttr("aria_orchestrator_workflow.test", "root_name", "item0"),
					resource.TestCheckResourceAttr("aria_orchestrator_workflow.test", "workflow_item", "[]"),
					resource.TestCheckResourceAttr("aria_orchestrator_workflow.test", "api_version", "6.0.0"),
					resource.TestCheckResourceAttr("aria_orchestrator_workflow.test", "editor_version", "2.0"),
					resource.TestCheckResourceAttr("aria_orchestrator_workflow.test", "force_delete", "true"),
					resource.TestCheckResourceAttr("aria_orchestrator_workflow.test", "wait_on_catalog", "false"),
					resource.TestCheckResourceAttrPair(
						"aria_orchestrator_workflow.test", "category_id",
						"aria_orchestrator_category.root", "id",
					),
				),
			},
			// Change category and Read testing
			{
				Config: `
resource "aria_orchestrator_category" "root" {
  name      = "TEST_ARIA_PROVIDER"
  type      = "WorkflowCategory"
  parent_id = ""
}

resource "aria_orchestrator_category" "test" {
  name      = "Test"
  type      = "WorkflowCategory"
  parent_id = aria_orchestrator_category.root.id
}

locals {
  input_forms = [
    {
      layout = {
        pages = []
      }
      options = {
        externalValidations = []
      }
      schema = {}
    }
  ]
}

resource "aria_orchestrator_workflow" "test" {
  name        = "Test Workflow Renamed"
  description = "Workflow generated by the acceptance tests of Aria provider (updated)."
  category_id = aria_orchestrator_category.test.id
  version     = "0.2.0"

  position = { x = 60, y = 10 }

  restart_mode            = 0 # skip
  resume_from_failed_mode = 2 # disabled

  attrib        = jsonencode([])
  presentation  = jsonencode({})
  workflow_item = jsonencode([])

  input_parameters  = []
  output_parameters = []

  input_forms = jsonencode(local.input_forms)

  force_delete    = true
  wait_on_catalog = false # Make tests faster

  lifecycle {
    postcondition {
      condition     = jsondecode(self.input_forms) == local.input_forms
      error_message = "Attribute Input Forms is not what we expected: ${self.input_forms}"
    }
  }
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("aria_orchestrator_workflow.test", "id"),
					resource.TestCheckResourceAttr("aria_orchestrator_workflow.test", "name", "Test Workflow Renamed"),
					resource.TestCheckResourceAttr("aria_orchestrator_workflow.test", "description", "Workflow generated by the acceptance tests of Aria provider (updated)."),
					resource.TestCheckResourceAttr("aria_orchestrator_workflow.test", "version", "0.2.0"),
					resource.TestMatchResourceAttr("aria_orchestrator_workflow.test", "version_id", regexp.MustCompile("[0-9a-f]{40}")),
					resource.TestCheckResourceAttr("aria_orchestrator_workflow.test", "position.x", "60"),
					resource.TestCheckResourceAttr("aria_orchestrator_workflow.test", "position.y", "10"),
					resource.TestCheckResourceAttr("aria_orchestrator_workflow.test", "allowed_operations", "vef"),
					resource.TestCheckResourceAttr("aria_orchestrator_workflow.test", "attrib", "[]"),
					resource.TestCheckResourceAttr("aria_orchestrator_workflow.test", "object_name", "workflow:name=generic"),
					resource.TestCheckResourceAttr("aria_orchestrator_workflow.test", "presentation", "{}"),
					resource.TestCheckResourceAttr("aria_orchestrator_workflow.test", "restart_mode", "0"),
					resource.TestCheckResourceAttr("aria_orchestrator_workflow.test", "resume_from_failed_mode", "2"),
					resource.TestCheckResourceAttr("aria_orchestrator_workflow.test", "root_name", "item0"),
					resource.TestCheckResourceAttr("aria_orchestrator_workflow.test", "workflow_item", "[]"),
					resource.TestCheckResourceAttr("aria_orchestrator_workflow.test", "api_version", "6.0.0"),
					resource.TestCheckResourceAttr("aria_orchestrator_workflow.test", "editor_version", "2.0"),
					resource.TestCheckResourceAttr("aria_orchestrator_workflow.test", "force_delete", "true"),
					resource.TestCheckResourceAttr("aria_orchestrator_workflow.test", "wait_on_catalog", "false"),
					resource.TestCheckResourceAttrPair(
						"aria_orchestrator_workflow.test", "category_id",
						"aria_orchestrator_category.test", "id",
					),
				),
			},
			// Change parameters and Read testing
			{
				Config: `
resource "aria_orchestrator_category" "root" {
  name      = "TEST_ARIA_PROVIDER"
  type      = "WorkflowCategory"
  parent_id = ""
}

resource "aria_orchestrator_category" "test" {
  name      = "Test"
  type      = "WorkflowCategory"
  parent_id = aria_orchestrator_category.root.id
}

locals {
  input_forms = [
    {
      layout = {
        pages = []
      }
      options = {
        externalValidations = []
      }
      schema = {}
    }
  ]
}

resource "aria_orchestrator_workflow" "test" {
  name        = "Test Workflow Renamed"
  description = "Workflow generated by the acceptance tests of Aria provider (updated)."
  category_id = aria_orchestrator_category.test.id
  version     = "1.0.0"

  position = { x = 60, y = 10 }

  restart_mode            = 0 # skip
  resume_from_failed_mode = 2 # disabled

  attrib        = jsonencode([])
  presentation  = jsonencode({})
  workflow_item = jsonencode([])

  input_parameters = [
    {
      name        = "vraHost"
      type        = "VRA:Host"
      description = ""
    },
    {
      name        = "deploymentId"
      type        = "string"
      description = ""
    }
  ]

  output_parameters = [
    {
      name       = "result"
      type       = "number"
      description = "0 = OK, Negative = Error, Positive = Number of values returned"
    },
    {
      name        = "errorText"
      type        = "string"
      description = "Error text, if any"
    },
    {
      name        = "outputText"
      type        = "string"
      description = "Result of running the SSH command"
    }
  ]

  input_forms = jsonencode(local.input_forms)

  force_delete    = true
  wait_on_catalog = false # Make tests faster

  lifecycle {
    postcondition {
      condition = self.input_parameters == tolist([
        {
          name        = "vraHost"
          type        = "VRA:Host"
          description = ""
        },
        {
          name        = "deploymentId"
          type        = "string"
          description = ""
        }
      ])
      error_message = "Input parameters is not what we expect: ${jsonencode(self.input_parameters)}"
    }
    postcondition {
      condition = self.output_parameters == tolist([
        {
          name       = "result"
          type       = "number"
          description = "0 = OK, Negative = Error, Positive = Number of values returned"
        },
        {
          name        = "errorText"
          type        = "string"
          description = "Error text, if any"
        },
        {
          name        = "outputText"
          type        = "string"
          description = "Result of running the SSH command"
        }
      ])
      error_message = "Output parameters is not what we expect: ${jsonencode(self.output_parameters)}"
    }
    postcondition {
      condition     = jsondecode(self.input_forms) == local.input_forms
      error_message = "Attribute Input Forms is not what we expected: ${self.input_forms}"
    }
  }
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("aria_orchestrator_workflow.test", "id"),
					resource.TestCheckResourceAttr("aria_orchestrator_workflow.test", "name", "Test Workflow Renamed"),
					resource.TestCheckResourceAttr("aria_orchestrator_workflow.test", "description", "Workflow generated by the acceptance tests of Aria provider (updated)."),
					resource.TestCheckResourceAttr("aria_orchestrator_workflow.test", "version", "1.0.0"),
					resource.TestMatchResourceAttr("aria_orchestrator_workflow.test", "version_id", regexp.MustCompile("[0-9a-f]{40}")),
					resource.TestCheckResourceAttr("aria_orchestrator_workflow.test", "position.x", "60"),
					resource.TestCheckResourceAttr("aria_orchestrator_workflow.test", "position.y", "10"),
					resource.TestCheckResourceAttr("aria_orchestrator_workflow.test", "allowed_operations", "vef"),
					resource.TestCheckResourceAttr("aria_orchestrator_workflow.test", "attrib", "[]"),
					resource.TestCheckResourceAttr("aria_orchestrator_workflow.test", "object_name", "workflow:name=generic"),
					resource.TestCheckResourceAttr("aria_orchestrator_workflow.test", "presentation", "{}"),
					resource.TestCheckResourceAttr("aria_orchestrator_workflow.test", "restart_mode", "0"),
					resource.TestCheckResourceAttr("aria_orchestrator_workflow.test", "resume_from_failed_mode", "2"),
					resource.TestCheckResourceAttr("aria_orchestrator_workflow.test", "root_name", "item0"),
					resource.TestCheckResourceAttr("aria_orchestrator_workflow.test", "workflow_item", "[]"),
					resource.TestCheckResourceAttr("aria_orchestrator_workflow.test", "api_version", "6.0.0"),
					resource.TestCheckResourceAttr("aria_orchestrator_workflow.test", "editor_version", "2.0"),
					resource.TestCheckResourceAttr("aria_orchestrator_workflow.test", "force_delete", "true"),
					resource.TestCheckResourceAttr("aria_orchestrator_workflow.test", "wait_on_catalog", "false"),
					resource.TestCheckResourceAttrPair(
						"aria_orchestrator_workflow.test", "category_id",
						"aria_orchestrator_category.test", "id",
					),
				),
			},
			// ImportState testing
			// FIXME https://github.com/davidfischer-ch/terraform-provider-aria/issues/122
			/*{
				ResourceName:      "aria_orchestrator_workflow.test",
				ImportState:       true,
				ImportStateVerify: true,

				// Prevent diff on force_delete & wait_on_catalog fields
				ImportStateVerifyIgnore: []string{"force_delete", "wait_on_catalog"},
			},*/
			// Delete testing automatically occurs in TestCase
		},
	})
}
