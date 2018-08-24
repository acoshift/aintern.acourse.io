default:
	# `make deploy` build and deploy to production
	# `make dev` starts server in localhost:8080

dev:
	live-server --mount=/:assets/ src/

deploy: clean public
	firebase deploy

clean:
	rm -rf public/

.PHONY: public
public:
	mkdir -p public
	cp -r assets/* public/
	html-minifier \
		--collapse-boolean-attributes \
		--collapse-inline-tag-whitespace \
		--collapse-whitespace \
		--minify-css \
		--minify-js \
		--minify-urls \
		--remove-attribute-quotes \
		--remove-comments \
		--remove-empty-attributes \
		--remove-optional-tags \
		--remove-redundant-attributes \
		--remove-script-type-attributes \
		--remove-style-link-type-attributes \
		--remove-tag-whitespace \
		--input-dir src \
		--output-dir public
	# find public -name '*.html' | while read f; do mv "$$f" "$${f%.html}"; done

setup:
	npm install -g live-server html-minifier
