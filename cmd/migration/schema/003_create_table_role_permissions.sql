-- +goose Up
CREATE TABLE "role_permissions" (
    "id" BIGSERIAL PRIMARY KEY,
    "role_id" BIGINT NOT NULL,
    "permission_id" BIGINT NOT NULL,
    "status" BOOLEAN NOT NULL,
    FOREIGN KEY ("role_id") REFERENCES "roles"("id"),
    FOREIGN KEY ("permission_id") REFERENCES "permissions"("id")
);

-- +goose Down
DROP TABLE "role_permissions";