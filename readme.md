# Postgres Container

Pull down an image for the latest stable release of Postgres

`docker pull postgres`

To pull down a version other than the latest stable release

`docker pull postgres:[tag_you_want]`

```
$ docker images
   >>>
REPOSITORY    TAG       IMAGE ID        CREATED        SIZE
postgres      latest    9907cacf0c01    2 weeks ago    314MB
```

Create a directory to serve as local host mount for persistent data store

```
## 1. Create a folder in a known location for you
mkdir -p $HOME/docker/volumes/postgres/templateApi

## 2. run the postgres image
$ docker run --rm --name pg-docker -e POSTGRES_PASSWORD=docker -d -p 5432:5432 -v $HOME/docker/volumes/postgres/templateApi:/var/lib/postgresql/data  postgres

## 3. check that the container is running
$ docker ps
>>>
CONTAINER ID        IMAGE               COMMAND                  CREATED             STATUS              PORTS                    NAMES
dfa570d6e843        postgres            "docker-entrypoint.s…"   27 hours ago        Up 3 seconds        0.0.0.0:5432->5432/tcp   pg-docker
```

enter the cli of the docker container

```
$ docker exec -it pg-docker bash
>>> Now you are in the container's bash console. Connect to the database
root@dfa570d6e843:/# psql -h localhost -U postgres -d postgres
>>>
psql (12.2 (Debian 12.2-2.pgdg100+1))
Type "help" for help.
postgres-# \l
List of databases
Name       |  Owner   | Encoding |  Collate   |   Ctype    |   ...
-----------+----------+----------+------------+------------+------postgres   | postgres |   UTF8   | en_US.utf8 | en_US.utf8 |   ...
```


Using Jetbrains Database Connection Tool

```
connection type: default    Driver: PostgreSQL
Host: localhost Port: 5432
Authentication: User & Password
User: <default is postgres, or enter username>
Password: <insert password>
Save: forever
database: <blank>
URL: jdbc:postgresql://localhost:5432/
```

