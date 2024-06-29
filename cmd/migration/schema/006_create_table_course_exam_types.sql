-- +goose Up
CREATE TABLE "course_exam_types" (
    "id" BIGSERIAL PRIMARY KEY,
    "course_id" BIGINT NOT NULL,
    "name" VARCHAR(64) NOT NULL,
    "max_ball" FLOAT NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP,
    "deleted_at" TIMESTAMP,
    FOREIGN KEY ("course_id") REFERENCES "courses"("id")
);

CREATE UNIQUE INDEX unique_exam_name_not_deleted ON "course_exam_types" ("name") WHERE "deleted_at" IS NULL;

-- +goose Down
DROP TABLE "course_exam_types";