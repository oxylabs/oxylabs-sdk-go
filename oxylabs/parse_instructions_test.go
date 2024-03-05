package oxylabs

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidateParseInstructions_EmptyArgs(t *testing.T) {
	fnNames := []string{"element_text", "length", "convert_to_float", "convert_to_int", "convert_to_str", "max", "min", "product"}
	for _, fnName := range fnNames {
		type args struct {
			instructions *map[string]interface{}
		}
		tests := []struct {
			name    string
			args    args
			wantErr bool
		}{
			{
				name: fmt.Sprintf("%s with no args", fnName),
				args: args{instructions: &map[string]interface{}{
					"_fns": []map[string]interface{}{
						{
							"_fn": fnName,
						},
					},
				}},
				wantErr: false,
			},
			{
				name: fmt.Sprintf("%s with args", fnName),
				args: args{instructions: &map[string]interface{}{
					"_fns": []map[string]interface{}{
						{
							"_fn":   fnName,
							"_args": "unnecessary arg",
						},
					},
				}},
				wantErr: true,
			},
			{
				name: fmt.Sprintf("%s with no args Fn struct", fnName),
				args: args{instructions: &map[string]interface{}{
					"_fns": []Fn{{
						Name: FnName(fnName),
					}},
				}},
				wantErr: false,
			},
			{
				name: fmt.Sprintf("%s with args Fn struct", fnName),
				args: args{instructions: &map[string]interface{}{
					"_fns": []Fn{{
						Name: FnName(fnName),
						Args: "unnecessary arg",
					}},
				}},
				wantErr: true,
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if err := ValidateParseInstructions(tt.args.instructions); (err != nil) != tt.wantErr {
					t.Errorf("ValidateParseInstructions() error = %v, wantErr %v", err, tt.wantErr)
				}
			})
		}
	}
}

func TestValidateParseInstructions_StringArrayArgs(t *testing.T) {
	fnNames := []string{"xpath", "xpath_one", "css", "css_one"}
	for _, fnName := range fnNames {
		type args struct {
			instructions *map[string]interface{}
		}
		tests := []struct {
			name    string
			args    args
			wantErr bool
		}{
			{
				name: fmt.Sprintf("%s with no args", fnName),
				args: args{instructions: &map[string]interface{}{
					"_fns": []map[string]interface{}{
						{
							"_fn": fnName,
						},
					},
				}},
				wantErr: true,
			},
			{
				name: fmt.Sprintf("%s with 1 string arg", fnName),
				args: args{instructions: &map[string]interface{}{
					"_fns": []map[string]interface{}{
						{
							"_fn":   fnName,
							"_args": []string{"1st arg"},
						},
					},
				}},
				wantErr: false,
			},
			{
				name: fmt.Sprintf("%s with 2 string args", fnName),
				args: args{instructions: &map[string]interface{}{
					"_fns": []map[string]interface{}{
						{
							"_fn":   fnName,
							"_args": []string{"1st arg", "2nd string"},
						},
					},
				}},
				wantErr: false,
			},
			{
				name: fmt.Sprintf("%s with 1 empty arg", fnName),
				args: args{instructions: &map[string]interface{}{
					"_fns": []map[string]interface{}{
						{
							"_fn":   fnName,
							"_args": []string{""},
						},
					},
				}},
				wantErr: true,
			},
			{
				name: fmt.Sprintf("%s with invalid type args", fnName),
				args: args{instructions: &map[string]interface{}{
					"_fns": []map[string]interface{}{
						{
							"_fn":   fnName,
							"_args": []interface{}{"string", 1},
						},
					},
				}},
				wantErr: true,
			},
			{
				name: fmt.Sprintf("%s with empty list of args", fnName),
				args: args{instructions: &map[string]interface{}{
					"_fns": []map[string]interface{}{
						{
							"_fn":   fnName,
							"_args": []string{},
						},
					},
				}},
				wantErr: true,
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if err := ValidateParseInstructions(tt.args.instructions); (err != nil) != tt.wantErr {
					t.Errorf("ValidateParseInstructions() error = %v, wantErr %v", err, tt.wantErr)
				}
			})
		}
	}
}

