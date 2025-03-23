// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasource

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/greatman/terraform-plugin-codegen-spec/code"
	"github.com/greatman/terraform-plugin-codegen-spec/datasource"
	specschema "github.com/greatman/terraform-plugin-codegen-spec/schema"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/convert"
	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/model"
)

func TestGeneratorInt32Attribute_New(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         *datasource.Int32Attribute
		expected      GeneratorInt32Attribute
		expectedError error
	}{
		"nil": {
			expectedError: fmt.Errorf("*datasource.Int32Attribute is nil"),
		},
		"computed": {
			input: &datasource.Int32Attribute{
				ComputedOptionalRequired: "computed",
			},
			expected: GeneratorInt32Attribute{
				ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Computed),
				CustomType:               convert.NewCustomTypePrimitive(nil, nil, "name"),
				Validators:               convert.NewValidators(convert.ValidatorTypeInt32, specschema.CustomValidators{}),
			},
		},
		"computed_optional": {
			input: &datasource.Int32Attribute{
				ComputedOptionalRequired: "computed_optional",
			},
			expected: GeneratorInt32Attribute{
				ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.ComputedOptional),
				CustomType:               convert.NewCustomTypePrimitive(nil, nil, "name"),
				Validators:               convert.NewValidators(convert.ValidatorTypeInt32, specschema.CustomValidators{}),
			},
		},
		"optional": {
			input: &datasource.Int32Attribute{
				ComputedOptionalRequired: "optional",
			},
			expected: GeneratorInt32Attribute{
				ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
				CustomType:               convert.NewCustomTypePrimitive(nil, nil, "name"),
				Validators:               convert.NewValidators(convert.ValidatorTypeInt32, specschema.CustomValidators{}),
			},
		},
		"required": {
			input: &datasource.Int32Attribute{
				ComputedOptionalRequired: "required",
			},
			expected: GeneratorInt32Attribute{
				ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Required),
				CustomType:               convert.NewCustomTypePrimitive(nil, nil, "name"),
				Validators:               convert.NewValidators(convert.ValidatorTypeInt32, specschema.CustomValidators{}),
			},
		},
		"custom_type": {
			input: &datasource.Int32Attribute{
				CustomType: &specschema.CustomType{
					Import: &code.Import{
						Path: "github.com/",
					},
					Type:      "my_type",
					ValueType: "myvalue_type",
				},
			},
			expected: GeneratorInt32Attribute{
				CustomType: convert.NewCustomTypePrimitive(&specschema.CustomType{
					Import: &code.Import{
						Path: "github.com/",
					},
					Type:      "my_type",
					ValueType: "myvalue_type",
				}, nil, "name"),
				Validators: convert.NewValidators(convert.ValidatorTypeInt32, nil),
			},
		},
		"deprecation_message": {
			input: &datasource.Int32Attribute{
				DeprecationMessage: pointer("deprecation message"),
			},
			expected: GeneratorInt32Attribute{
				CustomType:         convert.NewCustomTypePrimitive(nil, nil, "name"),
				DeprecationMessage: convert.NewDeprecationMessage(pointer("deprecation message")),
				Validators:         convert.NewValidators(convert.ValidatorTypeInt32, specschema.CustomValidators{}),
			},
		},
		"description": {
			input: &datasource.Int32Attribute{
				Description: pointer("description"),
			},
			expected: GeneratorInt32Attribute{
				CustomType:  convert.NewCustomTypePrimitive(nil, nil, "name"),
				Description: convert.NewDescription(pointer("description")),
				Validators:  convert.NewValidators(convert.ValidatorTypeInt32, specschema.CustomValidators{}),
			},
		},
		"sensitive": {
			input: &datasource.Int32Attribute{
				Sensitive: pointer(true),
			},
			expected: GeneratorInt32Attribute{
				CustomType: convert.NewCustomTypePrimitive(nil, nil, "name"),
				Sensitive:  convert.NewSensitive(pointer(true)),
				Validators: convert.NewValidators(convert.ValidatorTypeInt32, specschema.CustomValidators{}),
			},
		},
		"validators": {
			input: &datasource.Int32Attribute{
				Validators: specschema.Int32Validators{
					{
						Custom: &specschema.CustomValidator{
							Imports: []code.Import{
								{
									Path: "github.com/.../myvalidator",
								},
							},
							SchemaDefinition: "myvalidator.Validate()",
						},
					},
				},
			},
			expected: GeneratorInt32Attribute{
				CustomType: convert.NewCustomTypePrimitive(nil, nil, "name"),
				Validators: convert.NewValidators(convert.ValidatorTypeInt32, specschema.CustomValidators{
					&specschema.CustomValidator{
						Imports: []code.Import{
							{
								Path: "github.com/.../myvalidator",
							},
						},
						SchemaDefinition: "myvalidator.Validate()",
					},
				}),
			},
		},
	}

	for name, testCase := range testCases {

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := NewGeneratorInt32Attribute("name", testCase.input)

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestGeneratorInt32Attribute_Schema(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         GeneratorInt32Attribute
		expected      string
		expectedError error
	}{
		"custom-type": {
			input: GeneratorInt32Attribute{
				CustomType: convert.NewCustomTypePrimitive(
					&specschema.CustomType{
						Type: "my_custom_type",
					},
					nil,
					"int32_attribute",
				),
			},
			expected: `"int32_attribute": schema.Int32Attribute{
CustomType: my_custom_type,
},`,
		},

		"associated-external-type": {
			input: GeneratorInt32Attribute{
				CustomType: convert.NewCustomTypePrimitive(
					nil,
					&specschema.AssociatedExternalType{
						Type: "*api.ExtInt32",
					},
					"int32_attribute",
				),
			},
			expected: `"int32_attribute": schema.Int32Attribute{
CustomType: Int32AttributeType{},
},`,
		},

		"custom-type-overriding-associated-external-type": {
			input: GeneratorInt32Attribute{
				CustomType: convert.NewCustomTypePrimitive(
					&specschema.CustomType{
						Type: "my_custom_type",
					},
					&specschema.AssociatedExternalType{
						Type: "*api.ExtInt32",
					},
					"int32_attribute",
				),
			},
			expected: `"int32_attribute": schema.Int32Attribute{
CustomType: my_custom_type,
},`,
		},

		"required": {
			input: GeneratorInt32Attribute{
				ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Required),
			},
			expected: `"int32_attribute": schema.Int32Attribute{
Required: true,
},`,
		},

		"optional": {
			input: GeneratorInt32Attribute{
				ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Optional),
			},
			expected: `"int32_attribute": schema.Int32Attribute{
Optional: true,
},`,
		},

		"computed": {
			input: GeneratorInt32Attribute{
				ComputedOptionalRequired: convert.NewComputedOptionalRequired(specschema.Computed),
			},
			expected: `"int32_attribute": schema.Int32Attribute{
Computed: true,
},`,
		},

		"sensitive": {
			input: GeneratorInt32Attribute{
				Sensitive: convert.NewSensitive(pointer(true)),
			},
			expected: `"int32_attribute": schema.Int32Attribute{
Sensitive: true,
},`,
		},

		// TODO: Do we need separate description and markdown description?
		"description": {
			input: GeneratorInt32Attribute{
				Description: convert.NewDescription(pointer("description")),
			},
			expected: `"int32_attribute": schema.Int32Attribute{
Description: "description",
MarkdownDescription: "description",
},`,
		},

		"deprecation-message": {
			input: GeneratorInt32Attribute{
				DeprecationMessage: convert.NewDeprecationMessage(pointer("deprecated")),
			},
			expected: `"int32_attribute": schema.Int32Attribute{
DeprecationMessage: "deprecated",
},`,
		},

		"validators-empty": {
			input: GeneratorInt32Attribute{
				Validators: convert.NewValidators(convert.ValidatorTypeInt32, nil),
			},
			expected: `"int32_attribute": schema.Int32Attribute{
},`,
		},
		"validators": {
			input: GeneratorInt32Attribute{
				Validators: convert.NewValidators(convert.ValidatorTypeInt32, specschema.CustomValidators{
					{
						SchemaDefinition: "my_validator.Validate()",
					},
					{
						SchemaDefinition: "my_other_validator.Validate()",
					},
				}),
			},
			expected: `"int32_attribute": schema.Int32Attribute{
Validators: []validator.Int32{
my_validator.Validate(),
my_other_validator.Validate(),
},
},`,
		},
	}

	for name, testCase := range testCases {

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := testCase.input.Schema("int32_attribute")

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestGeneratorInt32Attribute_ModelField(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         GeneratorInt32Attribute
		expected      model.Field
		expectedError error
	}{
		"default": {
			expected: model.Field{
				Name:      "Int32Attribute",
				ValueType: "types.Int32",
				TfsdkName: "int32_attribute",
			},
		},
		"custom-type": {
			input: GeneratorInt32Attribute{
				CustomType: convert.NewCustomTypePrimitive(
					&specschema.CustomType{
						ValueType: "my_custom_value_type",
					},
					nil,
					"",
				),
			},
			expected: model.Field{
				Name:      "Int32Attribute",
				ValueType: "my_custom_value_type",
				TfsdkName: "int32_attribute",
			},
		},
		"associated-external-type": {
			input: GeneratorInt32Attribute{
				CustomType: convert.NewCustomTypePrimitive(
					nil,
					&specschema.AssociatedExternalType{
						Type: "*api.Int32Attribute",
					},
					"int32_attribute",
				),
			},
			expected: model.Field{
				Name:      "Int32Attribute",
				ValueType: "Int32AttributeValue",
				TfsdkName: "int32_attribute",
			},
		},
		"custom-type-overriding-associated-external-type": {
			input: GeneratorInt32Attribute{
				CustomType: convert.NewCustomTypePrimitive(
					&specschema.CustomType{
						ValueType: "my_custom_value_type",
					},
					&specschema.AssociatedExternalType{
						Type: "*api.Int32Attribute",
					},
					"",
				),
			},
			expected: model.Field{
				Name:      "Int32Attribute",
				ValueType: "my_custom_value_type",
				TfsdkName: "int32_attribute",
			},
		},
	}

	for name, testCase := range testCases {

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := testCase.input.ModelField("int32_attribute")

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
