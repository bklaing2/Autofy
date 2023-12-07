import 'unplugin-icons/types/svelte'
import type { SupabaseClient } from "@supabase/supabase-js";
import SpotifyWebApi from 'spotify-web-api-node'
import type { Database } from '$lib/types'

// See https://kit.svelte.dev/docs/types#app
// for information about these interfaces
declare global {
	namespace App {
		// interface Error {}
		interface Locals {
			supabase: SupabaseClient<Database>;
			spotify: SpotifyWebApi;
			signedIn?: boolean;
		}
		// interface PageData {}
		// interface Platform {}
	}
}

export {};
