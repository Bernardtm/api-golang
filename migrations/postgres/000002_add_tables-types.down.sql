-- Drop the tables in reverse order to avoid dependency issues
DROP TABLE IF EXISTS default_schema.menus CASCADE;


-- Drop the status table last, as it is referenced by other tables
DROP TABLE IF EXISTS default_schema.status CASCADE;
