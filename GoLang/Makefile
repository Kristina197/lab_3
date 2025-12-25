.PHONY: test coverage clean

test:
	@echo "Запуск тестов..."
	@go test -v ./...

coverage:
	@echo "Генерация покрытия..."
	@mkdir -p coverage
	@go test -coverprofile=coverage/coverage.out ./...
	@go tool cover -func=coverage/coverage.out
	@go tool cover -html=coverage/coverage.out -o coverage/report.html
	@echo "HTML отчет: coverage/report.html"

clean:
	@rm -rf coverage coverage-html
	@rm -f coverage.out coverage.html