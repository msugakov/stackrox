package postgres

//go:generate pg-table-bindings-wrapper --type=storage.ProcessBaselineResults --search-category PROCESS_BASELINE_RESULTS --postgres-migration-seq 44 --migrate-from "rocksdb:processWhitelistResults"
