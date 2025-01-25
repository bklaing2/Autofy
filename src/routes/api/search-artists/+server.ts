import type { RequestHandler } from './$types';
import { json } from '@sveltejs/kit'


export const GET: RequestHandler = async function({ url, locals }) {
	const { spotify } = locals

	const query = url.searchParams.get('q')
	if (!query) return json('')

	const artists = await spotify.searchArtists(query)
	return json(artists.body.artists?.items);
}
