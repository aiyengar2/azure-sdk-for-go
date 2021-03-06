package backupapi

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/recoveryservices/mgmt/2017-07-01/backup"
	"github.com/Azure/go-autorest/autorest"
)

// ProtectionIntentClientAPI contains the set of methods on the ProtectionIntentClient type.
type ProtectionIntentClientAPI interface {
	CreateOrUpdate(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, intentObjectName string, parameters backup.ProtectionIntentResource) (result backup.ProtectionIntentResource, err error)
	Delete(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, intentObjectName string) (result autorest.Response, err error)
	Get(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, intentObjectName string) (result backup.ProtectionIntentResource, err error)
	Validate(ctx context.Context, azureRegion string, parameters backup.PreValidateEnableBackupRequest) (result backup.PreValidateEnableBackupResponse, err error)
}

var _ ProtectionIntentClientAPI = (*backup.ProtectionIntentClient)(nil)

// StatusClientAPI contains the set of methods on the StatusClient type.
type StatusClientAPI interface {
	Get(ctx context.Context, azureRegion string, parameters backup.StatusRequest) (result backup.StatusResponse, err error)
}

var _ StatusClientAPI = (*backup.StatusClient)(nil)

// FeatureSupportClientAPI contains the set of methods on the FeatureSupportClient type.
type FeatureSupportClientAPI interface {
	Validate(ctx context.Context, azureRegion string, parameters backup.BasicFeatureSupportRequest) (result backup.AzureVMResourceFeatureSupportResponse, err error)
}

var _ FeatureSupportClientAPI = (*backup.FeatureSupportClient)(nil)

// ProtectionIntentGroupClientAPI contains the set of methods on the ProtectionIntentGroupClient type.
type ProtectionIntentGroupClientAPI interface {
	List(ctx context.Context, vaultName string, resourceGroupName string, filter string, skipToken string) (result backup.ProtectionIntentResourceListPage, err error)
	ListComplete(ctx context.Context, vaultName string, resourceGroupName string, filter string, skipToken string) (result backup.ProtectionIntentResourceListIterator, err error)
}

var _ ProtectionIntentGroupClientAPI = (*backup.ProtectionIntentGroupClient)(nil)

// UsageSummariesClientAPI contains the set of methods on the UsageSummariesClient type.
type UsageSummariesClientAPI interface {
	List(ctx context.Context, vaultName string, resourceGroupName string, filter string, skipToken string) (result backup.ManagementUsageList, err error)
}

var _ UsageSummariesClientAPI = (*backup.UsageSummariesClient)(nil)

// EnginesClientAPI contains the set of methods on the EnginesClient type.
type EnginesClientAPI interface {
	Get(ctx context.Context, vaultName string, resourceGroupName string, backupEngineName string, filter string, skipToken string) (result backup.EngineBaseResource, err error)
	List(ctx context.Context, vaultName string, resourceGroupName string, filter string, skipToken string) (result backup.EngineBaseResourceListPage, err error)
	ListComplete(ctx context.Context, vaultName string, resourceGroupName string, filter string, skipToken string) (result backup.EngineBaseResourceListIterator, err error)
}

var _ EnginesClientAPI = (*backup.EnginesClient)(nil)

// ProtectionContainerRefreshOperationResultsClientAPI contains the set of methods on the ProtectionContainerRefreshOperationResultsClient type.
type ProtectionContainerRefreshOperationResultsClientAPI interface {
	Get(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, operationID string) (result autorest.Response, err error)
}

var _ ProtectionContainerRefreshOperationResultsClientAPI = (*backup.ProtectionContainerRefreshOperationResultsClient)(nil)

// ProtectableContainersClientAPI contains the set of methods on the ProtectableContainersClient type.
type ProtectableContainersClientAPI interface {
	List(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, filter string) (result backup.ProtectableContainerResourceListPage, err error)
	ListComplete(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, filter string) (result backup.ProtectableContainerResourceListIterator, err error)
}

var _ ProtectableContainersClientAPI = (*backup.ProtectableContainersClient)(nil)

// ProtectionContainersClientAPI contains the set of methods on the ProtectionContainersClient type.
type ProtectionContainersClientAPI interface {
	Get(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string) (result backup.ProtectionContainerResource, err error)
	Inquire(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, filter string) (result autorest.Response, err error)
	Refresh(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, filter string) (result autorest.Response, err error)
	Register(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, parameters backup.ProtectionContainerResource) (result backup.ProtectionContainerResource, err error)
	Unregister(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string) (result autorest.Response, err error)
}

var _ ProtectionContainersClientAPI = (*backup.ProtectionContainersClient)(nil)

// WorkloadItemsClientAPI contains the set of methods on the WorkloadItemsClient type.
type WorkloadItemsClientAPI interface {
	List(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, filter string, skipToken string) (result backup.WorkloadItemResourceListPage, err error)
	ListComplete(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, filter string, skipToken string) (result backup.WorkloadItemResourceListIterator, err error)
}

var _ WorkloadItemsClientAPI = (*backup.WorkloadItemsClient)(nil)

// ProtectionContainerOperationResultsClientAPI contains the set of methods on the ProtectionContainerOperationResultsClient type.
type ProtectionContainerOperationResultsClientAPI interface {
	Get(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, operationID string) (result backup.ProtectionContainerResource, err error)
}

