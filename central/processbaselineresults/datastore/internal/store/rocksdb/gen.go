package rocksdb

//go:generate rocksdb-bindings-wrapper --type=ProcessBaselineResults --bucket=processWhitelistResults --key-func=GetDeploymentId() --migrate-seq 43 --migrate-to process_baseline_results
