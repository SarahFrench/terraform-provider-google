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
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/privateca/CaPool.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/examples/base_configs/iam_test_file.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package privateca_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/hashicorp/terraform-provider-google/google/acctest"
	"github.com/hashicorp/terraform-provider-google/google/envvar"
	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
)

func TestAccPrivatecaCaPoolIamBindingGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":           acctest.RandString(t, 10),
		"role":                    "roles/privateca.certificateManager",
		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccPrivatecaCaPoolIamBinding_basicGenerated(context),
			},
			{
				ResourceName:      "google_privateca_ca_pool_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/caPools/%s roles/privateca.certificateManager", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-my-pool%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Test Iam Binding update
				Config: testAccPrivatecaCaPoolIamBinding_updateGenerated(context),
			},
			{
				ResourceName:      "google_privateca_ca_pool_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/caPools/%s roles/privateca.certificateManager", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-my-pool%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccPrivatecaCaPoolIamMemberGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":           acctest.RandString(t, 10),
		"role":                    "roles/privateca.certificateManager",
		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				// Test Iam Member creation (no update for member, no need to test)
				Config: testAccPrivatecaCaPoolIamMember_basicGenerated(context),
			},
			{
				ResourceName:      "google_privateca_ca_pool_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/caPools/%s roles/privateca.certificateManager user:admin@hashicorptest.com", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-my-pool%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccPrivatecaCaPoolIamPolicyGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":           acctest.RandString(t, 10),
		"role":                    "roles/privateca.certificateManager",
		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccPrivatecaCaPoolIamPolicy_basicGenerated(context),
				Check:  resource.TestCheckResourceAttrSet("data.google_privateca_ca_pool_iam_policy.foo", "policy_data"),
			},
			{
				ResourceName:      "google_privateca_ca_pool_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/caPools/%s", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-my-pool%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccPrivatecaCaPoolIamPolicy_emptyBinding(context),
			},
			{
				ResourceName:      "google_privateca_ca_pool_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/caPools/%s", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-my-pool%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccPrivatecaCaPoolIamBindingGenerated_withCondition(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":           acctest.RandString(t, 10),
		"role":                    "roles/privateca.certificateManager",
		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccPrivatecaCaPoolIamBinding_withConditionGenerated(context),
			},
			{
				ResourceName:      "google_privateca_ca_pool_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/caPools/%s roles/privateca.certificateManager %s", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-my-pool%s", context["random_suffix"]), context["condition_title"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccPrivatecaCaPoolIamBindingGenerated_withAndWithoutCondition(t *testing.T) {
	// Multiple fine-grained resources
	acctest.SkipIfVcr(t)
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":           acctest.RandString(t, 10),
		"role":                    "roles/privateca.certificateManager",
		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccPrivatecaCaPoolIamBinding_withAndWithoutConditionGenerated(context),
			},
			{
				ResourceName:      "google_privateca_ca_pool_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/caPools/%s roles/privateca.certificateManager", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-my-pool%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:      "google_privateca_ca_pool_iam_binding.foo2",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/caPools/%s roles/privateca.certificateManager %s", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-my-pool%s", context["random_suffix"]), context["condition_title"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:      "google_privateca_ca_pool_iam_binding.foo3",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/caPools/%s roles/privateca.certificateManager %s", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-my-pool%s", context["random_suffix"]), context["condition_title_no_desc"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccPrivatecaCaPoolIamMemberGenerated_withCondition(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":           acctest.RandString(t, 10),
		"role":                    "roles/privateca.certificateManager",
		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccPrivatecaCaPoolIamMember_withConditionGenerated(context),
			},
			{
				ResourceName:      "google_privateca_ca_pool_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/caPools/%s roles/privateca.certificateManager user:admin@hashicorptest.com %s", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-my-pool%s", context["random_suffix"]), context["condition_title"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccPrivatecaCaPoolIamMemberGenerated_withAndWithoutCondition(t *testing.T) {
	// Multiple fine-grained resources
	acctest.SkipIfVcr(t)
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":           acctest.RandString(t, 10),
		"role":                    "roles/privateca.certificateManager",
		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccPrivatecaCaPoolIamMember_withAndWithoutConditionGenerated(context),
			},
			{
				ResourceName:      "google_privateca_ca_pool_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/caPools/%s roles/privateca.certificateManager user:admin@hashicorptest.com", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-my-pool%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:      "google_privateca_ca_pool_iam_member.foo2",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/caPools/%s roles/privateca.certificateManager user:admin@hashicorptest.com %s", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-my-pool%s", context["random_suffix"]), context["condition_title"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:      "google_privateca_ca_pool_iam_member.foo3",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/caPools/%s roles/privateca.certificateManager user:admin@hashicorptest.com %s", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-my-pool%s", context["random_suffix"]), context["condition_title_no_desc"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccPrivatecaCaPoolIamPolicyGenerated_withCondition(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":           acctest.RandString(t, 10),
		"role":                    "roles/privateca.certificateManager",
		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	// Test should have 2 bindings: one with a description and one without. Any < chars are converted to a unicode character by the API.
	expectedPolicyData := acctest.Nprintf(`{"bindings":[{"condition":{"description":"%{condition_desc}","expression":"%{condition_expr}","title":"%{condition_title}"},"members":["user:admin@hashicorptest.com"],"role":"%{role}"},{"condition":{"expression":"%{condition_expr}","title":"%{condition_title}-no-description"},"members":["user:admin@hashicorptest.com"],"role":"%{role}"}]}`, context)
	expectedPolicyData = strings.Replace(expectedPolicyData, "<", "\\u003c", -1)

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccPrivatecaCaPoolIamPolicy_withConditionGenerated(context),
				Check: resource.ComposeAggregateTestCheckFunc(
					// TODO(SarahFrench) - uncomment once https://github.com/GoogleCloudPlatform/magic-modules/pull/6466 merged
					// resource.TestCheckResourceAttr("data.google_iam_policy.foo", "policy_data", expectedPolicyData),
					resource.TestCheckResourceAttr("google_privateca_ca_pool_iam_policy.foo", "policy_data", expectedPolicyData),
					resource.TestCheckResourceAttrWith("data.google_iam_policy.foo", "policy_data", tpgresource.CheckGoogleIamPolicy),
				),
			},
			{
				ResourceName:      "google_privateca_ca_pool_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/caPools/%s", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-my-pool%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccPrivatecaCaPoolIamMember_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_privateca_ca_pool" "default" {
  name = "tf-test-my-pool%{random_suffix}"
  location = "us-central1"
  tier = "ENTERPRISE"
  publishing_options {
    publish_ca_cert = true
    publish_crl = true
  }
  labels = {
    foo = "bar"
  }
}

resource "google_privateca_ca_pool_iam_member" "foo" {
  ca_pool = google_privateca_ca_pool.default.id
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}
`, context)
}

func testAccPrivatecaCaPoolIamPolicy_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_privateca_ca_pool" "default" {
  name = "tf-test-my-pool%{random_suffix}"
  location = "us-central1"
  tier = "ENTERPRISE"
  publishing_options {
    publish_ca_cert = true
    publish_crl = true
  }
  labels = {
    foo = "bar"
  }
}

data "google_iam_policy" "foo" {
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
  }
}

resource "google_privateca_ca_pool_iam_policy" "foo" {
  ca_pool = google_privateca_ca_pool.default.id
  policy_data = data.google_iam_policy.foo.policy_data
}

data "google_privateca_ca_pool_iam_policy" "foo" {
  ca_pool = google_privateca_ca_pool.default.id
  depends_on = [
    google_privateca_ca_pool_iam_policy.foo
  ]
}
`, context)
}

func testAccPrivatecaCaPoolIamPolicy_emptyBinding(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_privateca_ca_pool" "default" {
  name = "tf-test-my-pool%{random_suffix}"
  location = "us-central1"
  tier = "ENTERPRISE"
  publishing_options {
    publish_ca_cert = true
    publish_crl = true
  }
  labels = {
    foo = "bar"
  }
}

data "google_iam_policy" "foo" {
}

resource "google_privateca_ca_pool_iam_policy" "foo" {
  ca_pool = google_privateca_ca_pool.default.id
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccPrivatecaCaPoolIamBinding_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_privateca_ca_pool" "default" {
  name = "tf-test-my-pool%{random_suffix}"
  location = "us-central1"
  tier = "ENTERPRISE"
  publishing_options {
    publish_ca_cert = true
    publish_crl = true
  }
  labels = {
    foo = "bar"
  }
}

resource "google_privateca_ca_pool_iam_binding" "foo" {
  ca_pool = google_privateca_ca_pool.default.id
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}
`, context)
}

func testAccPrivatecaCaPoolIamBinding_updateGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_privateca_ca_pool" "default" {
  name = "tf-test-my-pool%{random_suffix}"
  location = "us-central1"
  tier = "ENTERPRISE"
  publishing_options {
    publish_ca_cert = true
    publish_crl = true
  }
  labels = {
    foo = "bar"
  }
}

resource "google_privateca_ca_pool_iam_binding" "foo" {
  ca_pool = google_privateca_ca_pool.default.id
  role = "%{role}"
  members = ["user:admin@hashicorptest.com", "user:gterraformtest1@gmail.com"]
}
`, context)
}

