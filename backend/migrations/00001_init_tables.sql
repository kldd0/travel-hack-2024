-- +goose Up
-- +goose StatementBegin

create table if not exists tours
(
    id               serial PRIMARY KEY NOT NULL,
    title            varchar(255)       NOT NULL,
    location         varchar(255)       NOT NULL,
    category         varchar(255)       NOT NULL,
    tags             varchar[]          NOT NULL,
    description      text,
    nights_count      int,
    program          text[],
    included         text[],
    not_included     text[],
    difficulty_level varchar(255)       NOT NULL,
    comfort_level    varchar(255)       NOT NULL,
    tour_dates       date[],
    important_info   text,
    head_media       text[],
    program_media    text[],
    acc_media        text[],
    map              text,
    faq              text,
    rating           int
);

create unique index projects_id_idx on tours (id);

create table if not exists reviews (
    id             serial PRIMARY KEY NOT NULL,
    tour_id        int,
    liked          varchar[],
    username       varchar(255),
    positive       text,
    negative       text,
    local_resident boolean,
    type           varchar(255),
    frequency      varchar(255),
    video          text,
    foreign key (tour_id) references tours (id)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

drop table if exists reviews;

drop table if exists tours;

-- +goose StatementEnd
