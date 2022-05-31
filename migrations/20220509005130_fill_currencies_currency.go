package migrations

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"github.com/pressly/goose/v3"
	"gitlab.ozon.dev/zBlur/homework-2/internal/domain"
	"strings"
)

func init() {
	goose.AddMigration(upFillCurrencies, downFillCurrencies)
}

var fiatCurrencies = map[string]string{
	"EUR": "European euro",
	"USD": "United States dollar",
	"AFN": "Afghan afghani",
	"ALL": "Albanian lek",
	"DZD": "Algerian dinar",
	"XCD": "East Caribbean dollar",
	"AOA": "Angolan kwanza",
	"ARS": "Argentine peso",
	"AMD": "Armenian dram",
	"AWG": "Aruban florin",
	"SHP": "Saint Helena pound",
	"AUD": "Australian dollar",
	"AZN": "Azerbaijan manat",
	"BSD": "Bahamian dollar",
	"BHD": "Bahraini dinar",
	"BDT": "Bangladeshi taka",
	"BBD": "Barbadian dollar",
	"BYN": "Belarusian ruble",
	"BZD": "Belize dollar",
	"XOF": "West African CFA franc",
	"BMD": "Bermudian dollar",
	"BTN": "Bhutanese ngultrum",
	"BOB": "Bolivian boliviano",
	"BAM": "Bosnia and Herzegovina convertible mark",
	"BWP": "Botswana pula",
	"BRL": "Brazilian real",
	"BND": "Brunei dollar",
	"BGN": "Bulgarian lev",
	"CVE": "Cabo Verdean escudo",
	"BIF": "Burundi franc",
	"KHR": "Cambodian riel",
	"XAF": "Central African CFA franc",
	"CAD": "Canadian dollar",
	"KYD": "Cayman Islands dollar",
	"NZD": "New Zealand dollar",
	"CLP": "Chilean peso",
	"CNY": "Chinese Yuan Renminbi",
	"COP": "Colombian peso",
	"KMF": "Comorian franc",
	"CDF": "Congolese franc",
	"CRC": "Costa Rican colon",
	"HRK": "Croatian kuna",
	"CUP": "Cuban peso",
	"CZK": "Czech koruna",
	"ANG": "Netherlands Antillean guilder",
	"DKK": "Danish krone",
	"DJF": "Djiboutian franc",
	"DOP": "Dominican peso",
	"EGP": "Egyptian pound",
	"ERN": "Eritrean nakfa",
	"SZL": "Swazi lilangeni",
	"ETB": "Ethiopian birr",
	"FKP": "Falkland Islands pound",
	"FJD": "Fijian dollar",
	"XPF": "CFP franc",
	"GMD": "Gambian dalasi",
	"GEL": "Georgian lari",
	"GHS": "Ghanaian cedi",
	"GIP": "Gibraltar pound",
	"GTQ": "Guatemalan quetzal",
	"GGP": "Guernsey Pound",
	"GNF": "Guinean franc",
	"GYD": "Guyanese dollar",
	"HTG": "Haitian gourde",
	"HNL": "Honduran lempira",
	"HKD": "Hong Kong dollar",
	"HUF": "Hungarian forint",
	"ISK": "Icelandic krona",
	"INR": "Indian rupee",
	"IDR": "Indonesian rupiah",
	"XDR": "SDR (Special Drawing Right)",
	"IRR": "Iranian rial",
	"IQD": "Iraqi dinar",
	"IMP": "Manx pound",
	"ILS": "Israeli new shekel",
	"JMD": "Jamaican dollar",
	"JPY": "Japanese yen",
	"JEP": "Jersey pound",
	"JOD": "Jordanian dinar",
	"KZT": "Kazakhstani tenge",
	"KES": "Kenyan shilling",
	"KWD": "Kuwaiti dinar",
	"KGS": "Kyrgyzstani som",
	"LAK": "Lao kip",
	"LBP": "Lebanese pound",
	"LSL": "Lesotho loti",
	"LRD": "Liberian dollar",
	"LYD": "Libyan dinar",
	"CHF": "Swiss franc",
	"MOP": "Macanese pataca",
	"MGA": "Malagasy ariary",
	"MWK": "Malawian kwacha",
	"MYR": "Malaysian ringgit",
	"MVR": "Maldivian rufiyaa",
	"MRU": "Mauritanian ouguiya",
	"MUR": "Mauritian rupee",
	"MXN": "Mexican peso",
	"MDL": "Moldovan leu",
	"MNT": "Mongolian tugrik",
	"MAD": "Moroccan dirham",
	"MZN": "Mozambican metical",
	"MMK": "Myanmar kyat",
	"NAD": "Namibian dollar",
	"NPR": "Nepalese rupee",
	"NIO": "Nicaraguan cordoba",
	"NGN": "Nigerian naira",
	"KPW": "North Korean won",
	"MKD": "Macedonian denar",
	"NOK": "Norwegian krone",
	"OMR": "Omani rial",
	"PKR": "Pakistani rupee",
	"PGK": "Papua New Guinean kina",
	"PYG": "Paraguayan guarani",
	"PEN": "Peruvian sol",
	"PHP": "Philippine peso",
	"PLN": "Polish zloty",
	"QAR": "Qatari riyal",
	"RON": "Romanian leu",
	"RUB": "Russian ruble",
	"RWF": "Rwandan franc",
	"WST": "Samoan tala",
	"STN": "Sao Tome and Principe dobra",
	"SAR": "Saudi Arabian riyal",
	"RSD": "Serbian dinar",
	"SCR": "Seychellois rupee",
	"SLL": "Sierra Leonean leone",
	"SGD": "Singapore dollar",
	"SBD": "Solomon Islands dollar",
	"SOS": "Somali shilling",
	"ZAR": "South African rand",
	"GBP": "Pound sterling",
	"KRW": "South Korean won",
	"SSP": "South Sudanese pound",
	"LKR": "Sri Lankan rupee",
	"SDG": "Sudanese pound",
	"SRD": "Surinamese dollar",
	"SEK": "Swedish krona",
	"SYP": "Syrian pound",
	"TWD": "New Taiwan dollar",
	"TJS": "Tajikistani somoni",
	"TZS": "Tanzanian shilling",
	"THB": "Thai baht",
	"TOP": "Tongan pa’anga",
	"TTD": "Trinidad and Tobago dollar",
	"TND": "Tunisian dinar",
	"TRY": "Turkish lira",
	"TMT": "Turkmen manat",
	"UGX": "Ugandan shilling",
	"UAH": "Ukrainian hryvnia",
	"AED": "UAE dirham",
	"UYU": "Uruguayan peso",
	"UZS": "Uzbekistani som",
	"VUV": "Vanuatu vatu",
	"VES": "Venezuelan bolivar",
	"VND": "Vietnamese dong",
	"YER": "Yemeni rial",
	"ZMW": "Zambian kwacha",
}

