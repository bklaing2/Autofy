import { sql } from 'drizzle-orm';
import { boolean, char, integer, pgTable, text, timestamp, varchar, uuid } from 'drizzle-orm/pg-core';

export const usersTable = pgTable('users', {
  id: uuid().primaryKey().default(sql`gen_random_uuid()`),
  name: varchar({ length: 255 }).notNull(),
  age: integer().notNull(),
  email: varchar({ length: 255 }).notNull().unique(),
});

export const playlistsTable = pgTable('playlists', {
  id: char({ length: 22 }).primaryKey().notNull(),
  userId: uuid().notNull().references(() => usersTable.id, { onUpdate: 'cascade', onDelete: 'cascade' }),
  artists: text().array().notNull().default(sql`ARRAY[]::text[]`),
  followedArtists: text().array().notNull().default(sql`ARRAY[]::text[]`),
  updateWhenArtistPosts: boolean().notNull().default(false),
  updateWhenUserFollowsArtist: boolean().notNull().default(false),
  updateWhenUserUnfollowsArtist: boolean().notNull().default(false),
  updatedAt: timestamp().notNull().defaultNow(),
});

export const tokensTable = pgTable('tokens', {
  userId: uuid().primaryKey().notNull().references(() => usersTable.id, { onUpdate: 'cascade', onDelete: 'cascade' }),
  refreshToken: text().notNull(),
  accessToken: text().notNull(),
});

