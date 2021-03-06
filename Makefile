tag ?= latest
version ?= $(shell yq e '.version' helm/Chart.yaml)
clean-cmd = docker compose down --remove-orphans --volumes

smoke-test:
	docker compose up -d jwks
	sleep 3
	IMAGE_TAG=$(tag) docker compose up -d prod

prod-image:
	IMAGE_TAG=$(tag) docker compose build prod

push-prod:
	IMAGE_TAG=$(tag) docker compose push prod

dev:
	docker compose up --build dev jwks

cluster-dev:
	skaffold dev

test: clean
	docker compose up -d jwks
	docker compose run --no-deps test
	$(clean-cmd)

dev-test: clean
	docker compose run --no-deps dev-test
	$(clean-cmd)

clean:
	$(clean-cmd)
	go clean

helm-chart:
	@helm package helm/chart

publish-helm:
	@curl --user "$(CHART_AUTH_USER):$(CHART_AUTH_PASS)" \
        -F "chart=@im-job-$(version).tgz" \
        https://helm-charts.fitfit.dk/api/charts

swagger-check-install:
	which swagger || go get -u github.com/go-swagger/go-swagger/cmd/swagger

swagger-clean:
	rm -rf swagger/sdk/*
	rm -f swagger/swagger.yaml

swagger-docs: swagger-check-install
	swagger generate spec -o swagger/swagger.yaml -x swagger/sdk --scan-models
	swagger validate swagger/swagger.yaml

swagger-client: swagger-check-install
	swagger generate client -f swagger/swagger.yaml -t swagger/sdk

swagger: swagger-clean swagger-docs swagger-client

di:
	wire gen ./internal/di

.PHONY: binary docker-image push-docker-image dev test dev-test helm-chart publish-helm