var cryptoCurrencies = map[string]string{
	"USDT": "Tether",
	"USDC": "USD Coin",
	"BUSD": "Binance USD",
}

type Currency struct {
	Code  string `db:"code"`
	Title string `db:"title"`
	Type  string `db:"type"`
}

const CurrencyFieldsNum = 3

func upFillCurrencies(tx *sql.Tx) error {

	var placeholders []string
	var values []interface{}

	for code, title := range fiatCurrencies {
		placeholders = append(
			placeholders,
			fmt.Sprintf("($%d,$%d,$%d)",
				len(placeholders)*CurrencyFieldsNum+1,
				len(placeholders)*CurrencyFieldsNum+2,
				len(placeholders)*CurrencyFieldsNum+3,
			),
		)
		values = append(values, code, title, domain.CurrencyFiatType)
	}
	for code, title := range cryptoCurrencies {
		placeholders = append(
			placeholders,
			fmt.Sprintf("($%d,$%d,$%d)",
				len(placeholders)*CurrencyFieldsNum+1,
				len(placeholders)*CurrencyFieldsNum+2,
				len(placeholders)*CurrencyFieldsNum+3,
			),
		)
		values = append(values, code, title, domain.CurrencyCryptoType)
	}

	insertQuery := fmt.Sprintf(
		`INSERT INTO currencies_currency (Code,Title,Type) VALUES %s`,
		strings.Join(placeholders, ","),
	)

	_, err := tx.Exec(insertQuery, values...)
	if err != nil {
		return errors.Wrap(err, "failed to insert multiple records at once")
	}

	return nil
}

func downFillCurrencies(tx *sql.Tx) error {

	var placeholders []string
	var values []interface{}

	for code, _ := range fiatCurrencies {
		placeholders = append(
			placeholders,
			fmt.Sprintf("$%d", len(placeholders)+1),
		)
		values = append(values, code)
	}
	for code, _ := range cryptoCurrencies {
		placeholders = append(
			placeholders,
			fmt.Sprintf("$%d", len(placeholders)+1),
		)
		values = append(values, code)
	}

	deleteQuery := fmt.Sprintf(
		`DELETE FROM currencies_currency WHERE Code IN (%s)`,
		strings.Join(placeholders, ","),
	)

	_, err := tx.Exec(deleteQuery, values...)
	if err != nil {
		return errors.Wrap(err, "failed to delete multiple records at once")
	}

	return nil
}
