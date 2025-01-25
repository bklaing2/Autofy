import { SPOTIFY_CLIENT_ID, SPOTIFY_CLIENT_SECRET, SPOTIFY_REDIRECT_URI } from "$env/static/private";
import SpotifyWebApi from "spotify-web-api-node";


export default async function Spotify(accessToken?: string, refreshToken?: string, valid?: boolean) {
  const spotify = new SpotifyWebApi({ accessToken, refreshToken })

  // User is logged in with a valid access token
  if (accessToken && valid) return spotify

  // User is not logged in
  if (!accessToken || !refreshToken) return spotify

  // Refresh access token
  spotify.setClientId(SPOTIFY_CLIENT_ID)
  spotify.setClientSecret(SPOTIFY_CLIENT_SECRET)
  spotify.setRedirectURI(SPOTIFY_REDIRECT_URI)

  const { body: tokens } = await spotify.refreshAccessToken()
  spotify.setAccessToken(tokens.access_token)
  spotify.setRefreshToken(tokens.refresh_token || '')

  return spotify
}
