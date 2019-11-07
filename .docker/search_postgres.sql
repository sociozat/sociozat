
#### Create View
CREATE VIEW searches AS

  SELECT
    'topic' AS type,
    topics.name as title,
    topics.slug
  FROM topics
  UNION

  SELECT
    'user' as type,
    users.username as title,
    users.slug
  FROM users

  UNION

 SELECT
    'channel' as type,
    channels.name as title,
    channels.slug
  FROM channels

#### INDEXES
CREATE INDEX index_channels_on_body ON channels USING gin(to_tsvector('simple', name));
CREATE INDEX index_topics_on_body ON topics USING gin(to_tsvector('simple', name));
CREATE INDEX index_users_on_name ON users USING gin(to_tsvector('simple', username));