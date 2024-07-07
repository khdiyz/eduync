-- +goose Up
ALTER TABLE "student_actions" DROP COLUMN "action_text";

ALTER TABLE "student_actions" ADD COLUMN "group_id" BIGINT NOT NULL;

ALTER TABLE "student_actions" DROP COLUMN "action_time";

-- +goose Down
ALTER TABLE "student_actions" ADD COLUMN "action_time" TIME NOT NULL;

ALTER TABLE "student_actions" DROP COLUMN "group_id";

ALTER TABLE "student_actions" ADD COLUMN "action_text" TEXT;


