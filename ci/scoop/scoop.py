import hashlib
import json

version = open("gp-version/version","r").read().replace("\n","")
hash = hashlib.sha256(open("phlow-artifact-windows-s3/git-phlow-"+ version +"-windows-amd64.zip","r").read()).hexdigest()


d = {
    "version":  version,
    "license":  "mit",
    "extract_dir":  "",
    "url": "https://github.com/Praqma/git-phlow/releases/download/v"+ version +"/git-phlow-"+version+"-windows-amd64.zip",
    "depends": "git",
    "homepage": "https://github.com/praqma/git-phlow",
    "hash":  hash,
    "bin":  "git-phlow.exe",
    "notes": "type 'git phlow' for help" ,
    "autoupdate": {
        "url": "https://github.com/Praqma/git-phlow/releases/download/v"+version+"/git-phlow-"+version+"-windows-amd64.zip"
    }
}


file = open("scoop-bucket/git-phlow.json","w")
file.write(json.dumps(d, indent=2))