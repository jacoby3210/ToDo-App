include .env
export

export PROJECT_ROOT := $(shell cygpath -m "$(CURDIR)")

## Environments

env-cleanup:   ## Удалить данные окружения (volume)
	@read -p "Очистить volume? Опасность утери данных. [y/N]: " ans; \
	if [ "$$ans" = "y" ]; then \
		docker compose down todoapp-postgres port-forwarder && \
		rm -rf ${PROJECT_ROOT}/out/pgdata; \
	fi

env-down:      ## Остановить окружение
	@docker compose down todoapp-postgres

env-up:        ## Запустить окружение
	@docker compose up -d todoapp-postgres

## Migrations

migrate-create:   ## Создать новую версию схемы
	@if [ -z "$(seq)" ]; then \
		echo "Необходимо указать номер версии миграции через seq!"; \
		exit 1; \
	fi
	@echo "PROJECT_ROOT=$(PROJECT_ROOT)"
	@docker compose run --rm todoapp-postgres-migrate \
		create \
		-ext sql \
		-dir ./migrations \
		-seq "$(seq)"
