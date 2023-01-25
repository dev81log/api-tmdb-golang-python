# create dockerfile from mysql image
FROM mysql:5.7
# set environment variables
ENV MYSQL_ROOT_PASSWORD=root
ENV MYSQL_DATABASE=movies_db
ENV MYSQL_USER=asher
ENV MYSQL_PASSWORD=asher123
# copy sql file to docker-entrypoint-initdb.d
COPY ./movies_db_backup.sql /docker-entrypoint-initdb.d

EXPOSE 3306