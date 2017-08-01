import hashlib
import json
import subprocess


version = open("gp-version/version","r").read().replace("\n","")

def clone():
    output = subprocess.check_output(["git","clone","scoop-bucket","scoop-bucket-modified"])

def scoop_release():
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

    file = open("scoop-bucket-modified/git-phlow.json","w")
    file.write(json.dumps(d, indent=2))


def commit():
    scm = "git"
    folder = "scoop-bucket-modified"
    output = subprocess.check_output([scm,"config","--global","user.email","concourse@praqma.net"])
    print output
    output = subprocess.check_output([scm,"config","--global","user.name","concourse"])
    print output
    
    output = subprocess.check_output([scm,"-C",folder,"add","--all"])
    print output
    output = subprocess.check_output([scm,"-C",folder,"status"])
    print output
    output = subprocess.check_output([scm,"-C",folder,"commit","-m","released " + version])
    print output

clone()
scoop_release()
commit()
    