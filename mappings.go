package countries

//nolint:gochecknoglobals // this is intended for performance reasons.
var mappings = []Mapping{
	{
		Alpha2: "AU",
		Alpha3: "AUS",
		Translations: map[string]Translation{
			"de": {
				Common:      "Australien",              //nolint:misspell // correct german translation
				Official:    "Commonwealth Australien", //nolint:misspell // correct german translation
				Nationality: "Australisch",
			},
			"gb": {
				Common:      "Australia",
				Official:    "Commonwealth of Australia",
				Nationality: "Australian",
			},
		},
	},
	{
		Alpha2: "AT",
		Alpha3: "AUT",
		Translations: map[string]Translation{
			"de": {
				Common:      "Österreich",
				Official:    "Republik Österreich",
				Nationality: "Österreichisch",
			},
			"gb": {
				Common:      "Austria",
				Official:    "Republic of Austria",
				Nationality: "Austrian",
			},
		},
	},
	{
		Alpha2: "AZ",
		Alpha3: "AZE",
		Translations: map[string]Translation{
			"de": {
				Common:      "Aserbaidschan",
				Official:    "Republik Aserbaidschan",
				Nationality: "Aserbaidschaner",
			},
			"gb": {
				Common:      "Azerbaijan",
				Official:    "Republic of Azerbaijan",
				Nationality: "Azerbaijani",
			},
		},
	},
	{
		Alpha2: "AL",
		Alpha3: "ALB",
		Translations: map[string]Translation{
			"de": {
				Common:      "Albanien",
				Official:    "Republik Albanien",
				Nationality: "Albanisch",
			},
			"gb": {
				Common:      "Albania",
				Official:    "Republic of Albania",
				Nationality: "Albanian",
			},
		},
	},
	{
		Alpha2: "DZ",
		Alpha3: "DZA",
		Translations: map[string]Translation{
			"de": {
				Common:      "Algerien",
				Official:    "Demokratische Volksrepublik Algerien",
				Nationality: "Algerisch",
			},
			"gb": {
				Common:      "Algeria",
				Official:    "People's Democratic Republic of Algeria",
				Nationality: "Algerian",
			},
		},
	},
	{
		Alpha2: "AS",
		Alpha3: "ASM",
		Translations: map[string]Translation{
			"de": {
				Official:    "Amerikanisch-Samoa",
				Common:      "Amerikanisch-Samoa",
				Nationality: "Samoaner aus Amerikanisch-Samoa",
			},
			"gb": {
				Common:      "American Samoa",
				Official:    "American Samoa",
				Nationality: "American Samoan",
			},
		},
	},
	{
		Alpha2: "AI",
		Alpha3: "AIA",
		Translations: map[string]Translation{
			"de": {
				Common:      "Anguilla",
				Official:    "Anguilla",
				Nationality: "Anguillianer",
			},
			"gb": {
				Common:      "Anguilla",
				Official:    "Anguilla",
				Nationality: "Anguillan",
			},
		},
	},
	{
		Alpha2: "AO",
		Alpha3: "AGO",
		Translations: map[string]Translation{
			"de": {
				Common:      "Angola",
				Official:    "Republik Angola",
				Nationality: "Angolanisch",
			},
			"gb": {
				Common:      "Angola",
				Official:    "Republic of Angola",
				Nationality: "Angolan",
			},
		},
	},
	{
		Alpha2: "AD",
		Alpha3: "AND",
		Translations: map[string]Translation{
			"de": {
				Common:      "Andorra",
				Official:    "Fürstentum Andorra",
				Nationality: "Andorranisch",
			},
			"gb": {
				Common:      "Andorra",
				Official:    "Principality of Andorra",
				Nationality: "Andorran",
			},
		},
	},
	{
		Alpha2: "AQ",
		Alpha3: "ATA",
		Translations: map[string]Translation{
			"de": {
				Common:   "Antarktis",
				Official: "Antarktika",
			},
			"gb": {
				Common:   "Antarctica",
				Official: "Antarctica",
			},
		},
	},
	{
		Alpha2: "AG",
		Alpha3: "ATG",
		Translations: map[string]Translation{
			"de": {
				Common:      "Antigua und Barbuda",
				Official:    "Antigua und Barbuda",
				Nationality: "Antiguanisch",
			},
			"gb": {
				Common:      "Antigua and Barbuda",
				Official:    "Antigua and Barbuda",
				Nationality: "Antiguans and Barbudans",
			},
		},
	},
	{
		Alpha2: "AE",
		Alpha3: "ARE",
		Translations: map[string]Translation{
			"de": {
				Official:    "Vereinigte Arabische Emirate",
				Common:      "Vereinigte Arabische Emirate",
				Nationality: "Emirati",
			},
			"gb": {
				Common:      "United Arab Emirates",
				Official:    "United Arab Emirates",
				Nationality: "Emirati",
			},
		},
	},
	{
		Alpha2: "AR",
		Alpha3: "ARG",
		Translations: map[string]Translation{
			"de": {
				Official:    "Argentinische Republik",
				Common:      "Argentinien",
				Nationality: "Argentinisch",
			},
			"gb": {
				Common:      "Argentina",
				Official:    "Argentine Republic",
				Nationality: "Argentine",
			},
		},
	},
	{
		Alpha2: "AM",
		Alpha3: "ARM",
		Translations: map[string]Translation{
			"de": {
				Common:      "Armenien",
				Official:    "Republik Armenien",
				Nationality: "Armenisch",
			},
			"gb": {
				Common:      "Armenia",
				Official:    "Republic of Armenia",
				Nationality: "Armenian",
			},
		},
	},
	{
		Alpha2: "AW",
		Alpha3: "ABW",
		Translations: map[string]Translation{
			"de": {
				Common:      "Aruba",
				Official:    "Aruba",
				Nationality: "Aruber",
			},
			"gb": {
				Common:      "Aruba",
				Official:    "Aruba",
				Nationality: "Aruban",
			},
		},
	},
	{
		Alpha2: "AF",
		Alpha3: "AFG",
		Translations: map[string]Translation{
			"de": {
				Common:      "Afghanistan",
				Official:    "Islamische Republik Afghanistan",
				Nationality: "Afghanisch",
			},
			"gb": {
				Common:      "Afghanistan",
				Official:    "Islamic Republic of Afghanistan",
				Nationality: "Afghan",
			},
		},
	},
	{
		Alpha2: "BS",
		Alpha3: "BHS",
		Translations: map[string]Translation{
			"de": {
				Official:    "Commonwealth der Bahamas",
				Common:      "Bahamas",
				Nationality: "Bahamaisch",
			},
			"gb": {
				Common:      "Bahamas",
				Official:    "Commonwealth of the Bahamas",
				Nationality: "Bahamian",
			},
		},
	},
	{
		Alpha2: "BD",
		Alpha3: "BGD",
		Translations: map[string]Translation{
			"de": {
				Common:      "Bangladesch",               //nolint:misspell // correct german translation
				Official:    "Volksrepublik Bangladesch", //nolint:misspell // correct german translation
				Nationality: "Bangladeschisch",
			},
			"gb": {
				Common:      "Bangladesh",
				Official:    "People's Republic of Bangladesh",
				Nationality: "Bangladeshi",
			},
		},
	},
	{
		Alpha2: "BB",
		Alpha3: "BRB",
		Translations: map[string]Translation{
			"de": {
				Common:      "Barbados",
				Official:    "Barbados",
				Nationality: "Barbadisch",
			},
			"gb": {
				Common:      "Barbados",
				Official:    "Barbados",
				Nationality: "Barbadian",
			},
		},
	},
	{
		Alpha2: "BH",
		Alpha3: "BHR",
		Translations: map[string]Translation{
			"de": {
				Official:    "Königreich Bahrain",
				Common:      "Bahrain",
				Nationality: "Bahrainisch",
			},
			"gb": {
				Common:      "Bahrain",
				Official:    "Kingdom of Bahrain",
				Nationality: "Bahraini",
			},
		},
	},
	{
		Alpha2: "BY",
		Alpha3: "BLR",
		Translations: map[string]Translation{
			"de": {
				Common:      "Weißrussland",
				Official:    "Republik Belarus",
				Nationality: "Belarussisch",
			},
			"gb": {
				Common:      "Belarus",
				Official:    "Republic of Belarus",
				Nationality: "Belarusian",
			},
		},
	},
	{
		Alpha2: "BZ",
		Alpha3: "BLZ",
		Translations: map[string]Translation{
			"de": {
				Common:      "Belize",
				Official:    "Belize",
				Nationality: "Belizisch",
			},
			"gb": {
				Common:      "Belize",
				Official:    "Belize",
				Nationality: "Belizean",
			},
		},
	},
	{
		Alpha2: "BE",
		Alpha3: "BEL",
		Translations: map[string]Translation{
			"de": {
				Common:      "Belgien",
				Official:    "Königreich Belgien",
				Nationality: "Belgisch",
			},
			"gb": {
				Common:      "Belgium",
				Official:    "Kingdom of Belgium",
				Nationality: "Belgian",
			},
		},
	},
	{
		Alpha2: "BJ",
		Alpha3: "BEN",
		Translations: map[string]Translation{
			"de": {
				Official:    "Republik Benin",
				Common:      "Benin",
				Nationality: "Beninisch",
			},
			"gb": {
				Common:      "Benin",
				Official:    "Republic of Benin",
				Nationality: "Beninese",
			},
		},
	},
	{
		Alpha2: "BM",
		Alpha3: "BMU",
		Translations: map[string]Translation{
			"de": {
				Official:    "Bermuda",
				Common:      "Bermuda",
				Nationality: "Bermuder",
			},
			"gb": {
				Common:      "Bermuda",
				Official:    "Bermuda",
				Nationality: "Bermudian",
			},
		},
	},
	{
		Alpha2: "BG",
		Alpha3: "BGR",
		Translations: map[string]Translation{
			"de": {
				Common:      "Bulgarien",
				Official:    "Republik Bulgarien",
				Nationality: "Bulgarisch",
			},
			"gb": {
				Common:      "Bulgaria",
				Official:    "Republic of Bulgaria",
				Nationality: "Bulgarian",
			},
		},
	},
	{
		Alpha2: "BO",
		Alpha3: "BOL",
		Translations: map[string]Translation{
			"de": {
				Official:    "Multinationaler Staat von Bolivien",
				Common:      "Bolivien",
				Nationality: "Bolivianisch",
			},
			"gb": {
				Common:      "Bolivia",
				Official:    "Plurinational State of Bolivia",
				Nationality: "Bolivian",
			},
		},
	},
	{
		Alpha2: "BA",
		Alpha3: "BIH",
		Translations: map[string]Translation{
			"de": {
				Common:      "Bosnien und Herzegowina",
				Official:    "Bosnien und Herzegowina",
				Nationality: "Bosnisch-herzegowinisch",
			},
			"gb": {
				Common:      "Bosnia and Herzegovina",
				Official:    "Bosnia and Herzegovina",
				Nationality: "Bosnian and Herzegovinian",
			},
		},
	},
	{
		Alpha2: "BW",
		Alpha3: "BWA",
		Translations: map[string]Translation{
			"de": {
				Common:      "Botswana",
				Official:    "Republik Botswana",
				Nationality: "Botsuanisch",
			},
			"gb": {
				Common:      "Botswana",
				Official:    "Republic of Botswana",
				Nationality: "Botswanan",
			},
		},
	},
	{
		Alpha2: "BR",
		Alpha3: "BRA",
		Translations: map[string]Translation{
			"de": {
				Common:      "Brasilien",
				Official:    "Föderative Republik Brasilien",
				Nationality: "Brasilianisch",
			},
			"gb": {
				Common:      "Brazil",
				Official:    "Federative Republic of Brazil",
				Nationality: "Brazilian",
			},
		},
	},
	{
		Alpha2: "IO",
		Alpha3: "IOT",
		Translations: map[string]Translation{
			"de": {
				Common:   "Britisches Territorium im Indischen Ozean",
				Official: "Britisch-Indischer Ozean",
			},
			"gb": {
				Common:   "British Indian Ocean Territory",
				Official: "British Indian Ocean Territory",
			},
		},
	},
	{
		Alpha2: "BN",
		Alpha3: "BRN",
		Translations: map[string]Translation{
			"de": {
				Common:      "Brunei",
				Official:    "Nation von Brunei, Wohnung des Friedens",
				Nationality: "Bruneiisch",
			},
			"gb": {
				Common:      "Brunei",
				Official:    "Nation of Brunei, Abode of Peace",
				Nationality: "Bruneian",
			},
		},
	},
	{
		Alpha2: "BF",
		Alpha3: "BFA",
		Translations: map[string]Translation{
			"de": {
				Common:      "Burkina Faso",
				Official:    "Burkina Faso",
				Nationality: "Burkinisch",
			},
			"gb": {
				Common:      "Burkina Faso",
				Official:    "Burkina Faso",
				Nationality: "Burkinabe",
			},
		},
	},
	{
		Alpha2: "BI",
		Alpha3: "BDI",
		Translations: map[string]Translation{
			"de": {
				Common:      "Burundi",
				Official:    "Republik Burundi",
				Nationality: "Burundisch",
			},
			"gb": {
				Common:      "Burundi",
				Official:    "Republic of Burundi",
				Nationality: "Burundian",
			},
		},
	},
	{
		Alpha2: "BT",
		Alpha3: "BTN",
		Translations: map[string]Translation{
			"de": {
				Common:      "Bhutan",
				Official:    "Königreich Bhutan",
				Nationality: "Bhutanisch",
			},
			"gb": {
				Common:      "Bhutan",
				Official:    "Kingdom of Bhutan",
				Nationality: "Bhutanese",
			},
		},
	},
	{
		Alpha2: "VU",
		Alpha3: "VUT",
		Translations: map[string]Translation{
			"de": {
				Common:      "Vanuatu",
				Official:    "Vanuatu",
				Nationality: "Vanuatuisch",
			},
			"gb": {
				Common:      "Vanuatu",
				Official:    "Republic of Vanuatu",
				Nationality: "Ni-Vanuatu",
			},
		},
	},
	{
		Alpha2: "VA",
		Alpha3: "VAT",
		Translations: map[string]Translation{
			"de": {
				Common:      "Vatikanstadt",
				Official:    "Staat Vatikanstadt",
				Nationality: "Vatikanisch",
			},
			"gb": {
				Common:      "Vatican City",
				Official:    "Vatican City State",
				Nationality: "Vatican",
			},
		},
	},
	{
		Alpha2: "GB",
		Alpha3: "GBR",
		Translations: map[string]Translation{
			"de": {
				Common:      "Vereinigtes Königreich",
				Official:    "Vereinigtes Königreich Großbritannien und Nordirland",
				Nationality: "Britisch",
			},
			"gb": {
				Common:      "United Kingdom",
				Official:    "United Kingdom of Great Britain and Northern Ireland",
				Nationality: "British",
			},
		},
	},
	{
		Alpha2: "HU",
		Alpha3: "HUN",
		Translations: map[string]Translation{
			"de": {
				Common:      "Ungarn",
				Official:    "Ungarn",
				Nationality: "Ungarisch",
			},
			"gb": {
				Common:      "Hungary",
				Official:    "Hungary",
				Nationality: "Hungarian",
			},
		},
	},
	{
		Alpha2: "VE",
		Alpha3: "VEN",
		Translations: map[string]Translation{
			"de": {
				Common:      "Venezuela",
				Official:    "Bolivarische Republik Venezuela",
				Nationality: "Venezolanisch",
			},
			"gb": {
				Common:      "Venezuela",
				Official:    "Bolivarian Republic of Venezuela",
				Nationality: "Venezuelan",
			},
		},
	},
	{
		Alpha2: "VG",
		Alpha3: "VGB",
		Translations: map[string]Translation{
			"de": {
				Common:      "Britische Jungferninseln",
				Official:    "Jungferninseln",
				Nationality: "Britisch-Virgin-Islander",
			},
			"gb": {
				Common:      "British Virgin Islands",
				Official:    "Virgin Islands",
				Nationality: "British Virgin Islander",
			},
		},
	},
	{
		Alpha2: "VI",
		Alpha3: "VIR",
		Translations: map[string]Translation{
			"de": {
				Official:    "Jungferninseln der Vereinigten Staaten",
				Common:      "Amerikanische Jungferninseln",
				Nationality: "Ameriakner",
			},
			"gb": {
				Common:      "United States Virgin Islands",
				Official:    "Virgin Islands of the United States",
				Nationality: "American",
			},
		},
	},
	{
		Alpha2: "TL",
		Alpha3: "TLS",
		Translations: map[string]Translation{
			"de": {
				Common:      "Timor-Leste",
				Official:    "Demokratische Republik Timor-Leste",
				Nationality: "Osttimorese",
			},
			"gb": {
				Common:      "Timor-Leste",
				Official:    "Democratic Republic of Timor-Leste",
				Nationality: "East Timorese",
			},
		},
	},
	{
		Alpha2: "VN",
		Alpha3: "VNM",
		Translations: map[string]Translation{
			"de": {
				Common:      "Vietnam",
				Official:    "Sozialistische Republik Vietnam",
				Nationality: "Vietnamesisch",
			},
			"gb": {
				Common:      "Vietnam",
				Official:    "Socialist Republic of Vietnam",
				Nationality: "Vietnamese",
			},
		},
	},
	{
		Alpha2: "GA",
		Alpha3: "GAB",
		Translations: map[string]Translation{
			"de": {
				Common:      "Gabun",
				Official:    "Gabunische Republik",
				Nationality: "Gabunisch",
			},
			"gb": {
				Common:      "Gabon",
				Official:    "Gabonese Republic",
				Nationality: "Gabonese",
			},
		},
	},
	{
		Alpha2: "HT",
		Alpha3: "HTI",
		Translations: map[string]Translation{
			"de": {
				Common:      "Haiti",
				Official:    "Republik Haiti",
				Nationality: "Haitianisch",
			},
			"gb": {
				Common:      "Haiti",
				Official:    "Republic of Haiti",
				Nationality: "Haitian",
			},
		},
	},
	{
		Alpha2: "GY",
		Alpha3: "GUY",
		Translations: map[string]Translation{
			"de": {
				Official:    "Guayana",
				Common:      "Französisch Guyana",
				Nationality: "Guyanisch",
			},
			"gb": {
				Common:      "Guyana",
				Official:    "Co-operative Republic of Guyana",
				Nationality: "Guyanese",
			},
		},
	},
	{
		Alpha2: "GM",
		Alpha3: "GMB",
		Translations: map[string]Translation{
			"de": {
				Common:      "Gambia",
				Official:    "Republik Gambia",
				Nationality: "Gambisch",
			},
			"gb": {
				Common:      "Gambia",
				Official:    "Republic of the Gambia",
				Nationality: "Gambian",
			},
		},
	},
	{
		Alpha2: "GH",
		Alpha3: "GHA",
		Translations: map[string]Translation{
			"de": {
				Common:      "Ghana",
				Official:    "Republik Ghana",
				Nationality: "Ghanaisch",
			},
			"gb": {
				Common:      "Ghana",
				Official:    "Republic of Ghana",
				Nationality: "Ghanaian",
			},
		},
	},
	{
		Alpha2: "GP",
		Alpha3: "GLP",
		Translations: map[string]Translation{
			"de": {
				Common:      "Guadeloupe",
				Official:    "Guadeloupe",
				Nationality: "Franzose",
			},
			"gb": {
				Common:      "Guadeloupe",
				Official:    "Guadeloupe",
				Nationality: "French",
			},
		},
	},
	{
		Alpha2: "GT",
		Alpha3: "GTM",
		Translations: map[string]Translation{
			"de": {
				Common:      "Guatemala",
				Official:    "Republik Guatemala",
				Nationality: "Guatemaltekisch",
			},
			"gb": {
				Common:      "Guatemala",
				Official:    "Republic of Guatemala",
				Nationality: "Guatemalan",
			},
		},
	},
	{
		Alpha2: "GN",
		Alpha3: "GIN",
		Translations: map[string]Translation{
			"de": {
				Official:    "Republik Guinea",
				Common:      "Guinea",
				Nationality: "Guineisch",
			},
			"gb": {
				Common:      "Guinea",
				Official:    "Republic of Guinea",
				Nationality: "Guinean",
			},
		},
	},
	{
		Alpha2: "GW",
		Alpha3: "GNB",
		Translations: map[string]Translation{
			"de": {
				Common:      "Guinea-Bissau",
				Official:    "Republik Guinea-Bissau",
				Nationality: "Guinea-bissauisch",
			},
			"gb": {
				Common:      "Guinea-Bissau",
				Official:    "Republic of Guinea-Bissau",
				Nationality: "Guinea-Bissauan",
			},
		},
	},
	{
		Alpha2: "DE",
		Alpha3: "DEU",
		Translations: map[string]Translation{
			"de": {
				Common:      "Deutschland",
				Official:    "Bundesrepublik Deutschland",
				Nationality: "Deutsch",
			},
			"gb": {
				Common:      "Germany",
				Official:    "Federal Republic of Germany",
				Nationality: "German",
			},
		},
	},
	{
		Alpha2: "GI",
		Alpha3: "GIB",
		Translations: map[string]Translation{
			"de": {
				Common:      "Gibraltar",
				Official:    "Gibraltar",
				Nationality: "Gibraltarian",
			},
			"gb": {
				Common:      "Gibraltar",
				Official:    "Gibraltar",
				Nationality: "Gibraltarer",
			},
		},
	},
	{
		Alpha2: "HN",
		Alpha3: "HND",
		Translations: map[string]Translation{
			"de": {
				Common:      "Honduras",
				Official:    "Republik Honduras",
				Nationality: "Honduranisch",
			},
			"gb": {
				Common:      "Honduras",
				Official:    "Republic of Honduras",
				Nationality: "Honduran",
			},
		},
	},
	{
		Alpha2: "HK",
		Alpha3: "HKG",
		Translations: map[string]Translation{
			"de": {
				Common:      "Hongkong",
				Official:    "Sonderverwaltungszone der Volksrepublik China",
				Nationality: "Hongkonger",
			},
			"gb": {
				Common:      "Hong Kong",
				Official:    "Hong Kong Special Administrative Region of the People's Republic of China",
				Nationality: "Hong Konger",
			},
		},
	},
	{
		Alpha2: "GD",
		Alpha3: "GRD",
		Translations: map[string]Translation{
			"de": {
				Official:    "Grenada",
				Common:      "Grenada",
				Nationality: "Grenadisch",
			},
			"gb": {
				Common:      "Grenada",
				Official:    "Grenada",
				Nationality: "Grenadian",
			},
		},
	},
	{
		Alpha2: "GL",
		Alpha3: "GRL",
		Translations: map[string]Translation{
			"de": {
				Common:      "Grönland",
				Official:    "Grönland",
				Nationality: "Grönländer",
			},
			"gb": {
				Common:      "Greenland",
				Official:    "Greenland",
				Nationality: "Greenlander",
			},
		},
	},
	{
		Alpha2: "GR",
		Alpha3: "GRC",
		Translations: map[string]Translation{
			"de": {
				Common:      "Griechenland",
				Official:    "Hellenische Republik",
				Nationality: "Griechisch",
			},
			"gb": {
				Common:      "Greece",
				Official:    "Hellenic Republic",
				Nationality: "Greek",
			},
		},
	},
	{
		Alpha2: "GE",
		Alpha3: "GEO",
		Translations: map[string]Translation{
			"de": {
				Official:    "Georgia",
				Common:      "Georgien",
				Nationality: "Georgisch",
			},
			"gb": {
				Common:      "Georgia",
				Official:    "Georgia",
				Nationality: "Georgian",
			},
		},
	},
	{
		Alpha2: "GU",
		Alpha3: "GUM",
		Translations: map[string]Translation{
			"de": {
				Common:      "Guam",
				Official:    "Guam",
				Nationality: "Amerikaner",
			},
			"gb": {
				Common:      "Guam",
				Official:    "Guam",
				Nationality: "American",
			},
		},
	},
	{
		Alpha2: "DK",
		Alpha3: "DNK",
		Translations: map[string]Translation{
			"de": {
				Official:    "Königreich Dänemark",
				Common:      "Dänemark",
				Nationality: "Dänisch",
			},
			"gb": {
				Common:      "Denmark",
				Official:    "Kingdom of Denmark",
				Nationality: "Danish",
			},
		},
	},
	{
		Alpha2: "CD",
		Alpha3: "COD",
		Translations: map[string]Translation{
			"de": {
				Common:      "Kongo (Dem. Rep.)",
				Official:    "Demokratische Republik Kongo",
				Nationality: "Kongolesisch",
			},
			"gb": {
				Common:      "DR Congo",
				Official:    "Democratic Republic of the Congo",
				Nationality: "Congolese",
			},
		},
	},
	{
		Alpha2: "DJ",
		Alpha3: "DJI",
		Translations: map[string]Translation{
			"de": {
				Common:      "Dschibuti",
				Official:    "Republik Dschibuti",
				Nationality: "Dschibutisch",
			},
			"gb": {
				Common:      "Djibouti",
				Official:    "Republic of Djibouti",
				Nationality: "Djiboutian",
			},
		},
	},
	{
		Alpha2: "DM",
		Alpha3: "DMA",
		Translations: map[string]Translation{
			"de": {
				Common:      "Dominica",
				Official:    "Commonwealth von Dominica",
				Nationality: "Dominicanisch",
			},
			"gb": {
				Common:      "Dominica",
				Official:    "Commonwealth of Dominica",
				Nationality: "Dominican",
			},
		},
	},
	{
		Alpha2: "DO",
		Alpha3: "DOM",
		Translations: map[string]Translation{
			"de": {
				Official:    "Dominikanische Republik",
				Common:      "Dominikanische Republik",
				Nationality: "Dominikanisch",
			},
			"gb": {
				Common:      "Dominican Republic",
				Official:    "Dominican Republic",
				Nationality: "Dominikaner",
			},
		},
	},
	{
		Alpha2: "EG",
		Alpha3: "EGY",
		Translations: map[string]Translation{
			"de": {
				Common:      "Ägypten",
				Official:    "Arabische Republik Ägypten",
				Nationality: "Ägyptisch",
			},
			"gb": {
				Common:      "Egypt",
				Official:    "Arab Republic of Egypt",
				Nationality: "Egyptian",
			},
		},
	},
	{
		Alpha2: "ZM",
		Alpha3: "ZMB",
		Translations: map[string]Translation{
			"de": {
				Common:      "Sambia",
				Official:    "Republik Sambia",
				Nationality: "Sambisch",
			},
			"gb": {
				Common:      "Zambia",
				Official:    "Republic of Zambia",
				Nationality: "Zambian",
			},
		},
	},
	{
		Alpha2: "EH",
		Alpha3: "ESH",
		Translations: map[string]Translation{
			"de": {
				Common:   "Westsahara",
				Official: "Demokratische Arabische Republik Sahara",
			},
			"gb": {
				Common:   "Western Sahara",
				Official: "Sahrawi Arab Democratic Republic",
			},
		},
	},
	{
		Alpha2: "ZW",
		Alpha3: "ZWE",
		Translations: map[string]Translation{
			"de": {
				Common:      "Simbabwe",
				Official:    "Republik Simbabwe",
				Nationality: "Simbabwisch",
			},
			"gb": {
				Common:      "Zimbabwe",
				Official:    "Republic of Zimbabwe",
				Nationality: "Zimbabwean",
			},
		},
	},
	{
		Alpha2: "IL",
		Alpha3: "ISR",
		Translations: map[string]Translation{
			"de": {
				Common:      "Israel",
				Official:    "Staat Israel",
				Nationality: "Israelisch",
			},
			"gb": {
				Common:      "Israel",
				Official:    "State of Israel",
				Nationality: "Israeli",
			},
		},
	},
	{
		Alpha2: "IN",
		Alpha3: "IND",
		Translations: map[string]Translation{
			"de": {
				Common:      "Indien",
				Official:    "Republik Indien",
				Nationality: "Indisch",
			},
			"gb": {
				Common:      "India",
				Official:    "Republic of India",
				Nationality: "Indian",
			},
		},
	},
	{
		Alpha2: "ID",
		Alpha3: "IDN",
		Translations: map[string]Translation{
			"de": {
				Common:      "Indonesien",          //nolint:misspell // correct german translation
				Official:    "Republik Indonesien", //nolint:misspell // correct german translation
				Nationality: "Indonesisch",
			},
			"gb": {
				Common:   "Indonesia",
				Official: "Republic of Indonesia",
			},
		},
	},
	{
		Alpha2: "JO",
		Alpha3: "JOR",
		Translations: map[string]Translation{
			"de": {
				Official:    "Haschemitisches Königreich Jordanien",
				Common:      "Jordanien",
				Nationality: "Jordanisch",
			},
			"gb": {
				Common:      "Jordan",
				Official:    "Hashemite Kingdom of Jordan",
				Nationality: "Jordanian",
			},
		},
	},
	{
		Alpha2: "IQ",
		Alpha3: "IRQ",
		Translations: map[string]Translation{
			"de": {
				Common:      "Irak",
				Official:    "Republik Irak",
				Nationality: "Irakisch",
			},
			"gb": {
				Common:      "Iraq",
				Official:    "Republic of Iraq",
				Nationality: "Iraqi",
			},
		},
	},
	{
		Alpha2: "IR",
		Alpha3: "IRN",
		Translations: map[string]Translation{
			"de": {
				Common:      "Iran",
				Official:    "Islamische Republik Iran",
				Nationality: "Iranisch",
			},
			"gb": {
				Common:      "Iran",
				Official:    "Islamic Republic of Iran",
				Nationality: "Iranian",
			},
		},
	},
	{
		Alpha2: "IE",
		Alpha3: "IRL",
		Translations: map[string]Translation{
			"de": {
				Official:    "Republik Irland",
				Common:      "Irland",
				Nationality: "Irisch",
			},
			"gb": {
				Common:      "Ireland",
				Official:    "Republic of Ireland",
				Nationality: "Irish",
			},
		},
	},
	{
		Alpha2: "IS",
		Alpha3: "ISL",
		Translations: map[string]Translation{
			"de": {
				Common:      "Island",
				Official:    "Island",
				Nationality: "Isländisch",
			},
			"gb": {
				Common:      "Iceland",
				Official:    "Iceland",
				Nationality: "Icelander",
			},
		},
	},
	{
		Alpha2: "ES",
		Alpha3: "ESP",
		Translations: map[string]Translation{
			"de": {
				Common:      "Spanien",
				Official:    "Königreich Spanien",
				Nationality: "Spanisch", //nolint:misspell // correct german translation
			},
			"gb": {
				Common:      "Spain",
				Official:    "Kingdom of Spain",
				Nationality: "Spanish",
			},
		},
	},
	{
		Alpha2: "IT",
		Alpha3: "ITA",
		Translations: map[string]Translation{
			"de": {
				Official:    "Italienische Republik",
				Common:      "Italien",
				Nationality: "Italienisch",
			},
			"gb": {
				Common:      "Italy",
				Official:    "Italian Republic",
				Nationality: "Italian",
			},
		},
	},
	{
		Alpha2: "YE",
		Alpha3: "YEM",
		Translations: map[string]Translation{
			"de": {
				Common:      "Jemen",
				Official:    "Republik Jemen",
				Nationality: "Jemenitisch",
			},
			"gb": {
				Common:      "Yemen",
				Official:    "Republic of Yemen",
				Nationality: "Yemeni",
			},
		},
	},
	{
		Alpha2: "KZ",
		Alpha3: "KAZ",
		Translations: map[string]Translation{
			"de": {
				Common:      "Kasachstan",
				Official:    "Republik Kasachstan",
				Nationality: "Kasachisch",
			},
			"gb": {
				Common:      "Kazakhstan",
				Official:    "Republic of Kazakhstan",
				Nationality: "Kazakhstani",
			},
		},
	},
	{
		Alpha2: "KY",
		Alpha3: "CYM",
		Translations: map[string]Translation{
			"de": {
				Common:      "Kaimaninseln",
				Official:    "Cayman-Inseln",
				Nationality: "Kaimaninsel-Bewohner",
			},
			"gb": {
				Common:      "Cayman Islands",
				Official:    "Cayman Islands",
				Nationality: "Caymanian",
			},
		},
	},
	{
		Alpha2: "KH",
		Alpha3: "KHM",
		Translations: map[string]Translation{
			"de": {
				Official:    "Königreich Kambodscha",
				Common:      "Kambodscha",
				Nationality: "Kambodschanisch",
			},
			"gb": {
				Common:      "Cambodia",
				Official:    "Kingdom of Cambodia",
				Nationality: "Cambodian",
			},
		},
	},
	{
		Alpha2: "CM",
		Alpha3: "CMR",
		Translations: map[string]Translation{
			"de": {
				Common:      "Kamerun",
				Official:    "Republik Kamerun",
				Nationality: "Kamerunisch",
			},
			"gb": {
				Common:      "Cameroon",
				Official:    "Republic of Cameroon",
				Nationality: "Cameroonian",
			},
		},
	},
	{
		Alpha2: "CA",
		Alpha3: "CAN",
		Translations: map[string]Translation{
			"de": {
				Official:    "Kanada",
				Common:      "Kanada",
				Nationality: "Kanadisch",
			},
			"gb": {
				Common:      "Canada",
				Official:    "Canada",
				Nationality: "Canadian",
			},
		},
	},
	{
		Alpha2: "QA",
		Alpha3: "QAT",
		Translations: map[string]Translation{
			"de": {
				Common:      "Katar",
				Official:    "Staat Katar",
				Nationality: "Katarisch",
			},
			"gb": {
				Common:      "Qatar",
				Official:    "State of Qatar",
				Nationality: "Qatari",
			},
		},
	},
	{
		Alpha2: "KE",
		Alpha3: "KEN",
		Translations: map[string]Translation{
			"de": {
				Common:      "Kenia",
				Official:    "Republik Kenia",
				Nationality: "Kenianisch",
			},
			"gb": {
				Common:      "Kenya",
				Official:    "Republic of Kenya",
				Nationality: "Kenyan",
			},
		},
	},
	{
		Alpha2: "CY",
		Alpha3: "CYP",
		Translations: map[string]Translation{
			"de": {
				Official:    "Republik Zypern",
				Common:      "Zypern",
				Nationality: "Zyprisch",
			},
			"gb": {
				Common:      "Cyprus",
				Official:    "Republic of Cyprus",
				Nationality: "Cypriot",
			},
		},
	},
	{
		Alpha2: "KI",
		Alpha3: "KIR",
		Translations: map[string]Translation{
			"de": {
				Common:      "Kiribati",
				Official:    "Unabhängige und souveräne Republik Kiribati",
				Nationality: "Kiribatisch",
			},
			"gb": {
				Common:      "Kiribati",
				Official:    "Independent and Sovereign Republic of Kiribati",
				Nationality: "Kiribati",
			},
		},
	},
	{
		Alpha2: "CN",
		Alpha3: "CHN",
		Translations: map[string]Translation{
			"de": {
				Official:    "Volksrepublik China",
				Common:      "China",
				Nationality: "Chinesisch",
			},
			"gb": {
				Common:      "China",
				Official:    "People's Republic of China",
				Nationality: "Chinese",
			},
		},
	},
	{
		Alpha2: "CC",
		Alpha3: "CCK",
		Translations: map[string]Translation{
			"de": {
				Common:      "Kokosinseln",
				Official:    "Gebiet der Cocos (Keeling) Islands",
				Nationality: "Australier",
			},
			"gb": {
				Common:      "Cocos (Keeling) Islands",
				Official:    "Territory of the Cocos (Keeling) Islands",
				Nationality: "Australian",
			},
		},
	},
	{
		Alpha2: "CO",
		Alpha3: "COL",
		Translations: map[string]Translation{
			"de": {
				Common:      "Kolumbien",
				Official:    "Republik Kolumbien",
				Nationality: "Kolumbianisch",
			},
			"gb": {
				Common:      "Colombia",
				Official:    "Republic of Colombia",
				Nationality: "Colombian",
			},
		},
	},
	{
		Alpha2: "KM",
		Alpha3: "COM",
		Translations: map[string]Translation{
			"de": {
				Common:   "Union der Komoren",
				Official: "Komorisch",
			},
			"gb": {
				Common:   "Comoros",
				Official: "Union of the Comoros",
			},
		},
	},
	{
		Alpha2: "CG",
		Alpha3: "COG",
		Translations: map[string]Translation{
			"de": {
				Common:      "Kongo",
				Official:    "Republik Kongo",
				Nationality: "Kongolesisch",
			},
			"gb": {
				Common:      "Republic of the Congo",
				Official:    "Republic of the Congo",
				Nationality: "Congolese",
			},
		},
	},
	{
		Alpha2: "KP",
		Alpha3: "PRK",
		Translations: map[string]Translation{
			"de": {
				Common:      "Nordkorea",
				Official:    "Demokratische Volksrepublik Korea",
				Nationality: "Koreanisch",
			},
			"gb": {
				Common:      "North Korea",
				Official:    "Democratic People's Republic of Korea",
				Nationality: "North Korean",
			},
		},
	},
	{
		Alpha2: "KR",
		Alpha3: "KOR",
		Translations: map[string]Translation{
			"de": {
				Common:      "Südkorea",
				Official:    "Republik Korea",
				Nationality: "Koreanisch",
			},
			"gb": {
				Common:      "South Korea",
				Official:    "Republic of Korea",
				Nationality: "Südkoreaner",
			},
		},
	},
	{
		Alpha2: "CR",
		Alpha3: "CRI",
		Translations: map[string]Translation{
			"de": {
				Official:    "Republik Costa Rica",
				Common:      "Costa Rica",
				Nationality: "Costa-ricanisch",
			},
			"gb": {
				Common:      "Costa Rica",
				Official:    "Republic of Costa Rica",
				Nationality: "Costa Rican",
			},
		},
	},
	{
		Alpha2: "CI",
		Alpha3: "CIV",
		Translations: map[string]Translation{
			"de": {
				Common:      "Elfenbeinküste",
				Official:    "Republik Côte d'Ivoire",
				Nationality: "Ivorisch",
			},
			"gb": {
				Common:      "Ivory Coast",
				Official:    "Republic of Côte d'Ivoire",
				Nationality: "Ivorian",
			},
		},
	},
	{
		Alpha2: "CU",
		Alpha3: "CUB",
		Translations: map[string]Translation{
			"de": {
				Common:      "Kuba",
				Official:    "Republik Kuba",
				Nationality: "Kubanisch",
			},
			"gb": {
				Common:      "Cuba",
				Official:    "Republic of Cuba",
				Nationality: "Cuban",
			},
		},
	},
	{
		Alpha2: "KW",
		Alpha3: "KWT",
		Translations: map[string]Translation{
			"de": {
				Common:      "Kuwait",
				Official:    "Staat Kuwait",
				Nationality: "Kuwaitisch",
			},
			"gb": {
				Common:      "Kuwait",
				Official:    "State of Kuwait",
				Nationality: "Kuwaiti",
			},
		},
	},
	{
		Alpha2: "KG",
		Alpha3: "KGZ",
		Translations: map[string]Translation{
			"de": {
				Common:      "Kirgisistan",
				Official:    "Kirgisische Republik",
				Nationality: "Kirgisisch",
			},
			"gb": {
				Common:      "Kyrgyzstan",
				Official:    "Kyrgyz Republic",
				Nationality: "Kyrgyz",
			},
		},
	},
	{
		Alpha2: "LA",
		Alpha3: "LAO",
		Translations: map[string]Translation{
			"de": {
				Common:      "Laos",
				Official:    "Laos, Demokratische Volksrepublik",
				Nationality: "Laotisch",
			},
			"gb": {
				Common:      "Laos",
				Official:    "Lao People's Democratic Republic",
				Nationality: "Laotian",
			},
		},
	},
	{
		Alpha2: "LV",
		Alpha3: "LVA",
		Translations: map[string]Translation{
			"de": {
				Official:    "Republik Lettland",
				Common:      "Lettland",
				Nationality: "Lettisch",
			},
			"gb": {
				Common:      "Latvia",
				Official:    "Republic of Latvia",
				Nationality: "Latvian",
			},
		},
	},
	{
		Alpha2: "LS",
		Alpha3: "LSO",
		Translations: map[string]Translation{
			"de": {
				Common:      "Lesotho",
				Official:    "Königreich Lesotho",
				Nationality: "Lesothisch",
			},
			"gb": {
				Common:      "Lesotho",
				Official:    "Kingdom of Lesotho",
				Nationality: "Mosotho",
			},
		},
	},
	{
		Alpha2: "LR",
		Alpha3: "LBR",
		Translations: map[string]Translation{
			"de": {
				Official:    "Republik Liberia",
				Common:      "Liberia",
				Nationality: "Liberianisch",
			},
			"gb": {
				Common:      "Liberia",
				Official:    "Republic of Liberia",
				Nationality: "Liberian",
			},
		},
	},
	{
		Alpha2: "LB",
		Alpha3: "LBN",
		Translations: map[string]Translation{
			"de": {
				Official:    "Libanesische Republik",
				Common:      "Libanon",
				Nationality: "Libanesisch",
			},
			"gb": {
				Common:      "Lebanon",
				Official:    "Lebanese Republic",
				Nationality: "Lebanese",
			},
		},
	},
	{
		Alpha2: "LY",
		Alpha3: "LBY",
		Translations: map[string]Translation{
			"de": {
				Common:      "Libyen",
				Official:    "Staat Libyen",
				Nationality: "Libysch",
			},
			"gb": {
				Common:      "Libya",
				Official:    "State of Libya",
				Nationality: "Libyan",
			},
		},
	},
	{
		Alpha2: "LT",
		Alpha3: "LTU",
		Translations: map[string]Translation{
			"de": {
				Common:      "Litauen",
				Official:    "Republik Litauen",
				Nationality: "Litauisch",
			},
			"gb": {
				Common:      "Lithuania",
				Official:    "Republic of Lithuania",
				Nationality: "Lithuanian",
			},
		},
	},
	{
		Alpha2: "LI",
		Alpha3: "LIE",
		Translations: map[string]Translation{
			"de": {
				Common:      "Liechtenstein",
				Official:    "Fürstentum Liechtenstein",
				Nationality: "Liechtensteinisch",
			},
			"gb": {
				Common:      "Liechtenstein",
				Official:    "Principality of Liechtenstein",
				Nationality: "Liechtensteiner",
			},
		},
	},
	{
		Alpha2: "LU",
		Alpha3: "LUX",
		Translations: map[string]Translation{
			"de": {
				Common:      "Luxemburg",
				Official:    "Großherzogtum Luxemburg",
				Nationality: "Luxemburgisch",
			},
			"gb": {
				Common:      "Luxembourg",
				Official:    "Grand Duchy of Luxembourg",
				Nationality: "Luxembourger",
			},
		},
	},
	{
		Alpha2: "MU",
		Alpha3: "MUS",
		Translations: map[string]Translation{
			"de": {
				Common:      "Mauritius",
				Official:    "Republik Mauritius",
				Nationality: "Mauritisch",
			},
			"gb": {
				Common:      "Mauritius",
				Official:    "Republic of Mauritius",
				Nationality: "Mauritian",
			},
		},
	},
	{
		Alpha2: "MR",
		Alpha3: "MRT",
		Translations: map[string]Translation{
			"de": {
				Official:    "Islamische Republik Mauretanien",
				Common:      "Mauretanien",
				Nationality: "Mauretanisch",
			},
			"gb": {
				Common:      "Mauritania",
				Official:    "Islamic Republic of Mauritania",
				Nationality: "Mauritanian",
			},
		},
	},
	{
		Alpha2: "MG",
		Alpha3: "MDG",
		Translations: map[string]Translation{
			"de": {
				Common:      "Madagaskar",
				Official:    "Republik Madagascar",
				Nationality: "Madagassisch",
			},
			"gb": {
				Common:      "Madagascar",
				Official:    "Republic of Madagascar",
				Nationality: "Malagasy",
			},
		},
	},
	{
		Alpha2: "YT",
		Alpha3: "MYT",
		Translations: map[string]Translation{
			"de": {
				Official:    "Übersee-Département Mayotte",
				Common:      "Mayotte",
				Nationality: "Franzose",
			},
			"gb": {
				Common:      "Mayotte",
				Official:    "Department of Mayotte",
				Nationality: "French",
			},
		},
	},
	{
		Alpha2: "MO",
		Alpha3: "MAC",
		Translations: map[string]Translation{
			"de": {
				Common:      "Macao",
				Official:    "Sonderverwaltungsregion Macau der Volksrepublik China",
				Nationality: "Macaoer",
			},
			"gb": {
				Common:      "Macau",
				Official:    "Macao Special Administrative Region of the People's Republic of China",
				Nationality: "Macao",
			},
		},
	},
	{
		Alpha2: "MK",
		Alpha3: "MKD",
		Translations: map[string]Translation{
			"de": {
				Common:      "Mazedonien",
				Official:    "Republik Mazedonien",
				Nationality: "Mazedonier",
			},
			"gb": {
				Common:      "Macedonia",
				Official:    "Republic of Macedonia",
				Nationality: "Macedonian",
			},
		},
	},
	{
		Alpha2: "MW",
		Alpha3: "MWI",
		Translations: map[string]Translation{
			"de": {
				Common:      "Malawi",
				Official:    "Republik Malawi",
				Nationality: "Malawisch",
			},
			"gb": {
				Common:      "Malawi",
				Official:    "Republic of Malawi",
				Nationality: "Malawian",
			},
		},
	},
	{
		Alpha2: "MY",
		Alpha3: "MYS",
		Translations: map[string]Translation{
			"de": {
				Official:    "Malaysia",
				Common:      "Malaysia",
				Nationality: "Malaysisch",
			},
			"gb": {
				Common:      "Malaysia",
				Official:    "Malaysia",
				Nationality: "Malaysian",
			},
		},
	},
	{
		Alpha2: "ML",
		Alpha3: "MLI",
		Translations: map[string]Translation{
			"de": {
				Official:    "Republik Mali",
				Common:      "Mali",
				Nationality: "Malisch",
			},
			"gb": {
				Common:      "Mali",
				Official:    "Republic of Mali",
				Nationality: "Malian",
			},
		},
	},
	{
		Alpha2: "MV",
		Alpha3: "MDV",
		Translations: map[string]Translation{
			"de": {
				Common:      "Malediven",
				Official:    "Republik Malediven",
				Nationality: "Maledivisch",
			},
			"gb": {
				Common:      "Maldives",
				Official:    "Republic of the Maldives",
				Nationality: "Maldivian",
			},
		},
	},
	{
		Alpha2: "MT",
		Alpha3: "MLT",
		Translations: map[string]Translation{
			"de": {
				Official:    "Republik Malta",
				Common:      "Malta",
				Nationality: "Maltesisch",
			},
			"gb": {
				Common:      "Malta",
				Official:    "Republic of Malta",
				Nationality: "Maltese",
			},
		},
	},
	{
		Alpha2: "MP",
		Alpha3: "MNP",
		Translations: map[string]Translation{
			"de": {
				Common:      "Nördliche Marianen",
				Official:    "Commonwealth der Nördlichen Marianen",
				Nationality: "Amerikaner",
			},
			"gb": {
				Common:      "Northern Mariana Islands",
				Official:    "Commonwealth of the Northern Mariana Islands",
				Nationality: "American",
			},
		},
	},
	{
		Alpha2: "MA",
		Alpha3: "MAR",
		Translations: map[string]Translation{
			"de": {
				Common:      "Marokko",
				Official:    "Königreich Marokko",
				Nationality: "Marokkanisch",
			},
			"gb": {
				Common:      "Morocco",
				Official:    "Kingdom of Morocco",
				Nationality: "Moroccan",
			},
		},
	},
	{
		Alpha2: "MQ",
		Alpha3: "MTQ",
		Translations: map[string]Translation{
			"de": {
				Common:      "Martinique",
				Official:    "Martinique",
				Nationality: "Franzose",
			},
			"gb": {
				Common:      "Martinique",
				Official:    "Martinique",
				Nationality: "French",
			},
		},
	},
	{
		Alpha2: "MH",
		Alpha3: "MHL",
		Translations: map[string]Translation{
			"de": {
				Common:      "Marshallinseln",
				Official:    "Republik der Marshall-Inseln",
				Nationality: "Marshallisch",
			},
			"gb": {
				Common:      "Marshall Islands",
				Official:    "Republic of the Marshall Islands",
				Nationality: "Marshallese",
			},
		},
	},
	{
		Alpha2: "MX",
		Alpha3: "MEX",
		Translations: map[string]Translation{
			"de": {
				Common:      "Mexiko",
				Official:    "Vereinigte Mexikanische Staaten",
				Nationality: "Mexikanisch",
			},
			"gb": {
				Common:      "Mexico",
				Official:    "United Mexican States",
				Nationality: "Mexican",
			},
		},
	},
	{
		Alpha2: "FM",
		Alpha3: "FSM",
		Translations: map[string]Translation{
			"de": {
				Common:      "Mikronesien",
				Official:    "Föderierte Staaten von Mikronesien",
				Nationality: "Mikronesisch",
			},
			"gb": {
				Common:   "Micronesia",
				Official: "Federated States of Micronesia",
			},
		},
	},
	{
		Alpha2: "MZ",
		Alpha3: "MOZ",
		Translations: map[string]Translation{
			"de": {
				Official:    "Republik Mosambik",
				Common:      "Mosambik",
				Nationality: "Mosambikanisch",
			},
			"gb": {
				Common:      "Mozambique",
				Official:    "Republic of Mozambique",
				Nationality: "Micronesian",
			},
		},
	},
	{
		Alpha2: "MD",
		Alpha3: "MDA",
		Translations: map[string]Translation{
			"de": {
				Common:      "Moldawie",
				Official:    "Republik Moldau",
				Nationality: "Moldauisch",
			},
			"gb": {
				Common:      "Moldova",
				Official:    "Republic of Moldova",
				Nationality: "Moldovan",
			},
		},
	},
	{
		Alpha2: "MC",
		Alpha3: "MCO",
		Translations: map[string]Translation{
			"de": {
				Common:      "Monaco",
				Official:    "Fürstentum Monaco",
				Nationality: "Monegassisch",
			},
			"gb": {
				Common:      "Monaco",
				Official:    "Principality of Monaco",
				Nationality: "Monacan",
			},
		},
	},
	{
		Alpha2: "MN",
		Alpha3: "MNG",
		Translations: map[string]Translation{
			"de": {
				Official:    "Mongolei",
				Common:      "Mongolei",
				Nationality: "Mongolisch",
			},
			"gb": {
				Common:      "Mongolia",
				Official:    "Mongolia",
				Nationality: "Mongolian",
			},
		},
	},
	{
		Alpha2: "MS",
		Alpha3: "MSR",
		Translations: map[string]Translation{
			"de": {
				Official:    "Montserrat",
				Common:      "Montserrat",
				Nationality: "Montserrat-Bewohner",
			},
			"gb": {
				Common:      "Montserrat",
				Official:    "Montserrat",
				Nationality: "Montserratian",
			},
		},
	},
	{
		Alpha2: "MM",
		Alpha3: "MMR",
		Translations: map[string]Translation{
			"de": {
				Common:      "Myanmar",
				Official:    "Republik der Union von Myanmar",
				Nationality: "Myanmarisch",
			},
			"gb": {
				Common:      "Myanmar",
				Official:    "Republic of the Union of Myanmar",
				Nationality: "Myanmar (Burmese)",
			},
		},
	},
	{
		Alpha2: "NA",
		Alpha3: "NAM",
		Translations: map[string]Translation{
			"de": {
				Common:      "Namibia",
				Official:    "Republik Namibia",
				Nationality: "Namibisch",
			},
			"gb": {
				Common:      "Namibia",
				Official:    "Republic of Namibia",
				Nationality: "Namibian",
			},
		},
	},
	{
		Alpha2: "NR",
		Alpha3: "NRU",
		Translations: map[string]Translation{
			"de": {
				Common:      "Nauru",
				Official:    "Republik Nauru",
				Nationality: "Nauruisch",
			},
			"gb": {
				Common:      "Nauru",
				Official:    "Republic of Nauru",
				Nationality: "Nauruan",
			},
		},
	},
	{
		Alpha2: "NP",
		Alpha3: "NPL",
		Translations: map[string]Translation{
			"de": {
				Common:      "Népal",
				Official:    "Demokratischen Bundesrepublik Nepal",
				Nationality: "Nepalesisch",
			},
			"gb": {
				Common:      "Nepal",
				Official:    "Federal Democratic Republic of Nepal",
				Nationality: "Nepalese",
			},
		},
	},
	{
		Alpha2: "NE",
		Alpha3: "NER",
		Translations: map[string]Translation{
			"de": {
				Official:    "Republik Niger",
				Common:      "Niger",
				Nationality: "Nigrisch",
			},
			"gb": {
				Common:      "Niger",
				Official:    "Republic of Niger",
				Nationality: "Nigerien",
			},
		},
	},
	{
		Alpha2: "NG",
		Alpha3: "NGA",
		Translations: map[string]Translation{
			"de": {
				Common:      "Nigeria",
				Official:    "Bundesrepublik Nigeria",
				Nationality: "Nigerianisch",
			},
			"gb": {
				Common:      "Nigeria",
				Official:    "Federal Republic of Nigeria",
				Nationality: "Nigerian",
			},
		},
	},
	{
		Alpha2: "NL",
		Alpha3: "NLD",
		Translations: map[string]Translation{
			"de": {
				Official:    "Niederlande",
				Common:      "Niederlande",
				Nationality: "Niederländisch",
			},
			"gb": {
				Common:      "Netherlands",
				Official:    "Netherlands",
				Nationality: "Dutch",
			},
		},
	},
	{
		Alpha2: "NI",
		Alpha3: "NIC",
		Translations: map[string]Translation{
			"de": {
				Official:    "Republik Nicaragua",
				Common:      "Nicaragua",
				Nationality: "Nicaraguanisch",
			},
			"gb": {
				Common:      "Nicaragua",
				Official:    "Republic of Nicaragua",
				Nationality: "Nicaraguan",
			},
		},
	},
	{
		Alpha2: "NU",
		Alpha3: "NIU",
		Translations: map[string]Translation{
			"de": {
				Common:      "Niue",
				Official:    "Niue",
				Nationality: "Niueanisch",
			},
			"gb": {
				Common:      "Niue",
				Official:    "Niue",
				Nationality: "Niueaner",
			},
		},
	},
	{
		Alpha2: "NZ",
		Alpha3: "NZL",
		Translations: map[string]Translation{
			"de": {
				Common:      "Neuseeland",
				Official:    "Neuseeland",
				Nationality: "Neuseeländisch",
			},
			"gb": {
				Common:      "New Zealand",
				Official:    "New Zealand",
				Nationality: "New Zealander",
			},
		},
	},
	{
		Alpha2: "NC",
		Alpha3: "NCL",
		Translations: map[string]Translation{
			"de": {
				Official:    "Neukaledonien",
				Common:      "Neukaledonien",
				Nationality: "Franzose",
			},
			"gb": {
				Common:      "New Caledonia",
				Official:    "New Caledonia",
				Nationality: "French",
			},
		},
	},
	{
		Alpha2: "NO",
		Alpha3: "NOR",
		Translations: map[string]Translation{
			"de": {
				Official:    "Königreich Norwegen",
				Common:      "Norwegen",
				Nationality: "Norwegisch",
			},
			"gb": {
				Common:      "Norway",
				Official:    "Kingdom of Norway",
				Nationality: "Norwegian",
			},
		},
	},
	{
		Alpha2: "OM",
		Alpha3: "OMN",
		Translations: map[string]Translation{
			"de": {
				Common:      "Oman",
				Official:    "Sultanat Oman",
				Nationality: "Omanisch",
			},
			"gb": {
				Common:      "Oman",
				Official:    "Sultanate of Oman",
				Nationality: "Omani",
			},
		},
	},
	{
		Alpha2: "BV",
		Alpha3: "BVT",
		Translations: map[string]Translation{
			"de": {
				Official: "Bouvet-Insel",
				Common:   "Bouvetinsel",
			},
			"gb": {
				Common:   "Bouvet Island",
				Official: "Bouvet Island",
			},
		},
	},
	{
		Alpha2: "IM",
		Alpha3: "IMN",
		Translations: map[string]Translation{
			"de": {
				Official:    "Isle of Man",
				Common:      "Insel Man",
				Nationality: "Manxer",
			},
			"gb": {
				Common:      "Isle of Man",
				Official:    "Isle of Man",
				Nationality: "Manx",
			},
		},
	},
	{
		Alpha2: "NF",
		Alpha3: "NFK",
		Translations: map[string]Translation{
			"de": {
				Common:      "Norfolkinsel",
				Official:    "Gebiet der Norfolk-Insel",
				Nationality: "Manx",
			},
			"gb": {
				Common:      "Norfolk Island",
				Official:    "Territory of Norfolk Island",
				Nationality: "Manx",
			},
		},
	},
	{
		Alpha2: "PN",
		Alpha3: "PCN",
		Translations: map[string]Translation{
			"de": {
				Official:    "Pitcairn Inselgruppe",
				Common:      "Pitcairn",
				Nationality: "Pitcairn-Inselbewohner",
			},
			"gb": {
				Common:      "Pitcairn Islands",
				Official:    "Pitcairn Group of Islands",
				Nationality: "Pitcairn Islander",
			},
		},
	},
	{
		Alpha2: "CX",
		Alpha3: "CXR",
		Translations: map[string]Translation{
			"de": {
				Official:    "Gebiet der Weihnachtsinsel",
				Common:      "Weihnachtsinsel",
				Nationality: "Weihnachtsinselbewohner",
			},
			"gb": {
				Common:      "Christmas Island",
				Official:    "Territory of Christmas Island",
				Nationality: "Christmas Islander",
			},
		},
	},
	{
		Alpha2: "SH",
		Alpha3: "SHN",
		Translations: map[string]Translation{
			"de": {
				Common:      "St. Helena",
				Official:    "St. Helena, Ascension und Tristan da Cunha",
				Nationality: "St. Helenianer",
			},
			"gb": {
				Common:      "Saint Helena",
				Official:    "Saint Helena, Ascension and Tristan da Cunha",
				Nationality: "Saint Helenian",
			},
		},
	},
	{
		Alpha2: "WF",
		Alpha3: "WLF",
		Translations: map[string]Translation{
			"de": {
				Common:      "Wallis und Futuna",
				Official:    "Gebiet der Wallis und Futuna",
				Nationality: "Wallis und Futuna Inselbewohner",
			},
			"gb": {
				Common:      "Wallis and Futuna",
				Official:    "Territory of the Wallis and Futuna Islands",
				Nationality: "Wallisian and Futuna",
			},
		},
	},
	{
		Alpha2: "HM",
		Alpha3: "HMD",
		Translations: map[string]Translation{
			"de": {
				Official: "Heard und McDonaldinseln",
				Common:   "Heard und die McDonaldinseln",
			},
			"gb": {
				Common:   "Heard Island and McDonald Islands",
				Official: "Heard Island and McDonald Islands",
			},
		},
	},
	{
		Alpha2: "CV",
		Alpha3: "CPV",
		Translations: map[string]Translation{
			"de": {
				Official:    "Republik Cabo Verde",
				Common:      "Kap Verde",
				Nationality: "Kap-Verdier",
			},
			"gb": {
				Common:      "Cape Verde",
				Official:    "Republic of Cabo Verde",
				Nationality: "Cape Verdean",
			},
		},
	},
	{
		Alpha2: "CK",
		Alpha3: "COK",
		Translations: map[string]Translation{
			"de": {
				Common:      "Cookinseln",
				Official:    "Cook-Inseln",
				Nationality: "Cook-Inselbewohner",
			},
			"gb": {
				Common:      "Cook Islands",
				Official:    "Cook Islands",
				Nationality: "Cook Islander",
			},
		},
	},
	{
		Alpha2: "WS",
		Alpha3: "WSM",
		Translations: map[string]Translation{
			"de": {
				Common:      "Samoa",
				Official:    "Unabhängige Staat Samoa",
				Nationality: "Samoanisch",
			},
			"gb": {
				Common:      "Samoa",
				Official:    "Independent State of Samoa",
				Nationality: "Samoan",
			},
		},
	},
	{
		Alpha2: "SJ",
		Alpha3: "SJM",
		Translations: map[string]Translation{
			"de": {
				Common:      "Spitzbergen",
				Official:    "Inselgruppe Spitzbergen",
				Nationality: "Norwegian",
			},
			"gb": {
				Common:      "Svalbard and Jan Mayen",
				Official:    "Svalbard og Jan Mayen",
				Nationality: "Norweger",
			},
		},
	},
	{
		Alpha2: "TC",
		Alpha3: "TCA",
		Translations: map[string]Translation{
			"de": {
				Common:      "Turks-und Caicosinseln",
				Official:    "Turks und Caicos Inseln",
				Nationality: "Inselbewohner von Turks- und Caicosinseln",
			},
			"gb": {
				Common:      "Turks and Caicos Islands",
				Official:    "Turks and Caicos Islands",
				Nationality: "Turks and Caicos Islander",
			},
		},
	},
	{
		Alpha2: "UM",
		Alpha3: "UMI",
		Translations: map[string]Translation{
			"de": {
				Official: "USA, kleinere ausgelagerte Inseln",
				Common:   "Kleinere Inselbesitzungen der Vereinigten Staaten",
			},
			"gb": {
				Common:   "United States Minor Outlying Islands",
				Official: "United States Minor Outlying Islands",
			},
		},
	},
	{
		Alpha2: "PK",
		Alpha3: "PAK",
		Translations: map[string]Translation{
			"de": {
				Common:      "Pakistan",
				Official:    "Islamische Republik Pakistan",
				Nationality: "Pakistanisch",
			},
			"gb": {
				Common:      "Pakistan",
				Official:    "Islamic Republic of Pakistan",
				Nationality: "Pakistani",
			},
		},
	},
	{
		Alpha2: "PW",
		Alpha3: "PLW",
		Translations: map[string]Translation{
			"de": {
				Official:    "Palau",
				Common:      "Palau",
				Nationality: "Palauisch",
			},
			"gb": {
				Common:      "Palau",
				Official:    "Republic of Palau",
				Nationality: "Palauan",
			},
		},
	},
	{
		Alpha2: "PS",
		Alpha3: "PSE",
		Translations: map[string]Translation{
			"de": {
				Common:      "Palästina",
				Official:    "Staat Palästina",
				Nationality: "Palästinenser",
			},
			"gb": {
				Common:      "Palestine",
				Official:    "State of Palestine",
				Nationality: "Palestinian",
			},
		},
	},
	{
		Alpha2: "PA",
		Alpha3: "PAN",
		Translations: map[string]Translation{
			"de": {
				Official:    "Republik Panama",
				Common:      "Panama",
				Nationality: "Panamaisch",
			},
			"gb": {
				Common:      "Panama",
				Official:    "Republic of Panama",
				Nationality: "Panamanian",
			},
		},
	},
	{
		Alpha2: "PG",
		Alpha3: "PNG",
		Translations: map[string]Translation{
			"de": {
				Common:      "Papua-Neuguinea",
				Official:    "Unabhängige Staat Papua-Neuguinea",
				Nationality: "Papua-neuguineisch",
			},
			"gb": {
				Common:      "Papua New Guinea",
				Official:    "Independent State of Papua New Guinea",
				Nationality: "Papua New Guinean",
			},
		},
	},
	{
		Alpha2: "PY",
		Alpha3: "PRY",
		Translations: map[string]Translation{
			"de": {
				Common:      "Paraguay",
				Official:    "Republik Paraguay",
				Nationality: "Paraguayisch",
			},
			"gb": {
				Common:      "Paraguay",
				Official:    "Republic of Paraguay",
				Nationality: "Paraguayan",
			},
		},
	},
	{
		Alpha2: "PE",
		Alpha3: "PER",
		Translations: map[string]Translation{
			"de": {
				Official:    "Republik Peru",
				Common:      "Peru",
				Nationality: "Peruanisch",
			},
			"gb": {
				Common:      "Peru",
				Official:    "Republic of Peru",
				Nationality: "Peruvian",
			},
		},
	},
	{
		Alpha2: "PL",
		Alpha3: "POL",
		Translations: map[string]Translation{
			"de": {
				Common:      "Polen",
				Official:    "Republik Polen",
				Nationality: "Polnisch",
			},
			"gb": {
				Common:      "Poland",
				Official:    "Republic of Poland",
				Nationality: "Polish",
			},
		},
	},
	{
		Alpha2: "PT",
		Alpha3: "PRT",
		Translations: map[string]Translation{
			"de": {
				Common:      "Portugal",
				Official:    "Portugiesische Republik",
				Nationality: "Portugiesisch",
			},
			"gb": {
				Common:      "Portugal",
				Official:    "Portuguese Republic",
				Nationality: "Portuguese",
			},
		},
	},
	{
		Alpha2: "PR",
		Alpha3: "PRI",
		Translations: map[string]Translation{
			"de": {
				Common:      "Puerto Rico",
				Official:    "Commonwealth von Puerto Rico",
				Nationality: "Puerto-Ricaner",
			},
			"gb": {
				Common:      "Puerto Rico",
				Official:    "Commonwealth of Puerto Rico",
				Nationality: "Puerto Rican",
			},
		},
	},
	{
		Alpha2: "RE",
		Alpha3: "REU",
		Translations: map[string]Translation{
			"de": {
				Official:    "Réunion",
				Common:      "Réunion",
				Nationality: "Franzose",
			},
			"gb": {
				Common:      "Réunion",
				Official:    "Réunion Island",
				Nationality: "French",
			},
		},
	},
	{
		Alpha2: "RU",
		Alpha3: "RUS",
		Translations: map[string]Translation{
			"de": {
				Common:      "Russland",
				Official:    "Russische Föderation",
				Nationality: "Russisch",
			},
			"gb": {
				Common:      "Russia",
				Official:    "Russian Federation",
				Nationality: "Russian",
			},
		},
	},
	{
		Alpha2: "RW",
		Alpha3: "RWA",
		Translations: map[string]Translation{
			"de": {
				Official:    "Republik Ruanda",
				Common:      "Ruanda",
				Nationality: "Ruandisch",
			},
			"gb": {
				Common:      "Rwanda",
				Official:    "Republic of Rwanda",
				Nationality: "Rwandan",
			},
		},
	},
	{
		Alpha2: "RO",
		Alpha3: "ROU",
		Translations: map[string]Translation{
			"de": {
				Official:    "Rumänien",
				Common:      "Rumänien",
				Nationality: "Rumänisch",
			},
			"gb": {
				Common:      "Romania",
				Official:    "Romania",
				Nationality: "Romanian",
			},
		},
	},
	{
		Alpha2: "SV",
		Alpha3: "SLV",
		Translations: map[string]Translation{
			"de": {
				Common:      "El Salvador",
				Official:    "Republik El Salvador",
				Nationality: "Salvadorianisch",
			},
			"gb": {
				Common:      "El Salvador",
				Official:    "Republic of El Salvador",
				Nationality: "Salvadoran",
			},
		},
	},
	{
		Alpha2: "SM",
		Alpha3: "SMR",
		Translations: map[string]Translation{
			"de": {
				Common:      "San Marino",
				Official:    "Republik San Marino",
				Nationality: "San-marinesisch",
			},
			"gb": {
				Common:      "San Marino",
				Official:    "Most Serene Republic of San Marino",
				Nationality: "San Marinese",
			},
		},
	},
	{
		Alpha2: "ST",
		Alpha3: "STP",
		Translations: map[string]Translation{
			"de": {
				Common:      "São Tomé und Príncipe",
				Official:    "Demokratische Republik São Tomé und Príncipe",
				Nationality: "sro-tomйisch",
			},
			"gb": {
				Common:      "São Tomé and Príncipe",
				Official:    "Democratic Republic of São Tomé and Príncipe",
				Nationality: "Sao Tomean",
			},
		},
	},
	{
		Alpha2: "SA",
		Alpha3: "SAU",
		Translations: map[string]Translation{
			"de": {
				Official:    "Königreich Saudi-Arabien",
				Common:      "Saudi-Arabien",
				Nationality: "Saudi-arabisch",
			},
			"gb": {
				Common:      "Saudi Arabia",
				Official:    "Kingdom of Saudi Arabia",
				Nationality: "Saudi Arabian",
			},
		},
	},
	{
		Alpha2: "SZ",
		Alpha3: "SWZ",
		Translations: map[string]Translation{
			"de": {
				Common:      "Swasiland",
				Official:    "Königreich Swasiland",
				Nationality: "Swasi",
			},
			"gb": {
				Common:      "Swaziland",
				Official:    "Kingdom of Swaziland",
				Nationality: "Swasi",
			},
		},
	},
	{
		Alpha2: "SC",
		Alpha3: "SYC",
		Translations: map[string]Translation{
			"de": {
				Common:      "Seychellen",
				Official:    "Republik der Seychellen",
				Nationality: "Seychellisch",
			},
			"gb": {
				Common:      "Seychelles",
				Official:    "Republic of Seychelles",
				Nationality: "Seychellois",
			},
		},
	},
	{
		Alpha2: "SN",
		Alpha3: "SEN",
		Translations: map[string]Translation{
			"de": {
				Common:      "Senegal",
				Official:    "Republik Senegal",
				Nationality: "Senegalesisch",
			},
			"gb": {
				Common:      "Senegal",
				Official:    "Republic of Senegal",
				Nationality: "Senegalese",
			},
		},
	},
	{
		Alpha2: "PM",
		Alpha3: "SPM",
		Translations: map[string]Translation{
			"de": {
				Common:      "Saint-Pierre und Miquelon",
				Official:    "St. Pierre und Miquelon",
				Nationality: "Saint-Pierrais und Miquelonnais",
			},
			"gb": {
				Common:      "Saint Pierre and Miquelon",
				Official:    "Saint Pierre and Miquelon",
				Nationality: "Saint-Pierrais and Miquelonnais",
			},
		},
	},
	{
		Alpha2: "VC",
		Alpha3: "VCT",
		Translations: map[string]Translation{
			"de": {
				Official:    "St. Vincent und die Grenadinen",
				Common:      "Saint Vincent und die Grenadinen",
				Nationality: "Vincentisch",
			},
			"gb": {
				Common:      "Saint Vincent and the Grenadines",
				Official:    "Saint Vincent and the Grenadines",
				Nationality: "Vincentian",
			},
		},
	},
	{
		Alpha2: "KN",
		Alpha3: "KNA",
		Translations: map[string]Translation{
			"de": {
				Common:      "Saint Christopher und Nevis",
				Official:    "Föderation von Saint Kitts und Nevisa",
				Nationality: "Kittitianer und Nevisianer",
			},
			"gb": {
				Common:      "Saint Kitts and Nevis",
				Official:    "Federation of Saint Christopher and Nevisa",
				Nationality: "Kittitian and Nevisian",
			},
		},
	},
	{
		Alpha2: "LC",
		Alpha3: "LCA",
		Translations: map[string]Translation{
			"de": {
				Common:      "Saint Lucia",
				Official:    "St. Lucia",
				Nationality: "Lucianisch",
			},
			"gb": {
				Common:      "Saint Lucia",
				Official:    "Saint Lucia",
				Nationality: "Saint Lucian",
			},
		},
	},
	{
		Alpha2: "SG",
		Alpha3: "SGP",
		Translations: map[string]Translation{
			"de": {
				Official:    "Republik Singapur",
				Common:      "Singapur",
				Nationality: "Singapurisch",
			},
			"gb": {
				Common:      "Singapore",
				Official:    "Republic of Singapore",
				Nationality: "Singaporean",
			},
		},
	},
	{
		Alpha2: "SY",
		Alpha3: "SYR",
		Translations: map[string]Translation{
			"de": {
				Common:      "Syrien",
				Official:    "Arabische Republik Syrien",
				Nationality: "Syrisch",
			},
			"gb": {
				Common:      "Syria",
				Official:    "Syrian Arab Republic",
				Nationality: "Syrian",
			},
		},
	},
	{
		Alpha2: "SK",
		Alpha3: "SVK",
		Translations: map[string]Translation{
			"de": {
				Common:      "Slowakei",
				Official:    "Slowakische Republik",
				Nationality: "Slowakisch",
			},
			"gb": {
				Common:      "Slovakia",
				Official:    "Slovak Republic",
				Nationality: "Slovak",
			},
		},
	},
	{
		Alpha2: "SI",
		Alpha3: "SVN",
		Translations: map[string]Translation{
			"de": {
				Common:      "Slowenien",
				Official:    "Republik Slowenien",
				Nationality: "Slowenisch",
			},
			"gb": {
				Common:      "Slovenia",
				Official:    "Republic of Slovenia",
				Nationality: "Slovenian",
			},
		},
	},
	{
		Alpha2: "US",
		Alpha3: "USA",
		Translations: map[string]Translation{
			"de": {
				Official:    "Vereinigte Staaten von Amerika",
				Common:      "Vereinigte Staaten von Amerika",
				Nationality: "Amerikanisch",
			},
			"gb": {
				Common:      "United States",
				Official:    "United States of America",
				Nationality: "American",
			},
		},
	},
	{
		Alpha2: "SB",
		Alpha3: "SLB",
		Translations: map[string]Translation{
			"de": {
				Common:      "Salomonen",
				Official:    "Salomon-Inseln",
				Nationality: "Salomonisch",
			},
			"gb": {
				Common:      "Solomon Islands",
				Official:    "Solomon Islands",
				Nationality: "Solomon Islander",
			},
		},
	},
	{
		Alpha2: "SO",
		Alpha3: "SOM",
		Translations: map[string]Translation{
			"de": {
				Official:    "Bundesrepublik Somalia",
				Common:      "Somalia",
				Nationality: "Somalisch",
			},
			"gb": {
				Common:      "Somalia",
				Official:    "Federal Republic of Somalia",
				Nationality: "Somali",
			},
		},
	},
	{
		Alpha2: "SD",
		Alpha3: "SDN",
		Translations: map[string]Translation{
			"de": {
				Official:    "Republik Sudan",
				Common:      "Sudan",
				Nationality: "Sudanesisch",
			},
			"gb": {
				Common:      "Sudan",
				Official:    "Republic of the Sudan",
				Nationality: "Sudanese",
			},
		},
	},
	{
		Alpha2: "SR",
		Alpha3: "SUR",
		Translations: map[string]Translation{
			"de": {
				Official:    "Republik Suriname",
				Common:      "Suriname",
				Nationality: "Surinamisch",
			},
			"gb": {
				Common:      "Suriname",
				Official:    "Republic of Suriname",
				Nationality: "Surinamese",
			},
		},
	},
	{
		Alpha2: "SL",
		Alpha3: "SLE",
		Translations: map[string]Translation{
			"de": {
				Common:      "Sierra Leone",
				Official:    "Republik Sierra Leone",
				Nationality: "Sierra-leonisch",
			},
			"gb": {
				Common:      "Sierra Leone",
				Official:    "Republic of Sierra Leone",
				Nationality: "Sierra Leonean",
			},
		},
	},
	{
		Alpha2: "TJ",
		Alpha3: "TJK",
		Translations: map[string]Translation{
			"de": {
				Official:    "Republik Tadschikistan",
				Common:      "Tadschikistan",
				Nationality: "Tadschikisch",
			},
			"gb": {
				Common:      "Tajikistan",
				Official:    "Republic of Tajikistan",
				Nationality: "Tajik",
			},
		},
	},
	{
		Alpha2: "TW",
		Alpha3: "TWN",
		Translations: map[string]Translation{
			"de": {
				Common:      "Taiwan",
				Official:    "Republik China (Taiwan)",
				Nationality: "Taiwaner",
			},
			"gb": {
				Common:      "Taiwan",
				Official:    "Republic of China (Taiwan)",
				Nationality: "Taiwanese",
			},
		},
	},
	{
		Alpha2: "TH",
		Alpha3: "THA",
		Translations: map[string]Translation{
			"de": {
				Official:    "Königreich Thailand",
				Common:      "Thailand",
				Nationality: "Thailändisch",
			},
			"gb": {
				Common:      "Thailand",
				Official:    "Kingdom of Thailand",
				Nationality: "Thai",
			},
		},
	},
	{
		Alpha2: "TZ",
		Alpha3: "TZA",
		Translations: map[string]Translation{
			"de": {
				Official:    "Vereinigte Republik Tansania",
				Common:      "Tansania",
				Nationality: "Tansanisch",
			},
			"gb": {
				Common:      "Tanzania",
				Official:    "United Republic of Tanzania",
				Nationality: "Tanzanian",
			},
		},
	},
	{
		Alpha2: "TG",
		Alpha3: "TGO",
		Translations: map[string]Translation{
			"de": {
				Official:    "Republik Togo",
				Common:      "Togo",
				Nationality: "Togoisch",
			},
			"gb": {
				Common:      "Togo",
				Official:    "Togolese Republic",
				Nationality: "Togolese",
			},
		},
	},
	{
		Alpha2: "TK",
		Alpha3: "TKL",
		Translations: map[string]Translation{
			"de": {
				Official:    "Tokelau",
				Common:      "Tokelau",
				Nationality: "Tokelauer",
			},
			"gb": {
				Common:      "Tokelau",
				Official:    "Tokelau",
				Nationality: "Tokelauan",
			},
		},
	},
	{
		Alpha2: "TO",
		Alpha3: "TON",
		Translations: map[string]Translation{
			"de": {
				Common:      "Tonga",
				Official:    "Königreich Tonga",
				Nationality: "Tongaisch",
			},
			"gb": {
				Common:      "Tonga",
				Official:    "Kingdom of Tonga",
				Nationality: "Tongan",
			},
		},
	},
	{
		Alpha2: "TT",
		Alpha3: "TTO",
		Translations: map[string]Translation{
			"de": {
				Common:      "Trinidad und Tobago",
				Official:    "Republik Trinidad und Tobago",
				Nationality: "Trinidaner und Tobager",
			},
			"gb": {
				Common:      "Trinidad and Tobago",
				Official:    "Republic of Trinidad and Tobago",
				Nationality: "Trinidadian and Tobagonian",
			},
		},
	},
	{
		Alpha2: "TV",
		Alpha3: "TUV",
		Translations: map[string]Translation{
			"de": {
				Common:      "Tuvalu",
				Official:    "Tuvalu",
				Nationality: "Tuvaluisch",
			},
			"gb": {
				Common:      "Tuvalu",
				Official:    "Tuvalu",
				Nationality: "Tuvaluan",
			},
		},
	},
	{
		Alpha2: "TN",
		Alpha3: "TUN",
		Translations: map[string]Translation{
			"de": {
				Common:      "Tunesien",
				Official:    "Tunesische Republik",
				Nationality: "Tunesisch",
			},
			"gb": {
				Common:      "Tunisia",
				Official:    "Tunisian Republic",
				Nationality: "Tunisian",
			},
		},
	},
	{
		Alpha2: "TM",
		Alpha3: "TKM",
		Translations: map[string]Translation{
			"de": {
				Common:      "Turkmenistan",
				Official:    "Turkmenistan",
				Nationality: "Turkmenisch",
			},
			"gb": {
				Common:      "Turkmenistan",
				Official:    "Turkmenistan",
				Nationality: "Turkmen",
			},
		},
	},
	{
		Alpha2: "TR",
		Alpha3: "TUR",
		Translations: map[string]Translation{
			"de": {
				Common:      "Türkei",
				Official:    "Republik Türkei",
				Nationality: "Türkisch",
			},
			"gb": {
				Common:      "Turkey",
				Official:    "Republic of Turkey",
				Nationality: "Turkish",
			},
		},
	},
	{
		Alpha2: "UG",
		Alpha3: "UGA",
		Translations: map[string]Translation{
			"de": {
				Common:      "Uganda",
				Official:    "Republik Uganda",
				Nationality: "Ugandisch",
			},
			"gb": {
				Common:      "Uganda",
				Official:    "Republic of Uganda",
				Nationality: "Ugandan",
			},
		},
	},
	{
		Alpha2: "UZ",
		Alpha3: "UZB",
		Translations: map[string]Translation{
			"de": {
				Common:      "Usbekistan",
				Official:    "Republik Usbekistan",
				Nationality: "Usbekisch",
			},
			"gb": {
				Common:      "Uzbekistan",
				Official:    "Republic of Uzbekistan",
				Nationality: "Uzbek",
			},
		},
	},
	{
		Alpha2: "UA",
		Alpha3: "UKR",
		Translations: map[string]Translation{
			"de": {
				Official:    "Ukraine",
				Common:      "Ukraine",
				Nationality: "Ukrainisch",
			},
			"gb": {
				Common:      "Ukraine",
				Official:    "Ukraine",
				Nationality: "Ukrainian",
			},
		},
	},
	{
		Alpha2: "UY",
		Alpha3: "URY",
		Translations: map[string]Translation{
			"de": {
				Official:    "Republik Östlich des Uruguay",
				Common:      "Uruguay",
				Nationality: "Uruguayisch",
			},
			"gb": {
				Common:      "Uruguay",
				Official:    "Oriental Republic of Uruguay",
				Nationality: "Uruguayan",
			},
		},
	},
	{
		Alpha2: "FO",
		Alpha3: "FRO",
		Translations: map[string]Translation{
			"de": {
				Common:      "Färöer-Inseln",
				Official:    "Färöer",
				Nationality: "Färinger",
			},
			"gb": {
				Common:      "Faroe Islands",
				Official:    "Faroe Islands",
				Nationality: "Faroe Islander",
			},
		},
	},
	{
		Alpha2: "FJ",
		Alpha3: "FJI",
		Translations: map[string]Translation{
			"de": {
				Common:      "Fidschi",
				Official:    "Republik Fidschi",
				Nationality: "Fidschianisch",
			},
			"gb": {
				Common:      "Fiji",
				Official:    "Republic of Fiji",
				Nationality: "Fijian",
			},
		},
	},
	{
		Alpha2: "PH",
		Alpha3: "PHL",
		Translations: map[string]Translation{
			"de": {
				Common:      "Philippinen",
				Official:    "Republik der Philippinen",
				Nationality: "Philippinisch",
			},
			"gb": {
				Common:      "Philippines",
				Official:    "Republic of the Philippines",
				Nationality: "Filipino",
			},
		},
	},
	{
		Alpha2: "FI",
		Alpha3: "FIN",
		Translations: map[string]Translation{
			"de": {
				Official:    "Republik Finnland",
				Common:      "Finnland",
				Nationality: "Finnisch", //nolint:misspell // correct german translation
			},
			"gb": {
				Common:      "Finland",
				Official:    "Republic of Finland",
				Nationality: "Finnish",
			},
		},
	},
	{
		Alpha2: "FK",
		Alpha3: "FLK",
		Translations: map[string]Translation{
			"de": {
				Common:      "Falklandinseln",
				Official:    "Falkland-Inseln",
				Nationality: "Finnish",
			},
			"gb": {
				Common:      "Falkland Islands",
				Official:    "Falkland Islands",
				Nationality: "Finnish",
			},
		},
	},
	{
		Alpha2: "FR",
		Alpha3: "FRA",
		Translations: map[string]Translation{
			"de": {
				Common:      "Frankreich",
				Official:    "Französische Republik",
				Nationality: "Französisch",
			},
			"gb": {
				Common:      "France",
				Official:    "French Republic",
				Nationality: "French",
			},
		},
	},
	{
		Alpha2: "GF",
		Alpha3: "GUF",
		Translations: map[string]Translation{
			"de": {
				Common:      "Französisch Guyana",
				Official:    "Guayana",
				Nationality: "Französisch-Guianisch",
			},
			"gb": {
				Common:      "French Guiana",
				Official:    "Guiana",
				Nationality: "French Guianese",
			},
		},
	},
	{
		Alpha2: "PF",
		Alpha3: "PYF",
		Translations: map[string]Translation{
			"de": {
				Common:      "Französisch-Polynesien",
				Official:    "Französisch-Polynesien",
				Nationality: "Französisch-Polynesien",
			},
			"gb": {
				Common:      "French Polynesia",
				Official:    "French Polynesia",
				Nationality: "French Polynesian",
			},
		},
	},
	{
		Alpha2: "TF",
		Alpha3: "ATF",
		Translations: map[string]Translation{
			"de": {
				Common:   "Französische Süd-und Antarktisgebiete",
				Official: "Gebiet der Französisch Süd-und Antarktisgebiete",
			},
			"gb": {
				Common:   "French Southern and Antarctic Lands",
				Official: "Territory of the French Southern and Antarctic Lands",
			},
		},
	},
	{
		Alpha2: "HR",
		Alpha3: "HRV",
		Translations: map[string]Translation{
			"de": {
				Official:    "Republik Kroatien",
				Common:      "Kroatien",
				Nationality: "Kroatisch",
			},
			"gb": {
				Common:      "Croatia",
				Official:    "Republic of Croatia",
				Nationality: "Croatia",
			},
		},
	},
	{
		Alpha2: "CF",
		Alpha3: "CAF",
		Translations: map[string]Translation{
			"de": {
				Common:      "Zentralafrikanische Republik",
				Official:    "Zentralafrikanische Republik",
				Nationality: "Zentralafrikanisch",
			},
			"gb": {
				Common:      "Central African Republic",
				Official:    "Central African Republic",
				Nationality: "Central African Republic",
			},
		},
	},
	{
		Alpha2: "TD",
		Alpha3: "TCD",
		Translations: map[string]Translation{
			"de": {
				Common:      "Tschad",
				Official:    "Republik Tschad",
				Nationality: "Tschadisch",
			},
			"gb": {
				Common:      "Chad",
				Official:    "Republic of Chad",
				Nationality: "Chad",
			},
		},
	},
	{
		Alpha2: "CZ",
		Alpha3: "CZE",
		Translations: map[string]Translation{
			"de": {
				Official:    "Tschechische Republik",
				Common:      "Tschechische Republik",
				Nationality: "Tschechisch",
			},
			"gb": {
				Common:      "Czech Republic",
				Official:    "Czech Republic",
				Nationality: "Czech",
			},
		},
	},
	{
		Alpha2: "CL",
		Alpha3: "CHL",
		Translations: map[string]Translation{
			"de": {
				Common:      "Chile",
				Official:    "Republik Chile",
				Nationality: "Chilenisch",
			},
			"gb": {
				Common:      "Chile",
				Official:    "Republic of Chile",
				Nationality: "Chilean",
			},
		},
	},
	{
		Alpha2: "CH",
		Alpha3: "CHE",
		Translations: map[string]Translation{
			"de": {
				Official:    "Schweizerische Eidgenossenschaft",
				Common:      "Schweiz",
				Nationality: "Schweizerisch",
			},
			"gb": {
				Common:      "Switzerland",
				Official:    "Swiss Confederation",
				Nationality: "Swiss",
			},
		},
	},
	{
		Alpha2: "SE",
		Alpha3: "SWE",
		Translations: map[string]Translation{
			"de": {
				Common:      "Schweden",
				Official:    "Königreich Schweden",
				Nationality: "Schwedisch",
			},
			"gb": {
				Common:      "Sweden",
				Official:    "Kingdom of Sweden",
				Nationality: "Swedish",
			},
		},
	},
	{
		Alpha2: "LK",
		Alpha3: "LKA",
		Translations: map[string]Translation{
			"de": {
				Official:    "Demokratische Sozialistische Republik Sri Lanka",
				Common:      "Sri Lanka",
				Nationality: "Sri-lankisch",
			},
			"gb": {
				Common:      "Sri Lanka",
				Official:    "Democratic Socialist Republic of Sri Lanka",
				Nationality: "Sri Lankan",
			},
		},
	},
	{
		Alpha2: "EC",
		Alpha3: "ECU",
		Translations: map[string]Translation{
			"de": {
				Official:    "Republik Ecuador",
				Common:      "Ecuador",
				Nationality: "Ecuadorianisch",
			},
			"gb": {
				Common:      "Ecuador",
				Official:    "Republic of Ecuador",
				Nationality: "Ecuadorian",
			},
		},
	},
	{
		Alpha2: "GQ",
		Alpha3: "GNQ",
		Translations: map[string]Translation{
			"de": {
				Common:      "Äquatorialguinea",
				Official:    "Republik Äquatorialguinea",
				Nationality: "Äquatorial-guineisch",
			},
			"gb": {
				Common:      "Equatorial Guinea",
				Official:    "Republic of Equatorial Guinea",
				Nationality: "Equatorial Guinean",
			},
		},
	},
	{
		Alpha2: "ER",
		Alpha3: "ERI",
		Translations: map[string]Translation{
			"de": {
				Common:      "Eritrea",
				Official:    "Staat Eritrea",
				Nationality: "Eritreisch",
			},
			"gb": {
				Common:      "Eritrea",
				Official:    "State of Eritrea",
				Nationality: "Eritrean",
			},
		},
	},
	{
		Alpha2: "EE",
		Alpha3: "EST",
		Translations: map[string]Translation{
			"de": {
				Common:      "Estland",
				Official:    "Republik Estland",
				Nationality: "Estnisch",
			},
			"gb": {
				Common:      "Estonia",
				Official:    "Republic of Estonia",
				Nationality: "Estonian",
			},
		},
	},
	{
		Alpha2: "ET",
		Alpha3: "ETH",
		Translations: map[string]Translation{
			"de": {
				Official:    "Demokratische Bundesrepublik Äthiopien",
				Common:      "Äthiopien",
				Nationality: "Äthiopisch",
			},
			"gb": {
				Common:      "Ethiopia",
				Official:    "Federal Democratic Republic of Ethiopia",
				Nationality: "Ethiopian",
			},
		},
	},
	{
		Alpha2: "ZA",
		Alpha3: "ZAF",
		Translations: map[string]Translation{
			"de": {
				Common:      "Republic Südafrika",
				Official:    "Republik Südafrika",
				Nationality: "Südafrikanisch",
			},
			"gb": {
				Common:      "South Africa",
				Official:    "Republic of South Africa",
				Nationality: "South African",
			},
		},
	},
	{
		Alpha2: "GS",
		Alpha3: "SGS",
		Translations: map[string]Translation{
			"de": {
				Common:      "Südgeorgien und die Südlichen Sandwichinseln",
				Official:    "Südgeorgien und die Südlichen Sandwichinseln",
				Nationality: "Südgeorgisch und Südliche Sandwichinseln",
			},
			"gb": {
				Common:      "South Georgia",
				Official:    "South Georgia and the South Sandwich Islands",
				Nationality: "South Georgia and the South Sandwich Islands",
			},
		},
	},
	{
		Alpha2: "JM",
		Alpha3: "JAM",
		Translations: map[string]Translation{
			"de": {
				Common:      "Jamaika",
				Official:    "Jamaika",
				Nationality: "Jamaikanisch",
			},
			"gb": {
				Common:      "Jamaica",
				Official:    "Jamaica",
				Nationality: "Jamaican",
			},
		},
	},
	{
		Alpha2: "ME",
		Alpha3: "MNE",
		Translations: map[string]Translation{
			"de": {
				Official:    "Montenegro",
				Common:      "Montenegro",
				Nationality: "Montenegrinisch",
			},
			"gb": {
				Common:      "Montenegro",
				Official:    "Montenegro",
				Nationality: "Montenegrin",
			},
		},
	},
	{
		Alpha2: "BL",
		Alpha3: "BLM",
		Translations: map[string]Translation{
			"de": {
				Official:    "Gebietskörperschaft Saint -Barthélemy",
				Common:      "Saint-Barthélemy",
				Nationality: "Saint-Barthélemois",
			},
			"gb": {
				Common:      "Saint Barthélemy",
				Official:    "Collectivity of Saint Barthélemy",
				Nationality: "Saint-Barthélemois",
			},
		},
	},
	{
		Alpha2: "SX",
		Alpha3: "SXM",
		Translations: map[string]Translation{
			"de": {
				Common:      "Sint Maarten",
				Official:    "Sint Maarten",
				Nationality: "Sint Maartener",
			},
			"gb": {
				Common:      "Sint Maarten",
				Official:    "Sint Maarten",
				Nationality: "Sint Maartener",
			},
		},
	},
	{
		Alpha2: "RS",
		Alpha3: "SRB",
		Translations: map[string]Translation{
			"de": {
				Common:      "Serbien",
				Official:    "Republik Serbien",
				Nationality: "Serbisch",
			},
			"gb": {
				Common:      "Serbia",
				Official:    "Republic of Serbia",
				Nationality: "Serbian",
			},
		},
	},
	{
		Alpha2: "AX",
		Alpha3: "ALA",
		Translations: map[string]Translation{
			"de": {
				Common:      "Åland",
				Official:    "Åland-Inseln",
				Nationality: "Ålandisch",
			},
			"gb": {
				Common:      "Åland Islands",
				Official:    "Åland Islands",
				Nationality: "Ålandish",
			},
		},
	},
	{
		Alpha2: "BQ",
		Alpha3: "BES",
		Translations: map[string]Translation{
			"de": {
				Common:      "Karibische Niederlande",
				Official:    "Bonaire, Sint Eustatius und Saba",
				Nationality: "Bonaire, Sint Eustatius und Sabaner",
			},
			"gb": {
				Common:      "Caribbean Netherlands",
				Official:    "Bonaire, Sint Eustatius and Saba",
				Nationality: "Bonaire, Sint Eustatius and Saban",
			},
		},
	},
	{
		Alpha2: "GG",
		Alpha3: "GGY",
		Translations: map[string]Translation{
			"de": {
				Official:    "Guernsey",
				Common:      "Guernsey",
				Nationality: "Guernseyer",
			},
			"gb": {
				Common:      "Guernsey",
				Official:    "Bailiwick of Guernsey",
				Nationality: "Guernsey",
			},
		},
	},
	{
		Alpha2: "JE",
		Alpha3: "JEY",
		Translations: map[string]Translation{
			"de": {
				Official:    "Vogtei Jersey",
				Common:      "Jersey",
				Nationality: "Jerseyer",
			},
			"gb": {
				Common:      "Jersey",
				Official:    "Bailiwick of Jersey",
				Nationality: "Jersey",
			},
		},
	},
	{
		Alpha2: "CW",
		Alpha3: "CUW",
		Translations: map[string]Translation{
			"de": {
				Common:      "Curaçao",
				Official:    "Curaçao",
				Nationality: "Curaçaoer",
			},
			"gb": {
				Common:      "Curaçao",
				Official:    "Country of Curaçao",
				Nationality: "Curaçaoan",
			},
		},
	},
	{
		Alpha2: "MF",
		Alpha3: "MAF",
		Translations: map[string]Translation{
			"de": {
				Common:      "Saint Martin",
				Official:    "St. Martin",
				Nationality: "Saint-Martiner",
			},
			"gb": {
				Common:      "Saint Martin",
				Official:    "Saint Martin",
				Nationality: "Saint-Martinois",
			},
		},
	},
	{
		Alpha2: "SS",
		Alpha3: "SSD",
		Translations: map[string]Translation{
			"de": {
				Official:    "Republik Südsudan",
				Common:      "Südsudan",
				Nationality: "Südsudanesisch",
			},
			"gb": {
				Common:      "South Sudan",
				Official:    "Republic of South Sudan",
				Nationality: "South Sudanese",
			},
		},
	},
	{
		Alpha2: "JP",
		Alpha3: "JPN",
		Translations: map[string]Translation{
			"de": {
				Common:      "Japan",
				Official:    "Japan",
				Nationality: "Japanisch",
			},
			"gb": {
				Common:      "Japan",
				Official:    "Japan",
				Nationality: "Japanese",
			},
		},
	},
}
