-- Seed Script: SuperAdmin, Pages, Roles, and Permissions
-- Run this after migrations to set up initial data

-- =============================================
-- 1. INSERT PAGES (all routes from your system)
-- =============================================

INSERT INTO pages (name, route, icon, status) VALUES
('Projects', '/projects', 'folder', 'active'),
('Workdays', '/workdays', 'calendar', 'active'),
('Workday Materials', '/workday-materials', 'package', 'active'),
('Workday Labor', '/workday-labor', 'users', 'active'),
('Workday Equipment', '/workday-equipment', 'tool', 'active'),
('Work Categories', '/work-categories', 'layers', 'active'),
('Work Subcategories', '/work-subcategories', 'list', 'active'),
('Expenses', '/expenses', 'credit-card', 'active'),
('Income', '/income', 'dollar-sign', 'active'),
('Debtors', '/debtors', 'user-minus', 'active'),
('Users', '/users', 'users', 'active'),
('User Roles', '/user-roles', 'user-check', 'active'),
('Roles', '/roles', 'shield', 'active'),
('Pages', '/pages', 'file-text', 'active'),
('Role Pages', '/role-pages', 'link', 'active')
ON CONFLICT DO NOTHING;

-- =============================================
-- 2. INSERT SUPERADMIN ROLE
-- =============================================

INSERT INTO roles (name, description) VALUES
('SuperAdmin', 'Full system access - can manage all resources and permissions')
ON CONFLICT DO NOTHING;

-- =============================================
-- 3. ASSIGN ALL PERMISSIONS TO SUPERADMIN
-- =============================================

-- Get the SuperAdmin role ID and assign full permissions to all pages
DO $$
DECLARE
    superadmin_role_id INTEGER;
    page_record RECORD;
    permissions_json VARCHAR(255);
BEGIN
    -- Get SuperAdmin role ID
    SELECT id INTO superadmin_role_id FROM roles WHERE name = 'SuperAdmin';

    -- Loop through all pages and assign full permissions
    FOR page_record IN SELECT id, route FROM pages LOOP
        -- Set permissions based on page type
        IF page_record.route = '/users' THEN
            permissions_json := '["read","write","delete","updatePassword","updateStatus"]';
        ELSE
            permissions_json := '["read","write","delete"]';
        END IF;

        -- Insert role-page mapping (skip if exists)
        INSERT INTO rolePages (roleId, pageId, permissions)
        VALUES (superadmin_role_id, page_record.id, permissions_json)
        ON CONFLICT DO NOTHING;
    END LOOP;
END $$;

-- =============================================
-- 4. CREATE SUPERADMIN USER
-- =============================================

-- Password: Admin@123 (bcrypt hashed)
-- You should change this password after first login!
INSERT INTO users (username, email, password, fullName, phone, jobTitle, status) VALUES
('superadmin', 'admin@example.com', '$2a$10$N9qo8uLOickgx2ZMRZoMy.MqrqFnTqwy0Rl/NrJHkE9R8FBKX7g.C', 'Super Admin', '0000000000', 'System Administrator', 'active')
ON CONFLICT (username) DO NOTHING;

-- =============================================
-- 5. ASSIGN SUPERADMIN ROLE TO SUPERADMIN USER
-- =============================================

DO $$
DECLARE
    superadmin_user_id INTEGER;
    superadmin_role_id INTEGER;
BEGIN
    SELECT id INTO superadmin_user_id FROM users WHERE username = 'superadmin';
    SELECT id INTO superadmin_role_id FROM roles WHERE name = 'SuperAdmin';

    -- Assign role to user (skip if exists)
    INSERT INTO userRoles (userId, roleId)
    VALUES (superadmin_user_id, superadmin_role_id)
    ON CONFLICT DO NOTHING;
END $$;

-- =============================================
-- VERIFICATION QUERIES (optional - run to verify)
-- =============================================

-- Check pages
-- SELECT * FROM pages;

-- Check roles
-- SELECT * FROM roles;

-- Check role-page permissions
-- SELECT r.name as role, p.name as page, rp.permissions
-- FROM rolePages rp
-- JOIN roles r ON rp.roleId = r.id
-- JOIN pages p ON rp.pageId = p.id;

-- Check user roles
-- SELECT u.username, r.name as role
-- FROM userRoles ur
-- JOIN users u ON ur.userId = u.id
-- JOIN roles r ON ur.roleId = r.id;
