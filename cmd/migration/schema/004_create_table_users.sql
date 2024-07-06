-- +goose Up
CREATE TABLE "users" (
    "id" BIGSERIAL PRIMARY KEY,
    "full_name" VARCHAR(64) NOT NULL,
    "phone_number" VARCHAR(64) NOT NULL,
    "birth_date" DATE NOT NULL,
    "role_id" BIGINT NOT NULL,
    "username" VARCHAR(64) NOT NULL,
    "password" TEXT NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP,
    "deleted_at" TIMESTAMP,
    FOREIGN KEY ("role_id") REFERENCES "roles"("id")
);

CREATE UNIQUE INDEX unique_username_not_deleted ON "users" ("username") WHERE "deleted_at" IS NULL;

INSERT INTO "users" (
    "full_name",
    "phone_number",
    "birth_date",
    "role_id",
    "username",
    "password"
) VALUES (
    'Super Admin',
    '+998901234567',
    CURRENT_DATE::DATE,
    (SELECT "id" FROM "roles" WHERE "name" = 'SUPER ADMIN'),
    'superadmin',
    '6868616833326975734848736b6a64731481b6ea0a413bbee37199d0483c70388be8ea4c'
);

-- +goose Down
DROP TABLE "users";