package postgres

//go:generate pg-table-bindings-wrapper --type=storage.ActiveComponent --search-category ACTIVE_COMPONENT --references storage.Deployment,storage.ImageComponent --postgres-migration-seq 58 --migrate-from "rocksdb:active_components"
