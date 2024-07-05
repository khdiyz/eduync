-- +goose Up
ALTER TABLE "students" ADD COLUMN "birth_year" VARCHAR(4);

ALTER TABLE "students" RENAME COLUMN "phone_number_1" TO "phone_number";
ALTER TABLE "students" RENAME COLUMN "phone_number_2" TO "parent_phone";
ALTER TABLE "students" RENAME COLUMN "description" TO "address";

-- +goose Down
ALTER TABLE "students" RENAME COLUMN "address" TO "description";
ALTER TABLE "students" RENAME COLUMN "parent_phone" TO "phone_number_2";
ALTER TABLE "students" RENAME COLUMN "phone_number" TO "phone_number_1";

ALTER TABLE "students" DROP COLUMN "birth_year";

