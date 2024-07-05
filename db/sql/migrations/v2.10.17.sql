ALTER TABLE runner ADD column IF NOT EXISTS inventory_id INT REFERENCES project__inventory(id) ON DELETE SET NULL;
