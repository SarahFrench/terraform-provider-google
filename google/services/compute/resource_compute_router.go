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
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/compute/Router.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package compute

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
	"github.com/hashicorp/terraform-provider-google/google/verify"
)

// customizeDiff func for additional checks on google_compute_router properties:
func resourceComputeRouterCustomDiff(_ context.Context, diff *schema.ResourceDiff, meta interface{}) error {

	block := diff.Get("bgp.0").(map[string]interface{})
	advertiseMode := block["advertise_mode"]
	advertisedGroups := block["advertised_groups"].([]interface{})
	advertisedIPRanges := block["advertised_ip_ranges"].([]interface{})

	if advertiseMode == "DEFAULT" && len(advertisedGroups) != 0 {
		return fmt.Errorf("Error in bgp: advertised_groups cannot be specified when using advertise_mode DEFAULT")
	}
	if advertiseMode == "DEFAULT" && len(advertisedIPRanges) != 0 {
		return fmt.Errorf("Error in bgp: advertised_ip_ranges cannot be specified when using advertise_mode DEFAULT")
	}

	return nil
}

func ResourceComputeRouter() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeRouterCreate,
		Read:   resourceComputeRouterRead,
		Update: resourceComputeRouterUpdate,
		Delete: resourceComputeRouterDelete,

		Importer: &schema.ResourceImporter{
			State: resourceComputeRouterImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		CustomizeDiff: customdiff.All(
			resourceComputeRouterCustomDiff,
			tpgresource.DefaultProviderProject,
		),

		Schema: map[string]*schema.Schema{
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: verify.ValidateGCEName,
				Description: `Name of the resource. The name must be 1-63 characters long, and
comply with RFC1035. Specifically, the name must be 1-63 characters
long and match the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?'
which means the first character must be a lowercase letter, and all
following characters must be a dash, lowercase letter, or digit,
except the last character, which cannot be a dash.`,
			},
			"network": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
				Description:      `A reference to the network to which this router belongs.`,
			},
			"bgp": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: `BGP information specific to this router.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"asn": {
							Type:         schema.TypeInt,
							Required:     true,
							ValidateFunc: verify.ValidateRFC6996Asn,
							Description: `Local BGP Autonomous System Number (ASN). Must be an RFC6996
private ASN, either 16-bit or 32-bit. The value will be fixed for
this router resource. All VPN tunnels that link to this router
will have the same local ASN.`,
						},
						"advertise_mode": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: verify.ValidateEnum([]string{"DEFAULT", "CUSTOM", ""}),
							Description:  `User-specified flag to indicate which mode to use for advertisement. Default value: "DEFAULT" Possible values: ["DEFAULT", "CUSTOM"]`,
							Default:      "DEFAULT",
						},
						"advertised_groups": {
							Type:     schema.TypeList,
							Optional: true,
							Description: `User-specified list of prefix groups to advertise in custom mode.
This field can only be populated if advertiseMode is CUSTOM and
is advertised to all peers of the router. These groups will be
advertised in addition to any specified prefixes. Leave this field
blank to advertise no custom groups.

This enum field has the one valid value: ALL_SUBNETS`,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"advertised_ip_ranges": {
							Type:     schema.TypeList,
							Optional: true,
							Description: `User-specified list of individual IP ranges to advertise in
custom mode. This field can only be populated if advertiseMode
is CUSTOM and is advertised to all peers of the router. These IP
ranges will be advertised in addition to any specified groups.
Leave this field blank to advertise no custom IP ranges.`,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"range": {
										Type:     schema.TypeString,
										Required: true,
										Description: `The IP range to advertise. The value must be a
CIDR-formatted string.`,
									},
									"description": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: `User-specified description for the IP range.`,
									},
								},
							},
						},
						"identifier_range": {
							Type:     schema.TypeString,
							Computed: true,
							Optional: true,
							Description: `Explicitly specifies a range of valid BGP Identifiers for this Router.
