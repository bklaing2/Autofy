import type { Cookies } from "@sveltejs/kit";
import Tokens from "$lib/server/tokens";
import SpotifyWebApi from "spotify-web-api-node";
import Supabase from "$lib/server/supabase";

async function Spotify (cookies: Cookies) {
  const { accessToken, refreshToken, valid } = Tokens.get(cookies)
  const spotify = new SpotifyWebApi({ accessToken: accessToken })
  if (accessToken && valid) return spotify

  if (!accessToken || !refreshToken) {
    console.log('Missing access token or refresh token, signing the user out')
    const supabase = await Supabase(cookies)
    await supabase.auth.signOut()
    Tokens.clear(cookies)
    return spotify
  }

  console.log('Refreshing token from Spotify')
  spotify.setClientId(process.env.SPOTIFY_CLIENT_ID!)
  spotify.setClientSecret(process.env.SPOTIFY_CLIENT_SECRET!)
  spotify.setRefreshToken(refreshToken)

  const { body: tokens } = await spotify.refreshAccessToken()
  await Tokens.save(tokens.access_token, tokens.refresh_token, cookies)
  return spotify
}

export async function Service (accessToken: string, refreshToken: string) {
  const spotify = new SpotifyWebApi({
    clientId: process.env.SPOTIFY_CLIENT_ID!,
    clientSecret: process.env.SPOTIFY_CLIENT_SECRET!,
    accessToken: accessToken,
    refreshToken: refreshToken
  })

  const data = await spotify.refreshAccessToken()
  return { spotify: spotify, accessToken: data.body.access_token }
}

export default Spotify