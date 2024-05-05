-- This is for the local dev setup of this project.
-- Make sure to follow the setup docs when being used in production.
CREATE DATABASE DBScores;

CREATE USER 'devuser'@'localhost' IDENTIFIED BY 'devpassword';

GRANT ALL PRIVILEGES ON DBScores.* TO 'devuser'@'%';
GRANT ALL PRIVILEGES ON DBScores.* TO 'devuser'@'localhost';

USE DBScores;

CREATE TABLE Scores (
    GameId TINYTEXT NOT NULL,
    Name TINYTEXT NOT NULL,
    Score INT UNSIGNED NOT NULL
);

CREATE TABLE Games (
    GameId TINYTEXT NOT NULL,
    ApiKey TINYTEXT NOT NULL,
    MaxNameLength INT UNSIGNED NOT NULL,
    MaxScore INT UNSIGNED NOT NULL,
    PRIMARY KEY (GameId)
)