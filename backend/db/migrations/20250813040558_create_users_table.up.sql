create table public.users
(
    id         uuid not null primary key,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name       text not null,
    username   text not null
        constraint uni_users_username
            unique,
    email      text not null
        constraint uni_users_email
            unique,
    password   text not null,
    role_id    bigint
        constraint fk_users_role
            references public.roles
);

create index idx_users_deleted_at
    on public.users (deleted_at);

