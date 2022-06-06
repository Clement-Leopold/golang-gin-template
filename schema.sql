-- version 1 bussiness table 
CREATE TABLE IF NOT EXISTS public.t_accounts
(
    auto_id integer NOT NULL DEFAULT nextval('t_accounts_auto_id_seq'::regclass),
    id uuid NOT NULL,
    name character varying(50) COLLATE pg_catalog."default" NOT NULL,
    dob character varying(100) COLLATE pg_catalog."default" NOT NULL,
    address character varying(255) COLLATE pg_catalog."default" NOT NULL,
    description character varying(255) COLLATE pg_catalog."default",
    created_at timestamp without time zone NOT NULL,
    x_coordinate integer NOT NULL DEFAULT 0,
    y_coordinate integer NOT NULL DEFAULT 0,
    CONSTRAINT t_accounts_pkey PRIMARY KEY (auto_id),
    CONSTRAINT t_accounts_id_key UNIQUE (id),
    CONSTRAINT t_accounts_name_key UNIQUE (name)
)


CREATE TABLE IF NOT EXISTS public.t_followings
(
    auto_id integer NOT NULL DEFAULT nextval('t_followings_auto_id_seq'::regclass),
    u_id uuid NOT NULL,
    following_id uuid NOT NULL,
    distance real NOT NULL DEFAULT 0,
    CONSTRAINT t_followings_pkey PRIMARY KEY (auto_id)
)

CREATE INDEX IF NOT EXISTS f_u_id_idx
    ON public.t_followings USING btree
    (following_id ASC NULLS LAST, u_id ASC NULLS LAST)
    TABLESPACE pg_default;

