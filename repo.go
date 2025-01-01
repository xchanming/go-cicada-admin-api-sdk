package go_cicada_admin_sdk

type Repository struct {
	ClientService

	ProductKeywordDictionary *ProductKeywordDictionaryRepository

	ProductSorting *ProductSortingRepository

	ProductStream *ProductStreamRepository

	ProductVisibility *ProductVisibilityRepository

	OrderDelivery *OrderDeliveryRepository

	Order *OrderRepository

	ProductCrossSellingAssignedProducts *ProductCrossSellingAssignedProductsRepository

	ProductCustomFieldSet *ProductCustomFieldSetRepository

	ProductExport *ProductExportRepository

	ProductProperty *ProductPropertyRepository

	ProductSortingTranslation *ProductSortingTranslationRepository

	ProductTranslation *ProductTranslationRepository

	Document *DocumentRepository

	OrderDeliveryPosition *OrderDeliveryPositionRepository

	OrderTag *OrderTagRepository

	ProductMedia *ProductMediaRepository

	ProductPrice *ProductPriceRepository

	Locale *LocaleRepository

	OrderLineItem *OrderLineItemRepository

	OrderTransaction *OrderTransactionRepository

	ProductCategoryTree *ProductCategoryTreeRepository

	ProductConfiguratorSetting *ProductConfiguratorSettingRepository

	ProductFeatureSetTranslation *ProductFeatureSetTranslationRepository

	ProductReview *ProductReviewRepository

	ProductStreamFilter *ProductStreamFilterRepository

	OrderAddress *OrderAddressRepository

	ProductStreamTranslation *ProductStreamTranslationRepository

	ProductManufacturer *ProductManufacturerRepository

	ProductOption *ProductOptionRepository

	ProductCrossSellingTranslation *ProductCrossSellingTranslationRepository

	Currency *CurrencyRepository

	OrderCustomer *OrderCustomerRepository

	ProductCategory *ProductCategoryRepository

	ProductCrossSelling *ProductCrossSellingRepository

	ProductManufacturerTranslation *ProductManufacturerTranslationRepository

	ProductSearchKeyword *ProductSearchKeywordRepository

	ProductTag *ProductTagRepository

	LocaleTranslation *LocaleTranslationRepository

	Product *ProductRepository

	ProductFeatureSet *ProductFeatureSetRepository

	Language *LanguageRepository
}

func NewRepository(client ClientService) Repository {
	repo := Repository{
		ClientService: client,
	}

	repo.ProductKeywordDictionary = (*ProductKeywordDictionaryRepository)(&client)

	repo.ProductSorting = (*ProductSortingRepository)(&client)

	repo.ProductStream = (*ProductStreamRepository)(&client)

	repo.ProductVisibility = (*ProductVisibilityRepository)(&client)

	repo.OrderDelivery = (*OrderDeliveryRepository)(&client)

	repo.Order = (*OrderRepository)(&client)

	repo.ProductCrossSellingAssignedProducts = (*ProductCrossSellingAssignedProductsRepository)(&client)

	repo.ProductCustomFieldSet = (*ProductCustomFieldSetRepository)(&client)

	repo.ProductExport = (*ProductExportRepository)(&client)

	repo.ProductProperty = (*ProductPropertyRepository)(&client)

	repo.ProductSortingTranslation = (*ProductSortingTranslationRepository)(&client)

	repo.ProductTranslation = (*ProductTranslationRepository)(&client)

	repo.Document = (*DocumentRepository)(&client)

	repo.OrderDeliveryPosition = (*OrderDeliveryPositionRepository)(&client)

	repo.OrderTag = (*OrderTagRepository)(&client)

	repo.ProductMedia = (*ProductMediaRepository)(&client)

	repo.ProductPrice = (*ProductPriceRepository)(&client)

	repo.Locale = (*LocaleRepository)(&client)

	repo.OrderLineItem = (*OrderLineItemRepository)(&client)

	repo.OrderTransaction = (*OrderTransactionRepository)(&client)

	repo.ProductCategoryTree = (*ProductCategoryTreeRepository)(&client)

	repo.ProductConfiguratorSetting = (*ProductConfiguratorSettingRepository)(&client)

	repo.ProductFeatureSetTranslation = (*ProductFeatureSetTranslationRepository)(&client)

	repo.ProductReview = (*ProductReviewRepository)(&client)

	repo.ProductStreamFilter = (*ProductStreamFilterRepository)(&client)

	repo.OrderAddress = (*OrderAddressRepository)(&client)

	repo.ProductStreamTranslation = (*ProductStreamTranslationRepository)(&client)

	repo.ProductManufacturer = (*ProductManufacturerRepository)(&client)

	repo.ProductOption = (*ProductOptionRepository)(&client)

	repo.ProductCrossSellingTranslation = (*ProductCrossSellingTranslationRepository)(&client)

	repo.Currency = (*CurrencyRepository)(&client)

	repo.OrderCustomer = (*OrderCustomerRepository)(&client)

	repo.ProductCategory = (*ProductCategoryRepository)(&client)

	repo.ProductCrossSelling = (*ProductCrossSellingRepository)(&client)

	repo.ProductManufacturerTranslation = (*ProductManufacturerTranslationRepository)(&client)

	repo.ProductSearchKeyword = (*ProductSearchKeywordRepository)(&client)

	repo.ProductTag = (*ProductTagRepository)(&client)

	repo.LocaleTranslation = (*LocaleTranslationRepository)(&client)

	repo.Product = (*ProductRepository)(&client)

	repo.ProductFeatureSet = (*ProductFeatureSetRepository)(&client)

	repo.Language = (*LanguageRepository)(&client)

	return repo
}
