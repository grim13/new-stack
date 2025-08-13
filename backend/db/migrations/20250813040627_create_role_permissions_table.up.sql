create table public.role_permissions
(
    id            bigserial
        primary key,
    created_at    timestamp with time zone,
    updated_at    timestamp with time zone,
    deleted_at    timestamp with time zone,
    role_id       bigint not null
        constraint fk_role_permissions_role
            references public.roles,
    permission_id bigint not null
        constraint fk_role_permissions_permission
            references public.permissions
);

create index idx_role_permissions_deleted_at
    on public.role_permissions (deleted_at);

