create table public.user_roles
(
    id         bigserial
        primary key,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    user_id    uuid   not null
        constraint fk_user_roles_user
            references public.users,
    role_id    bigint not null
        constraint fk_user_roles_role
            references public.roles
);

create index idx_user_roles_deleted_at
    on public.user_roles (deleted_at);

