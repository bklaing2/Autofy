import { redirect } from "@sveltejs/kit";

export function load({ locals }) {
	if (!locals.signedIn) throw redirect(303, `/`)
}