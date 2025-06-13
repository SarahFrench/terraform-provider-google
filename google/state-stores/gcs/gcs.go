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

// Would match existing config that can be set in a backend block
// I.e. https://developer.hashicorp.com/terraform/language/backend/gcs#configuration-variables
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

	// IMPLEMENTATION
	// The gcs backend's schema could be copied here
	//     https://github.com/hashicorp/terraform/blob/540512e27b881144c12bb165e6ecc76b430f009f/internal/backend/remote-state/gcs/backend.go#L44-L132
}

func (d *GcsStateStore) ValidateConfig(context.Context, storage.ValidateConfigRequest, *storage.ValidateConfigResponse) {
	// Validate config data and raise diagnostic errors and warnings when appropriate

	// All validation is expected to be offline.

	// IMPLEMENTATION
	// The gcs backend does not include a ValidateConfig/PrepareConfig method (tl;dr they're different names for the same thing).
	// Instead, the gcs backend embeds backendbase.Base:
	//     https://github.com/hashicorp/terraform/blob/540512e27b881144c12bb165e6ecc76b430f009f/internal/backend/remote-state/gcs/backend.go#L32
	// And backendbase.Base implements PrepareConfig:
	//     https://github.com/hashicorp/terraform/blob/540512e27b881144c12bb165e6ecc76b430f009f/internal/backend/backendbase/base.go#L57
	//
	// New validation logic would need to be implemented in this new ValidateConfig method.
	// This method would need to validated values and check ENVs.
	// Note: Core would check that the config matches the state store's schema.
}

func (d *GcsStateStore) Configure(ctx context.Context, req storage.ConfigureRequest, resp *storage.ConfigureResponse) {

	// Provider-level config information available to use when configuring the state store
	if req.ProviderData == nil {
		return
	}

	// Configure an instance of the state store in the provider server

	// IMPLEMENTATION
	// The gcs backend's Configure method would need to be re-implemented here
	//     https://github.com/hashicorp/terraform/blob/540512e27b881144c12bb165e6ecc76b430f009f/internal/backend/remote-state/gcs/backend.go#L135-L293
	// This logic would need to be adapted to account for what is currently in the gcs backend's implementation of a remote state manager
	//     https://github.com/hashicorp/terraform/blob/540512e27b881144c12bb165e6ecc76b430f009f/internal/backend/remote-state/gcs/client.go#L22
	//     Currently the Configure method assembles a remote state manager, but that abstraction won't exist in PSS

}

// Lock a specific state
func (d *GcsStateStore) Lock(ctx context.Context, req storage.LockRequest, resp *storage.LockResponse) {

	// Perform locking implementation to lock the state for the specific state/workspace
	// Return an identifier for that lock

	// IMPLEMENTATION
	// The Lock method on the remote state manager in the gcs backend would need to be re-implemented here
	//     https://github.com/hashicorp/terraform/blob/540512e27b881144c12bb165e6ecc76b430f009f/internal/backend/remote-state/gcs/client.go#L92-L120
	// This isn't a simple cut+paste, as PSS does not include the idea of state managers outside of Core.
}

// Unlock a specific state
func (d *GcsStateStore) Unlock(ctx context.Context, req storage.UnlockRequest, resp *storage.UnlockResponse) {

	// Perform unlocking implementation to unlock the state for the specific state/workspace
	// Requires knowledge about the lock id

	// IMPLEMENTATION
	// The Unlock method on the remote state manager in the gcs backend would need to be re-implemented here
	//     https://github.com/hashicorp/terraform/blob/540512e27b881144c12bb165e6ecc76b430f009f/internal/backend/remote-state/gcs/client.go#L122-L135
	// This isn't a simple cut+paste, as PSS does not include the idea of state managers outside of Core.

}

// Read a specific state from its location and return that value to core
func (d *GcsStateStore) Read(ctx context.Context, req storage.ReadStateRequest, resp *storage.ReadStateResponse) {

	// Read the file at the location that corresponds to the given state/environment
	// If missing, return empty state (what diags?)

	// IMPLEMENTATION
	// The Get method on the remote state manager in the gcs backend would need to be re-implemented here
	//     https://github.com/hashicorp/terraform/blob/540512e27b881144c12bb165e6ecc76b430f009f/internal/backend/remote-state/gcs/client.go#L32-L60
	// This isn't a simple cut+paste, as PSS does not include the idea of state managers outside of Core.

}

// Write a state sent from core to the specific state's location
func (d *GcsStateStore) Write(ctx context.Context, req storage.WriteStateRequest, resp *storage.WriteStateResponse) {

	// Create or overwrite the file at the location that corresponds to the given state/environment

	// IMPLEMENTATION
	// The Put method on the remote state manager in the gcs backend would need to be re-implemented here
	//     https://github.com/hashicorp/terraform/blob/540512e27b881144c12bb165e6ecc76b430f009f/internal/backend/remote-state/gcs/client.go#L62-L79
	// This isn't a simple cut+paste, as PSS does not include the idea of state managers outside of Core.

}

// List all states that can be managed by this state store
func (d *GcsStateStore) GetStates(ctx context.Context, req storage.StatesRequest, resp *storage.StatesResponse) {

	// Return list of all .tfstate files in the bucket at the configured prefix value (i.e location in the bucket)

	// IMPLEMENTATION
	// The logic in the gcs backend's Workspaces method would need to be re-implemented here
	//     https://github.com/hashicorp/terraform/blob/540512e27b881144c12bb165e6ecc76b430f009f/internal/backend/remote-state/gcs/backend_state.go#L29-L61

}

func (d *GcsStateStore) DeleteState(ctx context.Context, req storage.DeleteStateRequest, resp *storage.DeleteStateResponse) {

	// Delete the .tfstate file in the bucket at the configured prefix value (i.e location in the bucket) that
	// corresponds to the named state/workspace included in the request

	// IMPLEMENTATION
	// The logic in the gcs backend's DeleteWorkspace method would need to be re-implemented here
	//     https://github.com/hashicorp/terraform/blob/540512e27b881144c12bb165e6ecc76b430f009f/internal/backend/remote-state/gcs/backend_state.go#L64-L75
	// This would include pulling some logic out of the gcs backend's implementation of the remote client state manager
	//     https://github.com/hashicorp/terraform/blob/540512e27b881144c12bb165e6ecc76b430f009f/internal/backend/remote-state/gcs/client.go#L81
}
