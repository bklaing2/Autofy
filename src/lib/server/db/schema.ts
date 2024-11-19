import { sql } from 'drizzle-orm';
import { boolean, char, pgTable, text, timestamp } from 'drizzle-orm/pg-core';

export const playlistsTable = pgTable('playlists', {
  id: char({ length: 22 }).primaryKey().notNull(),
  userId: text().notNull(),
  artists: text().array().notNull().default(sql`ARRAY[]::text[]`),
  followedArtists: text().array().notNull().default(sql`ARRAY[]::text[]`),
  updateWhenArtistPosts: boolean().notNull().default(false),
  updateWhenUserFollowsArtist: boolean().notNull().default(false),
  updateWhenUserUnfollowsArtist: boolean().notNull().default(false),
  updatedAt: timestamp().notNull().defaultNow(),
});

export const tokensTable = pgTable('tokens', {
  userId: text().primaryKey().notNull(),
  refreshToken: text().notNull(),
  accessToken: text().notNull(),
});