var _ ProtectionContainerOperationResultsClientAPI = (*backup.ProtectionContainerOperationResultsClient)(nil)

// BackupsClientAPI contains the set of methods on the BackupsClient type.
type BackupsClientAPI interface {
	Trigger(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, protectedItemName string, parameters backup.RequestResource) (result autorest.Response, err error)
}

var _ BackupsClientAPI = (*backup.BackupsClient)(nil)

// ProtectedItemOperationStatusesClientAPI contains the set of methods on the ProtectedItemOperationStatusesClient type.
type ProtectedItemOperationStatusesClientAPI interface {
	Get(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, protectedItemName string, operationID string) (result backup.OperationStatus, err error)
}

var _ ProtectedItemOperationStatusesClientAPI = (*backup.ProtectedItemOperationStatusesClient)(nil)

// ItemLevelRecoveryConnectionsClientAPI contains the set of methods on the ItemLevelRecoveryConnectionsClient type.
type ItemLevelRecoveryConnectionsClientAPI interface {
	Provision(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, protectedItemName string, recoveryPointID string, parameters backup.ILRRequestResource) (result autorest.Response, err error)
	Revoke(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, protectedItemName string, recoveryPointID string) (result autorest.Response, err error)
}

var _ ItemLevelRecoveryConnectionsClientAPI = (*backup.ItemLevelRecoveryConnectionsClient)(nil)

// OperationResultsClientAPI contains the set of methods on the OperationResultsClient type.
type OperationResultsClientAPI interface {
	Get(ctx context.Context, vaultName string, resourceGroupName string, operationID string) (result autorest.Response, err error)
}

var _ OperationResultsClientAPI = (*backup.OperationResultsClient)(nil)

// OperationStatusesClientAPI contains the set of methods on the OperationStatusesClient type.
type OperationStatusesClientAPI interface {
	Get(ctx context.Context, vaultName string, resourceGroupName string, operationID string) (result backup.OperationStatus, err error)
}

var _ OperationStatusesClientAPI = (*backup.OperationStatusesClient)(nil)

// ProtectionPoliciesClientAPI contains the set of methods on the ProtectionPoliciesClient type.
type ProtectionPoliciesClientAPI interface {
	Delete(ctx context.Context, vaultName string, resourceGroupName string, policyName string) (result autorest.Response, err error)
}

var _ ProtectionPoliciesClientAPI = (*backup.ProtectionPoliciesClient)(nil)

// ProtectionPolicyOperationStatusesClientAPI contains the set of methods on the ProtectionPolicyOperationStatusesClient type.
type ProtectionPolicyOperationStatusesClientAPI interface {
	Get(ctx context.Context, vaultName string, resourceGroupName string, policyName string, operationID string) (result backup.OperationStatus, err error)
}

var _ ProtectionPolicyOperationStatusesClientAPI = (*backup.ProtectionPolicyOperationStatusesClient)(nil)

// ProtectableItemsClientAPI contains the set of methods on the ProtectableItemsClient type.
type ProtectableItemsClientAPI interface {
	List(ctx context.Context, vaultName string, resourceGroupName string, filter string, skipToken string) (result backup.WorkloadProtectableItemResourceListPage, err error)
	ListComplete(ctx context.Context, vaultName string, resourceGroupName string, filter string, skipToken string) (result backup.WorkloadProtectableItemResourceListIterator, err error)
}

var _ ProtectableItemsClientAPI = (*backup.ProtectableItemsClient)(nil)

// ProtectionContainersGroupClientAPI contains the set of methods on the ProtectionContainersGroupClient type.
type ProtectionContainersGroupClientAPI interface {
	List(ctx context.Context, vaultName string, resourceGroupName string, filter string) (result backup.ProtectionContainerResourceListPage, err error)
	ListComplete(ctx context.Context, vaultName string, resourceGroupName string, filter string) (result backup.ProtectionContainerResourceListIterator, err error)
}

var _ ProtectionContainersGroupClientAPI = (*backup.ProtectionContainersGroupClient)(nil)

// SecurityPINsClientAPI contains the set of methods on the SecurityPINsClient type.
type SecurityPINsClientAPI interface {
	Get(ctx context.Context, vaultName string, resourceGroupName string) (result backup.TokenInformation, err error)
}

var _ SecurityPINsClientAPI = (*backup.SecurityPINsClient)(nil)

// ResourceStorageConfigsClientAPI contains the set of methods on the ResourceStorageConfigsClient type.
type ResourceStorageConfigsClientAPI interface {
	Get(ctx context.Context, vaultName string, resourceGroupName string) (result backup.ResourceConfigResource, err error)
	Patch(ctx context.Context, vaultName string, resourceGroupName string, parameters backup.ResourceConfigResource) (result autorest.Response, err error)
	Update(ctx context.Context, vaultName string, resourceGroupName string, parameters backup.ResourceConfigResource) (result backup.ResourceConfigResource, err error)
}

var _ ResourceStorageConfigsClientAPI = (*backup.ResourceStorageConfigsClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result backup.ClientDiscoveryResponsePage, err error)
	ListComplete(ctx context.Context) (result backup.ClientDiscoveryResponseIterator, err error)
}

var _ OperationsClientAPI = (*backup.OperationsClient)(nil)
