package main

type AutoGenerated []struct {
	StoreID int `json:"storeId"`
	Menu    []struct {
		ID                int    `json:"id"`
		Name              string `json:"name"`
		MenuSubcategories []struct {
			ID                             int    `json:"id"`
			Name                           string `json:"name"`
			MaxNumberItemsAllowedPerBasket int    `json:"maxNumberItemsAllowedPerBasket"`
			PhotoFilename                  string `json:"photoFilename"`
			HideFromCustomer               bool   `json:"hideFromCustomer"`
			HideOnCarousel                 bool   `json:"hideOnCarousel"`
			PriceByWeight                  bool   `json:"priceByWeight"`
			MenuItems                      []struct {
				ID                                        int     `json:"id"`
				MenuSubcategoryID                         int     `json:"menuSubcategoryId"`
				Name                                      string  `json:"name"`
				Description                               string  `json:"description"`
				Ingredients                               string  `json:"ingredients"`
				AdditionalInfo                            string  `json:"additionalInfo"`
				PreparationInfo                           string  `json:"preparationInfo"`
				StoringInfo                               string  `json:"storingInfo"`
				BestBeforeDays                            int     `json:"bestBeforeDays"`
				BestBeforeHours                           int     `json:"bestBeforeHours"`
				AgeRestriction                            int     `json:"ageRestriction"`
				PhotoFilename                             string  `json:"photoFilename"`
				PreparationTimeInSeconds                  int     `json:"preparationTimeInSeconds"`
				RequiredNoticeInSeconds                   int     `json:"requiredNoticeInSeconds"`
				IsFMS                                     bool    `json:"isFMS"`
				MonOpenHour                               int     `json:"monOpenHour,omitempty"`
				MonOpenMinute                             int     `json:"monOpenMinute,omitempty"`
				MonCloseHour                              int     `json:"monCloseHour,omitempty"`
				MonCloseMinute                            int     `json:"monCloseMinute,omitempty"`
				BypassTimeResOnKiosk                      bool    `json:"bypassTimeResOnKiosk"`
				BypassTimeResOnEpos                       bool    `json:"bypassTimeResOnEpos"`
				BypassTimeResOnMobile                     bool    `json:"bypassTimeResOnMobile"`
				BypassTimeResOnWebapp                     bool    `json:"bypassTimeResOnWebapp"`
				Hidden                                    bool    `json:"hidden"`
				Recommended                               bool    `json:"recommended"`
				Popular                                   bool    `json:"popular"`
				Collectible                               bool    `json:"collectible"`
				Deliverable                               bool    `json:"deliverable"`
				ExcludedFromOfficeCollection              bool    `json:"excludedFromOfficeCollection"`
				ExcludedFromHospitality                   bool    `json:"excludedFromHospitality"`
				PrintAdditionalInfo                       bool    `json:"printAdditionalInfo"`
				PrintStoringInfo                          bool    `json:"printStoringInfo"`
				PrintIngredients                          bool    `json:"printIngredients"`
				PrintPreparationInfo                      bool    `json:"printPreparationInfo"`
				PrintNetWeight                            bool    `json:"printNetWeight"`
				PrintBestBefore                           bool    `json:"printBestBefore"`
				PrintItemName                             bool    `json:"printItemName"`
				PrintAllergens                            bool    `json:"printAllergens"`
				PrintNutrients                            bool    `json:"printNutrients"`
				PrintAdditives                            bool    `json:"printAdditives"`
				PrintCuisines                             bool    `json:"printCuisines"`
				PrintDietaryLabels                        bool    `json:"printDietaryLabels"`
				PrintTime                                 bool    `json:"printTime"`
				PrintShelfEdgeLabel                       bool    `json:"printShelfEdgeLabel"`
				PrintForClickAndCollect                   bool    `json:"printForClickAndCollect"`
				CompleteForScanAndGo                      bool    `json:"completeForScanAndGo"`
				PriceByWeight                             bool    `json:"priceByWeight"`
				ShowForWeighAndPay                        bool    `json:"showForWeighAndPay"`
				AlcoholABV                                float64 `json:"alcoholABV"`
				AlcoholABVNew                             float64 `json:"alcoholABVNew"`
				AlcoholUnit                               float64 `json:"alcoholUnit"`
				CarbonFootprint                           float64 `json:"carbonFootprint"`
				CarbonFootprintCategory                   string  `json:"carbonFootprintCategory"`
				StoringTemperature                        string  `json:"storingTemperature"`
				TemperatureUnit                           string  `json:"temperatureUnit"`
				MinQuantity                               int     `json:"minQuantity"`
				MaxQuantity                               int     `json:"maxQuantity"`
				TaxRate                                   float64 `json:"taxRate"`
				TaxRateTA                                 float64 `json:"taxRateTA"`
				DiscountRate                              float64 `json:"discountRate"`
				ShowForScanAndGo                          bool    `json:"showForScanAndGo"`
				CanCancel                                 bool    `json:"canCancel"`
				CanEdit                                   bool    `json:"canEdit"`
				CancellationTimeInSeconds                 int     `json:"cancellationTimeInSeconds"`
				EditTimeInSeconds                         int     `json:"editTimeInSeconds"`
				CanCancelAfterAccept                      bool    `json:"canCancelAfterAccept"`
				CanEditAfterAccept                        bool    `json:"canEditAfterAccept"`
				CancellationTimeInSecondsAfterAccept      int     `json:"cancellationTimeInSecondsAfterAccept"`
				EditTimeInSecondsAfterAccept              int     `json:"editTimeInSecondsAfterAccept"`
				CanCancelBeforeCollection                 bool    `json:"canCancelBeforeCollection"`
				CanEditBeforeCollection                   bool    `json:"canEditBeforeCollection"`
				CancellationTimeInSecondsBeforeCollection int     `json:"cancellationTimeInSecondsBeforeCollection"`
				EditTimeInSecondsBeforeCollection         int     `json:"editTimeInSecondsBeforeCollection"`
				MultistepModifiers                        bool    `json:"multistepModifiers"`
				TaxCategory                               int     `json:"taxCategory"`
				MenuItemLevels                            []struct {
					ID                  int      `json:"id"`
					Name                string   `json:"name"`
					Price               int      `json:"price"`
					Hidden              bool     `json:"hidden"`
					RecycleDeposit      int      `json:"recycleDeposit"`
					CateringDeposit     int      `json:"cateringDeposit"`
					ReplenishPeriod     int      `json:"replenishPeriod"`
					RestockBuffer       int      `json:"restockBuffer"`
					ParMinimumStock     int      `json:"parMinimumStock"`
					ParMaximumStock     int      `json:"parMaximumStock"`
					ManualRestockAmount int      `json:"manualRestockAmount"`
					DaysUntilDepletion  int      `json:"daysUntilDepletion"`
					NetWeight           float64  `json:"netWeight"`
					NetWeightUnit       string   `json:"netWeightUnit"`
					NutriScore          string   `json:"nutriScore"`
					FixedDiscount       int      `json:"fixedDiscount"`
					QrCode              string   `json:"qrCode"`
					StorageLevel        int      `json:"storageLevel"`
					InventoryLevel      int      `json:"inventoryLevel"`
					PhotoFilename       string   `json:"photoFilename"`
					CarbonFootprint     float64  `json:"carbonFootprint"`
					QrCodes             []string `json:"qrCodes"`
					PriceGroups         struct {
					} `json:"priceGroups"`
					LinkedModifiers []any `json:"linkedModifiers"`
					EnergyLevels    []struct {
						Amount           float64 `json:"amount"`
						DailyRecommended int     `json:"dailyRecommended"`
						Unit             string  `json:"unit"`
						Ordinality       int     `json:"ordinality"`
						Name             string  `json:"name"`
						ID               int     `json:"id"`
						Fr               struct {
							Name string `json:"name"`
						} `json:"fr"`
						Es struct {
							Name string `json:"name"`
						} `json:"es"`
					} `json:"energyLevels"`
					OtherNutrients []struct {
						Amount           float64 `json:"amount"`
						DailyRecommended int     `json:"dailyRecommended"`
						Unit             string  `json:"unit"`
						Ordinality       int     `json:"ordinality"`
						Name             string  `json:"name"`
						ID               int     `json:"id"`
						Es               struct {
							Name string `json:"name"`
						} `json:"es,omitempty"`
						Us struct {
							Name string `json:"name"`
						} `json:"us,omitempty"`
						En struct {
							Name string `json:"name"`
						} `json:"en,omitempty"`
						Ja struct {
							Name string `json:"name"`
						} `json:"ja,omitempty"`
						Fr struct {
							Name string `json:"name"`
						} `json:"fr,omitempty"`
					} `json:"otherNutrients"`
					Energy struct {
						Amount           float64 `json:"amount"`
						DailyRecommended int     `json:"dailyRecommended"`
						Unit             string  `json:"unit"`
						Ordinality       int     `json:"ordinality"`
						Name             string  `json:"name"`
						ID               int     `json:"id"`
						Fr               struct {
							Name string `json:"name"`
						} `json:"fr"`
						Es struct {
							Name string `json:"name"`
						} `json:"es"`
					} `json:"Energy"`
					Allergens []struct {
						ID            int `json:"id"`
						Contamination int `json:"contamination"`
					} `json:"allergens"`
					Additives     []any `json:"additives"`
					Cuisines      []any `json:"cuisines"`
					DietaryLabels []struct {
						ID   int `json:"id"`
						Type int `json:"type"`
					} `json:"dietaryLabels"`
				} `json:"menuItemLevels"`
				MenuItemModifierGroups               []any `json:"menuItemModifierGroups"`
				MenuItemModifierGroupsMultipleSelect []any `json:"menuItemModifierGroupsMultipleSelect"`
				SuggestedUpsells                     []any `json:"suggestedUpsells"`
				MenuItemCustomerGroups               []any `json:"menuItemCustomerGroups"`
				MenuItemOptionalPhotoFileNames       []any `json:"menuItemOptionalPhotoFileNames"`
				OpeningPeriods                       []any `json:"openingPeriods"`
				MenuItemLevelNutrientAmounts         struct {
					Num14050604 struct {
					} `json:"14050604"`
				} `json:"menuItemLevelNutrientAmounts,omitempty"`
				ActiveAllergens               []int    `json:"activeAllergens"`
				AllergenImages                []string `json:"allergenImages"`
				MinimumItemCalories           string   `json:"minimumItemCalories"`
				TueOpenHour                   int      `json:"tueOpenHour,omitempty"`
				TueOpenMinute                 int      `json:"tueOpenMinute,omitempty"`
				TueCloseHour                  int      `json:"tueCloseHour,omitempty"`
				TueCloseMinute                int      `json:"tueCloseMinute,omitempty"`
				MenuItemLevelNutrientAmounts0 struct {
					Num14050252 struct {
					} `json:"14050252"`
					Num18339990 struct {
					} `json:"18339990"`
				} `json:"menuItemLevelNutrientAmounts,omitempty"`
				MenuItemLevelNutrientAmounts1 struct {
					Num19102838 struct {
					} `json:"19102838"`
				} `json:"menuItemLevelNutrientAmounts,omitempty"`
				WedOpenHour                   int `json:"wedOpenHour,omitempty"`
				WedOpenMinute                 int `json:"wedOpenMinute,omitempty"`
				WedCloseHour                  int `json:"wedCloseHour,omitempty"`
				WedCloseMinute                int `json:"wedCloseMinute,omitempty"`
				MenuItemLevelNutrientAmounts2 struct {
					Num25075402 struct {
					} `json:"25075402"`
				} `json:"menuItemLevelNutrientAmounts,omitempty"`
				MenuItemLevelNutrientAmounts3 struct {
					Num25075396 struct {
					} `json:"25075396"`
				} `json:"menuItemLevelNutrientAmounts,omitempty"`
			} `json:"menuItems"`
		} `json:"menuSubcategories"`
	} `json:"menu"`
}
