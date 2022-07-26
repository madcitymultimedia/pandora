package schema

import "github.com/hashicorp/pandora/tools/sdk/resourcemanager"

func (b Builder) identityTopLevelFieldsWithinResourceID(input resourcemanager.ResourceIdDefinition) (*map[string]FieldDefinition, error) {
	out := make(map[string]FieldDefinition, 0)

	// TODO: don't forget that the resource name itself needs to be included
	// TODO: mappings

	if len(input.Segments) > 2 {
		parentSegments := input.Segments[0 : len(input.Segments)-2]
		if segmentsContainResource(parentSegments) {
			// find the parent Resource ID and use that
			parentResourceIdName := ""
			for name, id := range b.resourceIds {
				if segmentsMatch(id.Segments, parentSegments) {
					parentResourceIdName = name
					break
				}
			}
			if parentResourceIdName != "" {
				parentResourceSchemaField := convertToSnakeCase(parentResourceIdName)
				out[parentResourceSchemaField] = FieldDefinition{
					Definition: resourcemanager.ApiObjectDefinition{
						Type:          resourcemanager.ReferenceApiObjectDefinitionType,
						ReferenceName: &parentResourceIdName,
					},
					// since this is included in the Resource ID it's implicitly Required/ForceNew
					Required: true,
					ForceNew: true,
				}
			}

			// TODO: perhaps add the components here instead, aside from Subscription/Tenant ID etc?
		}
	}

	return &out, nil
}

func segmentsMatch(first []resourcemanager.ResourceIdSegment, second []resourcemanager.ResourceIdSegment) bool {
	if len(first) != len(second) {
		return false
	}

	for i, firstVal := range first {
		secondVal := second[i]
		if firstVal.Type != secondVal.Type {
			return false
		}

		if firstVal.Type == resourcemanager.StaticSegment || firstVal.Type == resourcemanager.ResourceProviderSegment {
			if firstVal.FixedValue == nil || secondVal.FixedValue == nil {
				return false
			}

			if *firstVal.FixedValue != *secondVal.FixedValue {
				return false
			}
		}

		if firstVal.Type == resourcemanager.ConstantSegment {
			if firstVal.ConstantReference == nil || secondVal.ConstantReference == nil {
				return false
			}

			if *firstVal.ConstantReference != *secondVal.ConstantReference {
				return false
			}
		}
	}

	return true
}

func segmentsContainResource(input []resourcemanager.ResourceIdSegment) bool {
	penultimateSegmentIsStatic := false
	lastSegmentIsUserSpecifiable := false
	if len(input) >= 4 {
		penultimateSegment := input[len(input)-2]
		lastSegment := input[len(input)-1]

		penultimateSegmentIsStatic = penultimateSegment.Type == resourcemanager.StaticSegment
		lastSegmentIsUserSpecifiable = lastSegment.Type == resourcemanager.UserSpecifiedSegment
	}
	return penultimateSegmentIsStatic && lastSegmentIsUserSpecifiable
}