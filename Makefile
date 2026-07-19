.DEFAULT_GOAL := help

.PHONY: help fmt lint test test-race test-integration generate up down

help:
	@printf '%s\n' 'StockFlow bootstrap commands:'
	@printf '%s\n' '  make fmt              Format Go code (not available yet)'
	@printf '%s\n' '  make lint             Run static checks (not available yet)'
	@printf '%s\n' '  make test             Run tests (not available yet)'
	@printf '%s\n' '  make test-race        Run race tests (not available yet)'
	@printf '%s\n' '  make test-integration Run integration tests (not available yet)'
	@printf '%s\n' '  make generate         Generate code (not available yet)'
	@printf '%s\n' '  make up               Start local infrastructure (not available yet)'
	@printf '%s\n' '  make down             Stop local infrastructure (not available yet)'

fmt lint test test-race test-integration generate up down:
	@printf 'StockFlow bootstrap: target "%s" is not implemented yet; no Go modules or runtime infrastructure have been added.\n' '$@'
	@exit 1
