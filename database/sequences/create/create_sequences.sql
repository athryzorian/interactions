-- SEQUENCE: public.business_user_id_seq

-- DROP SEQUENCE IF EXISTS public.business_user_id_seq;

CREATE SEQUENCE IF NOT EXISTS public.business_user_id_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 2147483647
    CACHE 1;

-- SEQUENCE: public.business_user_roles_id_seq

-- DROP SEQUENCE IF EXISTS public.business_user_roles_id_seq;

CREATE SEQUENCE IF NOT EXISTS public.business_user_roles_id_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 2147483647
    CACHE 1;

-- SEQUENCE: public.countries_id_seq

-- DROP SEQUENCE IF EXISTS public.countries_id_seq;

CREATE SEQUENCE IF NOT EXISTS public.countries_id_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 2147483647
    CACHE 1;

-- SEQUENCE: public.customer_user_id_seq

-- DROP SEQUENCE IF EXISTS public.customer_user_id_seq;

CREATE SEQUENCE IF NOT EXISTS public.customer_user_id_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 2147483647
    CACHE 1;

-- SEQUENCE: public.customer_user_roles_id_seq

-- DROP SEQUENCE IF EXISTS public.customer_user_roles_id_seq;

CREATE SEQUENCE IF NOT EXISTS public.customer_user_roles_id_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 2147483647
    CACHE 1;

-- SEQUENCE: public.organization_designations_id_seq

-- DROP SEQUENCE IF EXISTS public.organization_designations_id_seq;

CREATE SEQUENCE IF NOT EXISTS public.organization_designations_id_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 2147483647
    CACHE 1;

-- SEQUENCE: public.organizations_id_seq

-- DROP SEQUENCE IF EXISTS public.organizations_id_seq;

CREATE SEQUENCE IF NOT EXISTS public.organizations_id_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 2147483647
    CACHE 1;

-- SEQUENCE: public.professions_id_seq

-- DROP SEQUENCE IF EXISTS public.professions_id_seq;

CREATE SEQUENCE IF NOT EXISTS public.professions_id_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 2147483647
    CACHE 1;

-- SEQUENCE: public.user_types_id_seq

-- DROP SEQUENCE IF EXISTS public.user_types_id_seq;

CREATE SEQUENCE IF NOT EXISTS public.user_types_id_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 2147483647
    CACHE 1;