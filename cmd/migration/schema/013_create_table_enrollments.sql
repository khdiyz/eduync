-- +goose Up
CREATE TABLE "enrollments" (
    "id" BIGSERIAL PRIMARY KEY,
    "student_id" BIGINT NOT NULL,
    "group_id" BIGINT NOT NULL,
    "join_date" DATE NOT NULL,
    "left_date" DATE,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP,
    "deleted_at" TIMESTAMP,
    FOREIGN KEY ("student_id") REFERENCES "students"("id"),
    FOREIGN KEY ("group_id") REFERENCES "groups"("id")
);

-- +goose Down
DROP TABLE "enrollments";