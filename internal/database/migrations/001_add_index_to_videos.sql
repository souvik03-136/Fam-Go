
-- Add an index on published_at for optimization
CREATE INDEX idx_videos_published_at ON videos (published_at);

-- Optional: Add an index on title if needed
CREATE INDEX idx_videos_title ON videos (title);
