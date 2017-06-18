CREATE SEQUENCE public.test_id_seq
  INCREMENT 1
  MINVALUE 1
  MAXVALUE 9223372036854775807
  START 1
  CACHE 1;
ALTER TABLE public.test_id_seq
  OWNER TO postgres;


CREATE TABLE public.test
(
  id bigint NOT NULL DEFAULT nextval('test_id_seq'::regclass),
  uid uuid NOT NULL DEFAULT uuid_generate_v4(),
  data character varying(100) NOT NULL,
  CONSTRAINT pk_test PRIMARY KEY (id),
  CONSTRAINT ux_test_uid UNIQUE (uid)
)
WITH (
  OIDS=FALSE
);
ALTER TABLE public.test
  OWNER TO postgres;
