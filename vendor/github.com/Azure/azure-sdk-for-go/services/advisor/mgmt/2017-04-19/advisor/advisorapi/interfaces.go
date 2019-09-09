package advisorapi

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
	"github.com/Azure/azure-sdk-for-go/services/advisor/mgmt/2017-04-19/advisor"
	"github.com/Azure/go-autorest/autorest"
	"github.com/satori/go.uuid"
)

// RecommendationMetadataClientAPI contains the set of methods on the RecommendationMetadataClient type.
type RecommendationMetadataClientAPI interface {
	Get(ctx context.Context, name string) (result advisor.SetObject, err error)
	List(ctx context.Context) (result advisor.MetadataEntityListResultPage, err error)
}

var _ RecommendationMetadataClientAPI = (*advisor.RecommendationMetadataClient)(nil)

// ConfigurationsClientAPI contains the set of methods on the ConfigurationsClient type.
type ConfigurationsClientAPI interface {
	CreateInResourceGroup(ctx context.Context, configContract advisor.ConfigData, resourceGroup string) (result advisor.ARMErrorResponseBody, err error)
	CreateInSubscription(ctx context.Context, configContract advisor.ConfigData) (result advisor.ARMErrorResponseBody, err error)
	ListByResourceGroup(ctx context.Context, resourceGroup string) (result advisor.ConfigurationListResult, err error)
	ListBySubscription(ctx context.Context) (result advisor.ConfigurationListResultPage, err error)
}

var _ ConfigurationsClientAPI = (*advisor.ConfigurationsClient)(nil)

// RecommendationsClientAPI contains the set of methods on the RecommendationsClient type.
type RecommendationsClientAPI interface {
	Generate(ctx context.Context) (result autorest.Response, err error)
	Get(ctx context.Context, resourceURI string, recommendationID string) (result advisor.ResourceRecommendationBase, err error)
	GetGenerateStatus(ctx context.Context, operationID uuid.UUID) (result autorest.Response, err error)
	List(ctx context.Context, filter string, top *int32, skipToken string) (result advisor.ResourceRecommendationBaseListResultPage, err error)
}

var _ RecommendationsClientAPI = (*advisor.RecommendationsClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result advisor.OperationEntityListResultPage, err error)
}

var _ OperationsClientAPI = (*advisor.OperationsClient)(nil)

// SuppressionsClientAPI contains the set of methods on the SuppressionsClient type.
type SuppressionsClientAPI interface {
	Create(ctx context.Context, resourceURI string, recommendationID string, name string, suppressionContract advisor.SuppressionContract) (result advisor.SuppressionContract, err error)
	Delete(ctx context.Context, resourceURI string, recommendationID string, name string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceURI string, recommendationID string, name string) (result advisor.SuppressionContract, err error)
	List(ctx context.Context, top *int32, skipToken string) (result advisor.SuppressionContractListResultPage, err error)
}

var _ SuppressionsClientAPI = (*advisor.SuppressionsClient)(nil)
