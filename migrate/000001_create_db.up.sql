CREATE ROLE wibuser WITH LOGIN PASSWORD 'wibpass';


CREATE DATABASE wibdb
    WITH
    OWNER = wibuser
    ENCODING = 'UTF8'
    CONNECTION LIMIT = -1
    IS_TEMPLATE = False;

CREATE TABLE data_links
(
    key_link varchar(36)  default ''::character varying not null
        constraint data_links_pk
            primary key,
    key_data varchar(150) default ''::character varying,
    created_at timestamptz
    type     integer      default 0                     not null,
    payload  jsonb                                      not null
);

ALTER TABLE data_links
    OWNER TO wibuser;
