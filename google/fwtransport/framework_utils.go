// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package fwtransport

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/hashicorp/terraform-provider-google/google/fwmodels"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

const uaEnvVar = "TF_APPEND_USER_AGENT"

func CompileUserAgentString(ctx context.Context, name, tfVersion, provVersion string) string {
	ua := fmt.Sprintf("Terraform/%s (+https://www.terraform.io) Terraform-Plugin-SDK/%s %s/%s", tfVersion, "terraform-plugin-framework", name, provVersion)

	if add := os.Getenv(uaEnvVar); add != "" {
		add = strings.TrimSpace(add)
		if len(add) > 0 {
			ua += " " + add
			tflog.Debug(ctx, fmt.Sprintf("Using modified User-Agent: %s", ua))
		}
	}

	return ua
}

// Remove fwtransport.GetCurrentUserEmailFramework
// Instead use transport.GetCurrentUserEmail

// GenerateFrameworkUserAgentString returns the provider-level USerAgent string with a module name appended, if provided in the config.
// Will only change the UserAgent if provider_meta is set in a module. See: https://developer.hashicorp.com/terraform/language/settings#passing-metadata-to-providers
// This new implementation is necessary over tpgresource.GenerateUserAgentString because it needs to accept the ProviderMeta values that the plugin-framework
// passes inside CreateRequest, ReadRequest, UpdateRequest, and DeleteRequest data.
func GenerateFrameworkUserAgentString(metaData *fwmodels.ProviderMetaModel, currUserAgent string) string {
	if metaData != nil && !metaData.ModuleName.IsNull() && metaData.ModuleName.ValueString() != "" {
		return strings.Join([]string{currUserAgent, metaData.ModuleName.ValueString()}, " ")
	}

	return currUserAgent
}

func HandleDatasourceNotFoundError(ctx context.Context, err error, state *tfsdk.State, resource string, diags *diag.Diagnostics) {
	if transport_tpg.IsGoogleApiErrorWithCode(err, 404) {
		tflog.Warn(ctx, fmt.Sprintf("Removing %s because it's gone", resource))
		// The resource doesn't exist anymore
		state.RemoveResource(ctx)
	}

	diags.AddError(fmt.Sprintf("Error when reading or editing %s", resource), err.Error())
}
