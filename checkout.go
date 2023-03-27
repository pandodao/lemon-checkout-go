package lemon

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
)

type (
	CheckoutAttributes struct {
		StoreID   int64  `json:"store_id"`
		VariantID int64  `json:"variant_id"`
		URL       string `json:"url"`

		ProductOptions struct {
			EanbledVariants []int64 `json:"enabled_variants"`
			RedirectURL     string  `json:"redirect_url"`
		} `json:"product_options"`

		CheckoutOptions interface{} `json:"checkout_options"`

		CheckoutData CheckoutData `json:"checkout_data"`

		ExpiresAt string `json:"expires_at"`
		TestMode  bool   `json:"test_mode"`
	}

	CheckoutData struct {
		Email          string   `json:"email,omitempty"`
		Name           string   `json:"name,omitempty"`
		BillingAddress []string `json:"billing_address,omitempty"`
		TaxNumber      string   `json:"tax_number,omitempty"`
		DiscountCode   string   `json:"discount_code,omitempty"`

		Custom any `json:"custom"`
	}

	CheckoutResult struct {
		ResourceObject
		Attributes CheckoutAttributes `json:"attributes"`
	}
)

func (c *Client) CreateCheckoutSimple(ctx context.Context, storeID int64, variantIDs []int64, redirectURL string, checkoutData CheckoutData) (*CheckoutResult, error) {
	attrs := CheckoutAttributes{}
	attrs.StoreID = storeID
	attrs.VariantID = variantIDs[0]
	attrs.ProductOptions.EanbledVariants = variantIDs
	attrs.ProductOptions.RedirectURL = redirectURL
	attrs.CheckoutData = checkoutData

	rels := Relationships{}
	rels.Store.Data.Type = "stores"
	rels.Store.Data.ID = strconv.FormatInt(storeID, 10)
	rels.Variant.Data.Type = "variants"
	rels.Variant.Data.ID = strconv.FormatInt(variantIDs[0], 10)

	return c.CreateCheckout(ctx, attrs, rels)
}

func (c *Client) CreateCheckout(ctx context.Context, attrs CheckoutAttributes, rels Relationships) (*CheckoutResult, error) {
	requestWrapper := &GeneralWrapper{
		Data: ResourceObject{
			Type:          "checkouts",
			Attributes:    attrs,
			Relationships: rels,
		},
	}

	output := CheckoutResult{}

	if err := c.request(ctx, http.MethodPost, "/checkouts", requestWrapper, &output); err != nil {
		fmt.Printf("c.Do err: %v\n", err)
		return nil, err
	}

	return &output, nil
}
