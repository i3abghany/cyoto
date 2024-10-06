SRC_FILES=main.go utils.go visitor.go parse.go interpret.go interpret_test.go

test: cyoto
	go test

cyoto: $(SRC_FILES) pkg/parser/kyotogrammar*
	go build -o cyoto

gen: pkg/parser/kyotogrammar%

fmt: $(SRC_FILES)
	go fmt

pkg/parser/kyotogrammar%: KyotoGrammar.g4
	mkdir -p ./pkg/parser
	antlr4 -Dlanguage=Go ./KyotoGrammar.g4 -o ./pkg/parser -visitor

clean:
	rm -f ./pkg/parser/*
	rm -f cyoto

.PHONY: gen clean test fmt
