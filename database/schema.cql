-- Refactored users table
CREATE TABLE IF NOT EXISTS users (
    id bigint PRIMARY KEY,
    username text,
    alias_name text,
    email text,
    created_at timestamp
);

-- Refactored user_preferences table
CREATE TABLE IF NOT EXISTS user_preferences (
    user_id bigint,
    preference_id bigint,
    value text,
    PRIMARY KEY (user_id, preference_id)
);

CREATE CUSTOM INDEX user_preferences_value_idx ON messaging.user_preferences (value)
  USING 'StorageAttachedIndex';

-- Refactored channels table
CREATE TABLE IF NOT EXISTS channels (
    id bigint PRIMARY KEY,
    name text,
    display_name text,
    description text,
    creator_id bigint,
    is_private boolean,
    created_at timestamp,
    updated_at timestamp
);

CREATE CUSTOM INDEX channels_creator_id_idx ON messaging.channels (creator_id)
  USING 'StorageAttachedIndex';

-- Refactored messages table
CREATE TABLE IF NOT EXISTS messages (
    channel_id bigint,
    bucket int,
    message_id bigint,
    author_id bigint,
    content text,
    created_at timestamp,
    PRIMARY KEY ((channel_id, bucket), message_id)
) WITH CLUSTERING ORDER BY (message_id DESC);
