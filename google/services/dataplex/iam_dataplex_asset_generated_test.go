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
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/dataplex/Asset.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/examples/base_configs/iam_test_file.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package dataplex_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/hashicorp/terraform-provider-google/google/acctest"
	"github.com/hashicorp/terraform-provider-google/google/envvar"
)

func TestAccDataplexAssetIamBindingGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
		"role":          "roles/viewer",
		"project_name":  envvar.GetTestProjectFromEnv(),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataplexAssetIamBinding_basicGenerated(context),
			},
			{
				ResourceName:      "google_dataplex_asset_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/lakes/%s/zones/%s/assets/%s roles/viewer", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-lake%s", context["random_suffix"]), fmt.Sprintf("tf-test-zone%s", context["random_suffix"]), fmt.Sprintf("tf-test-asset%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Test Iam Binding update
				Config: testAccDataplexAssetIamBinding_updateGenerated(context),
			},
			{
				ResourceName:      "google_dataplex_asset_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/lakes/%s/zones/%s/assets/%s roles/viewer", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-lake%s", context["random_suffix"]), fmt.Sprintf("tf-test-zone%s", context["random_suffix"]), fmt.Sprintf("tf-test-asset%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccDataplexAssetIamMemberGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
		"role":          "roles/viewer",
		"project_name":  envvar.GetTestProjectFromEnv(),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				// Test Iam Member creation (no update for member, no need to test)
				Config: testAccDataplexAssetIamMember_basicGenerated(context),
			},
			{
				ResourceName:      "google_dataplex_asset_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/lakes/%s/zones/%s/assets/%s roles/viewer user:admin@hashicorptest.com", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-lake%s", context["random_suffix"]), fmt.Sprintf("tf-test-zone%s", context["random_suffix"]), fmt.Sprintf("tf-test-asset%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccDataplexAssetIamPolicyGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
		"role":          "roles/viewer",
		"project_name":  envvar.GetTestProjectFromEnv(),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataplexAssetIamPolicy_basicGenerated(context),
				Check:  resource.TestCheckResourceAttrSet("data.google_dataplex_asset_iam_policy.foo", "policy_data"),
			},
			{
				ResourceName:      "google_dataplex_asset_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/lakes/%s/zones/%s/assets/%s", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-lake%s", context["random_suffix"]), fmt.Sprintf("tf-test-zone%s", context["random_suffix"]), fmt.Sprintf("tf-test-asset%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccDataplexAssetIamPolicy_emptyBinding(context),
			},
			{
				ResourceName:      "google_dataplex_asset_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/lakes/%s/zones/%s/assets/%s", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-lake%s", context["random_suffix"]), fmt.Sprintf("tf-test-zone%s", context["random_suffix"]), fmt.Sprintf("tf-test-asset%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccDataplexAssetIamMember_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_storage_bucket" "primary_bucket" {
  name          = "dataplex-bucket-%{random_suffix}"
  location      = "us-central1"
  uniform_bucket_level_access = true
  lifecycle {
    ignore_changes = [
      labels
    ]
  }

  project = "%{project_name}"
}

resource "google_dataplex_lake" "example" {
  name         = "tf-test-lake%{random_suffix}"
  location     = "us-central1"
  project = "%{project_name}"
}


resource "google_dataplex_zone" "example" {
  name         = "tf-test-zone%{random_suffix}"
  location     = "us-central1"
  lake = google_dataplex_lake.example.name
  type = "RAW"

  discovery_spec {
    enabled = false
  }


  resource_spec {
    location_type = "SINGLE_REGION"
  }

  project = "%{project_name}"
}



resource "google_dataplex_asset" "example" {
  name          = "tf-test-asset%{random_suffix}"
  location      = "us-central1"

  lake = google_dataplex_lake.example.name
  dataplex_zone = google_dataplex_zone.example.name
  discovery_spec {
    enabled = false
  }

  resource_spec {
    name = "projects/%{project_name}/buckets/dataplex-bucket-%{random_suffix}"
    type = "STORAGE_BUCKET"
  }

  project = "%{project_name}"
  depends_on = [
    google_storage_bucket.primary_bucket
  ]
}

