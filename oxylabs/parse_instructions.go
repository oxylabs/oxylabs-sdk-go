package oxylabs

import "fmt"

type FnName string

const (
	ElementText FnName = "element_text"
	Xpath       FnName = "xpath"
	XpathOne    FnName = "xpath_one"
	Css         FnName = "css"
	CssOne      FnName = "css_one"

	AmountFromString      FnName = "amount_from_string"
	AmountRangeFromString FnName = "amount_range_from_string"
	Join                  FnName = "join"
	RegexFindAll          FnName = "regex_find_all"
	RegexSearch           FnName = "regex_search"
	RegexSubstring        FnName = "regex_substring"

	Length         FnName = "length"
	SelectNth      FnName = "select_nth"
	ConvertToFloat FnName = "convert_to_float"
	ConvertToInt   FnName = "convert_to_int"
	ConvertToStr   FnName = "convert_to_str"

	Average FnName = "average"
	Max     FnName = "max"
	Min     FnName = "min"
	Product FnName = "product"
)

type Fn struct {
	Name FnName `json:"_fn"`
	Args any    `json:"_args,omitempty"`
}

func ValidateParseInstructions(instructions *map[string]interface{}) error {
	if instructions == nil {
		return fmt.Errorf("parse instructions cannot be nil")
	}

	for k, v := range *instructions {
		switch k {
		case "_fns":
			if err := validateFns(v); err != nil {
				return err
			}
		default:
			vv, ok := v.(map[string]interface{})
			if !ok {
				return fmt.Errorf("invalid parse instructions format")
			}
			if err := ValidateParseInstructions(&vv); err != nil {
				return err
			}
		}
	}

	return nil
}

func validateFns(fns interface{}) error {
	if fns == nil {
		return fmt.Errorf("_fns cannot be nil")
	}
	switch v := fns.(type) {
	case []Fn:
		for _, f := range v {
			if err := validateFn(f); err != nil {
				return err
			}
		}
	case []map[string]interface{}:
		for _, f := range v {
			fnName, ok := f["_fn"]
			if !ok {
				return fmt.Errorf("_fn must be set")
			}
			switch v := fnName.(type) {
			case string:
				if err := validateFn(Fn{Name: FnName(v), Args: f["_args"]}); err != nil {
					return err
				}
			case FnName:
				if err := validateFn(Fn{Name: v, Args: f["_args"]}); err != nil {
					return err
				}
			default:
				return fmt.Errorf("_fn must be string")
			}
		}
	default:
		return fmt.Errorf("invalid _fns format")
	}

	return nil
}

func validateFn(fn Fn) error {
	var err error
	switch fn.Name {
	case "":
		return fmt.Errorf("_fn cannot be empty")
	case ElementText, Length, ConvertToFloat, ConvertToInt, ConvertToStr, Max, Min, Product:
		err = validateEmpty(fn.Args)
	case Xpath, XpathOne, Css, CssOne:
		err = validateStringArray(fn.Args)
	case AmountFromString, AmountRangeFromString, RegexFindAll:
		err = validateString(fn.Args)
	case Join:
		err = validateOptionalString(fn.Args)
	case RegexSearch, RegexSubstring:
		err = validateListStringOptionalInt(fn.Args)
	case SelectNth:
		err = validateNonZeroInt(fn.Args)
	case Average:
		err = validateOptionalInt(fn.Args)
	}

	if err != nil {
		return fmt.Errorf("_fn %s invalid: %w", fn.Name, err)
	}
	return nil
}

func isEmpty(args any) bool {
	return args == nil
}

func validateEmpty(args any) error {
	if !isEmpty(args) {
		return fmt.Errorf("_args must be empty")
	}

	return nil
}

func validateStringArray(args any) error {
	a, ok := args.([]string)
	if !ok {
		return fmt.Errorf("_args must be of type []string")
	}
	if len(a) < 1 {
		return fmt.Errorf("_args cannot be empty")
	}
	for _, e := range a {
		if e == "" {
			return fmt.Errorf("_args cannot have empty elements")
		}
	}

	return nil
}

func validateString(args any) error {
	arg, ok := args.(string)
	if !ok {
		return fmt.Errorf("_args must be of type string")
	}
	if arg == "" {
		return fmt.Errorf("_args cannot be empty")
	}

	return nil
}

func validateOptionalString(args any) error {
	if isEmpty(args) {
		return nil
	}

	_, ok := args.(string)
	if !ok {
		return fmt.Errorf("_args must be of type string")
	}

	return nil
}

func validateNonZeroInt(args any) error {
	arg, ok := args.(int)
	if !ok {
		return fmt.Errorf("_args must be of type int")
	}
	if arg == 0 {
		return fmt.Errorf("_args cannot be 0")
	}

	return nil
}

func validateOptionalInt(args any) error {
	if isEmpty(args) {
		return nil
	}

	_, ok := args.(int)
	if !ok {
		return fmt.Errorf("_args must be of type int")
	}

	return nil
}

func validateListStringOptionalInt(args any) error {
	a, ok := args.([]any)
	if !ok {
		return fmt.Errorf("_args must be non empty list of arguments")
	}

	arg, ok := a[0].(string)
	if !ok {
		return fmt.Errorf("_args first argument must be string")
	}
	if arg == "" {
		return fmt.Errorf("_args first argument cannot be empty")
	}

	if len(a) < 2 {
		return nil
	}

	_, ok = a[1].(int)
	if !ok {
		return fmt.Errorf("_args second argument must be int")
	}

	return nil
}
