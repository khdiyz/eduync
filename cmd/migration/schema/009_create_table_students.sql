-- +goose Up
CREATE TABLE "students" (
    "id" BIGSERIAL PRIMARY KEY,
    "full_name" VARCHAR(64) NOT NULL,
    "phone_number_1" VARCHAR(64) NOT NULL,
    "phone_number_2" VARCHAR(64),
    "description" TEXT,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP,
    "deleted_at" TIMESTAMP
);

-- +goose Down
DROP TABLE "students";