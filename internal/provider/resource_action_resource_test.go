// Copyright (c) State of Geneva (Switzerland)
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const SELF string = "aria_resource_action.test"

func TestAccResourceActionResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{

			// Create and Read testing
			{
				Config: `
variable "test_project_id" {
  description = "Project where to generate test resources."
  type        = string
}

resource "aria_abx_action" "machine_test" {
  name            = "ARIA_PROVIDER_MACHINE_TEST"
  description     = "Temporary action generated by Aria provider's acceptance tests."
  runtime_name    = "python"
  memory_in_mb    = 128
  timeout_seconds = 60
  entrypoint      = "handler"
  dependencies    = []
  constants       = []
  secrets         = []
  source          = <<EOT
import os


def handler(*args, **kwargs):
    print('Global symbols :', globals())
    print('Environment variables :', os.environ)
    print('Call Arguments: ', args, kwargs)
EOT
  shared     = true
  project_id = var.test_project_id
}

resource "aria_resource_action" "test" {
  name          = aria_abx_action.machine_test.name
  display_name  = "Reset"
  description   = "Reset the machine."
  resource_type = "Cloud.vSphere.Machine"
  status        = "DRAFT"
  project_id    = var.test_project_id
  runnable_item = {
    id                = aria_abx_action.machine_test.id
    name              = aria_abx_action.machine_test.name
    project_id        = aria_abx_action.machine_test.project_id
    type              = "abx.action"
    input_parameters  = []
    output_parameters = []
  }
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(SELF, "id"),
					resource.TestCheckResourceAttr(SELF, "name", "ARIA_PROVIDER_MACHINE_TEST"),
					resource.TestCheckResourceAttr(SELF, "display_name", "Reset"),
					resource.TestCheckResourceAttr(SELF, "description", "Reset the machine."),
					resource.TestCheckResourceAttr(SELF, "resource_type", "Cloud.vSphere.Machine"),
					resource.TestCheckResourceAttr(SELF, "status", "DRAFT"),
					resource.TestCheckResourceAttrSet(SELF, "project_id"),
					resource.TestCheckResourceAttrSet(SELF, "org_id"),

					// Form definition generated by the platform
					resource.TestCheckResourceAttrSet(SELF, "form_definition.id"),
					resource.TestCheckResourceAttr(SELF, "form_definition.name", "ARIA_PROVIDER_MACHINE_TEST"),
					resource.TestCheckResourceAttr(SELF, "form_definition.type", "requestForm"),
					resource.TestCheckResourceAttrSet(SELF, "form_definition.form"),
					resource.TestCheckResourceAttr(SELF, "form_definition.form_format", "JSON"),
					resource.TestCheckResourceAttr(SELF, "form_definition.styles", ""),
					resource.TestCheckResourceAttrPair(SELF, "form_definition.source_id", SELF, "id"),
					resource.TestCheckResourceAttr(SELF, "form_definition.source_type", "resource.action"),
					resource.TestCheckResourceAttrSet(SELF, "form_definition.tenant"),
					resource.TestCheckResourceAttr(SELF, "form_definition.status", "ON"),
				),
			},

			// Update (miscellaneous) testing
			{
				Config: `
variable "test_project_id" {
  description = "Project where to generate test resources."
  type        = string
}

resource "aria_abx_action" "machine_test" {
  name            = "ARIA_PROVIDER_MACHINE_TEST"
  description     = "Temporary action generated by Aria provider's acceptance tests."
  runtime_name    = "python"
  memory_in_mb    = 128
  timeout_seconds = 60
  entrypoint      = "handler"
  dependencies    = []
  constants       = []
  secrets         = []
  source          = <<EOT
import os


def handler(*args, **kwargs):
    print('Global symbols :', globals())
    print('Environment variables :', os.environ)
    print('Call Arguments :', args, kwargs)
EOT
  shared     = true
  project_id = var.test_project_id
}

resource "aria_resource_action" "test" {
  name          = "${aria_abx_action.machine_test.name}_RENAMED"
  display_name  = "RESET"
  description   = "Reset the machine..."
  resource_type = "Cloud.vSphere.Machine"
  status        = "RELEASED"
  project_id    = var.test_project_id
  runnable_item = {
    id                = aria_abx_action.machine_test.id
    name              = aria_abx_action.machine_test.name
    project_id        = aria_abx_action.machine_test.project_id
    type              = "abx.action"
    input_parameters  = []
    output_parameters = []
  }
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(SELF, "id"),
					resource.TestCheckResourceAttr(SELF, "name", "ARIA_PROVIDER_MACHINE_TEST_RENAMED"),
					resource.TestCheckResourceAttr(SELF, "display_name", "RESET"),
					resource.TestCheckResourceAttr(SELF, "description", "Reset the machine..."),
					resource.TestCheckResourceAttr(SELF, "resource_type", "Cloud.vSphere.Machine"),
					resource.TestCheckResourceAttr(SELF, "status", "RELEASED"),
					resource.TestCheckResourceAttrSet(SELF, "project_id"),
					resource.TestCheckResourceAttrSet(SELF, "org_id"),

					// Form definition generated by the platform
					resource.TestCheckResourceAttrSet(SELF, "form_definition.id"),
					resource.TestCheckResourceAttr(SELF, "form_definition.name", "ARIA_PROVIDER_MACHINE_TEST_RENAMED"),
					resource.TestCheckResourceAttr(SELF, "form_definition.type", "requestForm"),
					resource.TestCheckResourceAttrSet(SELF, "form_definition.form"),
					resource.TestCheckResourceAttr(SELF, "form_definition.form_format", "JSON"),
					resource.TestCheckResourceAttr(SELF, "form_definition.styles", ""),
					resource.TestCheckResourceAttrPair(SELF, "form_definition.source_id", SELF, "id"),
					resource.TestCheckResourceAttr(SELF, "form_definition.source_type", "resource.action"),
					resource.TestCheckResourceAttrSet(SELF, "form_definition.tenant"),
					resource.TestCheckResourceAttr(SELF, "form_definition.status", "ON"),
				),
			},

			// Update (display name) testing
			{
				Config: `
variable "test_project_id" {
  description = "Project where to generate test resources."
  type        = string
}

resource "aria_abx_action" "machine_test" {
  name            = "ARIA_PROVIDER_MACHINE_TEST"
  description     = "Temporary action generated by Aria provider's acceptance tests."
  runtime_name    = "python"
  memory_in_mb    = 128
  timeout_seconds = 60
  entrypoint      = "handler"
  dependencies    = []
  constants       = []
  secrets         = []
  source          = <<EOT
import os


def handler(*args, **kwargs):
    print('Global symbols :', globals())
    print('Environment variables :', os.environ)
    print('Call Arguments :', args, kwargs)
EOT
  shared     = true
  project_id = var.test_project_id
}

resource "aria_resource_action" "test" {
  name          = "${aria_abx_action.machine_test.name}_RENAMED"
  display_name  = "RESET_BIS"
  description   = "Reset the machine..."
  resource_type = "Cloud.vSphere.Machine"
  status        = "RELEASED"
  project_id    = var.test_project_id
  runnable_item = {
    id                = aria_abx_action.machine_test.id
    name              = aria_abx_action.machine_test.name
    project_id        = aria_abx_action.machine_test.project_id
    type              = "abx.action"
    input_parameters  = []
    output_parameters = []
  }
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(SELF, "id"),
					resource.TestCheckResourceAttr(SELF, "name", "ARIA_PROVIDER_MACHINE_TEST_RENAMED"),
					resource.TestCheckResourceAttr(SELF, "display_name", "RESET_BIS"),
					resource.TestCheckResourceAttr(SELF, "description", "Reset the machine..."),
					resource.TestCheckResourceAttr(SELF, "resource_type", "Cloud.vSphere.Machine"),
					resource.TestCheckResourceAttr(SELF, "status", "RELEASED"),
					resource.TestCheckResourceAttrSet(SELF, "project_id"),
					resource.TestCheckResourceAttrSet(SELF, "org_id"),

					// Form definition generated by the platform
					resource.TestCheckResourceAttrSet(SELF, "form_definition.id"),
					resource.TestCheckResourceAttr(SELF, "form_definition.name", "ARIA_PROVIDER_MACHINE_TEST_RENAMED"),
					resource.TestCheckResourceAttr(SELF, "form_definition.type", "requestForm"),
					resource.TestCheckResourceAttrSet(SELF, "form_definition.form"),
					resource.TestCheckResourceAttr(SELF, "form_definition.form_format", "JSON"),
					resource.TestCheckResourceAttr(SELF, "form_definition.styles", ""),
					resource.TestCheckResourceAttrPair(SELF, "form_definition.source_id", SELF, "id"),
					resource.TestCheckResourceAttr(SELF, "form_definition.source_type", "resource.action"),
					resource.TestCheckResourceAttrSet(SELF, "form_definition.tenant"),
					resource.TestCheckResourceAttr(SELF, "form_definition.status", "ON"),
				),
			},

			// Unscoping testing
			{
				Config: `
variable "test_project_id" {
  description = "Project where to generate test resources."
  type        = string
}

resource "aria_abx_action" "machine_test" {
  name            = "ARIA_PROVIDER_MACHINE_TEST"
  description     = "Temporary action generated by Aria provider's acceptance tests."
  runtime_name    = "python"
  memory_in_mb    = 128
  timeout_seconds = 60
  entrypoint      = "handler"
  dependencies    = []
  constants       = []
  secrets         = []
  source          = <<EOT
import os


def handler(*args, **kwargs):
    print('Global symbols :', globals())
    print('Environment variables :', os.environ)
    print('Call Arguments :', args, kwargs)
EOT
  shared     = true
  project_id = var.test_project_id
}

resource "aria_resource_action" "test" {
  name          = "${aria_abx_action.machine_test.name}_RENAMED"
  display_name  = "RESET"
  description   = "Reset the machine..."
  resource_type = "Cloud.vSphere.Machine"
  status        = "DRAFT"
  project_id    = ""
  runnable_item = {
    id                = aria_abx_action.machine_test.id
    name              = aria_abx_action.machine_test.name
    project_id        = aria_abx_action.machine_test.project_id
    type              = "abx.action"
    input_parameters  = []
    output_parameters = []
  }
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(SELF, "id"),
					resource.TestCheckResourceAttr(SELF, "name", "ARIA_PROVIDER_MACHINE_TEST_RENAMED"),
					resource.TestCheckResourceAttr(SELF, "display_name", "RESET"),
					resource.TestCheckResourceAttr(SELF, "description", "Reset the machine..."),
					resource.TestCheckResourceAttr(SELF, "resource_type", "Cloud.vSphere.Machine"),
					resource.TestCheckResourceAttr(SELF, "status", "DRAFT"),
					resource.TestCheckResourceAttr(SELF, "project_id", ""),
					resource.TestCheckResourceAttrSet(SELF, "org_id"),

					// Form definition generated by the platform
					resource.TestCheckResourceAttrSet(SELF, "form_definition.id"),
					resource.TestCheckResourceAttr(SELF, "form_definition.name", "ARIA_PROVIDER_MACHINE_TEST_RENAMED"),
					resource.TestCheckResourceAttr(SELF, "form_definition.type", "requestForm"),
					resource.TestCheckResourceAttrSet(SELF, "form_definition.form"),
					resource.TestCheckResourceAttr(SELF, "form_definition.form_format", "JSON"),
					resource.TestCheckResourceAttr(SELF, "form_definition.styles", ""),
					resource.TestCheckResourceAttrPair(SELF, "form_definition.source_id", SELF, "id"),
					resource.TestCheckResourceAttr(SELF, "form_definition.source_type", "resource.action"),
					resource.TestCheckResourceAttrSet(SELF, "form_definition.tenant"),
					resource.TestCheckResourceAttr(SELF, "form_definition.status", "ON"),
				),
			},

			// Add criteria testing
			{
				Config: `
variable "test_project_id" {
  description = "Project where to generate test resources."
  type        = string
}

resource "aria_abx_action" "machine_test" {
  name            = "ARIA_PROVIDER_MACHINE_TEST"
  description     = "Temporary action generated by Aria provider's acceptance tests."
  runtime_name    = "python"
  memory_in_mb    = 128
  timeout_seconds = 60
  entrypoint      = "handler"
  dependencies    = []
  constants       = []
  secrets         = []
  source          = <<EOT
import os


def handler(*args, **kwargs):
    print('Global symbols :', globals())
    print('Environment variables :', os.environ)
    print('Call Arguments :', args, kwargs)
EOT
  shared     = true
  project_id = var.test_project_id
}

resource "aria_resource_action" "test" {
  name          = "${aria_abx_action.machine_test.name}_RENAMED"
  display_name  = "RESET"
  description   = "Reset the machine..."
  resource_type = "Cloud.vSphere.Machine"
  status        = "DRAFT"
  project_id    = ""
  runnable_item = {
    id                = aria_abx_action.machine_test.id
    name              = aria_abx_action.machine_test.name
    project_id        = aria_abx_action.machine_test.project_id
    type              = "abx.action"
    input_parameters  = []
    output_parameters = []
  }

  criteria = jsonencode({
    matchExpression = [
      {
        and = [
          {
            key      = "$${properties.osType}"
            operator = "eq"
            value    = "WINDOWS"
          }
        ]
      }
    ]
  })
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(SELF, "id"),
					resource.TestCheckResourceAttr(SELF, "name", "ARIA_PROVIDER_MACHINE_TEST_RENAMED"),
					resource.TestCheckResourceAttr(SELF, "display_name", "RESET"),
					resource.TestCheckResourceAttr(SELF, "description", "Reset the machine..."),
					resource.TestCheckResourceAttr(SELF, "resource_type", "Cloud.vSphere.Machine"),
					resource.TestCheckResourceAttr(SELF, "status", "DRAFT"),
					resource.TestCheckResourceAttr(SELF, "project_id", ""),
					resource.TestCheckResourceAttrSet(SELF, "org_id"),

					// Form definition generated by the platform
					resource.TestCheckResourceAttrSet(SELF, "form_definition.id"),
					resource.TestCheckResourceAttr(SELF, "form_definition.name", "ARIA_PROVIDER_MACHINE_TEST_RENAMED"),
					resource.TestCheckResourceAttr(SELF, "form_definition.type", "requestForm"),
					resource.TestCheckResourceAttrSet(SELF, "form_definition.form"),
					resource.TestCheckResourceAttr(SELF, "form_definition.form_format", "JSON"),
					resource.TestCheckResourceAttr(SELF, "form_definition.styles", ""),
					resource.TestCheckResourceAttrPair(SELF, "form_definition.source_id", SELF, "id"),
					resource.TestCheckResourceAttr(SELF, "form_definition.source_type", "resource.action"),
					resource.TestCheckResourceAttrSet(SELF, "form_definition.tenant"),
					resource.TestCheckResourceAttr(SELF, "form_definition.status", "ON"),
				),
			},

			// Update criteria testing
			{
				Config: `
variable "test_project_id" {
  description = "Project where to generate test resources."
  type        = string
}

resource "aria_abx_action" "machine_test" {
  name            = "ARIA_PROVIDER_MACHINE_TEST"
  description     = "Temporary action generated by Aria provider's acceptance tests."
  runtime_name    = "python"
  memory_in_mb    = 128
  timeout_seconds = 60
  entrypoint      = "handler"
  dependencies    = []
  constants       = []
  secrets         = []
  source          = <<EOT
import os


def handler(*args, **kwargs):
    print('Global symbols :', globals())
    print('Environment variables :', os.environ)
    print('Call Arguments :', args, kwargs)
EOT
  shared     = true
  project_id = var.test_project_id
}

resource "aria_resource_action" "test" {
  name          = "${aria_abx_action.machine_test.name}_RENAMED"
  display_name  = "RESET"
  description   = "Reset the machine..."
  resource_type = "Cloud.vSphere.Machine"
  status        = "DRAFT"
  project_id    = ""
  runnable_item = {
    id                = aria_abx_action.machine_test.id
    name              = aria_abx_action.machine_test.name
    project_id        = aria_abx_action.machine_test.project_id
    type              = "abx.action"
    input_parameters  = []
    output_parameters = []
  }

  criteria = jsonencode({
    matchExpression = [
      {
        and = [
          {
            key      = "$${properties.osType}"
            operator = "eq"
            value    = "WINDOWS"
          },
          {
            key      = "$${properties.totalMemoryMB}"
            operator = "greaterThan"
            value    = "1024"
          },
          {
            key      = "$${properties.tags}"
            operator = "hasAny"
            value = {
              matchExpression = [
                {
                  and = [
                    {
                      key      = "key"
                      operator = "eq"
                      value    = "env"
                    },
                    {
                      key      = "value"
                      operator = "eq"
                      value    = "REC"
                    }
                  ]
                }
              ]
            }
          }
        ]
      }
    ]
  })
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(SELF, "id"),
					resource.TestCheckResourceAttr(SELF, "name", "ARIA_PROVIDER_MACHINE_TEST_RENAMED"),
					resource.TestCheckResourceAttr(SELF, "display_name", "RESET"),
					resource.TestCheckResourceAttr(SELF, "description", "Reset the machine..."),
					resource.TestCheckResourceAttr(SELF, "resource_type", "Cloud.vSphere.Machine"),
					resource.TestCheckResourceAttr(SELF, "status", "DRAFT"),
					resource.TestCheckResourceAttr(SELF, "project_id", ""),
					resource.TestCheckResourceAttrSet(SELF, "org_id"),

					// Form definition generated by the platform
					resource.TestCheckResourceAttrSet(SELF, "form_definition.id"),
					resource.TestCheckResourceAttr(SELF, "form_definition.name", "ARIA_PROVIDER_MACHINE_TEST_RENAMED"),
					resource.TestCheckResourceAttr(SELF, "form_definition.type", "requestForm"),
					resource.TestCheckResourceAttrSet(SELF, "form_definition.form"),
					resource.TestCheckResourceAttr(SELF, "form_definition.form_format", "JSON"),
					resource.TestCheckResourceAttr(SELF, "form_definition.styles", ""),
					resource.TestCheckResourceAttrPair(SELF, "form_definition.source_id", SELF, "id"),
					resource.TestCheckResourceAttr(SELF, "form_definition.source_type", "resource.action"),
					resource.TestCheckResourceAttrSet(SELF, "form_definition.tenant"),
					resource.TestCheckResourceAttr(SELF, "form_definition.status", "ON"),
				),
			},
			// ImportState testing
			/* TODO https://github.com/davidfischer-ch/terraform-provider-aria/issues/32
			   {
			     ResourceName:      SELF,
			     ImportState:       true,
			     ImportStateVerify: true,
			   }, */
			// Delete testing automatically occurs in TestCase
			// TODO Check https://developer.hashicorp.com/terraform/plugin/sdkv2/testing/acceptance-tests/testcase#checkdestroy
		},
	})
}
