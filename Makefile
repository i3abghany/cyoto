Kyoto.exe: gen main.go
	go build

gen: $(wildcard pkg/parser/*)
	mkdir -p ./pkg/parser
	antlr4 -Dlanguage=Go ./KyotoGrammar.g4 -o ./pkg/parser -no-listener

clean:
	rm -f ./pkg/parser/*
	rm -f Kyoto.exe
