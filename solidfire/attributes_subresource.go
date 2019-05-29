package solidfire

import "github.com/hashicorp/terraform/helper/schema"

func schemaResourceAttributes() *schema.Schema {
	return &schema.Schema{
		// Attributes subresource
		Type:     schema.TypeMap,
		Optional: true,
		Elem:     &schema.Schema{Type: schema.TypeString},
	}
}

func schemaDataSourceAttributes() *schema.Schema {
	return &schema.Schema{
		// Attributes subresource
		Type:     schema.TypeMap,
		Computed: true,
		Elem:     &schema.Schema{Type: schema.TypeString},
	}
}
