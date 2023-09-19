-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

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
-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- DROP CONSTRAINT fk_check_in_terminal_id ON transactions;
-- DROP CONSTRAINT fk_check_out_terminal_id ON transactions;
DROP TABLE IF EXISTS transactions;;
-- +goose StatementEnd