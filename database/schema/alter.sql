--------------------------------------------------------------------------------------------------

ALTER TABLE IF EXISTS public.professions
    OWNER to myuser;

ALTER SEQUENCE public.professions_id_seq
    OWNED BY public.professions.id;

ALTER SEQUENCE public.professions_id_seq
    OWNER TO myuser;

--------------------------------------------------------------------------------------------------
