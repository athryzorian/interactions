-- Table: public.business_user_roles

-- DROP TABLE IF EXISTS public.business_user_roles;

CREATE TABLE IF NOT EXISTS public.business_user_roles
(
    role character varying COLLATE pg_catalog."default" NOT NULL,
    description character varying COLLATE pg_catalog."default" NOT NULL,
    id integer NOT NULL DEFAULT nextval('business_user_roles_id_seq'::regclass),
    CONSTRAINT business_user_roles_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.business_user_roles
    OWNER to myuser;

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
    CONSTRAINT professions_pkey PRIMARY KEY (id),
    CONSTRAINT professions_name_abbreviation_key UNIQUE (name)
        INCLUDE(abbreviation)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.professions
    OWNER to myuser;

-- Table: public.organizations

-- DROP TABLE IF EXISTS public.organizations;

CREATE TABLE IF NOT EXISTS public.organizations
(
    id integer NOT NULL DEFAULT nextval('organizations_id_seq'::regclass),
    name character varying COLLATE pg_catalog."default" NOT NULL,
    profession_id integer NOT NULL,
    description character varying COLLATE pg_catalog."default",
    created_on date NOT NULL,
    updated_on date NOT NULL DEFAULT CURRENT_DATE,
    created_by integer NOT NULL,
    updated_by integer NOT NULL,
    CONSTRAINT organizations_pkey PRIMARY KEY (id),
    CONSTRAINT created_by FOREIGN KEY (created_by)
        REFERENCES public.organizations (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID,
    CONSTRAINT profession_id FOREIGN KEY (profession_id)
        REFERENCES public.professions (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID,
    CONSTRAINT updated_by FOREIGN KEY (updated_by)  
        REFERENCES public.organizations (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.organizations
    OWNER to myuser;


-- Table: public.business_user

-- DROP TABLE IF EXISTS public.business_user;

CREATE TABLE IF NOT EXISTS public.business_user
(
    id integer NOT NULL DEFAULT nextval('business_user_id_seq'::regclass),
    firebase_user_id character varying COLLATE pg_catalog."default" NOT NULL,
    parent_organization integer NOT NULL,
    business_user_role integer NOT NULL,
    created_on date NOT NULL,
    updated_on date NOT NULL,
    created_by integer NOT NULL,
    updated_by integer,
    CONSTRAINT business_user_pkey PRIMARY KEY (id),
    CONSTRAINT business_user_role FOREIGN KEY (business_user_role)
        REFERENCES public.business_user_roles (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION,
    CONSTRAINT parent_organization FOREIGN KEY (parent_organization)
        REFERENCES public.organizations (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.business_user
    OWNER to myuser;

-- Table: public.countries

-- DROP TABLE IF EXISTS public.countries;

CREATE TABLE IF NOT EXISTS public.countries
(
    id integer NOT NULL DEFAULT nextval('countries_id_seq'::regclass),
    name character varying COLLATE pg_catalog."default" NOT NULL,
    abbreviation character varying COLLATE pg_catalog."default" NOT NULL,
    country_code smallint NOT NULL,
    CONSTRAINT countries_pkey PRIMARY KEY (id),
    CONSTRAINT countries_name_abbreviation_key UNIQUE (name)
        INCLUDE(abbreviation)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.countries
    OWNER to myuser;

-- Table: public.customer_user_roles

-- DROP TABLE IF EXISTS public.customer_user_roles;

CREATE TABLE IF NOT EXISTS public.customer_user_roles
(
    id integer NOT NULL DEFAULT nextval('customer_user_roles_id_seq'::regclass),
    role character varying COLLATE pg_catalog."default" NOT NULL,
    description character varying COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT customer_user_roles_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.customer_user_roles
    OWNER to myuser;

-- Table: public.customer_users

-- DROP TABLE IF EXISTS public.customer_users;

CREATE TABLE IF NOT EXISTS public.customer_users
(
    id integer NOT NULL DEFAULT nextval('customer_user_id_seq'::regclass),
    firebase_user_id character varying COLLATE pg_catalog."default" NOT NULL,
    parent_organization integer NOT NULL,
    customer_user_role integer NOT NULL,
    created_on date NOT NULL,
    updated_on date NOT NULL,
    created_by integer NOT NULL,
    updated_by integer NOT NULL,
    CONSTRAINT customer_user_pkey PRIMARY KEY (id),
    CONSTRAINT customer_user_role FOREIGN KEY (customer_user_role)
        REFERENCES public.customer_user_roles (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID,
    CONSTRAINT parent_organization FOREIGN KEY (parent_organization)
        REFERENCES public.organizations (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.customer_users
    OWNER to myuser;

-- Table: public.org_designations

-- DROP TABLE IF EXISTS public.org_designations;

CREATE TABLE IF NOT EXISTS public.org_designations
(
    id integer NOT NULL DEFAULT nextval('organization_designations_id_seq'::regclass),
    designation character varying COLLATE pg_catalog."default" NOT NULL,
    description character varying COLLATE pg_catalog."default",
    parent_organization integer NOT NULL,
    created_on date NOT NULL,
    updated_on date NOT NULL,
    created_by integer NOT NULL,
    updated_by integer NOT NULL,
    CONSTRAINT organization_designations_pkey PRIMARY KEY (id),
    CONSTRAINT parent_organization FOREIGN KEY (parent_organization)
        REFERENCES public.organizations (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.org_designations
    OWNER to myuser;

COMMENT ON TABLE public.org_designations
    IS 'Designations offered to various employees of organization.';


-- Table: public.user_types

-- DROP TABLE IF EXISTS public.user_types;

CREATE TABLE IF NOT EXISTS public.user_types
(
    id integer NOT NULL DEFAULT nextval('user_types_id_seq'::regclass),
    name character varying COLLATE pg_catalog."default" NOT NULL,
    description character varying COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT user_types_pkey PRIMARY KEY (id),
    CONSTRAINT user_types_name_key UNIQUE (name)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.user_types
    OWNER to myuser;
