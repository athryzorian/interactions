
--------------------------------------------------------------------------------------------------
-- SEQUENCE: public.professions_id_seq

-- DROP SEQUENCE IF EXISTS public.professions_id_seq;

CREATE SEQUENCE IF NOT EXISTS public.professions_id_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 2147483647
    CACHE 1;


-- Table: public.professions

-- DROP TABLE IF EXISTS public.professions;

CREATE TABLE IF NOT EXISTS public.professions
(
    id integer NOT NULL DEFAULT nextval('professions_id_seq'::regclass),
    name character varying COLLATE pg_catalog."default" NOT NULL,
    abbreviation character varying COLLATE pg_catalog."default",
    logo bytea,
    is_enabled boolean NOT NULL DEFAULT false,
    description character varying COLLATE pg_catalog."default",
    CONSTRAINT professions_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

--------------------------------------------------------------------------------------------------

