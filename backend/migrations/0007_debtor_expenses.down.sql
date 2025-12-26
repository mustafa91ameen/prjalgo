-- Remove debtorId column from expenses table
DROP INDEX IF EXISTS idx_expenses_debtorid;
ALTER TABLE expenses DROP COLUMN IF EXISTS debtorId;
