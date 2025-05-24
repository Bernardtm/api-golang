CREATE SCHEMA default_schema;


CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

create or replace function default_schema.uuid_generate_v7()
returns uuid
as $$
select encode(
    set_bit(
      set_bit(
        overlay(uuid_send(gen_random_uuid())
                placing substring(int8send(floor(extract(epoch from clock_timestamp()) * 1000)::bigint) from 3)
                from 1 for 6
        ),
        52, 1
      ),
      53, 1
    ),
    'hex')::uuid;
$$
language SQL
volatile;


DO $$
BEGIN
    EXECUTE (
        SELECT string_agg(
            'GRANT EXECUTE ON FUNCTION ' || ns.nspname || '.' || p.proname || '(' || oidvectortypes(p.proargtypes) || ') TO PUBLIC;',
            ' '
        )
        FROM pg_proc p
        JOIN pg_namespace ns ON ns.oid = p.pronamespace
        WHERE ns.nspname = 'default_schema'
    );
END;
$$;