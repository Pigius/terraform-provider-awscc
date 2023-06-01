// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package networkmanager

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_networkmanager_connect_peer", connectPeerDataSource)
}

// connectPeerDataSource returns the Terraform awscc_networkmanager_connect_peer data source.
// This Terraform data source corresponds to the CloudFormation AWS::NetworkManager::ConnectPeer resource.
func connectPeerDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: BgpOptions
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Bgp options for connect peer.",
		//	  "properties": {
		//	    "PeerAsn": {
		//	      "type": "number"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"bgp_options": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: PeerAsn
				"peer_asn": schema.Float64Attribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Bgp options for connect peer.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Configuration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Configuration of the connect peer.",
		//	  "properties": {
		//	    "BgpConfigurations": {
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "additionalProperties": false,
		//	        "description": "Bgp configuration for connect peer",
		//	        "properties": {
		//	          "CoreNetworkAddress": {
		//	            "description": "The address of a core network.",
		//	            "type": "string"
		//	          },
		//	          "CoreNetworkAsn": {
		//	            "description": "The ASN of the Coret Network.",
		//	            "type": "number"
		//	          },
		//	          "PeerAddress": {
		//	            "description": "The address of a core network Connect peer.",
		//	            "type": "string"
		//	          },
		//	          "PeerAsn": {
		//	            "description": "The ASN of the Connect peer.",
		//	            "type": "number"
		//	          }
		//	        },
		//	        "type": "object"
		//	      },
		//	      "type": "array"
		//	    },
		//	    "CoreNetworkAddress": {
		//	      "description": "The IP address of a core network.",
		//	      "type": "string"
		//	    },
		//	    "InsideCidrBlocks": {
		//	      "description": "The inside IP addresses used for a Connect peer configuration.",
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "type": "string"
		//	      },
		//	      "type": "array"
		//	    },
		//	    "PeerAddress": {
		//	      "description": "The IP address of the Connect peer.",
		//	      "type": "string"
		//	    },
		//	    "Protocol": {
		//	      "description": "Tunnel protocol type (Only support GRE for now)",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: BgpConfigurations
				"bgp_configurations": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
					NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: CoreNetworkAddress
							"core_network_address": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The address of a core network.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: CoreNetworkAsn
							"core_network_asn": schema.Float64Attribute{ /*START ATTRIBUTE*/
								Description: "The ASN of the Coret Network.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: PeerAddress
							"peer_address": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The address of a core network Connect peer.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: PeerAsn
							"peer_asn": schema.Float64Attribute{ /*START ATTRIBUTE*/
								Description: "The ASN of the Connect peer.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
					}, /*END NESTED OBJECT*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: CoreNetworkAddress
				"core_network_address": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The IP address of a core network.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: InsideCidrBlocks
				"inside_cidr_blocks": schema.ListAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Description: "The inside IP addresses used for a Connect peer configuration.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: PeerAddress
				"peer_address": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The IP address of the Connect peer.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: Protocol
				"protocol": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Tunnel protocol type (Only support GRE for now)",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Configuration of the connect peer.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ConnectAttachmentId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the attachment to connect.",
		//	  "type": "string"
		//	}
		"connect_attachment_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the attachment to connect.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ConnectPeerId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the Connect peer.",
		//	  "type": "string"
		//	}
		"connect_peer_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the Connect peer.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: CoreNetworkAddress
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The IP address of a core network.",
		//	  "type": "string"
		//	}
		"core_network_address": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The IP address of a core network.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: CoreNetworkId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the core network.",
		//	  "type": "string"
		//	}
		"core_network_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the core network.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: CreatedAt
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Connect peer creation time.",
		//	  "type": "string"
		//	}
		"created_at": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Connect peer creation time.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: EdgeLocation
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Connect peer Regions where edges are located.",
		//	  "type": "string"
		//	}
		"edge_location": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Connect peer Regions where edges are located.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: InsideCidrBlocks
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The inside IP addresses used for a Connect peer configuration.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "type": "string"
		//	  },
		//	  "type": "array"
		//	}
		"inside_cidr_blocks": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "The inside IP addresses used for a Connect peer configuration.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: PeerAddress
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The IP address of the Connect peer.",
		//	  "type": "string"
		//	}
		"peer_address": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The IP address of the Connect peer.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: State
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "State of the connect peer.",
		//	  "type": "string"
		//	}
		"state": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "State of the connect peer.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An array of key-value pairs to apply to this resource.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A key-value pair to associate with a resource.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key",
		//	      "Value"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"tags": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "An array of key-value pairs to apply to this resource.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::NetworkManager::ConnectPeer",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::NetworkManager::ConnectPeer").WithTerraformTypeName("awscc_networkmanager_connect_peer")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"bgp_configurations":    "BgpConfigurations",
		"bgp_options":           "BgpOptions",
		"configuration":         "Configuration",
		"connect_attachment_id": "ConnectAttachmentId",
		"connect_peer_id":       "ConnectPeerId",
		"core_network_address":  "CoreNetworkAddress",
		"core_network_asn":      "CoreNetworkAsn",
		"core_network_id":       "CoreNetworkId",
		"created_at":            "CreatedAt",
		"edge_location":         "EdgeLocation",
		"inside_cidr_blocks":    "InsideCidrBlocks",
		"key":                   "Key",
		"peer_address":          "PeerAddress",
		"peer_asn":              "PeerAsn",
		"protocol":              "Protocol",
		"state":                 "State",
		"tags":                  "Tags",
		"value":                 "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
