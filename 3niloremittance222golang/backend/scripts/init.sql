-- Enable UUID extension for generating UUIDs
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Create custom ENUM types
DO $$ BEGIN
    CREATE TYPE "Role" AS ENUM ('DIASPORA_USER', 'MERCHANT_ADMIN', 'PLATFORM_ADMIN');
EXCEPTION
    WHEN duplicate_object THEN null;
END $$;

DO $$ BEGIN
    CREATE TYPE "BusinessType" AS ENUM ('SCHOOL', 'HOSPITAL', 'SUPERMARKET', 'PHARMACY', 'UTILITY_COMPANY');
EXCEPTION
    WHEN duplicate_object THEN null;
END $$;

DO $$ BEGIN
    CREATE TYPE "TransactionStatus" AS ENUM ('PENDING', 'PROCESSING', 'COMPLETED', 'FAILED', 'REFUNDED', 'CANCELLED');
EXCEPTION
    WHEN duplicate_object THEN null;
END $$;

DO $$ BEGIN
    CREATE TYPE "PayoutStatus" AS ENUM ('PENDING', 'PROCESSING', 'COMPLETED', 'FAILED');
EXCEPTION
    WHEN duplicate_object THEN null;
END $$;

-- Create a sample audit function
CREATE OR REPLACE FUNCTION log_audit()
RETURNS TRIGGER AS $$
BEGIN
    INSERT INTO audit_logs (user_id, action, resource, resource_id, details, ip_address, user_agent)
    VALUES (
        current_setting('myapp.current_user', true)::text,
        TG_OP,
        TG_TABLE_NAME,
        NEW.id,
        jsonb_build_object('new', row_to_json(NEW), 'old', row_to_json(OLD)),
        current_setting('myapp.client_ip', true),
        current_setting('myapp.user_agent', true)
    );
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;
