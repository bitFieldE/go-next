DO $$ 
BEGIN
  IF NOT EXISTS (SELECT 1 FROM pg_database WHERE datname = 'go_next_backend') THEN
    CREATE DATABASE go_next_backend;
  END IF;
END $$;