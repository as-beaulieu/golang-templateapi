# Mocks

Mockgen is being used to automate mocks for testing

`./mock-gen.sh`

mock-gen shell is standardizing locations and interfaces targeted for mocks

`Be advised` - mockgen must be pointed to lowest level interfaces for mocking. 
When pointed to an interface compositioned from interfaces will give an error

```
Loading input failed: src/service/service.go:10:2: unknown embedded interface HealthReporter
```

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
$ docker run --name pg-docker -e POSTGRES_PASSWORD=docker -d -p 5432:5432 -v $HOME/docker/volumes/postgres/templateApi:/var/lib/postgresql/data  postgres
```

— rm: Automatically remove the container and it’s associated file system upon exit. In general, if we are running lots of short term containers, it is good practice to to pass rm flag to the docker run command for automatic cleanup and avoid disk space issues. We can always use the v option (described below) to persist data beyond the lifecycle of a container
— name: An identifying name for the container. We can choose any name we want. Note that two existing (even if they are stopped) containers cannot have the same name. In order to re-use a name, you would either need pass the rm flag to the docker run command or explicitly remove the container by using the command docker rm [container name].
-e: Expose environment variable of name POSTGRES_PASSWORD with value docker to the container. This environment variable sets the superuser password for PostgreSQL. We can set POSTGRES_PASSWORD to anything we like. I just choose it to be docker for demonstration. There are additional environment variables you can set. These include POSTGRES_USER and POSTGRES_DB. POSTGRES_USER sets the superuser name. If not provided, the superuser name defaults to postgres. POSTGRES_DB sets the name of the default database to setup. If not provided, it defaults to the value of POSTGRES_USER.
-d: Launches the container in detached mode or in other words, in the background.
-p: Bind port 5432 on localhost to port 5432 within the container. This option enables applications running out side of the container to be able to connect to the Postgres server running inside the container.
-v: Mount $HOME/docker/volumes/postgres on the host machine to the container side volume path /var/lib/postgresql/data created inside the container. This ensures that postgres data persists even after the container is removed.

```
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

Running automation docker from dockerfile

docker build -f Dockerfile.populate-users -t populate-users --build-arg POSTGRES_PORT=${POSTGRES_PORT} .

docker run -rm -i populate-users -e POSTGRES_PORT=5432

