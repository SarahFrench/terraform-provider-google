// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/gkehub2/Scope.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/examples/base_configs/iam_test_file.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package gkehub2_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/hashicorp/terraform-provider-google/google/acctest"
	"github.com/hashicorp/terraform-provider-google/google/envvar"
)

func TestAccGKEHub2ScopeIamBindingGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
		"role":          "roles/viewer",
		"project":       envvar.GetTestProjectFromEnv(),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccGKEHub2ScopeIamBinding_basicGenerated(context),
			},
			{
				ResourceName:      "google_gke_hub_scope_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/global/scopes/%s roles/viewer", envvar.GetTestProjectFromEnv(), fmt.Sprintf("tf-test-my-scope%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Test Iam Binding update
				Config: testAccGKEHub2ScopeIamBinding_updateGenerated(context),
			},
			{
				ResourceName:      "google_gke_hub_scope_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/global/scopes/%s roles/viewer", envvar.GetTestProjectFromEnv(), fmt.Sprintf("tf-test-my-scope%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccGKEHub2ScopeIamMemberGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
		"role":          "roles/viewer",
		"project":       envvar.GetTestProjectFromEnv(),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				// Test Iam Member creation (no update for member, no need to test)
				Config: testAccGKEHub2ScopeIamMember_basicGenerated(context),
			},
			{
				ResourceName:      "google_gke_hub_scope_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/global/scopes/%s roles/viewer user:admin@hashicorptest.com", envvar.GetTestProjectFromEnv(), fmt.Sprintf("tf-test-my-scope%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccGKEHub2ScopeIamPolicyGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
		"role":          "roles/viewer",
		"project":       envvar.GetTestProjectFromEnv(),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccGKEHub2ScopeIamPolicy_basicGenerated(context),
				Check:  resource.TestCheckResourceAttrSet("data.google_gke_hub_scope_iam_policy.foo", "policy_data"),
			},
			{
				ResourceName:      "google_gke_hub_scope_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/global/scopes/%s", envvar.GetTestProjectFromEnv(), fmt.Sprintf("tf-test-my-scope%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccGKEHub2ScopeIamPolicy_emptyBinding(context),
			},
			{
				ResourceName:      "google_gke_hub_scope_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/global/scopes/%s", envvar.GetTestProjectFromEnv(), fmt.Sprintf("tf-test-my-scope%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccGKEHub2ScopeIamMember_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_gke_hub_scope" "scope" {
  scope_id = "tf-test-my-scope%{random_suffix}"
  namespace_labels = {
      keyb = "valueb"
      keya = "valuea"
      keyc = "valuec" 
  }
  labels = {
      keyb = "valueb"
      keya = "valuea"
      keyc = "valuec" 
  }
}

resource "google_gke_hub_scope_iam_member" "foo" {
  project = google_gke_hub_scope.scope.project
  scope_id = google_gke_hub_scope.scope.scope_id
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}
`, context)
}

func testAccGKEHub2ScopeIamPolicy_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_gke_hub_scope" "scope" {
  scope_id = "tf-test-my-scope%{random_suffix}"
  namespace_labels = {
      keyb = "valueb"
      keya = "valuea"
      keyc = "valuec" 
  }
  labels = {
      keyb = "valueb"
      keya = "valuea"
      keyc = "valuec" 
  }
}

data "google_iam_policy" "foo" {
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
  }
}

resource "google_gke_hub_scope_iam_policy" "foo" {
  project = google_gke_hub_scope.scope.project
  scope_id = google_gke_hub_scope.scope.scope_id
  policy_data = data.google_iam_policy.foo.policy_data
}

data "google_gke_hub_scope_iam_policy" "foo" {
  project = google_gke_hub_scope.scope.project
  scope_id = google_gke_hub_scope.scope.scope_id
  depends_on = [
    google_gke_hub_scope_iam_policy.foo
  ]
}
`, context)
}

func testAccGKEHub2ScopeIamPolicy_emptyBinding(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_gke_hub_scope" "scope" {
  scope_id = "tf-test-my-scope%{random_suffix}"
  namespace_labels = {
      keyb = "valueb"
      keya = "valuea"
      keyc = "valuec" 
  }
  labels = {
      keyb = "valueb"
      keya = "valuea"
      keyc = "valuec" 
  }
}

data "google_iam_policy" "foo" {
}

resource "google_gke_hub_scope_iam_policy" "foo" {
  project = google_gke_hub_scope.scope.project
  scope_id = google_gke_hub_scope.scope.scope_id
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccGKEHub2ScopeIamBinding_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_gke_hub_scope" "scope" {
  scope_id = "tf-test-my-scope%{random_suffix}"
  namespace_labels = {
      keyb = "valueb"
      keya = "valuea"
      keyc = "valuec" 
  }
  labels = {
      keyb = "valueb"
      keya = "valuea"
      keyc = "valuec" 
  }
}

resource "google_gke_hub_scope_iam_binding" "foo" {
  project = google_gke_hub_scope.scope.project
  scope_id = google_gke_hub_scope.scope.scope_id
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}
`, context)
}

func testAccGKEHub2ScopeIamBinding_updateGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_gke_hub_scope" "scope" {
  scope_id = "tf-test-my-scope%{random_suffix}"
  namespace_labels = {
      keyb = "valueb"
      keya = "valuea"
      keyc = "valuec" 
  }
  labels = {
      keyb = "valueb"
      keya = "valuea"
      keyc = "valuec" 
  }
}

resource "google_gke_hub_scope_iam_binding" "foo" {
  project = google_gke_hub_scope.scope.project
  scope_id = google_gke_hub_scope.scope.scope_id
  role = "%{role}"
  members = ["user:admin@hashicorptest.com", "user:gterraformtest1@gmail.com"]
}
`, context)
}
