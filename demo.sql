create table user_infos (
    id bigserial primary key not null,
    name varchar(128) not null default '',
    create_time timestamp not null default current_timestamp
);

create unique index uk_user_name on user_infos(name);

create type match_state as enum ('liked', 'disliked', 'matched');

create table relationships (
    id bigserial primary key not null,
    user_id bigint not null default 0,
    other_user_id bigint not null default 0,
    state match_state not null default 'liked',
    create_time timestamp not null default current_timestamp
);

create unique index uk_relationship_pair on relationships(user_id, other_user_id);