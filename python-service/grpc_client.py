import grpc

class GRPCClient:
    def __init__(self, host: str = "localhost", port: int = 50051):
        self.host = host
        self.port = port
        self.channel = None
        self.stub = None

    def connect(self):
        try:
            self.channel = grpc.insecure_channel(f"{self.host}:{self.port}")
            print(f"Connected to gRPC server at {self.host}:{self.port}")
            return True
        except Exception as e:
            print(f"Failed to connect to gRPC server: {e}")
            return False

    def close(self):
        if self.channel:
            self.channel.close()

    async def process_data(self, payload: str) -> dict:
        return {
            "grpc_result": f"Mock processed: {payload}",
            "status": 200,
            "note": "gRPC stub not generated - see README for instructions"
        }
