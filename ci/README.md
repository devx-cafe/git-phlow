# Our Concourse pipelines

* Infrastructure project, Bosh + Concourse etc. as code here: https://github.com/Praqma/praqma-concourse
* The pipeline itself here where you are, in `pipeline.yml`
* `pipeline.yml` uses a lot of secrets, they need to be used when uploading the pipeline. We have a `git-phlow-cred.yml` available so you can easily load the pipeline with that:

    fly -t praqma-concourse-ci set-pipeline --pipeline git-phlow --config pipeline.yml --load-vars-from git-phlow-cred.yml


The pipeline name should always be `git-phlow` as we reference the URL to this in docs etc.
