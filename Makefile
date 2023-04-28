COMPOSE := docker-compose
.PHONY: list up down clean

# Optional, add swarmlab-compose instead of docker-compose

list:
	@echo "Available commands:"
	@echo "  up        : Start the containers."
	@echo "  down      : Stop the containers."
	@echo "  list      : List all running containers."
	@echo "  clean     : Stop and remove all containers."

up:
	$(COMPOSE) up -d

down:
	$(COMPOSE) down

list:
	$(COMPOSE) ps

clean:
	$(COMPOSE) down --volumes --remove-orphans