resource "google_dataplex_asset_iam_member" "foo" {
  project = google_dataplex_asset.example.project
  location = google_dataplex_asset.example.location
  lake = google_dataplex_asset.example.lake
  dataplex_zone = google_dataplex_asset.example.dataplex_zone
  asset = google_dataplex_asset.example.name
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}
`, context)
}

func testAccDataplexAssetIamPolicy_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_storage_bucket" "primary_bucket" {
  name          = "dataplex-bucket-%{random_suffix}"
  location      = "us-central1"
  uniform_bucket_level_access = true
  lifecycle {
    ignore_changes = [
      labels
    ]
  }

  project = "%{project_name}"
}

resource "google_dataplex_lake" "example" {
  name         = "tf-test-lake%{random_suffix}"
  location     = "us-central1"
  project = "%{project_name}"
}


resource "google_dataplex_zone" "example" {
  name         = "tf-test-zone%{random_suffix}"
  location     = "us-central1"
  lake = google_dataplex_lake.example.name
  type = "RAW"

  discovery_spec {
    enabled = false
  }


  resource_spec {
    location_type = "SINGLE_REGION"
  }

  project = "%{project_name}"
}



resource "google_dataplex_asset" "example" {
  name          = "tf-test-asset%{random_suffix}"
  location      = "us-central1"

  lake = google_dataplex_lake.example.name
  dataplex_zone = google_dataplex_zone.example.name
  discovery_spec {
    enabled = false
  }

  resource_spec {
    name = "projects/%{project_name}/buckets/dataplex-bucket-%{random_suffix}"
    type = "STORAGE_BUCKET"
  }

  project = "%{project_name}"
  depends_on = [
    google_storage_bucket.primary_bucket
  ]
}

data "google_iam_policy" "foo" {
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
  }
}

resource "google_dataplex_asset_iam_policy" "foo" {
  project = google_dataplex_asset.example.project
  location = google_dataplex_asset.example.location
  lake = google_dataplex_asset.example.lake
  dataplex_zone = google_dataplex_asset.example.dataplex_zone
  asset = google_dataplex_asset.example.name
  policy_data = data.google_iam_policy.foo.policy_data
}

data "google_dataplex_asset_iam_policy" "foo" {
  project = google_dataplex_asset.example.project
  location = google_dataplex_asset.example.location
  lake = google_dataplex_asset.example.lake
  dataplex_zone = google_dataplex_asset.example.dataplex_zone
  asset = google_dataplex_asset.example.name
  depends_on = [
    google_dataplex_asset_iam_policy.foo
  ]
}
`, context)
}

func testAccDataplexAssetIamPolicy_emptyBinding(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_storage_bucket" "primary_bucket" {
  name          = "dataplex-bucket-%{random_suffix}"
  location      = "us-central1"
  uniform_bucket_level_access = true
  lifecycle {
    ignore_changes = [
      labels
    ]
  }

  project = "%{project_name}"
}

resource "google_dataplex_lake" "example" {
  name         = "tf-test-lake%{random_suffix}"
  location     = "us-central1"
  project = "%{project_name}"
}


resource "google_dataplex_zone" "example" {
  name         = "tf-test-zone%{random_suffix}"
  location     = "us-central1"
  lake = google_dataplex_lake.example.name
  type = "RAW"

  discovery_spec {
    enabled = false
  }


  resource_spec {
    location_type = "SINGLE_REGION"
  }

  project = "%{project_name}"
}



resource "google_dataplex_asset" "example" {
  name          = "tf-test-asset%{random_suffix}"
  location      = "us-central1"

  lake = google_dataplex_lake.example.name
  dataplex_zone = google_dataplex_zone.example.name
  discovery_spec {
    enabled = false
  }

  resource_spec {
    name = "projects/%{project_name}/buckets/dataplex-bucket-%{random_suffix}"
    type = "STORAGE_BUCKET"
  }

  project = "%{project_name}"
  depends_on = [
    google_storage_bucket.primary_bucket
  ]
}

