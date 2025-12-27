-- Force the migration version to 6 (since we verified tables exist) and clear dirty flag
UPDATE schema_migrations SET version = 6, dirty = false;
