package solidfire

import "github.com/hashicorp/terraform/helper/schema"

func schemaAttributes() *schema.Schema {
	return &schema.Schema{
		// Attributes subresource
		Type:     schema.TypeMap,
		Optional: true,
	}
}
