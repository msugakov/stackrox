package postgres

//go:generate pg-table-bindings-wrapper --type=storage.SimpleAccessScope --postgres-migration-seq 56 --migrate-from "rocksdb:simple_access_scopes"
