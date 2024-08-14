// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package resourcemanager

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-google/google/fwmodels"
	"github.com/hashicorp/terraform-provider-google/google/fwresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

// Ensure the data source satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &GoogleClientConfigDataSource{}
	_ datasource.DataSourceWithConfigure = &GoogleClientConfigDataSource{}
	_ fwresource.LocationDescriber       = &GoogleClientConfigModel{}
)

func NewGoogleClientConfigDataSource() datasource.DataSource {
	return &GoogleClientConfigDataSource{}
}

type GoogleClientConfigDataSource struct {
	providerConfig *transport_tpg.Config
}

type GoogleClientConfigModel struct {
	Id types.String `tfsdk:"id"`
	// Below use string data type because we reuse SDK config logic to configure the PF provider
	// This propagates the SDK type system (only for provider-level config values)
	Project     string `tfsdk:"project"`
	Region      string `tfsdk:"region"`
	Zone        string `tfsdk:"zone"`
	AccessToken string `tfsdk:"access_token"`
}

func (m *GoogleClientConfigModel) GetLocationDescription(providerConfig *transport_tpg.Config) fwresource.LocationDescription {
	return fwresource.LocationDescription{
		RegionSchemaField: types.StringValue("region"),
		ZoneSchemaField:   types.StringValue("zone"),
		ProviderRegion:    providerConfig.Region,
		ProviderZone:      providerConfig.Zone,
	}
}

func (d *GoogleClientConfigDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_client_config"
}

func (d *GoogleClientConfigDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {

	resp.Schema = schema.Schema{

		Description:         "Use this data source to access the configuration of the Google Cloud provider.",
		MarkdownDescription: "Use this data source to access the configuration of the Google Cloud provider.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				Description:         "The ID of this data source in Terraform state. It is created in a projects/{{project}}/regions/{{region}}/zones/{{zone}} format and is NOT used by the data source in requests to Google APIs.",
				MarkdownDescription: "The ID of this data source in Terraform state. It is created in a projects/{{project}}/regions/{{region}}/zones/{{zone}} format and is NOT used by the data source in requests to Google APIs.",
			},
			"project": schema.StringAttribute{
				Description:         "The ID of the project to apply any resources to.",
				MarkdownDescription: "The ID of the project to apply any resources to.",
				Computed:            true,
			},
			"region": schema.StringAttribute{
				Description:         "The region to operate under.",
				MarkdownDescription: "The region to operate under.",
				Computed:            true,
			},
			"zone": schema.StringAttribute{
				Description:         "The zone to operate under.",
				MarkdownDescription: "The zone to operate under.",
				Computed:            true,
			},
			"access_token": schema.StringAttribute{
				Description:         "The OAuth2 access token used by the client to authenticate against the Google Cloud API.",
				MarkdownDescription: "The OAuth2 access token used by the client to authenticate against the Google Cloud API.",
				Computed:            true,
				Sensitive:           true,
			},
		},
	}
}

func (d *GoogleClientConfigDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	p, ok := req.ProviderData.(*transport_tpg.Config)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *transport_tpg.Config, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	// Required for accessing project, region, zone and tokenSource
	d.providerConfig = p
}

func (d *GoogleClientConfigDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data GoogleClientConfigModel
	var metaData *fwmodels.ProviderMetaModel
	var diags diag.Diagnostics

	// Read Provider meta into the meta model
	resp.Diagnostics.Append(req.ProviderMeta.Get(ctx, &metaData)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	locationInfo := data.GetLocationDescription(d.providerConfig)
	region, _ := locationInfo.GetRegion()
	zone, _ := locationInfo.GetZone()

	data.Id = types.StringValue(fmt.Sprintf("projects/%s/regions/%s/zones/%s", d.providerConfig.Project, region.String(), zone.String()))
	data.Project = d.providerConfig.Project
	data.Region = region.ValueString()
	data.Zone = zone.ValueString()

	token, err := d.providerConfig.TokenSource.Token()
	if err != nil {
		diags.AddError("Error setting access_token", err.Error())
		return
	}
	data.AccessToken = token.AccessToken

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
