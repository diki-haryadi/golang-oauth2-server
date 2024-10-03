CREATE TABLE buyback_calculations (
    id UUID DEFAULT uuid_generate_v4 () PRIMARY KEY,  -- Unique identifier for each calculation
    purchase_uid UUID NOT NULL REFERENCES purchases(id) ON DELETE CASCADE, -- Reference to the purchase
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,    -- Creation timestamp
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
