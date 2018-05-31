FROM golang

COPY /test-sg /opt/testapp/test-sg
COPY /wait-for-postgres.sh /opt/testapp/wait-for-postgres.sh

RUN apt-get update && apt-get install postgresql-client -y && chmod 777 /opt/testapp/wait-for-postgres.sh

EXPOSE 8080

CMD  /opt/testapp/wait-for-postgres.sh postgres "/opt/testapp/test-sg migrate" "/opt/testapp/test-sg server"