func TestValidateParseInstructions_StringArg(t *testing.T) {
	fnNames := []string{"amount_from_string", "amount_range_from_string", "regex_find_all"}
	for _, fnName := range fnNames {
		type args struct {
			instructions *map[string]interface{}
		}
		tests := []struct {
			name    string
			args    args
			wantErr bool
		}{
			{
				name: fmt.Sprintf("%s with no args", fnName),
				args: args{instructions: &map[string]interface{}{
					"_fns": []map[string]interface{}{
						{
							"_fn": fnName,
						},
					},
				}},
				wantErr: true,
			},
			{
				name: fmt.Sprintf("%s with 1 string arg", fnName),
				args: args{instructions: &map[string]interface{}{
					"_fns": []map[string]interface{}{
						{
							"_fn":   fnName,
							"_args": "1st arg",
						},
					},
				}},
				wantErr: false,
			},
			{
				name: fmt.Sprintf("%s with empty string arg", fnName),
				args: args{instructions: &map[string]interface{}{
					"_fns": []map[string]interface{}{
						{
							"_fn":   fnName,
							"_args": "",
						},
					},
				}},
				wantErr: true,
			},
			{
				name: fmt.Sprintf("%s with int arg", fnName),
				args: args{instructions: &map[string]interface{}{
					"_fns": []map[string]interface{}{
						{
							"_fn":   fnName,
							"_args": 1,
						},
					},
				}},
				wantErr: true,
			},
			{
				name: fmt.Sprintf("%s with invalid type args", fnName),
				args: args{instructions: &map[string]interface{}{
					"_fns": []map[string]interface{}{
						{
							"_fn":   fnName,
							"_args": []interface{}{"string", 1},
						},
					},
				}},
				wantErr: true,
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if err := ValidateParseInstructions(tt.args.instructions); (err != nil) != tt.wantErr {
					t.Errorf("ValidateParseInstructions() error = %v, wantErr %v", err, tt.wantErr)
				}
			})
		}
	}
}

func TestValidateParseInstructions_OptionalStringArgs(t *testing.T) {
	fnNames := []string{"join"}
	for _, fnName := range fnNames {
		type args struct {
			instructions *map[string]interface{}
		}
		tests := []struct {
			name    string
			args    args
			wantErr bool
		}{
			{
				name: fmt.Sprintf("%s with no args", fnName),
				args: args{instructions: &map[string]interface{}{
					"_fns": []map[string]interface{}{
						{
							"_fn": fnName,
						},
					},
				}},
				wantErr: false,
			},
			{
				name: fmt.Sprintf("%s with 1 string arg", fnName),
				args: args{instructions: &map[string]interface{}{
					"_fns": []map[string]interface{}{
						{
							"_fn":   fnName,
							"_args": "1st arg",
						},
					},
				}},
				wantErr: false,
			},
			{
				name: fmt.Sprintf("%s with empty string arg", fnName),
				args: args{instructions: &map[string]interface{}{
					"_fns": []map[string]interface{}{
						{
							"_fn":   fnName,
							"_args": "",
						},
					},
				}},
				wantErr: false,
			},
			{
				name: fmt.Sprintf("%s with int arg", fnName),
				args: args{instructions: &map[string]interface{}{
					"_fns": []map[string]interface{}{
						{
							"_fn":   fnName,
							"_args": 1,
						},
					},
				}},
				wantErr: true,
			},
			{
				name: fmt.Sprintf("%s with invalid type args", fnName),
				args: args{instructions: &map[string]interface{}{
					"_fns": []map[string]interface{}{
						{
							"_fn":   fnName,
							"_args": []interface{}{"string", 1},
						},
					},
				}},
				wantErr: true,
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if err := ValidateParseInstructions(tt.args.instructions); (err != nil) != tt.wantErr {
					t.Errorf("ValidateParseInstructions() error = %v, wantErr %v", err, tt.wantErr)
				}
			})
		}
	}
}

