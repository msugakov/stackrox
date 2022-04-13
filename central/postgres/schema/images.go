// Code generated by pg-bindings generator. DO NOT EDIT.

package schema

import (
	"github.com/stackrox/stackrox/pkg/postgres"
)

var (
	// CreateTableImagesStmt holds the create statement for table `Images`.
	CreateTableImagesStmt = &postgres.CreateStmts{
		Table: `
               create table if not exists images (
                   Id varchar,
                   Name_Registry varchar,
                   Name_Remote varchar,
                   Name_Tag varchar,
                   Name_FullName varchar,
                   Metadata_V1_Created timestamp,
                   Metadata_V1_User varchar,
                   Metadata_V1_Command text[],
                   Metadata_V1_Entrypoint text[],
                   Metadata_V1_Volumes text[],
                   Metadata_V1_Labels jsonb,
                   Scan_ScanTime timestamp,
                   Scan_OperatingSystem varchar,
                   Signature_Fetched timestamp,
                   Components integer,
                   Cves integer,
                   FixableCves integer,
                   LastUpdated timestamp,
                   RiskScore numeric,
                   TopCvss numeric,
                   serialized bytea,
                   PRIMARY KEY(Id)
               )
               `,
		Indexes: []string{},
		Children: []*postgres.CreateStmts{
			&postgres.CreateStmts{
				Table: `
               create table if not exists images_Layers (
                   images_Id varchar,
                   idx integer,
                   Instruction varchar,
                   Value varchar,
                   PRIMARY KEY(images_Id, idx),
                   CONSTRAINT fk_parent_table_0 FOREIGN KEY (images_Id) REFERENCES images(Id) ON DELETE CASCADE
               )
               `,
				Indexes: []string{
					"create index if not exists imagesLayers_idx on images_Layers using btree(idx)",
				},
				Children: []*postgres.CreateStmts{},
			},
		},
	}
)
