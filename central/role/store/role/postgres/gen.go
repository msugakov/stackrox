package postgres

//go:generate pg-table-bindings-wrapper --type=storage.Role --postgres-migration-seq 50 --migrate-from "rocksdb:roles"
