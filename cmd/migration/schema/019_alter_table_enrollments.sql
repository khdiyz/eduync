-- +goose Up
CREATE TYPE enrollment_status AS ENUM (
    'ACTIVE',
    'INACTIVE',
    'FROZEN'
);

ALTER TABLE "enrollments" ADD COLUMN "status" enrollment_status NOT NULL;

CREATE UNIQUE INDEX unique_student_group_deleted
ON enrollments (student_id, group_id)
WHERE deleted_at IS NULL;

-- +goose Down
DROP INDEX unique_student_group_deleted;

ALTER TABLE "enrollments" DROP COLUMN "status";

DROP TYPE "enrollment_status";