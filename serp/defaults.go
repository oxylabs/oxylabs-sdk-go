package serp

import (
	"reflect"

	"github.com/mslmio/oxylabs-sdk-go/oxylabs"
)

// Function to set default values for serp scrapers.
func SetDefaults(opt interface{}) {
	val := reflect.ValueOf(opt).Elem()

	// Loop through the fields of the struct.
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := val.Type().Field(i)

		// Set domain.
		if fieldType.Name == "Domain" && field.String() == "" {
			field.SetString(string(oxylabs.DOMAIN_COM))
		}

		// Set start page.
		if fieldType.Name == "StartPage" && field.Int() == 0 {
			field.SetInt(1)
		}

		// Set pages.
		if fieldType.Name == "Pages" && field.Int() == 0 {
			field.SetInt(1)
		}

		// Set limit.
		if fieldType.Name == "Limit" && field.Int() == 0 {
			field.SetInt(10)
		}

		// Set user agent.
		if fieldType.Name == "UserAgent" && field.String() == "" {
			field.SetString(string(oxylabs.UA_DESKTOP))
		}
	}
}
