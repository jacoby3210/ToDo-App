include .env
export

export PROJECT_ROOT := $(shell pwd)

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
