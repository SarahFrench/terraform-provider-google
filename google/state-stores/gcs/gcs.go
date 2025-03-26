// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package gcs

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/storage"
)

var _ storage.StateStore = GcsStateStore{}

func NewGcsStateStore() storage.StateStore {
	return &GcsStateStore{
		name: "gcs",
	}
}

type GcsStateStore struct {
	name string
}

type GcsStateStoreModel struct {
	Bucket           types.String `tfsdk:"bucket"`
	Prefix           types.String `tfsdk:"prefix"`
	EncryptionKey    types.String `tfsdk:"encryption_key"`
	KmsEncryptionKey types.String `tfsdk:"kms_encryption_key"`

	// Also in provider model
	AccessToken                        types.String `tfsdk:"access_token"`
	Credentials                        types.String `tfsdk:"credentials"`
	ImpersonateServiceAccount          types.String `tfsdk:"impersonate_service_account"`
	ImpersonateServiceAccountDelegates types.List   `tfsdk:"impersonate_service_account_delegates"`
	StorageCustomEndpoint              types.String `tfsdk:"storage_custom_endpoint"`
}

func (f GcsStateStore) Metadata(ctx context.Context, req storage.MetadataRequest, resp *storage.MetadataResponse) {
	resp.Name = f.name
}

func (d *GcsStateStore) Schema(ctx context.Context, req storage.SchemaRequest, resp *storage.SchemaResponse) {

	// Defines and returns the schema of the state store

	// Schema would match what's currently configurable though a backend block today
}

func (d *GcsStateStore) Configure(ctx context.Context, req storage.ConfigureRequest, resp *storage.ConfigureResponse) {

	// Provider-level config information available to use when configuring the state store
	if req.ProviderData == nil {
		return
	}

	// Configure an instance of the state store in the provider server
}

// Lock a specific state
func (d *GcsStateStore) Lock(ctx context.Context, req storage.LockRequest, resp *storage.LockResponse) {

	// Perform locking implementation to lock the state for the specific state/workspace
	// Return an identifier for that lock

}

// Unlock a specific state
func (d *GcsStateStore) Unlock(ctx context.Context, req storage.UnlockRequest, resp *storage.UnlockResponse) {

	// Perform unlocking implementation to unlock the state for the specific state/workspace
	// Requires knowledge about the lock id

}

// Read a specific state from its location and return that value to core
func (d *GcsStateStore) ReadState(ctx context.Context, req storage.ReadStateRequest, resp *storage.ReadStateResponse) {

	// Read the file at the location that corresponds to the given state/environment
	// If missing, return empty state (what diags?)

}

// Write a state sent from core to the specific state's location
func (d *GcsStateStore) WriteState(ctx context.Context, req storage.WriteStateRequest, resp *storage.WriteStateResponse) {

	// Create or overwrite the file at the location that corresponds to the given state/environment

}

// List all states that can be managed by this state store
func (d *GcsStateStore) States(ctx context.Context, req storage.StatesRequest, resp *storage.StatesResponse) {

	// Return list of all .tfstate files in the bucket at the configured prefix value (i.e location in the bucket)

}

func (d *GcsStateStore) DeleteState(ctx context.Context, req storage.DeleteStateRequest, resp *storage.DeleteStateResponse) {

	// Delete the .tfstate file in the bucket at the configured prefix value (i.e location in the bucket) that
	// corresponds to the named state/workspace included in the request

}
