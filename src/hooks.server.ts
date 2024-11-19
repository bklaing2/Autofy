import type { Handle } from '@sveltejs/kit'
import db from '$lib/server/db'
import Spotify from '$lib/server/spotify'
import Tokens from '$lib/server/tokens'


export const handle: Handle = async ({ event, resolve }) => {
  const cookies = event.cookies
  const tokens = Tokens.get(cookies)

  const spotify = await Spotify(tokens.accessToken, tokens.refreshToken, tokens.valid)
  const accessToken = spotify.getAccessToken(), refreshToken = spotify.getRefreshToken()
  const signedIn = !!(accessToken && refreshToken)

  await Tokens.save(spotify.getAccessToken(), spotify.getRefreshToken(), cookies)

  event.locals = { db, spotify, signedIn }

  return resolve(event, {
    filterSerializedResponseHeaders(name) { return name === 'content-range' }
  })
}
