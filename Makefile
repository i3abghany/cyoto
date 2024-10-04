Kyoto.exe: gen main.go
	go build

gen: KyotoGrammar.g4
	mkdir -p ./pkg/parser
	antlr4 -Dlanguage=Go ./KyotoGrammar.g4 -o ./pkg/parser

clean:
	rm -f ./pkg/parser/*
	rm -f Kyoto.exe

.PHONY: gen clean
