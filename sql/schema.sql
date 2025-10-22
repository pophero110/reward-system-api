CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    email TEXT NOT NULL UNIQUE,
    hashed_password TEXT NOT NULL,
    role TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS quests (
    id SERIAL PRIMARY KEY,
    quest_board_id INT NOT NULL,
    title TEXT NOT NULL,
    description TEXT,
    due_date_time TIMESTAMPTZ,
    comments TEXT,
    reward TEXT,
    status VARCHAR(50) NOT NULL DEFAULT 'available',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_quests_quest_board_id 
ON quests(quest_board_id);

CREATE OR REPLACE FUNCTION update_timestamp()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Users table trigger
DROP TRIGGER IF EXISTS trigger_users_updated_at ON users;
CREATE TRIGGER trigger_users_updated_at
BEFORE UPDATE ON users
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();

-- Quests table trigger
DROP TRIGGER IF EXISTS trigger_quests_updated_at ON quests;
CREATE TRIGGER trigger_quests_updated_at
BEFORE UPDATE ON quests
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();
