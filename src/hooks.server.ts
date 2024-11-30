import type { Handle } from '@sveltejs/kit'
import db from '$lib/server/db'
import Spotify from '$lib/server/spotify'
import Tokens from '$lib/server/tokens'
import { usersTable } from '$lib/server/db/schema'


export const handle: Handle = async ({ event, resolve }) => {
  const cookies = event.cookies
  const tokens = Tokens.get(cookies)

  const spotify = await Spotify(tokens.accessToken, tokens.refreshToken, tokens.valid)
  const accessToken = spotify.getAccessToken()
  const refreshToken = spotify.getRefreshToken()
  await Tokens.save(spotify.getAccessToken(), spotify.getRefreshToken(), cookies)

  const signedIn = !!(accessToken && refreshToken)


  if (signedIn) {
    const user = await spotify.getMe()
    await db
      .insert(usersTable)
      .values({ id: user.body.id, accessToken, refreshToken })
      .onConflictDoUpdate({ target: usersTable.id, set: { accessToken, refreshToken } });
  }

  event.locals = { db, spotify, signedIn }

  return resolve(event, {
    filterSerializedResponseHeaders(name) { return name === 'content-range' }
  })
}
