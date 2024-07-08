create table urls (
    id serial primary key,
    "url" text not null,
    alias text not null unique
);