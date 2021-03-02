import os

from dotenv import load_dotenv

load_dotenv()

ADMIN_PASSWORD = os.environ.get("ADMIN_PASSWORD", "CCNSccns")
