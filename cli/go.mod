module github.com/dc-lab/sky/cli

go 1.14

require (
	github.com/dc-lab/sky/data_manager/client v0.0.0-00010101000000-000000000000
	github.com/mitchellh/go-homedir v1.1.0
	github.com/spf13/cobra v1.0.0
	github.com/spf13/viper v1.6.3
)

replace github.com/dc-lab/sky/data_manager/client => ../data_manager/client
