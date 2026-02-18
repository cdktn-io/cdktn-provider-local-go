// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package provider


type LocalProviderConfig struct {
	// Alias name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/local/2.7.0/docs#alias LocalProvider#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
}

