package oxylabs

type UserAgent string

const (
	UA_MOBILE          UserAgent = "mobile"
	UA_TABLET          UserAgent = "tablet"
	UA_DESKTOP         UserAgent = "desktop"
	UA_MOBILE_IOS      UserAgent = "mobile_ios"
	UA_TABLET_IOS      UserAgent = "tablet_ios"
	UA_DESKTOP_EDGE    UserAgent = "desktop_edge"
	UA_DESKTOP_OPERA   UserAgent = "desktop_opera"
	UA_DESKTOP_SAFARI  UserAgent = "desktop_safari"
	UA_MOBILE_ANDROID  UserAgent = "mobile_android"
	UA_DESKTOP_CHROME  UserAgent = "desktop_chrome"
	UA_TABLET_ANDROID  UserAgent = "tablet_android"
	UA_DESKTOP_FIREFOX UserAgent = "desktop_firefox"
)

func IsUserAgentValid(ua UserAgent) bool {
	switch ua {
	case
		UA_MOBILE,
		UA_TABLET,
		UA_DESKTOP,
		UA_MOBILE_IOS,
		UA_TABLET_IOS,
		UA_DESKTOP_EDGE,
		UA_DESKTOP_OPERA,
		UA_DESKTOP_SAFARI,
		UA_MOBILE_ANDROID,
		UA_DESKTOP_CHROME,
		UA_TABLET_ANDROID,
		UA_DESKTOP_FIREFOX:
		return true
	default:
		return false
	}
}

type Render string

const (
	PNG  Render = "png"
	HTML Render = "html"
)

func IsRenderValid(render Render) bool {
	switch render {
	case
		PNG,
		HTML:
		return true
	default:
		return false
	}
}

type Source string

const (
	GoogleUrl           Source = "google"
	GoogleAds           Source = "google_ads"
	GoogleHotels        Source = "google_hotels"
	GoogleSearch        Source = "google_search"
	GoogleImages        Source = "google_images"
	GoogleSuggestions   Source = "google_suggest"
	GoogleTravelHotels  Source = "google_travel_hotels"
	GoogleTrendsExplore Source = "google_trends_explore"

	BingUrl    Source = "bing"
	BingSearch Source = "bing_search"

	YandexUrl    Source = "yandex"
	YandexSearch Source = "yandex_search"

	BaiduUrl    Source = "baidu"
	BaiduSearch Source = "baidu_search"

	GoogleShoppingUrl    Source = "google_shopping"
	GoogleShoppingSearch Source = "google_shopping_search"

	Wayfair       Source = "wayfair"
	WayfairSearch Source = "wayfair_search"
)

type Domain string

