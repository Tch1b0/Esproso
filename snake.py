from flask import Flask, jsonify, request
from http import HTTPStatus

app = Flask(__name__)

@app.route("/")
def snake():
    data = {
        "apiversion": 1,
        "author": "Tch1b0",
        "color": "#888888",
        "head": "caffeine",
        "tail": "coffee",
        "version": "0.0.1"
    }
    return jsonify(data)

@app.route("/start")
def start():
    return jsonify({"status": "ok"}), HTTPStatus.OK

@app.route("move")
def move():
    return jsonify({"move": "up"})

@app.route("end")
def end():
    return jsonify({"status": "ok"}), HTTPStatus.OK

app.run("localhost", 3000)