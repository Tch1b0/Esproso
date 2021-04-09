from flask import Flask, jsonify, request
from http import HTTPStatus

app = Flask(__name__)

@app.route("/")
def snake():
    data = {
        "apiversion": 1,
        "author": "Tch1b0",
        "color": "#400080",
        "head": "caffeine",
        "tail": "coffee",
        "version": "0.0.1"
    }
    return jsonify(data), HTTPStatus.OK

@app.route("/start")
def start():
    return jsonify({"status": "ok"}), HTTPStatus.OK

@app.route("/move")
def move():
    return jsonify({"move": "up"}), HTTPStatus.OK

@app.route("/end")
def end():
    return jsonify({"status": "ok"}), HTTPStatus.OK

if __name__ == "__main__":
    app.run(host="0.0.0.0", port=5001, debug=True)