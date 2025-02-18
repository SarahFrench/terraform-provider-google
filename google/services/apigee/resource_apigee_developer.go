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
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/apigee/Developer.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package apigee

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

func ResourceApigeeDeveloper() *schema.Resource {
	return &schema.Resource{
		Create: resourceApigeeDeveloperCreate,
		Read:   resourceApigeeDeveloperRead,
		Update: resourceApigeeDeveloperUpdate,
		Delete: resourceApigeeDeveloperDelete,

		Importer: &schema.ResourceImporter{
			State: resourceApigeeDeveloperImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"email": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Email address of the developer. This value is used to uniquely identify the developer in Apigee hybrid. Note that the email address has to be in lowercase only..`,
			},
			"first_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `First name of the developer.`,
			},
			"last_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `Last name of the developer.`,
			},
			"org_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `The Apigee Organization associated with the Apigee instance,
in the format 'organizations/{{org_name}}'.`,
			},
			"user_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `User name of the developer. Not used by Apigee hybrid.`,
			},
			"attributes": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: `Developer attributes (name/value pairs). The custom attribute limit is 18.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: `Key of the attribute`,
						},
						"value": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: `Value of the attribute`,
						},
					},
				},
			},
			"created_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Time at which the developer was created in milliseconds since epoch.`,
			},
			"last_modified_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Time at which the developer was last modified in milliseconds since epoch.`,
			},
			"organizatio_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Name of the Apigee organization in which the developer resides.`,
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Status of the developer. Valid values are active and inactive.`,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceApigeeDeveloperCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	emailProp, err := expandApigeeDeveloperEmail(d.Get("email"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("email"); !tpgresource.IsEmptyValue(reflect.ValueOf(emailProp)) && (ok || !reflect.DeepEqual(v, emailProp)) {
		obj["email"] = emailProp
	}
	firstNameProp, err := expandApigeeDeveloperFirstName(d.Get("first_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("first_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(firstNameProp)) && (ok || !reflect.DeepEqual(v, firstNameProp)) {
		obj["firstName"] = firstNameProp
	}
	lastNameProp, err := expandApigeeDeveloperLastName(d.Get("last_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("last_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(lastNameProp)) && (ok || !reflect.DeepEqual(v, lastNameProp)) {
		obj["lastName"] = lastNameProp
	}
	userNameProp, err := expandApigeeDeveloperUserName(d.Get("user_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("user_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(userNameProp)) && (ok || !reflect.DeepEqual(v, userNameProp)) {
		obj["userName"] = userNameProp
	}
	attributesProp, err := expandApigeeDeveloperAttributes(d.Get("attributes"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("attributes"); !tpgresource.IsEmptyValue(reflect.ValueOf(attributesProp)) && (ok || !reflect.DeepEqual(v, attributesProp)) {
		obj["attributes"] = attributesProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{ApigeeBasePath}}{{org_id}}/developers")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Developer: %#v", obj)
	billingProject := ""

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
		return fmt.Errorf("Error creating Developer: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "{{org_id}}/developers/{{email}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating Developer %q: %#v", d.Id(), res)

	return resourceApigeeDeveloperRead(d, meta)
}

func resourceApigeeDeveloperRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{ApigeeBasePath}}{{org_id}}/developers/{{email}}")
	if err != nil {
		return err
	}

	billingProject := ""

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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("ApigeeDeveloper %q", d.Id()))
	}

	if err := d.Set("email", flattenApigeeDeveloperEmail(res["email"], d, config)); err != nil {
		return fmt.Errorf("Error reading Developer: %s", err)
	}
	if err := d.Set("first_name", flattenApigeeDeveloperFirstName(res["firstName"], d, config)); err != nil {
		return fmt.Errorf("Error reading Developer: %s", err)
	}
	if err := d.Set("last_name", flattenApigeeDeveloperLastName(res["lastName"], d, config)); err != nil {
		return fmt.Errorf("Error reading Developer: %s", err)
	}
	if err := d.Set("user_name", flattenApigeeDeveloperUserName(res["userName"], d, config)); err != nil {
		return fmt.Errorf("Error reading Developer: %s", err)
	}
	if err := d.Set("attributes", flattenApigeeDeveloperAttributes(res["attributes"], d, config)); err != nil {
		return fmt.Errorf("Error reading Developer: %s", err)
	}
	if err := d.Set("organizatio_name", flattenApigeeDeveloperOrganizatioName(res["organizatioName"], d, config)); err != nil {
		return fmt.Errorf("Error reading Developer: %s", err)
	}
	if err := d.Set("status", flattenApigeeDeveloperStatus(res["status"], d, config)); err != nil {
		return fmt.Errorf("Error reading Developer: %s", err)
	}
	if err := d.Set("created_at", flattenApigeeDeveloperCreatedAt(res["createdAt"], d, config)); err != nil {
		return fmt.Errorf("Error reading Developer: %s", err)
	}
	if err := d.Set("last_modified_at", flattenApigeeDeveloperLastModifiedAt(res["lastModifiedAt"], d, config)); err != nil {
		return fmt.Errorf("Error reading Developer: %s", err)
	}

	return nil
}

func resourceApigeeDeveloperUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	obj := make(map[string]interface{})
	emailProp, err := expandApigeeDeveloperEmail(d.Get("email"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("email"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, emailProp)) {
		obj["email"] = emailProp
	}
	firstNameProp, err := expandApigeeDeveloperFirstName(d.Get("first_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("first_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, firstNameProp)) {
		obj["firstName"] = firstNameProp
	}
	lastNameProp, err := expandApigeeDeveloperLastName(d.Get("last_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("last_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, lastNameProp)) {
		obj["lastName"] = lastNameProp
	}
	userNameProp, err := expandApigeeDeveloperUserName(d.Get("user_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("user_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, userNameProp)) {
		obj["userName"] = userNameProp
	}
	attributesProp, err := expandApigeeDeveloperAttributes(d.Get("attributes"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("attributes"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, attributesProp)) {
		obj["attributes"] = attributesProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{ApigeeBasePath}}{{org_id}}/developers/{{email}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Developer %q: %#v", d.Id(), obj)
	headers := make(http.Header)

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "PUT",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutUpdate),
		Headers:   headers,
	})

	if err != nil {
		return fmt.Errorf("Error updating Developer %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating Developer %q: %#v", d.Id(), res)
	}

	return resourceApigeeDeveloperRead(d, meta)
}

func resourceApigeeDeveloperDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	url, err := tpgresource.ReplaceVars(d, config, "{{ApigeeBasePath}}{{org_id}}/developers/{{email}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)

	log.Printf("[DEBUG] Deleting Developer %q", d.Id())
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
		return transport_tpg.HandleNotFoundError(err, d, "Developer")
	}

	log.Printf("[DEBUG] Finished deleting Developer %q: %#v", d.Id(), res)
	return nil
}

func resourceApigeeDeveloperImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)

	// current import_formats cannot import fields with forward slashes in their value
	if err := tpgresource.ParseImportId([]string{"(?P<email>.+)"}, d, config); err != nil {
		return nil, err
	}

	nameParts := strings.Split(d.Get("email").(string), "/")
	if len(nameParts) == 4 {
		// `organizations/{{org_name}}/developers/{{email}}`
		orgId := fmt.Sprintf("organizations/%s", nameParts[1])
		if err := d.Set("org_id", orgId); err != nil {
			return nil, fmt.Errorf("Error setting org_id: %s", err)
		}
		if err := d.Set("email", nameParts[3]); err != nil {
			return nil, fmt.Errorf("Error setting email: %s", err)
		}
	} else if len(nameParts) == 3 {
		// `organizations/{{org_name}}/{{email}}`
		orgId := fmt.Sprintf("organizations/%s", nameParts[1])
		if err := d.Set("org_id", orgId); err != nil {
			return nil, fmt.Errorf("Error setting org_id: %s", err)
		}
		if err := d.Set("email", nameParts[2]); err != nil {
			return nil, fmt.Errorf("Error setting name: %s", err)
		}
	} else {
		return nil, fmt.Errorf(
			"Saw %s when the name is expected to have shape %s or %s",
			d.Get("name"),
			"organizations/{{org_name}}/developers/{{email}}",
			"organizations/{{org_name}}/{{email}}")
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "{{org_id}}/developers/{{email}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenApigeeDeveloperEmail(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenApigeeDeveloperFirstName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenApigeeDeveloperLastName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenApigeeDeveloperUserName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenApigeeDeveloperAttributes(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := make([]interface{}, 0, len(l))
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed = append(transformed, map[string]interface{}{
			"name":  flattenApigeeDeveloperAttributesName(original["name"], d, config),
			"value": flattenApigeeDeveloperAttributesValue(original["value"], d, config),
		})
	}
	return transformed
}
func flattenApigeeDeveloperAttributesName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenApigeeDeveloperAttributesValue(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenApigeeDeveloperOrganizatioName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenApigeeDeveloperStatus(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenApigeeDeveloperCreatedAt(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenApigeeDeveloperLastModifiedAt(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandApigeeDeveloperEmail(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApigeeDeveloperFirstName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApigeeDeveloperLastName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApigeeDeveloperUserName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApigeeDeveloperAttributes(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedName, err := expandApigeeDeveloperAttributesName(original["name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["name"] = transformedName
		}

		transformedValue, err := expandApigeeDeveloperAttributesValue(original["value"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedValue); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["value"] = transformedValue
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandApigeeDeveloperAttributesName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApigeeDeveloperAttributesValue(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
