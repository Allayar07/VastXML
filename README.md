create table vast
(
    id varchar,
    title varchar not null,
    is_Skipable boolean not null,
    skip_time integer,
    adsDuration integer not null,
    uri varchar
);
