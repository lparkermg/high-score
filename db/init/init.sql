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