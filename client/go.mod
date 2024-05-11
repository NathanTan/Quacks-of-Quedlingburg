module main

go 1.21.4

replace types => ../types/

require (
	github.com/gorilla/websocket v1.5.1
	types v0.0.0
)

require golang.org/x/net v0.22.0 // indirect
