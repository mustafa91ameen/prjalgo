-- Refresh Tokens Table
CREATE TABLE refreshTokens (
    id SERIAL PRIMARY KEY,
    userId INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    tokenHash VARCHAR(255) NOT NULL UNIQUE,
    expiresAt TIMESTAMP NOT NULL,
    revoked BOOLEAN DEFAULT FALSE,
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_refreshTokens_userId ON refreshTokens(userId);
CREATE INDEX idx_refreshTokens_tokenHash ON refreshTokens(tokenHash);
