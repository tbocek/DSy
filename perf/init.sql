-- Create users table
CREATE TABLE IF NOT EXISTS users (
                                     id SERIAL PRIMARY KEY,
                                     username VARCHAR(50) NOT NULL,
    email VARCHAR(100) NOT NULL
    );

-- Insert sample data
INSERT INTO users (username, email)
VALUES
    ('john_doe', 'john@example.com'),
    ('jane_smith', 'jane@example.com'),
    ('bob_johnson', 'bob@example.com'),
    ('alice_wonder', 'alice@example.com'),
    ('charlie_brown', 'charlie@example.com');