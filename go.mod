module example.com/mod

go 1.21

replace quacks => ./quacks/

replace client => ./client/

replace game => ./game/

require quacks v0.0.0-00010101000000-000000000000

require github.com/looplab/fsm v1.0.1 // indirect
