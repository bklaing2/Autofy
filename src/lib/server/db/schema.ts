import { relations, sql } from 'drizzle-orm';
import { boolean, char, pgTable, text, timestamp } from 'drizzle-orm/pg-core';

export const usersTable = pgTable('users', {
  id: text().primaryKey().notNull(),
  refreshToken: text().notNull(),
  accessToken: text().notNull(),
});

export const usersRelations = relations(usersTable, ({ many }) => ({
  playlists: many(playlistsTable),
}));

export const playlistsTable = pgTable('playlists', {
  id: char({ length: 22 }).primaryKey().notNull(),
  userId: text().notNull().references(() => usersTable.id, { onDelete: 'cascade' }),
  artists: text().array().notNull().default(sql`ARRAY[]::text[]`),
  followedArtists: text().array().notNull().default(sql`ARRAY[]::text[]`),
  updateWhenArtistPosts: boolean().notNull().default(false),
  updateWhenUserFollowsArtist: boolean().notNull().default(false),
  updateWhenUserUnfollowsArtist: boolean().notNull().default(false),
  updatedAt: timestamp().notNull().defaultNow(),
});

export const playlistsRelations = relations(playlistsTable, ({ one }) => ({
  user: one(usersTable, { fields: [playlistsTable.userId], references: [usersTable.id] }),
}));
