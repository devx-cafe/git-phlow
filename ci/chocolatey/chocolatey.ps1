
ls
$version = Get-Content .\gp-version\version

Expand-Archive .\phlow-artifact-windows-s3\git-phlow-*.zip -DestinationPath .\phlow-artifact-windows-s3

ls .\phlow-artifact-windows-s3

mv phlow-artifact-windows-s3/git-phlow.exe git-phlow/ci/chocolatey/tools/

cd git-phlow/ci/chocolatey
ls tools/

$NuSpecFilePath = ".\git-phlow.nuspec"
$file = Get-Item -Path $NuSpecFilePath
[ xml ]$fileContents = Get-Content -Path $NuSpecFilePath
$fileContents.package.metadata.version = '{0}' -f $version
$fileContents.Save($file.Fullname)

$token = [Environment]::GetEnvironmentVariable("chocotoken","User")

choco pack

choco push --api-key $token

exit $lastexitcode