-- SEQUENCE: public.business_user_id_seq

-- DROP SEQUENCE IF EXISTS public.business_user_id_seq;

ALTER SEQUENCE public.business_user_id_seq
    OWNED BY public.business_user.id;

ALTER SEQUENCE public.business_user_id_seq
    OWNER TO myuser;

-- SEQUENCE: public.business_user_roles_id_seq

-- DROP SEQUENCE IF EXISTS public.business_user_roles_id_seq;

ALTER SEQUENCE public.business_user_roles_id_seq
    OWNED BY public.business_user_roles.id;

ALTER SEQUENCE public.business_user_roles_id_seq
    OWNER TO myuser;

-- SEQUENCE: public.countries_id_seq

-- DROP SEQUENCE IF EXISTS public.countries_id_seq;

ALTER SEQUENCE public.countries_id_seq
    OWNED BY public.countries.id;

ALTER SEQUENCE public.countries_id_seq
    OWNER TO myuser;

-- SEQUENCE: public.customer_user_id_seq

-- DROP SEQUENCE IF EXISTS public.customer_user_id_seq;

ALTER SEQUENCE public.customer_user_id_seq
    OWNED BY public.customer_users.id;

ALTER SEQUENCE public.customer_user_id_seq
    OWNER TO myuser;

-- SEQUENCE: public.customer_user_roles_id_seq

-- DROP SEQUENCE IF EXISTS public.customer_user_roles_id_seq;

ALTER SEQUENCE public.customer_user_roles_id_seq
    OWNED BY public.customer_user_roles.id;

ALTER SEQUENCE public.customer_user_roles_id_seq
    OWNER TO myuser;

-- SEQUENCE: public.organization_designations_id_seq

-- DROP SEQUENCE IF EXISTS public.organization_designations_id_seq;

ALTER SEQUENCE public.organization_designations_id_seq
    OWNED BY public.org_designations.id;

ALTER SEQUENCE public.organization_designations_id_seq
    OWNER TO myuser;

-- SEQUENCE: public.organizations_id_seq

-- DROP SEQUENCE IF EXISTS public.organizations_id_seq;

ALTER SEQUENCE public.organizations_id_seq
    OWNED BY public.organizations.id;

ALTER SEQUENCE public.organizations_id_seq
    OWNER TO myuser;

-- SEQUENCE: public.professions_id_seq

-- DROP SEQUENCE IF EXISTS public.professions_id_seq;

ALTER SEQUENCE public.professions_id_seq
    OWNED BY public.professions.id;

ALTER SEQUENCE public.professions_id_seq
    OWNER TO myuser;

-- SEQUENCE: public.user_types_id_seq

-- DROP SEQUENCE IF EXISTS public.user_types_id_seq;

ALTER SEQUENCE public.user_types_id_seq
    OWNED BY public.user_types.id;

ALTER SEQUENCE public.user_types_id_seq
    OWNER TO myuser;