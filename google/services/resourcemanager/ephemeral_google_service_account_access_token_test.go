// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package resourcemanager_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-google/google/acctest"
)

func TestAccEphemeralGoogleServiceAccountAccessToken_basic(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccEphemeralGoogleServiceAccountAccessToken_basic(context),
			},
		},
	})
}

func testAccEphemeralGoogleServiceAccountAccessToken_basic(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_service_account" "default" {
  account_id   = "tf-test-%{random_suffix}"
  display_name = "Acceptance test impersonated service account"
}

ephemeral "google_service_account_access_token" "test" {
  target_service_account   = google_service_account.default.email
  scopes                   = ["userinfo-email", "cloud-platform"]
}
`, context)
}
