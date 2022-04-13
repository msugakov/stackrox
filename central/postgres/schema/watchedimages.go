// Code generated by pg-bindings generator. DO NOT EDIT.

package schema

import (
	"github.com/stackrox/stackrox/pkg/postgres"
)

var (
	// CreateTableWatchedimagesStmt holds the create statement for table `Watchedimages`.
	CreateTableWatchedimagesStmt = &postgres.CreateStmts{
		Table: `
               create table if not exists watchedimages (
                   Name varchar,
                   serialized bytea,
                   PRIMARY KEY(Name)
               )
               `,
		Indexes:  []string{},
		Children: []*postgres.CreateStmts{},
	}
)
