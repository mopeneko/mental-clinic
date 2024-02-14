.PHONY: gen
gen:
	goa gen github.com/mopeneko/mental-clinic/api/design -o api

.PHONY: example
example:
	goa example github.com/mopeneko/mental-clinic/api/design -o api
