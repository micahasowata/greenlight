CREATE TABLE IF NOT EXISTS permissions (
    id bigserial PRIMARY KEY,
    code text NOT NULL
);

CREATE TABLE IF NOT EXISTS users_permissions (
    user_id bigint REFERENCES users ON DELETE CASCADE,
    permission_id bigint REFERENCES permissions ON DELETE CASCADE PRIMARY KEY (user_id, permission_id)
);

INSERT INTO
    permissions (code)
VALUES
    ('movies:read'),
    ('movies:write');