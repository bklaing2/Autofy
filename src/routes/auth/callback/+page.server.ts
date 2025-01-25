import { SPOTIFY_CLIENT_ID, SPOTIFY_CLIENT_SECRET, SPOTIFY_REDIRECT_URI } from "$env/static/private";
import type { PageServerLoad } from './$types'
import { redirect } from "@sveltejs/kit";
import Tokens from "$lib/server/tokens";


export const load: PageServerLoad = async ({ locals, url, cookies }) => {
  const { spotify } = locals
  const authorizationCode = url.searchParams.get('code')

  if (!authorizationCode)
    throw redirect(303, '/')

  // User is logging in
  spotify.setClientId(SPOTIFY_CLIENT_ID)
  spotify.setClientSecret(SPOTIFY_CLIENT_SECRET)
  spotify.setRedirectURI(SPOTIFY_REDIRECT_URI)

  const { body: tokens } = await spotify.authorizationCodeGrant(authorizationCode)
  await Tokens.save(tokens.access_token, tokens.refresh_token, cookies)
  throw redirect(303, '/')
}
