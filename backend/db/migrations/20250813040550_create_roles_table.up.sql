create table public.roles
(
    id   bigserial
        primary key,
    name text not null
        constraint uni_roles_name
            unique
);