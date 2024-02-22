DIST_OUT := ../dist-go
PACKAGE := needle

GREEN := "\033[32m"
RESET := "\033[0m"

antlr-go:
	@antlr4 -Dlanguage=Go -package $(PACKAGE) NeedleLexer.g4 NeedleParser.g4
	@if [ ! -d "$(DIST_OUT)" ]; then mkdir -p "$(DIST_OUT)"; fi
	@find $(DIST_OUT) -maxdepth 1 -type f -exec echo Deleting {} \; -exec rm -f {} \;
	@mv ./*.go $(DIST_OUT)
	@echo $(GREEN) "All antlr4 files generated successfully!" $(RESET)
