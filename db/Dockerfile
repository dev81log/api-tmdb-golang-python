# use the official MySQL 5.7 image as the base image
FROM mysql:5.7

# set the environment variables for MySQL
ENV MYSQL_ROOT_PASSWORD=root
ENV MYSQL_DATABASE=movies_db
ENV MYSQL_USER=asher
ENV MYSQL_PASSWORD=asher123

# copy the SQL file to docker-entrypoint-initdb.d
COPY movies_db_backup.sql /docker-entrypoint-initdb.d/

# change the working directory
WORKDIR /docker-entrypoint-initdb.d/

# set the file permissions for the SQL file
RUN chmod 755 /docker-entrypoint-initdb.d/movies_db_backup.sql

# expose the default MySQL port
EXPOSE 3306

# start the MySQL server
CMD ["mysqld"]
