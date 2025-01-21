import { Resource } from 'sst'
import { SendMessageCommand, SQSClient } from '@aws-sdk/client-sqs'
import { Playlist } from '$lib/types'

const client = new SQSClient({})
const SQS_QUEUE_URL = Resource.UpdatePlaylists.url

async function updatePlaylist(playlistId: Playlist['id']) {
  const command = new SendMessageCommand({
    QueueUrl: SQS_QUEUE_URL,
    DelaySeconds: 10,
    MessageBody: playlistId
  })

  const response = await client.send(command)
  console.log(response)
  return response
}

const queue = {
  updatePlaylist
}

export default queue
export type Queue = typeof queue
