consumer:
	@cd consumer && go run .

producer:
	@cd producer && go run .

.PHONY: consumer
.PHONY: producer
