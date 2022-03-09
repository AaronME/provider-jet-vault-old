package generic

import (
	"github.com/crossplane/terrajet/pkg/config"
	"github.com/crossplane/terrajet/pkg/types/comments"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vault_generic_secret", func(r *config.Resource) {

		// we need to override the default group that terrajet generated for
		// this resource, which would be "vault"
		r.ShortGroup = "generic"

		r.ExternalName = config.IdentifierFromProvider

		// Add field for
		c := "String that will be written as the secret data at the given path.\n"
		comment, err := comments.New(c, comments.WithTFTag("-"))
		if err != nil {
			panic(errors.Wrap(err, "cannot build comment for string_data"))
		}
		r.TerraformResource.Schema["StringData"] = &schema.Schema{
			Type:         schema.TypeString,
			Sensitive:    false,
			Optional:     false,
			ExactlyOneOf: []string{"DataJSON", "StringData"},
			Description:  comment.String(),
		}

		// Fix requirement on data_json field

		// // This results in multiple Ref and Selector fields in addition to
		// // the required one I'm trying to remove.
		// r.References = config.References{
		// 	"data_json": {},
		// }

		// // This results in an error on generation that multiple resources have the same name
		// r.TerraformResource.Schema["DataJSON"] = &schema.Schema{
		// 	Type:         schema.TypeString,
		// 	Sensitive:    true,
		// 	Optional:     true,
		// 	ExactlyOneOf: []string{"DataJSON", "StringData"},
		// 	Description:  comment.String(),
		// }

		// // This results in a nil pointer panic
		// r.TerraformResource.Schema["DataJSON"].ExactlyOneOf = []string{"DataJSONSecretRef", "StringDataSecretRef"}
	})
}