func TestValidateParseInstructions_ListStringOptionalIntArgs(t *testing.T) {
	fnNames := []string{"regex_search", "regex_substring"}
	for _, fnName := range fnNames {
		type args struct {
			instructions *map[string]interface{}
		}
		tests := []struct {
			name    string
			args    args
			wantErr bool
		}{
			{
				name: fmt.Sprintf("%s with no args", fnName),
				args: args{instructions: &map[string]interface{}{
					"_fns": []map[string]interface{}{
						{
							"_fn": fnName,
						},
					},
				}},
				wantErr: true,
			},
			{
				name: fmt.Sprintf("%s with 1 string arg", fnName),
				args: args{instructions: &map[string]interface{}{
					"_fns": []map[string]interface{}{
						{
							"_fn":   fnName,
							"_args": "1st arg",
						},
					},
				}},
				wantErr: true,
			},
			{
				name: fmt.Sprintf("%s with 1 string arg in a list", fnName),
				args: args{instructions: &map[string]interface{}{
					"_fns": []map[string]interface{}{
						{
							"_fn":   fnName,
							"_args": []interface{}{"1st arg"},
						},
					},
				}},
				wantErr: false,
			},
			{
				name: fmt.Sprintf("%s with empty string arg", fnName),
				args: args{instructions: &map[string]interface{}{
					"_fns": []map[string]interface{}{
						{
							"_fn":   fnName,
							"_args": "",
						},
					},
				}},
				wantErr: true,
			},
			{
				name: fmt.Sprintf("%s with int arg", fnName),
				args: args{instructions: &map[string]interface{}{
					"_fns": []map[string]interface{}{
						{
							"_fn":   fnName,
							"_args": 1,
						},
					},
				}},
				wantErr: true,
			},
			{
				name: fmt.Sprintf("%s with valid string and int args", fnName),
				args: args{instructions: &map[string]interface{}{
					"_fns": []map[string]interface{}{
						{
							"_fn":   fnName,
							"_args": []interface{}{"string", 1},
						},
					},
				}},
				wantErr: false,
			},
			{
				name: fmt.Sprintf("%s with invalid type args", fnName),
				args: args{instructions: &map[string]interface{}{
					"_fns": []map[string]interface{}{
						{
							"_fn":   fnName,
							"_args": []interface{}{"string", "1"},
						},
					},
				}},
				wantErr: true,
			},
			{
				name: fmt.Sprintf("%s with invalid list of args", fnName),
				args: args{instructions: &map[string]interface{}{
					"_fns": []map[string]interface{}{
						{
							"_fn":   fnName,
							"_args": []interface{}{1, "1"},
						},
					},
				}},
				wantErr: true,
			},
			{
				name: fmt.Sprintf("%s with invalid list of args with empty string", fnName),
				args: args{instructions: &map[string]interface{}{
					"_fns": []map[string]interface{}{
						{
							"_fn":   fnName,
							"_args": []interface{}{""},
						},
					},
				}},
				wantErr: true,
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if err := ValidateParseInstructions(tt.args.instructions); (err != nil) != tt.wantErr {
					t.Errorf("ValidateParseInstructions() error = %v, wantErr %v", err, tt.wantErr)
				}
			})
		}
	}
}

func TestValidateParseInstructions_IntArg(t *testing.T) {
	fnNames := []string{"select_nth"}
	for _, fnName := range fnNames {
		type args struct {
			instructions *map[string]interface{}
		}
		tests := []struct {
			name    string
			args    args
			wantErr bool
		}{
			{
				name: fmt.Sprintf("%s with no args", fnName),
				args: args{instructions: &map[string]interface{}{
					"_fns": []map[string]interface{}{
						{
							"_fn": fnName,
						},
					},
				}},
				wantErr: true,
			},
			{
				name: fmt.Sprintf("%s with 1 string arg", fnName),
				args: args{instructions: &map[string]interface{}{
					"_fns": []map[string]interface{}{
						{
							"_fn":   fnName,
							"_args": "1st arg",
						},
					},
				}},
				wantErr: true,
			},
			{
				name: fmt.Sprintf("%s with empty string arg", fnName),
				args: args{instructions: &map[string]interface{}{
					"_fns": []map[string]interface{}{
						{
							"_fn":   fnName,
							"_args": "",
						},
					},
				}},
				wantErr: true,
			},
			{
				name: fmt.Sprintf("%s with int arg", fnName),
				args: args{instructions: &map[string]interface{}{
					"_fns": []map[string]interface{}{
						{
							"_fn":   fnName,
							"_args": 1,
						},
					},
				}},
				wantErr: false,
			},
			{
				name: fmt.Sprintf("%s with 0 int arg", fnName),
				args: args{instructions: &map[string]interface{}{
					"_fns": []map[string]interface{}{
						{
							"_fn":   fnName,
							"_args": 0,
						},
					},
				}},
				wantErr: true,
			},
			{
				name: fmt.Sprintf("%s with invalid type args", fnName),
				args: args{instructions: &map[string]interface{}{
					"_fns": []map[string]interface{}{
						{
							"_fn":   fnName,
							"_args": []interface{}{"string", 1},
						},
					},
				}},
				wantErr: true,
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if err := ValidateParseInstructions(tt.args.instructions); (err != nil) != tt.wantErr {
					t.Errorf("ValidateParseInstructions() error = %v, wantErr %v", err, tt.wantErr)
				}
			})
		}
	}
}

