CREATE DATABASE DBScores;
GRANT ALL PRIVILEGES ON DBScores.* TO 'devuser'@'%' IDENTIFIED BY 'mysql';
GRANT ALL PRIVILEGES ON DBScores.* TO 'devuser'@'localhost' IDENTIFIED BY 'mysql';

USE DBScores

CREATE TABLE Scores (
    [GameId] TINYTEXT
    [Name] TINYTEXT
    [Score] UNSIGNED INT
)