-- +goose Up
CREATE TABLE "lids" (
    "id" BIGSERIAL PRIMARY KEY,
    "full_name" VARCHAR(64) NOT NULL,
    "phone_number" VARCHAR(64), 
    "course_id" BIGINT NOT NULL,
    "description" TEXT,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP,
    "deleted_at" TIMESTAMP,
    FOREIGN KEY ("course_id") REFERENCES "courses"("id")
);

-- +goose Down
DROP TABLE "lids";