## concourse task for releasing a package on chocolatey
$env:Path = [System.Environment]::GetEnvironmentVariable("Path","Machine")
$ErrorActionPreference = "Stop"

ls   #debugging

# read version from file
$version = Get-Content .\gp-version\version

# list artifact for debugging purpose
ls .\phlow-artifact-windows-s3

mv phlow-artifact-windows-s3/git-phlow.exe git-phlow/ci/chocolatey/tools/

cd git-phlow/ci/chocolatey
ls tools/

# update the nuspec file
$NuSpecFilePath = ".\git-phlow.nuspec"
$file = Get-Item -Path $NuSpecFilePath
[ xml ]$fileContents = Get-Content -Path $NuSpecFilePath
$fileContents.package.metadata.version = '{0}' -f $version
$fileContents.Save($file.Fullname)

$token = $env:chocotoken

choco pack

choco push --api-key $token

exit $LastExitCode