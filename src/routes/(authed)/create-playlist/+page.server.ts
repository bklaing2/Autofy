import { redirect } from '@sveltejs/kit'


export const actions = {
  default: async ({ request, fetch }) => {
    await fetch('/playlists', { method: 'POST', body: await request.formData() })
    throw redirect(303, '/playlists')
  }
}