-- This is for the local dev setup of this project.
-- Make sure to follow the setup docs when being used in production.
CREATE DATABASE DBScores;

CREATE USER 'devuser'@'localhost' IDENTIFIED BY 'devpassword';

GRANT ALL PRIVILEGES ON DBScores.* TO 'devuser'@'%';
GRANT ALL PRIVILEGES ON DBScores.* TO 'devuser'@'localhost';

USE DBScores;

CREATE TABLE Scores (
    GameId TINYTEXT,
    Name TINYTEXT,
    Score INT UNSIGNED
);

CREATE TABLE Games (
    GameId TINYTEXT
    ApiKey TINYTEXT
    MaxNameLength INT UNSIGNED
    MaxScore INT UNSIGNED
)