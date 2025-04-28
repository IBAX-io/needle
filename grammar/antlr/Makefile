DIST_OUT := ../dist-go
DIST_OUT_JS := ../dist-js
DIST_OUT_PY := ../dist-py
DIST_OUT_JAVA :=../dist-java

PACKAGE := needle

GREEN := "\033[32m"
RESET := "\033[0m"

antlr-go:
	@antlr4 -Dlanguage=Go -package $(PACKAGE) NeedleLexer.g4 NeedleParser.g4
	@if [ ! -d "$(DIST_OUT)" ]; then mkdir -p "$(DIST_OUT)"; fi
	@find $(DIST_OUT) -maxdepth 1 -type f -exec echo Deleting {} \; -exec rm -f {} \;
	@mv ./*.go $(DIST_OUT)
	@echo $(GREEN) "All antlr4 files generated successfully!" $(RESET)

antlr-js:
	@antlr4 -Dlanguage=JavaScript NeedleLexer.g4 NeedleParser.g4
	@if [ ! -d "$(DIST_OUT_JS)" ]; then mkdir -p "$(DIST_OUT_JS)"; fi
	@find $(DIST_OUT_JS) -maxdepth 1 -type f -exec echo Deleting {} \; -exec rm -f {} \;
	@mv ./*.js $(DIST_OUT_JS)
	@echo $(GREEN) "All antlr4 files generated successfully for JavaScript!" $(RESET)

antlr-py:
	@antlr4 -Dlanguage=Python3 NeedleLexer.g4 NeedleParser.g4
	@if [ ! -d "$(DIST_OUT_PY)" ]; then mkdir -p "$(DIST_OUT_PY)"; fi
	@find $(DIST_OUT_PY) -maxdepth 1 -type f -exec echo Deleting {} \; -exec rm -f {} \;
	@mv ./*.py $(DIST_OUT_PY)
	@echo $(GREEN) "All antlr4 files generated successfully for Python 3!" $(RESET)

antlr-java:
	@antlr4 -Dlanguage=Java NeedleLexer.g4 NeedleParser.g4
	@if [ ! -d "$(DIST_OUT_JAVA)" ]; then mkdir -p "$(DIST_OUT_JAVA)"; fi
	@find $(DIST_OUT_JAVA) -maxdepth 1 -type f -exec echo Deleting {} \; -exec rm -f {} \;
	@mv ./*.java $(DIST_OUT_JAVA)
	@echo $(GREEN) "All antlr4 files generated successfully for Java!" $(RESET)