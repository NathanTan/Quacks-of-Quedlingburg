module main

go 1.21

replace types => ../types/
replace quacks => ../quacks/

require (
	github.com/anthdm/hollywood v0.0.0-20240115210651-dd34702ee21f
	github.com/gorilla/websocket v1.5.1
	types v0.0.0
)

require (
	github.com/DataDog/gostackparse v0.7.0 // indirect
	github.com/klauspost/cpuid/v2 v2.2.7 // indirect
	github.com/zeebo/xxh3 v1.0.2 // indirect
	golang.org/x/net v0.22.0 // indirect
	golang.org/x/sys v0.18.0 // indirect
	google.golang.org/protobuf v1.33.0 // indirect
)
