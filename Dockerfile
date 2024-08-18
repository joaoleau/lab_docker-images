FROM python:3.12.3-slim
LABEL maintainer="Joao"

RUN apt-get update && apt-get install -y git && rm -rf /var/lib/apt/lists/*

WORKDIR /app

RUN git clone https://github.com/joaoleau/devops.git \
    && cp -r devops/dev_projetos/flaskapp-metrics/* . \
    && pip install -r requirements.txt \
    && rm -rf devops

EXPOSE $PORT

CMD ["python", "app.py"]
