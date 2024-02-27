sync:
	uv pip compile pyproject.toml -o requirements.txt
	uv pip sync requirements.txt

temporal:
	temporal server start-dev --ui-port 8080