func testAccPrivatecaCaPoolIamBinding_withConditionGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_privateca_ca_pool" "default" {
  name = "tf-test-my-pool%{random_suffix}"
  location = "us-central1"
  tier = "ENTERPRISE"
  publishing_options {
    publish_ca_cert = true
    publish_crl = true
  }
  labels = {
    foo = "bar"
  }
}

resource "google_privateca_ca_pool_iam_binding" "foo" {
  ca_pool = google_privateca_ca_pool.default.id
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
  condition {
    title       = "%{condition_title}"
    description = "%{condition_desc}"
    expression  = "%{condition_expr}"
  }
}
`, context)
}

func testAccPrivatecaCaPoolIamBinding_withAndWithoutConditionGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_privateca_ca_pool" "default" {
  name = "tf-test-my-pool%{random_suffix}"
  location = "us-central1"
  tier = "ENTERPRISE"
  publishing_options {
    publish_ca_cert = true
    publish_crl = true
  }
  labels = {
    foo = "bar"
  }
}

resource "google_privateca_ca_pool_iam_binding" "foo" {
  ca_pool = google_privateca_ca_pool.default.id
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}

resource "google_privateca_ca_pool_iam_binding" "foo2" {
  ca_pool = google_privateca_ca_pool.default.id
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
  condition {
    title       = "%{condition_title}"
    description = "%{condition_desc}"
    expression  = "%{condition_expr}"
  }
}

resource "google_privateca_ca_pool_iam_binding" "foo3" {
  ca_pool = google_privateca_ca_pool.default.id
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
  condition {
    # Check that lack of description doesn't cause any issues
    # Relates to issue : https://github.com/hashicorp/terraform-provider-google/issues/8701
    title       = "%{condition_title_no_desc}"
    expression  = "%{condition_expr_no_desc}"
  }
}
`, context)
}

