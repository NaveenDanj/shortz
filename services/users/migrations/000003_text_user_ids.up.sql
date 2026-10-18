ALTER TABLE follows
    DROP CONSTRAINT follows_follower_id_fkey,
    DROP CONSTRAINT follows_followee_id_fkey;

ALTER TABLE users ALTER COLUMN id DROP DEFAULT;
ALTER TABLE users ALTER COLUMN id TYPE TEXT USING id::text;

ALTER TABLE follows
    ALTER COLUMN follower_id TYPE TEXT USING follower_id::text,
    ALTER COLUMN followee_id TYPE TEXT USING followee_id::text;

ALTER TABLE follows
    ADD CONSTRAINT follows_follower_id_fkey FOREIGN KEY (follower_id) REFERENCES users(id) ON DELETE CASCADE,
    ADD CONSTRAINT follows_followee_id_fkey FOREIGN KEY (followee_id) REFERENCES users(id) ON DELETE CASCADE;
