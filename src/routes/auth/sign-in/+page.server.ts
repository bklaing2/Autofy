import { SPOTIFY_CLIENT_ID, SPOTIFY_REDIRECT_URI } from "$env/static/private";
import type { Actions } from './$types';
import { redirect } from '@sveltejs/kit'


const SCOPES = [
  'user-follow-read',
  'playlist-read-private',
  'playlist-modify-public',
  'playlist-modify-private',
  'ugc-image-upload'
]


export const actions = {
  default: async ({ locals }) => {
    const { spotify } = locals
    spotify.setClientId(SPOTIFY_CLIENT_ID)
    spotify.setRedirectURI(SPOTIFY_REDIRECT_URI)

    const authorizeUrl = spotify.createAuthorizeURL(SCOPES, 'login')
    throw redirect(303, authorizeUrl)
  }
} satisfies Actions
