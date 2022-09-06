ifeq (wkhtmltopdf,$(firstword $(MAKECMDGOALS)))
  # use the rest as arguments for "run"
  RUN_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
  # ...and turn them into do-nothing targets
  $(eval $(RUN_ARGS):;@:)
endif

build:
	docker build -t wkhtmltopdf .

img-cls:
	docker image rm wkhtmltopdf

wkhtmltopdf: build
	docker run --rm -v $(PWD):/app/files wkhtmltopdf $(RUN_ARGS)
	make img-cls