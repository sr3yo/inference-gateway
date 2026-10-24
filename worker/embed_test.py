from sentence_transformers import SentenceTransformer
model = SentenceTransformer("all-MiniLM-L6-v2")

vec = model.encode("hello")

print(len(vec))
print(vec[:5])
