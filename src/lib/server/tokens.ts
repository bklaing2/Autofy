import { ENV } from "$env/static/private"
import type { Cookies } from "@sveltejs/kit"
import Supabase from "./supabase"

type OptionalString = string | null | undefined

const secure = ENV! !== 'DEV'


function get(cookies: Cookies) {
  return {
    accessToken: cookies.get('spotify_access_token'),
    refreshToken: cookies.get('spotify_refresh_token'),
    valid: !!cookies.get('spotify_token_valid')
  }
}


async function save(accessToken: OptionalString, refreshToken: OptionalString, cookies: Cookies) {
  cookies.set('spotify_access_token', accessToken ?? '', { path: '/', secure: secure })
  cookies.set('spotify_token_valid', new Date().toString(), { path: '/', secure: secure, maxAge: 3600 })
  if (!refreshToken) return

  cookies.set('spotify_refresh_token', refreshToken ?? '', { path: '/', secure: secure })
  if (!accessToken) return

  const supabase = await Supabase(cookies)
  await supabase
    .from('tokens')
    .upsert({ access_token: accessToken, refresh_token: refreshToken })
}


function clear(cookies: Cookies) {
  cookies.delete('spotify_access_token', { path: '/', secure: secure })
  cookies.delete('spotify_refresh_token', { path: '/', secure: secure })
  cookies.delete('spotify_token_valid', { path: '/', secure: secure })
}


const Tokens = { get, save, clear }

export default Tokens
