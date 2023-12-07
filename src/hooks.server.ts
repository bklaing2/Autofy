import type { Handle } from '@sveltejs/kit'
import Supabase from '$lib/server/supabase'
import Spotify from '$lib/server/spotify'


export const handle: Handle = async ({ event, resolve }) => {
  const cookies = event.cookies

  const supabase = await Supabase(cookies, event.url.searchParams.get('code'))

  event.locals = {
    supabase: supabase,
    spotify: await Spotify(cookies),
    signedIn: !!(await supabase.auth.getSession()).data.session?.user
  }

  return resolve(event, {
    filterSerializedResponseHeaders(name) { return name === 'content-range' }
  })
}