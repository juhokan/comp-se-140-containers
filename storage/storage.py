from flask import Flask, request, Response
import os

app = Flask(__name__)
LOG_FILE = '/data/log.txt'

os.makedirs('/data', exist_ok=True)

@app.route('/log', methods=['POST'])
def append_log():
    record = request.data.decode('utf-8')
    with open(LOG_FILE, 'a') as f:
        f.write(record + '\n')
    return '', 204

@app.route('/log', methods=['GET'])
def get_log():
    if not os.path.exists(LOG_FILE):
        return Response('', mimetype='text/plain')
    with open(LOG_FILE, 'r') as f:
        content = f.read()
    return Response(content, mimetype='text/plain')

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000)
