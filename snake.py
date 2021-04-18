from flask import Flask, jsonify
from http import HTTPStatus

app = Flask(__name__)

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
    return jsonify({"status": "ok"}), HTTPStatus.OK

@app.route("/move", methods=["POST"])
def move():
    return jsonify({"move": "up"}), HTTPStatus.OK

@app.route("/end", methods=["POST"])
def end():
    return jsonify({"status": "ok"}), HTTPStatus.OK

if __name__ == "__main__":
    app.run(host="0.0.0.0", debug=True, threaded=True)