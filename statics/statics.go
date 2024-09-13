package statics

var ClubNames = map[string]string{
	"PSG": "Paris Saint-Germain",
	"ATM": "Atletico Madrid",
	"BAY": "Bayer Leverkusen",
	"FUL": "Fulham",
	"SEV": "Sevilla",
	"RBL": "RB Leipzig",
	"LAZ": "Lazio",
	"STR": "Strasbourg",
	"EVE": "Everton",
	"MET": "Metz",
	"CHE": "Chelsea",
	"BET": "Real Betis",
	"CRY": "Crystal Palace",
	"HOF": "Hoffenheim",
	"NEW": "Newcastle United",
	"CEL": "Celta Vigo",
	"BMG": "Borussia M.Gladbach",
	"BUR": "Burnley",
	"SOU": "Southampton",
	"FRI": "Freiburg",
	"ARS": "Arsenal",
	"SAS": "Sassuolo",
	"LEI": "Leicester",
	"JUV": "Juventus",
	"VIL": "Villarreal",
	"MCI": "Manchester City",
	"TOT": "Tottenham",
	"ATA": "Atalanta",
	"LEE": "Leeds",
	"LEN": "Lens",
	"LIV": "Liverpool",
	"BAR": "Barcelona",
	"STU": "VfB Stuttgart",
	"SOC": "Real Sociedad",
	"SAM": "Sampdoria",
	"TOR": "Torino",
	"BHA": "Brighton",
	"INT": "Inter",
	"WVW": "Wolverhampton Wanderers",
	"NIC": "Nice",
	"STE": "Saint-Etienne",
	"RMA": "Real Madrid",
	"MIL": "AC Milan",
	"REN": "Rennes",
	"VAL": "Valencia",
	"REI": "Reims",
	"NAP": "Napoli",
	"LIL": "Lille",
	"WOL": "Wolfsburg",
	"EIN": "Eintracht Frankfurt",
	"MON": "Monaco",
	"BRE": "Brest",
	"WHU": "West Ham",
	"MUN": "Manchester United",
	"AVL": "Aston Villa",
	"DOR": "Borussia Dortmund",
	"LYO": "Lyon",
	"ROM": "Roma",
	"MAR": "Marseille",
	"VER": "Verona",
	"MPL": "Montpellier",
	"NAN": "Nantes",
	"UNB": "Union Berlin",
	"GEN": "Genoa",
	"MAI": "Mainz 05",
	"CAG": "Cagliari",
	"BOR": "Bordeaux",
	"SHU": "Sheffield United",
	"ATH": "Athletic Bilbao",
	"UDI": "Udinese",
	"BOL": "Bologna",
	"SCH": "Schalke 04",
	"WER": "Werder Bremen",
	"NIM": "Nimes",
	"ELC": "Elche",
	"ARM": "Arminia Bielefeld",
	"EIB": "Eibar",
	"CRO": "Crotone",
	"ALA": "Deportivo Alaves",
	"FIO": "Fiorentina",
	"AUG": "Augsburg",
	"HUE": "SD Huesca",
	"KOL": "FC Koln",
	"GET": "Getafe",
	"VAD": "Real Valladolid",
	"LOR": "Lorient",
	"LEV": "Levante",
	"HER": "Hertha Berlin",
	"WBA": "West Bromwich Albion",
	"PAR": "Parma Calcio 1913",
	"BEN": "Benevento",
	"ANG": "Angers",
	"GRA": "Granada",
	"DIJ": "Dijon",
	"OSA": "Osasuna",
	"SPE": "Spezia",
	"CAD": "Cadiz",
}

var ClubTournaments = map[string]string{
	"BUND": "Bundesliga",
	"LIG1": "Ligue 1",
	"SER":  "Serie A",
	"PREM": "Premier League",
	"LAL":  "LaLiga",
}

var SortVals = map[string]string{
	"Gls":   "Goals",
	"Possn": "Possession",
	"Pass":  "Pass",
	"Rat":   "Rating",
}

var SortValsStad = map[string]string{
	"Cap": "Capacity",
	"Pop": "Population",
}

