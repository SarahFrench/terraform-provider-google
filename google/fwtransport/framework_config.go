// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package fwtransport

import (
	"context"
	"net/http"
	"time"

	"golang.org/x/oauth2"

	"google.golang.org/api/option"

	"github.com/hashicorp/terraform-plugin-framework/types"

	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

// No longer used
type FrameworkProviderConfig struct {
	// DCLConfig
	// AccessToken
	// Credentials
	// ImpersonateServiceAccount
	// ImpersonateServiceAccountDelegates
	Project        types.String
	Region         types.String
	BillingProject types.String
	Zone           types.String
	UniverseDomain types.String
	Scopes         types.List
	// BatchingConfig
	UserProjectOverride types.Bool
	// RequestReason
	// RequestTimeout
	// DefaultLabels
	// AddTerraformAttributionLabel
	// TerraformAttributionLabelAdditionStrategy
	PollInterval time.Duration

	Client             *http.Client
	Context            context.Context
	UserAgent          string
	gRPCLoggingOptions []option.ClientOption

	TokenSource oauth2.TokenSource

	RequestBatcherIam          *transport_tpg.RequestBatcher
	RequestBatcherServiceUsage *transport_tpg.RequestBatcher

	// paths for client setup
	AccessApprovalBasePath           string
	AccessContextManagerBasePath     string
	ActiveDirectoryBasePath          string
	AlloydbBasePath                  string
	ApigeeBasePath                   string
	AppEngineBasePath                string
	ApphubBasePath                   string
	ArtifactRegistryBasePath         string
	BeyondcorpBasePath               string
	BiglakeBasePath                  string
	BigQueryBasePath                 string
	BigqueryAnalyticsHubBasePath     string
	BigqueryConnectionBasePath       string
	BigqueryDatapolicyBasePath       string
	BigqueryDataTransferBasePath     string
	BigqueryReservationBasePath      string
	BigtableBasePath                 string
	BillingBasePath                  string
	BinaryAuthorizationBasePath      string
	BlockchainNodeEngineBasePath     string
	CertificateManagerBasePath       string
	CloudAssetBasePath               string
	CloudBuildBasePath               string
	Cloudbuildv2BasePath             string
	ClouddeployBasePath              string
	ClouddomainsBasePath             string
	CloudFunctionsBasePath           string
	Cloudfunctions2BasePath          string
	CloudIdentityBasePath            string
	CloudIdsBasePath                 string
	CloudQuotasBasePath              string
	CloudRunBasePath                 string
	CloudRunV2BasePath               string
	CloudSchedulerBasePath           string
	CloudTasksBasePath               string
	ComposerBasePath                 string
	ComputeBasePath                  string
	ContainerAnalysisBasePath        string
	ContainerAttachedBasePath        string
	CoreBillingBasePath              string
	DatabaseMigrationServiceBasePath string
	DataCatalogBasePath              string
	DataFusionBasePath               string
	DataLossPreventionBasePath       string
	DataPipelineBasePath             string
	DataplexBasePath                 string
	DataprocBasePath                 string
	DataprocMetastoreBasePath        string
	DatastoreBasePath                string
	DatastreamBasePath               string
	DeploymentManagerBasePath        string
	DialogflowBasePath               string
	DialogflowCXBasePath             string
	DiscoveryEngineBasePath          string
	DNSBasePath                      string
	DocumentAIBasePath               string
	DocumentAIWarehouseBasePath      string
	EdgecontainerBasePath            string
	EdgenetworkBasePath              string
	EssentialContactsBasePath        string
	FilestoreBasePath                string
	FirebaseAppCheckBasePath         string
	FirestoreBasePath                string
	GKEBackupBasePath                string
	GKEHubBasePath                   string
	GKEHub2BasePath                  string
	GkeonpremBasePath                string
	HealthcareBasePath               string
	IAM2BasePath                     string
	IAMBetaBasePath                  string
	IAMWorkforcePoolBasePath         string
	IapBasePath                      string
	IdentityPlatformBasePath         string
	IntegrationConnectorsBasePath    string
	IntegrationsBasePath             string
	KMSBasePath                      string
	LoggingBasePath                  string
	LookerBasePath                   string
	MemcacheBasePath                 string
	MigrationCenterBasePath          string
	MLEngineBasePath                 string
	MonitoringBasePath               string
	NetappBasePath                   string
	NetworkConnectivityBasePath      string
	NetworkManagementBasePath        string
	NetworkSecurityBasePath          string
	NetworkServicesBasePath          string
	NotebooksBasePath                string
	OrgPolicyBasePath                string
	OSConfigBasePath                 string
	OSLoginBasePath                  string
	PrivatecaBasePath                string
	PrivilegedAccessManagerBasePath  string
	PublicCABasePath                 string
	PubsubBasePath                   string
	PubsubLiteBasePath               string
	RedisBasePath                    string
	ResourceManagerBasePath          string
	SecretManagerBasePath            string
	SecureSourceManagerBasePath      string
	SecurityCenterBasePath           string
	SecurityCenterManagementBasePath string
	SecurityCenterV2BasePath         string
	SecuritypostureBasePath          string
	ServiceManagementBasePath        string
	ServiceNetworkingBasePath        string
	ServiceUsageBasePath             string
	SiteVerificationBasePath         string
	SourceRepoBasePath               string
	SpannerBasePath                  string
	SQLBasePath                      string
	StorageBasePath                  string
	StorageInsightsBasePath          string
	StorageTransferBasePath          string
	TagsBasePath                     string
	TPUBasePath                      string
	VertexAIBasePath                 string
	VmwareengineBasePath             string
	VPCAccessBasePath                string
	WorkbenchBasePath                string
	WorkflowsBasePath                string
}
