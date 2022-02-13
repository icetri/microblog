CREATE TABLE blog
(
    id        serial
        constraint blog_pk
            primary key,
    title      text default ''::character varying not null,
    anous     text default ''::character varying not null,
    full_text text default ''::character varying not null,
    datenow      text default ''::character varying not null,
    username  text default ''::character varying not null
);

CREATE UNIQUE INDEX blog_id_uindex
    ON blog (id);