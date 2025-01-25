/// <reference path="./.sst/platform/config.d.ts" />


const DB_CONFIG = {
  username: "postgres",
  password: "password",
  host: "localhost",
  port: 5432,
  database: "local"
}


export default $config({

  app(input) {
    return {
      name: "autofy",
      removal: input?.stage === "production" ? "retain" : "remove",
      protect: ["production"].includes(input?.stage),
      home: "aws",
    };
  },

  async run() {
    const vpc = new sst.aws.Vpc("Vpc", { bastion: true, nat: "ec2" });
    const db = new sst.aws.Postgres("Db", {
      vpc,
      proxy: true,
      dev: DB_CONFIG
    });

    new sst.x.DevCommand("DbLocal", {
      dev: {
        command: `docker run --name autofy-db --rm -p ${DB_CONFIG.port}:5432 -v autofy-db:/var/lib/postgresql/data -e POSTGRES_USER=${DB_CONFIG.username} -e POSTGRES_PASSWORD=${DB_CONFIG.password} -e POSTGRES_DB=${DB_CONFIG.database} postgres:17.2-alpine`,
      },
    });

    new sst.x.DevCommand("DrizzleStudio", {
      link: [db],
      dev: {
        command: "bunx drizzle-kit studio",
      },
    });

    const playlistsToUpdateQueue = new sst.aws.Queue("PlaylistsToUpdate");
    const tracksToUpdateQueue = new sst.aws.Queue("TracksToUpdate");

    new sst.aws.SvelteKit("Frontend", {
      link: [db, playlistsToUpdateQueue],
      dev: {
        command: "bun dev"
      }
    });


    const queuePlaylistsFunction = new sst.aws.Function("QueuePlaylists", {
      runtime: "go",
      handler: "./lambdas/queue_playlists",
      link: [db, playlistsToUpdateQueue],
    })

    new sst.aws.Cron("QueuePlaylistsCron", {
      function: queuePlaylistsFunction.arn,
      schedule: "rate(7 days)"
    })

    const updatePlaylistFunction = new sst.aws.Function("UpdatePlaylist", {
      runtime: "go",
      handler: "./lambdas/update_playlist",
      link: [db, tracksToUpdateQueue],
    });

    const updateTracksFunction = new sst.aws.Function("UpdateTracks", {
      runtime: "go",
      handler: "./lambdas/update_tracks",
      link: [db],
    });

    playlistsToUpdateQueue.subscribe(updatePlaylistFunction.arn)
    tracksToUpdateQueue.subscribe(updateTracksFunction.arn)
  }
});
