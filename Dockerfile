FROM python:3.8-slim

ENV PYTHONUNBUFFERED=1
ENV PATH=$PATH:/root/.local/bin
WORKDIR /project

COPY . /project/

# Install pipx
RUN python3 -m pip install --upgrade pip && \
    python3 -m pip install --user pipx --no-warn-script-location && \
    python3 -m pipx ensurepath && \
    # Install poetry
    # apk --no-cache add gcc g++ curl musl-dev postgresql-dev && \
    apt-get update -y && \
    apt-get install curl -y && \
    curl -sSL https://raw.githubusercontent.com/python-poetry/poetry/master/get-poetry.py | python - && \
    . $HOME/.poetry/env && \
    poetry config virtualenvs.create false && \
    poetry install --no-dev --no-interaction --no-ansi

CMD python manage.py migrate && python manage.py runserver 0.0.0.0:8000 --insecure
