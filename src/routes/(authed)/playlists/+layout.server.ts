import type { LayoutServerLoad } from './$types'
import { and, eq } from 'drizzle-orm'
import { playlistsTable } from '$lib/server/db/schema'


export const load: LayoutServerLoad = async ({ locals }) => {
  const { db, spotify, signedIn } = locals

  if (!signedIn) return { playlists: [] as string[] }
  const user = (await spotify.getMe()).body.id

  const data = await db
    .select({ id: playlistsTable.id })
    .from(playlistsTable)
    .where(eq(playlistsTable.userId, user))

  let playlists = data.map(p => p.id)
  if (playlists.length === 0) return { playlists }

  // const exists = await spotify.getUserPlaylists()
  // for (let i = playlists.length - 1; i >= 0; i--) {
  //   const playlist = playlists[i]
  //   if (exists.body.items.find(p => p.id === playlist)) continue
  //
  //   db
  //     .delete(playlistsTable)
  //     .where(and(
  //       eq(playlistsTable.userId, user),
  //       eq(playlistsTable.id, playlist)
  //     ))
  //
  //   playlists = [...playlists.slice(0, i), ...playlists.slice(i + 1)]
  // }

  return { playlists: playlists }
}
