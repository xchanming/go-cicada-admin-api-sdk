package go_cicada_admin_sdk

import (
	"net/http"

	"time"
)

type ProductRepository ClientService

func (t ProductRepository) Search(ctx ApiContext, criteria Criteria) (*ProductCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/product", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(ProductCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t ProductRepository) SearchAll(ctx ApiContext, criteria Criteria) (*ProductCollection, *http.Response, error) {
	if criteria.Limit == 0 {
		criteria.Limit = 50
	}

	if criteria.Page == 0 {
		criteria.Page = 1
	}

	c, resp, err := t.Search(ctx, criteria)

	if err != nil {
		return c, resp, err
	}

	for {
		criteria.Page++

		nextC, nextResp, nextErr := t.Search(ctx, criteria)

		if nextErr != nil {
			return c, nextResp, nextErr
		}

		if len(nextC.Data) == 0 {
			break
		}

		c.Data = append(c.Data, nextC.Data...)
	}

	c.Total = int64(len(c.Data))

	return c, resp, err
}

func (t ProductRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/product", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(SearchIdsResponse)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t ProductRepository) Upsert(ctx ApiContext, entity []Product) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"product": {
		Entity:  "product",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t ProductRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"product": {
		Entity:  "product",
		Action:  "delete",
		Payload: payload,
	}})
}

type Product struct {
	MainVariantId string `json:"mainVariantId,omitempty"`

	PurchaseUnit float64 `json:"purchaseUnit,omitempty"`

	WhitelistIds interface{} `json:"whitelistIds,omitempty"`

	ManufacturerId string `json:"manufacturerId,omitempty"`

	Stock float64 `json:"stock,omitempty"`

	Available bool `json:"available,omitempty"`

	BlacklistIds interface{} `json:"blacklistIds,omitempty"`

	MetaDescription string `json:"metaDescription,omitempty"`

	Children []Product `json:"children,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	VersionId string `json:"versionId,omitempty"`

	PurchaseSteps float64 `json:"purchaseSteps,omitempty"`

	MarkAsTopseller bool `json:"markAsTopseller,omitempty"`

	Parent *Product `json:"parent,omitempty"`

	SearchKeywords []ProductSearchKeyword `json:"searchKeywords,omitempty"`

	CategoriesRo []Category `json:"categoriesRo,omitempty"`

	CustomFieldSets []CustomFieldSet `json:"customFieldSets,omitempty"`

	TaxId string `json:"taxId,omitempty"`

	Weight float64 `json:"weight,omitempty"`

	Keywords string `json:"keywords,omitempty"`

	SeoUrls []SeoUrl `json:"seoUrls,omitempty"`

	AutoIncrement float64 `json:"autoIncrement,omitempty"`

	AvailableStock float64 `json:"availableStock,omitempty"`

	ConfiguratorSettings []ProductConfiguratorSetting `json:"configuratorSettings,omitempty"`

	Visibilities []ProductVisibility `json:"visibilities,omitempty"`

	FeatureSetId string `json:"featureSetId,omitempty"`

	ProductReviews []ProductReview `json:"productReviews,omitempty"`

	ProductManufacturerVersionId string `json:"productManufacturerVersionId,omitempty"`

	MinPurchase float64 `json:"minPurchase,omitempty"`

	ReferenceUnit float64 `json:"referenceUnit,omitempty"`

	Width float64 `json:"width,omitempty"`

	PropertyIds interface{} `json:"propertyIds,omitempty"`

	DeliveryTime *DeliveryTime `json:"deliveryTime,omitempty"`

	Manufacturer *ProductManufacturer `json:"manufacturer,omitempty"`

	Id string `json:"id,omitempty"`

	ListingPrices interface{} `json:"listingPrices,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	MainCategories []MainCategory `json:"mainCategories,omitempty"`

	Variation interface{} `json:"variation,omitempty"`

	Tax *Tax `json:"tax,omitempty"`

	ParentVersionId string `json:"parentVersionId,omitempty"`

	Active bool `json:"active,omitempty"`

	IsCloseout bool `json:"isCloseout,omitempty"`

	ManufacturerNumber string `json:"manufacturerNumber,omitempty"`

	ChildCount float64 `json:"childCount,omitempty"`

	MetaTitle string `json:"metaTitle,omitempty"`

	PackUnit string `json:"packUnit,omitempty"`

	RestockTime float64 `json:"restockTime,omitempty"`

	MaxPurchase float64 `json:"maxPurchase,omitempty"`

	RatingAverage float64 `json:"ratingAverage,omitempty"`

	Categories []Category `json:"categories,omitempty"`

	Tags []Tag `json:"tags,omitempty"`

	Ean string `json:"ean,omitempty"`

	TagIds interface{} `json:"tagIds,omitempty"`

	Description string `json:"description,omitempty"`

	Prices []ProductPrice `json:"prices,omitempty"`

	PurchasePrices interface{} `json:"purchasePrices,omitempty"`

	Name string `json:"name,omitempty"`

	OrderLineItems []OrderLineItem `json:"orderLineItems,omitempty"`

	CoverId string `json:"coverId,omitempty"`

	DeliveryTimeId string `json:"deliveryTimeId,omitempty"`

	PurchasePrice float64 `json:"purchasePrice,omitempty"`

	Length float64 `json:"length,omitempty"`

	PackUnitPlural string `json:"packUnitPlural,omitempty"`

	Unit *Unit `json:"unit,omitempty"`

	Media []ProductMedia `json:"media,omitempty"`

	Options []PropertyGroupOption `json:"options,omitempty"`

	UnitId string `json:"unitId,omitempty"`

	Price interface{} `json:"price,omitempty"`

	DisplayGroup string `json:"displayGroup,omitempty"`

	Properties []PropertyGroupOption `json:"properties,omitempty"`

	FeatureSet *ProductFeatureSet `json:"featureSet,omitempty"`

	Translated interface{} `json:"translated,omitempty"`

	VariantRestrictions interface{} `json:"variantRestrictions,omitempty"`

	CategoryTree interface{} `json:"categoryTree,omitempty"`

	Translations []ProductTranslation `json:"translations,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	ShippingFree bool `json:"shippingFree,omitempty"`

	ReleaseDate time.Time `json:"releaseDate,omitempty"`

	Cover *ProductMedia `json:"cover,omitempty"`

	CrossSellings []ProductCrossSelling `json:"crossSellings,omitempty"`

	CrossSellingAssignedProducts []ProductCrossSellingAssignedProducts `json:"crossSellingAssignedProducts,omitempty"`

	ParentId string `json:"parentId,omitempty"`

	ProductMediaVersionId string `json:"productMediaVersionId,omitempty"`

	ProductNumber string `json:"productNumber,omitempty"`

	ConfiguratorGroupConfig interface{} `json:"configuratorGroupConfig,omitempty"`

	Height float64 `json:"height,omitempty"`

	OptionIds interface{} `json:"optionIds,omitempty"`

	CustomFieldSetSelectionActive bool `json:"customFieldSetSelectionActive,omitempty"`
}

type ProductCollection struct {
	EntityCollection

	Data []Product `json:"data"`
}
