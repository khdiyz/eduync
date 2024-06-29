-- +goose Up
CREATE TABLE "roles" (
    "id" BIGSERIAL PRIMARY KEY,
    "name" VARCHAR(64) NOT NULL,
    "description" TEXT,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP,
    "deleted_at" TIMESTAMP
);

CREATE UNIQUE INDEX unique_role_name_not_deleted ON "roles" ("name") WHERE "deleted_at" IS NULL;

INSERT INTO "roles" ("name") VALUES 
('SUPER ADMIN'),
('ADMIN'),
('MENTOR');

-- +goose Down
DROP TABLE "roles";