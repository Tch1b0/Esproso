from flask import Flask, jsonify, request
from http import HTTPStatus
from brain import Brain

app = Flask(__name__)
activeGames = {}


@app.after_request
def setJson(response):
    response.headers.add("Content-Type", "application/json")
    return response


@app.route("/")
def snake():
    data = {
        "apiversion": "1",
        "author": "Tch1b0",
        "color": "#400080",
        "head": "caffeine",
        "tail": "coffee",
        "version": "0.0.1"
    }
    response = app.make_response(data)
    response.mimetype = "application/json"
    return jsonify(data), HTTPStatus.OK


@app.route("/start", methods=["POST"])
def start():
    b = Brain()
    b.id = request.form["game"]["id"]
    b.board = request.form["board"]
    b.turn = 0
    activeGames[b.id] = b
    return jsonify({"status": "ok"}), HTTPStatus.OK


@app.route("/move", methods=["POST"])
def move():
    b = activeGames[request.form["game"]["id"]]
    b.turn = request.form["turn"]
    b.snake = request.form["you"]
    move = b.move()[0]
    b.lastMove = move
    return jsonify({"move": move}), HTTPStatus.OK


@app.route("/end", methods=["POST"])
def end():
    del activeGames[request.form["game"]["id"]]
    return jsonify({"status": "ok"}), HTTPStatus.OK


if __name__ == "__main__":
    app.run(host="0.0.0.0", debug=True, threaded=True)
