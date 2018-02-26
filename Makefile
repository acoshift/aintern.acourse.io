WEB=aintern.acourse.io

default:
	# `make deploy` build and deploy to production
	# `make dev` starts server in localhost:8080

dev:
	live-server --mount=/:assets/ src/

deploy: clean build
	gsutil -m rm -rf gs://$(WEB)/**
	gsutil -m -h "Cache-Control: public, max-age=3600" cp -r assets/* gs://$(WEB)
	gsutil \
		-m \
		-h "Cache-Control: public, max-age=3600" \
		-h "Content-Type: text/html" \
		cp -r build/* gs://$(WEB)

clean:
	rm -rf build/

.PHONY: build
build:
	mkdir -p build
	html-minifier \
		--html5 \
		--collapse-boolean-attributes \
		--collapse-inline-tag-whitespace \
		--collapse-whitespace \
		--minify-css \
		--minify-js \
		--minify-ur-ls \
		--remove-attribute-quotes \
		--remove-comments \
		--remove-empty-attributes \
		--remove-optional-tags \
		--remove-redundant-attributes \
		--remove-script-type-attributes \
		--remove-style-link-type-attributes \
		--remove-tag-whitespace \
		--input-dir src \
		--output-dir build
	find build -name '*.html' | while read f; do mv "$$f" "$${f%.html}"; done

setup:
	npm install -g live-server html-minifier
