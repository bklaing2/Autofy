import { Resource } from 'sst'
import { SendMessageCommand, SQSClient } from '@aws-sdk/client-sqs'
import { DBPlaylist } from '$lib/types'

const client = new SQSClient({})
const SQS_QUEUE_URL = Resource.PlaylistsToUpdate.url

async function updatePlaylist(updates: Omit<DBPlaylist, 'updatedAt'>) {
  const command = new SendMessageCommand({
    QueueUrl: SQS_QUEUE_URL,
    DelaySeconds: 10,
    MessageBody: JSON.stringify({ updates })
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
