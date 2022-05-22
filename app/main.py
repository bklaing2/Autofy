import os
from urllib.parse import urlparse
from flask import Flask
import redis
import pymongo


# Set up server
app = Flask(__name__)

# Set up Redis
url = urlparse(os.environ.get("REDIS_TLS_URL"))
r = redis.Redis(host=url.hostname, port=url.port, username=url.username, password=url.password, ssl=True, ssl_cert_reqs=None)

# Set up Mongo
mongo = pymongo.MongoClient(os.environ.get('MONGO_URL'))
db = mongo.autofy
playlists = db.playlists


redis_test = r.get('test')
mongo_test = playlists.find_one()

@app.route('/')
def index():
    return f'REDIS: {redis_test}<br>MONGO: {mongo_test}'