func testAccPrivatecaCaPoolIamMember_withConditionGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_privateca_ca_pool" "default" {
  name = "tf-test-my-pool%{random_suffix}"
  location = "us-central1"
  tier = "ENTERPRISE"
  publishing_options {
    publish_ca_cert = true
    publish_crl = true
  }
  labels = {
    foo = "bar"
  }
}

resource "google_privateca_ca_pool_iam_member" "foo" {
  ca_pool = google_privateca_ca_pool.default.id
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
  condition {
    title       = "%{condition_title}"
    description = "%{condition_desc}"
    expression  = "%{condition_expr}"
  }
}
`, context)
}

func testAccPrivatecaCaPoolIamMember_withAndWithoutConditionGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_privateca_ca_pool" "default" {
  name = "tf-test-my-pool%{random_suffix}"
  location = "us-central1"
  tier = "ENTERPRISE"
  publishing_options {
    publish_ca_cert = true
    publish_crl = true
  }
  labels = {
    foo = "bar"
  }
}

resource "google_privateca_ca_pool_iam_member" "foo" {
  ca_pool = google_privateca_ca_pool.default.id
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}

resource "google_privateca_ca_pool_iam_member" "foo2" {
  ca_pool = google_privateca_ca_pool.default.id
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
  condition {
    title       = "%{condition_title}"
    description = "%{condition_desc}"
    expression  = "%{condition_expr}"
  }
}

resource "google_privateca_ca_pool_iam_member" "foo3" {
  ca_pool = google_privateca_ca_pool.default.id
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
  condition {
    # Check that lack of description doesn't cause any issues
    # Relates to issue : https://github.com/hashicorp/terraform-provider-google/issues/8701
    title       = "%{condition_title_no_desc}"
    expression  = "%{condition_expr_no_desc}"
  }
}
`, context)
}

func testAccPrivatecaCaPoolIamPolicy_withConditionGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_privateca_ca_pool" "default" {
  name = "tf-test-my-pool%{random_suffix}"
  location = "us-central1"
  tier = "ENTERPRISE"
  publishing_options {
    publish_ca_cert = true
    publish_crl = true
  }
  labels = {
    foo = "bar"
  }
}

data "google_iam_policy" "foo" {
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
    condition {
      # Check that lack of description doesn't cause any issues
      # Relates to issue : https://github.com/hashicorp/terraform-provider-google/issues/8701
      title       = "%{condition_title_no_desc}"
      expression  = "%{condition_expr_no_desc}"
    }
  }
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
    condition {
      title       = "%{condition_title}"
      description = "%{condition_desc}"
      expression  = "%{condition_expr}"
    }
  }
}

resource "google_privateca_ca_pool_iam_policy" "foo" {
  ca_pool = google_privateca_ca_pool.default.id
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}
