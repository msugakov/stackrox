// Code generated by pg-bindings generator. DO NOT EDIT.

package schema

import (
	"github.com/stackrox/stackrox/pkg/postgres"
)

var (

	// CreateTableImageComponentRelationStmt holds the create statement for table `ImageComponentRelation`.
	CreateTableImageComponentRelationStmt = &postgres.CreateStmts{
		Table: `
               create table if not exists image_component_relation (
                   image_components_Name varchar,
                   image_components_Version varchar,
                   image_components_OperatingSystem varchar,
                   Location varchar,
                   ImageId varchar,
                   ImageComponentId varchar,
                   serialized bytea,
                   PRIMARY KEY(image_components_Name, image_components_Version, image_components_OperatingSystem, ImageId, ImageComponentId),
                   CONSTRAINT fk_parent_table_0 FOREIGN KEY (ImageId) REFERENCES images(Id) ON DELETE CASCADE,
                   CONSTRAINT fk_parent_table_1 FOREIGN KEY (ImageComponentId, image_components_Name, image_components_Version, image_components_OperatingSystem) REFERENCES image_components(Id, Name, Version, OperatingSystem) ON DELETE CASCADE
               )
               `,
		Indexes: []string{},
	}
)
