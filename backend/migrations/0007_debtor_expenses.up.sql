-- Add debtorId column to expenses table for linking debt payments
ALTER TABLE expenses ADD COLUMN debtorId INTEGER REFERENCES debtors(id) ON DELETE SET NULL;

-- Create index for faster lookups
CREATE INDEX idx_expenses_debtorid ON expenses(debtorId) WHERE debtorId IS NOT NULL;
