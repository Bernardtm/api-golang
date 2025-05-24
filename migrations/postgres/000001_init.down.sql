-- Drop the GRANT EXECUTE privileges on the functions in the default_schema
DO $$
DECLARE
    grant_cmd TEXT;
BEGIN
    FOR grant_cmd IN 
        SELECT 'REVOKE EXECUTE ON FUNCTION ' || ns.nspname || '.' || p.proname || '(' || oidvectortypes(p.proargtypes) || ') FROM PUBLIC;'
        FROM pg_proc p
        JOIN pg_namespace ns ON ns.oid = p.pronamespace
        WHERE ns.nspname = 'default_schema'
    LOOP
        EXECUTE grant_cmd;
    END LOOP;
END;
$$;

-- Drop the uuid_generate_v7 function
DROP FUNCTION IF EXISTS default_schema.uuid_generate_v7();

-- Drop the extension if it is no longer used
DROP EXTENSION IF EXISTS "uuid-ossp";

-- Drop the schema
DROP SCHEMA IF EXISTS default_schema CASCADE;