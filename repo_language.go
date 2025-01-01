package go_cicada_admin_sdk

import (
	"net/http"

	"time"
)

type LanguageRepository ClientService

func (t LanguageRepository) Search(ctx ApiContext, criteria Criteria) (*LanguageCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/language", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(LanguageCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t LanguageRepository) SearchAll(ctx ApiContext, criteria Criteria) (*LanguageCollection, *http.Response, error) {
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

func (t LanguageRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/language", criteria)

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

func (t LanguageRepository) Upsert(ctx ApiContext, entity []Language) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"language": {
		Entity:  "language",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t LanguageRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"language": {
		Entity:  "language",
		Action:  "delete",
		Payload: payload,
	}})
}

type Language struct {
	LocaleId string `json:"localeId,omitempty"`

	UnitTranslations []UnitTranslation `json:"unitTranslations,omitempty"`

	MailHeaderFooterTranslations []MailHeaderFooterTranslation `json:"mailHeaderFooterTranslations,omitempty"`

	ProductKeywordDictionaries []ProductKeywordDictionary `json:"productKeywordDictionaries,omitempty"`

	ThemeTranslations []ThemeTranslation `json:"themeTranslations,omitempty"`

	SalutationTranslations []SalutationTranslation `json:"salutationTranslations,omitempty"`

	Locale *Locale `json:"locale,omitempty"`

	TranslationCode *Locale `json:"translationCode,omitempty"`

	SalesChannelDomains []SalesChannelDomain `json:"salesChannelDomains,omitempty"`

	NumberRangeTypeTranslations []NumberRangeTypeTranslation `json:"numberRangeTypeTranslations,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	TranslationCodeId string `json:"translationCodeId,omitempty"`

	Name string `json:"name,omitempty"`

	Customers []Customer `json:"customers,omitempty"`

	PaymentMethodTranslations []PaymentMethodTranslation `json:"paymentMethodTranslations,omitempty"`

	ProductReviews []ProductReview `json:"productReviews,omitempty"`

	ProductSortingTranslations []ProductSortingTranslation `json:"productSortingTranslations,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	Parent *Language `json:"parent,omitempty"`

	Children []Language `json:"children,omitempty"`

	CountryStateTranslations []CountryStateTranslation `json:"countryStateTranslations,omitempty"`

	ProductTranslations []ProductTranslation `json:"productTranslations,omitempty"`

	ProductFeatureSetTranslations []ProductFeatureSetTranslation `json:"productFeatureSetTranslations,omitempty"`

	Orders []Order `json:"orders,omitempty"`

	CategoryTranslations []CategoryTranslation `json:"categoryTranslations,omitempty"`

	CountryTranslations []CountryTranslation `json:"countryTranslations,omitempty"`

	PropertyGroupOptionTranslations []PropertyGroupOptionTranslation `json:"propertyGroupOptionTranslations,omitempty"`

	CustomerGroupTranslations []CustomerGroupTranslation `json:"customerGroupTranslations,omitempty"`

	SalesChannelTranslations []SalesChannelTranslation `json:"salesChannelTranslations,omitempty"`

	ProductStreamTranslations []ProductStreamTranslation `json:"productStreamTranslations,omitempty"`

	StateMachineTranslations []StateMachineTranslation `json:"stateMachineTranslations,omitempty"`

	ProductSearchKeywords []ProductSearchKeyword `json:"productSearchKeywords,omitempty"`

	ProductManufacturerTranslations []ProductManufacturerTranslation `json:"productManufacturerTranslations,omitempty"`

	CmsPageTranslations []CmsPageTranslation `json:"cmsPageTranslations,omitempty"`

	PropertyGroupTranslations []PropertyGroupTranslation `json:"propertyGroupTranslations,omitempty"`

	MailTemplateTranslations []MailTemplateTranslation `json:"mailTemplateTranslations,omitempty"`

	MailTemplateTypeTranslations []MailTemplateTypeTranslation `json:"mailTemplateTypeTranslations,omitempty"`

	NewsletterRecipients []NewsletterRecipient `json:"newsletterRecipients,omitempty"`

	PluginTranslations []PluginTranslation `json:"pluginTranslations,omitempty"`

	StateMachineStateTranslations []StateMachineStateTranslation `json:"stateMachineStateTranslations,omitempty"`

	DeliveryTimeTranslations []DeliveryTimeTranslation `json:"deliveryTimeTranslations,omitempty"`

	SeoUrlTranslations []SeoUrl `json:"seoUrlTranslations,omitempty"`

	TaxRuleTypeTranslations []TaxRuleTypeTranslation `json:"taxRuleTypeTranslations,omitempty"`

	ProductCrossSellingTranslations []ProductCrossSellingTranslation `json:"productCrossSellingTranslations,omitempty"`

	Id string `json:"id,omitempty"`

	SalesChannelDefaultAssignments []SalesChannel `json:"salesChannelDefaultAssignments,omitempty"`

	MediaTranslations []MediaTranslation `json:"mediaTranslations,omitempty"`

	SalesChannelTypeTranslations []SalesChannelTypeTranslation `json:"salesChannelTypeTranslations,omitempty"`

	PromotionTranslations []PromotionTranslation `json:"promotionTranslations,omitempty"`

	NumberRangeTranslations []NumberRangeTranslation `json:"numberRangeTranslations,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	SalesChannels []SalesChannel `json:"salesChannels,omitempty"`

	ParentId string `json:"parentId,omitempty"`

	CurrencyTranslations []CurrencyTranslation `json:"currencyTranslations,omitempty"`

	LocaleTranslations []LocaleTranslation `json:"localeTranslations,omitempty"`

	ShippingMethodTranslations []ShippingMethodTranslation `json:"shippingMethodTranslations,omitempty"`

	CmsSlotTranslations []CmsSlotTranslation `json:"cmsSlotTranslations,omitempty"`

	ImportExportProfileTranslations []ImportExportProfileTranslation `json:"importExportProfileTranslations,omitempty"`
}

type LanguageCollection struct {
	EntityCollection

	Data []Language `json:"data"`
}
