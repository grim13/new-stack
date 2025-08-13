create table public.permissions
(
    id          bigserial
        primary key,
    name        text not null
        constraint uni_permissions_name
            unique,
    description text not null
);