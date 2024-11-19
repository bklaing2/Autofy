import { DATABASE_URL } from "$env/static/private";
import { drizzle } from 'drizzle-orm/node-postgres';

const db = drizzle({ connection: DATABASE_URL, casing: 'snake_case' });

export default db
export type Database = typeof db
