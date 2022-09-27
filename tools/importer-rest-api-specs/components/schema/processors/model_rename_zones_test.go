package processors

import (
	"testing"

	"github.com/hashicorp/pandora/tools/sdk/resourcemanager"
)

func TestProcessModel_RenameZones_Valid(t *testing.T) {
	testData := []struct {
		modelNameInput string
		modelsInput    map[string]resourcemanager.TerraformSchemaModelDefinition
		expected       map[string]resourcemanager.TerraformSchemaModelDefinition
	}{
		{
			modelNameInput: "Disco",
			modelsInput: map[string]resourcemanager.TerraformSchemaModelDefinition{
				"Disco": {
					Fields: map[string]resourcemanager.TerraformSchemaFieldDefinition{
						"AvailabilityZones": {
							ObjectDefinition: resourcemanager.TerraformSchemaFieldObjectDefinition{
								Type: resourcemanager.TerraformSchemaFieldTypeList,
								NestedObject: &resourcemanager.TerraformSchemaFieldObjectDefinition{
									Type: resourcemanager.TerraformSchemaFieldTypeString,
								},
							},
						},
					},
				},
			},
			expected: map[string]resourcemanager.TerraformSchemaModelDefinition{
				"Disco": {
					Fields: map[string]resourcemanager.TerraformSchemaFieldDefinition{
						"Zones": {
							ObjectDefinition: resourcemanager.TerraformSchemaFieldObjectDefinition{
								Type: resourcemanager.TerraformSchemaFieldTypeList,
								NestedObject: &resourcemanager.TerraformSchemaFieldObjectDefinition{
									Type: resourcemanager.TerraformSchemaFieldTypeString,
								},
							},
						},
					},
				},
			},
		},
	}

	for _, v := range testData {
		t.Logf("[DEBUG] Testing %s", v.modelNameInput)

		actual, err := modelRenameZones{}.ProcessModel(v.modelNameInput, v.modelsInput[v.modelNameInput], v.modelsInput)
		if err != nil {
			if v.expected == nil {
				continue
			}

			t.Fatalf("error: %+v", err)
		}

		if !modelDefinitionsMatch(t, actual, v.expected) {
			t.Fatalf("Expected %+v but got %+v", v.expected, actual)
		}
	}
}

func TestProcessModel_RenameZone_Valid(t *testing.T) {
	testData := []struct {
		modelNameInput string
		modelsInput    map[string]resourcemanager.TerraformSchemaModelDefinition
		expected       map[string]resourcemanager.TerraformSchemaModelDefinition
	}{
		{
			modelNameInput: "Disco",
			modelsInput: map[string]resourcemanager.TerraformSchemaModelDefinition{
				"Disco": {
					Fields: map[string]resourcemanager.TerraformSchemaFieldDefinition{
						"AvailabilityZone": {
							ObjectDefinition: resourcemanager.TerraformSchemaFieldObjectDefinition{
								Type: resourcemanager.TerraformSchemaFieldTypeString,
							},
						},
					},
				},
			},
			expected: map[string]resourcemanager.TerraformSchemaModelDefinition{
				"Disco": {
					Fields: map[string]resourcemanager.TerraformSchemaFieldDefinition{
						"Zone": {
							ObjectDefinition: resourcemanager.TerraformSchemaFieldObjectDefinition{
								Type: resourcemanager.TerraformSchemaFieldTypeString,
							},
						},
					},
				},
			},
		},
	}

	for _, v := range testData {
		t.Logf("[DEBUG] Testing %s", v.modelNameInput)

		actual, err := modelRenameZones{}.ProcessModel(v.modelNameInput, v.modelsInput[v.modelNameInput], v.modelsInput)
		if err != nil {
			if v.expected == nil {
				continue
			}

			t.Fatalf("error: %+v", err)
		}

		if !modelDefinitionsMatch(t, actual, v.expected) {
			t.Fatalf("Expected %+v but got %+v", v.expected, actual)
		}
	}
}