It is provided as a link-local IPv4 range (from 169.254.0.0/16), of
size at least /30, even if the BGP sessions are over IPv6. It must
not overlap with any IPv4 BGP session ranges. Other vendors commonly
call this router ID.`,
						},
						"keepalive_interval": {
							Type:     schema.TypeInt,
							Optional: true,
							Description: `The interval in seconds between BGP keepalive messages that are sent
to the peer. Hold time is three times the interval at which keepalive
messages are sent, and the hold time is the maximum number of seconds
allowed to elapse between successive keepalive messages that BGP
receives from a peer.

BGP will use the smaller of either the local hold time value or the
peer's hold time value as the hold time for the BGP connection
between the two peers. If set, this value must be between 20 and 60.
The default is 20.`,
							Default: 20,
						},
					},
				},
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `An optional description of this resource.`,
			},
			"encrypted_interconnect_router": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
				Description: `Indicates if a router is dedicated for use with encrypted VLAN
attachments (interconnectAttachments).`,
			},
			"region": {
				Type:             schema.TypeString,
				Computed:         true,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
				Description:      `Region where the router resides.`,
			},
			"creation_timestamp": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Creation timestamp in RFC3339 text format.`,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"self_link": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceComputeRouterCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	nameProp, err := expandComputeRouterName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	descriptionProp, err := expandComputeRouterDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); ok || !reflect.DeepEqual(v, descriptionProp) {
		obj["description"] = descriptionProp
	}
	networkProp, err := expandComputeRouterNetwork(d.Get("network"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("network"); !tpgresource.IsEmptyValue(reflect.ValueOf(networkProp)) && (ok || !reflect.DeepEqual(v, networkProp)) {
		obj["network"] = networkProp
	}
	bgpProp, err := expandComputeRouterBgp(d.Get("bgp"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("bgp"); ok || !reflect.DeepEqual(v, bgpProp) {
		obj["bgp"] = bgpProp
	}
	encryptedInterconnectRouterProp, err := expandComputeRouterEncryptedInterconnectRouter(d.Get("encrypted_interconnect_router"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("encrypted_interconnect_router"); !tpgresource.IsEmptyValue(reflect.ValueOf(encryptedInterconnectRouterProp)) && (ok || !reflect.DeepEqual(v, encryptedInterconnectRouterProp)) {
		obj["encryptedInterconnectRouter"] = encryptedInterconnectRouterProp
	}
	regionProp, err := expandComputeRouterRegion(d.Get("region"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("region"); !tpgresource.IsEmptyValue(reflect.ValueOf(regionProp)) && (ok || !reflect.DeepEqual(v, regionProp)) {
		obj["region"] = regionProp
	}

	lockName, err := tpgresource.ReplaceVars(d, config, "router/{{region}}/{{name}}")
	if err != nil {
		return err
	}
	transport_tpg.MutexStore.Lock(lockName)
	defer transport_tpg.MutexStore.Unlock(lockName)

	url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/routers")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Router: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Router: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "POST",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutCreate),
		Headers:   headers,
	})
	if err != nil {
		return fmt.Errorf("Error creating Router: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/regions/{{region}}/routers/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = ComputeOperationWaitTime(
		config, res, project, "Creating Router", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create Router: %s", err)
	}

	log.Printf("[DEBUG] Finished creating Router %q: %#v", d.Id(), res)

	return resourceComputeRouterRead(d, meta)
}

func resourceComputeRouterRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/routers/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Router: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "GET",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Headers:   headers,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("ComputeRouter %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Router: %s", err)
	}

	if err := d.Set("creation_timestamp", flattenComputeRouterCreationTimestamp(res["creationTimestamp"], d, config)); err != nil {
		return fmt.Errorf("Error reading Router: %s", err)
	}
	if err := d.Set("name", flattenComputeRouterName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading Router: %s", err)
	}
	if err := d.Set("description", flattenComputeRouterDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading Router: %s", err)
	}
	if err := d.Set("network", flattenComputeRouterNetwork(res["network"], d, config)); err != nil {
		return fmt.Errorf("Error reading Router: %s", err)
	}
	if err := d.Set("bgp", flattenComputeRouterBgp(res["bgp"], d, config)); err != nil {
		return fmt.Errorf("Error reading Router: %s", err)
	}
	if err := d.Set("encrypted_interconnect_router", flattenComputeRouterEncryptedInterconnectRouter(res["encryptedInterconnectRouter"], d, config)); err != nil {
		return fmt.Errorf("Error reading Router: %s", err)
	}
	if err := d.Set("region", flattenComputeRouterRegion(res["region"], d, config)); err != nil {
		return fmt.Errorf("Error reading Router: %s", err)
	}
	if err := d.Set("self_link", tpgresource.ConvertSelfLinkToV1(res["selfLink"].(string))); err != nil {
		return fmt.Errorf("Error reading Router: %s", err)
	}

	return nil
}

func resourceComputeRouterUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Router: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	descriptionProp, err := expandComputeRouterDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); ok || !reflect.DeepEqual(v, descriptionProp) {
		obj["description"] = descriptionProp
	}
	bgpProp, err := expandComputeRouterBgp(d.Get("bgp"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("bgp"); ok || !reflect.DeepEqual(v, bgpProp) {
		obj["bgp"] = bgpProp
	}

	lockName, err := tpgresource.ReplaceVars(d, config, "router/{{region}}/{{name}}")
	if err != nil {
		return err
	}
	transport_tpg.MutexStore.Lock(lockName)
	defer transport_tpg.MutexStore.Unlock(lockName)

	url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/routers/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Router %q: %#v", d.Id(), obj)
	headers := make(http.Header)

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "PATCH",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutUpdate),
		Headers:   headers,
	})

	if err != nil {
		return fmt.Errorf("Error updating Router %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating Router %q: %#v", d.Id(), res)
	}

	err = ComputeOperationWaitTime(
		config, res, project, "Updating Router", userAgent,
		d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return err
	}

	return resourceComputeRouterRead(d, meta)
}

func resourceComputeRouterDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Router: %s", err)
	}
	billingProject = project

	lockName, err := tpgresource.ReplaceVars(d, config, "router/{{region}}/{{name}}")
	if err != nil {
		return err
	}
	transport_tpg.MutexStore.Lock(lockName)
	defer transport_tpg.MutexStore.Unlock(lockName)

	url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/routers/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)

	log.Printf("[DEBUG] Deleting Router %q", d.Id())
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "DELETE",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutDelete),
		Headers:   headers,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "Router")
	}

	err = ComputeOperationWaitTime(
		config, res, project, "Deleting Router", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting Router %q: %#v", d.Id(), res)
	return nil
}

func resourceComputeRouterImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^projects/(?P<project>[^/]+)/regions/(?P<region>[^/]+)/routers/(?P<name>[^/]+)$",
		"^(?P<project>[^/]+)/(?P<region>[^/]+)/(?P<name>[^/]+)$",
		"^(?P<region>[^/]+)/(?P<name>[^/]+)$",
		"^(?P<name>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/regions/{{region}}/routers/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenComputeRouterCreationTimestamp(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeRouterName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeRouterDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeRouterNetwork(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return tpgresource.ConvertSelfLinkToV1(v.(string))
}

func flattenComputeRouterBgp(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["asn"] =
		flattenComputeRouterBgpAsn(original["asn"], d, config)
	transformed["advertise_mode"] =
		flattenComputeRouterBgpAdvertiseMode(original["advertiseMode"], d, config)
	transformed["advertised_groups"] =
		flattenComputeRouterBgpAdvertisedGroups(original["advertisedGroups"], d, config)
	transformed["advertised_ip_ranges"] =
		flattenComputeRouterBgpAdvertisedIpRanges(original["advertisedIpRanges"], d, config)
	transformed["keepalive_interval"] =
		flattenComputeRouterBgpKeepaliveInterval(original["keepaliveInterval"], d, config)
	transformed["identifier_range"] =
		flattenComputeRouterBgpIdentifierRange(original["identifierRange"], d, config)
	return []interface{}{transformed}
}
func flattenComputeRouterBgpAsn(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := tpgresource.StringToFixed64(strVal); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func flattenComputeRouterBgpAdvertiseMode(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeRouterBgpAdvertisedGroups(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeRouterBgpAdvertisedIpRanges(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	apiData := make([]map[string]interface{}, 0, len(l))
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		apiData = append(apiData, map[string]interface{}{
			"description": original["description"],
			"range":       original["range"],
		})
	}
	configData := []map[string]interface{}{}
	if v, ok := d.GetOk("bgp.0.advertised_ip_ranges"); ok {
		for _, item := range v.([]interface{}) {
			configData = append(configData, item.(map[string]interface{}))
		}
	}
	sorted, err := tpgresource.SortMapsByConfigOrder(configData, apiData, "range")
	if err != nil {
		log.Printf("[ERROR] Could not support API response for advertisedIpRanges.0.range: %s", err)
		return apiData
	}
	return sorted
}

func flattenComputeRouterBgpKeepaliveInterval(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := tpgresource.StringToFixed64(strVal); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func flattenComputeRouterBgpIdentifierRange(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeRouterEncryptedInterconnectRouter(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeRouterRegion(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return tpgresource.NameFromSelfLinkStateFunc(v)
}

func expandComputeRouterName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRouterDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRouterNetwork(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	f, err := tpgresource.ParseGlobalFieldValue("networks", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for network: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputeRouterBgp(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedAsn, err := expandComputeRouterBgpAsn(original["asn"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAsn); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["asn"] = transformedAsn
	}

	transformedAdvertiseMode, err := expandComputeRouterBgpAdvertiseMode(original["advertise_mode"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAdvertiseMode); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["advertiseMode"] = transformedAdvertiseMode
	}

	transformedAdvertisedGroups, err := expandComputeRouterBgpAdvertisedGroups(original["advertised_groups"], d, config)
	if err != nil {
		return nil, err
	} else {
		transformed["advertisedGroups"] = transformedAdvertisedGroups
	}

	transformedAdvertisedIpRanges, err := expandComputeRouterBgpAdvertisedIpRanges(original["advertised_ip_ranges"], d, config)
	if err != nil {
		return nil, err
	} else {
		transformed["advertisedIpRanges"] = transformedAdvertisedIpRanges
	}

	transformedKeepaliveInterval, err := expandComputeRouterBgpKeepaliveInterval(original["keepalive_interval"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedKeepaliveInterval); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["keepaliveInterval"] = transformedKeepaliveInterval
	}

	transformedIdentifierRange, err := expandComputeRouterBgpIdentifierRange(original["identifier_range"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedIdentifierRange); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["identifierRange"] = transformedIdentifierRange
	}

	return transformed, nil
}

func expandComputeRouterBgpAsn(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRouterBgpAdvertiseMode(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRouterBgpAdvertisedGroups(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRouterBgpAdvertisedIpRanges(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedRange, err := expandComputeRouterBgpAdvertisedIpRangesRange(original["range"], d, config)
		if err != nil {
			return nil, err
		} else {
			transformed["range"] = transformedRange
		}

		transformedDescription, err := expandComputeRouterBgpAdvertisedIpRangesDescription(original["description"], d, config)
		if err != nil {
			return nil, err
		} else {
			transformed["description"] = transformedDescription
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandComputeRouterBgpAdvertisedIpRangesRange(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRouterBgpAdvertisedIpRangesDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRouterBgpKeepaliveInterval(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRouterBgpIdentifierRange(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRouterEncryptedInterconnectRouter(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRouterRegion(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	f, err := tpgresource.ParseGlobalFieldValue("regions", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for region: %s", err)
	}
	return f.RelativeLink(), nil
}
