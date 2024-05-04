# High Score
A lightweight and open source high score service.

## Intent
The intent for this project is to be a relatively small and simple high score service that can be used in the cloud with minimal setup.

For example, all you would need to do is setup a the services container as you would any other with environment variables, etc and then have it run in the cloud.

## Requirements
### Storage
A MySQL Database, setup by using the sql script in `.\db\init\`. (Make sure to change the user and password setup)

### Configuration
The service uses environment variables from the container to setup various aspects of service configuration (DB, Hosting etc) and has a different config.json file to handle game specific details.

**config.json structure**
```
{
    "games":{ // The over all dictionary structure
        "gameId":{ // Individual game entry.
            "apiKey": string, // The key to access the api.
            "maxNameLength": int, // The max length the name value can be.
            "maxScore": int // The max score that can be submitted.
        },
        "gameId2": {
            "apiKey": string,
            "maxNameLength": int,
            "maxScore": int
        }
    }
}
```

## Setup
### Prerequisits
- Change the `config/config.json` file with game configs specific to your projects.

### Local (Dev)
Local setup is made to be as straight forward as possible, to the point that you would only need to run `docker compose up` to use the service with all default settings and config.

### Production
This setup assumes you have a working knowledge of the cloud service or system you want to set this up on.

A lot of the production setup will depend on what cloud service you use, or if you run it on prem, so this will cover the basics for a production setup.

1) Setup your instance of MySQL with a locked down user that only has the required permissions to access the `DBScores` database and the `Scores` table. Use the `db\init\init.sql` as a base for this.

2) Configure the container you want to use with the environment settings required (See the `Settings` section for more details)

3) Setup a mount for reading the config file. The service will either look in `./config/config.json` or `/etc/highscore/config.json` so binding to one of those will work. (See `Configuration` section for the structure)

4) Have the container grab the image `mrlparker/high-score:latest` from docker hub.

## Settings
Each of the following environment variables will need the prefix of `HIGHSCORE_SERVICE_` to work (eg `HIGHSCORE_SERVICE_HOST`).

`HOST` - The host to use for the server.

`PORT` - The port for the server to listen on.

`DB_USER` - The username for db access

`DB_PASS` - The password for db access

`DB_ADDRESS` - The host and port for the db (eg. localhost:1234)

`DB_NAME` - The name of the database to use.
