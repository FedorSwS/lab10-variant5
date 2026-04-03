from contextlib import asynccontextmanager
from fastapi import FastAPI, HTTPException
from pydantic import BaseModel
import httpx
from grpc_client import GRPCClient

GO_SERVICE_URL = "http://go-service:8080"
GO_GRPC_URL = "go-service:50051"

class Address(BaseModel):
    city: str
    street: str

class User(BaseModel):
    id: int
    name: str
    email: str
    address: Address

grpc_client = GRPCClient()

@asynccontextmanager
async def lifespan(app: FastAPI):
    yield
    grpc_client.close()

app = FastAPI(title="Python FastAPI Proxy Service", lifespan=lifespan)

@app.post("/api/forward-user", response_model=User)
async def forward_user(user: User):
    async with httpx.AsyncClient(timeout=10.0) as client:
        try:
            response = await client.post(
                f"{GO_SERVICE_URL}/api/users",
                json=user.model_dump()
            )
            response.raise_for_status()
            return User(**response.json())
        except httpx.HTTPError as e:
            raise HTTPException(status_code=502, detail=f"Go service error: {str(e)}")

@app.get("/api/users")
async def get_users():
    async with httpx.AsyncClient(timeout=10.0) as client:
        try:
            response = await client.get(f"{GO_SERVICE_URL}/api/users")
            response.raise_for_status()
            return response.json()
        except httpx.HTTPError as e:
            raise HTTPException(status_code=502, detail=f"Go service error: {str(e)}")

@app.post("/api/grpc-process")
async def grpc_process(payload: str = "hello grpc"):
    return await grpc_client.process_data(payload)
