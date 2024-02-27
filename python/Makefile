sync:
	uv pip compile pyproject.toml -o requirements.txt
	uv pip sync requirements.txt

worker:
	python3 src/worker.py

client:
	python3 src/client.py
