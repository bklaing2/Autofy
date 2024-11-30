import type { InferSelectModel } from "drizzle-orm";
import type { playlistsTable } from "./server/db/schema";


export type DBPlaylist = InferSelectModel<typeof playlistsTable> { }

export type Artist = { id: string; name: string; img?: string };

export interface Playlist extends Omit<DBPlaylist, 'artists' | 'followedArtists'> {
  artists: Artist[];
  followedArtists: boolean;
}


