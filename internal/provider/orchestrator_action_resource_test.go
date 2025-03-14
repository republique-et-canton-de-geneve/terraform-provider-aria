// Copyright (c) State of Geneva (Switzerland)
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccOrchestratorActionResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: `
locals {
	script = <<EOT
// Get VRA Host by type
var vrahosts = VraHostManager.findHostsByType("vra-onprem");
for each ( var host in vrahosts ) {
    var vRAHost = host;
    System.warn("vRA Host is : "+vRAHost.vraHost)
    }
return vRAHost;

EOT
}

resource "aria_orchestrator_action" "test" {
  name                 = "getVRAHost"
  module               = "aria_provider_tests"
  fqn                  = "aria_provider_tests/getVRAHost"
  description          = "Temporary action generated by Aria provider's acceptance tests."
  version              = "1.0.0"
  runtime              = ""
  runtime_memory_limit = 128 * 1024 * 1024
  runtime_timeout      = 30
  script               = local.script
  input_parameters     = []
  output_type          = "VRA:Host"

  lifecycle {
    postcondition {
      condition     = self.script == local.script
      error_message = "Script must be ${local.script}, actual ${self.script}"
    }
  }
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("aria_orchestrator_action.test", "id"),
					resource.TestCheckResourceAttr("aria_orchestrator_action.test", "name", "getVRAHost"),
					resource.TestCheckResourceAttr("aria_orchestrator_action.test", "module", "aria_provider_tests"),
					resource.TestCheckResourceAttr("aria_orchestrator_action.test", "fqn", "aria_provider_tests/getVRAHost"),
					resource.TestCheckResourceAttr("aria_orchestrator_action.test", "description", "Temporary action generated by Aria provider's acceptance tests."),
					resource.TestCheckResourceAttr("aria_orchestrator_action.test", "version", "1.0.0"),
					resource.TestCheckResourceAttr("aria_orchestrator_action.test", "runtime", ""),
					resource.TestCheckResourceAttr("aria_orchestrator_action.test", "runtime_memory_limit", "134217728"),
					resource.TestCheckResourceAttr("aria_orchestrator_action.test", "runtime_timeout", "30"),
					resource.TestCheckResourceAttr("aria_orchestrator_action.test", "output_type", "VRA:Host"),
				),
			},
			// Update and Read testing
			{
				Config: `
locals {
	script = <<EOT
print('Hello World')
EOT
}

