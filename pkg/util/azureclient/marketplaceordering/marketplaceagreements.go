package marketplaceordering

//go:generate mockgen -destination=../../../util/mocks/mock_azureclient/mock_$GOPACKAGE/$GOPACKAGE.go github.com/openshift/openshift-azure/pkg/util/azureclient/$GOPACKAGE MarketPlaceAgreementsClient
//go:generate gofmt -s -l -w ../../../util/mocks/mock_azureclient/mock_$GOPACKAGE/$GOPACKAGE.go
//go:generate goimports -local=github.com/openshift/openshift-azure -e -w ../../../util/mocks/mock_azureclient/mock_$GOPACKAGE/$GOPACKAGE.go

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/marketplaceordering/mgmt/2015-06-01/marketplaceordering"
	"github.com/Azure/go-autorest/autorest"
	"github.com/sirupsen/logrus"

	"github.com/openshift/openshift-azure/pkg/util/azureclient"
)

// MarketPlaceAgreementsClient is a minimal interface for azure MarketPlaceAgreementsClient
type MarketPlaceAgreementsClient interface {
	Create(ctx context.Context, publisherID string, offerID string, planID string, parameters marketplaceordering.AgreementTerms) (result marketplaceordering.AgreementTerms, err error)
	Get(ctx context.Context, publisherID string, offerID string, planID string) (result marketplaceordering.AgreementTerms, err error)
}

type marketPlaceAgreementsClient struct {
	marketplaceordering.MarketplaceAgreementsClient
}

var _ MarketPlaceAgreementsClient = &marketPlaceAgreementsClient{}

// NewMarketPlaceAgreementsClient creates a new MarketPlaceAgreementsClient
func NewMarketPlaceAgreementsClient(ctx context.Context, log *logrus.Entry, subscriptionID string, authorizer autorest.Authorizer) MarketPlaceAgreementsClient {
	client := marketplaceordering.NewMarketplaceAgreementsClient(subscriptionID)
	azureclient.SetupClient(ctx, log, "marketplaceordering.MarketplaceAgreementsClient", &client.Client, authorizer)

	return &marketPlaceAgreementsClient{
		MarketplaceAgreementsClient: client,
	}
}
