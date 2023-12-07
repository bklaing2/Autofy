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
    const { data, error } = await locals.supabase.auth.signInWithOAuth({
      provider: 'spotify',
      options: { scopes: SCOPES.join(' ') }
    })

    
    if (error) throw new Error(error.message)
    throw redirect(303, data.url as string)
  }
}