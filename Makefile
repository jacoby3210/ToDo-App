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


migrate-action:	 ## Выполнить миграцию (up/down)
	@if [ -z "$(action)" ]; then \
		echo "Необходимо указать действие миграции через action!"; \
		exit 1; \
	fi
	@docker compose run --rm todoapp-postgres-migrate \
		-path ./migrations \
		-database postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@todoapp-postgres:5432/$(POSTGRES_DB)?sslmode=disable \
		"$(action)"

migrate-down:	 ## Откатить миграцию
	@make migrate-action action=down

migrate-up:	 ## Применить миграцию
	@make migrate-action action=up