func TestValidateParseInstructions_OptionalIntArg(t *testing.T) {
	fnNames := []string{"average"}
	for _, fnName := range fnNames {
		type args struct {
			instructions *map[string]interface{}
		}
		tests := []struct {
			name    string
			args    args
			wantErr bool
		}{
			{
				name: fmt.Sprintf("%s with no args", fnName),
				args: args{instructions: &map[string]interface{}{
					"_fns": []map[string]interface{}{
						{
							"_fn": fnName,
						},
					},
				}},
				wantErr: false,
			},
			{
				name: fmt.Sprintf("%s with 1 string arg", fnName),
				args: args{instructions: &map[string]interface{}{
					"_fns": []map[string]interface{}{
						{
							"_fn":   fnName,
							"_args": "1st arg",
						},
					},
				}},
				wantErr: true,
			},
			{
				name: fmt.Sprintf("%s with empty string arg", fnName),
				args: args{instructions: &map[string]interface{}{
					"_fns": []map[string]interface{}{
						{
							"_fn":   fnName,
							"_args": "",
						},
					},
				}},
				wantErr: true,
			},
			{
				name: fmt.Sprintf("%s with int arg", fnName),
				args: args{instructions: &map[string]interface{}{
					"_fns": []map[string]interface{}{
						{
							"_fn":   fnName,
							"_args": 0,
						},
					},
				}},
				wantErr: false,
			},
			{
				name: fmt.Sprintf("%s with invalid type args", fnName),
				args: args{instructions: &map[string]interface{}{
					"_fns": []map[string]interface{}{
						{
							"_fn":   fnName,
							"_args": []interface{}{"string", 1},
						},
					},
				}},
				wantErr: true,
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if err := ValidateParseInstructions(tt.args.instructions); (err != nil) != tt.wantErr {
					t.Errorf("ValidateParseInstructions() error = %v, wantErr %v", err, tt.wantErr)
				}
			})
		}
	}
}

func TestValidateParseInstructions_NilInstructions(t *testing.T) {
	err := ValidateParseInstructions(nil)
	assert.Error(t, err)
}

func TestValidateParseInstructions_MissingFn(t *testing.T) {
	err := ValidateParseInstructions(&map[string]interface{}{
		"element": "invalid item",
	})
	assert.Error(t, err)
}

func TestValidateParseInstructions_InvalidInstructions(t *testing.T) {
	err := ValidateParseInstructions(&map[string]interface{}{
		"element": map[string]interface{}{
			"_fns": []map[string]interface{}{
				{
					"element": "invalid item",
				},
			},
		},
	})
	assert.Error(t, err)
}

func TestValidateParseInstructions_EmptyFn(t *testing.T) {
	err := ValidateParseInstructions(&map[string]interface{}{
		"element": map[string]interface{}{
			"_fns": []map[string]interface{}{
				{
					"_fn": "",
				},
			},
		},
	})
	assert.Error(t, err)
}

func TestValidateParseInstructions_NilFns(t *testing.T) {
	err := ValidateParseInstructions(&map[string]interface{}{
		"element": map[string]interface{}{
			"_fns": nil,
		},
	})
	assert.Error(t, err)
}

func TestValidateParseInstructions_InvalidFnType(t *testing.T) {
	err := ValidateParseInstructions(&map[string]interface{}{
		"element": map[string]interface{}{
			"_fns": []map[string]interface{}{
				{
					"_fn": 1,
				},
			},
		},
	})
	assert.Error(t, err)
}

func TestValidateParseInstructions_InvalidFnsFormat(t *testing.T) {
	err := ValidateParseInstructions(&map[string]interface{}{
		"element": map[string]interface{}{
			"_fns": map[string]interface{}{},
		},
	})
	assert.Error(t, err)
}
