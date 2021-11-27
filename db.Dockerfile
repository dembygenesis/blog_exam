FROM mysql:8.0.12

COPY ./database/*.sql /docker-entrypoint-initdb.d/