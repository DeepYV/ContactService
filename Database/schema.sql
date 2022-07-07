CREATE TABLE public.contact
(
    "Number" text,
    "Name" text,
    PRIMARY KEY ("Number")
);

ALTER TABLE IF EXISTS public.contact
    OWNER to postgres;