resource "aria_orchestrator_action" "test" {
  name                 = "getVRAHost"
  module               = "aria_provider_tests_bis"
  fqn                  = "aria_provider_tests_bis/getVRAHost"
  description          = "Temporary action generated by Aria provider's acceptance tests (bis)."
  version              = "1.0.1"
  runtime              = "python:3.10"
  runtime_memory_limit = 10000000
  runtime_timeout      = 5
  script               = local.script
  input_parameters     = []
  output_type          = "string"

  lifecycle {
    postcondition {
      condition     = self.script == local.script
      error_message = "Script must be ${local.script}, actual ${self.script}"
    }
  }
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("aria_orchestrator_action.test", "id"),
					resource.TestCheckResourceAttr("aria_orchestrator_action.test", "name", "getVRAHost"),
					resource.TestCheckResourceAttr("aria_orchestrator_action.test", "module", "aria_provider_tests_bis"),
					resource.TestCheckResourceAttr("aria_orchestrator_action.test", "fqn", "aria_provider_tests_bis/getVRAHost"),
					resource.TestCheckResourceAttr("aria_orchestrator_action.test", "description", "Temporary action generated by Aria provider's acceptance tests (bis)."),
					resource.TestCheckResourceAttr("aria_orchestrator_action.test", "version", "1.0.1"),
					resource.TestCheckResourceAttr("aria_orchestrator_action.test", "runtime", "python:3.10"),
					resource.TestCheckResourceAttr("aria_orchestrator_action.test", "runtime_memory_limit", "10000000"),
					resource.TestCheckResourceAttr("aria_orchestrator_action.test", "runtime_timeout", "5"),
					resource.TestCheckResourceAttr("aria_orchestrator_action.test", "output_type", "string"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "aria_orchestrator_action.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccOrchestratorActionWithInputParametersResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: `
locals {
	script = <<EOT
if (vraHost == null || deploymentId == null) return null;

var url = "/deployment/api/deployments/" + deploymentId;
var deployment = System.getModule("com.vmware.vra.extensibility.plugin.rest").getObjectFromUrl(vraHost, url);

return deployment;

EOT
}

resource "aria_orchestrator_action" "test" {
  name                 = "getDeploymentById" # You have to manage this boilerplate for some time
  module               = "aria_provider_tests"  # Having name, module and fqn = module/name...
  fqn                  = "aria_provider_tests/getDeploymentById"
  description          = "Return the deployment object matching given ID."
  version              = "1.0.0"
  runtime              = "" # javascript
  runtime_memory_limit = 0
  runtime_timeout      = 0
  script               = local.script
  output_type          = "Any"

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

  lifecycle {
    postcondition {
      condition     = self.script == local.script
      error_message = "Script must be ${local.script}, actual ${self.script}"
    }
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
  }
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("aria_orchestrator_action.test", "id"),
					resource.TestCheckResourceAttr("aria_orchestrator_action.test", "name", "getDeploymentById"),
					resource.TestCheckResourceAttr("aria_orchestrator_action.test", "module", "aria_provider_tests"),
					resource.TestCheckResourceAttr("aria_orchestrator_action.test", "fqn", "aria_provider_tests/getDeploymentById"),
					resource.TestCheckResourceAttr("aria_orchestrator_action.test", "description", "Return the deployment object matching given ID."),
					resource.TestCheckResourceAttr("aria_orchestrator_action.test", "version", "1.0.0"),
					resource.TestCheckResourceAttr("aria_orchestrator_action.test", "runtime", ""),
					resource.TestCheckResourceAttr("aria_orchestrator_action.test", "runtime_memory_limit", "0"),
					resource.TestCheckResourceAttr("aria_orchestrator_action.test", "runtime_timeout", "0"),
					resource.TestCheckResourceAttr("aria_orchestrator_action.test", "output_type", "Any"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "aria_orchestrator_action.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccOrchestratorActionDeleteConvergeResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create interlinked resources that Terraform doesn't know they are, three levels deep
			{
				Config: `
# The way the configuration is written Terraform is not aware that test_b depends on test_a, and
# test_c depends on test_b. Destroy should normally do not work properly if deletion wasn't retried.
# Automatically in case of conflits.

resource "aria_orchestrator_action" "test_a" {
  name                 = "actionA"
  module               = "aria_provider_tests"
  fqn                  = "aria_provider_tests/actionA"
  description          = "An action used by actionB."
  version              = "1.0.0"
  runtime              = "" # javascript
  runtime_memory_limit = 0
  runtime_timeout      = 0
  script               = ""
  input_parameters     = []
  output_type          = "Any"
  # force_delete = true
}

resource "aria_orchestrator_action" "test_b" {
  name                 = "actionB"
  module               = "aria_provider_tests"
  fqn                  = "aria_provider_tests/actionB"
  description          = "An action using actionA."
  version              = "1.0.0"
  runtime              = "" # javascript
  runtime_memory_limit = 0
  runtime_timeout      = 0
  script               = <<EOT
 var actionA = System.getModule("aria_provider_tests").actionA();
 EOT
  input_parameters     = []
  output_type          = "Any"
}

resource "aria_orchestrator_action" "test_c" {
  name                 = "actionC"
  module               = "aria_provider_tests"
  fqn                  = "aria_provider_tests/actionC"
  description          = "An action using actionB."
  version              = "1.0.0"
  runtime              = "" # javascript
  runtime_memory_limit = 0
  runtime_timeout      = 0
  script               = <<EOT
 var actionA = System.getModule("aria_provider_tests").actionB();
 EOT
  input_parameters     = []
  output_type          = "Any"
}
`,
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccOrchestratorActionForceDeleteResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create interlinked resources that Terraform doesn't know they are
			{
				Config: `
# The way the configuration is written Terraform is not aware that test_e depends on test_d.
# Destroy of test_d should not be possible unless test_e is destroyed too.
# In this case test_d's force_delete is true so it will be forced.

resource "aria_orchestrator_action" "test_d" {
  name                 = "actionD"
  module               = "aria_provider_tests"
  fqn                  = "aria_provider_tests/actionD"
  description          = "An action used by actionE."
  version              = "1.0.0"
  runtime              = "" # javascript
  runtime_memory_limit = 0
  runtime_timeout      = 0
  script               = ""
  input_parameters     = []
  output_type          = "Any"
  force_delete         = true
}

resource "aria_orchestrator_action" "test_e" {
  name                 = "actionE"
  module               = "aria_provider_tests"
  fqn                  = "aria_provider_tests/actionE"
  description          = "An action using actionD."
  version              = "1.0.0"
  runtime              = "" # javascript
  runtime_memory_limit = 0
  runtime_timeout      = 0
  script               = <<EOT
 var actionA = System.getModule("aria_provider_tests").actionD();
 EOT
  input_parameters     = []
  output_type          = "Any"
}
`,
			},
			// Destroy test_d shouldn't be possible if not forced. Here it is...
			{
				Config: `
resource "aria_orchestrator_action" "test_e" {
  name                 = "actionE"
  module               = "aria_provider_tests"
  fqn                  = "aria_provider_tests/actionE"
  description          = "An action using actionD."
  version              = "1.0.0"
  runtime              = "" # javascript
  runtime_memory_limit = 0
  runtime_timeout      = 0
  script               = <<EOT
 var actionA = System.getModule("aria_provider_tests").actionD();
 EOT
  input_parameters     = []
  output_type          = "Any"
}
`,
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
