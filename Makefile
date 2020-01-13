##########################################
#   _____  _____   _____   _____  ____   #
#  / ____||_   _| / ____| / ____|/ __ \  #				
# | |       | |  | (___  | |    | |  | | #				
# | |       | |   \___ \ | |    | |  | | #
# | |____  _| |_  ____) || |____| |__| | #
#  \_____||_____||_____/  \_____|\____/  #
#                                        #
##########################################


GO_VERSION = "1.13"

ENV        ?= develop

BINARIES = $(shell ls cmd/)

EOC      = \033[0m
BLUE     = \033[34m
GREEN    = \033[32m
RED      = \033[31m
WHITE    = \033[37m

DB		 = ./database/schema/swapi.dat


.PHONY: up build logs ps clean fclean help unit-tests integration-tests lint stop

help:
	@echo "$(GREEN)[$@]$(EOC): up                - Run development or production environment (default: develop)"
	@echo "$(GREEN)[$@]$(EOC): logs              - Print logs container. You can add the c=service_name option to view the specific log of a container."
	@echo "$(GREEN)[$@]$(EOC): ps                - List containers"
	@echo "$(GREEN)[$@]$(EOC): unit-tests        - Run unit tests"
	@echo "$(GREEN)[$@]$(EOC): integration-tests - Run integration tests"
	@echo "$(GREEN)[$@]$(EOC): clean             - Stop and remove containers"
	@echo "$(GREEN)[$@]$(EOC): fclean            - Stop and remove containers/volumes"
	@echo "$(GREEN)[$@]$(EOC): help              - Print this help"

base-image:
	@echo "$(BLUE)[$@]$(EOC): Build cisco/base-golang (${GO_VERSION})"
	docker build -f build/tools/Golang.dockerfile --build-arg GO_VERSION=${GO_VERSION} -t cisco/base-golang .

up: base-image
	@echo "$(BLUE)[$@]$(EOC): Run ${ENV} environment"
	docker-compose -f docker-compose.yml -f build/${ENV}/docker-compose.yml up --build -d

build: base-image
	@echo "$(BLUE)[$@]$(EOC): Build ${ENV} environment"
	docker-compose -f docker-compose.yml -f build/${ENV}/docker-compose.yml build

ps:
	docker-compose -f docker-compose.yml -f build/${ENV}/docker-compose.yml ps

logs:
	docker-compose -f docker-compose.yml -f build/${ENV}/docker-compose.yml logs -f --tail 500 ${c}

lint: base-image
	docker run --rm cisco/base-golang bash -c 'go get -u golang.org/x/lint/golint && golint -set_exit_status ./... && go vet ./...'

unit-tests: base-image
	docker run -e DB=${DB} --rm cisco/base-golang bash -c 'go test -v ./...'
stop:
	@echo "$(BLUE)[$@]$(EOC): Stop containers"
	docker-compose -f docker-compose.yml -f build/${ENV}/docker-compose.yml stop ${c}

clean:
	@echo "$(BLUE)[$@]$(EOC): Remove and stop containers"
	docker-compose -f docker-compose.yml -f build/${ENV}/docker-compose.yml down

fclean:
	@echo "$(BLUE)[$@]$(EOC): Remove volumes and containers"
	docker-compose -f docker-compose.yml -f build/${ENV}/docker-compose.yml down -v