func TestProcessModel_RenameZones_Invalid(t *testing.T) {
	testData := []struct {
		modelNameInput string
		modelsInput    map[string]resourcemanager.TerraformSchemaModelDefinition
		expected       map[string]resourcemanager.TerraformSchemaModelDefinition
	}{
		{
			modelNameInput: "Leopard",
			modelsInput: map[string]resourcemanager.TerraformSchemaModelDefinition{
				"Leopard": {
					Fields: map[string]resourcemanager.TerraformSchemaFieldDefinition{
						"Id": {
							ObjectDefinition: resourcemanager.TerraformSchemaFieldObjectDefinition{
								Type:          resourcemanager.TerraformSchemaFieldTypeReference,
								ReferenceName: stringPointer("SubResource"),
							},
						},
						"Weight": {
							ObjectDefinition: resourcemanager.TerraformSchemaFieldObjectDefinition{
								Type: resourcemanager.TerraformSchemaFieldTypeInteger,
							},
						},
					},
				},
				"SubResource": {
					Fields: map[string]resourcemanager.TerraformSchemaFieldDefinition{
						"Id": {
							ObjectDefinition: resourcemanager.TerraformSchemaFieldObjectDefinition{
								Type: resourcemanager.TerraformSchemaFieldTypeString,
							},
						},
					},
				},
			},
			// unchanged
			expected: map[string]resourcemanager.TerraformSchemaModelDefinition{
				"Leopard": {
					Fields: map[string]resourcemanager.TerraformSchemaFieldDefinition{
						"Id": {
							ObjectDefinition: resourcemanager.TerraformSchemaFieldObjectDefinition{
								Type:          resourcemanager.TerraformSchemaFieldTypeReference,
								ReferenceName: stringPointer("SubResource"),
							},
						},
						"Weight": {
							ObjectDefinition: resourcemanager.TerraformSchemaFieldObjectDefinition{
								Type: resourcemanager.TerraformSchemaFieldTypeInteger,
							},
						},
					},
				},
				"SubResource": {
					Fields: map[string]resourcemanager.TerraformSchemaFieldDefinition{
						"Id": {
							ObjectDefinition: resourcemanager.TerraformSchemaFieldObjectDefinition{
								Type: resourcemanager.TerraformSchemaFieldTypeString,
							},
						},
					},
				},
			},
		},
		{
			modelNameInput: "Meerkat",
			modelsInput: map[string]resourcemanager.TerraformSchemaModelDefinition{
				"Meerkat": {
					Fields: map[string]resourcemanager.TerraformSchemaFieldDefinition{
						"Id": {
							ObjectDefinition: resourcemanager.TerraformSchemaFieldObjectDefinition{
								Type:          resourcemanager.TerraformSchemaFieldTypeReference,
								ReferenceName: stringPointer("SubResource"),
							},
						},
					},
				},
				"SubResource": {
					Fields: map[string]resourcemanager.TerraformSchemaFieldDefinition{
						"Id": {
							ObjectDefinition: resourcemanager.TerraformSchemaFieldObjectDefinition{
								Type: resourcemanager.TerraformSchemaFieldTypeString,
							},
						},
						"Weight": {
							ObjectDefinition: resourcemanager.TerraformSchemaFieldObjectDefinition{
								Type: resourcemanager.TerraformSchemaFieldTypeInteger,
							},
						},
					},
				},
			},
			// unchanged since the nested model doesn't match
			expected: map[string]resourcemanager.TerraformSchemaModelDefinition{
				"Meerkat": {
					Fields: map[string]resourcemanager.TerraformSchemaFieldDefinition{
						"Id": {
							ObjectDefinition: resourcemanager.TerraformSchemaFieldObjectDefinition{
								Type:          resourcemanager.TerraformSchemaFieldTypeReference,
								ReferenceName: stringPointer("SubResource"),
							},
						},
					},
				},
				"SubResource": {
					Fields: map[string]resourcemanager.TerraformSchemaFieldDefinition{
						"Id": {
							ObjectDefinition: resourcemanager.TerraformSchemaFieldObjectDefinition{
								Type: resourcemanager.TerraformSchemaFieldTypeString,
							},
						},
						"Weight": {
							ObjectDefinition: resourcemanager.TerraformSchemaFieldObjectDefinition{
								Type: resourcemanager.TerraformSchemaFieldTypeInteger,
							},
						},
					},
				},
			},
		},
	}

	for _, v := range testData {
		t.Logf("[DEBUG] Testing %s", v.modelNameInput)

		actual, err := modelRenameZones{}.ProcessModel(v.modelNameInput, v.modelsInput[v.modelNameInput], v.modelsInput)
		if err != nil {
			if v.expected == nil {
				continue
			}

			t.Fatalf("error: %+v", err)
		}

		if !modelDefinitionsMatch(t, actual, v.expected) {
			t.Fatalf("Expected %+v but got %+v", v.expected, actual)
		}
	}
}