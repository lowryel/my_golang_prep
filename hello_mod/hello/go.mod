module ex/hello

go 1.20

replace ex/greetings => ../greetings

require (
	ex/greetings v0.0.0-00010101000000-000000000000
	loop.go/loop v0.0.0-00010101000000-000000000000
)

replace loop.go/loop => ../loops
