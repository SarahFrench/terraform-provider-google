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
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/eventarc/GoogleChannelConfig.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package eventarc

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

func ResourceEventarcGoogleChannelConfig() *schema.Resource {
	return &schema.Resource{
		Create: resourceEventarcGoogleChannelConfigCreate,
		Read:   resourceEventarcGoogleChannelConfigRead,
		Update: resourceEventarcGoogleChannelConfigUpdate,
		Delete: resourceEventarcGoogleChannelConfigDelete,

		Importer: &schema.ResourceImporter{
			State: resourceEventarcGoogleChannelConfigImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		CustomizeDiff: customdiff.All(
			tpgresource.DefaultProviderProject,
		),

		Schema: map[string]*schema.Schema{
			"location": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The location for the resource`,
			},
			"name": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
				Description:      `Required. The resource name of the config. Must be in the format of, 'projects/{project}/locations/{location}/googleChannelConfig'.`,
			},
			"crypto_key_name": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
				Description:      `Optional. Resource name of a KMS crypto key (managed by the user) used to encrypt/decrypt their event data. It must match the pattern 'projects/*/locations/*/keyRings/*/cryptoKeys/*'.`,
			},
			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Output only. The last-modified time.`,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceEventarcGoogleChannelConfigCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	nameProp, err := expandEventarcGoogleChannelConfigName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	cryptoKeyNameProp, err := expandEventarcGoogleChannelConfigCryptoKeyName(d.Get("crypto_key_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("crypto_key_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(cryptoKeyNameProp)) && (ok || !reflect.DeepEqual(v, cryptoKeyNameProp)) {
		obj["cryptoKeyName"] = cryptoKeyNameProp
	}

	url, err := tpgresource.ReplaceVarsForId(d, config, "{{EventarcBasePath}}projects/{{project}}/locations/{{location}}/googleChannelConfig?updateMask=cryptoKeyName")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new GoogleChannelConfig: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for GoogleChannelConfig: %s", err)
	}
	billingProject = strings.TrimPrefix(project, "projects/")

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "PATCH",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutCreate),
		Headers:   headers,
	})
	if err != nil {
		return fmt.Errorf("Error creating GoogleChannelConfig: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVarsForId(d, config, "projects/{{project}}/locations/{{location}}/googleChannelConfig")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating GoogleChannelConfig %q: %#v", d.Id(), res)

	return resourceEventarcGoogleChannelConfigRead(d, meta)
}

func resourceEventarcGoogleChannelConfigRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVarsForId(d, config, "{{EventarcBasePath}}projects/{{project}}/locations/{{location}}/googleChannelConfig")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for GoogleChannelConfig: %s", err)
	}
	billingProject = strings.TrimPrefix(project, "projects/")

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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("EventarcGoogleChannelConfig %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading GoogleChannelConfig: %s", err)
	}

	if err := d.Set("name", flattenEventarcGoogleChannelConfigName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading GoogleChannelConfig: %s", err)
	}
	if err := d.Set("update_time", flattenEventarcGoogleChannelConfigUpdateTime(res["updateTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading GoogleChannelConfig: %s", err)
	}
	if err := d.Set("crypto_key_name", flattenEventarcGoogleChannelConfigCryptoKeyName(res["cryptoKeyName"], d, config)); err != nil {
		return fmt.Errorf("Error reading GoogleChannelConfig: %s", err)
	}

	return nil
}

func resourceEventarcGoogleChannelConfigUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for GoogleChannelConfig: %s", err)
	}
	billingProject = strings.TrimPrefix(project, "projects/")

	obj := make(map[string]interface{})
	cryptoKeyNameProp, err := expandEventarcGoogleChannelConfigCryptoKeyName(d.Get("crypto_key_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("crypto_key_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, cryptoKeyNameProp)) {
		obj["cryptoKeyName"] = cryptoKeyNameProp
	}

	url, err := tpgresource.ReplaceVarsForId(d, config, "{{EventarcBasePath}}projects/{{project}}/locations/{{location}}/googleChannelConfig")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating GoogleChannelConfig %q: %#v", d.Id(), obj)
	headers := make(http.Header)
	updateMask := []string{}

	if d.HasChange("crypto_key_name") {
		updateMask = append(updateMask, "cryptoKeyName")
	}
	// updateMask is a URL parameter but not present in the schema, so ReplaceVars
	// won't set it
	url, err = transport_tpg.AddQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	// if updateMask is empty we are not updating anything so skip the post
	if len(updateMask) > 0 {
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
			return fmt.Errorf("Error updating GoogleChannelConfig %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating GoogleChannelConfig %q: %#v", d.Id(), res)
		}

	}

	return resourceEventarcGoogleChannelConfigRead(d, meta)
}

func resourceEventarcGoogleChannelConfigDelete(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[WARNING] Eventarc GoogleChannelConfig resources"+
		" cannot be deleted from Google Cloud. The resource %s will be removed from Terraform"+
		" state, but will still be present on Google Cloud.", d.Id())
	d.SetId("")

	return nil
}

func resourceEventarcGoogleChannelConfigImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/googleChannelConfig$",
		"^(?P<project>[^/]+)/(?P<location>[^/]+)$",
		"^(?P<location>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVarsForId(d, config, "projects/{{project}}/locations/{{location}}/googleChannelConfig")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenEventarcGoogleChannelConfigName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenEventarcGoogleChannelConfigUpdateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenEventarcGoogleChannelConfigCryptoKeyName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandEventarcGoogleChannelConfigName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandEventarcGoogleChannelConfigCryptoKeyName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
