-- Migration: seed initial admin user (idempotent)
-- Default credentials:
-- email: admin@example.com
-- password: admin123

INSERT INTO users (email, password_hash, full_name, role, is_active)
VALUES (
    'admin@example.com',
    crypt('admin123', gen_salt('bf')),
    'System Admin',
    'admin',
    TRUE
)
ON CONFLICT (email) DO NOTHING;