var StadCountry = map[string]string{
	"Germany":                  "GER",
	"Norway":                   "NOR",
	"Nigeria":                  "NGA",
	"England":                  "ENG",
	"Ireland":                  "IRL",
	"South Korea":              "KOR",
	"Cambodia":                 "CAM",
	"Belarus":                  "BLR",
	"Finland":                  "FIN",
	"Colombia":                 "COL",
	"Iceland":                  "ISL",
	"Mexico":                   "MEX",
	"Latvia":                   "LVA",
	"Uruguay":                  "URU",
	"China":                    "CHN",
	"Gibraltar":                "GIB",
	"Uzbekistan":               "UZB",
	"Poland":                   "POL",
	"Scotland":                 "SCO",
	"Bosnia-Herzegovina":       "BIH",
	"Ukraine":                  "UKR",
	"United States of America": "USA",
	"Tunisia":                  "TUN",
	"Jordan":                   "JOR",
	"Austria":                  "AUT",
	"Iran":                     "IRN",
	"Armenia":                  "ARM",
	"Turkmenistan":             "TKM",
	"Ecuador":                  "ECU",
	"Turkey":                   "TUR",
	"Gambia":                   "GAM",
	"Oman":                     "OMA",
	"Morocco":                  "MAR",
	"Chile":                    "CHI",
	"Italy":                    "ITA",
	"Israel":                   "ISR",
	"Lithuania":                "LTU",
	"Equatorial Guinea":        "EQG",
	"North Korea":              "PKR",
	"United Arab Emirates":     "UAE",
	"Malaysia":                 "MAS",
	"Hong Kong":                "HKG",
	"Laos":                     "LAO",
	"Luxemburg":                "LUX",
	"Lebanon":                  "LIB",
	"Greece":                   "GRE",
	"Honduras":                 "HON",
	"Bolivia":                  "BOL",
	"Costa Rica":               "CRC",
	"Estonia":                  "EST",
	"New Zeland":               "NZL",
	"Brazil":                   "BRA",
	"Saudi Arabia":             "KSA",
	"Belgium":                  "BEL",
	"Czech Republic":           "CZE",
	"Albania":                  "ALB",
	"Burkina Faso":             "BFA",
	"Ivory Coast":              "CIV",
	"Kuwait":                   "KUW",
	"Argentina":                "ARG",
	"Gabon":                    "GAB",
	"Nepal":                    "NEP",
	"Greenland":                "GRL",
	"Malawi":                   "MWI",
	"Faroe Islands":            "FRO",
	"Taiwan":                   "TPE",
	"Iraq":                     "IRQ",
	"Japan":                    "JPN",
	"Northern Ireland":         "NIR",
	"Peru":                     "PER",
	"Moldova":                  "MDA",
	"Algeria":                  "ALG",
	"Guatemala":                "GUA",
	"Netherlands":              "NED",
	"Syria":                    "SYR",
	"Hungary":                  "HUN",
	"Venezuela":                "VEN",
	"India":                    "IND",
	"Montenegro":               "MNE",
	"Trinidad and Tobago":      "TRI",
	"Republic of South Africa": "RSA",
	"Singapore":                "SIN",
	"Cyprus":                   "CYP",
	"Azerbaijan":               "AZE",
	"Georgia":                  "GEO",
	"Kenya":                    "KEN",
	"Burma":                    "MYA",
	"Denmark":                  "DEN",
	"Thailand":                 "THA",
	"Portugal":                 "POR",
	"Bulgaria":                 "BUL",
	"Mozambique":               "MOZ",
	"Tanzania":                 "TAN",
	"Bangladesh":               "BAN",
	"Slovakia":                 "SVK",
	"Croatia":                  "CRO",
	"Wales":                    "WAL",
	"Uganda":                   "UGA",
	"Rwanda":                   "RWA",
	"Eritrea":                  "ERI",
	"Vietnam":                  "VIE",
	"Romania":                  "ROU",
	"Egypt":                    "EGY",
	"Mali":                     "MLI",
	"Australia":                "AUS",
	"Qatar":                    "QAT",
	"Angola":                   "ANG",
	"Cameroon":                 "CMR",
	"Botswana":                 "BOT",
	"France":                   "FRA",
	"Senegal":                  "SEN",
	"Lesotho":                  "LES",
	"Spain":                    "ESP",
	"Russia":                   "RUS",
	"Kazakhstan":               "KAZ",
	"Sweden":                   "SWE",
	"Ghana":                    "GHA",
	"Congo":                    "CGO",
	"Brunei":                   "BRU",
	"Paraguay":                 "PAR",
	"Serbia":                   "SER",
	"Slovenia":                 "SVN",
	"Macedonia":                "MKD",
	"Switzerland":              "SUI",
	"Malta":                    "MLT",
	"Butan":                    "BHU",
	"Canada":                   "CAN",
	"Burundi":                  "BDI",
	"Indonesia":                "IDN",
}
