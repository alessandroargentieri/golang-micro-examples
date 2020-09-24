$ systemctl status docker
$ systemctl restart docker
$ docker pull mysql
$ docker run --name mymysql -e MYSQL_ROOT_PASSWORD=password -d mysql:latest
$ docker ps
$ docker exec -it fe7fb92a7205 sh
# echo $MYSQL_ROOT_PASSWORD
# mysql --user=root --password=$MYSQL_ROOT_PASSWORD
mysql> create database db_example; -- Creates the new database
mysql> use db_example;
mysql> create table user ( id smallint unsigned not null auto_increment, name varchar(20) not null, email varchar(50), constraint pk_user primary key (id) );
mysql> create user 'springuser'@'%' identified by 'ThePassword'; -- Creates the user
mysql> grant all on db_example.* to 'springuser'@'%'; -- Gives all privileges to the new user on the newly created database
mysql> exit
# exit
$ alias docker-ip='sudo docker inspect --format="{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}'
$ docker-ip fe7fb92a7205
$ docker ps

verify the IP and the port and add the ENV var MYSQL_HOST in the run configuration with the value got from "docker-ip <SHA-mysql>"
then run the server and add some values to the DB:
$ curl -X POST 'http://localhost:8080/add' -d '{"name":"First", "email":"someemail@someemailprovider.com"}' -H 'Content-Type:application/json'
$ curl -X POST 'http://localhost:8080/add' -d '{"name":"Second", "email":"anotheremail@someemailprovider.com"}' -H 'Content-Type:application/json'

then ask for results:
$ curl 'http://localhost:8080/get' -H 'Content-Type:application/json'
