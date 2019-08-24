--CREATE TABLE IF NOT EXISTS book(
--   ISBN int PRIMARY KEY,
--   Name varchar(255) NOT NULL,
--   Description varchar(255) NOT NULL,
--   Author varchar(255) NOT NULL,
--   Rating int NOT NULL
--;

CREATE TABLE public.book
(
    "ISBN" integer NOT NULL,
    "Name" text COLLATE pg_catalog."default" NOT NULL,
    "Description" text COLLATE pg_catalog."default" NOT NULL,
    "Author" text COLLATE pg_catalog."default" NOT NULL,
    "Rating" integer NOT NULL,
    CONSTRAINT book_pkey PRIMARY KEY ("ISBN")
)
WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;

ALTER TABLE public.book
    OWNER to postgres;