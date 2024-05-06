# High Score
A lightweight and open source high score service.

[![ko-fi](https://ko-fi.com/img/githubbutton_sm.svg)](https://ko-fi.com/A1371A99)

## Intent
The intent for this project is to be a relatively small and simple high score service that can be used in the cloud with minimal setup.

For example, all you would need to do is setup a the services container as you would any other with environment variables, setup your mysql db and then have it run in the cloud.

## Requirements
### Storage
A MySQL Database, setup by using the sql script in `.\db\init\`. (Make sure to change the user and password setup)

### Configuration
The service uses environment variables from the container to setup various aspects of service configuration (DB, Hosting etc)

To configure the games that are available to the api, you will need to update the `Games` table in `DBScores` with the details required.

## Setup
### Prerequisits
- Change the `db/init/update-games.sql` file with game configs specific to your projects.

### Local (Dev)
Local setup is made to be as straight forward as possible, to the point that you would only need to run `docker compose up` to use the service with all default settings and config.

### Production
This setup assumes you have a working knowledge of the cloud service or system you want to set this up on.

A lot of the production setup will depend on what cloud service you use, or if you run it on prem, so this will cover the basics for a production setup.

1) Setup your instance of MySQL with a locked down user that only has the required permissions to access the `DBScores` database and the `Scores` + `Games` tables. Use the `db\init\init.sql` as a base for this.

2) Configure the container you want to use with the environment settings required (See the `Settings` section for more details)

3) Have the container grab the image `mrlparker/high-score:latest` from docker hub.

## Settings
Each of the following environment variables will need the prefix of `HIGHSCORE_SERVICE_` to work (eg `HIGHSCORE_SERVICE_HOST`).

`HOST` - The host to use for the server.

`PORT` - The port for the server to listen on.

`ROUTE_BASE` - The base url sections or all routes (for example with a value or `/api/` `scores` will become `/api/scores`)

`DB_USER` - The username for db access

`DB_PASS` - The password for db access

`DB_ADDRESS` - The host and port for the db (eg. localhost:1234)

`DB_NAME` - The name of the database to use.

`DB_GAMETABLE_NAME` - The name of the game table.

`DB_SCORESTABLE_NAME` - The name of the scores table.
