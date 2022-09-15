CREATE TABLE accounts (
    -- Internal account ID. Never changes.
    account_id UUID NOT NULL DEFAULT gen_random_uuid() PRIMARY KEY,

    -- Identity used to tied this account to OIDC.
    -- OpenID Connect Core, 2. ID Token: “It MUST NOT exceed 255 ASCII
    -- characters in length”.
    account_oidc_sub STRING(255) NOT NULL UNIQUE,

    --- Copy/cache of user data retrieved from OIDC IdP on login. Currently this
    --- is only updated on first login, but we should find a way to trigger
    --- a re-retrieval.
    -- Display name preferred by user.
    -- Self-limiting ourselves to 255 unicode codepoints here. This is also
    -- supposedly what keycloak also defaults to for user attributes.
    account_display_name STRING(255) NOT NULL
);