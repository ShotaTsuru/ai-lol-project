import os
import requests
from fast_mcp import MCPServer, Tool, Param
from dotenv import load_dotenv

load_dotenv()

RIOT_API_KEY = os.getenv("RIOT_API_KEY", "")
BASE_URL = "https://kr.api.riotgames.com"
DDRAGON_URL = "https://ddragon.leagueoflegends.com"

headers = {"X-Riot-Token": RIOT_API_KEY}

server = MCPServer(
    name="League of Legends MCP Server",
    description="League of Legends data and analysis MCP server",
    version="1.0.0",
)


@server.tool(
    name="get_summoner_info",
    description="サモナーの基本情報を取得",
    params=[
        Param("summoner_name", str, "サモナー名"),
        Param("region", str, "リージョン（kr, na1, euw1など）", default="kr"),
    ],
)
def get_summoner_info(summoner_name, region="kr"):
    url = f"https://{region}.api.riotgames.com/lol/summoner/v4/summoners/by-name/{summoner_name}"
    resp = requests.get(url, headers=headers)
    if resp.status_code != 200:
        return {"error": f"サモナー情報の取得に失敗しました: {resp.text}"}
    data = resp.json()
    return {
        "summonerId": data.get("id"),
        "accountId": data.get("accountId"),
        "puuid": data.get("puuid"),
        "name": data.get("name"),
        "profileIconId": data.get("profileIconId"),
        "level": data.get("summonerLevel"),
        "revisionDate": data.get("revisionDate"),
    }


if __name__ == "__main__":
    if not RIOT_API_KEY:
        print("❌ RIOT_API_KEY is not set in .env")
        exit(1)
    print("🎮 League of Legends MCP Server starting...")
    server.run(host="0.0.0.0", port=8080)
