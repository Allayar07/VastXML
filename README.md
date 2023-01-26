create table vast
(
    id serial not null,
    title varchar not null,
    is_Skipable boolean not null,
    skip_time integer,
    ads_height integer,
    ads_width integer,
    adsDuration integer not null,
    uri varchar
);
