-- swith to supportdb

CREATE TABLE IF NOT EXISTS tickets (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(255),
    last_name VARCHAR(255),
    email VARCHAR(255),
    issue VARCHAR(255),
    priority VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO tickets
    (first_name, last_name, email, issue, priority)
VALUES
    ('John', 'Doe', 'yL1pB@example.com', 'My issue', 'High'),
    ('Jane', 'Doe', 'jyL1pB@example.com', 'My issue', 'Medium'),
    ('Jill', 'Doe', 'jyL1pB@example.com', 'My issue', 'Low');