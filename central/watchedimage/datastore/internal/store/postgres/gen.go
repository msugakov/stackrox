package postgres

//go:generate pg-table-bindings-wrapper --type=storage.WatchedImage --postgres-migration-seq 58 --migrate-from "rocksdb:watchedimages"
