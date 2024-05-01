# High Score
A lightweight and open source high score service.

## Intent
The intent for this project is to be a relatively small and simple high score service that can be used in the cloud with minimal setup.

For example, all you would need to do is setup a the services container as you would any other with environment variables, etc and then have it run in the cloud.

## Requirements
### Storage
A SQL Database, setup by running the included script.

### Configuration
The service uses environment variables from the container to configure the games, their respective limits and their attached api keys. (This could change during the course of development)