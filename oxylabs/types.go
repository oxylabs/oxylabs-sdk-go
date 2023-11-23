package oxylabs

type UserAgent string

var (
	UA_DESKTOP         UserAgent = "desktop"
	UA_DESKTOP_CHROME  UserAgent = "desktop_chrome"
	UA_DESKTOP_EDGE    UserAgent = "desktop_edge"
	UA_DESKTOP_FIREFOX UserAgent = "desktop_firefox"
	UA_DESKTOP_OPERA   UserAgent = "desktop_opera"
	UA_DESKTOP_SAFARI  UserAgent = "desktop_safari"
	UA_MOBILE          UserAgent = "mobile"
	UA_MOBILE_ANDROID  UserAgent = "mobile_android"
	UA_MOBILE_IOS      UserAgent = "mobile_ios"
	UA_TABLET          UserAgent = "tablet"
	UA_TABLET_ANDROID  UserAgent = "tablet_android"
	UA_TABLET_IOS      UserAgent = "tablet_ios"
)

func IsUserAgentValid(ua string) bool {
	switch UserAgent(ua) {
	case
		UA_DESKTOP,
		UA_DESKTOP_CHROME,
		UA_DESKTOP_EDGE,
		UA_DESKTOP_FIREFOX,
		UA_DESKTOP_OPERA,
		UA_DESKTOP_SAFARI,
		UA_MOBILE,
		UA_MOBILE_ANDROID,
		UA_MOBILE_IOS,
		UA_TABLET,
		UA_TABLET_ANDROID,
		UA_TABLET_IOS:
		return true
	default:
		return false
	}
}

type Render string

var (
	HTML Render = "html"
	PNG  Render = "png"
)

func IsRenderValid(render Render) bool {
	switch render {
	case
		HTML,
		PNG:
		return true
	default:
		return false
	}
}

type Domain string

var (
	DOMAIN_COM Domain = "com"
	DOMAIN_RU  Domain = "ru"
	DOMAIN_UA  Domain = "ua"
	DOMAIN_BY  Domain = "by"
	DOMAIN_KZ  Domain = "kz"
	DOMAIN_TR  Domain = "tr"
)

type Locale string

var (
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
