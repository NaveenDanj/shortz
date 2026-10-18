ALTER TABLE follows
    DROP CONSTRAINT follows_follower_id_fkey,
    DROP CONSTRAINT follows_followee_id_fkey;

ALTER TABLE follows
    ALTER COLUMN follower_id TYPE UUID USING follower_id::uuid,
    ALTER COLUMN followee_id TYPE UUID USING followee_id::uuid;

ALTER TABLE users ALTER COLUMN id TYPE UUID USING id::uuid;
ALTER TABLE users ALTER COLUMN id SET DEFAULT gen_random_uuid();

ALTER TABLE follows
    ADD CONSTRAINT follows_follower_id_fkey FOREIGN KEY (follower_id) REFERENCES users(id) ON DELETE CASCADE,
    ADD CONSTRAINT follows_followee_id_fkey FOREIGN KEY (followee_id) REFERENCES users(id) ON DELETE CASCADE;
