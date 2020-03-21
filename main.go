package main

import "github.com/godbus/dbus"
import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Covid struct {
	Global struct {
		Afghanistan struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Afghanistan"`
		Albania struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Albania"`
		Algeria struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Algeria"`
		Andorra struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Andorra"`
		Angola struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Angola"`
		AntiguaAndBarbuda struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Antigua and Barbuda"`
		Argentina struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Argentina"`
		Armenia struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Armenia"`
		Australia struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Australia"`
		Austria struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Austria"`
		Azerbaijan struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Azerbaijan"`
		Bahamas struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Bahamas"`
		Bahrain struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Bahrain"`
		Bangladesh struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Bangladesh"`
		Barbados struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Barbados"`
		Belarus struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Belarus"`
		Belgium struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Belgium"`
		Benin struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Benin"`
		Bhutan struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Bhutan"`
		Bolivia struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Bolivia"`
		BosniaAndHerzegovina struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Bosnia and Herzegovina"`
		Brazil struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Brazil"`
		Brunei struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Brunei"`
		Bulgaria struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Bulgaria"`
		BurkinaFaso struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Burkina Faso"`
		Cambodia struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Cambodia"`
		Cameroon struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Cameroon"`
		Canada struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Canada"`
		CapeVerde struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Cape Verde"`
		CentralAfricanRepublic struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Central African Republic"`
		Chad struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Chad"`
		Chile struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Chile"`
		China struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"China"`
		Colombia struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Colombia"`
		CostaRica struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Costa Rica"`
		Croatia struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Croatia"`
		Cuba struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Cuba"`
		DRCongo struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"DR Congo"`
		Denmark struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Denmark"`
		Djibouti struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Djibouti"`
		DominicanRepublic struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Dominican Republic"`
		Ecuador struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Ecuador"`
		Egypt struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Egypt"`
		ElSalvador struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"El Salvador"`
		EquatorialGuinea struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Equatorial Guinea"`
		Estonia struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Estonia"`
		Eswatini struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Eswatini"`
		Ethiopia struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Ethiopia"`
		Fiji struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Fiji"`
		Finland struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Finland"`
		France struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"France"`
		Gabon struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Gabon"`
		Gambia struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Gambia"`
		Georgia struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Georgia"`
		Germany struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Germany"`
		Ghana struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Ghana"`
		Greece struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Greece"`
		Guatemala struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Guatemala"`
		Guinea struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Guinea"`
		Guyana struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Guyana"`
		Haiti struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Haiti"`
		Honduras struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Honduras"`
		Hungary struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Hungary"`
		Iceland struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Iceland"`
		India struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"India"`
		Indonesia struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Indonesia"`
		Iran struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Iran"`
		Iraq struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Iraq"`
		Ireland struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Ireland"`
		Israel struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Israel"`
		Italy struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Italy"`
		IvoryCoast struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Ivory Coast"`
		Jamaica struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Jamaica"`
		Japan struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Japan"`
		Jordan struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Jordan"`
		Kazakhstan struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Kazakhstan"`
		Kenya struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Kenya"`
		Kosovo struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Kosovo"`
		KreuzfahrtschiffeUndSonstige struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Kreuzfahrtschiffe und Sonstige"`
		Kuwait struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Kuwait"`
		Kyrgyzstan struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Kyrgyzstan"`
		Latvia struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Latvia"`
		Lebanon struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Lebanon"`
		Liberia struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Liberia"`
		Liechtenstein struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Liechtenstein"`
		Lithuania struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Lithuania"`
		Luxembourg struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Luxembourg"`
		Madagascar struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Madagascar"`
		Malaysia struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Malaysia"`
		Maldives struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Maldives"`
		Malta struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Malta"`
		Martinique struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Martinique"`
		Mauritania struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Mauritania"`
		Mauritius struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Mauritius"`
		Mexico struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Mexico"`
		Moldawien struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Moldawien"`
		Monaco struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Monaco"`
		Mongolia struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Mongolia"`
		Montenegro struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Montenegro"`
		Morocco struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Morocco"`
		Namibia struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Namibia"`
		Nepal struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Nepal"`
		Netherlands struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Netherlands"`
		NewZealand struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"New Zealand"`
		Nicaragua struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Nicaragua"`
		Niger struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Niger"`
		Nigeria struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Nigeria"`
		Nordmazedonien struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Nordmazedonien"`
		Norway struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Norway"`
		Oman struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Oman"`
		Osttimor struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Osttimor"`
		Pakistan struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Pakistan"`
		Panama struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Panama"`
		PapuaNeuguinea struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Papua-Neuguinea"`
		Paraguay struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Paraguay"`
		Peru struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Peru"`
		Philippines struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Philippines"`
		Poland struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Poland"`
		Portugal struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Portugal"`
		Qatar struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Qatar"`
		RepublicOfTheCongo struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Republic of the Congo"`
		Romania struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Romania"`
		Russia struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Russia"`
		Rwanda struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Rwanda"`
		SaintLucia struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Saint Lucia"`
		SaintVincentAndTheGrenadines struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Saint Vincent and the Grenadines"`
		SanMarino struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"San Marino"`
		SaudiArabia struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Saudi Arabia"`
		Senegal struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Senegal"`
		Serbia struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Serbia"`
		Seychelles struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Seychelles"`
		Singapore struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Singapore"`
		Slovakia struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Slovakia"`
		Slovenia struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Slovenia"`
		Somalia struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Somalia"`
		SouthAfrica struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"South Africa"`
		SouthKorea struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"South Korea"`
		Spain struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Spain"`
		SriLanka struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Sri Lanka"`
		Sudan struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Sudan"`
		Suriname struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Suriname"`
		Sweden struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Sweden"`
		Switzerland struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Switzerland"`
		Taiwan struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Taiwan"`
		Tanzania struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Tanzania"`
		Thailand struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Thailand"`
		Togo struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Togo"`
		TrinidadAndTobago struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Trinidad and Tobago"`
		Tschechien struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Tschechien"`
		Tunisia struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Tunisia"`
		Turkey struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Turkey"`
		Ukraine struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Ukraine"`
		UnitedArabEmirates struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"United Arab Emirates"`
		UnitedKingdom struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"United Kingdom"`
		UnitedStates struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"United States"`
		Uruguay struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Uruguay"`
		Uzbekistan struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Uzbekistan"`
		VaticanCity struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Vatican City"`
		Venezuela struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Venezuela"`
		Vietnam struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Vietnam"`
		Zambia struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Zambia"`
		Zimbabwe struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Zimbabwe"`
		Zypern struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Zypern"`
	} `json:"global"`
	Usa struct {
		Alabama struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Alabama"`
		Alaska struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Alaska"`
		AmerikanischeJungferninseln struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Amerikanische Jungferninseln"`
		Arizona struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Arizona"`
		Arkansas struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Arkansas"`
		California struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"California"`
		Colorado struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Colorado"`
		Connecticut struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Connecticut"`
		Delaware struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Delaware"`
		DistrictOfColumbia struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"District of Columbia"`
		Florida struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Florida"`
		Georgia struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Georgia"`
		Guam struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Guam"`
		Hawaii struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Hawaii"`
		Idaho struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Idaho"`
		Illinois struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Illinois"`
		Indiana struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Indiana"`
		Iowa struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Iowa"`
		Kansas struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Kansas"`
		Kentucky struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Kentucky"`
		KreuzfahrtschiffGrandPrincess struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Kreuzfahrtschiff "Grand Princess""`
		Louisiana struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Louisiana"`
		Maine struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Maine"`
		Maryland struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Maryland"`
		Massachusetts struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Massachusetts"`
		Michigan struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Michigan"`
		Minnesota struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Minnesota"`
		Mississippi struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Mississippi"`
		Missouri struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Missouri"`
		Montana struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Montana"`
		Nebraska struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Nebraska"`
		Nevada struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Nevada"`
		NewHampshire struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"New Hampshire"`
		NewJersey struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"New Jersey"`
		NewMexico struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"New Mexico"`
		NewYork struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"New York"`
		NorthCarolina struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"North Carolina"`
		NorthDakota struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"North Dakota"`
		Ohio struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Ohio"`
		Oklahoma struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Oklahoma"`
		Oregon struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Oregon"`
		Pennsylvania struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Pennsylvania"`
		PuertoRico struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Puerto Rico"`
		RhodeIsland struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Rhode Island"`
		SouthCarolina struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"South Carolina"`
		SouthDakota struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"South Dakota"`
		Tennessee struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Tennessee"`
		Texas struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Texas"`
		Utah struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Utah"`
		Vermont struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Vermont"`
		Virginia struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Virginia"`
		Washington struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Washington"`
		WestVirginia struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"West Virginia"`
		Wisconsin struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Wisconsin"`
		Wyoming struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Wyoming"`
	} `json:"usa"`
	Canada struct {
		Alberta struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Alberta"`
		BritishColumbia struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"British Columbia"`
		KreuzfahrtschiffGrandPrincess struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Kreuzfahrtschiff "Grand Princess""`
		Manitoba struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Manitoba"`
		NewBrunswick struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"New Brunswick"`
		NewfoundlandAndLabrador struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Newfoundland and Labrador"`
		NovaScotia struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Nova Scotia"`
		Ontario struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Ontario"`
		PrinceEdwardIsland struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Prince Edward Island"`
		Quebec struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Quebec"`
		Saskatchewan struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Saskatchewan"`
	} `json:"canada"`
	Germany struct {
		BadenWRttemberg struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Baden-Württemberg"`
		Bayern struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Bayern"`
		Berlin struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Berlin"`
		Brandenburg struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Brandenburg"`
		Bremen struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Bremen"`
		Hamburg struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Hamburg"`
		Hessen struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Hessen"`
		MecklenburgVorpommern struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Mecklenburg-Vorpommern"`
		Niedersachsen struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Niedersachsen"`
		NordrheinWestfalen struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Nordrhein-Westfalen"`
		Repatriierte struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Repatriierte"`
		RheinlandPfalz struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Rheinland-Pfalz"`
		Saarland struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Saarland"`
		Sachsen struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Sachsen"`
		SachsenAnhalt struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Sachsen-Anhalt"`
		SchleswigHolstein struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Schleswig-Holstein"`
		ThRingen struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Thüringen"`
	} `json:"germany"`
	China struct {
		Anhui struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Anhui"`
		Beijing struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Beijing"`
		Chongqing struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Chongqing"`
		Fujian struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Fujian"`
		Gansu struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Gansu"`
		Guangdong struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Guangdong"`
		Guangxi struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Guangxi"`
		Guizhou struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Guizhou"`
		Hainan struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Hainan"`
		Hebei struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Hebei"`
		Heilongjiang struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Heilongjiang"`
		Henan struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Henan"`
		HongKong struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Hong Kong"`
		Hubei struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Hubei"`
		Hunan struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Hunan"`
		Jiangsu struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Jiangsu"`
		Jiangxi struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Jiangxi"`
		Jilin struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Jilin"`
		Liaoning struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Liaoning"`
		Macau struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Macau"`
		NeiMongol struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Nei Mongol"`
		Ningxia struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Ningxia"`
		Qinghai struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Qinghai"`
		Shaanxi struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Shaanxi"`
		Shandong struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Shandong"`
		Shanghai struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Shanghai"`
		Shanxi struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Shanxi"`
		Sichuan struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Sichuan"`
		Tianjin struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Tianjin"`
		Tibet struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Tibet"`
		XinjiangUygur struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Xinjiang Uygur"`
		Yunnan struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Yunnan"`
		Zhejiang struct {
			Updated   int64 `json:"updated"`
			Confirmed int   `json:"confirmed"`
			Recovered int   `json:"recovered"`
			Deaths    int   `json:"deaths"`
		} `json:"Zhejiang"`
	} `json:"china"`
}

func main() {

	resp, _ := http.Get("https://montanaflynn.github.io/covid-19/data/current.json")
	//fmt.Println(resp)
	var in Covid

	json.NewDecoder(resp.Body).Decode(&in)
	//	fmt.Println(in)

	conn, err := dbus.ConnectSessionBus()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	obj := conn.Object("org.freedesktop.Notifications", "/org/freedesktop/Notifications")
	call := obj.Call("org.freedesktop.Notifications.Notify", 0, "", uint32(0),
		"", "Current TN Deaths", string(in.Usa.Tennessee.Deaths), []string{},
		map[string]dbus.Variant{}, int32(5000))
	if call.Err != nil {
		panic(call.Err)
	}
}