const (
	DOMAIN_RU       Domain = "ru"
	DOMAIN_UA       Domain = "ua"
	DOMAIN_TR       Domain = "tr"
	DOMAIN_CN       Domain = "cn"
	DOMAIN_COM_AI   Domain = "com.ai"
	DOMAIN_COM_PR   Domain = "com.pr"
	DOMAIN_SR       Domain = "sr"
	DOMAIN_ML       Domain = "ml"
	DOMAIN_COM_LB   Domain = "com.lb"
	DOMAIN_BF       Domain = "bf"
	DOMAIN_FM       Domain = "fm"
	DOMAIN_COM_MX   Domain = "com.mx"
	DOMAIN_BJ       Domain = "bj"
	DOMAIN_EE       Domain = "ee"
	DOMAIN_MV       Domain = "mv"
	DOMAIN_NE       Domain = "ne"
	DOMAIN_AT       Domain = "at"
	DOMAIN_GG       Domain = "gg"
	DOMAIN_AE       Domain = "ae"
	DOMAIN_CO_UZ    Domain = "co.uz"
	DOMAIN_AM       Domain = "am"
	DOMAIN_COM_SA   Domain = "com.sa"
	DOMAIN_TL       Domain = "tl"
	DOMAIN_COM_NA   Domain = "com.na"
	DOMAIN_COM_BH   Domain = "com.bh"
	DOMAIN_DK       Domain = "dk"
	DOMAIN_COM_SB   Domain = "com.sb"
	DOMAIN_RO       Domain = "ro"
	DOMAIN_BY       Domain = "by"
	DOMAIN_COM_CO   Domain = "com.co"
	DOMAIN_COM_GI   Domain = "com.gi"
	DOMAIN_CO_ID    Domain = "co.id"
	DOMAIN_MS       Domain = "ms"
	DOMAIN_COM_NG   Domain = "com.ng"
	DOMAIN_IS       Domain = "is"
	DOMAIN_COM_EG   Domain = "com.eg"
	DOMAIN_COM_ET   Domain = "com.et"
	DOMAIN_COM_AF   Domain = "com.af"
	DOMAIN_CH       Domain = "ch"
	DOMAIN_CO_AO    Domain = "co.ao"
	DOMAIN_CL       Domain = "cl"
	DOMAIN_CO_ZA    Domain = "co.za"
	DOMAIN_COM_NF   Domain = "com.nf"
	DOMAIN_DK_RO    Domain = "ro"
	DOMAIN_MD       Domain = "md"
	DOMAIN_ES       Domain = "es"
	DOMAIN_BJ_YO    Domain = "bj"
	DOMAIN_HU       Domain = "hu"
	DOMAIN_DJ       Domain = "dj"
	DOMAIN_COM_MT   Domain = "com.mt"
	DOMAIN_COM_EC   Domain = "com.ec"
	DOMAIN_CO_IN    Domain = "co.in"
	DOMAIN_LK       Domain = "lk"
	DOMAIN_CO_KE    Domain = "co.ke"
	DOMAIN_GY       Domain = "gy"
	DOMAIN_BE       Domain = "be"
	DOMAIN_VG       Domain = "vg"
	DOMAIN_CO_BW    Domain = "co.bw"
	DOMAIN_COM_VN   Domain = "com.vn"
	DOMAIN_CO_TZ    Domain = "co.tz"
	DOMAIN_NE_HA    Domain = "ne"
	DOMAIN_CO_ZW    Domain = "co.zw"
	DOMAIN_TO       Domain = "to"
	DOMAIN_KZ       Domain = "kz"
	DOMAIN_COM_UY   Domain = "com.uy"
	DOMAIN_IQ       Domain = "iq"
	DOMAIN_COM_TW   Domain = "com.tw"
	DOMAIN_RW       Domain = "rw"
	DOMAIN_AD       Domain = "ad"
	DOMAIN_COM_LY   Domain = "com.ly"
	DOMAIN_AL       Domain = "al"
	DOMAIN_CO_IL    Domain = "co.il"
	DOMAIN_KI       Domain = "ki"
	DOMAIN_COM      Domain = "com"
	DOMAIN_MU       Domain = "mu"
	DOMAIN_SC       Domain = "sc"
	DOMAIN_COM_HK   Domain = "com.hk"
	DOMAIN_COM_PA   Domain = "com.pa"
	DOMAIN_CA       Domain = "ca"
	DOMAIN_GE       Domain = "ge"
	DOMAIN_COM_GT   Domain = "com.gt"
	DOMAIN_LI       Domain = "li"
	DOMAIN_COM_KH   Domain = "com.kh"
	DOMAIN_CO_CR    Domain = "co.cr"
	DOMAIN_COM_BO   Domain = "com.bo"
	DOMAIN_CO_VE    Domain = "co.ve"
	DOMAIN_COM_NI   Domain = "com.ni"
	DOMAIN_TD       Domain = "td"
	DOMAIN_CF       Domain = "cf"
	DOMAIN_TK       Domain = "tk"
	DOMAIN_BI       Domain = "bi"
	DOMAIN_MG       Domain = "mg"
	DOMAIN_COM_BD   Domain = "com.bd"
	DOMAIN_COM_BZ   Domain = "com.bz"
	DOMAIN_GM       Domain = "gm"
	DOMAIN_LA       Domain = "la"
	DOMAIN_COM_KW   Domain = "com.kw"
	DOMAIN_CM       Domain = "cm"
	DOMAIN_HT       Domain = "ht"
	DOMAIN_NO       Domain = "no"
	DOMAIN_COM_FJ   Domain = "com.fj"
	DOMAIN_TM       Domain = "tm"
	DOMAIN_COM_SL   Domain = "com.sl"
	DOMAIN_COM_MM   Domain = "com.mm"
	DOMAIN_IM       Domain = "im"
	DOMAIN_SI       Domain = "si"
	DOMAIN_COM_QA   Domain = "com.qa"
	DOMAIN_COM_PE   Domain = "com.pe"
	DOMAIN_CD       Domain = "cd"
	DOMAIN_TT       Domain = "tt"
	DOMAIN_COM_TR   Domain = "com.tr"
	DOMAIN_TG       Domain = "tg"
	DOMAIN_CO_LS    Domain = "co.ls"
	DOMAIN_GR       Domain = "gr"
	DOMAIN_GL       Domain = "gl"
	DOMAIN_MK       Domain = "mk"
	DOMAIN_CO_ZM    Domain = "co.zm"
	DOMAIN_COM_PH   Domain = "com.ph"
	DOMAIN_IT       Domain = "it"
	DOMAIN_CO_JP    Domain = "co.jp"
	DOMAIN_WS       Domain = "ws"
	DOMAIN_COM_AR   Domain = "com.ar"
	DOMAIN_CO_MZ    Domain = "co.mz"
	DOMAIN_AZ       Domain = "az"
	DOMAIN_CO_CK    Domain = "co.ck"
	DOMAIN_FI       Domain = "fi"
	DOMAIN_COM_BN   Domain = "com.bn"
	DOMAIN_PT       Domain = "pt"
	DOMAIN_COM_TJ   Domain = "com.tj"
	DOMAIN_COM_CY   Domain = "com.cy"
	DOMAIN_CV       Domain = "cv"
	DOMAIN_COM_MY   Domain = "com.my"
	DOMAIN_IE       Domain = "ie"
	DOMAIN_COM_SG   Domain = "com.sg"
	DOMAIN_DE       Domain = "de"
	DOMAIN_BA       Domain = "ba"
	DOMAIN_LU       Domain = "lu"
	DOMAIN_BG       Domain = "bg"
	DOMAIN_CO_VI    Domain = "co.vi"
	DOMAIN_COM_OM   Domain = "com.om"
	DOMAIN_AS       Domain = "as"
	DOMAIN_DZ       Domain = "dz"
	DOMAIN_FR       Domain = "fr"
	DOMAIN_LV       Domain = "lv"
	DOMAIN_LT       Domain = "lt"
	DOMAIN_PS       Domain = "ps"
	DOMAIN_SE       Domain = "se"
	DOMAIN_CG       Domain = "cg"
	DOMAIN_NR       Domain = "nr"
	DOMAIN_CO_UG    Domain = "co.ug"
	DOMAIN_COM_VC   Domain = "com.vc"
	DOMAIN_JO       Domain = "jo"
	DOMAIN_CO_TH    Domain = "co.th"
	DOMAIN_RS       Domain = "rs"
	DOMAIN_BS       Domain = "bs"
	DOMAIN_COM_PK   Domain = "com.pk"
	DOMAIN_CO_UK    Domain = "co.uk"
	DOMAIN_SO       Domain = "so"
	DOMAIN_GA       Domain = "ga"
	DOMAIN_COM_UA   Domain = "com.ua"
	DOMAIN_HR       Domain = "hr"
	DOMAIN_COM_CU   Domain = "com.cu"
	DOMAIN_SK       Domain = "sk"
	DOMAIN_COM_NP   Domain = "com.np"
	DOMAIN_NU       Domain = "nu"
	DOMAIN_MN       Domain = "mn"
	DOMAIN_VU       Domain = "vu"
	DOMAIN_NL       Domain = "nl"
	DOMAIN_PT_ST    Domain = "st"
	DOMAIN_COM_BR   Domain = "com.br"
	DOMAIN_TH       Domain = "co.th"
	DOMAIN_MW       Domain = "mw"
	DOMAIN_COM_PG   Domain = "com.pg"
	DOMAIN_PL       Domain = "pl"
	DOMAIN_CO_NZ    Domain = "co.nz"
	DOMAIN_KG       Domain = "kg"
	DOMAIN_CI       Domain = "ci"
	DOMAIN_SH       Domain = "sh"
	DOMAIN_COM_DO   Domain = "com.do"
	DOMAIN_SN       Domain = "sn"
	DOMAIN_COM_JM   Domain = "com.jm"
	DOMAIN_CO_MA    Domain = "co.ma"
	DOMAIN_COM_TN   Domain = "com.tn"
	DOMAIN_DM       Domain = "dm"
	DOMAIN_COM_SV   Domain = "com.sv"
	DOMAIN_COM_SG_2 Domain = "com.sg"
	DOMAIN_GP       Domain = "gp"
	DOMAIN_ME       Domain = "me"
	DOMAIN_COM_AG   Domain = "com.ag"
	DOMAIN_CZ       Domain = "cz"
	DOMAIN_COM_PY   Domain = "com.py"
	DOMAIN_MR_IN    Domain = "co.in"
	DOMAIN_COM_GH   Domain = "com.gh"
	DOMAIN_ST_LS    Domain = "co.ls"
	DOMAIN_BT       Domain = "bt"
	DOMAIN_RU_KZ    Domain = "kz"
	DOMAIN_IT_SM    Domain = "sm"
	DOMAIN_JE       Domain = "je"
	DOMAIN_TN       Domain = "tn"
	DOMAIN_COM_AU   Domain = "com.au"
	DOMAIN_ME_ME    Domain = "me"
	DOMAIN_PN       Domain = "pn"
	DOMAIN_HN       Domain = "hn"
	DOMAIN_CO_KR    Domain = "co.kr"
	DOMAIN_AR       Domain = "com.ar"
	DOMAIN_BO       Domain = "com.bo"
	DOMAIN_BZ       Domain = "com.bz"
	DOMAIN_UY       Domain = "com.uy"
	DOMAIN_COM_VE   Domain = "com.ve"
	DOMAIN_ID_TL    Domain = "tl"
)

type Locale string

const (
	LOCALE_EN Locale = "en"
	LOCALE_RU Locale = "ru"
	LOCALE_BY Locale = "by"
	LOCALE_DE Locale = "de"
	LOCALE_FR Locale = "fr"
	LOCALE_ID Locale = "id"
	LOCALE_KK Locale = "kk"
	LOCALE_TT Locale = "tt"
	LOCALE_TR Locale = "tr"
	LOCALE_UK Locale = "uk"
)