data "google_iam_policy" "foo" {
}

resource "google_dataplex_asset_iam_policy" "foo" {
  project = google_dataplex_asset.example.project
  location = google_dataplex_asset.example.location
  lake = google_dataplex_asset.example.lake
  dataplex_zone = google_dataplex_asset.example.dataplex_zone
  asset = google_dataplex_asset.example.name
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccDataplexAssetIamBinding_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_storage_bucket" "primary_bucket" {
  name          = "dataplex-bucket-%{random_suffix}"
  location      = "us-central1"
  uniform_bucket_level_access = true
  lifecycle {
    ignore_changes = [
      labels
    ]
  }

  project = "%{project_name}"
}

resource "google_dataplex_lake" "example" {
  name         = "tf-test-lake%{random_suffix}"
  location     = "us-central1"
  project = "%{project_name}"
}


resource "google_dataplex_zone" "example" {
  name         = "tf-test-zone%{random_suffix}"
  location     = "us-central1"
  lake = google_dataplex_lake.example.name
  type = "RAW"

  discovery_spec {
    enabled = false
  }


  resource_spec {
    location_type = "SINGLE_REGION"
  }

  project = "%{project_name}"
}



resource "google_dataplex_asset" "example" {
  name          = "tf-test-asset%{random_suffix}"
  location      = "us-central1"

  lake = google_dataplex_lake.example.name
  dataplex_zone = google_dataplex_zone.example.name
  discovery_spec {
    enabled = false
  }

  resource_spec {
    name = "projects/%{project_name}/buckets/dataplex-bucket-%{random_suffix}"
    type = "STORAGE_BUCKET"
  }

  project = "%{project_name}"
  depends_on = [
    google_storage_bucket.primary_bucket
  ]
}

resource "google_dataplex_asset_iam_binding" "foo" {
  project = google_dataplex_asset.example.project
  location = google_dataplex_asset.example.location
  lake = google_dataplex_asset.example.lake
  dataplex_zone = google_dataplex_asset.example.dataplex_zone
  asset = google_dataplex_asset.example.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}
`, context)
}

func testAccDataplexAssetIamBinding_updateGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_storage_bucket" "primary_bucket" {
  name          = "dataplex-bucket-%{random_suffix}"
  location      = "us-central1"
  uniform_bucket_level_access = true
  lifecycle {
    ignore_changes = [
      labels
    ]
  }

  project = "%{project_name}"
}

resource "google_dataplex_lake" "example" {
  name         = "tf-test-lake%{random_suffix}"
  location     = "us-central1"
  project = "%{project_name}"
}


resource "google_dataplex_zone" "example" {
  name         = "tf-test-zone%{random_suffix}"
  location     = "us-central1"
  lake = google_dataplex_lake.example.name
  type = "RAW"

  discovery_spec {
    enabled = false
  }


  resource_spec {
    location_type = "SINGLE_REGION"
  }

  project = "%{project_name}"
}



resource "google_dataplex_asset" "example" {
  name          = "tf-test-asset%{random_suffix}"
  location      = "us-central1"

  lake = google_dataplex_lake.example.name
  dataplex_zone = google_dataplex_zone.example.name
  discovery_spec {
    enabled = false
  }

  resource_spec {
    name = "projects/%{project_name}/buckets/dataplex-bucket-%{random_suffix}"
    type = "STORAGE_BUCKET"
  }

  project = "%{project_name}"
  depends_on = [
    google_storage_bucket.primary_bucket
  ]
}

resource "google_dataplex_asset_iam_binding" "foo" {
  project = google_dataplex_asset.example.project
  location = google_dataplex_asset.example.location
  lake = google_dataplex_asset.example.lake
  dataplex_zone = google_dataplex_asset.example.dataplex_zone
  asset = google_dataplex_asset.example.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com", "user:gterraformtest1@gmail.com"]
}
`, context)
}
