.PHONY: build lint fmt fmt-check test check

build lint fmt fmt-check test check:
	$(MAKE) -C backend $@
