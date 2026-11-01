from concurrent import futures

import grpc
from sentence_transformers import SentenceTransformer

import embedding_pb2
import embedding_pb2_grpc


class EmbeddingService(embedding_pb2_grpc.EmbeddingServiceServicer):
    def __init__(self):
        self.model = SentenceTransformer("all-MiniLM-L6-v2")

    def Embed(self, request, context):
        vectors = self.model.encode(list(request.texts))

        embeddings = [
            embedding_pb2.Embedding(values=vector.tolist())
            for vector in vectors
        ]

        return embedding_pb2.EmbedResponse(embeddings=embeddings)


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=4))

    embedding_pb2_grpc.add_EmbeddingServiceServicer_to_server(EmbeddingService(), server,)

    server.add_insecure_port("127.0.0.1:50051")
    server.start()
    print("gRPC worker listening on 127.0.0.1:50051", flush=True)
    server.wait_for_termination()


if __name__ == "__main__":
    serve()