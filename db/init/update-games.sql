-- Replace these values with the details you want to add. --
USE DBScores;

INSERT INTO Games (GameId, ApiKey, MaxNameLength, MaxScore) VALUES ('gameId1', 'apikey1', 3, 100) ON DUPLICATE KEY UPDATE ApiKey = 'apiKey1', MaxNameLength = 3, MaxScore = 100;
INSERT INTO Games (GameId, ApiKey, MaxNameLength, MaxScore) VALUES ('gameId2', 'apikey2', 5, 500) ON DUPLICATE KEY UPDATE ApiKey = 'apiKey2', MaxNameLength = 5, MaxScore = 500;