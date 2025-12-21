CREATE TABLE preferences (
    id SERIAL PRIMARY KEY,
    user_id VARCHAR(100) NOT NULL,
    channel VARCHAR(50) NOT NULL,
    enabled BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    UNIQUE(user_id, channel)
);

CREATE INDEX idx_preferences_user ON preferences(user_id);