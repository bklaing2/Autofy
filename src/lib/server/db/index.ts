import { Resource } from 'sst';
import { Pool } from 'pg';
import { drizzle } from 'drizzle-orm/node-postgres';
import * as schema from './schema';

const pool = new Pool({
  host: Resource.AutofyDb.host,
  port: Resource.AutofyDb.port,
  user: Resource.AutofyDb.username,
  password: Resource.AutofyDb.password,
  database: Resource.AutofyDb.database,
})

const db = drizzle(pool, { schema, casing: 'snake_case' });

export default db
export type Database = typeof db
