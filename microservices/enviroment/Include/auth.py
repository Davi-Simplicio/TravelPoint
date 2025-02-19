import os
import json
from google.oauth2.credentials import Credentials
from google_auth_oauthlib.flow import InstalledAppFlow
from google.auth.transport.requests import Request
from dotenv import load_dotenv

load_dotenv()

CLIENT_SECRET_FILE = os.getenv("CLIENT_SECRET_FILE")
TOKEN_FILE = os.getenv("TOKEN_FILE")
SCOPES = os.getenv("SCOPES").split(" ")

def get_credentials():
    creds = None

    # ✅ Check if token.json exists and load it
    if os.path.exists(TOKEN_FILE):
        with open(TOKEN_FILE, "r") as token:
            creds = Credentials.from_authorized_user_info(json.load(token), SCOPES)

    # ✅ If token is missing or expired, authenticate again
    if not creds or not creds.valid:
        if creds and creds.expired and creds.refresh_token:
            creds.refresh(Request())  # Refresh expired token
        else:
            # ✅ Open browser for user login with `access_type="offline"`
            flow = InstalledAppFlow.from_client_secrets_file(
                CLIENT_SECRET_FILE,
                SCOPES
            )
            creds = flow.run_local_server(port=8080, access_type="offline", prompt="consent")

            # ✅ Save the token to token.json
            with open(TOKEN_FILE, "w") as token:
                token.write(creds.to_json())
                print("✅ Token successfully saved to token.json!")

    return creds

# ✅ Run authentication
creds = get_credentials()
print("✅ Authentication completed, token.json is now saved!")
