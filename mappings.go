package countries

//nolint:gochecknoglobals // this is intended for performance reasons.
var mappings = []Mapping{
	{
		Alpha2: "AU",
		Alpha3: "AUS",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Australien",              //nolint:misspell // correct german translation
				Official:    "Commonwealth Australien", //nolint:misspell // correct german translation
				Nationality: "Australisch",
			},
			"en": {
				Common:      "Australia",
				Official:    "Commonwealth of Australia",
				Nationality: "Australian",
			},
		},
	},
	{
		Alpha2: "AT",
		Alpha3: "AUT",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Österreich",
				Official:    "Republik Österreich",
				Nationality: "Österreichisch",
			},
			"en": {
				Common:      "Austria",
				Official:    "Republic of Austria",
				Nationality: "Austrian",
			},
		},
	},
	{
		Alpha2: "AZ",
		Alpha3: "AZE",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Aserbaidschan",
				Official:    "Republik Aserbaidschan",
				Nationality: "Aserbaidschaner",
			},
			"en": {
				Common:      "Azerbaijan",
				Official:    "Republic of Azerbaijan",
				Nationality: "Azerbaijani",
			},
		},
	},
	{
		Alpha2: "AL",
		Alpha3: "ALB",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Albanien",
				Official:    "Republik Albanien",
				Nationality: "Albanisch",
			},
			"en": {
				Common:      "Albania",
				Official:    "Republic of Albania",
				Nationality: "Albanian",
			},
		},
	},
	{
		Alpha2: "DZ",
		Alpha3: "DZA",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Algerien",
				Official:    "Demokratische Volksrepublik Algerien",
				Nationality: "Algerisch",
			},
			"en": {
				Common:      "Algeria",
				Official:    "People's Democratic Republic of Algeria",
				Nationality: "Algerian",
			},
		},
	},
	{
		Alpha2: "AS",
		Alpha3: "ASM",
		Translations: map[Language]Translation{
			"de": {
				Official:    "Amerikanisch-Samoa",
				Common:      "Amerikanisch-Samoa",
				Nationality: "Samoaner aus Amerikanisch-Samoa",
			},
			"en": {
				Common:      "American Samoa",
				Official:    "American Samoa",
				Nationality: "American Samoan",
			},
		},
	},
	{
		Alpha2: "AI",
		Alpha3: "AIA",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Anguilla",
				Official:    "Anguilla",
				Nationality: "Anguillianer",
			},
			"en": {
				Common:      "Anguilla",
				Official:    "Anguilla",
				Nationality: "Anguillan",
			},
		},
	},
	{
		Alpha2: "AO",
		Alpha3: "AGO",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Angola",
				Official:    "Republik Angola",
				Nationality: "Angolanisch",
			},
			"en": {
				Common:      "Angola",
				Official:    "Republic of Angola",
				Nationality: "Angolan",
			},
		},
	},
	{
		Alpha2: "AD",
		Alpha3: "AND",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Andorra",
				Official:    "Fürstentum Andorra",
				Nationality: "Andorranisch",
			},
			"en": {
				Common:      "Andorra",
				Official:    "Principality of Andorra",
				Nationality: "Andorran",
			},
		},
	},
	{
		Alpha2: "AQ",
		Alpha3: "ATA",
		Translations: map[Language]Translation{
			"de": {
				Common:   "Antarktis",
				Official: "Antarktika",
			},
			"en": {
				Common:   "Antarctica",
				Official: "Antarctica",
			},
		},
	},
	{
		Alpha2: "AG",
		Alpha3: "ATG",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Antigua und Barbuda",
				Official:    "Antigua und Barbuda",
				Nationality: "Antiguanisch",
			},
			"en": {
				Common:      "Antigua and Barbuda",
				Official:    "Antigua and Barbuda",
				Nationality: "Antiguans and Barbudans",
			},
		},
	},
	{
		Alpha2: "AE",
		Alpha3: "ARE",
		Translations: map[Language]Translation{
			"de": {
				Official:    "Vereinigte Arabische Emirate",
				Common:      "Vereinigte Arabische Emirate",
				Nationality: "Emirati",
			},
			"en": {
				Common:      "United Arab Emirates",
				Official:    "United Arab Emirates",
				Nationality: "Emirati",
			},
		},
	},
	{
		Alpha2: "AR",
		Alpha3: "ARG",
		Translations: map[Language]Translation{
			"de": {
				Official:    "Argentinische Republik",
				Common:      "Argentinien",
				Nationality: "Argentinisch",
			},
			"en": {
				Common:      "Argentina",
				Official:    "Argentine Republic",
				Nationality: "Argentine",
			},
		},
	},
	{
		Alpha2: "AM",
		Alpha3: "ARM",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Armenien",
				Official:    "Republik Armenien",
				Nationality: "Armenisch",
			},
			"en": {
				Common:      "Armenia",
				Official:    "Republic of Armenia",
				Nationality: "Armenian",
			},
		},
	},
	{
		Alpha2: "AW",
		Alpha3: "ABW",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Aruba",
				Official:    "Aruba",
				Nationality: "Aruber",
			},
			"en": {
				Common:      "Aruba",
				Official:    "Aruba",
				Nationality: "Aruban",
			},
		},
	},
	{
		Alpha2: "AF",
		Alpha3: "AFG",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Afghanistan",
				Official:    "Islamische Republik Afghanistan",
				Nationality: "Afghanisch",
			},
			"en": {
				Common:      "Afghanistan",
				Official:    "Islamic Republic of Afghanistan",
				Nationality: "Afghan",
			},
		},
	},
	{
		Alpha2: "BS",
		Alpha3: "BHS",
		Translations: map[Language]Translation{
			"de": {
				Official:    "Commonwealth der Bahamas",
				Common:      "Bahamas",
				Nationality: "Bahamaisch",
			},
			"en": {
				Common:      "Bahamas",
				Official:    "Commonwealth of the Bahamas",
				Nationality: "Bahamian",
			},
		},
	},
	{
		Alpha2: "BD",
		Alpha3: "BGD",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Bangladesch",               //nolint:misspell // correct german translation
				Official:    "Volksrepublik Bangladesch", //nolint:misspell // correct german translation
				Nationality: "Bangladeschisch",
			},
			"en": {
				Common:      "Bangladesh",
				Official:    "People's Republic of Bangladesh",
				Nationality: "Bangladeshi",
			},
		},
	},
	{
		Alpha2: "BB",
		Alpha3: "BRB",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Barbados",
				Official:    "Barbados",
				Nationality: "Barbadisch",
			},
			"en": {
				Common:      "Barbados",
				Official:    "Barbados",
				Nationality: "Barbadian",
			},
		},
	},
	{
		Alpha2: "BH",
		Alpha3: "BHR",
		Translations: map[Language]Translation{
			"de": {
				Official:    "Königreich Bahrain",
				Common:      "Bahrain",
				Nationality: "Bahrainisch",
			},
			"en": {
				Common:      "Bahrain",
				Official:    "Kingdom of Bahrain",
				Nationality: "Bahraini",
			},
		},
	},
	{
		Alpha2: "BY",
		Alpha3: "BLR",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Weißrussland",
				Official:    "Republik Belarus",
				Nationality: "Belarussisch",
			},
			"en": {
				Common:      "Belarus",
				Official:    "Republic of Belarus",
				Nationality: "Belarusian",
			},
		},
	},
	{
		Alpha2: "BZ",
		Alpha3: "BLZ",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Belize",
				Official:    "Belize",
				Nationality: "Belizisch",
			},
			"en": {
				Common:      "Belize",
				Official:    "Belize",
				Nationality: "Belizean",
			},
		},
	},
	{
		Alpha2: "BE",
		Alpha3: "BEL",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Belgien",
				Official:    "Königreich Belgien",
				Nationality: "Belgisch",
			},
			"en": {
				Common:      "Belgium",
				Official:    "Kingdom of Belgium",
				Nationality: "Belgian",
			},
		},
	},
	{
		Alpha2: "BJ",
		Alpha3: "BEN",
		Translations: map[Language]Translation{
			"de": {
				Official:    "Republik Benin",
				Common:      "Benin",
				Nationality: "Beninisch",
			},
			"en": {
				Common:      "Benin",
				Official:    "Republic of Benin",
				Nationality: "Beninese",
			},
		},
	},
	{
		Alpha2: "BM",
		Alpha3: "BMU",
		Translations: map[Language]Translation{
			"de": {
				Official:    "Bermuda",
				Common:      "Bermuda",
				Nationality: "Bermuder",
			},
			"en": {
				Common:      "Bermuda",
				Official:    "Bermuda",
				Nationality: "Bermudian",
			},
		},
	},
	{
		Alpha2: "BG",
		Alpha3: "BGR",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Bulgarien",
				Official:    "Republik Bulgarien",
				Nationality: "Bulgarisch",
			},
			"en": {
				Common:      "Bulgaria",
				Official:    "Republic of Bulgaria",
				Nationality: "Bulgarian",
			},
		},
	},
	{
		Alpha2: "BO",
		Alpha3: "BOL",
		Translations: map[Language]Translation{
			"de": {
				Official:    "Multinationaler Staat von Bolivien",
				Common:      "Bolivien",
				Nationality: "Bolivianisch",
			},
			"en": {
				Common:      "Bolivia",
				Official:    "Plurinational State of Bolivia",
				Nationality: "Bolivian",
			},
		},
	},
	{
		Alpha2: "BA",
		Alpha3: "BIH",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Bosnien und Herzegowina",
				Official:    "Bosnien und Herzegowina",
				Nationality: "Bosnisch-herzegowinisch",
			},
			"en": {
				Common:      "Bosnia and Herzegovina",
				Official:    "Bosnia and Herzegovina",
				Nationality: "Bosnian and Herzegovinian",
			},
		},
	},
	{
		Alpha2: "BW",
		Alpha3: "BWA",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Botswana",
				Official:    "Republik Botswana",
				Nationality: "Botsuanisch",
			},
			"en": {
				Common:      "Botswana",
				Official:    "Republic of Botswana",
				Nationality: "Botswanan",
			},
		},
	},
	{
		Alpha2: "BR",
		Alpha3: "BRA",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Brasilien",
				Official:    "Föderative Republik Brasilien",
				Nationality: "Brasilianisch",
			},
			"en": {
				Common:      "Brazil",
				Official:    "Federative Republic of Brazil",
				Nationality: "Brazilian",
			},
		},
	},
	{
		Alpha2: "IO",
		Alpha3: "IOT",
		Translations: map[Language]Translation{
			"de": {
				Common:   "Britisches Territorium im Indischen Ozean",
				Official: "Britisch-Indischer Ozean",
			},
			"en": {
				Common:   "British Indian Ocean Territory",
				Official: "British Indian Ocean Territory",
			},
		},
	},
	{
		Alpha2: "BN",
		Alpha3: "BRN",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Brunei",
				Official:    "Nation von Brunei, Wohnung des Friedens",
				Nationality: "Bruneiisch",
			},
			"en": {
				Common:      "Brunei",
				Official:    "Nation of Brunei, Abode of Peace",
				Nationality: "Bruneian",
			},
		},
	},
	{
		Alpha2: "BF",
		Alpha3: "BFA",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Burkina Faso",
				Official:    "Burkina Faso",
				Nationality: "Burkinisch",
			},
			"en": {
				Common:      "Burkina Faso",
				Official:    "Burkina Faso",
				Nationality: "Burkinabe",
			},
		},
	},
	{
		Alpha2: "BI",
		Alpha3: "BDI",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Burundi",
				Official:    "Republik Burundi",
				Nationality: "Burundisch",
			},
			"en": {
				Common:      "Burundi",
				Official:    "Republic of Burundi",
				Nationality: "Burundian",
			},
		},
	},
	{
		Alpha2: "BT",
		Alpha3: "BTN",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Bhutan",
				Official:    "Königreich Bhutan",
				Nationality: "Bhutanisch",
			},
			"en": {
				Common:      "Bhutan",
				Official:    "Kingdom of Bhutan",
				Nationality: "Bhutanese",
			},
		},
	},
	{
		Alpha2: "VU",
		Alpha3: "VUT",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Vanuatu",
				Official:    "Vanuatu",
				Nationality: "Vanuatuisch",
			},
			"en": {
				Common:      "Vanuatu",
				Official:    "Republic of Vanuatu",
				Nationality: "Ni-Vanuatu",
			},
		},
	},
	{
		Alpha2: "VA",
		Alpha3: "VAT",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Vatikanstadt",
				Official:    "Staat Vatikanstadt",
				Nationality: "Vatikanisch",
			},
			"en": {
				Common:      "Vatican City",
				Official:    "Vatican City State",
				Nationality: "Vatican",
			},
		},
	},
	{
		Alpha2: "GB",
		Alpha3: "GBR",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Vereinigtes Königreich",
				Official:    "Vereinigtes Königreich Großbritannien und Nordirland",
				Nationality: "Britisch",
			},
			"en": {
				Common:      "United Kingdom",
				Official:    "United Kingdom of Great Britain and Northern Ireland",
				Nationality: "British",
			},
		},
	},
	{
		Alpha2: "HU",
		Alpha3: "HUN",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Ungarn",
				Official:    "Ungarn",
				Nationality: "Ungarisch",
			},
			"en": {
				Common:      "Hungary",
				Official:    "Hungary",
				Nationality: "Hungarian",
			},
		},
	},
	{
		Alpha2: "VE",
		Alpha3: "VEN",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Venezuela",
				Official:    "Bolivarische Republik Venezuela",
				Nationality: "Venezolanisch",
			},
			"en": {
				Common:      "Venezuela",
				Official:    "Bolivarian Republic of Venezuela",
				Nationality: "Venezuelan",
			},
		},
	},
	{
		Alpha2: "VG",
		Alpha3: "VGB",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Britische Jungferninseln",
				Official:    "Jungferninseln",
				Nationality: "Britisch-Virgin-Islander",
			},
			"en": {
				Common:      "British Virgin Islands",
				Official:    "Virgin Islands",
				Nationality: "British Virgin Islander",
			},
		},
	},
	{
		Alpha2: "VI",
		Alpha3: "VIR",
		Translations: map[Language]Translation{
			"de": {
				Official:    "Jungferninseln der Vereinigten Staaten",
				Common:      "Amerikanische Jungferninseln",
				Nationality: "Ameriakner",
			},
			"en": {
				Common:      "United States Virgin Islands",
				Official:    "Virgin Islands of the United States",
				Nationality: "American",
			},
		},
	},
	{
		Alpha2: "TL",
		Alpha3: "TLS",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Timor-Leste",
				Official:    "Demokratische Republik Timor-Leste",
				Nationality: "Osttimorese",
			},
			"en": {
				Common:      "Timor-Leste",
				Official:    "Democratic Republic of Timor-Leste",
				Nationality: "East Timorese",
			},
		},
	},
	{
		Alpha2: "VN",
		Alpha3: "VNM",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Vietnam",
				Official:    "Sozialistische Republik Vietnam",
				Nationality: "Vietnamesisch",
			},
			"en": {
				Common:      "Vietnam",
				Official:    "Socialist Republic of Vietnam",
				Nationality: "Vietnamese",
			},
		},
	},
	{
		Alpha2: "GA",
		Alpha3: "GAB",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Gabun",
				Official:    "Gabunische Republik",
				Nationality: "Gabunisch",
			},
			"en": {
				Common:      "Gabon",
				Official:    "Gabonese Republic",
				Nationality: "Gabonese",
			},
		},
	},
	{
		Alpha2: "HT",
		Alpha3: "HTI",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Haiti",
				Official:    "Republik Haiti",
				Nationality: "Haitianisch",
			},
			"en": {
				Common:      "Haiti",
				Official:    "Republic of Haiti",
				Nationality: "Haitian",
			},
		},
	},
	{
		Alpha2: "GY",
		Alpha3: "GUY",
		Translations: map[Language]Translation{
			"de": {
				Official:    "Guayana",
				Common:      "Französisch Guyana",
				Nationality: "Guyanisch",
			},
			"en": {
				Common:      "Guyana",
				Official:    "Co-operative Republic of Guyana",
				Nationality: "Guyanese",
			},
		},
	},
	{
		Alpha2: "GM",
		Alpha3: "GMB",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Gambia",
				Official:    "Republik Gambia",
				Nationality: "Gambisch",
			},
			"en": {
				Common:      "Gambia",
				Official:    "Republic of the Gambia",
				Nationality: "Gambian",
			},
		},
	},
	{
		Alpha2: "GH",
		Alpha3: "GHA",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Ghana",
				Official:    "Republik Ghana",
				Nationality: "Ghanaisch",
			},
			"en": {
				Common:      "Ghana",
				Official:    "Republic of Ghana",
				Nationality: "Ghanaian",
			},
		},
	},
	{
		Alpha2: "GP",
		Alpha3: "GLP",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Guadeloupe",
				Official:    "Guadeloupe",
				Nationality: "Franzose",
			},
			"en": {
				Common:      "Guadeloupe",
				Official:    "Guadeloupe",
				Nationality: "French",
			},
		},
	},
	{
		Alpha2: "GT",
		Alpha3: "GTM",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Guatemala",
				Official:    "Republik Guatemala",
				Nationality: "Guatemaltekisch",
			},
			"en": {
				Common:      "Guatemala",
				Official:    "Republic of Guatemala",
				Nationality: "Guatemalan",
			},
		},
	},
	{
		Alpha2: "GN",
		Alpha3: "GIN",
		Translations: map[Language]Translation{
			"de": {
				Official:    "Republik Guinea",
				Common:      "Guinea",
				Nationality: "Guineisch",
			},
			"en": {
				Common:      "Guinea",
				Official:    "Republic of Guinea",
				Nationality: "Guinean",
			},
		},
	},
	{
		Alpha2: "GW",
		Alpha3: "GNB",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Guinea-Bissau",
				Official:    "Republik Guinea-Bissau",
				Nationality: "Guinea-bissauisch",
			},
			"en": {
				Common:      "Guinea-Bissau",
				Official:    "Republic of Guinea-Bissau",
				Nationality: "Guinea-Bissauan",
			},
		},
	},
	{
		Alpha2: "DE",
		Alpha3: "DEU",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Deutschland",
				Official:    "Bundesrepublik Deutschland",
				Nationality: "Deutsch",
			},
			"en": {
				Common:      "Germany",
				Official:    "Federal Republic of Germany",
				Nationality: "German",
			},
		},
	},
	{
		Alpha2: "GI",
		Alpha3: "GIB",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Gibraltar",
				Official:    "Gibraltar",
				Nationality: "Gibraltarian",
			},
			"en": {
				Common:      "Gibraltar",
				Official:    "Gibraltar",
				Nationality: "Gibraltarer",
			},
		},
	},
	{
		Alpha2: "HN",
		Alpha3: "HND",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Honduras",
				Official:    "Republik Honduras",
				Nationality: "Honduranisch",
			},
			"en": {
				Common:      "Honduras",
				Official:    "Republic of Honduras",
				Nationality: "Honduran",
			},
		},
	},
	{
		Alpha2: "HK",
		Alpha3: "HKG",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Hongkong",
				Official:    "Sonderverwaltungszone der Volksrepublik China",
				Nationality: "Hongkonger",
			},
			"en": {
				Common:      "Hong Kong",
				Official:    "Hong Kong Special Administrative Region of the People's Republic of China",
				Nationality: "Hong Konger",
			},
		},
	},
	{
		Alpha2: "GD",
		Alpha3: "GRD",
		Translations: map[Language]Translation{
			"de": {
				Official:    "Grenada",
				Common:      "Grenada",
				Nationality: "Grenadisch",
			},
			"en": {
				Common:      "Grenada",
				Official:    "Grenada",
				Nationality: "Grenadian",
			},
		},
	},
	{
		Alpha2: "GL",
		Alpha3: "GRL",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Grönland",
				Official:    "Grönland",
				Nationality: "Grönländer",
			},
			"en": {
				Common:      "Greenland",
				Official:    "Greenland",
				Nationality: "Greenlander",
			},
		},
	},
	{
		Alpha2: "GR",
		Alpha3: "GRC",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Griechenland",
				Official:    "Hellenische Republik",
				Nationality: "Griechisch",
			},
			"en": {
				Common:      "Greece",
				Official:    "Hellenic Republic",
				Nationality: "Greek",
			},
		},
	},
	{
		Alpha2: "GE",
		Alpha3: "GEO",
		Translations: map[Language]Translation{
			"de": {
				Official:    "Georgia",
				Common:      "Georgien",
				Nationality: "Georgisch",
			},
			"en": {
				Common:      "Georgia",
				Official:    "Georgia",
				Nationality: "Georgian",
			},
		},
	},
	{
		Alpha2: "GU",
		Alpha3: "GUM",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Guam",
				Official:    "Guam",
				Nationality: "Amerikaner",
			},
			"en": {
				Common:      "Guam",
				Official:    "Guam",
				Nationality: "American",
			},
		},
	},
	{
		Alpha2: "DK",
		Alpha3: "DNK",
		Translations: map[Language]Translation{
			"de": {
				Official:    "Königreich Dänemark",
				Common:      "Dänemark",
				Nationality: "Dänisch",
			},
			"en": {
				Common:      "Denmark",
				Official:    "Kingdom of Denmark",
				Nationality: "Danish",
			},
		},
	},
	{
		Alpha2: "CD",
		Alpha3: "COD",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Kongo (Dem. Rep.)",
				Official:    "Demokratische Republik Kongo",
				Nationality: "Kongolesisch",
			},
			"en": {
				Common:      "DR Congo",
				Official:    "Democratic Republic of the Congo",
				Nationality: "Congolese",
			},
		},
	},
	{
		Alpha2: "DJ",
		Alpha3: "DJI",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Dschibuti",
				Official:    "Republik Dschibuti",
				Nationality: "Dschibutisch",
			},
			"en": {
				Common:      "Djibouti",
				Official:    "Republic of Djibouti",
				Nationality: "Djiboutian",
			},
		},
	},
	{
		Alpha2: "DM",
		Alpha3: "DMA",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Dominica",
				Official:    "Commonwealth von Dominica",
				Nationality: "Dominicanisch",
			},
			"en": {
				Common:      "Dominica",
				Official:    "Commonwealth of Dominica",
				Nationality: "Dominican",
			},
		},
	},
	{
		Alpha2: "DO",
		Alpha3: "DOM",
		Translations: map[Language]Translation{
			"de": {
				Official:    "Dominikanische Republik",
				Common:      "Dominikanische Republik",
				Nationality: "Dominikanisch",
			},
			"en": {
				Common:      "Dominican Republic",
				Official:    "Dominican Republic",
				Nationality: "Dominikaner",
			},
		},
	},
	{
		Alpha2: "EG",
		Alpha3: "EGY",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Ägypten",
				Official:    "Arabische Republik Ägypten",
				Nationality: "Ägyptisch",
			},
			"en": {
				Common:      "Egypt",
				Official:    "Arab Republic of Egypt",
				Nationality: "Egyptian",
			},
		},
	},
	{
		Alpha2: "ZM",
		Alpha3: "ZMB",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Sambia",
				Official:    "Republik Sambia",
				Nationality: "Sambisch",
			},
			"en": {
				Common:      "Zambia",
				Official:    "Republic of Zambia",
				Nationality: "Zambian",
			},
		},
	},
	{
		Alpha2: "EH",
		Alpha3: "ESH",
		Translations: map[Language]Translation{
			"de": {
				Common:   "Westsahara",
				Official: "Demokratische Arabische Republik Sahara",
			},
			"en": {
				Common:   "Western Sahara",
				Official: "Sahrawi Arab Democratic Republic",
			},
		},
	},
	{
		Alpha2: "ZW",
		Alpha3: "ZWE",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Simbabwe",
				Official:    "Republik Simbabwe",
				Nationality: "Simbabwisch",
			},
			"en": {
				Common:      "Zimbabwe",
				Official:    "Republic of Zimbabwe",
				Nationality: "Zimbabwean",
			},
		},
	},
	{
		Alpha2: "IL",
		Alpha3: "ISR",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Israel",
				Official:    "Staat Israel",
				Nationality: "Israelisch",
			},
			"en": {
				Common:      "Israel",
				Official:    "State of Israel",
				Nationality: "Israeli",
			},
		},
	},
	{
		Alpha2: "IN",
		Alpha3: "IND",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Indien",
				Official:    "Republik Indien",
				Nationality: "Indisch",
			},
			"en": {
				Common:      "India",
				Official:    "Republic of India",
				Nationality: "Indian",
			},
		},
	},
	{
		Alpha2: "ID",
		Alpha3: "IDN",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Indonesien",          //nolint:misspell // correct german translation
				Official:    "Republik Indonesien", //nolint:misspell // correct german translation
				Nationality: "Indonesisch",
			},
			"en": {
				Common:   "Indonesia",
				Official: "Republic of Indonesia",
			},
		},
	},
	{
		Alpha2: "JO",
		Alpha3: "JOR",
		Translations: map[Language]Translation{
			"de": {
				Official:    "Haschemitisches Königreich Jordanien",
				Common:      "Jordanien",
				Nationality: "Jordanisch",
			},
			"en": {
				Common:      "Jordan",
				Official:    "Hashemite Kingdom of Jordan",
				Nationality: "Jordanian",
			},
		},
	},
	{
		Alpha2: "IQ",
		Alpha3: "IRQ",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Irak",
				Official:    "Republik Irak",
				Nationality: "Irakisch",
			},
			"en": {
				Common:      "Iraq",
				Official:    "Republic of Iraq",
				Nationality: "Iraqi",
			},
		},
	},
	{
		Alpha2: "IR",
		Alpha3: "IRN",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Iran",
				Official:    "Islamische Republik Iran",
				Nationality: "Iranisch",
			},
			"en": {
				Common:      "Iran",
				Official:    "Islamic Republic of Iran",
				Nationality: "Iranian",
			},
		},
	},
	{
		Alpha2: "IE",
		Alpha3: "IRL",
		Translations: map[Language]Translation{
			"de": {
				Official:    "Republik Irland",
				Common:      "Irland",
				Nationality: "Irisch",
			},
			"en": {
				Common:      "Ireland",
				Official:    "Republic of Ireland",
				Nationality: "Irish",
			},
		},
	},
	{
		Alpha2: "IS",
		Alpha3: "ISL",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Island",
				Official:    "Island",
				Nationality: "Isländisch",
			},
			"en": {
				Common:      "Iceland",
				Official:    "Iceland",
				Nationality: "Icelander",
			},
		},
	},
	{
		Alpha2: "ES",
		Alpha3: "ESP",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Spanien",
				Official:    "Königreich Spanien",
				Nationality: "Spanisch", //nolint:misspell // correct german translation
			},
			"en": {
				Common:      "Spain",
				Official:    "Kingdom of Spain",
				Nationality: "Spanish",
			},
		},
	},
	{
		Alpha2: "IT",
		Alpha3: "ITA",
		Translations: map[Language]Translation{
			"de": {
				Official:    "Italienische Republik",
				Common:      "Italien",
				Nationality: "Italienisch",
			},
			"en": {
				Common:      "Italy",
				Official:    "Italian Republic",
				Nationality: "Italian",
			},
		},
	},
	{
		Alpha2: "YE",
		Alpha3: "YEM",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Jemen",
				Official:    "Republik Jemen",
				Nationality: "Jemenitisch",
			},
			"en": {
				Common:      "Yemen",
				Official:    "Republic of Yemen",
				Nationality: "Yemeni",
			},
		},
	},
	{
		Alpha2: "KZ",
		Alpha3: "KAZ",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Kasachstan",
				Official:    "Republik Kasachstan",
				Nationality: "Kasachisch",
			},
			"en": {
				Common:      "Kazakhstan",
				Official:    "Republic of Kazakhstan",
				Nationality: "Kazakhstani",
			},
		},
	},
	{
		Alpha2: "KY",
		Alpha3: "CYM",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Kaimaninseln",
				Official:    "Cayman-Inseln",
				Nationality: "Kaimaninsel-Bewohner",
			},
			"en": {
				Common:      "Cayman Islands",
				Official:    "Cayman Islands",
				Nationality: "Caymanian",
			},
		},
	},
	{
		Alpha2: "KH",
		Alpha3: "KHM",
		Translations: map[Language]Translation{
			"de": {
				Official:    "Königreich Kambodscha",
				Common:      "Kambodscha",
				Nationality: "Kambodschanisch",
			},
			"en": {
				Common:      "Cambodia",
				Official:    "Kingdom of Cambodia",
				Nationality: "Cambodian",
			},
		},
	},
	{
		Alpha2: "CM",
		Alpha3: "CMR",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Kamerun",
				Official:    "Republik Kamerun",
				Nationality: "Kamerunisch",
			},
			"en": {
				Common:      "Cameroon",
				Official:    "Republic of Cameroon",
				Nationality: "Cameroonian",
			},
		},
	},
	{
		Alpha2: "CA",
		Alpha3: "CAN",
		Translations: map[Language]Translation{
			"de": {
				Official:    "Kanada",
				Common:      "Kanada",
				Nationality: "Kanadisch",
			},
			"en": {
				Common:      "Canada",
				Official:    "Canada",
				Nationality: "Canadian",
			},
		},
	},
	{
		Alpha2: "QA",
		Alpha3: "QAT",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Katar",
				Official:    "Staat Katar",
				Nationality: "Katarisch",
			},
			"en": {
				Common:      "Qatar",
				Official:    "State of Qatar",
				Nationality: "Qatari",
			},
		},
	},
	{
		Alpha2: "KE",
		Alpha3: "KEN",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Kenia",
				Official:    "Republik Kenia",
				Nationality: "Kenianisch",
			},
			"en": {
				Common:      "Kenya",
				Official:    "Republic of Kenya",
				Nationality: "Kenyan",
			},
		},
	},
	{
		Alpha2: "CY",
		Alpha3: "CYP",
		Translations: map[Language]Translation{
			"de": {
				Official:    "Republik Zypern",
				Common:      "Zypern",
				Nationality: "Zyprisch",
			},
			"en": {
				Common:      "Cyprus",
				Official:    "Republic of Cyprus",
				Nationality: "Cypriot",
			},
		},
	},
	{
		Alpha2: "KI",
		Alpha3: "KIR",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Kiribati",
				Official:    "Unabhängige und souveräne Republik Kiribati",
				Nationality: "Kiribatisch",
			},
			"en": {
				Common:      "Kiribati",
				Official:    "Independent and Sovereign Republic of Kiribati",
				Nationality: "Kiribati",
			},
		},
	},
	{
		Alpha2: "CN",
		Alpha3: "CHN",
		Translations: map[Language]Translation{
			"de": {
				Official:    "Volksrepublik China",
				Common:      "China",
				Nationality: "Chinesisch",
			},
			"en": {
				Common:      "China",
				Official:    "People's Republic of China",
				Nationality: "Chinese",
			},
		},
	},
	{
		Alpha2: "CC",
		Alpha3: "CCK",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Kokosinseln",
				Official:    "Gebiet der Cocos (Keeling) Islands",
				Nationality: "Australier",
			},
			"en": {
				Common:      "Cocos (Keeling) Islands",
				Official:    "Territory of the Cocos (Keeling) Islands",
				Nationality: "Australian",
			},
		},
	},
	{
		Alpha2: "CO",
		Alpha3: "COL",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Kolumbien",
				Official:    "Republik Kolumbien",
				Nationality: "Kolumbianisch",
			},
			"en": {
				Common:      "Colombia",
				Official:    "Republic of Colombia",
				Nationality: "Colombian",
			},
		},
	},
	{
		Alpha2: "KM",
		Alpha3: "COM",
		Translations: map[Language]Translation{
			"de": {
				Common:   "Union der Komoren",
				Official: "Komorisch",
			},
			"en": {
				Common:   "Comoros",
				Official: "Union of the Comoros",
			},
		},
	},
	{
		Alpha2: "CG",
		Alpha3: "COG",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Kongo",
				Official:    "Republik Kongo",
				Nationality: "Kongolesisch",
			},
			"en": {
				Common:      "Republic of the Congo",
				Official:    "Republic of the Congo",
				Nationality: "Congolese",
			},
		},
	},
	{
		Alpha2: "KP",
		Alpha3: "PRK",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Nordkorea",
				Official:    "Demokratische Volksrepublik Korea",
				Nationality: "Koreanisch",
			},
			"en": {
				Common:      "North Korea",
				Official:    "Democratic People's Republic of Korea",
				Nationality: "North Korean",
			},
		},
	},
	{
		Alpha2: "KR",
		Alpha3: "KOR",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Südkorea",
				Official:    "Republik Korea",
				Nationality: "Koreanisch",
			},
			"en": {
				Common:      "South Korea",
				Official:    "Republic of Korea",
				Nationality: "Südkoreaner",
			},
		},
	},
	{
		Alpha2: "CR",
		Alpha3: "CRI",
		Translations: map[Language]Translation{
			"de": {
				Official:    "Republik Costa Rica",
				Common:      "Costa Rica",
				Nationality: "Costa-ricanisch",
			},
			"en": {
				Common:      "Costa Rica",
				Official:    "Republic of Costa Rica",
				Nationality: "Costa Rican",
			},
		},
	},
	{
		Alpha2: "CI",
		Alpha3: "CIV",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Elfenbeinküste",
				Official:    "Republik Côte d'Ivoire",
				Nationality: "Ivorisch",
			},
			"en": {
				Common:      "Ivory Coast",
				Official:    "Republic of Côte d'Ivoire",
				Nationality: "Ivorian",
			},
		},
	},
	{
		Alpha2: "CU",
		Alpha3: "CUB",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Kuba",
				Official:    "Republik Kuba",
				Nationality: "Kubanisch",
			},
			"en": {
				Common:      "Cuba",
				Official:    "Republic of Cuba",
				Nationality: "Cuban",
			},
		},
	},
	{
		Alpha2: "KW",
		Alpha3: "KWT",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Kuwait",
				Official:    "Staat Kuwait",
				Nationality: "Kuwaitisch",
			},
			"en": {
				Common:      "Kuwait",
				Official:    "State of Kuwait",
				Nationality: "Kuwaiti",
			},
		},
	},
	{
		Alpha2: "KG",
		Alpha3: "KGZ",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Kirgisistan",
				Official:    "Kirgisische Republik",
				Nationality: "Kirgisisch",
			},
			"en": {
				Common:      "Kyrgyzstan",
				Official:    "Kyrgyz Republic",
				Nationality: "Kyrgyz",
			},
		},
	},
	{
		Alpha2: "LA",
		Alpha3: "LAO",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Laos",
				Official:    "Laos, Demokratische Volksrepublik",
				Nationality: "Laotisch",
			},
			"en": {
				Common:      "Laos",
				Official:    "Lao People's Democratic Republic",
				Nationality: "Laotian",
			},
		},
	},
	{
		Alpha2: "LV",
		Alpha3: "LVA",
		Translations: map[Language]Translation{
			"de": {
				Official:    "Republik Lettland",
				Common:      "Lettland",
				Nationality: "Lettisch",
			},
			"en": {
				Common:      "Latvia",
				Official:    "Republic of Latvia",
				Nationality: "Latvian",
			},
		},
	},
	{
		Alpha2: "LS",
		Alpha3: "LSO",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Lesotho",
				Official:    "Königreich Lesotho",
				Nationality: "Lesothisch",
			},
			"en": {
				Common:      "Lesotho",
				Official:    "Kingdom of Lesotho",
				Nationality: "Mosotho",
			},
		},
	},
	{
		Alpha2: "LR",
		Alpha3: "LBR",
		Translations: map[Language]Translation{
			"de": {
				Official:    "Republik Liberia",
				Common:      "Liberia",
				Nationality: "Liberianisch",
			},
			"en": {
				Common:      "Liberia",
				Official:    "Republic of Liberia",
				Nationality: "Liberian",
			},
		},
	},
	{
		Alpha2: "LB",
		Alpha3: "LBN",
		Translations: map[Language]Translation{
			"de": {
				Official:    "Libanesische Republik",
				Common:      "Libanon",
				Nationality: "Libanesisch",
			},
			"en": {
				Common:      "Lebanon",
				Official:    "Lebanese Republic",
				Nationality: "Lebanese",
			},
		},
	},
	{
		Alpha2: "LY",
		Alpha3: "LBY",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Libyen",
				Official:    "Staat Libyen",
				Nationality: "Libysch",
			},
			"en": {
				Common:      "Libya",
				Official:    "State of Libya",
				Nationality: "Libyan",
			},
		},
	},
	{
		Alpha2: "LT",
		Alpha3: "LTU",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Litauen",
				Official:    "Republik Litauen",
				Nationality: "Litauisch",
			},
			"en": {
				Common:      "Lithuania",
				Official:    "Republic of Lithuania",
				Nationality: "Lithuanian",
			},
		},
	},
	{
		Alpha2: "LI",
		Alpha3: "LIE",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Liechtenstein",
				Official:    "Fürstentum Liechtenstein",
				Nationality: "Liechtensteinisch",
			},
			"en": {
				Common:      "Liechtenstein",
				Official:    "Principality of Liechtenstein",
				Nationality: "Liechtensteiner",
			},
		},
	},
	{
		Alpha2: "LU",
		Alpha3: "LUX",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Luxemburg",
				Official:    "Großherzogtum Luxemburg",
				Nationality: "Luxemburgisch",
			},
			"en": {
				Common:      "Luxembourg",
				Official:    "Grand Duchy of Luxembourg",
				Nationality: "Luxembourger",
			},
		},
	},
	{
		Alpha2: "MU",
		Alpha3: "MUS",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Mauritius",
				Official:    "Republik Mauritius",
				Nationality: "Mauritisch",
			},
			"en": {
				Common:      "Mauritius",
				Official:    "Republic of Mauritius",
				Nationality: "Mauritian",
			},
		},
	},
	{
		Alpha2: "MR",
		Alpha3: "MRT",
		Translations: map[Language]Translation{
			"de": {
				Official:    "Islamische Republik Mauretanien",
				Common:      "Mauretanien",
				Nationality: "Mauretanisch",
			},
			"en": {
				Common:      "Mauritania",
				Official:    "Islamic Republic of Mauritania",
				Nationality: "Mauritanian",
			},
		},
	},
	{
		Alpha2: "MG",
		Alpha3: "MDG",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Madagaskar",
				Official:    "Republik Madagascar",
				Nationality: "Madagassisch",
			},
			"en": {
				Common:      "Madagascar",
				Official:    "Republic of Madagascar",
				Nationality: "Malagasy",
			},
		},
	},
	{
		Alpha2: "YT",
		Alpha3: "MYT",
		Translations: map[Language]Translation{
			"de": {
				Official:    "Übersee-Département Mayotte",
				Common:      "Mayotte",
				Nationality: "Franzose",
			},
			"en": {
				Common:      "Mayotte",
				Official:    "Department of Mayotte",
				Nationality: "French",
			},
		},
	},
	{
		Alpha2: "MO",
		Alpha3: "MAC",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Macao",
				Official:    "Sonderverwaltungsregion Macau der Volksrepublik China",
				Nationality: "Macaoer",
			},
			"en": {
				Common:      "Macau",
				Official:    "Macao Special Administrative Region of the People's Republic of China",
				Nationality: "Macao",
			},
		},
	},
	{
		Alpha2: "MK",
		Alpha3: "MKD",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Mazedonien",
				Official:    "Republik Mazedonien",
				Nationality: "Mazedonier",
			},
			"en": {
				Common:      "Macedonia",
				Official:    "Republic of Macedonia",
				Nationality: "Macedonian",
			},
		},
	},
	{
		Alpha2: "MW",
		Alpha3: "MWI",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Malawi",
				Official:    "Republik Malawi",
				Nationality: "Malawisch",
			},
			"en": {
				Common:      "Malawi",
				Official:    "Republic of Malawi",
				Nationality: "Malawian",
			},
		},
	},
	{
		Alpha2: "MY",
		Alpha3: "MYS",
		Translations: map[Language]Translation{
			"de": {
				Official:    "Malaysia",
				Common:      "Malaysia",
				Nationality: "Malaysisch",
			},
			"en": {
				Common:      "Malaysia",
				Official:    "Malaysia",
				Nationality: "Malaysian",
			},
		},
	},
	{
		Alpha2: "ML",
		Alpha3: "MLI",
		Translations: map[Language]Translation{
			"de": {
				Official:    "Republik Mali",
				Common:      "Mali",
				Nationality: "Malisch",
			},
			"en": {
				Common:      "Mali",
				Official:    "Republic of Mali",
				Nationality: "Malian",
			},
		},
	},
	{
		Alpha2: "MV",
		Alpha3: "MDV",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Malediven",
				Official:    "Republik Malediven",
				Nationality: "Maledivisch",
			},
			"en": {
				Common:      "Maldives",
				Official:    "Republic of the Maldives",
				Nationality: "Maldivian",
			},
		},
	},
	{
		Alpha2: "MT",
		Alpha3: "MLT",
		Translations: map[Language]Translation{
			"de": {
				Official:    "Republik Malta",
				Common:      "Malta",
				Nationality: "Maltesisch",
			},
			"en": {
				Common:      "Malta",
				Official:    "Republic of Malta",
				Nationality: "Maltese",
			},
		},
	},
	{
		Alpha2: "MP",
		Alpha3: "MNP",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Nördliche Marianen",
				Official:    "Commonwealth der Nördlichen Marianen",
				Nationality: "Amerikaner",
			},
			"en": {
				Common:      "Northern Mariana Islands",
				Official:    "Commonwealth of the Northern Mariana Islands",
				Nationality: "American",
			},
		},
	},
	{
		Alpha2: "MA",
		Alpha3: "MAR",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Marokko",
				Official:    "Königreich Marokko",
				Nationality: "Marokkanisch",
			},
			"en": {
				Common:      "Morocco",
				Official:    "Kingdom of Morocco",
				Nationality: "Moroccan",
			},
		},
	},
	{
		Alpha2: "MQ",
		Alpha3: "MTQ",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Martinique",
				Official:    "Martinique",
				Nationality: "Franzose",
			},
			"en": {
				Common:      "Martinique",
				Official:    "Martinique",
				Nationality: "French",
			},
		},
	},
	{
		Alpha2: "MH",
		Alpha3: "MHL",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Marshallinseln",
				Official:    "Republik der Marshall-Inseln",
				Nationality: "Marshallisch",
			},
			"en": {
				Common:      "Marshall Islands",
				Official:    "Republic of the Marshall Islands",
				Nationality: "Marshallese",
			},
		},
	},
	{
		Alpha2: "MX",
		Alpha3: "MEX",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Mexiko",
				Official:    "Vereinigte Mexikanische Staaten",
				Nationality: "Mexikanisch",
			},
			"en": {
				Common:      "Mexico",
				Official:    "United Mexican States",
				Nationality: "Mexican",
			},
		},
	},
	{
		Alpha2: "FM",
		Alpha3: "FSM",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Mikronesien",
				Official:    "Föderierte Staaten von Mikronesien",
				Nationality: "Mikronesisch",
			},
			"en": {
				Common:   "Micronesia",
				Official: "Federated States of Micronesia",
			},
		},
	},
	{
		Alpha2: "MZ",
		Alpha3: "MOZ",
		Translations: map[Language]Translation{
			"de": {
				Official:    "Republik Mosambik",
				Common:      "Mosambik",
				Nationality: "Mosambikanisch",
			},
			"en": {
				Common:      "Mozambique",
				Official:    "Republic of Mozambique",
				Nationality: "Micronesian",
			},
		},
	},
	{
		Alpha2: "MD",
		Alpha3: "MDA",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Moldawie",
				Official:    "Republik Moldau",
				Nationality: "Moldauisch",
			},
			"en": {
				Common:      "Moldova",
				Official:    "Republic of Moldova",
				Nationality: "Moldovan",
			},
		},
	},
	{
		Alpha2: "MC",
		Alpha3: "MCO",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Monaco",
				Official:    "Fürstentum Monaco",
				Nationality: "Monegassisch",
			},
			"en": {
				Common:      "Monaco",
				Official:    "Principality of Monaco",
				Nationality: "Monacan",
			},
		},
	},
	{
		Alpha2: "MN",
		Alpha3: "MNG",
		Translations: map[Language]Translation{
			"de": {
				Official:    "Mongolei",
				Common:      "Mongolei",
				Nationality: "Mongolisch",
			},
			"en": {
				Common:      "Mongolia",
				Official:    "Mongolia",
				Nationality: "Mongolian",
			},
		},
	},
	{
		Alpha2: "MS",
		Alpha3: "MSR",
		Translations: map[Language]Translation{
			"de": {
				Official:    "Montserrat",
				Common:      "Montserrat",
				Nationality: "Montserrat-Bewohner",
			},
			"en": {
				Common:      "Montserrat",
				Official:    "Montserrat",
				Nationality: "Montserratian",
			},
		},
	},
	{
		Alpha2: "MM",
		Alpha3: "MMR",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Myanmar",
				Official:    "Republik der Union von Myanmar",
				Nationality: "Myanmarisch",
			},
			"en": {
				Common:      "Myanmar",
				Official:    "Republic of the Union of Myanmar",
				Nationality: "Myanmar (Burmese)",
			},
		},
	},
	{
		Alpha2: "NA",
		Alpha3: "NAM",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Namibia",
				Official:    "Republik Namibia",
				Nationality: "Namibisch",
			},
			"en": {
				Common:      "Namibia",
				Official:    "Republic of Namibia",
				Nationality: "Namibian",
			},
		},
	},
	{
		Alpha2: "NR",
		Alpha3: "NRU",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Nauru",
				Official:    "Republik Nauru",
				Nationality: "Nauruisch",
			},
			"en": {
				Common:      "Nauru",
				Official:    "Republic of Nauru",
				Nationality: "Nauruan",
			},
		},
	},
	{
		Alpha2: "NP",
		Alpha3: "NPL",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Népal",
				Official:    "Demokratischen Bundesrepublik Nepal",
				Nationality: "Nepalesisch",
			},
			"en": {
				Common:      "Nepal",
				Official:    "Federal Democratic Republic of Nepal",
				Nationality: "Nepalese",
			},
		},
	},
	{
		Alpha2: "NE",
		Alpha3: "NER",
		Translations: map[Language]Translation{
			"de": {
				Official:    "Republik Niger",
				Common:      "Niger",
				Nationality: "Nigrisch",
			},
			"en": {
				Common:      "Niger",
				Official:    "Republic of Niger",
				Nationality: "Nigerien",
			},
		},
	},
	{
		Alpha2: "NG",
		Alpha3: "NGA",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Nigeria",
				Official:    "Bundesrepublik Nigeria",
				Nationality: "Nigerianisch",
			},
			"en": {
				Common:      "Nigeria",
				Official:    "Federal Republic of Nigeria",
				Nationality: "Nigerian",
			},
		},
	},
	{
		Alpha2: "NL",
		Alpha3: "NLD",
		Translations: map[Language]Translation{
			"de": {
				Official:    "Niederlande",
				Common:      "Niederlande",
				Nationality: "Niederländisch",
			},
			"en": {
				Common:      "Netherlands",
				Official:    "Netherlands",
				Nationality: "Dutch",
			},
		},
	},
	{
		Alpha2: "NI",
		Alpha3: "NIC",
		Translations: map[Language]Translation{
			"de": {
				Official:    "Republik Nicaragua",
				Common:      "Nicaragua",
				Nationality: "Nicaraguanisch",
			},
			"en": {
				Common:      "Nicaragua",
				Official:    "Republic of Nicaragua",
				Nationality: "Nicaraguan",
			},
		},
	},
	{
		Alpha2: "NU",
		Alpha3: "NIU",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Niue",
				Official:    "Niue",
				Nationality: "Niueanisch",
			},
			"en": {
				Common:      "Niue",
				Official:    "Niue",
				Nationality: "Niueaner",
			},
		},
	},
	{
		Alpha2: "NZ",
		Alpha3: "NZL",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Neuseeland",
				Official:    "Neuseeland",
				Nationality: "Neuseeländisch",
			},
			"en": {
				Common:      "New Zealand",
				Official:    "New Zealand",
				Nationality: "New Zealander",
			},
		},
	},
	{
		Alpha2: "NC",
		Alpha3: "NCL",
		Translations: map[Language]Translation{
			"de": {
				Official:    "Neukaledonien",
				Common:      "Neukaledonien",
				Nationality: "Franzose",
			},
			"en": {
				Common:      "New Caledonia",
				Official:    "New Caledonia",
				Nationality: "French",
			},
		},
	},
	{
		Alpha2: "NO",
		Alpha3: "NOR",
		Translations: map[Language]Translation{
			"de": {
				Official:    "Königreich Norwegen",
				Common:      "Norwegen",
				Nationality: "Norwegisch",
			},
			"en": {
				Common:      "Norway",
				Official:    "Kingdom of Norway",
				Nationality: "Norwegian",
			},
		},
	},
	{
		Alpha2: "OM",
		Alpha3: "OMN",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Oman",
				Official:    "Sultanat Oman",
				Nationality: "Omanisch",
			},
			"en": {
				Common:      "Oman",
				Official:    "Sultanate of Oman",
				Nationality: "Omani",
			},
		},
	},
	{
		Alpha2: "BV",
		Alpha3: "BVT",
		Translations: map[Language]Translation{
			"de": {
				Official: "Bouvet-Insel",
				Common:   "Bouvetinsel",
			},
			"en": {
				Common:   "Bouvet Island",
				Official: "Bouvet Island",
			},
		},
	},
	{
		Alpha2: "IM",
		Alpha3: "IMN",
		Translations: map[Language]Translation{
			"de": {
				Official:    "Isle of Man",
				Common:      "Insel Man",
				Nationality: "Manxer",
			},
			"en": {
				Common:      "Isle of Man",
				Official:    "Isle of Man",
				Nationality: "Manx",
			},
		},
	},
	{
		Alpha2: "NF",
		Alpha3: "NFK",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Norfolkinsel",
				Official:    "Gebiet der Norfolk-Insel",
				Nationality: "Manx",
			},
			"en": {
				Common:      "Norfolk Island",
				Official:    "Territory of Norfolk Island",
				Nationality: "Manx",
			},
		},
	},
	{
		Alpha2: "PN",
		Alpha3: "PCN",
		Translations: map[Language]Translation{
			"de": {
				Official:    "Pitcairn Inselgruppe",
				Common:      "Pitcairn",
				Nationality: "Pitcairn-Inselbewohner",
			},
			"en": {
				Common:      "Pitcairn Islands",
				Official:    "Pitcairn Group of Islands",
				Nationality: "Pitcairn Islander",
			},
		},
	},
	{
		Alpha2: "CX",
		Alpha3: "CXR",
		Translations: map[Language]Translation{
			"de": {
				Official:    "Gebiet der Weihnachtsinsel",
				Common:      "Weihnachtsinsel",
				Nationality: "Weihnachtsinselbewohner",
			},
			"en": {
				Common:      "Christmas Island",
				Official:    "Territory of Christmas Island",
				Nationality: "Christmas Islander",
			},
		},
	},
	{
		Alpha2: "SH",
		Alpha3: "SHN",
		Translations: map[Language]Translation{
			"de": {
				Common:      "St. Helena",
				Official:    "St. Helena, Ascension und Tristan da Cunha",
				Nationality: "St. Helenianer",
			},
			"en": {
				Common:      "Saint Helena",
				Official:    "Saint Helena, Ascension and Tristan da Cunha",
				Nationality: "Saint Helenian",
			},
		},
	},
	{
		Alpha2: "WF",
		Alpha3: "WLF",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Wallis und Futuna",
				Official:    "Gebiet der Wallis und Futuna",
				Nationality: "Wallis und Futuna Inselbewohner",
			},
			"en": {
				Common:      "Wallis and Futuna",
				Official:    "Territory of the Wallis and Futuna Islands",
				Nationality: "Wallisian and Futuna",
			},
		},
	},
	{
		Alpha2: "HM",
		Alpha3: "HMD",
		Translations: map[Language]Translation{
			"de": {
				Official: "Heard und McDonaldinseln",
				Common:   "Heard und die McDonaldinseln",
			},
			"en": {
				Common:   "Heard Island and McDonald Islands",
				Official: "Heard Island and McDonald Islands",
			},
		},
	},
	{
		Alpha2: "CV",
		Alpha3: "CPV",
		Translations: map[Language]Translation{
			"de": {
				Official:    "Republik Cabo Verde",
				Common:      "Kap Verde",
				Nationality: "Kap-Verdier",
			},
			"en": {
				Common:      "Cape Verde",
				Official:    "Republic of Cabo Verde",
				Nationality: "Cape Verdean",
			},
		},
	},
	{
		Alpha2: "CK",
		Alpha3: "COK",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Cookinseln",
				Official:    "Cook-Inseln",
				Nationality: "Cook-Inselbewohner",
			},
			"en": {
				Common:      "Cook Islands",
				Official:    "Cook Islands",
				Nationality: "Cook Islander",
			},
		},
	},
	{
		Alpha2: "WS",
		Alpha3: "WSM",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Samoa",
				Official:    "Unabhängige Staat Samoa",
				Nationality: "Samoanisch",
			},
			"en": {
				Common:      "Samoa",
				Official:    "Independent State of Samoa",
				Nationality: "Samoan",
			},
		},
	},
	{
		Alpha2: "SJ",
		Alpha3: "SJM",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Spitzbergen",
				Official:    "Inselgruppe Spitzbergen",
				Nationality: "Norwegian",
			},
			"en": {
				Common:      "Svalbard and Jan Mayen",
				Official:    "Svalbard og Jan Mayen",
				Nationality: "Norweger",
			},
		},
	},
	{
		Alpha2: "TC",
		Alpha3: "TCA",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Turks-und Caicosinseln",
				Official:    "Turks und Caicos Inseln",
				Nationality: "Inselbewohner von Turks- und Caicosinseln",
			},
			"en": {
				Common:      "Turks and Caicos Islands",
				Official:    "Turks and Caicos Islands",
				Nationality: "Turks and Caicos Islander",
			},
		},
	},
	{
		Alpha2: "UM",
		Alpha3: "UMI",
		Translations: map[Language]Translation{
			"de": {
				Official: "USA, kleinere ausgelagerte Inseln",
				Common:   "Kleinere Inselbesitzungen der Vereinigten Staaten",
			},
			"en": {
				Common:   "United States Minor Outlying Islands",
				Official: "United States Minor Outlying Islands",
			},
		},
	},
	{
		Alpha2: "PK",
		Alpha3: "PAK",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Pakistan",
				Official:    "Islamische Republik Pakistan",
				Nationality: "Pakistanisch",
			},
			"en": {
				Common:      "Pakistan",
				Official:    "Islamic Republic of Pakistan",
				Nationality: "Pakistani",
			},
		},
	},
	{
		Alpha2: "PW",
		Alpha3: "PLW",
		Translations: map[Language]Translation{
			"de": {
				Official:    "Palau",
				Common:      "Palau",
				Nationality: "Palauisch",
			},
			"en": {
				Common:      "Palau",
				Official:    "Republic of Palau",
				Nationality: "Palauan",
			},
		},
	},
	{
		Alpha2: "PS",
		Alpha3: "PSE",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Palästina",
				Official:    "Staat Palästina",
				Nationality: "Palästinenser",
			},
			"en": {
				Common:      "Palestine",
				Official:    "State of Palestine",
				Nationality: "Palestinian",
			},
		},
	},
	{
		Alpha2: "PA",
		Alpha3: "PAN",
		Translations: map[Language]Translation{
			"de": {
				Official:    "Republik Panama",
				Common:      "Panama",
				Nationality: "Panamaisch",
			},
			"en": {
				Common:      "Panama",
				Official:    "Republic of Panama",
				Nationality: "Panamanian",
			},
		},
	},
	{
		Alpha2: "PG",
		Alpha3: "PNG",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Papua-Neuguinea",
				Official:    "Unabhängige Staat Papua-Neuguinea",
				Nationality: "Papua-neuguineisch",
			},
			"en": {
				Common:      "Papua New Guinea",
				Official:    "Independent State of Papua New Guinea",
				Nationality: "Papua New Guinean",
			},
		},
	},
	{
		Alpha2: "PY",
		Alpha3: "PRY",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Paraguay",
				Official:    "Republik Paraguay",
				Nationality: "Paraguayisch",
			},
			"en": {
				Common:      "Paraguay",
				Official:    "Republic of Paraguay",
				Nationality: "Paraguayan",
			},
		},
	},
	{
		Alpha2: "PE",
		Alpha3: "PER",
		Translations: map[Language]Translation{
			"de": {
				Official:    "Republik Peru",
				Common:      "Peru",
				Nationality: "Peruanisch",
			},
			"en": {
				Common:      "Peru",
				Official:    "Republic of Peru",
				Nationality: "Peruvian",
			},
		},
	},
	{
		Alpha2: "PL",
		Alpha3: "POL",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Polen",
				Official:    "Republik Polen",
				Nationality: "Polnisch",
			},
			"en": {
				Common:      "Poland",
				Official:    "Republic of Poland",
				Nationality: "Polish",
			},
		},
	},
	{
		Alpha2: "PT",
		Alpha3: "PRT",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Portugal",
				Official:    "Portugiesische Republik",
				Nationality: "Portugiesisch",
			},
			"en": {
				Common:      "Portugal",
				Official:    "Portuguese Republic",
				Nationality: "Portuguese",
			},
		},
	},
	{
		Alpha2: "PR",
		Alpha3: "PRI",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Puerto Rico",
				Official:    "Commonwealth von Puerto Rico",
				Nationality: "Puerto-Ricaner",
			},
			"en": {
				Common:      "Puerto Rico",
				Official:    "Commonwealth of Puerto Rico",
				Nationality: "Puerto Rican",
			},
		},
	},
	{
		Alpha2: "RE",
		Alpha3: "REU",
		Translations: map[Language]Translation{
			"de": {
				Official:    "Réunion",
				Common:      "Réunion",
				Nationality: "Franzose",
			},
			"en": {
				Common:      "Réunion",
				Official:    "Réunion Island",
				Nationality: "French",
			},
		},
	},
	{
		Alpha2: "RU",
		Alpha3: "RUS",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Russland",
				Official:    "Russische Föderation",
				Nationality: "Russisch",
			},
			"en": {
				Common:      "Russia",
				Official:    "Russian Federation",
				Nationality: "Russian",
			},
		},
	},
	{
		Alpha2: "RW",
		Alpha3: "RWA",
		Translations: map[Language]Translation{
			"de": {
				Official:    "Republik Ruanda",
				Common:      "Ruanda",
				Nationality: "Ruandisch",
			},
			"en": {
				Common:      "Rwanda",
				Official:    "Republic of Rwanda",
				Nationality: "Rwandan",
			},
		},
	},
	{
		Alpha2: "RO",
		Alpha3: "ROU",
		Translations: map[Language]Translation{
			"de": {
				Official:    "Rumänien",
				Common:      "Rumänien",
				Nationality: "Rumänisch",
			},
			"en": {
				Common:      "Romania",
				Official:    "Romania",
				Nationality: "Romanian",
			},
		},
	},
	{
		Alpha2: "SV",
		Alpha3: "SLV",
		Translations: map[Language]Translation{
			"de": {
				Common:      "El Salvador",
				Official:    "Republik El Salvador",
				Nationality: "Salvadorianisch",
			},
			"en": {
				Common:      "El Salvador",
				Official:    "Republic of El Salvador",
				Nationality: "Salvadoran",
			},
		},
	},
	{
		Alpha2: "SM",
		Alpha3: "SMR",
		Translations: map[Language]Translation{
			"de": {
				Common:      "San Marino",
				Official:    "Republik San Marino",
				Nationality: "San-marinesisch",
			},
			"en": {
				Common:      "San Marino",
				Official:    "Most Serene Republic of San Marino",
				Nationality: "San Marinese",
			},
		},
	},
	{
		Alpha2: "ST",
		Alpha3: "STP",
		Translations: map[Language]Translation{
			"de": {
				Common:      "São Tomé und Príncipe",
				Official:    "Demokratische Republik São Tomé und Príncipe",
				Nationality: "sro-tomйisch",
			},
			"en": {
				Common:      "São Tomé and Príncipe",
				Official:    "Democratic Republic of São Tomé and Príncipe",
				Nationality: "Sao Tomean",
			},
		},
	},
	{
		Alpha2: "SA",
		Alpha3: "SAU",
		Translations: map[Language]Translation{
			"de": {
				Official:    "Königreich Saudi-Arabien",
				Common:      "Saudi-Arabien",
				Nationality: "Saudi-arabisch",
			},
			"en": {
				Common:      "Saudi Arabia",
				Official:    "Kingdom of Saudi Arabia",
				Nationality: "Saudi Arabian",
			},
		},
	},
	{
		Alpha2: "SZ",
		Alpha3: "SWZ",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Swasiland",
				Official:    "Königreich Swasiland",
				Nationality: "Swasi",
			},
			"en": {
				Common:      "Swaziland",
				Official:    "Kingdom of Swaziland",
				Nationality: "Swasi",
			},
		},
	},
	{
		Alpha2: "SC",
		Alpha3: "SYC",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Seychellen",
				Official:    "Republik der Seychellen",
				Nationality: "Seychellisch",
			},
			"en": {
				Common:      "Seychelles",
				Official:    "Republic of Seychelles",
				Nationality: "Seychellois",
			},
		},
	},
	{
		Alpha2: "SN",
		Alpha3: "SEN",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Senegal",
				Official:    "Republik Senegal",
				Nationality: "Senegalesisch",
			},
			"en": {
				Common:      "Senegal",
				Official:    "Republic of Senegal",
				Nationality: "Senegalese",
			},
		},
	},
	{
		Alpha2: "PM",
		Alpha3: "SPM",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Saint-Pierre und Miquelon",
				Official:    "St. Pierre und Miquelon",
				Nationality: "Saint-Pierrais und Miquelonnais",
			},
			"en": {
				Common:      "Saint Pierre and Miquelon",
				Official:    "Saint Pierre and Miquelon",
				Nationality: "Saint-Pierrais and Miquelonnais",
			},
		},
	},
	{
		Alpha2: "VC",
		Alpha3: "VCT",
		Translations: map[Language]Translation{
			"de": {
				Official:    "St. Vincent und die Grenadinen",
				Common:      "Saint Vincent und die Grenadinen",
				Nationality: "Vincentisch",
			},
			"en": {
				Common:      "Saint Vincent and the Grenadines",
				Official:    "Saint Vincent and the Grenadines",
				Nationality: "Vincentian",
			},
		},
	},
	{
		Alpha2: "KN",
		Alpha3: "KNA",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Saint Christopher und Nevis",
				Official:    "Föderation von Saint Kitts und Nevisa",
				Nationality: "Kittitianer und Nevisianer",
			},
			"en": {
				Common:      "Saint Kitts and Nevis",
				Official:    "Federation of Saint Christopher and Nevisa",
				Nationality: "Kittitian and Nevisian",
			},
		},
	},
	{
		Alpha2: "LC",
		Alpha3: "LCA",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Saint Lucia",
				Official:    "St. Lucia",
				Nationality: "Lucianisch",
			},
			"en": {
				Common:      "Saint Lucia",
				Official:    "Saint Lucia",
				Nationality: "Saint Lucian",
			},
		},
	},
	{
		Alpha2: "SG",
		Alpha3: "SGP",
		Translations: map[Language]Translation{
			"de": {
				Official:    "Republik Singapur",
				Common:      "Singapur",
				Nationality: "Singapurisch",
			},
			"en": {
				Common:      "Singapore",
				Official:    "Republic of Singapore",
				Nationality: "Singaporean",
			},
		},
	},
	{
		Alpha2: "SY",
		Alpha3: "SYR",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Syrien",
				Official:    "Arabische Republik Syrien",
				Nationality: "Syrisch",
			},
			"en": {
				Common:      "Syria",
				Official:    "Syrian Arab Republic",
				Nationality: "Syrian",
			},
		},
	},
	{
		Alpha2: "SK",
		Alpha3: "SVK",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Slowakei",
				Official:    "Slowakische Republik",
				Nationality: "Slowakisch",
			},
			"en": {
				Common:      "Slovakia",
				Official:    "Slovak Republic",
				Nationality: "Slovak",
			},
		},
	},
	{
		Alpha2: "SI",
		Alpha3: "SVN",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Slowenien",
				Official:    "Republik Slowenien",
				Nationality: "Slowenisch",
			},
			"en": {
				Common:      "Slovenia",
				Official:    "Republic of Slovenia",
				Nationality: "Slovenian",
			},
		},
	},
	{
		Alpha2: "US",
		Alpha3: "USA",
		Translations: map[Language]Translation{
			"de": {
				Official:    "Vereinigte Staaten von Amerika",
				Common:      "Vereinigte Staaten von Amerika",
				Nationality: "Amerikanisch",
			},
			"en": {
				Common:      "United States",
				Official:    "United States of America",
				Nationality: "American",
			},
		},
	},
	{
		Alpha2: "SB",
		Alpha3: "SLB",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Salomonen",
				Official:    "Salomon-Inseln",
				Nationality: "Salomonisch",
			},
			"en": {
				Common:      "Solomon Islands",
				Official:    "Solomon Islands",
				Nationality: "Solomon Islander",
			},
		},
	},
	{
		Alpha2: "SO",
		Alpha3: "SOM",
		Translations: map[Language]Translation{
			"de": {
				Official:    "Bundesrepublik Somalia",
				Common:      "Somalia",
				Nationality: "Somalisch",
			},
			"en": {
				Common:      "Somalia",
				Official:    "Federal Republic of Somalia",
				Nationality: "Somali",
			},
		},
	},
	{
		Alpha2: "SD",
		Alpha3: "SDN",
		Translations: map[Language]Translation{
			"de": {
				Official:    "Republik Sudan",
				Common:      "Sudan",
				Nationality: "Sudanesisch",
			},
			"en": {
				Common:      "Sudan",
				Official:    "Republic of the Sudan",
				Nationality: "Sudanese",
			},
		},
	},
	{
		Alpha2: "SR",
		Alpha3: "SUR",
		Translations: map[Language]Translation{
			"de": {
				Official:    "Republik Suriname",
				Common:      "Suriname",
				Nationality: "Surinamisch",
			},
			"en": {
				Common:      "Suriname",
				Official:    "Republic of Suriname",
				Nationality: "Surinamese",
			},
		},
	},
	{
		Alpha2: "SL",
		Alpha3: "SLE",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Sierra Leone",
				Official:    "Republik Sierra Leone",
				Nationality: "Sierra-leonisch",
			},
			"en": {
				Common:      "Sierra Leone",
				Official:    "Republic of Sierra Leone",
				Nationality: "Sierra Leonean",
			},
		},
	},
	{
		Alpha2: "TJ",
		Alpha3: "TJK",
		Translations: map[Language]Translation{
			"de": {
				Official:    "Republik Tadschikistan",
				Common:      "Tadschikistan",
				Nationality: "Tadschikisch",
			},
			"en": {
				Common:      "Tajikistan",
				Official:    "Republic of Tajikistan",
				Nationality: "Tajik",
			},
		},
	},
	{
		Alpha2: "TW",
		Alpha3: "TWN",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Taiwan",
				Official:    "Republik China (Taiwan)",
				Nationality: "Taiwaner",
			},
			"en": {
				Common:      "Taiwan",
				Official:    "Republic of China (Taiwan)",
				Nationality: "Taiwanese",
			},
		},
	},
	{
		Alpha2: "TH",
		Alpha3: "THA",
		Translations: map[Language]Translation{
			"de": {
				Official:    "Königreich Thailand",
				Common:      "Thailand",
				Nationality: "Thailändisch",
			},
			"en": {
				Common:      "Thailand",
				Official:    "Kingdom of Thailand",
				Nationality: "Thai",
			},
		},
	},
	{
		Alpha2: "TZ",
		Alpha3: "TZA",
		Translations: map[Language]Translation{
			"de": {
				Official:    "Vereinigte Republik Tansania",
				Common:      "Tansania",
				Nationality: "Tansanisch",
			},
			"en": {
				Common:      "Tanzania",
				Official:    "United Republic of Tanzania",
				Nationality: "Tanzanian",
			},
		},
	},
	{
		Alpha2: "TG",
		Alpha3: "TGO",
		Translations: map[Language]Translation{
			"de": {
				Official:    "Republik Togo",
				Common:      "Togo",
				Nationality: "Togoisch",
			},
			"en": {
				Common:      "Togo",
				Official:    "Togolese Republic",
				Nationality: "Togolese",
			},
		},
	},
	{
		Alpha2: "TK",
		Alpha3: "TKL",
		Translations: map[Language]Translation{
			"de": {
				Official:    "Tokelau",
				Common:      "Tokelau",
				Nationality: "Tokelauer",
			},
			"en": {
				Common:      "Tokelau",
				Official:    "Tokelau",
				Nationality: "Tokelauan",
			},
		},
	},
	{
		Alpha2: "TO",
		Alpha3: "TON",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Tonga",
				Official:    "Königreich Tonga",
				Nationality: "Tongaisch",
			},
			"en": {
				Common:      "Tonga",
				Official:    "Kingdom of Tonga",
				Nationality: "Tongan",
			},
		},
	},
	{
		Alpha2: "TT",
		Alpha3: "TTO",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Trinidad und Tobago",
				Official:    "Republik Trinidad und Tobago",
				Nationality: "Trinidaner und Tobager",
			},
			"en": {
				Common:      "Trinidad and Tobago",
				Official:    "Republic of Trinidad and Tobago",
				Nationality: "Trinidadian and Tobagonian",
			},
		},
	},
	{
		Alpha2: "TV",
		Alpha3: "TUV",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Tuvalu",
				Official:    "Tuvalu",
				Nationality: "Tuvaluisch",
			},
			"en": {
				Common:      "Tuvalu",
				Official:    "Tuvalu",
				Nationality: "Tuvaluan",
			},
		},
	},
	{
		Alpha2: "TN",
		Alpha3: "TUN",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Tunesien",
				Official:    "Tunesische Republik",
				Nationality: "Tunesisch",
			},
			"en": {
				Common:      "Tunisia",
				Official:    "Tunisian Republic",
				Nationality: "Tunisian",
			},
		},
	},
	{
		Alpha2: "TM",
		Alpha3: "TKM",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Turkmenistan",
				Official:    "Turkmenistan",
				Nationality: "Turkmenisch",
			},
			"en": {
				Common:      "Turkmenistan",
				Official:    "Turkmenistan",
				Nationality: "Turkmen",
			},
		},
	},
	{
		Alpha2: "TR",
		Alpha3: "TUR",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Türkei",
				Official:    "Republik Türkei",
				Nationality: "Türkisch",
			},
			"en": {
				Common:      "Turkey",
				Official:    "Republic of Turkey",
				Nationality: "Turkish",
			},
		},
	},
	{
		Alpha2: "UG",
		Alpha3: "UGA",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Uganda",
				Official:    "Republik Uganda",
				Nationality: "Ugandisch",
			},
			"en": {
				Common:      "Uganda",
				Official:    "Republic of Uganda",
				Nationality: "Ugandan",
			},
		},
	},
	{
		Alpha2: "UZ",
		Alpha3: "UZB",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Usbekistan",
				Official:    "Republik Usbekistan",
				Nationality: "Usbekisch",
			},
			"en": {
				Common:      "Uzbekistan",
				Official:    "Republic of Uzbekistan",
				Nationality: "Uzbek",
			},
		},
	},
	{
		Alpha2: "UA",
		Alpha3: "UKR",
		Translations: map[Language]Translation{
			"de": {
				Official:    "Ukraine",
				Common:      "Ukraine",
				Nationality: "Ukrainisch",
			},
			"en": {
				Common:      "Ukraine",
				Official:    "Ukraine",
				Nationality: "Ukrainian",
			},
		},
	},
	{
		Alpha2: "UY",
		Alpha3: "URY",
		Translations: map[Language]Translation{
			"de": {
				Official:    "Republik Östlich des Uruguay",
				Common:      "Uruguay",
				Nationality: "Uruguayisch",
			},
			"en": {
				Common:      "Uruguay",
				Official:    "Oriental Republic of Uruguay",
				Nationality: "Uruguayan",
			},
		},
	},
	{
		Alpha2: "FO",
		Alpha3: "FRO",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Färöer-Inseln",
				Official:    "Färöer",
				Nationality: "Färinger",
			},
			"en": {
				Common:      "Faroe Islands",
				Official:    "Faroe Islands",
				Nationality: "Faroe Islander",
			},
		},
	},
	{
		Alpha2: "FJ",
		Alpha3: "FJI",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Fidschi",
				Official:    "Republik Fidschi",
				Nationality: "Fidschianisch",
			},
			"en": {
				Common:      "Fiji",
				Official:    "Republic of Fiji",
				Nationality: "Fijian",
			},
		},
	},
	{
		Alpha2: "PH",
		Alpha3: "PHL",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Philippinen",
				Official:    "Republik der Philippinen",
				Nationality: "Philippinisch",
			},
			"en": {
				Common:      "Philippines",
				Official:    "Republic of the Philippines",
				Nationality: "Filipino",
			},
		},
	},
	{
		Alpha2: "FI",
		Alpha3: "FIN",
		Translations: map[Language]Translation{
			"de": {
				Official:    "Republik Finnland",
				Common:      "Finnland",
				Nationality: "Finnisch", //nolint:misspell // correct german translation
			},
			"en": {
				Common:      "Finland",
				Official:    "Republic of Finland",
				Nationality: "Finnish",
			},
		},
	},
	{
		Alpha2: "FK",
		Alpha3: "FLK",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Falklandinseln",
				Official:    "Falkland-Inseln",
				Nationality: "Finnish",
			},
			"en": {
				Common:      "Falkland Islands",
				Official:    "Falkland Islands",
				Nationality: "Finnish",
			},
		},
	},
	{
		Alpha2: "FR",
		Alpha3: "FRA",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Frankreich",
				Official:    "Französische Republik",
				Nationality: "Französisch",
			},
			"en": {
				Common:      "France",
				Official:    "French Republic",
				Nationality: "French",
			},
		},
	},
	{
		Alpha2: "GF",
		Alpha3: "GUF",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Französisch Guyana",
				Official:    "Guayana",
				Nationality: "Französisch-Guianisch",
			},
			"en": {
				Common:      "French Guiana",
				Official:    "Guiana",
				Nationality: "French Guianese",
			},
		},
	},
	{
		Alpha2: "PF",
		Alpha3: "PYF",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Französisch-Polynesien",
				Official:    "Französisch-Polynesien",
				Nationality: "Französisch-Polynesien",
			},
			"en": {
				Common:      "French Polynesia",
				Official:    "French Polynesia",
				Nationality: "French Polynesian",
			},
		},
	},
	{
		Alpha2: "TF",
		Alpha3: "ATF",
		Translations: map[Language]Translation{
			"de": {
				Common:   "Französische Süd-und Antarktisgebiete",
				Official: "Gebiet der Französisch Süd-und Antarktisgebiete",
			},
			"en": {
				Common:   "French Southern and Antarctic Lands",
				Official: "Territory of the French Southern and Antarctic Lands",
			},
		},
	},
	{
		Alpha2: "HR",
		Alpha3: "HRV",
		Translations: map[Language]Translation{
			"de": {
				Official:    "Republik Kroatien",
				Common:      "Kroatien",
				Nationality: "Kroatisch",
			},
			"en": {
				Common:      "Croatia",
				Official:    "Republic of Croatia",
				Nationality: "Croatia",
			},
		},
	},
	{
		Alpha2: "CF",
		Alpha3: "CAF",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Zentralafrikanische Republik",
				Official:    "Zentralafrikanische Republik",
				Nationality: "Zentralafrikanisch",
			},
			"en": {
				Common:      "Central African Republic",
				Official:    "Central African Republic",
				Nationality: "Central African Republic",
			},
		},
	},
	{
		Alpha2: "TD",
		Alpha3: "TCD",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Tschad",
				Official:    "Republik Tschad",
				Nationality: "Tschadisch",
			},
			"en": {
				Common:      "Chad",
				Official:    "Republic of Chad",
				Nationality: "Chad",
			},
		},
	},
	{
		Alpha2: "CZ",
		Alpha3: "CZE",
		Translations: map[Language]Translation{
			"de": {
				Official:    "Tschechische Republik",
				Common:      "Tschechische Republik",
				Nationality: "Tschechisch",
			},
			"en": {
				Common:      "Czech Republic",
				Official:    "Czech Republic",
				Nationality: "Czech",
			},
		},
	},
	{
		Alpha2: "CL",
		Alpha3: "CHL",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Chile",
				Official:    "Republik Chile",
				Nationality: "Chilenisch",
			},
			"en": {
				Common:      "Chile",
				Official:    "Republic of Chile",
				Nationality: "Chilean",
			},
		},
	},
	{
		Alpha2: "CH",
		Alpha3: "CHE",
		Translations: map[Language]Translation{
			"de": {
				Official:    "Schweizerische Eidgenossenschaft",
				Common:      "Schweiz",
				Nationality: "Schweizerisch",
			},
			"en": {
				Common:      "Switzerland",
				Official:    "Swiss Confederation",
				Nationality: "Swiss",
			},
		},
	},
	{
		Alpha2: "SE",
		Alpha3: "SWE",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Schweden",
				Official:    "Königreich Schweden",
				Nationality: "Schwedisch",
			},
			"en": {
				Common:      "Sweden",
				Official:    "Kingdom of Sweden",
				Nationality: "Swedish",
			},
		},
	},
	{
		Alpha2: "LK",
		Alpha3: "LKA",
		Translations: map[Language]Translation{
			"de": {
				Official:    "Demokratische Sozialistische Republik Sri Lanka",
				Common:      "Sri Lanka",
				Nationality: "Sri-lankisch",
			},
			"en": {
				Common:      "Sri Lanka",
				Official:    "Democratic Socialist Republic of Sri Lanka",
				Nationality: "Sri Lankan",
			},
		},
	},
	{
		Alpha2: "EC",
		Alpha3: "ECU",
		Translations: map[Language]Translation{
			"de": {
				Official:    "Republik Ecuador",
				Common:      "Ecuador",
				Nationality: "Ecuadorianisch",
			},
			"en": {
				Common:      "Ecuador",
				Official:    "Republic of Ecuador",
				Nationality: "Ecuadorian",
			},
		},
	},
	{
		Alpha2: "GQ",
		Alpha3: "GNQ",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Äquatorialguinea",
				Official:    "Republik Äquatorialguinea",
				Nationality: "Äquatorial-guineisch",
			},
			"en": {
				Common:      "Equatorial Guinea",
				Official:    "Republic of Equatorial Guinea",
				Nationality: "Equatorial Guinean",
			},
		},
	},
	{
		Alpha2: "ER",
		Alpha3: "ERI",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Eritrea",
				Official:    "Staat Eritrea",
				Nationality: "Eritreisch",
			},
			"en": {
				Common:      "Eritrea",
				Official:    "State of Eritrea",
				Nationality: "Eritrean",
			},
		},
	},
	{
		Alpha2: "EE",
		Alpha3: "EST",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Estland",
				Official:    "Republik Estland",
				Nationality: "Estnisch",
			},
			"en": {
				Common:      "Estonia",
				Official:    "Republic of Estonia",
				Nationality: "Estonian",
			},
		},
	},
	{
		Alpha2: "ET",
		Alpha3: "ETH",
		Translations: map[Language]Translation{
			"de": {
				Official:    "Demokratische Bundesrepublik Äthiopien",
				Common:      "Äthiopien",
				Nationality: "Äthiopisch",
			},
			"en": {
				Common:      "Ethiopia",
				Official:    "Federal Democratic Republic of Ethiopia",
				Nationality: "Ethiopian",
			},
		},
	},
	{
		Alpha2: "ZA",
		Alpha3: "ZAF",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Republic Südafrika",
				Official:    "Republik Südafrika",
				Nationality: "Südafrikanisch",
			},
			"en": {
				Common:      "South Africa",
				Official:    "Republic of South Africa",
				Nationality: "South African",
			},
		},
	},
	{
		Alpha2: "GS",
		Alpha3: "SGS",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Südgeorgien und die Südlichen Sandwichinseln",
				Official:    "Südgeorgien und die Südlichen Sandwichinseln",
				Nationality: "Südgeorgisch und Südliche Sandwichinseln",
			},
			"en": {
				Common:      "South Georgia",
				Official:    "South Georgia and the South Sandwich Islands",
				Nationality: "South Georgia and the South Sandwich Islands",
			},
		},
	},
	{
		Alpha2: "JM",
		Alpha3: "JAM",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Jamaika",
				Official:    "Jamaika",
				Nationality: "Jamaikanisch",
			},
			"en": {
				Common:      "Jamaica",
				Official:    "Jamaica",
				Nationality: "Jamaican",
			},
		},
	},
	{
		Alpha2: "ME",
		Alpha3: "MNE",
		Translations: map[Language]Translation{
			"de": {
				Official:    "Montenegro",
				Common:      "Montenegro",
				Nationality: "Montenegrinisch",
			},
			"en": {
				Common:      "Montenegro",
				Official:    "Montenegro",
				Nationality: "Montenegrin",
			},
		},
	},
	{
		Alpha2: "BL",
		Alpha3: "BLM",
		Translations: map[Language]Translation{
			"de": {
				Official:    "Gebietskörperschaft Saint -Barthélemy",
				Common:      "Saint-Barthélemy",
				Nationality: "Saint-Barthélemois",
			},
			"en": {
				Common:      "Saint Barthélemy",
				Official:    "Collectivity of Saint Barthélemy",
				Nationality: "Saint-Barthélemois",
			},
		},
	},
	{
		Alpha2: "SX",
		Alpha3: "SXM",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Sint Maarten",
				Official:    "Sint Maarten",
				Nationality: "Sint Maartener",
			},
			"en": {
				Common:      "Sint Maarten",
				Official:    "Sint Maarten",
				Nationality: "Sint Maartener",
			},
		},
	},
	{
		Alpha2: "RS",
		Alpha3: "SRB",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Serbien",
				Official:    "Republik Serbien",
				Nationality: "Serbisch",
			},
			"en": {
				Common:      "Serbia",
				Official:    "Republic of Serbia",
				Nationality: "Serbian",
			},
		},
	},
	{
		Alpha2: "AX",
		Alpha3: "ALA",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Åland",
				Official:    "Åland-Inseln",
				Nationality: "Ålandisch",
			},
			"en": {
				Common:      "Åland Islands",
				Official:    "Åland Islands",
				Nationality: "Ålandish",
			},
		},
	},
	{
		Alpha2: "BQ",
		Alpha3: "BES",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Karibische Niederlande",
				Official:    "Bonaire, Sint Eustatius und Saba",
				Nationality: "Bonaire, Sint Eustatius und Sabaner",
			},
			"en": {
				Common:      "Caribbean Netherlands",
				Official:    "Bonaire, Sint Eustatius and Saba",
				Nationality: "Bonaire, Sint Eustatius and Saban",
			},
		},
	},
	{
		Alpha2: "GG",
		Alpha3: "GGY",
		Translations: map[Language]Translation{
			"de": {
				Official:    "Guernsey",
				Common:      "Guernsey",
				Nationality: "Guernseyer",
			},
			"en": {
				Common:      "Guernsey",
				Official:    "Bailiwick of Guernsey",
				Nationality: "Guernsey",
			},
		},
	},
	{
		Alpha2: "JE",
		Alpha3: "JEY",
		Translations: map[Language]Translation{
			"de": {
				Official:    "Vogtei Jersey",
				Common:      "Jersey",
				Nationality: "Jerseyer",
			},
			"en": {
				Common:      "Jersey",
				Official:    "Bailiwick of Jersey",
				Nationality: "Jersey",
			},
		},
	},
	{
		Alpha2: "CW",
		Alpha3: "CUW",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Curaçao",
				Official:    "Curaçao",
				Nationality: "Curaçaoer",
			},
			"en": {
				Common:      "Curaçao",
				Official:    "Country of Curaçao",
				Nationality: "Curaçaoan",
			},
		},
	},
	{
		Alpha2: "MF",
		Alpha3: "MAF",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Saint Martin",
				Official:    "St. Martin",
				Nationality: "Saint-Martiner",
			},
			"en": {
				Common:      "Saint Martin",
				Official:    "Saint Martin",
				Nationality: "Saint-Martinois",
			},
		},
	},
	{
		Alpha2: "SS",
		Alpha3: "SSD",
		Translations: map[Language]Translation{
			"de": {
				Official:    "Republik Südsudan",
				Common:      "Südsudan",
				Nationality: "Südsudanesisch",
			},
			"en": {
				Common:      "South Sudan",
				Official:    "Republic of South Sudan",
				Nationality: "South Sudanese",
			},
		},
	},
	{
		Alpha2: "JP",
		Alpha3: "JPN",
		Translations: map[Language]Translation{
			"de": {
				Common:      "Japan",
				Official:    "Japan",
				Nationality: "Japanisch",
			},
			"en": {
				Common:      "Japan",
				Official:    "Japan",
				Nationality: "Japanese",
			},
		},
	},
}
