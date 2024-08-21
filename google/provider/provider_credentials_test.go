// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package provider_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-google/google/acctest"
	"github.com/hashicorp/terraform-provider-google/google/envvar"
)

// Test different ways of providing credentials when using SDK and PF implemented resources + data sources
// Need to provision something to trigger errors when 'bad' credentials used
func TestAccProvider_credentials(t *testing.T) {
	testCases := map[string]func(t *testing.T){
		"configuring credentials as a path to a non-existent file results in an error": testAccProvider_credentialsInProviderBlock_badFilepathCausesError,
		"when credentials is set to an empty string there is an error":                 testAccProvider_credentialsInProviderBlock_configEmptyStringCausesError,
		"config takes precedence over environment variables":                           testAccProvider_credentialsInProviderBlock_configPrecedenceEnvironmentVariables,
		"assert precedence order of credentials environment variables":                 testAccProvider_credentialsInProviderBlock_precedenceOrderEnvironmentVariables,
	}

	for name, tc := range testCases {
		// shadow the tc variable into scope so that when
		// the loop continues, if t.Run hasn't executed tc(t)
		// yet, we don't have a race condition
		// see https://github.com/golang/go/wiki/CommonMistakes#using-goroutines-on-loop-iterator-variables
		tc := tc
		t.Run(name, func(t *testing.T) {
			tc(t)
		})
	}
}

func testAccProvider_credentialsInProviderBlock_configPrecedenceEnvironmentVariables(t *testing.T) {

	credentials := envvar.GetTestCredsFromEnv()

	// set all credentials env vars to 'bad' values
	badCreds := acctest.GenerateFakeCredentialsJson("test")
	for _, v := range envvar.CredsEnvVars {
		t.Setenv(v, badCreds)
	}

	context := map[string]interface{}{
		"credentials":   credentials,
		"resource_name": "tf-test-resource-" + acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		// No PreCheck for checking ENVs as we're altering them
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccProvider_credentialsInProviderBlock(context),
			},
		},
	})
}

func testAccProvider_credentialsInProviderBlock_configEmptyStringCausesError(t *testing.T) {

	context := map[string]interface{}{
		"credentials":   "",
		"resource_name": "tf-test-resource-" + acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config:      testAccProvider_credentialsInProviderBlock(context),
				ExpectError: regexp.MustCompile("expected a non-empty string"),
			},
		},
	})
}

func testAccProvider_credentialsInProviderBlock_badFilepathCausesError(t *testing.T) {

	pathToMissingFile := "./this/path/does/not/exist.json" // Doesn't exist

	context := map[string]interface{}{
		"credentials":   pathToMissingFile,
		"resource_name": "tf-test-resource-" + acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config:      testAccProvider_credentialsInProviderBlock(context),
				ExpectError: regexp.MustCompile("JSON credentials are not valid"),
			},
		},
	})
}

func testAccProvider_credentialsInProviderBlock_precedenceOrderEnvironmentVariables(t *testing.T) {

	/*
		These are all the ENVs for credentials, and they are in order of precedence.

		GOOGLE_CREDENTIALS
		GOOGLE_CLOUD_KEYFILE_JSON
		GCLOUD_KEYFILE_JSON
		GOOGLE_APPLICATION_CREDENTIALS
		GOOGLE_USE_DEFAULT_CREDENTIALS
	*/

	goodCredentials := envvar.GetTestCredsFromEnv()
	badCreds := acctest.GenerateFakeCredentialsJson("test")

	context := map[string]interface{}{
		"resource_name": "tf-test-resource-" + acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		// No PreCheck as we're altering ENVs
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				// Error as all ENVs set to 'bad' creds
				PreConfig: func() {
					for _, v := range envvar.CredsEnvVars {
						t.Setenv(v, badCreds)
					}
				},
				Config:      testAccProvider_credentials(context),
				ExpectError: regexp.MustCompile("private key should be a PEM or plain PKCS1 or PKCS8"),
			},
			{
				// GOOGLE_CREDENTIALS is used 1st if set
				PreConfig: func() {
					// good
					t.Setenv("GOOGLE_CREDENTIALS", goodCredentials) //used
					// bad
					t.Setenv("GOOGLE_CLOUD_KEYFILE_JSON", badCreds)
					t.Setenv("GCLOUD_KEYFILE_JSON", badCreds)
					t.Setenv("GOOGLE_APPLICATION_CREDENTIALS", badCreds)
				},
				Config: testAccProvider_credentials(context),
			},
			{
				// GOOGLE_CLOUD_KEYFILE_JSON is used 2nd
				PreConfig: func() {
					// unset
					t.Setenv("GOOGLE_CREDENTIALS", "")
					// good
					t.Setenv("GOOGLE_CLOUD_KEYFILE_JSON", goodCredentials) //used
					// bad
					t.Setenv("GCLOUD_KEYFILE_JSON", badCreds)
					t.Setenv("GOOGLE_APPLICATION_CREDENTIALS", badCreds)

				},
				Config: testAccProvider_credentials(context),
			},
			{
				// GOOGLE_CLOUD_KEYFILE_JSON is used 3rd
				PreConfig: func() {
					// unset
					t.Setenv("GOOGLE_CREDENTIALS", "")
					t.Setenv("GOOGLE_CLOUD_KEYFILE_JSON", "")
					// good
					t.Setenv("GCLOUD_KEYFILE_JSON", goodCredentials) //used
					// bad
					t.Setenv("GOOGLE_APPLICATION_CREDENTIALS", badCreds)
				},
				Config: testAccProvider_credentials(context),
			},
			{
				// GOOGLE_APPLICATION_CREDENTIALS is used 4th
				PreConfig: func() {
					// unset
					t.Setenv("GOOGLE_CREDENTIALS", "")
					t.Setenv("GOOGLE_CLOUD_KEYFILE_JSON", "")
					t.Setenv("GCLOUD_KEYFILE_JSON", "")
					// good
					t.Setenv("GOOGLE_APPLICATION_CREDENTIALS", goodCredentials) //used
				},
				Config: testAccProvider_credentials(context),
			},
		},
	})
}

func testAccProvider_credentials(context map[string]interface{}) string {
	return acctest.Nprintf(`
provider "google" {}

// SDK
resource "google_dns_managed_zone" "default" {
  name     = "%{resource_name}"
  dns_name = "dnssec.gcp.tfacc.hashicorptest.com."

  dnssec_config {
    state         = "off"
    non_existence = "nsec3"
  }
}

// Either SDK or PF depending on provider version
data "google_dns_managed_zone" "default" {
  name     = google_dns_managed_zone.default.name
}
`, context)
}

func testAccProvider_credentialsInProviderBlock(context map[string]interface{}) string {
	return acctest.Nprintf(`
provider "google" {
	credentials = "%{credentials}"
}

// SDK
resource "google_dns_managed_zone" "default" {
  name     = "%{resource_name}"
  dns_name = "dnssec.gcp.tfacc.hashicorptest.com."

  dnssec_config {
    state         = "off"
    non_existence = "nsec3"
  }
}

// Either SDK or PF depending on provider version
data "google_dns_managed_zone" "default" {
  name     = google_dns_managed_zone.default.name
}
`, context)
}
