CREATE TABLE IF NOT EXISTS users (
    id uuid PRIMARY KEY,
    username VARCHAR(50) NOT NULL,
    password VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS terminals (
    id uuid PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    longitude DECIMAL(9,6) NOT NULL,
    latitude DECIMAL(9,6) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS transactions (
    id uuid PRIMARY KEY,
    amount DECIMAL(9,2) NOT NULL,
    check_in_terminal_id uuid NOT NULL,
    check_out_terminal_id uuid,
    check_in_timestamp TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    check_out_timestamp TIMESTAMP,
    CONSTRAINT   fk_check_in_terminal_id FOREIGN KEY (check_in_terminal_id) REFERENCES terminals(id),
    CONSTRAINT   fk_check_out_terminal_id FOREIGN KEY (check_out_terminal_id) REFERENCES terminals(id)
);
