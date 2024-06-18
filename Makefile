.ONESHELL:
.PHONY: statik

statik:
	cd web && yarn build
	statik -src=./web/build

clean:
	rm -rf ./statik