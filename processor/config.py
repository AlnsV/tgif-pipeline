from typing import List, Optional

from pydantic import (
    BaseModel,
    BaseSettings,
    Field
)


class EngineConfig(BaseModel):
    indicator: str = "vwap"
    pairs: List[str] = ["BTC-USD", "ETH-USD", "ETH-BTC"]


class RabbitLogin(BaseModel):
    host: Optional[str] = Field("localhost", env="RABBIT_HOST")
    port: Optional[int] = Field(5672, env="RABBIT_PORT")
    username: Optional[str] = Field("guest", env="RABBIT_USER")
    password: Optional[str] = Field("guest", env="RABBIT_PWD")


class Settings(BaseSettings):
    engine: EngineConfig
    rabbit_login: RabbitLogin
    output_exchange: Optional[str] = "engine_output"
    coinbase_ws_url: Optional[str] = "wss://ws-feed.pro.coinbase.com"


cfg = Settings(
    engine=EngineConfig(),
    rabbit_login=RabbitLogin()
)