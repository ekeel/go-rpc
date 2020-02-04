build-example:
	scripts/build-example.sh

build-plugin:
	scripts/build-plugin.sh

clean: FORCE
	scripts/clean.sh

help:
	scripts/help.sh

FORCE: