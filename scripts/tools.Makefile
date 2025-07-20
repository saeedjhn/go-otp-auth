# This makefile should be used to hold functions/variables

ifeq ($(ARCH),x86_64)
	ARCH := amd64
else ifeq ($(ARCH),aarch64)
	ARCH := arm64 
endif

define github_url
    https://github.com/$(GITHUB)/releases/download/v$(VERSION)/$(ARCHIVE)
endef

# creates a directory scripts/bin.
scripts/bin:
	@ mkdir -p $@

# ~~~ Tools ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# ~~ [migrate] ~~~ https://github.com/golang-migrate/migrate ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

MIGRATE := $(shell command -v migrate || echo "scripts/bin/migrate")
migrate: scripts/bin/migrate ## Install migrate (database migration)

scripts/bin/migrate: VERSION := 4.18.1
scripts/bin/migrate: GITHUB  := golang-migrate/migrate
scripts/bin/migrate: ARCHIVE := migrate.$(OSTYPE)-$(ARCH).tar.gz
scripts/bin/migrate: scripts/bin
	@ if [ ! -f "$@" ]; then \
		printf "Install migrate... "; \
		curl -Ls $(shell echo $(call github_url) | tr A-Z a-z) | tar -zOxf - ./migrate > $@ && chmod +x $@; \
		echo "done."; \
	else \
		echo "migrate already installed."; \
	fi