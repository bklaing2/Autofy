import 'unplugin-icons/types/svelte'
import type { Database } from '$lib/server/db';
import type SpotifyWebApi from 'spotify-web-api-node'

// See https://kit.svelte.dev/docs/types#app
// for information about these interfaces
declare global {
  namespace App {
    // interface Error {}
    interface Locals {
      db: Database;
      spotify: SpotifyWebApi;
      signedIn?: boolean;
    }
    // interface PageData {}
    // interface Platform {}
  }
}

export { };
