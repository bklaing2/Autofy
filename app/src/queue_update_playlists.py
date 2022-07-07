import os
from rq import Queue
from pymongo import MongoClient

from app.src.update_playlist import update_playlist
from app.src.worker import conn


# Ping server to wake up dynos
os.system("curl https://auto-fy.herokuapp.com/")


# Set up queue
q = Queue(connection=conn)


# Set up Mongo
mongo = MongoClient(os.environ.get('MONGO_URL'))
db = mongo.autofy
playlists_coll = db.playlists



# Add all playlists to worker queues
print('Updating all playlists...')
for obj in playlists_coll.find({ 'token': { '$exists': True }}):
    q.enqueue(update_playlist, obj, job_timeout=1800) # 30 mins

print('Finished updating playlists')