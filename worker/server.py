from flask import Flask, request, jsonify
from sentence_transformers import SentenceTransformer

model = SentenceTransformer("all-MiniLM-L6-v2")

app = Flask(__name__)
@app.route("/embed", methods=["POST"])

def embed():
    data = request.get_json()
    texts = data["texts"]
    vectors = model.encode(texts)
    return jsonify({"embeddings": vectors.tolist()})

if __name__ == "__main__":
    app.run(port=8000)