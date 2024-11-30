import type { LayoutServerLoad } from './$types'
import { eq } from 'drizzle-orm'
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

  return { playlists }
}
