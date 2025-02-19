from fastapi import FastAPI
import uvicorn
import base64
from email.message import EmailMessage
from pydantic import BaseModel

import google.auth
from googleapiclient.discovery import build
from google.oauth2.credentials import Credentials
from googleapiclient.errors import HttpError

app = FastAPI()
class EmailRequest(BaseModel):
    email: str
    code: str


@app.post("/sendCode")
def send_code(request: EmailRequest):
    email = request.email
    code = request.code
    print("Sending code to email..." + email + " with code: " + code) 
    SCOPES = ["https://www.googleapis.com/auth/gmail.send"]
    creds = Credentials.from_authorized_user_file("token.json", SCOPES)
    
    try:
        service = build("gmail", "v1", credentials=creds)
        message = EmailMessage()
        
        message.set_content("Verification code: " + code)
        message["Subject"] = "Travel Point Verification code"
        message["From"] = "davi31102005@gmail.com"
        message["To"] = email
        
        enconded_message = base64.urlsafe_b64encode(message.as_bytes())
        
        create_message = {"raw": enconded_message.decode()}
        
        send_message = ( service.users().messages().send(userId="me", body=create_message).execute() )
        print(f'Message Id: {send_message["id"]}')
    except HttpError as error:
        print(f"An error occurred: {error}")
        send_message = None
        return send_message

if __name__ == "__main__":
    uvicorn.run(app, host="127.0.0.1", port=5000)
        