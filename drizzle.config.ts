import dotenv from 'dotenv';
import dotenvExpand from 'dotenv-expand';
import { defineConfig } from 'drizzle-kit';
import { Resource } from 'sst';

dotenvExpand.expand(dotenv.config())

export default defineConfig({
  out: './drizzle',
  schema: './src/lib/server/db/schema.ts',
  dialect: 'postgresql',
  casing: 'snake_case',
  dbCredentials: {
    host: Resource.Db.host,
    port: Resource.Db.port,
    user: Resource.Db.username,
    password: Resource.Db.password,
    database: Resource.Db.database,
    ssl: false
  }
});

